# PaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**PaymentResponseData**](PaymentResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaymentResponse

`func NewPaymentResponse() *PaymentResponse`

NewPaymentResponse instantiates a new PaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseWithDefaults

`func NewPaymentResponseWithDefaults() *PaymentResponse`

NewPaymentResponseWithDefaults instantiates a new PaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *PaymentResponse) GetData() PaymentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentResponse) GetDataOk() (*PaymentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentResponse) SetData(v PaymentResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *PaymentResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PaymentResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PaymentResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PaymentResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaymentResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *PaymentResponse) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PaymentResponse) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PaymentResponse) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PaymentResponse) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLog

`func (o *PaymentResponse) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *PaymentResponse) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *PaymentResponse) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *PaymentResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


