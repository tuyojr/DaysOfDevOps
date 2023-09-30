terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "5.6.2"
        }
    }
}

provider "aws" {
    region = "us-east-1"
}
