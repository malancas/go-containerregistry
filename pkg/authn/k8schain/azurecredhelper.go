package k8schain

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/docker/docker-credential-helpers/credentials"
)

type ACRHelper struct{}

func NewACRHelper() credentials.Helper {
	return &ACRHelper{}
}

func (a ACRHelper) Add(_ *credentials.Credentials) error {
	return fmt.Errorf("add is unimplemented")
}

func (a ACRHelper) Delete(_ string) error {
	return fmt.Errorf("delete is unimplemented")
}

func (a ACRHelper) Get(_ string) (string, string, error) {
	azCred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to obtain a credential: %w", err)
	}

	// We need to set the desired token policy to https://management.azure.com
	// to get a token that can be used to authenticate to the Azure Container Registry.
	opts := policy.TokenRequestOptions{
		Scopes: []string{"https://management.azure.com"},
	}
	token, err := azCred.GetToken(context.Background(), opts)
	if err != nil {
		return "", "", fmt.Errorf("failed to get token: %w", err)
	}

	return token.Token, "", nil
}

func (a ACRHelper) List() (map[string]string, error) {
	return nil, fmt.Errorf("list is unimplemented")
}

