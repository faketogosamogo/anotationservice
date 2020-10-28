package repository

import (
	"annotation-microservice/model"
	"testing"
)

func TestWordRepositorySlice_Index(t *testing.T) {
	repo := NewWordRepositorySlice()
	count := 10
	testWords := make([]model.Word, 0)
	for i:=0; i < count; i++ {
		testWords = append(testWords, model.Word{Text: "text"})
	}
	repo.AddRange(testWords)
	words, err := repo.Index()
	if err!=nil{
		t.Errorf("Ошибки не должно быть")
	}
	if len(words) != count{
		t.Errorf("В репозитории должно быть %v элементов, а на данный момент их: %v ", count, len(words))
	}
}
func TestWordRepositorySlice_RemoveAll(t *testing.T) {
	repo := NewWordRepositorySlice()
	count := 10
	testWords := make([]model.Word, 0)
	for i:=0; i < count; i++ {
		testWords = append(testWords, model.Word{Text: "text"})
	}
	repo.RemoveAll()
	if len(repo.words)!=0{
		t.Errorf("Репозиторий не очистился")
	}
}
func TestWordRepositorySlice_AddRange(t *testing.T) {
	repo := NewWordRepositorySlice()
	count := 10
	testWords := make([]model.Word, 0)
	for i:=0; i < count; i++ {
		testWords = append(testWords, model.Word{Text: "text"})
	}
	err := repo.AddRange(testWords)
	if err!=nil{
		t.Errorf("Ошибки не должно быть")
	}
	if count!=len(repo.words){
		t.Errorf("Добавлялось: %v, добавилось: %v", testWords, len(repo.words))
	}
}
func TestWordRepositorySlice_AddRange_Should_Return_err(t *testing.T) {
	repo := NewWordRepositorySlice()
	count := 10
	testWords := make([]model.Word, 0)
	for i:=0; i < count; i++ {
		testWords = append(testWords, model.Word{Text: ""})
	}
	err := repo.AddRange(testWords)
	if err==nil{
		t.Errorf("Должна вернуться ошибка")
	}
}