package handler

import (
	"github.com/robinsonvs/time-table-project/internal/core/service"
	"github.com/robinsonvs/time-table-project/internal/service/availabilityservice"
	"github.com/robinsonvs/time-table-project/internal/service/courseservice"
	"github.com/robinsonvs/time-table-project/internal/service/disciplineservice"
	"github.com/robinsonvs/time-table-project/internal/service/eligibledisciplineservice"
	"github.com/robinsonvs/time-table-project/internal/service/parameterizationservice"
	"github.com/robinsonvs/time-table-project/internal/service/professorservice"
	"github.com/robinsonvs/time-table-project/internal/service/semesterservice"
	"github.com/robinsonvs/time-table-project/internal/service/userservice"
	"net/http"
)

func NewHandler(userService userservice.UserService,
	courseService courseservice.CourseService,
	semesterService semesterservice.SemesterService,
	professorService professorservice.ProfessorService,
	disciplineService disciplineservice.DisciplineService,
	availabilityService availabilityservice.AvailabilityService,
	parameterizationService parameterizationservice.ParameterizationService,
	eligibleDisciplineService eligibledisciplineservice.EligibleDisciplineService,
	geneticAlgorithmService service.GeneticAlgorithmServiceInterface) Handler {
	return &handler{
		userService:               userService,
		courseService:             courseService,
		semesterService:           semesterService,
		professorService:          professorService,
		disciplineService:         disciplineService,
		availabilityService:       availabilityService,
		parameterizationService:   parameterizationService,
		eligibleDisciplineService: eligibleDisciplineService,
		geneticAlgorithmService:   geneticAlgorithmService,
	}
}

type handler struct {
	userService               userservice.UserService
	courseService             courseservice.CourseService
	semesterService           semesterservice.SemesterService
	professorService          professorservice.ProfessorService
	disciplineService         disciplineservice.DisciplineService
	availabilityService       availabilityservice.AvailabilityService
	parameterizationService   parameterizationservice.ParameterizationService
	eligibleDisciplineService eligibledisciplineservice.EligibleDisciplineService
	geneticAlgorithmService   service.GeneticAlgorithmServiceInterface
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	FindManyUsers(w http.ResponseWriter, r *http.Request)

	UpdateUserPassword(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	CreateCourse(w http.ResponseWriter, r *http.Request)
	UpdateCourse(w http.ResponseWriter, r *http.Request)
	DeleteCourse(w http.ResponseWriter, r *http.Request)
	GetCourseByID(w http.ResponseWriter, r *http.Request)
	FindManyCourses(w http.ResponseWriter, r *http.Request)

	CreateSemester(w http.ResponseWriter, r *http.Request)
	UpdateSemester(w http.ResponseWriter, r *http.Request)
	DeleteSemester(w http.ResponseWriter, r *http.Request)
	GetSemesterByID(w http.ResponseWriter, r *http.Request)
	FindManySemesters(w http.ResponseWriter, r *http.Request)

	CreateProfessor(w http.ResponseWriter, r *http.Request)
	UpdateProfessor(w http.ResponseWriter, r *http.Request)
	DeleteProfessor(w http.ResponseWriter, r *http.Request)
	GetProfessorByID(w http.ResponseWriter, r *http.Request)
	FindManyProfessors(w http.ResponseWriter, r *http.Request)

	CreateDiscipline(w http.ResponseWriter, r *http.Request)
	UpdateDiscipline(w http.ResponseWriter, r *http.Request)
	DeleteDiscipline(w http.ResponseWriter, r *http.Request)
	GetDisciplineByID(w http.ResponseWriter, r *http.Request)
	FindManyDisciplines(w http.ResponseWriter, r *http.Request)
	FindManyDisciplinesByCourseId(w http.ResponseWriter, r *http.Request)

	CreateAvailability(w http.ResponseWriter, r *http.Request)
	UpdateAvailability(w http.ResponseWriter, r *http.Request)
	DeleteAvailability(w http.ResponseWriter, r *http.Request)
	GetAvailabilityByID(w http.ResponseWriter, r *http.Request)
	FindManyAvailabilities(w http.ResponseWriter, r *http.Request)
	FindManyAvailabilitiesByProfessorId(w http.ResponseWriter, r *http.Request)

	CreateParameterization(w http.ResponseWriter, r *http.Request)
	UpdateParameterization(w http.ResponseWriter, r *http.Request)
	DeleteParameterization(w http.ResponseWriter, r *http.Request)
	GetParameterizationByID(w http.ResponseWriter, r *http.Request)
	FindManyParameterizations(w http.ResponseWriter, r *http.Request)
	FindManyParameterizationsBySemesterId(w http.ResponseWriter, r *http.Request)

	CreateEligibleDiscipline(w http.ResponseWriter, r *http.Request)
	DeleteEligibleDiscipline(w http.ResponseWriter, r *http.Request)

	GenerateProposal(w http.ResponseWriter, r *http.Request)
}
