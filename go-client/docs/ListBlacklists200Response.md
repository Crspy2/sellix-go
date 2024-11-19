# ListBlacklists200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**ListBlacklists200ResponseData**](ListBlacklists200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewListBlacklists200Response

`func NewListBlacklists200Response() *ListBlacklists200Response`

NewListBlacklists200Response instantiates a new ListBlacklists200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlacklists200ResponseWithDefaults

`func NewListBlacklists200ResponseWithDefaults() *ListBlacklists200Response`

NewListBlacklists200ResponseWithDefaults instantiates a new ListBlacklists200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListBlacklists200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBlacklists200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBlacklists200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBlacklists200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ListBlacklists200Response) GetData() ListBlacklists200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListBlacklists200Response) GetDataOk() (*ListBlacklists200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListBlacklists200Response) SetData(v ListBlacklists200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ListBlacklists200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ListBlacklists200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListBlacklists200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListBlacklists200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListBlacklists200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ListBlacklists200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListBlacklists200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListBlacklists200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListBlacklists200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *ListBlacklists200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListBlacklists200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListBlacklists200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListBlacklists200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


