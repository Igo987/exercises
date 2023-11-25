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
	Service := masker.Run(*path_read, *path_write)
	fmt.Println(Service)

}
