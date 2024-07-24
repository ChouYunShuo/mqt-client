/*
Core API

Testing LedgerEntriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/torpago/mqt-client"
)

func Test_openapi_LedgerEntriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LedgerEntriesAPIService GetAccountLedgerEntry", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var ledgerEntryToken string

		resp, httpRes, err := apiClient.LedgerEntriesAPI.GetAccountLedgerEntry(context.Background(), accountToken, ledgerEntryToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LedgerEntriesAPIService ListAccountLedgerEntries", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.LedgerEntriesAPI.ListAccountLedgerEntries(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LedgerEntriesAPIService ResendWebhookEvent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var eventType string
		var resourceToken string

		resp, httpRes, err := apiClient.LedgerEntriesAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
