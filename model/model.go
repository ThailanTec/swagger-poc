package model

import "github.com/google/uuid"

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type UserInput struct {
	Name  string `json:"name" binding:"required" example:"Thailan Santos"`
	Email string `json:"email" binding:"required" example:"thailan@gmail.com"`
}

type UserOutput struct {
	ID     uuid.UUID `json:"id" example:"c087a2ed-6694-49ad-884a-eecade5a2f90"`
	Status Status    `json:"status" example:"active"`
	Name   string    `json:"name" example:"Thailan Santos"`
	Email  string    `json:"email" example:"thailan@gmail.com"`
}

// ResourceResponse é a estrutura de resposta para o recurso criado
type ResourceResponse struct {
	ID   int    `json:"id" example:"c087a2ed-6694-49ad-884a-eecade5a2f90"`
	Name string `json:"name" example:"Thailan Santos"`
}

// ErrorResponse é a estrutura de resposta para erros
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request, erro ao criar usuário."`
}
