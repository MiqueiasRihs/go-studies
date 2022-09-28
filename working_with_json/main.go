package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar string
}

// type Foo struct {
// 	ean      string
// 	name     string
// 	fullname string
// 	price    string
// 	picture  string
// }

func main() {
	foo1 := new(Foo)
	getJson("https://www.obrazul.com.br/api/recruitment/products/", foo1)
	println("aaaaaaaaaa")
	println(foo1.Bar)
}
