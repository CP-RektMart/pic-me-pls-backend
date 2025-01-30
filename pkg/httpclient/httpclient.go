package httpclient

import (
	"net/url"

	"github.com/cockroachdb/errors"
	"github.com/valyala/fasthttp"
)

type Client struct {
	baseURL string
}

func New(baseURL string) (*Client, error) {
	if _, err := url.Parse(baseURL); err != nil {
		return nil, errors.Wrap(err, "invalid base URL")
	}

	return &Client{
		baseURL: baseURL,
	}, nil
}

type RequestOptions struct {
	path     string
	method   string
	Body     []byte
	Query    url.Values
	Header   map[string]string
	FormData url.Values
}

type HttpResponse struct {
	URL string
	fasthttp.Response
}

func (c *Client) request(reqOptions RequestOptions) (*HttpResponse, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(reqOptions.method)
	for k, v := range reqOptions.Header {
		req.Header.Set(k, v)
	}

	url := c.baseURL + reqOptions.path
	query := reqOptions.Query.Encode()
	if query != "" {
		url += "?" + query
	}
	req.SetRequestURI(url)

	if reqOptions.Body != nil {
		req.Header.SetContentType("application/json")
		req.SetBody(reqOptions.Body)
	} else if reqOptions.FormData != nil {
		req.Header.SetContentType("application/x-www-form-urlencoded")
		req.SetBodyString(reqOptions.FormData.Encode())
	}

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	httpResponse := HttpResponse{
		URL: url,
	}
	resp.CopyTo(&httpResponse.Response)

	return &httpResponse, nil
}

func (h *Client) Do(method, path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = method
	return h.request(reqOptions)
}

func (h *Client) Get(path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = fasthttp.MethodGet
	return h.request(reqOptions)
}

func (h *Client) Post(path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = fasthttp.MethodPost
	return h.request(reqOptions)
}

func (h *Client) Put(path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = fasthttp.MethodPut
	return h.request(reqOptions)
}

func (h *Client) Patch(path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = fasthttp.MethodPatch
	return h.request(reqOptions)
}

func (h *Client) Delete(path string, reqOptions RequestOptions) (*HttpResponse, error) {
	reqOptions.path = path
	reqOptions.method = fasthttp.MethodDelete
	return h.request(reqOptions)
}
