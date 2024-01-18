package service

import "server/pkg/repository"

type Authorization interface {

}

type Images interface {

}

type Documents interface {}

type Service struct {
	Authorization
	Images
	Documents
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}