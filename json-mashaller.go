package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type user struct {
	First, Last string
	Age         int
}

type userjson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type byAge []user

func (u byAge) Len() int {
	return len(u)
}

func (u byAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u byAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func run() {
	u1 := user{
		First: "Andrew",
		Last:  "Mititi",
		Age:   26,
	}

	u2 := user{
		First: "Rael",
		Last:  "Gesare",
		Age:   1,
	}

	users_array := []user{u1, u2}

	_, err := json.Marshal(users_array)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(os.Stdout).Encode(u1)

	users_json_data := `[{"First":"Andrew","Last":"Mititi","Age":26},{"First":"Rael","Last":"Gesare","Age":1}]`
	d := []userjson{}

	err = json.Unmarshal([]byte(users_json_data), &d)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(byAge(users_array))
	fmt.Println(users_array)
}
