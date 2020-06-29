#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

function haltOnFailure {
	if [ $? != 0 ]
	then
		echo "$1 failed"
		exit
	fi
}

echo "Deleting build artifacts..."
rm -v lambda lambda.zip

echo "Building Go function..."
go build -v -ldflags="-s -w" -o lambda lambda.go
haltOnFailure "go build"

echo "Setting permissions..."
chmod uga+x lambda
haltOnFailure "chmod"

echo "Preparing zip..."
zip --dot-size 1024k lambda.zip lambda
haltOnFailure "zip"

echo "Uploading to AWS..."
aws lambda update-function-code --function-name IoT-Button_send-text-message --zip-file fileb://lambda.zip
haltOnFailure "aws"

echo "Complete"
