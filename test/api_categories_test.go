/*
Sellix Developers API

Testing CategoriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sellix

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/crspy2/sellix-go"
)

func Test_sellix_CategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CategoriesAPIService CreateCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CategoriesAPI.CreateCategory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService DeleteCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uniqid string

		resp, httpRes, err := apiClient.CategoriesAPI.DeleteCategory(context.Background(), uniqid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService GetCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uniqid string

		resp, httpRes, err := apiClient.CategoriesAPI.GetCategory(context.Background(), uniqid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService ListCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CategoriesAPI.ListCategories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CategoriesAPIService UpdateCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uniqid string

		resp, httpRes, err := apiClient.CategoriesAPI.UpdateCategory(context.Background(), uniqid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}