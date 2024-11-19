# MiniGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | The ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **float32** | The unique identifyer for the shop. | [optional] 
**Title** | Pointer to **string** | Group title. | [optional] 
**ImageAttachment** | Pointer to **string** | DEPRECATED: Unique id of the image attachment for this product. | [optional] 
**Unlisted** | Pointer to **float32** | Whether or not the group is unlisted. | [optional] 
**SortPriority** | Pointer to **float32** | The priority of the group on the storefront. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the group. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the group has been edited. | [optional] 
**CategoryId** | Pointer to **string** | The category the group is a part of | [optional] 
**GroupId** | Pointer to **string** | The unique identifyer of the current group | [optional] 

## Methods

### NewMiniGroup

`func NewMiniGroup() *MiniGroup`

NewMiniGroup instantiates a new MiniGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiniGroupWithDefaults

`func NewMiniGroupWithDefaults() *MiniGroup`

NewMiniGroupWithDefaults instantiates a new MiniGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MiniGroup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MiniGroup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MiniGroup) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MiniGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *MiniGroup) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *MiniGroup) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *MiniGroup) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *MiniGroup) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *MiniGroup) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *MiniGroup) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *MiniGroup) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *MiniGroup) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *MiniGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MiniGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MiniGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MiniGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAttachment

`func (o *MiniGroup) GetImageAttachment() string`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *MiniGroup) GetImageAttachmentOk() (*string, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *MiniGroup) SetImageAttachment(v string)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *MiniGroup) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetUnlisted

`func (o *MiniGroup) GetUnlisted() float32`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *MiniGroup) GetUnlistedOk() (*float32, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *MiniGroup) SetUnlisted(v float32)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *MiniGroup) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *MiniGroup) GetSortPriority() float32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *MiniGroup) GetSortPriorityOk() (*float32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *MiniGroup) SetSortPriority(v float32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *MiniGroup) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MiniGroup) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MiniGroup) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MiniGroup) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MiniGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MiniGroup) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MiniGroup) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MiniGroup) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MiniGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCategoryId

`func (o *MiniGroup) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *MiniGroup) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *MiniGroup) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *MiniGroup) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetGroupId

`func (o *MiniGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MiniGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MiniGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MiniGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


