package repository

import "annotation-microservice/model"

type WordRepository interface {
	AddRange(words []model.Word) error
	Index() ([]model.Word, error)
	RemoveAll() error
}


