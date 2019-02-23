package astigandi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/asticode/go-astitools/http"
	"github.com/pkg/errors"
)

const (
	baseURL = "https://dns.api.gandi.net/api/v5"
)

// Client represents the client
type Client struct {
	apiKey string
	s      *astihttp.Sender
}

// New creates a new client
func New(c Configuration) *Client {
	return &Client{
		apiKey: c.APIKey,
		s:      astihttp.NewSender(c.Sender),
	}
}

func (c *Client) send(method, url string, reqPayload, respPayload interface{}) (err error) {
	// Create body
	var body io.Reader
	if reqPayload != nil {
		// Marshal
		buf := &bytes.Buffer{}
		if err = json.NewEncoder(buf).Encode(reqPayload); err != nil {
			err = errors.Wrapf(err, "astigandi: marshaling payload of %s request to %s failed", method, url)
			return
		}

		// Set body
		body = buf
	}

	// Create request
	var req *http.Request
	if req, err = http.NewRequest(method, baseURL+url, body); err != nil {
		err = errors.Wrapf(err, "astigandi: creating %s request to %s failed", method, url)
		return
	}

	// Add headers
	if reqPayload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("X-Api-Key", c.apiKey)

	// Send
	var resp *http.Response
	if resp, err = c.s.Send(req); err != nil {
		err = errors.Wrapf(err, "astigandi: sending %s request to %s failed", req.Method, req.URL.Path)
		return
	}
	defer resp.Body.Close()

	// Process error
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		// Unmarshal
		var e Error
		if err = json.NewDecoder(resp.Body).Decode(&e); err != nil {
			err = errors.Wrap(err, "astigandi: unmarshaling error failed")
			return
		}

		// Set error
		err = e
		return
	}

	// Unmarshal
	if respPayload != nil {
		if err = json.NewDecoder(resp.Body).Decode(respPayload); err != nil {
			err = errors.Wrap(err, "astigandi: unmarshaling failed")
			return
		}
	}
	return
}
