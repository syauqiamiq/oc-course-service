package handler

import "ocCourseService/services"

type handler struct {
	service services.Service
}

func NewHandler(service services.Service) *handler {
	return &handler{
		service: service,
	}
}
