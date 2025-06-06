package httpx

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	clientHeaders := make(http.Header)
	clientHeaders.Set("Content-Type", "application/json")
	clientHeaders.Set("User-Agent", "my-app")
	client.Headers = clientHeaders

	reqHeaders := make(http.Header)
	reqHeaders.Set("X-Request-Id", "ABC-123")

	headers := client.getRequestHeaders(reqHeaders)

	if len(headers) != 3 {
		t.Error("expected 3 headers")
	}

	if headers.Get("Content-Type") != "application/json" {
		t.Error("invalid content type")
	}

	if headers.Get("User-Agent") != "my-app" {
		t.Error("invalid user agent")
	}

	if headers.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no error expected when passing in a nil body")
		}

		if body != nil {
			t.Error("expected body to be nil")
		}
	})

	t.Run("bodyWithJSON", func(t *testing.T) {
		reqBody := []string{"a", "b"}

		body, err := client.getRequestBody("application/json", reqBody)
		if err != nil {
			t.Error("no error wase expected when marshalling slice as JSON")
		}

		if string(body) != `["a","b"]` {
			t.Error("invalid JSON body returned")
		}
	})

	t.Run("bodyWithXML", func(t *testing.T) {
		reqBody := []string{"a", "b"}

		body, err := client.getRequestBody("application/xml", reqBody)
		if err != nil {
			t.Error("no error wase expected when marshalling slice as xml")
		}

		if string(body) != `<string>a</string><string>b</string>` {
			t.Error("invalid xml body returned")
		}
	})

	t.Run("bodyWithDefaultJSON", func(t *testing.T) {
		reqBody := []string{"a", "b"}

		body, err := client.getRequestBody("", reqBody)
		if err != nil {
			t.Error("no error wase expected when marshalling slice as JSON")
		}

		if string(body) != `["a","b"]` {
			t.Error("invalid JSON body returned")
		}
	})
}
