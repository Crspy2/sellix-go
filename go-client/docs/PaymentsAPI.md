# \PaymentsAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompletePayment**](PaymentsAPI.md#CompletePayment) | **Put** /payments/{uniqid} | 
[**CreatePayment**](PaymentsAPI.md#CreatePayment) | **Post** /payments | 
[**VoidPayment**](PaymentsAPI.md#VoidPayment) | **Delete** /payments/{uniqid} | 



## CompletePayment

> CompletePayment200Response CompletePayment(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





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
	resp, r, err := apiClient.PaymentsAPI.CompletePayment(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CompletePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompletePayment`: CompletePayment200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CompletePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompletePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CompletePayment200Response**](CompletePayment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayment

> CreatePayment200Response CreatePayment(ctx).CreatePaymentRequest(createPaymentRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	createPaymentRequest := *openapiclient.NewCreatePaymentRequest("customer@gmail.com") // CreatePaymentRequest | payment JSON to be created
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.CreatePayment(context.Background()).CreatePaymentRequest(createPaymentRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePayment`: CreatePayment200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentRequest** | [**CreatePaymentRequest**](CreatePaymentRequest.md) | payment JSON to be created | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreatePayment200Response**](CreatePayment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidPayment

> VoidPayment200Response VoidPayment(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





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
	resp, r, err := apiClient.PaymentsAPI.VoidPayment(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.VoidPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidPayment`: VoidPayment200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.VoidPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**VoidPayment200Response**](VoidPayment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

