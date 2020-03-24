package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func req(ctx context.Context, s3Event events.S3Event) {
	rundeckJob := os.Getenv("RUNDECK_JOB_URL")

	for _, r := range s3Event.Records {
		s3 := r.S3
		reqBody, err := json.Marshal(map[string]string{
			"bucket":   s3.Bucket.Name,
			"filePath": s3.Object.Key,
		})
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post(rundeckJob, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
	}

}

func main() {
	lambda.Start(req)
}
