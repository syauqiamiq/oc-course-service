package services

import "ocCourseService/repository"

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {

	return &service{
		repository: repository,
	}
}


type Service interface {

}
