package process

import (
	"log"
	"math/rand"
	"time"

	"github.com/robinsonvs/time-table-project/internal/entity"
)

func InitializePopulation(size int, disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity, weeksToGenerate int, parameterization entity.ParameterizationEntity) []entity.Timetable {
	population := make([]entity.Timetable, size)
	for i := 0; i < size; i++ {
		population[i] = GenerateRandomTimetable(disciplines, professors, availabilities, weeksToGenerate, parameterization)
	}
	return population
}

func GenerateRandomTimetable(disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity, weeksToGenerate int, parameterization entity.ParameterizationEntity) entity.Timetable {
	var timetable entity.Timetable
	occupiedSlots := make(map[string][]time.Time)
	allocatedHours := make(map[int64]float64)
	allocatedCredits := int32(0)
	timeStartProcess := time.Date(2024, 10, 7, 0, 0, 0, 0, time.UTC)

	for week := 0; week < weeksToGenerate; week++ {
		//weekStart := time.Now().AddDate(0, 0, week*7)
		weekStart := timeStartProcess.AddDate(0, 0, week*7)

		for day := 0; day < 5; day++ {
			weekDay := weekStart.AddDate(0, 0, day)

			for _, shift := range []string{"Morning", "Afternoon", "Night"} {

				for _, discipline := range disciplines {

					//if allocatedCredits >= parameterization.MaxCreditsToOffer {
					//	break
					//}

					availableProfessors := FilterEligibleProfessors(discipline.ID, professors)
					if len(availableProfessors) == 0 {
						continue
					}
					professor := availableProfessors[rand.Intn(len(availableProfessors))]
					availableSlots := FilterAvailableSlots(professor.ID, availabilities)
					if len(availableSlots) == 0 {
						continue
					}

					slot := findMatchingSlot(shift, availableSlots, weekDay)
					if slot == nil {
						continue
					}

					startTime, endTime := GenerateNextAvailableTime(weekDay, slot.Shift, occupiedSlots[slot.Shift], slot.DayOfWeek, timetable.Classes)

					zeroTime := time.Time{}
					if startTime.Equal(zeroTime) {
						break
					}

					classDuration := endTime.Sub(startTime).Hours()
					if int(allocatedHours[professor.ID]+classDuration) > int(professor.HoursToAllocate) {
						continue
					}

					class := entity.ClassEntity{
						DayOfWeek:    slot.DayOfWeek,
						Shift:        slot.Shift,
						StartTime:    startTime,
						EndTime:      endTime,
						DisciplineID: discipline.ID,
						ProfessorID:  professor.ID,
					}

					timetable.Classes = append(timetable.Classes, class)
					occupiedSlots[slot.Shift] = append(occupiedSlots[slot.Shift], startTime)
					allocatedHours[professor.ID] += classDuration
					allocatedCredits += discipline.Credits

				}
			}

		}
	}

	return timetable
}

func findMatchingSlot(shift string, availableSlots []entity.AvailabilityEntity, weekDay time.Time) *entity.AvailabilityEntity {
	for i, availableSlot := range availableSlots {
		if availableSlot.DayOfWeek == weekDay.Weekday().String() && availableSlot.Shift == shift {
			return &availableSlots[i]
		}
	}
	return nil
}

func RunGeneticAlgorithm(disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity, parameterization entity.ParameterizationEntity, weeksToGenerate int) entity.Timetable {
	rand.Seed(time.Now().UnixNano())

	populationSize := 100
	population := InitializePopulation(populationSize, disciplines, professors, availabilities, weeksToGenerate, parameterization)

	for i := 0; i < 1000; i++ {
		newPopulation := make([]entity.Timetable, populationSize)
		for j := 0; j < populationSize; j++ {
			parent1 := TournamentSelection(population)
			parent2 := TournamentSelection(population)
			child := Crossover(parent1, parent2)
			Mutate(&child, disciplines, professors, availabilities)
			EvaluateFitness(&child, parameterization)
			newPopulation[j] = child
		}
		ReplacePopulation(population, newPopulation)
	}

	best := population[0]
	for _, individual := range population {
		if individual.Fitness > best.Fitness {
			best = individual
		}
	}

	return best
}

func GenerateNextAvailableTime(weekDay time.Time, shift string, occupiedTimes []time.Time, dayOfWeek string, classes []entity.ClassEntity) (time.Time, time.Time) {
	var startHour, endHour int
	now := weekDay
	switch shift {
	case "Morning":
		startHour = 8
		endHour = 12
	case "Afternoon":
		startHour = 13
		endHour = 18
	case "Night":
		startHour = 19
		endHour = 23
	default:
		log.Fatal("invalid shift")
	}
	for hour := startHour; hour < endHour; hour++ {
		startTime := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
		if !isTimeOccupied(startTime, occupiedTimes) && !isClassScheduled(classes, dayOfWeek, startTime, shift) {
			endTime := startTime.Add(time.Hour)
			return startTime, endTime
		}
	}

	lastSlot := time.Time{}
	return lastSlot, lastSlot.Add(time.Hour)
}

func isTimeOccupied(timeToCheck time.Time, occupiedTimes []time.Time) bool {
	for _, occupiedTime := range occupiedTimes {
		if occupiedTime.Equal(timeToCheck) {
			return true
		}
	}
	return false
}

func isClassScheduled(classes []entity.ClassEntity, dayOfWeek string, startTime time.Time, shift string) bool {
	for _, class := range classes {
		if class.DayOfWeek == dayOfWeek && class.Shift == shift && class.StartTime.Equal(startTime) {
			return true
		}
	}
	return false
}

