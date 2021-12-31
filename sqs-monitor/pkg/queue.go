package pkg

import "github.com/aws/aws-sdk-go/service/cloudwatch"

type Queue struct {
	QueueName        string
	ApplicationCount int
}

// GetMessage count takes a cloudwatch client an a list of queues
// and returns a Queue struct with queue name and application count
func GetMessageCount(cw *cloudwatch.Client, queueList []string) (Queue, error) {
	return Queue{}, nil
}

// getQueueMetric takes the name of the queue and using cloudwatch API retrieves
// the NumberOfMessagesReceived metric, which in turn returns a count of messages
// by queue.
func getQueueMetric(cw *cloudwatch.Client, queue string) int {
	return 0
}

// calculateNumberOfApps arbitrarily calculates the number of apps connected to each queue
func calculateNumberOfApps(messageCount int) int {
	return 0
}
