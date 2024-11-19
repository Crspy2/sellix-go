# CheckLicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Uniqid of the license product. | 
**Key** | **string** | License key purchase by the customer. | 
**HardwareId** | Pointer to **string** | Hardware ID to be assigned to the license. If passed, the hardware ID will be checked against the license key. | [optional] 

## Methods

### NewCheckLicenseRequest

`func NewCheckLicenseRequest(productId string, key string, ) *CheckLicenseRequest`

NewCheckLicenseRequest instantiates a new CheckLicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckLicenseRequestWithDefaults

`func NewCheckLicenseRequestWithDefaults() *CheckLicenseRequest`

NewCheckLicenseRequestWithDefaults instantiates a new CheckLicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CheckLicenseRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CheckLicenseRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CheckLicenseRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetKey

`func (o *CheckLicenseRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CheckLicenseRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CheckLicenseRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetHardwareId

`func (o *CheckLicenseRequest) GetHardwareId() string`

GetHardwareId returns the HardwareId field if non-nil, zero value otherwise.

### GetHardwareIdOk

`func (o *CheckLicenseRequest) GetHardwareIdOk() (*string, bool)`

GetHardwareIdOk returns a tuple with the HardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareId

`func (o *CheckLicenseRequest) SetHardwareId(v string)`

SetHardwareId sets HardwareId field to given value.

### HasHardwareId

`func (o *CheckLicenseRequest) HasHardwareId() bool`

HasHardwareId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


