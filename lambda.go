package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	DomainName string `json:"domainName"`
	Protocal   string `json:"protocol"`
	Port       string `json:"port"`
}

type Response struct {
	DNSEntries []string `json:"dnsEntries"`
	Reachable  bool     `json:"reachable"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input Request
	if err := json.Unmarshal([]byte(request.Body), &input); err == nil {
		response := &Response{}
		if ips, err := net.LookupIP(input.DomainName); err == nil {
			var dnsEntries []string
			for _, ip := range ips {
				dnsEntries = append(dnsEntries, fmt.Sprintf("%s IN A %s", input.DomainName, ip))
			}
			response.DNSEntries = dnsEntries
		}

		if _, err := net.DialTimeout(input.Protocal, fmt.Sprintf("%s:%d", input.DomainName, input.Port), time.Duration(500*time.Millisecond)); err == nil {
			response.Reachable = true
		} else {
			response.Reachable = false
		}

		if jsonBody, err := json.Marshal(response); err == nil {
			return events.APIGatewayProxyResponse{Body: string(jsonBody), StatusCode: 200}, err
		}
	} else {
		log.Printf("%v", err)
	}

	return events.APIGatewayProxyResponse{StatusCode: 500}, nil
}

func main() {
	lambda.Start(Handler)
}
