package services

import (
	models "attendance/Package/Models"
	repository "attendance/Package/Repository"
	"errors"
	"fmt"
	"strconv"
	"time"
	// "encoding/json"
	// "net/http"
	// "strconv"
)

var err error

type StudentService interface { //this is an abstraction layer, we don't want to expose the implementation to rest handlers directly thats why we are doing it via interfaces.
	GetStudentsFromDB() []models.Student
	GetStudentByID(params map[string]string) []models.Attendance
	GetAttendanceByClass(params map[string]string) []models.Attendance
	GetAttendanceById(params map[string]string) []models.Attendance
	CreateStudent(student models.Student) (models.Student, error)
	PunchIn(string, int) (models.Attendance, error)
}

type StudentServiceImpl struct { //this is a class in java for equivalency purpose
	// y repositoryService
	//here whatever variables we use outside the scope of this file we put define here so that we can use in services file.
	repoModels repository.Repo
}

type Student struct {
	StudentId string `json:"studentId"` //S1, S2,....
	Name      string `json:"studentName"`
	Class     int    `json:"studentClass"`
}

func NewStudentServiceImpl(repoModels repository.Repo) *StudentServiceImpl {
	handler := &StudentServiceImpl{
		repoModels: repoModels,
	}
	return handler
}

func (impl *StudentServiceImpl) GetStudentsFromDB() []models.Student {
	// var students repository.RepoImpl
	return impl.repoModels.GetAttendanceOfAllStudents()
	//  GetAttendanceOfAllStudents()
	// repository.DB.Find(&students)

}

func (impl *StudentServiceImpl) GetStudentByID(params map[string]string) []models.Attendance {
	return impl.repoModels.GetAttendanceByStudentId(params)
	//repository.DB.Raw("Select * from attendances where attendance_id=?", params["studentId"]).Scan(&attendance)
}

func (impl *StudentServiceImpl) GetAttendanceByClass(params map[string]string) []models.Attendance {
	// params := mux.Vars(r) // getting parameters from the request

	return impl.repoModels.GetAttendanceByClass(params)
}

func (impl *StudentServiceImpl) GetAttendanceById(params map[string]string) []models.Attendance {
	// params := mux.Vars(r) // getting parameters from the request

	return impl.repoModels.GetAttendanceByStudentId(params)
}

func (impl *StudentServiceImpl) CreateStudent(student models.Student) (models.Student, error) {

	return impl.repoModels.CreateStudent(student)
}

func (impl *StudentServiceImpl) PunchIn(attendanceId string, class int) (models.Attendance, error) {
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	punchInTime := strconv.Itoa(currentTime.Hour()) + ":" + strconv.Itoa(currentTime.Minute()) + ":" + strconv.Itoa(currentTime.Second())

	var newattendance models.Attendance
	newattendance.AttendanceId = attendanceId
	newattendance.Class = class

	attendanceId_DBState := impl.repoModels.PunchIn(&newattendance)
	// attendanceId_DBState := repository.DB.Model(&newattendance).Where("attendance_id=? AND class=? AND day=? AND month=? AND year=?", newattendance.AttendanceId, newattendance.Class, day, month, year).Last(&newattendance)

	if attendanceId_DBState.RecordNotFound() {

		fmt.Println("no record found hence appending current data to attendances table")
		newattendance.Day = day
		newattendance.Month = int(month)
		newattendance.Year = year
		newattendance.PunchInTime = punchInTime

		createdStudentRecord := impl.repoModels.CreateAttendanceRecord(&newattendance)
		// createdStudentRecord := DB.Create(&newattendance)
		er := createdStudentRecord.Error
		return newattendance, er
		// if err != nil {
		// 	json.NewEncoder(w).Encode(err)
		// 	return
		// } else {
		// 	json.NewEncoder(w).Encode(newattendance)
		// }
	} else {
		err := errors.New("Punch In Record for this ID already found in the database")
		return models.Attendance{}, err

	}

	// if dayFromDB.RecordNotFound() {
	// 	createdStudentRecord := DB.Create(&newattendance)
	// 	checkRecordInAttendanceDB(newattendance, w, createdStudentRecord)

	// } else if monthFromDB.RecordNotFound() {
	// 	createdStudentRecord := DB.Create(&newattendance)
	// 	checkRecordInAttendanceDB(newattendance, w, createdStudentRecord)

	// } else if yearFromDB.RecordNotFound() {
	// 	createdStudentRecord := DB.Create(&newattendance)
	// 	checkRecordInAttendanceDB(newattendance, w, createdStudentRecord)
	// } else {
	// 	DB.Model(&newattendance).Where("attendanceid=? AND day=? AND month=? AND year=?", newattendance.AttendanceId, newattendance.Day, newattendance.Month, newattendance.Year).Update("punchOUtTime", newattendance.PunchOutTime)
	// }
}

// func (impl *StudentServiceImpl) CreateTeacherStudent(w http.ResponseWriter, r *http.Request) models.Student {
// 	w.Header().Set("Content-Type", "application/json") //it repreesents the type of response the server is sending back to the client
// 	var student models.Student
// 	var newstudent models.Student
// 	//json.NewDecoder(r.Body).Decode(&student) //decoding the value from the Api and decoding it in the form of Student struct type
// 	json.NewDecoder(r.Body).Decode(&newstudent)

// 	//var id int64
// 	t := repository.DB.Last(&student)
// 	if t.RecordNotFound() {
// 		student.ID = 0
// 	}

// 	var newStudentId = "S" + strconv.Itoa(int(student.ID)+1)
// 	newstudent.StudentId = newStudentId

// 	createdStudent := repository.DB.Create(&newstudent)
// 	er := createdStudent.Error
// 	// DB.Model(&newstudent).Where("id=?", newstudent.ID).Update("studentId", newStudentId)

// 	if er != nil {
// 		json.NewEncoder(w).Encode(er)
// 	}
// 	// newstudent.StudentId = newStudentId
// 	return newstudent

// }
