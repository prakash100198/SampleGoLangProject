package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	StudentId string `json:"studentId"` //S1, S2,....
	Name      string `json:"studentName"`
	Class     int    `json:"studentClass"`
}

type Teacher struct {
	gorm.Model
	TeacherId string `json:"teacherId"` //T1, T2,....
	Name      string `json:"teacherName"`
}

type Attendance struct {
	gorm.Model
	AttendanceId string `json:"attendanceid"` //id would be of either student or teacher
	Class        int    `json:"class"`
	Day          int    ` json:"day"`
	Month        int    ` json:"month"`
	Year         int    ` json:"year"`
	PunchInTime  string `json:"punchInTime"`
	PunchOutTime string `json:"punchOutTime"`
}
type Model interface {
}
type ModelImpl struct {
	student    Student
	teacher    Teacher
	attendance Attendance
}

func NewModelImpl(student Student, teacher Teacher, attendance Attendance) *ModelImpl {
	model := &ModelImpl{
		student:    student,
		teacher:    teacher,
		attendance: attendance,
	}
	return model
}

