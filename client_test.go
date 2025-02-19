package paypal

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TEST_ClientID = "AePKFcOP3dwCy1HRe6-GRs9goCCERgIet7w7sQeHD172IF12Iay9lET7JjgSwF0VrP-EzArGwbQXDJEe"
	TEST_Secret   = "EM4MkXpLY-GFlsEHSYzXWY9mV0Kmv-N3dPXAX_Cid4lDeyG1Gw88923wXVWKpKfaDuSaAs1RUctrsMdn"
)

var (
	client *Client
)

func TestMain(m *testing.M) {
	client = New(TEST_ClientID, TEST_Secret)
	os.Exit(m.Run())
}

func randString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func TestGetAccessToken(t *testing.T) {

	token := client.GetAccessToken(context.Background())

	require.NotNil(t, token)

	token2 := client.GetAccessToken(context.Background())
	require.Equal(t, token, token2)

}

func TestOrder(t *testing.T) {
	t.Run("create_minimal_order", func(t *testing.T) {
		o, err := client.CreateOrder(context.TODO(), CreatePurchaseUnit("create_minimal_order", "USD", 100))

		require.NoError(t, err)
		require.NotNil(t, o)
		require.NotEmpty(t, o.ID)
		require.Equal(t, "CREATED", o.Status)

	})
}
