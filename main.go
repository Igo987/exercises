package main

import (
	"flag"
	"fmt"

	masker "github.com/Igo87/project/masker"
)

func main() {
	path_read := flag.String("path_read", "./src/links.txt", "path to write file")
	path_write := flag.String("path_write", "./src/links.txt", "path to read file")
	flag.Parse()
	somePres := masker.NewPresent()
	somePres.Path = *path_read
	someProd := masker.NewProduce()
	someProd.Path = *path_write
	newService := masker.NewService(somePres, someProd)

	if err := newService.Run(); err != nil {
		err = fmt.Errorf("an error occurred when starting the service: %s", err)
		fmt.Println(err)
	}

}
