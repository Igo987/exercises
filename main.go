package main

import (
	"flag"
	"fmt"

	masker "github.com/Igo87/project/pkg"
)

func main() {
	path_read := flag.String("path_read", "./src/links.txt", "path to write file")
	path_write := flag.String("path_write", "./src/links.txt", "path to read file")
	flag.Parse()
	someProd := masker.NewProduce()
	somePres := masker.NewPresent()
	newService := masker.NewService(someProd, somePres)
	err := newService.Run(*path_read, *path_write)
	if err != nil {
		err = fmt.Errorf("an error occurred when starting the service: %s", err)
		fmt.Println(err)
	}

}
