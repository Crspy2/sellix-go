# SubscriptionInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**SubscriptionInvoiceResponseData**](SubscriptionInvoiceResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriptionInvoiceResponse

`func NewSubscriptionInvoiceResponse() *SubscriptionInvoiceResponse`

NewSubscriptionInvoiceResponse instantiates a new SubscriptionInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionInvoiceResponseWithDefaults

`func NewSubscriptionInvoiceResponseWithDefaults() *SubscriptionInvoiceResponse`

NewSubscriptionInvoiceResponseWithDefaults instantiates a new SubscriptionInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubscriptionInvoiceResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionInvoiceResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionInvoiceResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionInvoiceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *SubscriptionInvoiceResponse) GetData() SubscriptionInvoiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionInvoiceResponse) GetDataOk() (*SubscriptionInvoiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionInvoiceResponse) SetData(v SubscriptionInvoiceResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SubscriptionInvoiceResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *SubscriptionInvoiceResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SubscriptionInvoiceResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SubscriptionInvoiceResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SubscriptionInvoiceResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *SubscriptionInvoiceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionInvoiceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionInvoiceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SubscriptionInvoiceResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *SubscriptionInvoiceResponse) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SubscriptionInvoiceResponse) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SubscriptionInvoiceResponse) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SubscriptionInvoiceResponse) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLog

`func (o *SubscriptionInvoiceResponse) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SubscriptionInvoiceResponse) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SubscriptionInvoiceResponse) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *SubscriptionInvoiceResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


