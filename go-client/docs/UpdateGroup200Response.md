# UpdateGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**CreateBlacklist200ResponseData**](CreateBlacklist200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateGroup200Response

`func NewUpdateGroup200Response() *UpdateGroup200Response`

NewUpdateGroup200Response instantiates a new UpdateGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroup200ResponseWithDefaults

`func NewUpdateGroup200ResponseWithDefaults() *UpdateGroup200Response`

NewUpdateGroup200ResponseWithDefaults instantiates a new UpdateGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateGroup200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateGroup200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateGroup200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateGroup200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UpdateGroup200Response) GetData() CreateBlacklist200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateGroup200Response) GetDataOk() (*CreateBlacklist200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateGroup200Response) SetData(v CreateBlacklist200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateGroup200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UpdateGroup200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UpdateGroup200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UpdateGroup200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UpdateGroup200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateGroup200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateGroup200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateGroup200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateGroup200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *UpdateGroup200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *UpdateGroup200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *UpdateGroup200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *UpdateGroup200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


