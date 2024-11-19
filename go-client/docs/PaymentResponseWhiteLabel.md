# PaymentResponseWhiteLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**PaymentResponseWhiteLabelData**](PaymentResponseWhiteLabelData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentResponseWhiteLabel

`func NewPaymentResponseWhiteLabel() *PaymentResponseWhiteLabel`

NewPaymentResponseWhiteLabel instantiates a new PaymentResponseWhiteLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseWhiteLabelWithDefaults

`func NewPaymentResponseWhiteLabelWithDefaults() *PaymentResponseWhiteLabel`

NewPaymentResponseWhiteLabelWithDefaults instantiates a new PaymentResponseWhiteLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentResponseWhiteLabel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponseWhiteLabel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponseWhiteLabel) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentResponseWhiteLabel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *PaymentResponseWhiteLabel) GetData() PaymentResponseWhiteLabelData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentResponseWhiteLabel) GetDataOk() (*PaymentResponseWhiteLabelData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentResponseWhiteLabel) SetData(v PaymentResponseWhiteLabelData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentResponseWhiteLabel) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentResponseWhiteLabel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentResponseWhiteLabel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentResponseWhiteLabel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentResponseWhiteLabel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLog

`func (o *PaymentResponseWhiteLabel) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *PaymentResponseWhiteLabel) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *PaymentResponseWhiteLabel) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *PaymentResponseWhiteLabel) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetError

`func (o *PaymentResponseWhiteLabel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PaymentResponseWhiteLabel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PaymentResponseWhiteLabel) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PaymentResponseWhiteLabel) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEnv

`func (o *PaymentResponseWhiteLabel) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PaymentResponseWhiteLabel) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PaymentResponseWhiteLabel) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PaymentResponseWhiteLabel) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


