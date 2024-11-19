# GetCustomerPaymentMethods200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**GetCustomerPaymentMethods200ResponseData**](GetCustomerPaymentMethods200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCustomerPaymentMethods200Response

`func NewGetCustomerPaymentMethods200Response() *GetCustomerPaymentMethods200Response`

NewGetCustomerPaymentMethods200Response instantiates a new GetCustomerPaymentMethods200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerPaymentMethods200ResponseWithDefaults

`func NewGetCustomerPaymentMethods200ResponseWithDefaults() *GetCustomerPaymentMethods200Response`

NewGetCustomerPaymentMethods200ResponseWithDefaults instantiates a new GetCustomerPaymentMethods200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetCustomerPaymentMethods200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCustomerPaymentMethods200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCustomerPaymentMethods200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCustomerPaymentMethods200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetCustomerPaymentMethods200Response) GetData() GetCustomerPaymentMethods200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerPaymentMethods200Response) GetDataOk() (*GetCustomerPaymentMethods200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerPaymentMethods200Response) SetData(v GetCustomerPaymentMethods200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetCustomerPaymentMethods200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *GetCustomerPaymentMethods200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetCustomerPaymentMethods200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetCustomerPaymentMethods200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetCustomerPaymentMethods200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *GetCustomerPaymentMethods200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetCustomerPaymentMethods200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetCustomerPaymentMethods200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetCustomerPaymentMethods200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *GetCustomerPaymentMethods200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetCustomerPaymentMethods200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetCustomerPaymentMethods200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetCustomerPaymentMethods200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


