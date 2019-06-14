package findmax

import (
	"reflect"
	"testing"
)

func TestFindMax(t *testing.T) {
	sl := []int{1, 40, 2, 3, 4, 5, 6, 7}

	res := FindMax(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	if reflect.ValueOf(res).Int() != 40 {
		t.Errorf("The maximum value should be %d", 40)
	}

	words := []string{"garb", "strict", "assigment", "singling"}
	res = FindMax(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	if reflect.ValueOf(res).String() != "assigment" {
		t.Errorf("The maximum value should be %s", "assigment")
	}
	type Book struct {
		Name  string
		Count int
	}
	books := []Book{
		{"book1", 3},
		{"book2", 17},
		{"book3", 30},
		{"book4", 50},
		{"book5", 20},
	}

	res = FindMax(books, func(i, j int) bool {
		return books[i].Count < books[j].Count
	})

	s := reflect.ValueOf(res)
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if typeOfT.Field(i).Name == "Count" {
			if f.Int() != 50 {
				t.Errorf("The maximum value should be %s", "assigment")
			}
		}
	}
}
