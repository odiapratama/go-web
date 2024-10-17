package main

import (
	"log"
	"os"
	"strconv"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func (s semester) TotalUnits() int {
	total := 0
	for _, c := range s.Courses {
		units, err := strconv.Atoi(c.Units)
		if err != nil {
			log.Fatalln(err)
		}
		total += units
	}
	return total
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"CSCI-40", "Introduction to Computer Programming", "4"},
				{"CSCI-130", "Introduction to Web Development", "3"},
				{"CSCI-140", "Introduction to UNIX", "1"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"CSCI-50", "Advanced Computer Programming", "4"},
				{"CSCI-150", "Advanced Web Development", "3"},
				{"CSCI-160", "Advanced UNIX", "1"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{"CSCI-60", "Database Management", "4"},
				{"CSCI-170", "Database Design", "3"},
				{"CSCI-180", "Database Administration", "1"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}
