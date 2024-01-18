package repository


type Authorization interface {

}

type Images interface {

}

type Documents interface {}

type Repository struct {
	Authorization
	Images
	Documents
}

func NewRepository() *Repository {
	return &Repository{}
}