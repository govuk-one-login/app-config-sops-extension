package main

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsops/sops/v3/cmd/sops/formats"
	"github.com/getsops/sops/v3/decrypt"
)

type SopsExtensionRequest struct {
	Content    string `json:"Content"`
	Parameters struct {
		Format string `json:"format"`
	} `json:"Parameters"`
}

type SopsExtensionResponse struct {
	Content string                  `json:"Content,omitempty"`
	Error   string                  `json:"Error,omitempty"`
	Message string                  `json:"Message,omitempty"`
	Details []AppConfigErrorDetails `json:"Details,omitempty"`
}

type AppConfigErrorDetails struct {
	Type   string `json:"Type"`
	Name   string `json:"Name"`
	Reason string `json:"Reason"`
}

func errorResponse(err error, message string) SopsExtensionResponse {
	return SopsExtensionResponse{
		Error:   message,
		Message: err.Error(),
		Details: []AppConfigErrorDetails{
			{
				Type:   "Error",
				Name:   message,
				Reason: err.Error(),
			},
		},
	}
}

func init() {
}

func handleRequest(ctx context.Context, event SopsExtensionRequest) (SopsExtensionResponse, error) {
	log.Print("Attempting to decrypt configuration with SOPS ...")
	content, err := base64.StdEncoding.DecodeString(event.Content)

	if err != nil {
		log.Printf("Error decoding base64 content: %v", err)
		return errorResponse(err, "Error decoding base64 content"), nil
	}

	decrypted, err := decrypt.DataWithFormat(content, formats.FormatFromString(event.Parameters.Format))

	if err != nil {
		log.Printf("Error decrypting configuration: %v", err)
		return errorResponse(err, "Error decrypting configuration"), nil
	}

	log.Print("Successfully decrypted configuration")
	return SopsExtensionResponse{
		Content: base64.StdEncoding.EncodeToString(decrypted),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}