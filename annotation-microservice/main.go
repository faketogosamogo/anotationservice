package main

import (
	"annotation-microservice/controller"
	 "annotation-microservice/repository/slice"
	"net/http"
)

func main(){
	wordRepository:=repository.WordRepositorySlice{}
	anagramController:=controller.NewAnagramController(&wordRepository)

	http.HandleFunc("/load", anagramController.Load)
	http.HandleFunc("/get", anagramController.GetAnagrams)
	http.ListenAndServe(":8080", nil)
}
