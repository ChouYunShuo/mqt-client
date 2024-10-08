/*
Core API

Testing AccountCardsAPIService

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

func Test_openapi_AccountCardsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountCardsAPIService CreateCardForAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.AccountCardsAPI.CreateCardForAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountCardsAPIService GetCardByAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var cardToken string

		resp, httpRes, err := apiClient.AccountCardsAPI.GetCardByAccount(context.Background(), accountToken, cardToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountCardsAPIService GetCardsByAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.AccountCardsAPI.GetCardsByAccount(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
