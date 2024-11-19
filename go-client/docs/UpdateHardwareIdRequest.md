# UpdateHardwareIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Uniqid of the license product. | 
**Key** | **string** | License key purchase by the customer. | 
**HardwareId** | **string** | Hardware ID to be assigned to the license. | 

## Methods

### NewUpdateHardwareIdRequest

`func NewUpdateHardwareIdRequest(productId string, key string, hardwareId string, ) *UpdateHardwareIdRequest`

NewUpdateHardwareIdRequest instantiates a new UpdateHardwareIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHardwareIdRequestWithDefaults

`func NewUpdateHardwareIdRequestWithDefaults() *UpdateHardwareIdRequest`

NewUpdateHardwareIdRequestWithDefaults instantiates a new UpdateHardwareIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *UpdateHardwareIdRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdateHardwareIdRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdateHardwareIdRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetKey

`func (o *UpdateHardwareIdRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateHardwareIdRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateHardwareIdRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetHardwareId

`func (o *UpdateHardwareIdRequest) GetHardwareId() string`

GetHardwareId returns the HardwareId field if non-nil, zero value otherwise.

### GetHardwareIdOk

`func (o *UpdateHardwareIdRequest) GetHardwareIdOk() (*string, bool)`

GetHardwareIdOk returns a tuple with the HardwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareId

`func (o *UpdateHardwareIdRequest) SetHardwareId(v string)`

SetHardwareId sets HardwareId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


