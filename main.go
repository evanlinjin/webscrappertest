package main

import (
	"fmt"
	"log"
	"flag"
	"golang.org/x/net/html"
	"os"
)

func main() {
	fStr := flag.String("f", "test.html", "file name")
	flag.Parse()

	f, e := os.Open(*fStr)
	if e != nil {
		log.Println(e)
		return
	}

	var name, email string

	z := html.NewTokenizer(f)
	for {
		switch z.Next() {
		case html.StartTagToken:
			t := z.Token()

			if t.Data != "span" {
				continue
			}

			var yes bool
			yes, name, email = AttrsRepresentGD(t.Attr)
			if yes {
				goto finish
			}

		case html.ErrorToken:
			return
		}
	}

finish:
	fmt.Println("SENDER EMAIL:", email)
	fmt.Println("SENDER NAME:", name)
}

func AttrsRepresentGD(attrs []html.Attribute) (yes bool, name, email string) {
	for _, a := range attrs {
		switch a.Key {
		case "class":
			if a.Val != "gD" {
				return
			}
			yes = true
		case "email":
			email = a.Val
		case "name":
			name = a.Val
		}
	}
	return
}