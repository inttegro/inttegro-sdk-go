package operatingpolicy

import (
	"context"

	"github.com/inttegro/inttegro-sdk-go/v12/internal/transport"
)

// Service manages reusable operating-policy templates.
type Service struct {
	client transport.Client
}

// Create creates version one of an operating policy.
func (s *Service) Create(ctx context.Context, params CreateParams) (*OperatingPolicy, error) {
	return s.send(ctx, "/operating_policies/create", params)
}

// Lookup returns the active version of an operating policy.
func (s *Service) Lookup(ctx context.Context, params LookupParams) (*OperatingPolicy, error) {
	return s.send(ctx, "/operating_policies/lookup", params)
}

// Update creates and activates a new immutable operating-policy version.
func (s *Service) Update(ctx context.Context, params UpdateParams) (*OperatingPolicy, error) {
	return s.send(ctx, "/operating_policies/update", params)
}

func (s *Service) send(ctx context.Context, path string, params any) (*OperatingPolicy, error) {
	var resp struct {
		OperatingPolicy OperatingPolicy `json:"operating_policy"`
	}
	if err := s.client.Do(ctx, "POST", path, params, &resp); err != nil {
		return nil, err
	}
	return &resp.OperatingPolicy, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
