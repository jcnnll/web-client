package httpx

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type httpClient struct {
	Headers http.Header
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	SetHeaders(headers http.Header)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body any) (*http.Response, error)
	Put(url string, headers http.Header, body any) (*http.Response, error)
	Patch(url string, headers http.Header, body any) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body any) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *httpClient) do(method, url string, headers http.Header, body any) (*http.Response, error) {
	client := http.Client{}

	reqHeaders := c.getRequestHeaders(headers)

	reqBody, err := c.getRequestBody(reqHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header = reqHeaders

	return client.Do(req)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	res := make(http.Header)

	// add instance headers:
	for header, val := range c.Headers {
		if len(val) > 0 {
			res.Set(header, val[0])
		}
	}

	// add headers for current req:
	for header, val := range requestHeaders {
		if len(val) > 0 {
			res.Set(header, val[0])
		}
	}

	return res
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}
