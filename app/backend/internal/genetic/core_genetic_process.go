package genetic

import (
	"github.com/robinsonvs/time-table-project/internal/entity"
	"time"
)

import (
	"math/rand"
)

func InitializePopulation(size int, disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity) []entity.Timetable {
	population := make([]entity.Timetable, size)
	for i := 0; i < size; i++ {
		population[i] = GenerateRandomTimetable(disciplines, professors, availabilities)
	}
	return population
}

func GenerateRandomTimetable(disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity) entity.Timetable {
	var timetable entity.Timetable
	for _, discipline := range disciplines {
		availableProfessors := FilterEligibleProfessors(discipline.ID, professors)
		if len(availableProfessors) == 0 {
			continue
		}
		professor := availableProfessors[rand.Intn(len(availableProfessors))]
		availableSlots := FilterAvailableSlots(professor.ID, availabilities)
		if len(availableSlots) == 0 {
			continue
		}
		slot := availableSlots[rand.Intn(len(availableSlots))]
		class := entity.ClassEntity{
			DayOfWeek:    slot.DayOfWeek,
			Shift:        slot.Shift,
			StartTime:    time.Now(),                    // Exemplo de horário
			EndTime:      time.Now().Add(1 * time.Hour), // Exemplo de horário
			DisciplineID: discipline.ID,
			ProfessorID:  professor.ID,
		}
		timetable.Classes = append(timetable.Classes, class)
	}
	return timetable
}

func FilterEligibleProfessors(disciplineID int64, professors []entity.ProfessorEntity) []entity.ProfessorEntity {
	// Filtrar os professores elegíveis para a disciplina
	var eligibleProfessors []entity.ProfessorEntity
	for _, professor := range professors {
		for _, eligibleDiscipline := range professor.Disciplines {
			if eligibleDiscipline.ID == disciplineID {
				eligibleProfessors = append(eligibleProfessors, professor)
			}
		}
	}
	return eligibleProfessors
}

func FilterAvailableSlots(professorID int64, availabilities []entity.AvailabilityEntity) []entity.AvailabilityEntity {
	// Filtrar os slots disponíveis para o professor
	var availableSlots []entity.AvailabilityEntity
	for _, availability := range availabilities {
		if availability.ProfessorID == professorID {
			availableSlots = append(availableSlots, availability)
		}
	}
	return availableSlots
}