func FilterEligibleProfessors(disciplineID int64, professors []entity.ProfessorEntity) []entity.ProfessorEntity {
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
	var availableSlots []entity.AvailabilityEntity
	for _, availability := range availabilities {
		if availability.ProfessorID == professorID {
			availableSlots = append(availableSlots, availability)
		}
	}
	return availableSlots
}

func EvaluateFitness(timetable *entity.Timetable, parameterization entity.ParameterizationEntity) {
	fitness := 0.0
	fitness += EvaluateCreditGoals(timetable, parameterization)
	fitness += EvaluateDistribution(timetable)
	fitness += EvaluateNoOverlaps(timetable)
	fitness += EvaluateTeacherHours(timetable, parameterization)
	timetable.Fitness = fitness
}

func EvaluateCreditGoals(timetable *entity.Timetable, parameterization entity.ParameterizationEntity) float64 {
	var creditCount int32 = 0
	for _, class := range timetable.Classes {
		for _, discipline := range parameterization.Disciplines {
			if class.DisciplineID == discipline.ID {
				creditCount += discipline.Credits
			}
		}
	}
	if creditCount <= parameterization.MaxCreditsToOffer {
		return 1.0
	}
	return 0.0
}

func EvaluateDistribution(timetable *entity.Timetable) float64 {
	return 1.0
}

func EvaluateNoOverlaps(timetable *entity.Timetable) float64 {
	for i, class1 := range timetable.Classes {
		for j, class2 := range timetable.Classes {
			if i != j && class1.DayOfWeek == class2.DayOfWeek && class1.Shift == class2.Shift {
				if class1.ProfessorID == class2.ProfessorID || class1.DisciplineID == class2.DisciplineID {
					if class1.StartTime == class2.StartTime { // test if you don't have two classes in a row with the same teacher
						return 0.0
					}
				}
			}
		}
	}
	return 1.0
}

func EvaluateTeacherHours(timetable *entity.Timetable, parameterization entity.ParameterizationEntity) float64 {
	teacherHours := make(map[int64]float64)
	for _, class := range timetable.Classes {
		teacherHours[class.ProfessorID] += class.EndTime.Sub(class.StartTime).Hours()
	}
	for _, professor := range parameterization.Professors {
		if int(teacherHours[professor.ID]) > int(professor.HoursToAllocate) {
			return 0.0 // Strongly penalizes schedules that exceed the teacher's availability
		}
	}
	return 1.0
}

func TournamentSelection(population []entity.Timetable) entity.Timetable {
	tournamentSize := 5
	tournament := make([]entity.Timetable, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		randomIndex := rand.Intn(len(population))
		tournament[i] = population[randomIndex]
	}
	best := tournament[0]
	for _, individual := range tournament {
		if individual.Fitness > best.Fitness {
			best = individual
		}
	}
	return best
}

func Crossover(parent1, parent2 entity.Timetable) entity.Timetable {
	if parent1.Classes == nil || parent2.Classes == nil {
		return entity.Timetable{
			Classes: []entity.ClassEntity{},
		}
	}

	if len(parent1.Classes) == 0 || len(parent2.Classes) == 0 {
		return entity.Timetable{
			Classes: []entity.ClassEntity{},
		}
	}

	// Define the crossover point as the smaller size between the two parents
	minLength := len(parent1.Classes)
	if len(parent2.Classes) < minLength {
		minLength = len(parent2.Classes)
	}

	// Generate a valid crossing point
	crossoverPoint := rand.Intn(minLength)

	// Raising a child with cross-class backgrounds
	child := entity.Timetable{}
	child.Classes = append(child.Classes, parent1.Classes[:crossoverPoint]...)
	child.Classes = append(child.Classes, parent2.Classes[crossoverPoint:]...)

	return child
}

func Mutate(timetable *entity.Timetable, disciplines []entity.DisciplineEntity, professors []entity.ProfessorEntity, availabilities []entity.AvailabilityEntity) {
	mutationRate := 0.01
	for i := range timetable.Classes {
		if rand.Float64() < mutationRate {
			// Filter eligible teachers for the current subject
			availableProfessors := FilterEligibleProfessors(timetable.Classes[i].DisciplineID, professors)
			if len(availableProfessors) == 0 {
				continue
			}

			// Select a new teacher randomly
			newProfessor := availableProfessors[rand.Intn(len(availableProfessors))]

			// Filter the availabilities of the new professor
			availableSlots := FilterAvailableSlots(newProfessor.ID, availabilities)
			if len(availableSlots) == 0 {
				continue
			}

			// Find a time slot compatible with the shift and day and week of the original classroom
			weekDay := timetable.Classes[i].StartTime // Setting the original day of the week
			slot := findMatchingSlot(timetable.Classes[i].Shift, availableSlots, weekDay)
			if slot == nil {
				continue
			}

			// Recalculate startTime and endTime based on new teacher availability
			startTime, endTime := GenerateNextAvailableTime(weekDay, slot.Shift, nil, slot.DayOfWeek, timetable.Classes)

			// Check if the time is valid (it must not be zero)
			zeroTime := time.Time{}
			if startTime.Equal(zeroTime) {
				continue
			}

			// Update the class in the timetable with the new teacher and recalculated times
			timetable.Classes[i].ProfessorID = newProfessor.ID
			timetable.Classes[i].StartTime = startTime
			timetable.Classes[i].EndTime = endTime
		}
	}
}

func ReplacePopulation(population, newPopulation []entity.Timetable) {
	copy(population, newPopulation)
}
