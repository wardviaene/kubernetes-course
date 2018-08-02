#!/bin/bash

set -e

RANDOM_STRING=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 8 | tr '[:upper:]' '[:lower:]' | head -n 1)

# it's important to set the AWS_REGION if not set. Change the default
DEFAULT_REGION="eu-west-1"
AWS_REGION="${AWS_REGION:-${DEFAULT_REGION}}"

if [ "$AWS_REGION" == "us-east-1" ] ; then
  aws s3api create-bucket --bucket helm-${RANDOM_STRING}
else
  aws s3api create-bucket --bucket helm-${RANDOM_STRING} --region $AWS_REGION --create-bucket-configuration LocationConstraint=${AWS_REGION}
fi

helm plugin install https://github.com/hypnoglow/helm-s3.git
helm s3 init s3://helm-${RANDOM_STRING}/charts
helm repo add my-charts s3://helm-lcxqoc24/charts
