
package models

const SubjectCount = 5

type Student struct {
    RollNo int
    Name   string
    Marks  [SubjectCount]int
}
