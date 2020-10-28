package service

import (
	"errors"
	"math"
	"strings"
)

//Первые два аргумента слова для проверки на анаграмму, caseSensitive учитывать ли регистр слов
func areStringsAnagram(first, second string, caseSensitive bool) bool{
	if len(first)!=len(second){
		return false
	}
	if caseSensitive{
		first = strings.ToLower(first)
		second = strings.ToLower(second)
	}
	for _, letter:= range first{
		if strings.Count(first, string(letter)) != strings.Count(second, string(letter)){
			return false
		}
	}
	return true
}
//Получения анаграм для слова
//word-слово для поиска
//from - словарь со словами
//caseSensitive - учитывать ли регистр
func GetAnagramsByString(word string, from []string, caseSensitive bool)([]string, error){

	if len(word)==0{
		return nil, errors.New("cлово для поиска не должно быть пустым")
	}
	/*anagrams:= make([]string, 0)
	for _, val:= range from{
		if areStringsAnagram(word, val, caseSensitive){
			anagrams = append(anagrams)
		}
	}
	*/
	anagrams:= make([]string, 0)

	countOfParts := 5
	splitedFrom, err := splitSlice(from, countOfParts)

	channels := make(chan []string)

	if err!=nil{
		return nil, err
	}
	for i:=0; i < len(splitedFrom); i++{
		go func(from []string, ch chan []string) {

			anagram := make([]string, 0)
			for _, val := range from{
				if areStringsAnagram(word, val, caseSensitive){
					anagram = append(anagram, val)
				}
			}
			ch <- anagram
		}(splitedFrom[i], channels)
	}
	for i:=0; i < countOfParts; i++{
		anagrams = append(anagrams, <-channels...)
	}
	return anagrams, nil
}

//Возвращает разделенный slice
func splitSlice(slice []string, countOfParts int)([][]string, error){
	if countOfParts<1{
		return nil, errors.New("количество частей не может быть меньше 1")
	}
	lenOfPart := int(math.Ceil(float64(len(slice))/float64(countOfParts)))
	parts :=  make([][]string, countOfParts)

	left:=0
	right:= lenOfPart

	for i:=0; i < countOfParts; i++{
		if left >= len(slice){
			parts[i] = make([]string, 0)
		}else{
			if right > len(slice){
				right = len(slice)
			}
			parts[i] = slice[left : right]
			left += lenOfPart
			right += lenOfPart
		}
	}
	return parts, nil
}
