package main

import (
	"context"
	"log"

	"github.com/leon-devs/course-docker_intro/basics/pkg/app"
)

func main() {
	ctx := context.Background()

	application, err := app.NewApplication(ctx)
	if err != nil {
		log.Fatalln("Error getting application: ", err)
	}

	log.Println("About to start application ...")

	if err = application.Start(ctx); err != nil {
		log.Fatalln("Error starting application: ", err)
	}
}
