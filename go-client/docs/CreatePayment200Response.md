# CreatePayment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**PaymentResponseWhiteLabelData**](PaymentResponseWhiteLabelData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreatePayment200Response

`func NewCreatePayment200Response() *CreatePayment200Response`

NewCreatePayment200Response instantiates a new CreatePayment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayment200ResponseWithDefaults

`func NewCreatePayment200ResponseWithDefaults() *CreatePayment200Response`

NewCreatePayment200ResponseWithDefaults instantiates a new CreatePayment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreatePayment200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreatePayment200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreatePayment200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreatePayment200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *CreatePayment200Response) GetData() PaymentResponseWhiteLabelData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreatePayment200Response) GetDataOk() (*PaymentResponseWhiteLabelData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreatePayment200Response) SetData(v PaymentResponseWhiteLabelData)`

SetData sets Data field to given value.

### HasData

`func (o *CreatePayment200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *CreatePayment200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreatePayment200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreatePayment200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreatePayment200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *CreatePayment200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreatePayment200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreatePayment200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreatePayment200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *CreatePayment200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *CreatePayment200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *CreatePayment200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *CreatePayment200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLog

`func (o *CreatePayment200Response) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *CreatePayment200Response) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *CreatePayment200Response) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *CreatePayment200Response) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


