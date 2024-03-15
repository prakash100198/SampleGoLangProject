package repositories

import (
	models "attendance/Models"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)



var err error


type Repo interface {
	GetAttendanceByStudentId(params map[string]string) []models.Attendance
	GetAttendanceOfAllStudents() []models.Student
	GetAttendanceByClass(params map[string]string) []models.Attendance
	GetAttendanceById(params map[string]string) []models.Attendance
	//Teachers functionality :-
	GetTeachers() []models.Teacher
	GetTeacherAttendanceById(params map[string]string) []models.Attendance

	//Creating student
	CreateStudent(student models.Student) (models.Student, error)
	PunchIn(*models.Attendance) *gorm.DB
	CreateAttendanceRecord(*models.Attendance) *gorm.DB
}

type RepoImpl struct {
	DB *gorm.DB
}

type Student struct {
	StudentId string `json:"studentId"` //S1, S2,....
	Name      string `json:"studentName"`
	Class     int    `json:"studentClass"`
}



func NewRepoImpl(DB *gorm.DB) *RepoImpl {
	handler := &RepoImpl{
		DB: DB,
	}
	return handler
}

func (impl *RepoImpl) GetAttendanceByStudentId(params map[string]string) []models.Attendance {
	var attendance []models.Attendance
	impl.DB.Raw("Select * from attendances where attendance_id=?", params["studentId"]).Scan(&attendance)
	return attendance
}

func (impl *RepoImpl) GetAttendanceOfAllStudents() []models.Student {
	var student []models.Student
	impl.DB.Find(&student)
	return student
}

func (impl *RepoImpl) GetAttendanceByClass(params map[string]string) []models.Attendance {
	class, _ := strconv.Atoi(params["class"])
	day, _ := strconv.Atoi(params["day"])
	month, _ := strconv.Atoi(params["month"])
	year, _ := strconv.Atoi(params["year"])

	// var attendanceresponse Attendanceresponse
	var attendance []models.Attendance

	// DB.Model(&attendance).Where("class=? AND day=? AND month=? AND year=?", class, day, month, year)
	impl.DB.Raw("Select * from attendances where class=? AND day=? AND month=? AND year=?", class, day, month, year).Scan(&attendance)
	return attendance
}

func (impl *RepoImpl) GetAttendanceById(params map[string]string) []models.Attendance {
	id := params["id"]
	// id := params["id"]
	month, _ := strconv.Atoi(params["month"])
	year, _ := strconv.Atoi(params["year"])

	var attendance []models.Attendance

	impl.DB.Raw("Select * From attendances where attendance_id=? AND month=? AND year=?", id, month, year).Scan(&attendance)
	return attendance
}

func (impl *RepoImpl) GetTeachers() []models.Teacher {
	var teachers []models.Teacher
	impl.DB.Find(&teachers)
	return teachers
}

func (impl *RepoImpl) GetTeacherAttendanceById(params map[string]string) []models.Attendance {
	var attendance []models.Attendance
	impl.DB.Raw("Select * from attendances where attendance_id=?", params["teacherId"]).Scan(&attendance)
	return attendance
}

func (impl *RepoImpl) CreateStudent(student models.Student) (models.Student, error) {
	// var id int64
	// var newstudent models.Student

	var laststudent models.Student

	t := impl.DB.Last(&laststudent)
	if t.RecordNotFound() {
		student.ID = 0
	} else if t.Error != nil {
		// json.NewEncoder(w).Encode(t.Error)
		return student, t.Error
	}
	var newStudentId = "S" + strconv.Itoa(int(laststudent.ID)+1)
	student.StudentId = newStudentId

	createdStudent := impl.DB.Create(&student)
	er := createdStudent.Error
	//impl.DB.Model(&newstudent).Where("id=?", student.ID).Update("studentId", newStudentId) // we can also fetch
	// newstudent.StudentId = newStudentId

	return student, er

}

func (impl *RepoImpl) PunchIn(attendance *models.Attendance) *gorm.DB {
	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	// punchInTime := strconv.Itoa(currentTime.Hour()) + ":" + strconv.Itoa(currentTime.Minute()) + ":" + strconv.Itoa(currentTime.Second())
	var newattendance models.Attendance
	newattendance.AttendanceId = attendance.AttendanceId
	newattendance.Class = attendance.Class

	db_state := impl.DB.Model(&attendance).Where("attendance_id=? AND class=? AND day=? AND month=? AND year=?", attendance.AttendanceId, attendance.Class, day, month, year).Last(&attendance)
	return db_state
	// if db_state.RecordNotFound() {

	// 	fmt.Println("no record found hence appending current data to attendances table")
	// 	newattendance.Day = day
	// 	newattendance.Month = int(month)
	// 	newattendance.Year = year
	// 	newattendance.PunchInTime = punchInTime

	// 	// createdStudentRecord := repository.DB.Create(&newattendance)
	// } else {

	// }

}
func (impl *RepoImpl) CreateAttendanceRecord(attendance *models.Attendance) *gorm.DB {
	createdStudentRecord := impl.DB.Create(&attendance)
	return createdStudentRecord
}
