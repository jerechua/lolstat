package myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestBuilder struct {
	host        string
	path        string
	scheme      string
	queryParams map[string]string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		scheme:      "http",
		queryParams: map[string]string{},
	}
}

func (rb *RequestBuilder) SetHost(host string) *RequestBuilder {
	rb.host = host
	return rb
}

func (rb *RequestBuilder) SetPath(path string) *RequestBuilder {
	rb.path = path
	return rb
}

func (rb *RequestBuilder) Secure() *RequestBuilder {
	rb.scheme = "https"
	return rb
}

func (rb *RequestBuilder) AddQueryParam(k, v string) *RequestBuilder {
	rb.queryParams[k] = v
	return rb
}

func (rb *RequestBuilder) Build() (*Request, error) {
	if rb.host == "" {
		return nil, fmt.Errorf("RequestBuilder host must be set.")
	}
	if rb.path == "" {
		return nil, fmt.Errorf("RequestBuilder path must be set.")
	}
	return &Request{
		host:        rb.host,
		path:        rb.path,
		scheme:      rb.scheme,
		queryParams: rb.queryParams,
	}, nil
}

type Request struct {
	host        string
	path        string
	scheme      string
	queryParams map[string]string
}

func NewRequest(host, path string) *Request {
	return &Request{
		host:        host,
		path:        path,
		scheme:      "http",
		queryParams: map[string]string{},
	}
}

func (r *Request) AddQueryParam(k, v string) error {
	val, ok := r.queryParams[k]
	if ok {
		return fmt.Errorf(
			"Query parameter for key %q already exists with value %q", k, val)
	}
	r.queryParams[k] = v
	return nil
}

func (r *Request) Secure() {
	r.scheme = "https"
}

// GetAsync if Get should be done via goroutine.
func (r *Request) GetAsync(cr chan *Response, ce chan error, done func()) {
	defer done()
	res, err := r.Get()
	cr <- res
	ce <- err
}

func (r *Request) Get() (*Response, error) {
	u := &url.URL{
		Host:   r.host,
		Path:   r.path,
		Scheme: r.scheme,
	}

	q := u.Query()
	for k, v := range r.queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req := &http.Request{
		Method: "GET",
		URL:    u,
	}

	hc := http.Client{}
	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		res:  res,
		body: body,
	}, nil
}

type Response struct {
	res  *http.Response
	body []byte
}

func (r *Response) OK() bool {
	return r.res.StatusCode == http.StatusOK
}

func (r *Response) StatusError() error {
	if r.res.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"Expected status to be %s(%d), but got %s(%d) instead",
			http.StatusText(http.StatusOK),
			http.StatusOK,
			http.StatusText(r.res.StatusCode),
			r.res.StatusCode)
	}
	return nil
}

func (r *Response) Body() []byte {
	return r.body
}
