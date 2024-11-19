# GetSubscription200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**GetSubscription200ResponseData**](GetSubscription200ResponseData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSubscription200Response

`func NewGetSubscription200Response() *GetSubscription200Response`

NewGetSubscription200Response instantiates a new GetSubscription200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscription200ResponseWithDefaults

`func NewGetSubscription200ResponseWithDefaults() *GetSubscription200Response`

NewGetSubscription200ResponseWithDefaults instantiates a new GetSubscription200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSubscription200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSubscription200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSubscription200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSubscription200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetSubscription200Response) GetData() GetSubscription200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSubscription200Response) GetDataOk() (*GetSubscription200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSubscription200Response) SetData(v GetSubscription200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetSubscription200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *GetSubscription200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetSubscription200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetSubscription200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetSubscription200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLog

`func (o *GetSubscription200Response) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *GetSubscription200Response) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *GetSubscription200Response) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *GetSubscription200Response) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetError

`func (o *GetSubscription200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetSubscription200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetSubscription200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetSubscription200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEnv

`func (o *GetSubscription200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetSubscription200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetSubscription200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetSubscription200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


