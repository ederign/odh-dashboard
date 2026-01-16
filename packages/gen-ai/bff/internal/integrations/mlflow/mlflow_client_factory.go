package mlflow

// MLFlowClientFactory interface for creating MLFlow clients
type MLFlowClientFactory interface {
	CreateClient(baseURL string) (MLFlowClientInterface, error)
}

// RealClientFactory creates real MLFlow clients using the SDK
type RealClientFactory struct{}

// NewRealClientFactory creates a factory for real MLFlow clients
func NewRealClientFactory() MLFlowClientFactory {
	return &RealClientFactory{}
}

// CreateClient creates a new real MLFlow client with the given base URL
func (f *RealClientFactory) CreateClient(baseURL string) (MLFlowClientInterface, error) {
	return NewSDKClient(baseURL)
}
