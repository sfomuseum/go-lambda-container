package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"time"
)

func HelloWorld(ctx context.Context) (string, error) {
	str := fmt.Sprintf("Hello world, %v", time.Now())
	return str, nil
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

		str, err := HelloWorld(ctx)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(str)

	case "lambda":
		lambda.Start(HelloWorld)
	default:
		log.Fatalf("Invalid mode '%s'", *mode)
	}
}
