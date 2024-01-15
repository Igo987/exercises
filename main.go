package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	masker "github.com/Igo87/project/masker"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go masker.Stop(ctx)

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
