# MiniProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Price** | Pointer to **string** | Product price. | [optional] 
**PriceDisplay** | Pointer to **string** | Product price in currency. | [optional] 
**PriceDiscount** | Pointer to **float32** | The discount price on the purchased goods. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Unlisted** | Pointer to **int32** | Whether or not the product is shown on the storefront. | [optional] 
**Title** | Pointer to **string** | The title of the product | [optional] 
**ImageAttachment** | Pointer to **string** | DEPRECATED: Unique id of the image attachment for this product. | [optional] 
**Description** | Pointer to **string** | The description of the product | [optional] 
**QuantityMin** | Pointer to **int32** | Minimum quantity purchasable of this product. | [optional] 
**QuantityMax** | Pointer to **int32** | Maximum quantity purchasable of this product. | [optional] 
**QuantityWarning** | Pointer to **int32** | At which product quantity should we send a webhook and email warning the merchant. | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldsArrayInner**](CustomFieldsArrayInner.md) |  | [optional] 
**Type** | Pointer to **string** | Product type. | [optional] 
**ShopId** | Pointer to **float32** | The shop ID to which this resource belongs. | [optional] 
**Gateways** | Pointer to **string** | A list of gateways available on this product. | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Minimum number of confirmations for a crypto payment to be accepted. | [optional] 
**Private** | Pointer to **int32** | Whether or not the product is private and cannot be purchased. | [optional] 
**Stock** | Pointer to **int32** | Stock of the current product, can be -1 for unlimited stock. | [optional] 
**SortPriority** | Pointer to **int32** | Sort order of this product. | [optional] 
**OnHold** | Pointer to **int32** | Whether or not the product is on hold and purchaseable. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the product. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**ProductId** | Pointer to **string** | The uniqid of the product | [optional] 
**GroupId** | Pointer to **string** | If the product is in a group, this is the group_id of that group. | [optional] 
**AverageScore** | Pointer to **string** | The average rating of the product. | [optional] 

## Methods

### NewMiniProduct

`func NewMiniProduct() *MiniProduct`

NewMiniProduct instantiates a new MiniProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiniProductWithDefaults

`func NewMiniProductWithDefaults() *MiniProduct`

NewMiniProductWithDefaults instantiates a new MiniProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqid

`func (o *MiniProduct) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *MiniProduct) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *MiniProduct) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *MiniProduct) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetPrice

`func (o *MiniProduct) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MiniProduct) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MiniProduct) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MiniProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplay

`func (o *MiniProduct) GetPriceDisplay() string`

GetPriceDisplay returns the PriceDisplay field if non-nil, zero value otherwise.

### GetPriceDisplayOk

`func (o *MiniProduct) GetPriceDisplayOk() (*string, bool)`

GetPriceDisplayOk returns a tuple with the PriceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplay

`func (o *MiniProduct) SetPriceDisplay(v string)`

SetPriceDisplay sets PriceDisplay field to given value.

### HasPriceDisplay

`func (o *MiniProduct) HasPriceDisplay() bool`

HasPriceDisplay returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *MiniProduct) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *MiniProduct) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *MiniProduct) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *MiniProduct) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetCurrency

`func (o *MiniProduct) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MiniProduct) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MiniProduct) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MiniProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnlisted

`func (o *MiniProduct) GetUnlisted() int32`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *MiniProduct) GetUnlistedOk() (*int32, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *MiniProduct) SetUnlisted(v int32)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *MiniProduct) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetTitle

`func (o *MiniProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MiniProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MiniProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MiniProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAttachment

`func (o *MiniProduct) GetImageAttachment() string`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *MiniProduct) GetImageAttachmentOk() (*string, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *MiniProduct) SetImageAttachment(v string)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *MiniProduct) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetDescription

