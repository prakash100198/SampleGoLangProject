package user_routes

import (
	"fmt"
	rest_handlers "attendance/RestHandlers"
	"github.com/gorilla/mux"
)

type AppRouterInterface interface {
	InitializeRouter(router *mux.Router)
}

type AppRouterInterfaceImpl struct {
	handlers  rest_handlers.RestHandler
	MuxRouter *mux.Router
}

func NewAppRouterInterfaceImpl(handlers rest_handlers.RestHandler) *AppRouterInterfaceImpl {
	handler := &AppRouterInterfaceImpl{
		handlers: handlers,
	}
	return handler
}

func (impl *AppRouterInterfaceImpl) InitializeRouter(r *mux.Router) {
	fmt.Println("initializing router..")
	//handler functions
	//student functions
	r.HandleFunc("/students", impl.handlers.GetStudents).Methods("GET")            //
	r.HandleFunc("/students/{studentId}", impl.handlers.GetStudent).Methods("GET") //

	//teacher functions
	r.HandleFunc("/teachers", impl.handlers.GetTeachers).Methods("GET")
	r.HandleFunc("/teachers/Id}", impl.handlers.GetTeacher).Methods("GET")
	//r.HandleFunc("/teachers/{teacherId}/attendance", GetTeacherAttandanceFromDB).Methods("GET")

	//principal functions
	r.HandleFunc("/principal/addStudent", impl.handlers.CreateStudent).Methods("POST") //
	// r.HandleFunc("/principal/addTeacher", rest_handlers.CreateTeacher).Methods("POST") //not implemented repetetive task

	//punchin/punchout triggering routes for student and teacher
	r.HandleFunc("/students/punchin", impl.handlers.Punchin).Methods("POST")
	r.HandleFunc("/teachers/punchin", impl.handlers.Punchin).Methods("POST")

	// r.HandleFunc("/students/punchout", rest_handlers.Punchout).Methods("POST")  // not implemented, repetetive task
	// r.HandleFunc("/teachers/punchout", rest_handlers.Punchout).Methods("POST")

	//this functionality is for the teacher to get attendance details for a particular class
	r.HandleFunc("/attendance/{class}/{day}/{month}/{year}", impl.handlers.GetAttandanceFromDB).Methods("GET") //

	//this functionality is for student/teacher to get his attendance details
	r.HandleFunc("/attendance/{id}/{month}/{year}", impl.handlers.GetOnePersonAttendance).Methods("GET") //

}
