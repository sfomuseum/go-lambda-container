package main

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"	
)

func HelloWorld(ctx context.Context) error {
	fmt.Println("Hello world")
	return nil
}

func main() {

	fs := flagset.NewFlagSet("sfomuseum")

	mode := fs.String("mode", "cli", "...")

	flagset.Parse(fs)

	ctx := context.Background()

	err := flagset.SetFlagsFromEnvVars(fs, "SFOMUSEUM")

	if err != nil {
		log.Fatalf("Failed to assign flags from environment variables, %v", err)
	}

	switch *mode {
	case "cli":
		HelloWorld(ctx)
	case "lambda":
		lambda.Start(HelloWorld)
	default:
		log.Fatalf("Invalid mode '%s'", *mode)
	}
}
	
	
