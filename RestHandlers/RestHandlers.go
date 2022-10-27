package rest_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// repository "attendance/Package/repository"
	models "attendance/Models"
	services "attendance/Services"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// var DB *gorm.DB
var err error

type RestHandler interface {
	GetStudents(w http.ResponseWriter, r *http.Request)
	GetStudent(w http.ResponseWriter, r *http.Request)
	GetAttandanceFromDB(w http.ResponseWriter, r *http.Request)
	GetOnePersonAttendance(w http.ResponseWriter, r *http.Request)
	GetTeachers(w http.ResponseWriter, r *http.Request)
	GetTeacher(w http.ResponseWriter, r *http.Request)
	CreateStudent(w http.ResponseWriter, r *http.Request)
	Punchin(w http.ResponseWriter, r *http.Request)
}
type RestHandlerImpl struct {
	studentServices services.StudentService
	teacherServices services.TeacherService
}

type PunchInInitialData struct {
	AttendanceId string `json:"attendanceid"` //id would be of either student or teacher
	Class        int    `json:"class"`
}

type Attendance struct {
	AttendanceId string `json:"attendanceid"` //id would be of either student or teacher
	Class        int    `json:"class"`
	Day          int    ` json:"day"`
	Month        int    ` json:"month"`
	Year         int    ` json:"year"`
	PunchInTime  string `json:"punchInTime"`
	PunchOutTime string `json:"punchOutTime"`
}

func NewRestHandlerImpl(studentServices services.StudentService, teacherServices services.TeacherService) *RestHandlerImpl {
	handler := &RestHandlerImpl{
		studentServices: studentServices,
		teacherServices: teacherServices,
	}
	return handler
}

func (impl *RestHandlerImpl) GetStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved for get students")

	// var students []models.Student
	// repository.DB.Find(&students)
	// var studentServiceInstance services.StudentServiceImpl
	// studentServiceInstance.GetStudentsFromDB()
	json.NewEncoder(w).Encode(impl.studentServices.GetStudentsFromDB())
}

func (impl *RestHandlerImpl) GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved to get student by id")

	params := mux.Vars(r) //to get the parameters from the url

	// //var student Student
	// var attendance []models.Attendance
	// repository.DB.Raw("Select attendance_id, class, day, month, year, punch_in_time, punch_out_time from attendances where attendance_id=?", params["studentId"]).Scan(&attendance)
	// //DB.Raw("Select * from students where student_id=?", params["studentId"]).Scan(&student)
	// json.NewEncoder(w).Encode(services.StudentService.GetStudentByID(&services.StudentServiceImpl{}, params))
	json.NewEncoder(w).Encode(impl.studentServices.GetStudentByID(params))
}

func (impl *RestHandlerImpl) GetAttandanceFromDB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request to get attendance record of student and teacher")
	params := mux.Vars(r) // getting parameters from the request
	// class, _ := strconv.Atoi(params["class"])
	// day, _ := strconv.Atoi(params["day"])
	// month, _ := strconv.Atoi(params["month"])
	// year, _ := strconv.Atoi(params["year"])

	// // var attendanceresponse Attendanceresponse
	// var attendance []models.Attendance

	// // DB.Model(&attendance).Where("class=? AND day=? AND month=? AND year=?", class, day, month, year)
	// repository.DB.Raw("Select attendance_id, class, day, month, year, punch_in_time, punch_out_time from attendances where class=? AND day=? AND month=? AND year=?", class, day, month, year).Scan(&attendance)
	json.NewEncoder(w).Encode(impl.studentServices.GetAttendanceByClass(params))
}

func (impl *RestHandlerImpl) GetOnePersonAttendance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request to get attendance record of one person")
	params := mux.Vars(r)
	// id := params["id"]
	// // id := params["id"]
	// month, _ := strconv.Atoi(params["month"])
	// year, _ := strconv.Atoi(params["year"])
	// var attendance []models.Attendance
	// repository.DB.Raw("Select attendance_id, class, day, month, year, punch_in_time, punch_out_time From attendances where attendance_id=? AND month=? AND year=?", id, month, year).Scan(&attendance)
	// for index := range attendance {
	json.NewEncoder(w).Encode(impl.studentServices.GetStudentByID(params))
	// }
}

func (impl *RestHandlerImpl) GetTeachers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved for get teachers")

	json.NewEncoder(w).Encode(impl.teacherServices.GetTeacher())
}

func (impl *RestHandlerImpl) GetTeacher(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved for get teacher by id")

	params := mux.Vars(r) //to get the parameters from the url

	// var attendance models.Attendance
	// repository.DB.Raw("Select * from attendances where attendance_id=?", params["teacherId"]).Scan(&attendance)
	// DB.Find(&teacher, params["teacherId"])
	json.NewEncoder(w).Encode(impl.teacherServices.GetTeacherById(params))
}

