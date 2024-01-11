package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type myStruct struct {
	ID       int
	Number   string
	Year     int
	Students []struct {
		LastName   string
		FirstName  string `json:"FirstName"`
		MiddleName string `json:"MiddleName"`
		Birthday   string `json:"Birthday"`
		Address    string `json:"Address"`
		Phone      string `json:"Phone"`
		Rating     []int
	}
}

func f() {

	file, err := ioutil.ReadFile("/Users/basty64/Programming/Projects/GolangProjects/GolangBaza/simple_algorithms/test.json")
	if err != nil {
		fmt.Println(err)
	}

	s := new(myStruct)

	if err := json.Unmarshal(file, s); err != nil {
		fmt.Println(err)
		return
	}
	students_amount := 0
	marks_amount := 0
	for i, students := range s.Students {
		marks_amount += len(students.Rating)
		students_amount = i + 1
	}
	var aver float64
	aver = float64(marks_amount) / float64(students_amount)
	file, _ = json.MarshalIndent(map[string]float64{"Average": aver}, "", "    ")
	fmt.Printf("%s", file)

}
