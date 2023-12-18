package request

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/vitorfhc/htb-client/pkg/consts"
)

type Request struct {
	ctx  context.Context
	opts *RequestOptions
}

type RequestOptions struct {
	AuthToken string
	Path      consts.HTBPath
	Method    string
	Query     url.Values
	Body      io.Reader
	JSON      bool
}

type RequestOption func(*RequestOptions)

func New(ctx context.Context, opts ...RequestOption) *Request {
	options := &RequestOptions{
		Method: http.MethodGet,
		Query:  url.Values{},
		Body:   nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	return &Request{
		opts: options,
		ctx:  ctx,
	}
}

func WithAuthToken(token string) RequestOption {
	return func(opts *RequestOptions) {
		opts.AuthToken = token
	}
}

func WithPath(path consts.HTBPath) RequestOption {
	return func(opts *RequestOptions) {
		opts.Path = path
	}
}

func WithMethod(method string) RequestOption {
	return func(opts *RequestOptions) {
		opts.Method = method
	}
}

func WithQuery(query url.Values) RequestOption {
	return func(opts *RequestOptions) {
		opts.Query = query
	}
}

func WithBody(body io.Reader) RequestOption {
	return func(opts *RequestOptions) {
		opts.Body = body
	}
}

func WithJSONBody(body io.Reader) RequestOption {
	return func(opts *RequestOptions) {
		opts.Body = body
		opts.JSON = true
	}
}

func (r *Request) Build() (*http.Request, error) {
	parsedURL, err := url.Parse(consts.HTBHost + string(r.opts.Path))
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	if r.opts.Query != nil {
		parsedURL.RawQuery = r.opts.Query.Encode()
	}

	request, err := http.NewRequestWithContext(r.ctx, r.opts.Method, parsedURL.String(), r.opts.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if r.opts.AuthToken != "" {
		request.Header.Set("Authorization", "Bearer "+r.opts.AuthToken)
	}

	if r.opts.JSON {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}
