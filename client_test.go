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
	TEST_ClientID = "AfHi2_hDAfVtrxAeZSCf0ni6ksHQcPhrHQPOED7Jj6fKjdpa2_e-hIgN1j7x8ReMseFocow9BfsId3oV"
	TEST_Secret   = "EJFIzsOAQ-fkvj3_e4PA2_mpHiukNRMFgA66Q0mdHGe8HW8F7ngOxOQXO7LmeFbvo8qQCydrsfqcII2l"
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
		require.Equal(t, StatusCreated, o.Status)

		co, err := client.CaptureOrder(context.TODO(), o.ID)

		require.Error(t, err)
		require.Nil(t, co)

		pe := err.(*PaypalError)
		require.Equal(t, 422, pe.StatusCode)
		require.Equal(t, "UNPROCESSABLE_ENTITY", pe.Name)
		require.Equal(t, "ORDER_NOT_APPROVED", pe.Details[0].Issue)

	})
}
