# \BlacklistsAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlacklist**](BlacklistsAPI.md#CreateBlacklist) | **Post** /blacklists | 
[**DeleteBlacklist**](BlacklistsAPI.md#DeleteBlacklist) | **Delete** /blacklists/{uniqid} | 
[**GetBlacklist**](BlacklistsAPI.md#GetBlacklist) | **Get** /blacklists/{uniqid} | 
[**ListBlacklists**](BlacklistsAPI.md#ListBlacklists) | **Get** /blacklists | 
[**UpdateBlacklist**](BlacklistsAPI.md#UpdateBlacklist) | **Put** /blacklists/{uniqid} | 



## CreateBlacklist

> CreateBlacklist200Response CreateBlacklist(ctx).CreateBlacklistRequest(createBlacklistRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	createBlacklistRequest := *openapiclient.NewCreateBlacklistRequest("EMAIL", "example@sellix.io") // CreateBlacklistRequest | blacklist JSON to be created
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistsAPI.CreateBlacklist(context.Background()).CreateBlacklistRequest(createBlacklistRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistsAPI.CreateBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlacklist`: CreateBlacklist200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistsAPI.CreateBlacklist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBlacklistRequest** | [**CreateBlacklistRequest**](CreateBlacklistRequest.md) | blacklist JSON to be created | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateBlacklist200Response**](CreateBlacklist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlacklist

> DeleteBlacklist200Response DeleteBlacklist(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





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
	resp, r, err := apiClient.BlacklistsAPI.DeleteBlacklist(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistsAPI.DeleteBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlacklist`: DeleteBlacklist200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistsAPI.DeleteBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**DeleteBlacklist200Response**](DeleteBlacklist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlacklist

> GetBlacklist200Response GetBlacklist(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





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
	resp, r, err := apiClient.BlacklistsAPI.GetBlacklist(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistsAPI.GetBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlacklist`: GetBlacklist200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistsAPI.GetBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetBlacklist200Response**](GetBlacklist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlacklists

> ListBlacklists200Response ListBlacklists(ctx).XSellixMerchant(xSellixMerchant).Page(page).Execute()





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
	resp, r, err := apiClient.BlacklistsAPI.ListBlacklists(context.Background()).XSellixMerchant(xSellixMerchant).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistsAPI.ListBlacklists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlacklists`: ListBlacklists200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistsAPI.ListBlacklists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlacklistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **page** | **int32** | If pagination is desired, the page to fetch of the response | 

### Return type

[**ListBlacklists200Response**](ListBlacklists200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlacklist

> UpdateBlacklist200Response UpdateBlacklist(ctx, uniqid).UpdateBlacklistRequest(updateBlacklistRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	updateBlacklistRequest := *openapiclient.NewUpdateBlacklistRequest() // UpdateBlacklistRequest | blacklist JSON to be updated
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistsAPI.UpdateBlacklist(context.Background(), uniqid).UpdateBlacklistRequest(updateBlacklistRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistsAPI.UpdateBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlacklist`: UpdateBlacklist200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistsAPI.UpdateBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBlacklistRequest** | [**UpdateBlacklistRequest**](UpdateBlacklistRequest.md) | blacklist JSON to be updated | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateBlacklist200Response**](UpdateBlacklist200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

