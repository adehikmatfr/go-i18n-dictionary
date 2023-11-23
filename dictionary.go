package dictionary

import (
	"fmt"
	"strings"
)

type Lang string
type Keyword string

const (
	EN Lang = "en"
	JP Lang = "jp"
)

type Module interface {
	SetLang(l Lang)
	SetWords(w map[Lang]map[Keyword]string)
	Format(k Keyword, formatParams ...interface{}) string
	Get(k Keyword) string
}

type Dictionary struct {
	Lang  Lang
	Words map[Lang]map[Keyword]string
}

func New() Module {
	return &Dictionary{
		Lang:  EN,
		Words: Words,
	}
}

func (d *Dictionary) getDictionary(l Lang) map[Keyword]string {
	dic, ok := d.Words[d.Lang]
	if !ok {
		dic = d.Words[EN]
	}

	return dic
}

func (d *Dictionary) SetLang(l Lang) {
	d.Lang = l
}

func (d *Dictionary) SetWords(w map[Lang]map[Keyword]string) {
	d.Words = w
}

func (d *Dictionary) Format(k Keyword, formatParams ...interface{}) string {
	dic := d.getDictionary(d.Lang)
	format := dic[k]

	for index, param := range formatParams {
		format = strings.ReplaceAll(format, fmt.Sprintf("{%d}", index), fmt.Sprintf("%v", param))
	}

	return format
}

func (d *Dictionary) Get(k Keyword) string {
	dic := d.getDictionary(d.Lang)
	return dic[k]
}
