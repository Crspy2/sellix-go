# CreateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Unlisted** | Pointer to **bool** |  | [optional] 
**SortPriority** | Pointer to **int32** |  | [optional] 
**ProductsBound** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGroupRequest

`func NewCreateGroupRequest(title string, ) *CreateGroupRequest`

NewCreateGroupRequest instantiates a new CreateGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRequestWithDefaults

`func NewCreateGroupRequestWithDefaults() *CreateGroupRequest`

NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateGroupRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGroupRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGroupRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUnlisted

`func (o *CreateGroupRequest) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *CreateGroupRequest) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *CreateGroupRequest) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *CreateGroupRequest) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *CreateGroupRequest) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *CreateGroupRequest) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *CreateGroupRequest) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *CreateGroupRequest) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetProductsBound

`func (o *CreateGroupRequest) GetProductsBound() []string`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *CreateGroupRequest) GetProductsBoundOk() (*[]string, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *CreateGroupRequest) SetProductsBound(v []string)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *CreateGroupRequest) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


