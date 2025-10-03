package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 20, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Helper to build a calculator with one grade per type and reduce redundant code.
func gradeHelper(a, e, s int) *GradeCalculator {
    g := NewGradeCalculator()
    g.AddGrade("assignment", a, Assignment)
    g.AddGrade("exam", e, Exam)
    g.AddGrade("essay", s, Essay)
    return g
}

func TestLetterBoundaries(t *testing.T) {
    t.Run("exact A at 90", func(t *testing.T) {
        got := gradeHelper(90, 90, 90).GetFinalGrade()
        if got != "A" {
            t.Fatalf("want A, got %s", got)
        }
    })
    t.Run("exact B at 80", func(t *testing.T) {
        got := gradeHelper(80, 80, 80).GetFinalGrade()
        if got != "B" {
            t.Fatalf("want B, got %s", got)
        }
    })
    t.Run("exact C at 70", func(t *testing.T) {
        got := gradeHelper(70, 70, 70).GetFinalGrade()
        if got != "C" {
            t.Fatalf("want C, got %s", got)
        }
    })
    t.Run("exact D at 60", func(t *testing.T) {
        got := gradeHelper(60, 60, 60).GetFinalGrade()
        if got != "D" {
            t.Fatalf("want D, got %s", got)
        }
    })
    t.Run("F below 60", func(t *testing.T) {
        got := gradeHelper(59, 59, 59).GetFinalGrade()
        if got != "F" {
            t.Fatalf("want F, got %s", got)
        }
    })
}

func TestAveragingWithinType(t *testing.T) {
    g := NewGradeCalculator()
    g.AddGrade("assignment 1", 100, Assignment)
    g.AddGrade("assignment 2", 0, Assignment) 
    g.AddGrade("exam 1", 100, Exam)
    g.AddGrade("essay 1", 100, Essay)
    if got := g.GetFinalGrade(); got != "C" {
        t.Fatalf("want C, got %s", got)
    }
}

func TestMixedValuesProduceB(t *testing.T) {
    got := gradeHelper(88, 82, 84).GetFinalGrade()
    if got != "B" {
        t.Fatalf("want B, got %s", got)
    }
}

func TestGetGradeCBoundary(t *testing.T) {
    g := NewGradeCalculator()
    g.AddGrade("assignment", 70, Assignment)
    g.AddGrade("exam", 70, Exam)
    g.AddGrade("essay", 70, Essay)
    if got := g.GetFinalGrade(); got != "C" {
        t.Fatalf("want C, got %s", got)
    }
}

func TestGetGradeDBoundary(t *testing.T) {
    g := NewGradeCalculator()
    g.AddGrade("assignment", 60, Assignment)
    g.AddGrade("exam", 60, Exam)
    g.AddGrade("essay", 60, Essay)
    if got := g.GetFinalGrade(); got != "D" {
        t.Fatalf("want D, got %s", got)
    }
}

// Test String()
func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Fatalf("Assignment.String() = %q, want %q", Assignment.String(), "assignment")
    }
    if Exam.String() != "exam" {
        t.Fatalf("Exam.String() = %q, want %q", Exam.String(), "exam")
    }
    if Essay.String() != "essay" {
        t.Fatalf("Essay.String() = %q, want %q", Essay.String(), "essay")
    }
}

// Test computeAverage covering all the branches:
func TestComputeAverageGrades(t *testing.T) {
    // empty slice
    if got := computeAverage([]Grade{}); got != 0 {
        t.Errorf("expect 0 for empty slice, got %v", got)
    }
    // one grade
    if got := computeAverage([]Grade{{Name: "g1", Grade: 75, Type: Assignment}}); got != 75 {
        t.Errorf("expect 75 for single grade, got %v", got)
    }
    // multiple grades: (80+90+100)/3 = 90
    grades := []Grade{
        {Name: "g1", Grade: 80, Type: Assignment},
        {Name: "g2", Grade: 90, Type: Assignment},
        {Name: "g3", Grade: 100, Type: Assignment},
    }
    if got := computeAverage(grades); got != 90 {
        t.Errorf("expect 90 for 80,90,100, got %v", got)
    }
}


// Test calculateNumericalGrade – exercise the paths for
func TestCalculateNumericalGrade(t *testing.T) {
    // empty -> 0
    g := NewGradeCalculator()
    if got := g.calculateNumericalGrade(); got != 0 {
        t.Errorf("expected 0 for empty grades, got %v", got)
    }

    // perfect grades -> 100
    g.AddGrade("assignment", 100, Assignment)
    g.AddGrade("exam", 100, Exam)
    g.AddGrade("essay", 100, Essay)
    if got := g.calculateNumericalGrade(); got != 100 {
        t.Errorf("expected 100 for perfect grades, got %v", got)
    }

    // mixed case to exercise truncation: 0.5*88 + 0.35*82 + 0.15*84 = 85.3 -> int(85.3)=85
    g2 := NewGradeCalculator()
    g2.AddGrade("a", 88, Assignment)
    g2.AddGrade("e", 82, Exam)
    g2.AddGrade("s", 84, Essay)
    if got := g2.calculateNumericalGrade(); got != 85 {
        t.Errorf("expected 85 for 88/82/84 weighted average, got %v", got)
    }
}