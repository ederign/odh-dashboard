package mlflow

import (
	"context"

	mlflow "github.com/opendatahub-io/mlflow-go/mlflow"
	"github.com/opendatahub-io/mlflow-go/mlflow/promptregistry"
)

// MLFlowClientInterface defines the interface for MLFlow Prompt Registry operations
type MLFlowClientInterface interface {
	ListPrompts(ctx context.Context) (*promptregistry.PromptList, error)
	LoadPrompt(ctx context.Context, name string) (*promptregistry.Prompt, error)
	ListPromptVersions(ctx context.Context, name string) (*promptregistry.PromptVersionList, error)
	RegisterPrompt(ctx context.Context, name, template string, opts ...promptregistry.RegisterOption) (*promptregistry.Prompt, error)
}

// SDKClient wraps the mlflow-go SDK client
type SDKClient struct {
	client *mlflow.Client
}

// NewSDKClient creates a new MLFlow client using the mlflow-go SDK
func NewSDKClient(baseURL string) (*SDKClient, error) {
	client, err := mlflow.NewClient(
		mlflow.WithTrackingURI(baseURL),
		mlflow.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return &SDKClient{client: client}, nil
}

// ListPrompts retrieves all prompts from the MLFlow Prompt Registry
func (c *SDKClient) ListPrompts(ctx context.Context) (*promptregistry.PromptList, error) {
	return c.client.PromptRegistry().ListPrompts(ctx)
}

// LoadPrompt retrieves a specific prompt by name
func (c *SDKClient) LoadPrompt(ctx context.Context, name string) (*promptregistry.Prompt, error) {
	return c.client.PromptRegistry().LoadPrompt(ctx, name)
}

// ListPromptVersions retrieves all versions of a prompt
func (c *SDKClient) ListPromptVersions(ctx context.Context, name string) (*promptregistry.PromptVersionList, error) {
	return c.client.PromptRegistry().ListPromptVersions(ctx, name)
}

// RegisterPrompt creates a new prompt or adds a version to an existing prompt
func (c *SDKClient) RegisterPrompt(ctx context.Context, name, template string, opts ...promptregistry.RegisterOption) (*promptregistry.Prompt, error) {
	return c.client.PromptRegistry().RegisterPrompt(ctx, name, template, opts...)
}
