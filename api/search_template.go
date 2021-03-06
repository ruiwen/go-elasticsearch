// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// SearchTemplateOption is a non-required SearchTemplate option that gets applied to an HTTP request.
type SearchTemplateOption func(r *transport.Request)

// WithSearchTemplateIndex - a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
func WithSearchTemplateIndex(index []string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateType - a comma-separated list of document types to search; leave empty to perform the operation on all types.
func WithSearchTemplateType(documentType []string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithSearchTemplateAllowNoIndices(allowNoIndices bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// SearchTemplateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type SearchTemplateExpandWildcards int

const (
	// SearchTemplateExpandWildcardsOpen can be used to set SearchTemplateExpandWildcards to "open"
	SearchTemplateExpandWildcardsOpen = iota
	// SearchTemplateExpandWildcardsClosed can be used to set SearchTemplateExpandWildcards to "closed"
	SearchTemplateExpandWildcardsClosed = iota
	// SearchTemplateExpandWildcardsNone can be used to set SearchTemplateExpandWildcards to "none"
	SearchTemplateExpandWildcardsNone = iota
	// SearchTemplateExpandWildcardsAll can be used to set SearchTemplateExpandWildcards to "all"
	SearchTemplateExpandWildcardsAll = iota
)

// WithSearchTemplateExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithSearchTemplateExpandWildcards(expandWildcards SearchTemplateExpandWildcards) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateExplain - specify whether to return detailed information about score computation as part of a hit.
func WithSearchTemplateExplain(explain bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithSearchTemplateIgnoreUnavailable(ignoreUnavailable bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplatePreference - specify the node or shard the operation should be performed on (default: random).
func WithSearchTemplatePreference(preference string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateProfile - specify whether to profile the query execution.
func WithSearchTemplateProfile(profile bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateRouting - a comma-separated list of specific routing values.
func WithSearchTemplateRouting(routing []string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateScroll - specify how long a consistent view of the index should be maintained for scrolled search.
func WithSearchTemplateScroll(scroll time.Duration) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// SearchTemplateSearchType - search operation type.
type SearchTemplateSearchType int

const (
	// SearchTemplateSearchTypeQueryThenFetch can be used to set SearchTemplateSearchType to "query_then_fetch"
	SearchTemplateSearchTypeQueryThenFetch = iota
	// SearchTemplateSearchTypeQueryAndFetch can be used to set SearchTemplateSearchType to "query_and_fetch"
	SearchTemplateSearchTypeQueryAndFetch = iota
	// SearchTemplateSearchTypeDfsQueryThenFetch can be used to set SearchTemplateSearchType to "dfs_query_then_fetch"
	SearchTemplateSearchTypeDfsQueryThenFetch = iota
	// SearchTemplateSearchTypeDfsQueryAndFetch can be used to set SearchTemplateSearchType to "dfs_query_and_fetch"
	SearchTemplateSearchTypeDfsQueryAndFetch = iota
)

// WithSearchTemplateSearchType - search operation type.
func WithSearchTemplateSearchType(searchType SearchTemplateSearchType) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
func WithSearchTemplateTypedKeys(typedKeys bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateBody - the search definition template and its params.
func WithSearchTemplateBody(body map[string]interface{}) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateErrorTrace - include the stack trace of returned errors.
func WithSearchTemplateErrorTrace(errorTrace bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithSearchTemplateFilterPath(filterPath []string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateHuman - return human readable values for statistics.
func WithSearchTemplateHuman(human bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateIgnore - ignores the specified HTTP status codes.
func WithSearchTemplateIgnore(ignore []int) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplatePretty - pretty format the returned JSON response.
func WithSearchTemplatePretty(pretty bool) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTemplateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSearchTemplateSourceParam(sourceParam string) SearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// SearchTemplate - see https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html for more info.
//
// options: optional parameters.
func (a *API) SearchTemplate(options ...SearchTemplateOption) (*SearchTemplateResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &SearchTemplateResponse{resp}, err
}

// SearchTemplateResponse is the response for SearchTemplate.
type SearchTemplateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *SearchTemplateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
