package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/robinsonvs/time-table-project/internal/common/utils"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func LoggerData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestData map[string]interface{}
		if r.Body != http.NoBody {
			// copy body
			CopyBody, _ := io.ReadAll(r.Body)
			// restore body
			r.Body = io.NopCloser(bytes.NewBuffer(CopyBody))
			if err := json.Unmarshal(CopyBody, &requestData); err != nil {
				slog.Error("error unmarshalling request data", err, slog.String("func", "LoggerData"))
			}
			if hasSensitiveData(requestData) {
				for key, value := range requestData {
					for _, keyword := range sensitiveKeywords {
						if containsSensitiveKeyword(key, keyword) || containsSensitiveKeyword(value, keyword) {
							requestData[key] = "[REDACTED]"
						}
					}
				}
			}
		} else {
			r.Body = http.NoBody
		}

		// get user in token
		var userID string
		var userEmail string
		user, err := utils.DecodeJwt(r)
		if err != nil {
			userID = "no token"
			userEmail = "no token"
		} else {
			userID = user.ID
			userEmail = user.Email
		}
		slog.Info("request_data",
			slog.Any("url", r.URL.Path),
			slog.Any("method", r.Method),
			slog.Any("query", r.URL.Query()),
			slog.Any("body", requestData),
			slog.Any("id", userID),
			slog.Any("email", userEmail),
		)

		next.ServeHTTP(w, r)
	})
}

var sensitiveKeywords = []string{"password"}

func hasSensitiveData(body map[string]interface{}) bool {
	for key, value := range body {
		for _, keyword := range sensitiveKeywords {
			if containsSensitiveKeyword(key, keyword) || containsSensitiveKeyword(value, keyword) {
				return true
			}
		}
	}
	return false
}

func containsSensitiveKeyword(data interface{}, keyword string) bool {
	str, ok := data.(string)
	if !ok {
		return false
	}
	return strings.Contains(strings.ToLower(str), keyword)
}
