package secretsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecretValue get the contents of a secretsmanager secret
func GetSecretValue(secretName string) (string, error) {
	// Load your AWS credentials and configuration. You can use environment variables, IAM roles, or AWS CLI configuration.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	// Create a Secrets Manager client
	client := secretsmanager.NewFromConfig(cfg)

	// Define the input for the GetSecretValue operation
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Call the GetSecretValue operation to retrieve the secret value
	result, err := client.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	// Check if the secret is a string
	if result.SecretString != nil {
		return *result.SecretString, nil
	}

	// Check if the secret is a binary value
	if result.SecretBinary != nil {
		return "", fmt.Errorf("Secret value is a binary. Not Supported")
	}

	return "", fmt.Errorf("Secret value is neither a string nor binary")
}
