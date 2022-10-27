package main

import (
	repository "attendance/Repository"
	rest_handlers "attendance/RestHandlers"
	user_routes "attendance/Router"
	services "attendance/Services"

	models "attendance/Models"
	"fmt"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// "fmt"

func DBConnectionProvider() *gorm.DB {
	dbURI := "host=localhost user=postgres dbname=postgres sslmode=disable password=prakash port=5432"

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

func InitializeEvent() *AppInterfaceImpl {

	wire.Build(
		user_routes.NewAppRouterInterfaceImpl,
		wire.Bind(new(user_routes.AppRouterInterface), new(*user_routes.AppRouterInterfaceImpl)),

		rest_handlers.NewRestHandlerImpl,
		wire.Bind(new(rest_handlers.RestHandler), new(*rest_handlers.RestHandlerImpl)),

		services.NewTeacherServiceImpl,
		wire.Bind(new(services.TeacherService), new(*services.TeacherServiceImpl)),

		services.NewStudentServiceImpl,
		wire.Bind(new(services.StudentService), new(*services.StudentServiceImpl)),

		repository.NewRepoImpl,
		wire.Bind(new(repository.Repo), new(*repository.RepoImpl)),

		// models.NewModelImpl,
		// wire.Bind(new(models.Model), new(models.ModelImpl)),
		NewApp,
		NewMuxProvider,
		DBConnectionProvider,
		// wire.Bind(new(AppInterface), new(*AppInterfaceImpl)),

	)
	return &AppInterfaceImpl{}
}
