/*
Core API

Testing DirectDepositAccountsAPIService

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

func Test_openapi_DirectDepositAccountsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DirectDepositAccountsAPIService CreateAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.CreateAccount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService CreateTransition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.CreateTransition(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetCDDInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetCDDInfo(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetDirectDepositAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetDirectDepositAccount(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetDirectDepositAccountTransition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetDirectDepositAccountTransition(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetTransitionList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userToken string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetTransitionList(context.Background(), userToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetUserDirectDepositAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetUserDirectDepositAccounts(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService GetUserForDirectDepositAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountNumber string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.GetUserForDirectDepositAccount(context.Background(), accountNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService Update", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.Update(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DirectDepositAccountsAPIService UpdateCDDInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var token string
		var cddtoken string

		resp, httpRes, err := apiClient.DirectDepositAccountsAPI.UpdateCDDInfo(context.Background(), token, cddtoken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
