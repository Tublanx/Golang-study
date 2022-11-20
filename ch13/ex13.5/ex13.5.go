package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수: %.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{19, 2174, 90.2}

	student2 := student

	PrintStudent(student2)
}
