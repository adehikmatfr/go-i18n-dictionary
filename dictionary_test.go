package dictionary

import (
	"reflect"
	"testing"
)

func TestSetLang(t *testing.T) {
	d := New().(*Dictionary)

	d.SetLang(JP)

	if d.Lang != JP {
		t.Errorf("Expected language %v, got %v", JP, d.Lang)
	}
}

var (
	HelloKeyword Keyword = "hallo_keyword"
	h            string  = "Hello, {0}!"
	hJP          string  = "こんにちは、{0}!"
	errf         string  = "Expected %v, got %v"
	words                = map[Lang]map[Keyword]string{
		EN: {
			HelloKeyword: h,
		},
		JP: {
			HelloKeyword: hJP,
		},
	}
)

func TestSetWords(t *testing.T) {
	d := New().(*Dictionary)

	d.SetWords(words)

	if reflect.TypeOf(d.Words) != reflect.TypeOf(words) {
		t.Errorf("Expected words %v, got %v", words, d.Words)
	}
}

func TestFormat(t *testing.T) {
	d := New().(*Dictionary)

	d.SetWords(words)

	result := d.Format(HelloKeyword, "John")

	expectedEN := "Hello, John!"
	if result != expectedEN {
		t.Errorf(errf, expectedEN, result)
	}

	d.SetLang(JP)
	resultJP := d.Format(HelloKeyword, "John")

	expectedJP := "こんにちは、John!"
	if resultJP != expectedJP {
		t.Errorf(errf, expectedJP, resultJP)
	}
}

func TestGet(t *testing.T) {
	d := New().(*Dictionary)

	d.SetWords(words)

	result := d.Get(HelloKeyword)

	expectedEN := h
	if result != expectedEN {
		t.Errorf(errf, expectedEN, result)
	}

	d.SetLang(JP)
	resultJP := d.Get(HelloKeyword)

	expectedJP := hJP
	if resultJP != expectedJP {
		t.Errorf(errf, expectedJP, resultJP)
	}
}
