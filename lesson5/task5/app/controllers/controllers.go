package controllers

import "app/usecase"

type AppControllers struct {
	uc usecase.Usecase
}

func NewAppControllers(uc usecase.Usecase) *AppControllers {
	return &AppControllers{
		uc: uc,
	}
}
