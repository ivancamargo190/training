// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/server"
	"github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/server/restapi"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(server.SwaggerJSON, server.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := restapi.NewUserAPI(swaggerSpec)
	server := server.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "User"
	parser.LongDescription = "This is the User Microservice, responsible for managing Users in Modern Bank."

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
