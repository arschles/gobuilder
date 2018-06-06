package main

import (
	"fmt"
	"log"
	"time"

	"github.com/arschles/gobuilder"
)

type user struct {
	Markdown  string
	CreatedAt time.Time
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
