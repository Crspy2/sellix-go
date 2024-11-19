# GetOrder200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 

## Methods

### NewGetOrder200ResponseData

`func NewGetOrder200ResponseData() *GetOrder200ResponseData`

NewGetOrder200ResponseData instantiates a new GetOrder200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrder200ResponseDataWithDefaults

`func NewGetOrder200ResponseDataWithDefaults() *GetOrder200ResponseData`

NewGetOrder200ResponseDataWithDefaults instantiates a new GetOrder200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GetOrder200ResponseData) GetOrder() Invoice`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetOrder200ResponseData) GetOrderOk() (*Invoice, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetOrder200ResponseData) SetOrder(v Invoice)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetOrder200ResponseData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


