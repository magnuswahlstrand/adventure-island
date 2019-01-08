package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type resource struct {
	path, name string
}

func main() {
	resources := []resource{
		resource{
			path: "resources/all.png",
			name: "All_png",
		},
		resource{
			path: "resources/character.png",
			name: "Character_png",
		},
		resource{
			path: "resources/character_color.png",
			name: "Character_color_png",
		},
		resource{
			path: "resources/objects.png",
			name: "Objects_png",
		},
	}

	f, err := os.Create("resources/resources.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package resources\n")
	for _, r := range resources {
		content, err := ioutil.ReadFile(r.path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(f, "var %s    = []byte(%q)\n", r.name, string(content))
	}
}
