package main

import (
	"encoding/json"
	"fmt"

	"github.com/digitalhouse-dev/dh-kit/meta"

	"github.com/digitalhouse-dev/dh-kit/response"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
	Age       int    `json:"age"`
}

func main() {
	u := &User{
		FirstName: "Nahuel",
		LastName:  "Costamagna",
		Mail:      "nlcostamagna@gmail.com",
		Age:       32,
	}

	print(response.OK("", u, nil, nil))

	print(response.OK("sarasa", u, nil, nil))

	m := meta.New(1, 30, 100)
	print(response.OK("sarasa", u, m, nil))

	print(response.Created("", u, nil, nil))
	print(response.Accepted("", u, nil, nil))

	print(response.BadRequest("my error"))
	print(response.NotFound("my error"))
	print(response.InternalServerError("my error"))

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
