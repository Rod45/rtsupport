package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

type User struct {
	Id   string `gorethink:"id"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
