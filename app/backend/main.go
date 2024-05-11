package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/robinsonvs/time-table-project/config/env"
	"github.com/robinsonvs/time-table-project/config/logger"
	_ "github.com/robinsonvs/time-table-project/docs"
	"github.com/robinsonvs/time-table-project/internal/database"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/handler/routes"
	"github.com/robinsonvs/time-table-project/internal/handler/userhandler"
	"github.com/robinsonvs/time-table-project/internal/repository/userrepository"
	"github.com/robinsonvs/time-table-project/internal/service/userservice"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

// @title Time Table Allocation API
// @version 1.0

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	//enableCors(router)

	// init routes
	routes.InitUserRoutes(router, newUserHandler)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The url pointing to API definition"
	))

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}

func enableCors(router *chi.Mux) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})
}
