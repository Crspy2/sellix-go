# SubscriptionTrialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**SubscriptionTrialResponseData**](SubscriptionTrialResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriptionTrialResponse

`func NewSubscriptionTrialResponse() *SubscriptionTrialResponse`

NewSubscriptionTrialResponse instantiates a new SubscriptionTrialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTrialResponseWithDefaults

`func NewSubscriptionTrialResponseWithDefaults() *SubscriptionTrialResponse`

NewSubscriptionTrialResponseWithDefaults instantiates a new SubscriptionTrialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubscriptionTrialResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionTrialResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionTrialResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionTrialResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *SubscriptionTrialResponse) GetData() SubscriptionTrialResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionTrialResponse) GetDataOk() (*SubscriptionTrialResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionTrialResponse) SetData(v SubscriptionTrialResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *SubscriptionTrialResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *SubscriptionTrialResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SubscriptionTrialResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SubscriptionTrialResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SubscriptionTrialResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *SubscriptionTrialResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionTrialResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionTrialResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SubscriptionTrialResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *SubscriptionTrialResponse) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *SubscriptionTrialResponse) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *SubscriptionTrialResponse) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *SubscriptionTrialResponse) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetLog

`func (o *SubscriptionTrialResponse) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SubscriptionTrialResponse) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SubscriptionTrialResponse) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *SubscriptionTrialResponse) HasLog() bool`

HasLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


