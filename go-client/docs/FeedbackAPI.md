# \FeedbackAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedback**](FeedbackAPI.md#GetFeedback) | **Get** /feedback/{uniqid} | 
[**ListFeedback**](FeedbackAPI.md#ListFeedback) | **Get** /feedback | 
[**ReplyFeedback**](FeedbackAPI.md#ReplyFeedback) | **Post** /feedback/reply/{uniqid} | 



## GetFeedback

> GetFeedback200Response GetFeedback(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





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
	resp, r, err := apiClient.FeedbackAPI.GetFeedback(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.GetFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedback`: GetFeedback200Response
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.GetFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetFeedback200Response**](GetFeedback200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeedback

> ListFeedback200Response ListFeedback(ctx).XSellixMerchant(xSellixMerchant).Page(page).Execute()





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
	resp, r, err := apiClient.FeedbackAPI.ListFeedback(context.Background()).XSellixMerchant(xSellixMerchant).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.ListFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeedback`: ListFeedback200Response
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.ListFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **page** | **int32** | If pagination is desired, the page to fetch of the response | 

### Return type

[**ListFeedback200Response**](ListFeedback200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyFeedback

> ReplyFeedback200Response ReplyFeedback(ctx, uniqid).ReplyFeedbackRequest(replyFeedbackRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	replyFeedbackRequest := *openapiclient.NewReplyFeedbackRequest("This is a feedback reply.") // ReplyFeedbackRequest | JSON that contains the reply
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedbackAPI.ReplyFeedback(context.Background(), uniqid).ReplyFeedbackRequest(replyFeedbackRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedbackAPI.ReplyFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyFeedback`: ReplyFeedback200Response
	fmt.Fprintf(os.Stdout, "Response from `FeedbackAPI.ReplyFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replyFeedbackRequest** | [**ReplyFeedbackRequest**](ReplyFeedbackRequest.md) | JSON that contains the reply | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**ReplyFeedback200Response**](ReplyFeedback200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

