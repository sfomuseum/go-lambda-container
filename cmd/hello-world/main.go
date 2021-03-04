package main

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sfomuseum/go-flags/flagset"
	"io"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"	
)

func ReadFile(ctx context.Context, bucket *blob.Bucket, path string) (string, error) {

	fh, err := bucket.NewReader(ctx, path, nil)

	if err != nil {
		return "", err
	}

	defer fh.Close()
	
	body, err := io.ReadAll(fh)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {

	fs := flagset.NewFlagSet("sfomuseum")

	bucket_uri := fs.String("bucket-uri", "", "file:///")
	mode := fs.String("mode", "cli", "...")

	flagset.Parse(fs)

	ctx := context.Background()

	err := flagset.SetFlagsFromEnvVars(fs, "SFOMUSEUM")

	if err != nil {
		log.Fatalf("Failed to assign flags from environment variables, %v", err)
	}

	bucket, err := blob.NewBucket(ctx, *bucket_uri)

	if err != nil {
		log.Fatalf("Failed to create bucket for '%s', %v", *bucket_uri, err)
	}
	
	switch *mode {
	case "cli":

		args := flag.Args()

		if len(args) == 0 {
			log.Fatal("Missing path to read")
		}
		
		path := args[0]
		
		str, err := ReadFile(ctx, bucket, path)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(str)
		
	case "lambda":

		handler := func(ctx) (string, error){
			return ReadFile(ctx, bucket, path)
		}
		
		lambda.Start(handler)
	default:
		log.Fatalf("Invalid mode '%s'", *mode)
	}
}
	
	