`func (o *MiniProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MiniProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MiniProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MiniProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantityMin

`func (o *MiniProduct) GetQuantityMin() int32`

GetQuantityMin returns the QuantityMin field if non-nil, zero value otherwise.

### GetQuantityMinOk

`func (o *MiniProduct) GetQuantityMinOk() (*int32, bool)`

GetQuantityMinOk returns a tuple with the QuantityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMin

`func (o *MiniProduct) SetQuantityMin(v int32)`

SetQuantityMin sets QuantityMin field to given value.

### HasQuantityMin

`func (o *MiniProduct) HasQuantityMin() bool`

HasQuantityMin returns a boolean if a field has been set.

### GetQuantityMax

`func (o *MiniProduct) GetQuantityMax() int32`

GetQuantityMax returns the QuantityMax field if non-nil, zero value otherwise.

### GetQuantityMaxOk

`func (o *MiniProduct) GetQuantityMaxOk() (*int32, bool)`

GetQuantityMaxOk returns a tuple with the QuantityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMax

`func (o *MiniProduct) SetQuantityMax(v int32)`

SetQuantityMax sets QuantityMax field to given value.

### HasQuantityMax

`func (o *MiniProduct) HasQuantityMax() bool`

HasQuantityMax returns a boolean if a field has been set.

### GetQuantityWarning

`func (o *MiniProduct) GetQuantityWarning() int32`

GetQuantityWarning returns the QuantityWarning field if non-nil, zero value otherwise.

### GetQuantityWarningOk

`func (o *MiniProduct) GetQuantityWarningOk() (*int32, bool)`

GetQuantityWarningOk returns a tuple with the QuantityWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWarning

`func (o *MiniProduct) SetQuantityWarning(v int32)`

SetQuantityWarning sets QuantityWarning field to given value.

### HasQuantityWarning

`func (o *MiniProduct) HasQuantityWarning() bool`

HasQuantityWarning returns a boolean if a field has been set.

### GetCustomFields

`func (o *MiniProduct) GetCustomFields() []CustomFieldsArrayInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MiniProduct) GetCustomFieldsOk() (*[]CustomFieldsArrayInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MiniProduct) SetCustomFields(v []CustomFieldsArrayInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MiniProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetType

`func (o *MiniProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MiniProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MiniProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MiniProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetShopId

`func (o *MiniProduct) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *MiniProduct) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *MiniProduct) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *MiniProduct) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetGateways

`func (o *MiniProduct) GetGateways() string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *MiniProduct) GetGatewaysOk() (*string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *MiniProduct) SetGateways(v string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *MiniProduct) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *MiniProduct) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *MiniProduct) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *MiniProduct) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *MiniProduct) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetPrivate

`func (o *MiniProduct) GetPrivate() int32`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *MiniProduct) GetPrivateOk() (*int32, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *MiniProduct) SetPrivate(v int32)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *MiniProduct) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetStock

`func (o *MiniProduct) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *MiniProduct) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *MiniProduct) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *MiniProduct) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetSortPriority

`func (o *MiniProduct) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *MiniProduct) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *MiniProduct) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *MiniProduct) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetOnHold

`func (o *MiniProduct) GetOnHold() int32`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *MiniProduct) GetOnHoldOk() (*int32, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *MiniProduct) SetOnHold(v int32)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *MiniProduct) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MiniProduct) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MiniProduct) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MiniProduct) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MiniProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MiniProduct) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MiniProduct) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MiniProduct) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MiniProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProductId

`func (o *MiniProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *MiniProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *MiniProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *MiniProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGroupId

`func (o *MiniProduct) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MiniProduct) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MiniProduct) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MiniProduct) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetAverageScore

`func (o *MiniProduct) GetAverageScore() string`

GetAverageScore returns the AverageScore field if non-nil, zero value otherwise.

### GetAverageScoreOk

`func (o *MiniProduct) GetAverageScoreOk() (*string, bool)`

GetAverageScoreOk returns a tuple with the AverageScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageScore

`func (o *MiniProduct) SetAverageScore(v string)`

SetAverageScore sets AverageScore field to given value.

### HasAverageScore

`func (o *MiniProduct) HasAverageScore() bool`

HasAverageScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


