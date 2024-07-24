/*
Core API

Testing PaymentsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_PaymentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentsAPIService CreatePayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.PaymentsAPI.CreatePayment(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService ListPayments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.PaymentsAPI.ListPayments(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService ReleasePaymentHold", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var paymentToken string

		resp, httpRes, err := apiClient.PaymentsAPI.ReleasePaymentHold(context.Background(), accountToken, paymentToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService ResendWebhookEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventType string
		var resourceToken string

		resp, httpRes, err := apiClient.PaymentsAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService RetrievePayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var paymentToken string

		resp, httpRes, err := apiClient.PaymentsAPI.RetrievePayment(context.Background(), accountToken, paymentToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentsAPIService TransitionPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountToken string
		var paymentToken string

		resp, httpRes, err := apiClient.PaymentsAPI.TransitionPayment(context.Background(), accountToken, paymentToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
