# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | The ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **float32** | The unique identifyer for the shop. | [optional] 
**Title** | Pointer to **string** | Group title. | [optional] 
**ImageAttachment** | Pointer to **string** | DEPRECATED: Unique id of the image attachment for this product. | [optional] 
**Unlisted** | Pointer to **bool** | Whether or not the group is unlisted. | [optional] 
**SortPriority** | Pointer to **float32** | The priority of the group on the storefront. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the group. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the group has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the group. | [optional] 
**ProductsBound** | Pointer to [**[]Product**](Product.md) |  | [optional] 
**ImageAttachmentInfo** | Pointer to [**ImageAttachment**](ImageAttachment.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the store the group is on. | [optional] 
**Theme** | Pointer to **string** | The theme of the group. | [optional] 
**DarkMode** | Pointer to **int32** | Whether or not the group is in dark mode. | [optional] 
**UiStyleConfiguration** | Pointer to **bool** | Whether or not the group has a custom UI style configuration. | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Group) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Group) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Group) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Group) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *Group) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Group) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Group) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Group) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *Group) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Group) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Group) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Group) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAttachment

`func (o *Group) GetImageAttachment() string`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *Group) GetImageAttachmentOk() (*string, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *Group) SetImageAttachment(v string)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *Group) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetUnlisted

`func (o *Group) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *Group) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *Group) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *Group) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *Group) GetSortPriority() float32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *Group) GetSortPriorityOk() (*float32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *Group) SetSortPriority(v float32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *Group) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Group) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Group) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Group) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Group) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Group) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Group) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Group) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Group) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Group) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Group) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Group) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Group) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProductsBound

`func (o *Group) GetProductsBound() []Product`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *Group) GetProductsBoundOk() (*[]Product, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *Group) SetProductsBound(v []Product)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *Group) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.

### GetImageAttachmentInfo

`func (o *Group) GetImageAttachmentInfo() ImageAttachment`

GetImageAttachmentInfo returns the ImageAttachmentInfo field if non-nil, zero value otherwise.

### GetImageAttachmentInfoOk

`func (o *Group) GetImageAttachmentInfoOk() (*ImageAttachment, bool)`

GetImageAttachmentInfoOk returns a tuple with the ImageAttachmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachmentInfo

`func (o *Group) SetImageAttachmentInfo(v ImageAttachment)`

SetImageAttachmentInfo sets ImageAttachmentInfo field to given value.

### HasImageAttachmentInfo

`func (o *Group) HasImageAttachmentInfo() bool`

HasImageAttachmentInfo returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTheme

`func (o *Group) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Group) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Group) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Group) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetDarkMode

`func (o *Group) GetDarkMode() int32`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *Group) GetDarkModeOk() (*int32, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *Group) SetDarkMode(v int32)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *Group) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetUiStyleConfiguration

`func (o *Group) GetUiStyleConfiguration() bool`

GetUiStyleConfiguration returns the UiStyleConfiguration field if non-nil, zero value otherwise.

### GetUiStyleConfigurationOk

`func (o *Group) GetUiStyleConfigurationOk() (*bool, bool)`

GetUiStyleConfigurationOk returns a tuple with the UiStyleConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiStyleConfiguration

`func (o *Group) SetUiStyleConfiguration(v bool)`

SetUiStyleConfiguration sets UiStyleConfiguration field to given value.

### HasUiStyleConfiguration

`func (o *Group) HasUiStyleConfiguration() bool`

HasUiStyleConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


