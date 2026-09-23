package order

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v10/search"
)

// Search finds order projections owned by the authenticated application.
func (s *Service) Search(ctx context.Context, params search.Params) (*search.Page, error) {
	var resp struct {
		Search search.Page `json:"search"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/search", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Search, nil
}
