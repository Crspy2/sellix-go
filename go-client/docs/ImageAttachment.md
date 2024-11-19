# ImageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uniqid** | Pointer to **string** |  | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**Storage** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**Extension** | Pointer to **string** |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this resource belongs. | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the image. | [optional] 

## Methods

### NewImageAttachment

`func NewImageAttachment() *ImageAttachment`

NewImageAttachment instantiates a new ImageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAttachmentWithDefaults

`func NewImageAttachmentWithDefaults() *ImageAttachment`

NewImageAttachmentWithDefaults instantiates a new ImageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageAttachment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImageAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ImageAttachment) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ImageAttachment) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ImageAttachment) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ImageAttachment) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *ImageAttachment) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *ImageAttachment) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *ImageAttachment) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *ImageAttachment) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetStorage

`func (o *ImageAttachment) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ImageAttachment) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ImageAttachment) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ImageAttachment) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetType

`func (o *ImageAttachment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageAttachment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageAttachment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageAttachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ImageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalName

`func (o *ImageAttachment) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *ImageAttachment) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *ImageAttachment) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *ImageAttachment) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetExtension

`func (o *ImageAttachment) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *ImageAttachment) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *ImageAttachment) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *ImageAttachment) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetShopId

`func (o *ImageAttachment) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ImageAttachment) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ImageAttachment) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ImageAttachment) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductId

`func (o *ImageAttachment) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ImageAttachment) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ImageAttachment) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ImageAttachment) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSize

`func (o *ImageAttachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageAttachment) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageAttachment) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageAttachment) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ImageAttachment) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageAttachment) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageAttachment) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ImageAttachment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


