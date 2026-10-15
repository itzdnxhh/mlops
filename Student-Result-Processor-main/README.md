# Student Result Processing System

A modular command-line application for processing student results using Go.

## 1. Project Overview

The application accepts student details and marks for five subjects. It calculates the total, average, percentage, grade, and pass/fail status.

## 2. Objectives

- Develop a command-line application using Go.
- Apply structs, arrays, slices, loops, functions, and packages.
- Organize the application into modules.
- Validate student details and subject marks.
- Use Git and GitHub for version control.

## 3. Technologies Used

- Go programming language
- Visual Studio Code
- Git
- GitHub

## 4. Project Structure

```text
StudentResultProcessor/
├── main.go
├── go.mod
├── README.md
├── models/
│   └── student.go
└── result/
    └── processor.go
```

## 5. Features

- Processes multiple students in one run.
- Accepts marks for five subjects.
- Validates marks between 0 and 100.
- Calculates total, average, and percentage.
- Assigns grades based on percentage.
- Marks a student as FAIL if any subject mark is below 35.

## 6. Grade Criteria

| Percentage | Grade |
|---|---|
| 90–100 | A+ |
| 80–89.99 | A |
| 70–79.99 | B |
| 60–69.99 | C |
| 50–59.99 | D |
| Below 50 | F |

A student must score at least 35 in every subject to pass.

## 7. How to Run

Ensure Go is installed, open a terminal in the project directory, and execute:

```bash
go run .
```

Follow the prompts to enter the student details and subject marks.

## 8. Testing

Test with valid marks, boundary values (0 and 100), invalid input, and marks below 35. Record the actual results produced by the program.

## 9. Conclusion

This project demonstrates Go programming, modular application design, student result processing, and Git/GitHub version control.