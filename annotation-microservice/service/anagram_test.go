package service

import (
	"reflect"
	"sort"
	"testing"
)

type areWordsAnagramScenario struct{
	firstWord string
	secondWord string
	result bool
}

type getAnagramsByWordScenario struct{
	word string
	result []string
}

func TestAreWordsAnagram(t *testing.T) {
	testScenarios := [...]areWordsAnagramScenario{
		{"верно", "верно", true},
		{"аааб", "ааб", false},
		{"", "",true},
		{"aabb", "baba",true},
	}
	for _, scen := range testScenarios {
		if areStringsAnagram(scen.firstWord, scen.secondWord, false) != scen.result{
			t.Errorf("При словах: %s, %s должно быть: %t", scen.firstWord, scen.secondWord, scen.result)
		}
	}
}

func TestSplitSlice(t *testing.T){
	expected := [][]string{
		[]string{"0", "1", "2"},
		[]string{"3", "4", "5"},
		[]string{"6", "7", "8"},
		[]string{"9", "10", "11"},
		[]string{"12"},
	}
	slice := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
				      "10", "11", "12"}
	splitedSlice, err := splitSlice(slice, 5)
	/*expected := [][]string{
		[]string{"0"},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	}
	slice:=[]string{"0"}
	splitedSlice, err := splitSlice(slice, 5)*/
	if err!=nil{
		t.Error("Ошибки не должно быть!")
	}
	if !reflect.DeepEqual(expected, splitedSlice){
		t.Errorf("Ожидалось: %s, вернолось: %s", expected, splitedSlice)
	}
}
func TestSplitSlice_Should_Return_err(t *testing.T){
	_, err := splitSlice(nil, 0)
	if err==nil{
		t.Error("Ошибка должна быть!")
	}
}

func TestGetAnagramsByString(t *testing.T){
	testDictionary := []string{"foobar", "aabb", "baba", "boofar", "test"}
	testScenarios := [...]getAnagramsByWordScenario{
		{"foobar", []string{"foobar", "boofar"}},
		{"foobar", []string{"foobar", "boofar"}},
		{"raboof", []string{"foobar", "boofar"}},
		{"abba", []string{"aabb", "baba"}},
		{"test", []string{"test"}},
		{"qwerty", []string{}},
	}

	for _, scen := range testScenarios{
		result, err := GetAnagramsByString(scen.word, testDictionary, false)
		if err!=nil{
			t.Error("Не должно быть ошибки")
		}
		sort.Strings(result)
		sort.Strings(scen.result)
		if !reflect.DeepEqual(result, scen.result){
			t.Errorf("Ошибка возврата при слове: %s, ожидалось: %s, вернулось: %s}", scen.word, scen.result, result)
		}
	}
}
func TestGetAnagramsByString_Should_Return_err(t *testing.T) {
	testDictionary := []string{"foobar", "aabb", "baba", "boofar", "test"}
	emptyWord := ""
	_, err := GetAnagramsByString(emptyWord, testDictionary, false)
	if err==nil{
		t.Error("Должна вернуться ошибка!")
	}
}


func BenchmarkGetAnagramsByString(b *testing.B) {
	testScenarios := make([]string, 0)
	for i:=0; i <= 1000000; i++{
		testScenarios = append(testScenarios, "test")
	}
	GetAnagramsByString("not test", testScenarios, false)
}