func (impl *RestHandlerImpl) CreateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved to create student")

	w.Header().Set("Content-Type", "application/json") //it repreesents the type of response the server is sending back to the client
	var student models.Student
	var resStudent models.Student
	json.NewDecoder(r.Body).Decode(&student) //decoding the value from the Api and decoding it in the form of Student struct type
	// json.NewDecoder(r.Body).Decode(&resStudent)

	// //var id int64
	// t := repository.DB.Last(&student)
	// if t.RecordNotFound() {
	// 	student.ID = 0
	// } else if t.Error != nil {
	// 	json.NewEncoder(w).Encode(t.Error)
	// }
	// // } else {
	// // 	json.NewEncoder(w).Encode(student.ID)
	// // }
	// var newStudentId = "S" + strconv.Itoa(int(student.ID)+1)
	// newstudent.StudentId = newStudentId

	// createdStudent := repository.DB.Create(&newstudent)
	// er := createdStudent.Error
	// DB.Model(&newstudent).Where("id=?", newstudent.ID).Update("studentId", newStudentId)
	// newstudent.StudentId = newStudentId
	resStudent, err := impl.studentServices.CreateStudent(student)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(resStudent)
	}

}

// func CreateTeacher(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("request recieved to create teacher")

// 	// w.Header().Set("Content-Type", "application/json")
// 	// var teacher models.Teacher
// 	// var newteacher models.Teacher

// 	// json.NewDecoder(r.Body).Decode(&newteacher) //decoding the value from the Api and decoding it in the form of Student struct type

// 	// t := repository.DB.Last(&teacher)
// 	// if t.RecordNotFound() {
// 	// 	teacher.ID = 0
// 	// } else if t.Error != nil {
// 	// 	json.NewEncoder(w).Encode(t.Error)
// 	// 	return
// 	// }
// 	// var newTeacherId = "T" + strconv.Itoa(int(teacher.ID)+1)
// 	// newteacher.TeacherId = newTeacherId

// 	// createdTeacher := repository.DB.Create(&newteacher)
// 	// er := createdTeacher.Error
// 	// // DB.Model(&newteacher).Where("id=?", newteacher.ID).Update("teacherId", newTeacherId)

// 	// if er != nil {
// 	// 	json.NewEncoder(w).Encode(er)
// 	// 	return
// 	// } else {
// 	// 	// newteacher.TeacherId = newTeacherId
// 	json.NewEncoder(w).Encode(services.CreateTeacherStudent(w, r))
// }

func (impl *RestHandlerImpl) Punchin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received for marking student attendance(punch in/ punch out)")
	w.Header().Set("Content-Type", "application/json")

	//var attendance Attendance
	//var newattendance models.Attendance
	var attendance models.Attendance
	json.NewDecoder(r.Body).Decode(&attendance)

	if attendance.AttendanceId == "" {
		var err error
		err = fmt.Errorf("please enter the id")
		json.NewEncoder(w).Encode(err)
		return
	}

	responseAttendance, er := impl.studentServices.PunchIn(attendance.AttendanceId, attendance.Class)

	if er != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewEncoder(w).Encode(responseAttendance)
	}

}

// func Punchout(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Request received for marking student attendance(punch in/ punch out)")
// 	w.Header().Set("Content-Type", "application/json")

// 	var newattendance models.Attendance

// 	json.NewDecoder(r.Body).Decode(&newattendance)

// 	if newattendance.AttendanceId == "" {
// 		var err error
// 		err = fmt.Errorf("please enter the id")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}
// 	currentTime := time.Now()
// 	year := currentTime.Year()
// 	month := currentTime.Month()
// 	day := currentTime.Day()
// 	punchOutTime := strconv.Itoa(currentTime.Hour()) + ":" + strconv.Itoa(currentTime.Minute()) + ":" + strconv.Itoa(currentTime.Second())

// 	attendanceId_DBstate := repository.DB.Model(&newattendance).Where("attendance_id=? AND class=? AND day=? AND month=? AND year=?", newattendance.AttendanceId, newattendance.Class, day, month, year)

// 	if attendanceId_DBstate.RecordNotFound() {
// 		var err error
// 		err = fmt.Errorf("please do punch in first")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	} else {
// 		repository.DB.Model(&newattendance).Where("attendance_id=? AND day=? AND month=? AND year=?", newattendance.AttendanceId, day, month, year).Update("punchOutTime", punchOutTime)
// 	}
// }
