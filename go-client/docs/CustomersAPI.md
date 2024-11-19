# \CustomersAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersAPI.md#CreateCustomer) | **Post** /customers | 
[**CustomerAffiliateBalance**](CustomersAPI.md#CustomerAffiliateBalance) | **Post** /customers/{id_or_email}/affiliate_balance | 
[**GetCustomer**](CustomersAPI.md#GetCustomer) | **Get** /customers/{id_or_email} | 
[**GetCustomerPaymentMethods**](CustomersAPI.md#GetCustomerPaymentMethods) | **Get** /customers/{id_or_email}/payment_methods | 
[**ListCustomers**](CustomersAPI.md#ListCustomers) | **Get** /customers | 
[**UpdateCustomer**](CustomersAPI.md#UpdateCustomer) | **Put** /customers/{id_or_email} | 



## CreateCustomer

> CreateCustomer200Response CreateCustomer(ctx).CreateCustomerRequest(createCustomerRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	createCustomerRequest := *openapiclient.NewCreateCustomerRequest("sample@gmail.com", "James", "Smith") // CreateCustomerRequest | customer JSON to be created
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).CreateCustomerRequest(createCustomerRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomer`: CreateCustomer200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerRequest** | [**CreateCustomerRequest**](CreateCustomerRequest.md) | customer JSON to be created | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateCustomer200Response**](CreateCustomer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAffiliateBalance

> CustomerAffiliateBalance200Response CustomerAffiliateBalance(ctx, idOrEmail).CustomerAffiliateBalanceRequest(customerAffiliateBalanceRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	idOrEmail := "sample6978a0e39d" // string | ID or email associated with the customer
	customerAffiliateBalanceRequest := *openapiclient.NewCustomerAffiliateBalanceRequest() // CustomerAffiliateBalanceRequest | JSON describing adjustments to affiliate balance
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomerAffiliateBalance(context.Background(), idOrEmail).CustomerAffiliateBalanceRequest(customerAffiliateBalanceRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomerAffiliateBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerAffiliateBalance`: CustomerAffiliateBalance200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomerAffiliateBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrEmail** | **string** | ID or email associated with the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAffiliateBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerAffiliateBalanceRequest** | [**CustomerAffiliateBalanceRequest**](CustomerAffiliateBalanceRequest.md) | JSON describing adjustments to affiliate balance | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CustomerAffiliateBalance200Response**](CustomerAffiliateBalance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> GetCustomer200Response GetCustomer(ctx, idOrEmail).XSellixMerchant(xSellixMerchant).Execute()





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
	idOrEmail := "sample6978a0e39d" // string | ID or email associated with the customer
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomer(context.Background(), idOrEmail).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomer`: GetCustomer200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrEmail** | **string** | ID or email associated with the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetCustomer200Response**](GetCustomer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerPaymentMethods

> GetCustomerPaymentMethods200Response GetCustomerPaymentMethods(ctx, idOrEmail).XSellixMerchant(xSellixMerchant).Execute()





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
	idOrEmail := "sample6978a0e39d" // string | ID or email associated with the customer
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.GetCustomerPaymentMethods(context.Background(), idOrEmail).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.GetCustomerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerPaymentMethods`: GetCustomerPaymentMethods200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.GetCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrEmail** | **string** | ID or email associated with the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetCustomerPaymentMethods200Response**](GetCustomerPaymentMethods200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> ListCustomers200Response ListCustomers(ctx).XSellixMerchant(xSellixMerchant).Page(page).Execute()





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
	resp, r, err := apiClient.CustomersAPI.ListCustomers(context.Background()).XSellixMerchant(xSellixMerchant).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomers`: ListCustomers200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **page** | **int32** | If pagination is desired, the page to fetch of the response | 

### Return type

[**ListCustomers200Response**](ListCustomers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> UpdateCustomer200Response UpdateCustomer(ctx, idOrEmail).UpdateCustomerRequest(updateCustomerRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	idOrEmail := "sample6978a0e39d" // string | ID or email associated with the customer
	updateCustomerRequest := *openapiclient.NewUpdateCustomerRequest() // UpdateCustomerRequest | customer JSON to be updated
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.UpdateCustomer(context.Background(), idOrEmail).UpdateCustomerRequest(updateCustomerRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.UpdateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomer`: UpdateCustomer200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrEmail** | **string** | ID or email associated with the customer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerRequest** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md) | customer JSON to be updated | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateCustomer200Response**](UpdateCustomer200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

