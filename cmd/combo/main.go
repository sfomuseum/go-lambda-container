package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aaronland/go-http-server"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"net/http"
)

func combo() string {
	return "Hello world"
}

func main() {

	fs := flagset.NewFlagSet("combo")

	mode := fs.String("mode", "cli", "Valid options are: cli, lambda, server")
	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI string")

	flagset.Parse(fs)

	ctx := context.Background()
	
	switch *mode {
	case "cli":
		fmt.Println(combo())
	case "lambda":
		lambda.Start(combo)
	case "server":

		fn := func(rsp http.ResponseWriter, req *http.Request) {

			str := combo()

			enc := json.NewEncoder(rsp)
			err := enc.Encode([]byte(str))

			if err != nil {
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		handler := http.HandlerFunc(fn)

		mux := http.NewServeMux()
		mux.Handle("/", handler)

		s, err := server.NewServer(ctx, *server_uri)

		if err != nil {
			log.Fatalf("Failed to create server for '%s', %v", server_uri, err)
		}

		log.Printf("Listening for requests at %s\n", s.Address())

		err = s.ListenAndServe(ctx, mux)

		if err != nil {
			log.Fatalf("Failed to serve requests for '%s', %v", server_uri, err)
		}

	default:
		log.Fatalf("Unsupported mode '%s'", *mode)
	}
}
