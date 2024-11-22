/*
Sellix Developers API

Testing InformationAPIService

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

func Test_sellix_InformationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InformationAPIService GetSelf", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InformationAPI.GetSelf(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
