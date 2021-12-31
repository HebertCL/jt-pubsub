package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	queue "github.com/HebertCL/jt-pubsub/sqs-monitor/pkg"
)

// Func main is the entrypoint.
// main will consume a list of queues and from there it will
// follow the logic: per queue, get a number of messages, calculate
// the amount of apps sending messages, return a result
func main() {
	// TODO: Avoid hardcoding profile
	queues := []string{"queue", "names", "here"}
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "localstack",
	})
	if err != nil {
		log.Fatalf("Unable to fetch AWS Credentials: %v", err)
	}

	svc := cloudwatch.New(sess)
	output, err := queue.GetMessageCount(svc, queues)
	if err != nil {
		log.Fatalf("An unexpected error occurred: %v", err)
	}

	// Some printing here to output
}
