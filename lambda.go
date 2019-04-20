package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	DomainName string `json:"domainName"`
	Protocal   string `json:"protocol"`
	Port       int    `json:"port"`
}

type Response struct {
	DNSEntries []string `json:"dnsEntries"`
	Reachable  bool     `json:"reachable"`
}

func Handler(ctx context.Context, request Request) (Response, error) {
	response := Response{}
	if ips, err := net.LookupIP(request.DomainName); err == nil {
		var dnsEntries []string
		for _, ip := range ips {
			dnsEntries = append(dnsEntries, fmt.Sprintf("%s IN A %s", request.DomainName, ip))
		}
		response.DNSEntries = dnsEntries
	}

	if _, err := net.DialTimeout(request.Protocal, fmt.Sprintf("%s:%d", request.DomainName, request.Port), time.Duration(500*time.Millisecond)); err == nil {
		response.Reachable = true
	} else {
		response.Reachable = false
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
