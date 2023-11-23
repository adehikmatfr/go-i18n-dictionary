package main

import (
	"fmt"

	dictionary "github.com/adehikmatfr/go-i18n-dictionary"
)

func main() {
	dic := dictionary.New()
	fmt.Println(dic.Get(dictionary.Hallo))
	dic.SetLang(dictionary.JP)
	fmt.Println(dic.Get(dictionary.Hallo))
	fmt.Println(dic.Format(dictionary.HalloNamef, "ade"))
}
