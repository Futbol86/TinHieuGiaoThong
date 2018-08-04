package main

import (
	"fmt"
	// "fullstackweb/lib"
	 "github.com/Futbol86/TinHieuGiaoThong/lib/languages.go"
	//"github.com/shijuvar/go-samples-thenewstack/packagedemo/lib"
)

func main() {
	lib.Add("dr","Dart")
	fmt.Println(lib.Get("dr"))
	languages:=lib.GetAll()
	for _, v := range languages {
		fmt.Println(v)
	}
}
