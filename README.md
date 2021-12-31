# jt-PubSub task

This repo contains all of the task related deliverables for the task.

- Completed: SNS/SQS Terraform modules using Localstack with output.
- WIP: sqs-monitor standalone app in Go

The application development took way more than the 1 hour I have destined to it, so what you will find there is an empty shell-like structure explaining briefly the idea on how to get the application built.
I used the approach of decoupling everything into a pkg folder and place the the application logic with its respective tests. The development of the application also raised me the question on the calculation
of apps connected to the queue. In this case, I wondered what kind of logic should you expect to see as this is something I haven't done before. Other than that, I would gladly provide more input on both tasks via call.