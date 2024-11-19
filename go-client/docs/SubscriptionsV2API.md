# \SubscriptionsV2API

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptionv2**](SubscriptionsV2API.md#CreateSubscriptionv2) | **Post** /subscriptions_v2 | 
[**DeleteSubscriptionv2**](SubscriptionsV2API.md#DeleteSubscriptionv2) | **Delete** /subscriptions_v2/{id} | 
[**UpdateSubscriptionv2**](SubscriptionsV2API.md#UpdateSubscriptionv2) | **Put** /subscriptions_v2/{id} | 



## CreateSubscriptionv2

> CreateSubscriptionv2200Response CreateSubscriptionv2(ctx).CreateSubscriptionv2Request(createSubscriptionv2Request).XSellixMerchant(xSellixMerchant).Execute()





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
	createSubscriptionv2Request := *openapiclient.NewCreateSubscriptionv2Request("customer.email@gmail.com", "67151e256f149", int32(1), "pp_1aed3fb1583653f925c260") // CreateSubscriptionv2Request | subscription creation JSON
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsV2API.CreateSubscriptionv2(context.Background()).CreateSubscriptionv2Request(createSubscriptionv2Request).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsV2API.CreateSubscriptionv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscriptionv2`: CreateSubscriptionv2200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsV2API.CreateSubscriptionv2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionv2Request** | [**CreateSubscriptionv2Request**](CreateSubscriptionv2Request.md) | subscription creation JSON | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateSubscriptionv2200Response**](CreateSubscriptionv2200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionv2

> CreateSubscriptionv2200Response DeleteSubscriptionv2(ctx, id).XSellixMerchant(xSellixMerchant).Execute()





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
	id := "sample6978a0e39d" // string | ID of the resource
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsV2API.DeleteSubscriptionv2(context.Background(), id).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsV2API.DeleteSubscriptionv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubscriptionv2`: CreateSubscriptionv2200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsV2API.DeleteSubscriptionv2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateSubscriptionv2200Response**](CreateSubscriptionv2200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscriptionv2

> CreateSubscriptionv2200Response UpdateSubscriptionv2(ctx, id).UpdateSubscriptionv2Request(updateSubscriptionv2Request).XSellixMerchant(xSellixMerchant).Execute()





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
	id := "sample6978a0e39d" // string | ID of the resource
	updateSubscriptionv2Request := *openapiclient.NewUpdateSubscriptionv2Request() // UpdateSubscriptionv2Request | subscription update JSON
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsV2API.UpdateSubscriptionv2(context.Background(), id).UpdateSubscriptionv2Request(updateSubscriptionv2Request).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsV2API.UpdateSubscriptionv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscriptionv2`: CreateSubscriptionv2200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsV2API.UpdateSubscriptionv2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubscriptionv2Request** | [**UpdateSubscriptionv2Request**](UpdateSubscriptionv2Request.md) | subscription update JSON | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateSubscriptionv2200Response**](CreateSubscriptionv2200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

