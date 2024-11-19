# ProductDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice this download is linked to. | [optional] 
**CustomerIp** | Pointer to **string** | The IP of the customer that downloaded the product. | [optional] 
**CustomerIsp** | Pointer to **string** | The ISP of the customer that downloaded the product. | [optional] 
**CustomerTimezone** | Pointer to **string** | The timezone of the customer that downloaded the product. | [optional] 
**CustomerCountry** | Pointer to **string** | The country of the customer that downloaded the product. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp, available if the download has been created. | [optional] 

## Methods

### NewProductDownload

`func NewProductDownload() *ProductDownload`

NewProductDownload instantiates a new ProductDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDownloadWithDefaults

`func NewProductDownloadWithDefaults() *ProductDownload`

NewProductDownloadWithDefaults instantiates a new ProductDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductDownload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductDownload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductDownload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductDownload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *ProductDownload) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ProductDownload) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ProductDownload) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ProductDownload) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetCustomerIp

`func (o *ProductDownload) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *ProductDownload) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *ProductDownload) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *ProductDownload) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetCustomerIsp

`func (o *ProductDownload) GetCustomerIsp() string`

GetCustomerIsp returns the CustomerIsp field if non-nil, zero value otherwise.

### GetCustomerIspOk

`func (o *ProductDownload) GetCustomerIspOk() (*string, bool)`

GetCustomerIspOk returns a tuple with the CustomerIsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIsp

`func (o *ProductDownload) SetCustomerIsp(v string)`

SetCustomerIsp sets CustomerIsp field to given value.

### HasCustomerIsp

`func (o *ProductDownload) HasCustomerIsp() bool`

HasCustomerIsp returns a boolean if a field has been set.

### GetCustomerTimezone

`func (o *ProductDownload) GetCustomerTimezone() string`

GetCustomerTimezone returns the CustomerTimezone field if non-nil, zero value otherwise.

### GetCustomerTimezoneOk

`func (o *ProductDownload) GetCustomerTimezoneOk() (*string, bool)`

GetCustomerTimezoneOk returns a tuple with the CustomerTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTimezone

`func (o *ProductDownload) SetCustomerTimezone(v string)`

SetCustomerTimezone sets CustomerTimezone field to given value.

### HasCustomerTimezone

`func (o *ProductDownload) HasCustomerTimezone() bool`

HasCustomerTimezone returns a boolean if a field has been set.

### GetCustomerCountry

`func (o *ProductDownload) GetCustomerCountry() string`

GetCustomerCountry returns the CustomerCountry field if non-nil, zero value otherwise.

### GetCustomerCountryOk

`func (o *ProductDownload) GetCustomerCountryOk() (*string, bool)`

GetCustomerCountryOk returns a tuple with the CustomerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountry

`func (o *ProductDownload) SetCustomerCountry(v string)`

SetCustomerCountry sets CustomerCountry field to given value.

### HasCustomerCountry

`func (o *ProductDownload) HasCustomerCountry() bool`

HasCustomerCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductDownload) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductDownload) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductDownload) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductDownload) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


