package repositories

import (
	"context"

	"github.com/opendatahub-io/gen-ai/internal/integrations/mlflow"
	"github.com/opendatahub-io/mlflow-go/mlflow/promptregistry"
)

// MLFlowPromptsRepository handles MLFlow Prompt Registry operations
type MLFlowPromptsRepository struct {
}

// NewMLFlowPromptsRepository creates a new MLFlow prompts repository
func NewMLFlowPromptsRepository() *MLFlowPromptsRepository {
	return &MLFlowPromptsRepository{}
}

// ListPrompts retrieves all prompts from the MLFlow Prompt Registry
func (r *MLFlowPromptsRepository) ListPrompts(
	client mlflow.MLFlowClientInterface,
	ctx context.Context,
) (*promptregistry.PromptList, error) {
	return client.ListPrompts(ctx)
}

// LoadPrompt retrieves a specific prompt by name
func (r *MLFlowPromptsRepository) LoadPrompt(
	client mlflow.MLFlowClientInterface,
	ctx context.Context,
	name string,
) (*promptregistry.Prompt, error) {
	return client.LoadPrompt(ctx, name)
}

// ListPromptVersions retrieves all versions of a prompt
func (r *MLFlowPromptsRepository) ListPromptVersions(
	client mlflow.MLFlowClientInterface,
	ctx context.Context,
	name string,
) (*promptregistry.PromptVersionList, error) {
	return client.ListPromptVersions(ctx, name)
}

// RegisterPrompt creates a new prompt or adds a version to an existing prompt
func (r *MLFlowPromptsRepository) RegisterPrompt(
	client mlflow.MLFlowClientInterface,
	ctx context.Context,
	name, template string,
	opts ...promptregistry.RegisterOption,
) (*promptregistry.Prompt, error) {
	return client.RegisterPrompt(ctx, name, template, opts...)
}
