package controller

import (
	repository "annotation-microservice/repository/slice"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestAnagramController_Load(t *testing.T) {
	wordRepository := repository.NewWordRepositorySlice()
	anagramController := NewAnagramController(wordRepository)

	words := [...]string{"foobar", "aabb", "baba", "boofar", "test"}
	wordsJson, _ := json.Marshal(words)
	body := bytes.NewReader(wordsJson)

	req, err := http.NewRequest("POST", "/Load", body)
	if err!=nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()
	loadHandler := http.HandlerFunc(anagramController.Load)

	loadHandler.ServeHTTP(respRecorder, req)
	if respRecorder.Code!=http.StatusOK{
		t.Errorf("Код должен быть 200, но он %v", respRecorder.Code)
	}
	wordIndex, err := wordRepository.Index()
	if err!=nil{
		t.Errorf("Произошла ошибка получения элеметов из репозитория: %s", err)
	}
	if len(wordIndex) != len(words){
		t.Errorf("Длина загруженного не совпадает с передаваемым: %s", err)
	}
}

func TestAnagramController_Load_Should_Return_Err(t *testing.T) {
	wordRepository := repository.NewWordRepositorySlice()
	anagramController := NewAnagramController(wordRepository)


	words := [...]string{"foobar", "", "baba", "boofar", "test"}
	wordsJson, _ := json.Marshal(words)
	body := bytes.NewReader(wordsJson)

	req, err := http.NewRequest("POST", "/Load", body)
	if err!=nil {
		t.Fatal(err)
	}

	respRecorder := httptest.NewRecorder()
	loadHandler := http.HandlerFunc(anagramController.Load)

	loadHandler.ServeHTTP(respRecorder, req)
	if respRecorder.Code!=http.StatusUnprocessableEntity{
		t.Errorf("Код должен быть 422, но он %v", respRecorder.Code)
	}
	wordIndex, err := wordRepository.Index()
	if err!=nil{
		t.Errorf("Произошла ошибка получения элеметов из репозитория: %s", err)
	}
	if len(wordIndex) != 0{
		t.Error("Репозиторий должен быть пуст!")
	}
}

func TestAnagramController_GetAnagrams(t *testing.T) {
	words := []string{"foobar", "aabb", "baba", "boofar", "test"}
	expected := []string{"foobar", "boofar"}

	wordRepository := repository.NewWordRepositorySlice()
	wordRepository.AddRange(castStringsToWords(words))

	anagramController := NewAnagramController(wordRepository)

	req, err := http.NewRequest("GET", "/get?word=foobar", nil)
	if err!=nil {
		t.Fatal(err)
	}
	respRecorder := httptest.NewRecorder()
	getHandler := http.HandlerFunc(anagramController.GetAnagrams)

	getHandler.ServeHTTP(respRecorder, req)

	wordsFromRespJson, err := ioutil.ReadAll(respRecorder.Body)
	if err!=nil{
		t.Errorf("Ошибка чтения тела ответа: %s", err)
	}
	wordsFromResp := make([]string, 0)
	err = json.Unmarshal(wordsFromRespJson, &wordsFromResp)
	if err!=nil{
		t.Errorf("Ошибка расшифровки полученных данных: %s", err)
	}
	sort.Strings(expected)
	sort.Strings(wordsFromResp)
	if !reflect.DeepEqual(expected, wordsFromResp){
		t.Errorf("Полученный ответ не соответствует ожидаемому")
	}
}
func TestAnagramController_GetAnagrams_Should_Return_nil(t *testing.T) {
	wordRepository := repository.NewWordRepositorySlice()
	anagramController := NewAnagramController(wordRepository)

	req, err := http.NewRequest("GET", "/get?word=foobar", nil)

	if err!=nil {
		t.Fatal(err)
	}
	respRecorder := httptest.NewRecorder()
	getHandler := http.HandlerFunc(anagramController.GetAnagrams)

	getHandler.ServeHTTP(respRecorder, req)

	if respRecorder.Body.Len() != 0{
		t.Errorf("Контроллер вернул не nil, длина тела ответа: %v", respRecorder.Body.Len())
	}
	if respRecorder.Code  != http.StatusNotFound{
		t.Errorf("Контроллер должен вернуть 404, но вернул %v", respRecorder.Code)
	}
}

func TestAnagramController_GetAnagrams_Should_Return_422(t *testing.T) {
	wordRepository := repository.NewWordRepositorySlice()
	anagramController := NewAnagramController(wordRepository)
	//не передаётся слово
	req, err := http.NewRequest("GET", "/get", nil)
	if err!=nil {
		t.Fatal(err)
	}
	respRecorder := httptest.NewRecorder()
	getHandler := http.HandlerFunc(anagramController.GetAnagrams)

	getHandler.ServeHTTP(respRecorder, req)

	if respRecorder.Code!=http.StatusUnprocessableEntity{
		t.Errorf("Код ответа долже быть 422, но он: %v", respRecorder.Code)
	}
}

