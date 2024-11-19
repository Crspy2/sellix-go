# \ProductsAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductsAPI.md#CreateProduct) | **Post** /products | 
[**CreateProductImage**](ProductsAPI.md#CreateProductImage) | **Post** /products?img&#x3D;true | 
[**DeleteProduct**](ProductsAPI.md#DeleteProduct) | **Delete** /products/{uniqid} | 
[**GetProduct**](ProductsAPI.md#GetProduct) | **Get** /products/{uniqid} | 
[**ListProducts**](ProductsAPI.md#ListProducts) | **Get** /products | 
[**UpdateProduct**](ProductsAPI.md#UpdateProduct) | **Put** /products/{uniqid} | 
[**UpdateProductImage**](ProductsAPI.md#UpdateProductImage) | **Put** /products/{uniqid}?img&#x3D;true | 



## CreateProduct

> CreateProduct200Response CreateProduct(ctx).CreateProductRequest(createProductRequest).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	createProductRequest := *openapiclient.NewCreateProductRequest("Software Activation Keys", float64(12.5), "Product description example.", "SERIALS") // CreateProductRequest | product JSON to be created
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.CreateProduct(context.Background()).CreateProductRequest(createProductRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.CreateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProduct`: CreateProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.CreateProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) | product JSON to be created | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**CreateProduct200Response**](CreateProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProductImage

> CreateProduct200Response CreateProductImage(ctx).Image(image).XSellixMerchant(xSellixMerchant).Data(data).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	image := os.NewFile(1234, "some_file") // *os.File | Image file to upload with the product.
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)
	data := *openapiclient.NewCreateProductImageRequestData("Software Activation Keys", float64(12.5), "Product description example.", "SERIALS") // CreateProductImageRequestData |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.CreateProductImage(context.Background()).Image(image).XSellixMerchant(xSellixMerchant).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.CreateProductImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProductImage`: CreateProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.CreateProductImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | ***os.File** | Image file to upload with the product. | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **data** | [**CreateProductImageRequestData**](CreateProductImageRequestData.md) |  | 

### Return type

[**CreateProduct200Response**](CreateProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProduct200Response DeleteProduct(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.DeleteProduct(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.DeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProduct`: DeleteProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.DeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**DeleteProduct200Response**](DeleteProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> GetProduct200Response GetProduct(ctx, uniqid).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProduct(context.Background(), uniqid).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProduct`: GetProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**GetProduct200Response**](GetProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ListProducts200Response ListProducts(ctx).XSellixMerchant(xSellixMerchant).Page(page).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)
	page := int32(1) // int32 | If pagination is desired, the page to fetch of the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ListProducts(context.Background()).XSellixMerchant(xSellixMerchant).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProducts`: ListProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **page** | **int32** | If pagination is desired, the page to fetch of the response | 

### Return type

[**ListProducts200Response**](ListProducts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> UpdateProduct200Response UpdateProduct(ctx, uniqid).UpdateProductRequest(updateProductRequest).XSellixMerchant(xSellixMerchant).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	updateProductRequest := *openapiclient.NewUpdateProductRequest() // UpdateProductRequest | product JSON to be updated
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.UpdateProduct(context.Background(), uniqid).UpdateProductRequest(updateProductRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.UpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProduct`: UpdateProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProductRequest** | [**UpdateProductRequest**](UpdateProductRequest.md) | product JSON to be updated | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateProduct200Response**](UpdateProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProductImage

> UpdateProduct200Response UpdateProductImage(ctx, uniqid).Image(image).XSellixMerchant(xSellixMerchant).Data(data).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/crspy2/sellix-go"
)

func main() {
	uniqid := "sample29a0e39d" // string | Uniqid of the resource
	image := os.NewFile(1234, "some_file") // *os.File | Image file to upload with the product.
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)
	data := *openapiclient.NewUpdateProductImageRequestData() // UpdateProductImageRequestData |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.UpdateProductImage(context.Background(), uniqid).Image(image).XSellixMerchant(xSellixMerchant).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.UpdateProductImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductImage`: UpdateProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.UpdateProductImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniqid** | **string** | Uniqid of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | Image file to upload with the product. | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 
 **data** | [**UpdateProductImageRequestData**](UpdateProductImageRequestData.md) |  | 

### Return type

[**UpdateProduct200Response**](UpdateProduct200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

