package main

import (
	"fmt"
	"log"
	"time"

	"github.com/arschles/gobuilder"
)

type user struct {
	Name         string
	EmailAddress string
	Markdown     string
	CreatedAt    time.Time
	Admin        bool
	PasswordHash string
	Author       *user
}

func (u user) ToBuilder() gobuilder.Builder {
	return gobuilder.NewBuilder("user", u)
}

func main() {
	u := user{Markdown: "this is some markdown!", CreatedAt: time.Now()}
	res, err := u.ToBuilder().Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
