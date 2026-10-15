
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"studentresult/models"
	"studentresult/result"
)

func readLine(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		text, err := reader.ReadString('\n')
		if err != nil && len(text) == 0 {
			fmt.Println("\nUnable to read input. Please try again.")
			continue
		}
		return strings.TrimSpace(text)
	}
}

func readInt(reader *bufio.Reader, prompt string, min, max int) int {
	for {
		text := readLine(reader, prompt)
		value, err := strconv.Atoi(text)

		if err == nil && value >= min && value <= max {
			return value
		}

		fmt.Printf("Enter a number between %d and %d.\n", min, max)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("===================================")
	fmt.Println(" STUDENT RESULT PROCESSING SYSTEM")
	fmt.Println("===================================")

	count := readInt(reader, "Enter number of students: ", 1, 100)

	students := make([]models.Student, 0, count)

	for i := 0; i < count; i++ {
		fmt.Printf("\n--- Student %d ---\n", i+1)

		rollNo := readInt(reader, "Enter roll number: ", 1, 999999)
		name := readLine(reader, "Enter student name: ")

		for name == "" {
			fmt.Println("Name cannot be empty.")
			name = readLine(reader, "Enter student name: ")
		}

		student := models.Student{
			RollNo: rollNo,
			Name:   name,
		}

		for subject := 0; subject < models.SubjectCount; subject++ {
			prompt := fmt.Sprintf("Enter mark for Subject %d (0-100): ", subject+1)
			student.Marks[subject] = readInt(reader, prompt, 0, 100)
		}

		students = append(students, student)
	}

	fmt.Println("\n========== RESULT SUMMARY ==========")

	for _, student := range students {
		res := result.ProcessStudent(student)

		fmt.Printf("\nRoll Number : %d\n", student.RollNo)
		fmt.Printf("Name        : %s\n", student.Name)
		fmt.Printf("Total       : %d / 500\n", res.Total)
		fmt.Printf("Average     : %.2f\n", res.Average)
		fmt.Printf("Percentage  : %.2f%%\n", res.Percentage)
		fmt.Printf("Grade       : %s\n", res.Grade)

		if res.Passed {
			fmt.Println("Status      : PASS")
		} else {
			fmt.Println("Status      : FAIL")
		}
	}

	fmt.Println("\n===================================")
}
