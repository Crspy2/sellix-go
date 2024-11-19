# ListGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**ListGroups200ResponseData**](ListGroups200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewListGroups200Response

`func NewListGroups200Response() *ListGroups200Response`

NewListGroups200Response instantiates a new ListGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroups200ResponseWithDefaults

`func NewListGroups200ResponseWithDefaults() *ListGroups200Response`

NewListGroups200ResponseWithDefaults instantiates a new ListGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListGroups200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListGroups200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListGroups200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListGroups200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ListGroups200Response) GetData() ListGroups200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListGroups200Response) GetDataOk() (*ListGroups200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListGroups200Response) SetData(v ListGroups200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ListGroups200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ListGroups200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListGroups200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListGroups200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListGroups200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ListGroups200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListGroups200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListGroups200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListGroups200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *ListGroups200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListGroups200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListGroups200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListGroups200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


