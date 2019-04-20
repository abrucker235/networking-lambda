#### lambda.go
Lambda function that takes json object as input.

#### apigw_lambda.go
Lambda function that takes `APIGatewayProxyRequest`

#### Build Uploadable Zip
```
GOOS=linux go build && zip -o networking-lambda.zip ./networking-lambda
```