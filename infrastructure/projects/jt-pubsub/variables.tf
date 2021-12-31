# Queue values
variable "queue_name" {
  description = "SQS Queue name"
  type        = string
}

# Topic values
variable "topic_name" {
  description = "SNS Topic name"
  type        = string
}

# Resource tagging
variable "environment" {
  description = "Environment name"
  type        = string
}

variable "project" {
  description = "Project name"
  type        = string
}