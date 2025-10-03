# ESEP Grade Calculator

This library is a simple utility to calculate a student's grades. The grades are based on three categories: assignments, essays, and exams. Grades are weighted and calculated into a final letter grade based on the grading scale.

Exam Weights:
 - Assignments = 50%
 - Exams = 35%
 - Essays = 15%

Numeric grade averages to letter grades:
 - A >= 90
 - B >= 80
 - C >= 70
 - D >= 60
 - F < 60

Runtime requirements:
Golang 1.24+

Setup:
```
git clone https://github.com/your-username/esep-grade-calculator.git
cd esep-grade-calculator
```

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```