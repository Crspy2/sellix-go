# CreateCategoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Unlisted** | Pointer to **bool** |  | [optional] 
**SortPriority** | Pointer to **int32** |  | [optional] 
**ProductsBound** | Pointer to **[]string** |  | [optional] 
**GroupsBound** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCategoryRequest

`func NewCreateCategoryRequest(title string, ) *CreateCategoryRequest`

NewCreateCategoryRequest instantiates a new CreateCategoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCategoryRequestWithDefaults

`func NewCreateCategoryRequestWithDefaults() *CreateCategoryRequest`

NewCreateCategoryRequestWithDefaults instantiates a new CreateCategoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateCategoryRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCategoryRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCategoryRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUnlisted

`func (o *CreateCategoryRequest) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *CreateCategoryRequest) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *CreateCategoryRequest) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *CreateCategoryRequest) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *CreateCategoryRequest) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *CreateCategoryRequest) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *CreateCategoryRequest) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *CreateCategoryRequest) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetProductsBound

`func (o *CreateCategoryRequest) GetProductsBound() []string`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *CreateCategoryRequest) GetProductsBoundOk() (*[]string, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *CreateCategoryRequest) SetProductsBound(v []string)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *CreateCategoryRequest) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.

### GetGroupsBound

`func (o *CreateCategoryRequest) GetGroupsBound() []string`

GetGroupsBound returns the GroupsBound field if non-nil, zero value otherwise.

### GetGroupsBoundOk

`func (o *CreateCategoryRequest) GetGroupsBoundOk() (*[]string, bool)`

GetGroupsBoundOk returns a tuple with the GroupsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsBound

`func (o *CreateCategoryRequest) SetGroupsBound(v []string)`

SetGroupsBound sets GroupsBound field to given value.

### HasGroupsBound

`func (o *CreateCategoryRequest) HasGroupsBound() bool`

HasGroupsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


