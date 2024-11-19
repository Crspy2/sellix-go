# ListFeedback200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**ListFeedback200ResponseData**](ListFeedback200ResponseData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewListFeedback200Response

`func NewListFeedback200Response() *ListFeedback200Response`

NewListFeedback200Response instantiates a new ListFeedback200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeedback200ResponseWithDefaults

`func NewListFeedback200ResponseWithDefaults() *ListFeedback200Response`

NewListFeedback200ResponseWithDefaults instantiates a new ListFeedback200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListFeedback200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFeedback200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFeedback200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListFeedback200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ListFeedback200Response) GetData() ListFeedback200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFeedback200Response) GetDataOk() (*ListFeedback200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFeedback200Response) SetData(v ListFeedback200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ListFeedback200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ListFeedback200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListFeedback200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListFeedback200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListFeedback200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLog

`func (o *ListFeedback200Response) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ListFeedback200Response) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ListFeedback200Response) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *ListFeedback200Response) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetError

`func (o *ListFeedback200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListFeedback200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListFeedback200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListFeedback200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEnv

`func (o *ListFeedback200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListFeedback200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListFeedback200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListFeedback200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


