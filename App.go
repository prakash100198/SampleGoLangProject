package main

import (
	// models "attendance/Package/Models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type AppInterfaceImpl struct {
	MuxRouter *mux.Router
	dB        *gorm.DB
}


func NewApp(MuxRouter *mux.Router, dB *gorm.DB) *AppInterfaceImpl {
	app := &AppInterfaceImpl{
		MuxRouter: MuxRouter,
		dB:        dB,
	}
	return app

}

// func (app *AppInterfaceImpl) InitializeRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	return r
// }

// func initialMigration() *gorm.DB {

// 	dbURI := "host=localhost user=postgres dbname=postgres sslmode=disable password=prakash port=5432"

// 	var dB *gorm.DB
// 	dB, err = gorm.Open("postgres", dbURI)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to Database")
// 	} else {
// 		fmt.Println("Successfully connected to database")
// 	}
// 	// defer DB.Close()
// 	dB.AutoMigrate(&models.Student{})
// 	dB.AutoMigrate(&models.Teacher{})
// 	dB.AutoMigrate(&models.Attendance{})
// 	return dB
// }

func (app *AppInterfaceImpl) Start() {

	// dbURI := "host=localhost user=postgres dbname=postgres sslmode=disable password=prakash port=5432"

	// var dB *gorm.DB
	// dB, err = gorm.Open("postgres", dbURI)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic("Cannot connect to Database")
	// } else {
	// 	fmt.Println("Successfully connected to database")
	// }
	// // defer DB.Close()
	// dB.AutoMigrate(&models.Student{})
	// dB.AutoMigrate(&models.Teacher{})
	// dB.AutoMigrate(&models.Attendance{})
	fmt.Println("Initializing Router")

	log.Fatal(http.ListenAndServe(":8090", app.MuxRouter))
}
