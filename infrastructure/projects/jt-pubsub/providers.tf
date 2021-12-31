provider "aws" {
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  # access_key = "foo"
  # secret_key = "bar"
  profile = "localstack"
  region  = "us-east-1"

  endpoints {
    sns = "http://localhost:4566"
    sqs = "http://localhost:4566"
  }

  default_tags {
    tags = {
      Environment = var.environment
      Project     = var.project
    }
  }
}
