package helper

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opendatahub-io/gen-ai/internal/constants"
	"github.com/opendatahub-io/gen-ai/internal/integrations/mlflow"
)

// GetContextMLFlowClientFromReq safely retrieves the MLFlow client from the HTTP request context.
// Returns an error if the client is not found or is nil.
func GetContextMLFlowClientFromReq(r *http.Request) (mlflow.MLFlowClientInterface, error) {
	return GetContextMLFlowClient(r.Context())
}

// GetContextMLFlowClient safely retrieves the MLFlow client from the given context.
// Returns an error if the client is not found or is nil.
func GetContextMLFlowClient(ctx context.Context) (mlflow.MLFlowClientInterface, error) {
	client, ok := ctx.Value(constants.MLFlowClientKey).(mlflow.MLFlowClientInterface)

	if !ok || client == nil {
		return nil, fmt.Errorf("missing MLFlow client in context - ensure AttachMLFlowClient middleware is used")
	}

	return client, nil
}
