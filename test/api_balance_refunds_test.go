/*
Core API

Testing BalanceRefundsAPIService

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

func Test_openapi_BalanceRefundsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BalanceRefundsAPIService CreateBalanceRefund", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.BalanceRefundsAPI.CreateBalanceRefund(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
