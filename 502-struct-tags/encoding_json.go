package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	F1 string `json:"filed1"`
	F2 string `json:"filed2,omitempty"`
	F3 string `json:"filed3,omitempty"`
	F4 string `json:"filed4,omitempty"`
}

func encodingJson() {
	fmt.Println("Struc Tags : encoding json")

	t := T{
		F1: "v1",
		F2: "",
		F3: "v3",
		F4: "v4",
	}

	s, _ := json.Marshal(t)
	fmt.Printf("%s \n ", s)
}
