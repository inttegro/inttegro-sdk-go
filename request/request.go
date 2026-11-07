// Package request defines controls shared by resource service calls.
package request

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Meta carries per-request controls that do not change an operation payload.
type Meta struct {
	// IdempotencyKey safely identifies retries for mutation requests.
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// Response selects how the successful response should be represented.
	Response *Response `json:"response,omitempty"`
}

// Response contains caller-controlled preferences for the successful response.
type Response struct {
	// Projection selects public properties independently for each named response
	// resource. Resource names and property paths are documented by the endpoint.
	Projection Projection `json:"projection,omitempty"`
}

// Projection maps a response resource name to its requested property selection.
//
// An endpoint validates resource names and property paths against its public
// response contract. Resources not present in the map retain their complete
// caller-visible representation.
type Projection map[string]PropertySelection

// PropertySelection includes or excludes public properties from one response
// resource. Construct one with Pick or Omit so the wire value always has one
// selection mode.
type PropertySelection struct {
	mode  string
	paths []string
}

// Pick includes only the supplied public property paths and the ancestors
// needed to contain them. Selecting a parent includes its complete
// caller-visible subtree.
func Pick(propertyPaths ...string) PropertySelection {
	return PropertySelection{mode: "pick", paths: append([]string(nil), propertyPaths...)}
}

// Omit excludes the supplied public property paths and all of their
// descendants.
func Omit(propertyPaths ...string) PropertySelection {
	return PropertySelection{mode: "omit", paths: append([]string(nil), propertyPaths...)}
}

// MarshalJSON implements json.Marshaler. Empty or blank property paths are
// rejected before an HTTP request can be encoded.
func (selection PropertySelection) MarshalJSON() ([]byte, error) {
	if selection.mode != "pick" && selection.mode != "omit" {
		return nil, fmt.Errorf("request: property selection must use Pick or Omit")
	}
	if len(selection.paths) == 0 {
		return nil, fmt.Errorf("request: %s must contain at least one property path", selection.mode)
	}
	for _, path := range selection.paths {
		if strings.TrimSpace(path) == "" {
			return nil, fmt.Errorf("request: %s property paths must not be blank", selection.mode)
		}
	}
	return json.Marshal(map[string][]string{selection.mode: append([]string(nil), selection.paths...)})
}

// Options contains controls passed alongside a request body.
type Options struct {
	IdempotencyKey string
}

// Option configures one resource request.
type Option func(*Options)

// WithIdempotencyKey sets the key used to safely retry a mutation.
func WithIdempotencyKey(key string) Option {
	return func(options *Options) {
		options.IdempotencyKey = key
	}
}

// Apply resolves a sequence of request options.
func Apply(options []Option) Options {
	var resolved Options
	for _, option := range options {
		if option != nil {
			option(&resolved)
		}
	}
	return resolved
}
