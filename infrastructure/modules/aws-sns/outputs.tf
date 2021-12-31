output "topic_id" {
  value = aws_sns_topic.this.id
}

output "topic_arn" {
  value = aws_sns_topic.this.arn
}

output "subscription_id" {
  value = aws_sns_topic_subscription.this.id
}