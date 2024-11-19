# \LicensesAPI

All URIs are relative to *https://dev.sellix.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckLicense**](LicensesAPI.md#CheckLicense) | **Post** /products/licensing/check | 
[**UpdateHardwareId**](LicensesAPI.md#UpdateHardwareId) | **Put** /products/licensing/hardware_id | 



## CheckLicense

> UpdateHardwareId200Response CheckLicense(ctx).CheckLicenseRequest(checkLicenseRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	checkLicenseRequest := *openapiclient.NewCheckLicenseRequest("5eb6b2bde8a50", "LICENSE-SAMPLE-00") // CheckLicenseRequest | product licensing assign hardware ID JSON
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.CheckLicense(context.Background()).CheckLicenseRequest(checkLicenseRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.CheckLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckLicense`: UpdateHardwareId200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.CheckLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkLicenseRequest** | [**CheckLicenseRequest**](CheckLicenseRequest.md) | product licensing assign hardware ID JSON | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateHardwareId200Response**](UpdateHardwareId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHardwareId

> UpdateHardwareId200Response UpdateHardwareId(ctx).UpdateHardwareIdRequest(updateHardwareIdRequest).XSellixMerchant(xSellixMerchant).Execute()





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
	updateHardwareIdRequest := *openapiclient.NewUpdateHardwareIdRequest("5eb6b2bde8a50", "LICENSE-SAMPLE-00", "098H52ST479QE053V2") // UpdateHardwareIdRequest | product licensing assign hardware ID JSON
	xSellixMerchant := "Sellix" // string | The name of the store to perform actions on via the API (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.UpdateHardwareId(context.Background()).UpdateHardwareIdRequest(updateHardwareIdRequest).XSellixMerchant(xSellixMerchant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.UpdateHardwareId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHardwareId`: UpdateHardwareId200Response
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.UpdateHardwareId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHardwareIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateHardwareIdRequest** | [**UpdateHardwareIdRequest**](UpdateHardwareIdRequest.md) | product licensing assign hardware ID JSON | 
 **xSellixMerchant** | **string** | The name of the store to perform actions on via the API | 

### Return type

[**UpdateHardwareId200Response**](UpdateHardwareId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

