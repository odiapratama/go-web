package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

type code struct {
	Code    int
	Descrip string
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(p2)

	var data []code

	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err2 := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err2)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
}
