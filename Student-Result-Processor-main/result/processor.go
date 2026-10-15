package result

import "studentresult/models"

type Result struct {
	Total      int
	Average    float64
	Percentage float64
	Grade      string
	Passed     bool
}

func ProcessStudent(s models.Student) Result {
	total := 0
	passed := true

	for _, mark := range s.Marks {
		total += mark
		if mark < 35 {
			passed = false
		}
	}

	average := float64(total) / float64(len(s.Marks))
	percentage := float64(total) / float64(len(s.Marks)*100) * 100

	grade := "F"
	switch {
	case percentage >= 90:
		grade = "A+"
	case percentage >= 80:
		grade = "A"
	case percentage >= 70:
		grade = "B"
	case percentage >= 60:
		grade = "C"
	case percentage >= 50:
		grade = "D"
	}

	return Result{
		Total:      total,
		Average:    average,
		Percentage: percentage,
		Grade:      grade,
		Passed:     passed,
	}
}
