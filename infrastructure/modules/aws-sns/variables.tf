variable "topic_name" {
  description = "SNS Topic name"
  type = string
}

variable "topic_endpoint" {
  description = "The endpoint for SNS subscription to a queue"
  type = string
}
