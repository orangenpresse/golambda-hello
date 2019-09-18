package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (interface{}, error) {
	name := request.PathParameters["name"]
	returnString := "Hello World!"

	if name != "" {
		returnString = fmt.Sprintf("Hello %s!", name)
	}

	response := events.APIGatewayProxyResponse{}
	headers := make(map[string]string)
	headers["Content-Type"] = "text/plain; charset=utf8"

	response.IsBase64Encoded = false
	response.StatusCode = 200
	response.Headers = headers
	response.Body = returnString

	return response, nil
}
