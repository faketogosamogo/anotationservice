package repository

import (
	"annotation-microservice/model"
	"errors"
)

type WordRepositorySlice struct {
	words []model.Word
}
func NewWordRepositorySlice() *WordRepositorySlice{
	return &WordRepositorySlice{
		make([]model.Word, 0),
	}
}
func(w *WordRepositorySlice) Index() ([]model.Word, error){
	return w.words, nil
}
func(w *WordRepositorySlice) RemoveAll() error{
	w.words = make([]model.Word, 0)
	return nil
}

func(w *WordRepositorySlice) AddRange(words []model.Word) error{
	tempWords := make([]model.Word, 0)

	for _, word := range words{
		if len(word.Text)==0{
			return errors.New("Текст в слове обязателен!")
		}
		tempWords = append(tempWords, word)
	}
	w.words = append(w.words, tempWords...)
	return nil
}
