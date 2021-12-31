# SQS module
module "queue" {
  source = "../../modules/aws-sqs"

  queue_name = var.queue_name
}

data "aws_sqs_queue" "queue" {
  depends_on = [
    module.queue
  ]

  name = var.queue_name
}

module "topic" {
  source = "../../modules/aws-sns"
  depends_on = [
    module.queue
  ]

  topic_name     = var.topic_name
  topic_endpoint = data.aws_sqs_queue.queue.arn
}