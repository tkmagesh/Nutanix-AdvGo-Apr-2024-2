package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"-"`
}

func main() {
	// Serialization
	p := Product{100, "Pen", 10}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(p); err != nil {
		fmt.Println("serailization error :", err)
	}

	// Deserialization
	r := strings.NewReader("{\"id\":100,\"name\":\"pencil\"}")
	decoder := json.NewDecoder(r)
	var p2 Product
	if err := decoder.Decode(&p2); err != nil {
		fmt.Println("deserailization error :", err)
	} else {
		fmt.Printf("%#v\n", p2)
	}
}
