# CategoryProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Price** | Pointer to **float64** | Product price. | [optional] 
**PriceDisplay** | Pointer to **float64** | Product price in currency. | [optional] 
**PriceDiscount** | Pointer to **float32** | The discount price on the purchased goods. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Unlisted** | Pointer to **int32** | Whether the product is listed on the storefront | [optional] 
**Title** | Pointer to **string** | The product title | [optional] 
**ImageAttachment** | Pointer to **string** | DEPRECATED: Unique id of the image attachment for this product. | [optional] 
**Description** | Pointer to **string** | Product description. | [optional] 
**QuantityMin** | Pointer to **int32** | Minimum quantity purchasable of this product. | [optional] 
**QuantityMax** | Pointer to **int32** | Maximum quantity purchasable of this product. | [optional] 
**QuantityWarning** | Pointer to **int32** | At which product quantity should we send a webhook and email warning the merchant. | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldsArrayInner**](CustomFieldsArrayInner.md) |  | [optional] 
**Type** | Pointer to **string** | Product type. | [optional] 
**ShopId** | Pointer to **float32** | The shop ID to which this resource belongs. | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Minimum number of confirmations for a crypto payment to be accepted. | [optional] 
**Private** | Pointer to **bool** | If private is true, the product is hidden on the storefront and cannot be bought with a direct link. | [optional] 
**Stock** | Pointer to **int32** | Stock of the current product, can be -1 for unlimited stock. | [optional] 
**SortPriority** | Pointer to **int32** | Sort order of this product. | [optional] 
**OnHold** | Pointer to **int32** | Whether the product cannot be bought but is shown in the storefront. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the product. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**ImageName** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**CategoryId** | Pointer to **string** | The category the product is a part of | [optional] 
**ProductId** | Pointer to **string** | The id of the current product | [optional] 

## Methods

### NewCategoryProduct

`func NewCategoryProduct() *CategoryProduct`

NewCategoryProduct instantiates a new CategoryProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryProductWithDefaults

`func NewCategoryProductWithDefaults() *CategoryProduct`

NewCategoryProductWithDefaults instantiates a new CategoryProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqid

`func (o *CategoryProduct) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *CategoryProduct) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *CategoryProduct) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *CategoryProduct) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetPrice

`func (o *CategoryProduct) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CategoryProduct) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CategoryProduct) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CategoryProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplay

`func (o *CategoryProduct) GetPriceDisplay() float64`

GetPriceDisplay returns the PriceDisplay field if non-nil, zero value otherwise.

### GetPriceDisplayOk

`func (o *CategoryProduct) GetPriceDisplayOk() (*float64, bool)`

GetPriceDisplayOk returns a tuple with the PriceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplay

`func (o *CategoryProduct) SetPriceDisplay(v float64)`

SetPriceDisplay sets PriceDisplay field to given value.

### HasPriceDisplay

`func (o *CategoryProduct) HasPriceDisplay() bool`

HasPriceDisplay returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *CategoryProduct) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *CategoryProduct) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *CategoryProduct) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *CategoryProduct) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetCurrency

`func (o *CategoryProduct) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CategoryProduct) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CategoryProduct) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CategoryProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnlisted

`func (o *CategoryProduct) GetUnlisted() int32`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *CategoryProduct) GetUnlistedOk() (*int32, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *CategoryProduct) SetUnlisted(v int32)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *CategoryProduct) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetTitle

`func (o *CategoryProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CategoryProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CategoryProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CategoryProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAttachment

`func (o *CategoryProduct) GetImageAttachment() string`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *CategoryProduct) GetImageAttachmentOk() (*string, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *CategoryProduct) SetImageAttachment(v string)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *CategoryProduct) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantityMin

`func (o *CategoryProduct) GetQuantityMin() int32`

GetQuantityMin returns the QuantityMin field if non-nil, zero value otherwise.

### GetQuantityMinOk

`func (o *CategoryProduct) GetQuantityMinOk() (*int32, bool)`

GetQuantityMinOk returns a tuple with the QuantityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMin

`func (o *CategoryProduct) SetQuantityMin(v int32)`

SetQuantityMin sets QuantityMin field to given value.

### HasQuantityMin

`func (o *CategoryProduct) HasQuantityMin() bool`

HasQuantityMin returns a boolean if a field has been set.

### GetQuantityMax

`func (o *CategoryProduct) GetQuantityMax() int32`

GetQuantityMax returns the QuantityMax field if non-nil, zero value otherwise.

### GetQuantityMaxOk

`func (o *CategoryProduct) GetQuantityMaxOk() (*int32, bool)`

GetQuantityMaxOk returns a tuple with the QuantityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMax

`func (o *CategoryProduct) SetQuantityMax(v int32)`

SetQuantityMax sets QuantityMax field to given value.

### HasQuantityMax

`func (o *CategoryProduct) HasQuantityMax() bool`

HasQuantityMax returns a boolean if a field has been set.

### GetQuantityWarning

`func (o *CategoryProduct) GetQuantityWarning() int32`

GetQuantityWarning returns the QuantityWarning field if non-nil, zero value otherwise.

### GetQuantityWarningOk

`func (o *CategoryProduct) GetQuantityWarningOk() (*int32, bool)`

GetQuantityWarningOk returns a tuple with the QuantityWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWarning

`func (o *CategoryProduct) SetQuantityWarning(v int32)`

SetQuantityWarning sets QuantityWarning field to given value.

### HasQuantityWarning

`func (o *CategoryProduct) HasQuantityWarning() bool`

HasQuantityWarning returns a boolean if a field has been set.

### GetCustomFields

`func (o *CategoryProduct) GetCustomFields() []CustomFieldsArrayInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CategoryProduct) GetCustomFieldsOk() (*[]CustomFieldsArrayInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CategoryProduct) SetCustomFields(v []CustomFieldsArrayInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CategoryProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetType

`func (o *CategoryProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoryProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoryProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CategoryProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetShopId

`func (o *CategoryProduct) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *CategoryProduct) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *CategoryProduct) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *CategoryProduct) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetGateways

`func (o *CategoryProduct) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CategoryProduct) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CategoryProduct) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CategoryProduct) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *CategoryProduct) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *CategoryProduct) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *CategoryProduct) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *CategoryProduct) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetPrivate

`func (o *CategoryProduct) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CategoryProduct) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CategoryProduct) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CategoryProduct) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetStock

`func (o *CategoryProduct) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CategoryProduct) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CategoryProduct) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CategoryProduct) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetSortPriority

`func (o *CategoryProduct) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *CategoryProduct) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *CategoryProduct) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *CategoryProduct) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetOnHold

`func (o *CategoryProduct) GetOnHold() int32`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *CategoryProduct) GetOnHoldOk() (*int32, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *CategoryProduct) SetOnHold(v int32)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *CategoryProduct) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CategoryProduct) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CategoryProduct) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CategoryProduct) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CategoryProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CategoryProduct) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CategoryProduct) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CategoryProduct) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CategoryProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetImageName

`func (o *CategoryProduct) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CategoryProduct) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CategoryProduct) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *CategoryProduct) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *CategoryProduct) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *CategoryProduct) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *CategoryProduct) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *CategoryProduct) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *CategoryProduct) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *CategoryProduct) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *CategoryProduct) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *CategoryProduct) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetCategoryId

`func (o *CategoryProduct) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryProduct) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryProduct) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CategoryProduct) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetProductId

`func (o *CategoryProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CategoryProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CategoryProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CategoryProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


