package operatingpolicy

import (
	"context"

	"github.com/inttegro/inttegro-sdk-go/v13/internal/transport"
)

// Service manages reusable operating-policy templates.
type Service struct {
	client transport.Client
}

// Create creates version one of an operating policy.
func (s *Service) Create(ctx context.Context, params CreateParams) (*OperatingPolicy, error) {
	return s.send(ctx, "/operating_policies/create", params)
}

// Lookup returns the latest version of an operating policy, or an exact
// immutable version when params specifies one.
func (s *Service) Lookup(ctx context.Context, params LookupParams) (*OperatingPolicy, error) {
	return s.send(ctx, "/operating_policies/lookup", params)
}

// Page lists immutable versions of one operating policy.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/operating_policies/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
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
