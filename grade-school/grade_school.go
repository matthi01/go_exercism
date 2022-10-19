package school

import "sort"

type School struct {
	grades []Grade
}

type Grade struct {
	grade    int
	students []string
}

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	for i, g := range s.grades {
		if g.grade == grade {
			studentList := append(s.grades[i].students, student)
			sort.Strings(studentList)
			s.grades[i].students = studentList
			return
		}
	}
	newGrade := Grade{grade: grade, students: append([]string{}, student)}
	gradesList := append(s.grades, newGrade)
	sort.Slice(gradesList, func(i, j int) bool {
		return gradesList[i].grade < gradesList[j].grade
	})
	s.grades = gradesList
}

func (s *School) Grade(level int) []string {
	for _, g := range s.grades {
		if g.grade == level {
			return g.students
		}
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	return s.grades
}
