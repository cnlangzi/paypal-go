// Package paypal provides a client for PayPal REST API

package paypal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"net/http"
	"strings"
)

const (
	SandboxHost = "https://api-m.sandbox.paypal.com"
)

type Client struct {
	httpClient  *http.Client
	credentials Credentials
	accessToken *AccessToken

	apiHost string
}

func New(clientID, secret string) *Client {
	return &Client{
		httpClient: &http.Client{},
		credentials: Credentials{
			ClientID: clientID,
			Secret:   secret,
		},

		apiHost: SandboxHost,
	}
}

func (c *Client) GetAccessToken(ctx context.Context) *AccessToken {
	if c.accessToken == nil || c.accessToken.WillExpires() {
		c.accessToken, _ = c.getAccessToken(ctx)
	}

	return c.accessToken
}

// curl -v -X POST "https://api-m.sandbox.paypal.com/v1/oauth2/token"\
// -u "CLIENT_ID:CLIENT_SECRET"\
// -H "Content-Type: application/x-www-form-urlencoded"\
// -d "grant_type=client_credentials"
func (c *Client) getAccessToken(ctx context.Context) (*AccessToken, error) {
	req, _ := http.NewRequest(http.MethodPost, c.apiHost+"/v1/oauth2/token", strings.NewReader("grant_type=client_credentials"))

	req.SetBasicAuth(c.credentials.ClientID, c.credentials.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req.WithContext(ctx))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("paypal: " + resp.Status)
	}

	var token AccessToken

	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, err
	}

	token.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn-30) * time.Second)

	return &token, nil
}

// CreateOrder creates a new order with the given purchase unit and options.
// It sends a request to the PayPal API to create the order and returns the created order details or an error.
func (c *Client) CreateOrder(ctx context.Context, unit PurchaseUnit, options ...CreateOrderOption) (*CreatedOrder, error) {

	cor := NewCreateOrderRequest(unit, options...)

	req, _ := c.createRequest(ctx, http.MethodPost, "/v2/checkout/orders/", cor)

	var o CreatedOrder

	e, err := c.doRequest(ctx, req, &o)

	if err != nil {
		return nil, err
	}

	if e != nil {
		return nil, e
	}

	return &o, nil
}

// CaptureOrder captures a previously created order by its ID.
// It sends a request to the PayPal API to capture the order and returns the captured order details or an error.
func (c *Client) CaptureOrder(ctx context.Context, orderID string) (*CapturedOrder, error) {
	req, _ := c.createRequest(ctx, http.MethodPost, "/v2/checkout/orders/"+orderID+"/capture", nil)

	var o CapturedOrder
	e, err := c.doRequest(ctx, req, &o)

	if err != nil {
		return nil, err
	}

	if e != nil {
		return nil, e
	}

	return &o, nil
}

func (c *Client) createRequest(ctx context.Context, method string, path string, body any) (*http.Request, error) {
	var buf []byte
	if body != nil {
		buf, _ = json.Marshal(body)
	}

	req, _ := http.NewRequest(method, c.apiHost+path, bytes.NewBuffer(buf))
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken(ctx).AccessToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, result any) (*PaypalError, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		var e PaypalError
		err = json.Unmarshal(buf, &e)
		if err != nil {
			return nil, err
		}
		e.StatusCode = resp.StatusCode
		return &e, nil
	}

	err = json.Unmarshal(buf, result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
