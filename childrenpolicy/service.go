package childrenpolicy

import (
	"context"

	"github.com/inttegro/inttegro-sdk-go/v13/internal/transport"
)

// Service provides gated access to the authenticated application's policy.
type Service struct {
	client transport.Client
}

// Lookup returns the authenticated application's latest children policy, or an
// exact immutable version when params specifies one.
func (s *Service) Lookup(ctx context.Context, options ...LookupParams) (*ChildrenPolicy, error) {
	params := LookupParams{}
	if len(options) > 0 {
		params = options[0]
	}
	var resp struct {
		ChildrenPolicy ChildrenPolicy `json:"children_policy"`
	}
	if err := s.client.Do(ctx, "POST", "/children_policies/lookup", params, &resp); err != nil {
		return nil, err
	}
	return &resp.ChildrenPolicy, nil
}

// Page lists immutable versions of the authenticated application's children policy.
func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/children_policies/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
