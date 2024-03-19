package services

import (
	models "attendance/Models"
	repository "attendance/Repository"
)

type TeacherService interface { //this is an abstraction layer, we don't want to expose the implementation to rest handlers directly thats why we are doing it via interfaces.
	GetTeacher() []models.Teacher
	GetTeacherById(params map[string]string) []models.Attendance
}

type TeacherServiceImpl struct { //this is a class in java for equivalency purpose
	// y repositoryService
	//here whatever variables we use outside the scope of this file we put define here so that we can use in services file.
	repoModels repository.Repo
}

func NewTeacherServiceImpl(repoModels repository.Repo) *TeacherServiceImpl {
	handler := &TeacherServiceImpl{
		repoModels: repoModels,
	}
	return handler

}

func (impl *TeacherServiceImpl) GetTeacher() []models.Teacher {
	return impl.repoModels.GetTeachers()

}

func (impl *TeacherServiceImpl) GetTeacherById(params map[string]string) []models.Attendance {
	return impl.repoModels.GetTeacherAttendanceById(params)
}
