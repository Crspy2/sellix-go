# ListedGroup

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
**ImageName** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**ProductsBound** | Pointer to [**[]MiniProduct**](MiniProduct.md) |  | [optional] 
**ProductsCount** | Pointer to **int32** | The number of products in the group. | [optional] 

## Methods

### NewListedGroup

`func NewListedGroup() *ListedGroup`

NewListedGroup instantiates a new ListedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedGroupWithDefaults

`func NewListedGroupWithDefaults() *ListedGroup`

NewListedGroupWithDefaults instantiates a new ListedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListedGroup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListedGroup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListedGroup) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ListedGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ListedGroup) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ListedGroup) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ListedGroup) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ListedGroup) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *ListedGroup) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ListedGroup) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ListedGroup) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ListedGroup) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *ListedGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListedGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListedGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListedGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAttachment

`func (o *ListedGroup) GetImageAttachment() string`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *ListedGroup) GetImageAttachmentOk() (*string, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *ListedGroup) SetImageAttachment(v string)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *ListedGroup) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetUnlisted

`func (o *ListedGroup) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *ListedGroup) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *ListedGroup) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *ListedGroup) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetSortPriority

`func (o *ListedGroup) GetSortPriority() float32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *ListedGroup) GetSortPriorityOk() (*float32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *ListedGroup) SetSortPriority(v float32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *ListedGroup) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListedGroup) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListedGroup) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListedGroup) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListedGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListedGroup) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListedGroup) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListedGroup) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListedGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListedGroup) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListedGroup) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListedGroup) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListedGroup) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetImageName

`func (o *ListedGroup) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListedGroup) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListedGroup) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ListedGroup) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *ListedGroup) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *ListedGroup) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *ListedGroup) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *ListedGroup) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *ListedGroup) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *ListedGroup) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *ListedGroup) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *ListedGroup) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetProductsBound

`func (o *ListedGroup) GetProductsBound() []MiniProduct`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *ListedGroup) GetProductsBoundOk() (*[]MiniProduct, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *ListedGroup) SetProductsBound(v []MiniProduct)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *ListedGroup) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.

### GetProductsCount

`func (o *ListedGroup) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *ListedGroup) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *ListedGroup) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *ListedGroup) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


