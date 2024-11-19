# ListCoupons200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**ListCoupons200ResponseData**](ListCoupons200ResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Env** | Pointer to **string** |  | [optional] 

## Methods

### NewListCoupons200Response

`func NewListCoupons200Response() *ListCoupons200Response`

NewListCoupons200Response instantiates a new ListCoupons200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCoupons200ResponseWithDefaults

`func NewListCoupons200ResponseWithDefaults() *ListCoupons200Response`

NewListCoupons200ResponseWithDefaults instantiates a new ListCoupons200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListCoupons200Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCoupons200Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCoupons200Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListCoupons200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ListCoupons200Response) GetData() ListCoupons200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCoupons200Response) GetDataOk() (*ListCoupons200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCoupons200Response) SetData(v ListCoupons200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *ListCoupons200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ListCoupons200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListCoupons200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListCoupons200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListCoupons200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ListCoupons200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListCoupons200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListCoupons200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListCoupons200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetEnv

`func (o *ListCoupons200Response) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ListCoupons200Response) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ListCoupons200Response) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ListCoupons200Response) HasEnv() bool`

HasEnv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


