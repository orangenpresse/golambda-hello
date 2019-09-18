#!/bin/bash
GOOS=linux go build main.go
zip -r handler.zip ./main

###############
# CREATE CODE
###############
# aws lambda create-function \
#  --region eu-central-1 \
#  --function-name hello-world \
#  --memory 512 \
#  --role [ROLE] \
#  --runtime go1.x \
#  --zip-file fileb://./handler.zip \
#  --profile [PROFILE] \
#  --handler main

###############
# UPDATE CODE
###############
aws lambda update-function-code \
    --region eu-central-1 \
    --function-name hello-world \
    --zip-file fileb://./handler.zip \
    --profile [PROFILE]

rm main
rm handler.zip
