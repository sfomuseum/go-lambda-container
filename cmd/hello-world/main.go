package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"log"
	"os"
	"time"
)

func HelloWorld(ctx context.Context) (string, error) {
	str := fmt.Sprintf("Hello world, %v", time.Now())
	return str, nil
}

func main() {

	fs := flagset.NewFlagSet("sfomuseum")

	mode := fs.String("mode", "cli", "Valid modes are: cli (command line), lambda.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Emit the phrase 'Hello world' and the current time.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		fs.PrintDefaults()
	}

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
