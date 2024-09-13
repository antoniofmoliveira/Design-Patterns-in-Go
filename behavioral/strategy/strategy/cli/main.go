package main

import (
	"flag"
	"log"
	"os"

	"srcgo/github.com/antoniofmoliveira/patterns/behavioral/strategy/strategy/shapes"
)

var output = flag.String("output", "text", "The output to use between 'text' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.Factory(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}
	err = activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
