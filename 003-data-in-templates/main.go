package main

import (
	"log"
	"os"
	"text/template"
)

type Passanger struct {
	Name, Surname string
	Age           int
	TransportId   string
}

type Transport struct {
	Id, Type, Model string
}

type Tickets struct {
	Passangers []Passanger
	Transports []Transport
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	passangers := []Passanger{
		{"John", "Doe", 25, "A123"},
		{"Jane", "Doe", 24, "B123"},
		{"Alice", "Doe", 23, "C123"},
	}

	transports := []Transport{
		{"A123", "Car", "BMW"},
		{"B123", "Plane", "Boeing"},
		{"C123", "Train", "TGV"},
	}

	tickets := Tickets{
		Passangers: passangers,
		Transports: transports,
	}

	err := tpl.Execute(os.Stdout, tickets)
	if err != nil {
		log.Fatalln(err)
	}
}
