package controller

import (
	"annotation-microservice/model"
	"annotation-microservice/repository"
	"annotation-microservice/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func castWordsToStrings(words []model.Word)[]string{
	stringWords:=make([]string, 0)
	for _,word := range words{
		stringWords = append(stringWords, word.Text)
	}
	return stringWords
}
func castStringsToWords(stringWords []string)[]model.Word{
	words := make([]model.Word, 0)
	for _, str:= range stringWords{
		words = append(words, model.Word{Text: str})
	}
	return words
}

type AnagramController struct{
	wordRepository repository.WordRepository
}

func NewAnagramController(wordRepo repository.WordRepository) *AnagramController{
	return &AnagramController{wordRepository: wordRepo}
}


func(c AnagramController) Load(w http.ResponseWriter, r *http.Request){
	if r.Method!=http.MethodPost{
		log.Println("Метод не POST")
		http.Error(w, "Неподдерживаемый метод, должен быть POST", http.StatusMethodNotAllowed)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Ошибка чтения запроса", http.StatusUnprocessableEntity)
	}

	stringsWords := make([]string, 0)
	err = json.Unmarshal(body, &stringsWords)
	if err!=nil{
		log.Print(err.Error())
		http.Error(w, "Данные переданны не в верном представлении", http.StatusUnprocessableEntity)
	}

	c.wordRepository.RemoveAll()

	err = c.wordRepository.AddRange(castStringsToWords(stringsWords))
	if err!=nil{
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	log.Printf("Было загружено: %v записей!", len(stringsWords))
	w.WriteHeader(200)
}
func(c AnagramController) GetAnagrams(w http.ResponseWriter, r *http.Request){
	if r.Method!=http.MethodGet{
		log.Print("Метод не GET")
		http.Error(w, "Неподдерживаемый метод, должен быть GET", http.StatusMethodNotAllowed)
	}
	word := r.URL.Query().Get("word")
	if len(word)==0{
		log.Print("Не передано слово для поиска!")
		http.Error(w, "Не передано слово для поиска!", http.StatusUnprocessableEntity)
	}
	words, err := c.wordRepository.Index()
	if err!=nil{
		log.Println("Ошибка получения слов err:", err)
		http.Error(w, "Ошибка получения слов", http.StatusInternalServerError)
	}
	anagrams, _ := service.GetAnagramsByString(word, castWordsToStrings(words), false)
	log.Printf("Поиск по слову %s", word)

	if len(anagrams)==0{
		log.Printf("Не найдено анаграм для слова: %v", word)
		w.WriteHeader(http.StatusNotFound)
	}else{
		anagramsJson, _ := json.Marshal(anagrams)
		w.Write(anagramsJson)
	}
}