package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Struct []struct {
	Id int `json:"global_id"`
}

func t() {
	//b := Struct{}
	s := new(Struct)
	file, err := ioutil.ReadFile("/Users/basty64/Programming/Projects/GolangProjects/GolangBaza/simple_algorithms/data-20190514T0100.json")
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(file, &s)
	var n int
	for _, ids := range *s {
		n += ids.Id
	}
	fmt.Println(n)
}
