# GetBlacklist200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**GetBlacklist200ResponseData**](GetBlacklist200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBlacklist200Response

`func NewGetBlacklist200Response() *GetBlacklist200Response`

NewGetBlacklist200Response instantiates a new GetBlacklist200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlacklist200ResponseWithDefaults

`func NewGetBlacklist200ResponseWithDefaults() *GetBlacklist200Response`

NewGetBlacklist200ResponseWithDefaults instantiates a new GetBlacklist200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetBlacklist200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBlacklist200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBlacklist200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBlacklist200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetBlacklist200Response) GetData() GetBlacklist200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBlacklist200Response) GetDataOk() (*GetBlacklist200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBlacklist200Response) SetData(v GetBlacklist200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetBlacklist200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *GetBlacklist200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetBlacklist200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetBlacklist200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetBlacklist200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *GetBlacklist200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetBlacklist200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetBlacklist200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetBlacklist200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *GetBlacklist200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GetBlacklist200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GetBlacklist200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GetBlacklist200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


