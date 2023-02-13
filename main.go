package main

import (
	// models "attendance/Package/Models"
	models "attendance/Models"
	repository "attendance/Repository"
	rest_handlers "attendance/RestHandlers"
	user_routes "attendance/Router"
	services "attendance/Services"
	"fmt"
	"log"
	"net/http"

	// "encoding/json"

	// "github.com/gorilla/mux"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	//"gorm.io/gorm"
)

var err error

type db struct {
	db *gorm.DB
}

func main() {
	fmt.Println("running main...")

	r := mux.NewRouter()
	// NewApp(r, dB)

	// app := InitializeEvent()
	// fmt.Println("inside main.go main function")

	// app.Start()
	db := initialMigration()
	repo := repository.NewRepoImpl(db)
	serviceStudent := services.NewStudentServiceImpl(repo)
	servicesTeacher := services.NewTeacherServiceImpl(repo)
	restHandler := rest_handlers.NewRestHandlerImpl(serviceStudent, servicesTeacher)
	userRoutes := user_routes.NewAppRouterInterfaceImpl(restHandler)
	// app:= NewApp(r, dB)

	// e: InitializeEvent()
	userRoutes.InitializeRouter(r)
	fmt.Println("initializing router..")
	log.Fatal(http.ListenAndServe(":32026", r))
}

func initialMigration() *gorm.DB {

	dbURI := "host=postgres-postgresql.devtron-demo user=postgres dbname=postgres sslmode=disable password=prakash port=5432"
	//postgres-devtron-demo-postgresql-0
	// url := "postgres://postgres:prakash@localhost:5432/postgres?sslmode=disable"
	
	var dB *gorm.DB
	dB, err = gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	} else {
		fmt.Println("Successfully connected to database")
	}
	// defer DB.Close()
	dB.AutoMigrate(&models.Student{})
	dB.AutoMigrate(&models.Teacher{})
	dB.AutoMigrate(&models.Attendance{})
	return dB
}
