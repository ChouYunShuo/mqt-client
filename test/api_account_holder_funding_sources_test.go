/*
Core API

Testing AccountHolderFundingSourcesAPIService

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

func Test_openapi_AccountHolderFundingSourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountHolderFundingSourcesAPIService GetFundingsourcesAchFundingsourcetoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetoken(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService GetFundingsourcesAchFundingsourcetokenVerificationamounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesAchFundingsourcetokenVerificationamounts(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService GetFundingsourcesBusinessBusinesstoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var businessToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesBusinessBusinesstoken(context.Background(), businessToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService GetFundingsourcesPaymentcardFundingsourcetoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesPaymentcardFundingsourcetoken(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService GetFundingsourcesUserUsertoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.GetFundingsourcesUserUsertoken(context.Background(), userToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PostFundingsourcesAch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesAch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PostFundingsourcesAchPartner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesAchPartner(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PostFundingsourcesPaymentcard", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PostFundingsourcesPaymentcard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PutFundingsourcesAchFundingsourcetoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesAchFundingsourcetoken(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PutFundingsourcesFundingsourcetoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetoken(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountHolderFundingSourcesAPIService PutFundingsourcesFundingsourcetokenDefault", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fundingSourceToken string

		resp, httpRes, err := apiClient.AccountHolderFundingSourcesAPI.PutFundingsourcesFundingsourcetokenDefault(context.Background(), fundingSourceToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
