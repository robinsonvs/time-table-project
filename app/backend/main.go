package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/robinsonvs/time-table-project/config/env"
	"github.com/robinsonvs/time-table-project/config/logger"
	_ "github.com/robinsonvs/time-table-project/docs"
	"github.com/robinsonvs/time-table-project/internal/database"
	"github.com/robinsonvs/time-table-project/internal/database/sqlc"
	"github.com/robinsonvs/time-table-project/internal/handler"
	"github.com/robinsonvs/time-table-project/internal/handler/routes"
	"github.com/robinsonvs/time-table-project/internal/repository/availabilityrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/courserepository"
	"github.com/robinsonvs/time-table-project/internal/repository/disciplinerepository"
	"github.com/robinsonvs/time-table-project/internal/repository/parameterizationrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/professorrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/semesterrepository"
	"github.com/robinsonvs/time-table-project/internal/repository/userrepository"
	"github.com/robinsonvs/time-table-project/internal/service/availabilityservice"
	"github.com/robinsonvs/time-table-project/internal/service/courseservice"
	"github.com/robinsonvs/time-table-project/internal/service/disciplineservice"
	"github.com/robinsonvs/time-table-project/internal/service/parameterizationservice"
	"github.com/robinsonvs/time-table-project/internal/service/professorservice"
	"github.com/robinsonvs/time-table-project/internal/service/semesterservice"
	"github.com/robinsonvs/time-table-project/internal/service/userservice"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

// @title Time Table API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
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

	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	courseRepo := courserepository.NewCourseRepository(dbConnection, queries)
	semesterRepo := semesterrepository.NewSemesterRepository(dbConnection, queries)
	professorRepo := professorrepository.NewProfessorRepository(dbConnection, queries)
	disciplineRepo := disciplinerepository.NewDisciplineRepository(dbConnection, queries)
	availabilityRepo := availabilityrepository.NewAvailabilityRepository(dbConnection, queries)
	parameterizationRepo := parameterizationrepository.NewParameterizationRepository(dbConnection, queries)

	newUserService := userservice.NewUserService(userRepo)
	newCourseService := courseservice.NewCourseService(courseRepo)
	newSemesterService := semesterservice.NewSemesterService(semesterRepo)
	newProfessorService := professorservice.NewProfessorService(professorRepo)
	newDisciplineService := disciplineservice.NewDisciplineService(disciplineRepo)
	newAvailabilityService := availabilityservice.NewAvailabilityService(availabilityRepo)
	newParameterizationService := parameterizationservice.NewParameterizationService(parameterizationRepo)

	newHandler := handler.NewHandler(newUserService,
		newCourseService, newSemesterService, newProfessorService,
		newDisciplineService, newAvailabilityService, newParameterizationService)

	//enableCors(router)

	// init routes
	routes.InitRoutes(router, newHandler)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
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
