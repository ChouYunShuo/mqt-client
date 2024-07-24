/*
Core API

Testing RewardRedemptionsBetaAPIService

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

func Test_openapi_RewardRedemptionsBetaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RewardRedemptionsBetaAPIService GetRedemption", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string
		var redemptionToken string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.GetRedemption(context.Background(), token, redemptionToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService PostRedemptionTransition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string
		var redemptionToken string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.PostRedemptionTransition(context.Background(), token, redemptionToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService PostRewardProgramRedemption", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.PostRewardProgramRedemption(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService RetrieveRedemptions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptions(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RewardRedemptionsBetaAPIService RetrieveRedemptionsBalance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.RewardRedemptionsBetaAPI.RetrieveRedemptionsBalance(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
