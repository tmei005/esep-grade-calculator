package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	var assignmentAvg, examAvg, essayAvg float64

	if len(gc.assignments) > 0 {
		sum := 0
		for _, grade := range gc.assignments {
			sum += grade.Grade
		}
		assignmentAvg = float64(sum) / float64(len(gc.assignments))
	}

	if len(gc.exams) > 0 {
		sum := 0
		for _, grade := range gc.exams {
			sum += grade.Grade
		}
		examAvg = float64(sum) / float64(len(gc.exams))
	}

	if len(gc.essays) > 0 {
		sum := 0
		for _, grade := range gc.essays {
			sum += grade.Grade
		}
		essayAvg = float64(sum) / float64(len(gc.essays))
	}

	finalScore := assignmentAvg*0.5 + examAvg*0.35 + essayAvg*0.15

	switch {
	case finalScore >= 90:
		return "A"
	case finalScore >= 80:
		return "B"
	case finalScore >= 70:
		return "C"
	case finalScore >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{Name: name, Grade: grade, Type: Assignment})
	case Exam:
		gc.exams = append(gc.exams, Grade{Name: name, Grade: grade, Type: Exam})
	case Essay:
		gc.essays = append(gc.essays, Grade{Name: name, Grade: grade, Type: Essay})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.assignments)
	exam_average := computeAverage(gc.exams)
	essay_average := computeAverage(gc.essays)

	weighted_grade := float64(assignment_average)*0.5 + float64(exam_average)*0.35 + float64(essay_average)*0.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}

	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}
	return sum / len(grades)
}
