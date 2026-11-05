package childrenpolicy

import (
	"context"

	"github.com/inttegro/inttegro-sdk-go/v12/internal/transport"
)

// Service provides gated access to the authenticated application's policy.
type Service struct {
	client transport.Client
}

// Lookup returns the authenticated application's active children policy.
func (s *Service) Lookup(ctx context.Context) (*ChildrenPolicy, error) {
	var resp struct {
		ChildrenPolicy ChildrenPolicy `json:"children_policy"`
	}
	if err := s.client.Do(ctx, "POST", "/children_policies/lookup", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.ChildrenPolicy, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
