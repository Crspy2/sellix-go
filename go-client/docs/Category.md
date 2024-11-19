# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this category belongs. | [optional] 
**Title** | Pointer to **string** | Title of the category. | [optional] 
**Unlisted** | Pointer to **bool** | Whether or not the category is visible on the merchant storefront. | [optional] 
**SortPriority** | Pointer to **int32** | Sort order of this category. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the category. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the category has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the category. | [optional] 
**ProductsBound** | Pointer to [**[]CategoryProduct**](CategoryProduct.md) | Array of products. Please note that the product object contains limited fields, to get the full product object please use the Products API endpoint. | [optional] 
**ProductsCount** | Pointer to **int32** | How many products are present in the products_bound array | [optional] 
**GroupsBound** | Pointer to [**[]CategoryGroupsBoundInner**](CategoryGroupsBoundInner.md) | Array of groups. | [optional] 
**GroupsCount** | Pointer to **int32** | How many groups are present in the groups_bound array | [optional] 

## Methods

### NewCategory

`func NewCategory() *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Category) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Category) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Category) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Category) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Category) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *Category) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Category) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Category) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Category) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *Category) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Category) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Category) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Category) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnlisted

`func (o *Category) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *Category) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *Category) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *Category) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *Category) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *Category) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *Category) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *Category) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Category) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Category) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Category) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Category) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Category) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Category) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Category) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Category) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Category) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Category) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Category) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Category) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProductsBound

`func (o *Category) GetProductsBound() []CategoryProduct`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *Category) GetProductsBoundOk() (*[]CategoryProduct, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *Category) SetProductsBound(v []CategoryProduct)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *Category) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.

### GetProductsCount

`func (o *Category) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *Category) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *Category) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *Category) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.

### GetGroupsBound

`func (o *Category) GetGroupsBound() []CategoryGroupsBoundInner`

GetGroupsBound returns the GroupsBound field if non-nil, zero value otherwise.

### GetGroupsBoundOk

`func (o *Category) GetGroupsBoundOk() (*[]CategoryGroupsBoundInner, bool)`

GetGroupsBoundOk returns a tuple with the GroupsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsBound

`func (o *Category) SetGroupsBound(v []CategoryGroupsBoundInner)`

SetGroupsBound sets GroupsBound field to given value.

### HasGroupsBound

`func (o *Category) HasGroupsBound() bool`

HasGroupsBound returns a boolean if a field has been set.

### GetGroupsCount

`func (o *Category) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *Category) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *Category) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.

### HasGroupsCount

`func (o *Category) HasGroupsCount() bool`

HasGroupsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


