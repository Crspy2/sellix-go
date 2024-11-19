# \WhitelistsAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWhitelist**](WhitelistsAPI.md#CreateWhitelist) | **Post** /whitelists | 
[**DeleteWhitelist**](WhitelistsAPI.md#DeleteWhitelist) | **Delete** /whitelists/{uniqid} | 
[**GetWhitelist**](WhitelistsAPI.md#GetWhitelist) | **Get** /whitelists/{uniqid} | 
[**ListWhitelists**](WhitelistsAPI.md#ListWhitelists) | **Get** /whitelists | 
[**UpdateWhitelist**](WhitelistsAPI.md#UpdateWhitelist) | **Put** /whitelists/{uniqid} | 



## CreateWhitelist

> CreateWhitelist200Response CreateWhitelist(ctx).CreateWhitelistRequest(createWhitelistRequest).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createWhitelistRequest := *openapiclient.NewCreateWhitelistRequest("EMAIL", "example@sellix.io") // CreateWhitelistRequest | whitelist JSON to be created
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhitelistsAPI.CreateWhitelist(context.Background()).CreateWhitelistRequest(createWhitelistRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsAPI.CreateWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWhitelist`: CreateWhitelist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhitelistsAPI.CreateWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWhitelistRequest** | [**CreateWhitelistRequest**](CreateWhitelistRequest.md) | whitelist JSON to be created | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateWhitelist200Response**](CreateWhitelist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWhitelist

> DeleteWhitelist200Response DeleteWhitelist(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhitelistsAPI.DeleteWhitelist(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsAPI.DeleteWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWhitelist`: DeleteWhitelist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhitelistsAPI.DeleteWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**DeleteWhitelist200Response**](DeleteWhitelist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhitelist

> GetWhitelist200Response GetWhitelist(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhitelistsAPI.GetWhitelist(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsAPI.GetWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhitelist`: GetWhitelist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhitelistsAPI.GetWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetWhitelist200Response**](GetWhitelist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWhitelists

> ListWhitelists200Response ListWhitelists(ctx).XSellixMerchant(xSellixMerchant).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)
	page := int32(1) // int32 | If pagination is desired, the page to fetch of the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhitelistsAPI.ListWhitelists(context.Background()).XSellixMerchant(xSellixMerchant).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsAPI.ListWhitelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWhitelists`: ListWhitelists200Response
	fmt.Fprintf(os.Stdout, "Response from `WhitelistsAPI.ListWhitelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWhitelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **page** | **int32** | If pagination is desired, the page to fetch of the response | 

### Return type

[**ListWhitelists200Response**](ListWhitelists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhitelist

> UpdateWhitelist200Response UpdateWhitelist(ctx, uniqid).UpdateWhitelistRequest(updateWhitelistRequest).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	updateWhitelistRequest := *openapiclient.NewUpdateWhitelistRequest() // UpdateWhitelistRequest | whitelist JSON to be updated
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WhitelistsAPI.UpdateWhitelist(context.Background(), uniqid).UpdateWhitelistRequest(updateWhitelistRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsAPI.UpdateWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhitelist`: UpdateWhitelist200Response
	fmt.Fprintf(os.Stdout, "Response from `WhitelistsAPI.UpdateWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWhitelistRequest** | [**UpdateWhitelistRequest**](UpdateWhitelistRequest.md) | whitelist JSON to be updated | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateWhitelist200Response**](UpdateWhitelist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

