# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this coupon belongs. | [optional] 
**Type** | Pointer to **string** | Coupon type, whether it&#39;s for a product or subscription. SUBSCRIPTION coupons are only for Sellix-own subscriptions, this field will always return PRODUCT. | [optional] 
**CouponType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | Coupon code to be used during the checkout phase. | [optional] 
**StripePromoId** | Pointer to **string** | The stripe ID for the promotion | [optional] 
**StripeCouponId** | Pointer to **string** | The stripe ID for the coupon | [optional] 
**UseType** | Pointer to **string** | If LIMITED, an array of products must be specified as this coupon will be used only with them. | [optional] 
**Discount** | Pointer to **float64** | Discount value for this coupon. | [optional] 
**Currency** | Pointer to **string** | Currency of the coupon discount value. | [optional] 
**Used** | Pointer to **int32** | How many times this coupon has been used. | [optional] 
**DisabledWithVolumeDiscounts** | Pointer to **bool** | Whether or not this coupon is valid if a volume discount is applied. | [optional] 
**AllRecurringBillInvoices** | Pointer to **bool** | Whether or not this coupon should be applied for each product SUBSCRIPTION renewal. | [optional] 
**MaxUses** | Pointer to **int32** | Maximum usage for this coupon. | [optional] 
**SmartContractAddress** | Pointer to **string** | The smart contract address for the coupon | [optional] 
**TokenId** | Pointer to **string** | The token ID for the coupon | [optional] 
**Blockchain** | Pointer to [**Blockchain**](Blockchain.md) |  | [optional] 
**ExpireAt** | Pointer to **string** | The datetime object corresponding to the expiration time of the coupon | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the coupon. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the category has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the coupon. | [optional] 
**ProductsBound** | Pointer to **[]string** | Array of product uniqids. Differs from the categories API as this endpoint does not need specific details about a product. Use the products API to get details about a single product. | [optional] 
**ProductsBoundExtended** | Pointer to **[]string** | Details about the products in the products_bound field. | [optional] 
**ProductsCount** | Pointer to **int32** | How many products are present in the products_bound array | [optional] 

## Methods

### NewCoupon

`func NewCoupon() *Coupon`

NewCoupon instantiates a new Coupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponWithDefaults

`func NewCouponWithDefaults() *Coupon`

NewCouponWithDefaults instantiates a new Coupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Coupon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Coupon) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Coupon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Coupon) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Coupon) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Coupon) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Coupon) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *Coupon) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Coupon) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Coupon) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Coupon) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *Coupon) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Coupon) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Coupon) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Coupon) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCouponType

`func (o *Coupon) GetCouponType() string`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *Coupon) GetCouponTypeOk() (*string, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *Coupon) SetCouponType(v string)`

SetCouponType sets CouponType field to given value.

### HasCouponType

`func (o *Coupon) HasCouponType() bool`

HasCouponType returns a boolean if a field has been set.

### GetCode

`func (o *Coupon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Coupon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Coupon) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Coupon) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStripePromoId

`func (o *Coupon) GetStripePromoId() string`

GetStripePromoId returns the StripePromoId field if non-nil, zero value otherwise.

### GetStripePromoIdOk

`func (o *Coupon) GetStripePromoIdOk() (*string, bool)`

GetStripePromoIdOk returns a tuple with the StripePromoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePromoId

`func (o *Coupon) SetStripePromoId(v string)`

SetStripePromoId sets StripePromoId field to given value.

### HasStripePromoId

`func (o *Coupon) HasStripePromoId() bool`

HasStripePromoId returns a boolean if a field has been set.

### GetStripeCouponId

`func (o *Coupon) GetStripeCouponId() string`

GetStripeCouponId returns the StripeCouponId field if non-nil, zero value otherwise.

### GetStripeCouponIdOk

`func (o *Coupon) GetStripeCouponIdOk() (*string, bool)`

GetStripeCouponIdOk returns a tuple with the StripeCouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCouponId

`func (o *Coupon) SetStripeCouponId(v string)`

SetStripeCouponId sets StripeCouponId field to given value.

### HasStripeCouponId

`func (o *Coupon) HasStripeCouponId() bool`

HasStripeCouponId returns a boolean if a field has been set.

### GetUseType

`func (o *Coupon) GetUseType() string`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *Coupon) GetUseTypeOk() (*string, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *Coupon) SetUseType(v string)`

SetUseType sets UseType field to given value.

### HasUseType

`func (o *Coupon) HasUseType() bool`

HasUseType returns a boolean if a field has been set.

### GetDiscount

`func (o *Coupon) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Coupon) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Coupon) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Coupon) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCurrency

`func (o *Coupon) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Coupon) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Coupon) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Coupon) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsed

`func (o *Coupon) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Coupon) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Coupon) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Coupon) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetDisabledWithVolumeDiscounts

`func (o *Coupon) GetDisabledWithVolumeDiscounts() bool`

GetDisabledWithVolumeDiscounts returns the DisabledWithVolumeDiscounts field if non-nil, zero value otherwise.

### GetDisabledWithVolumeDiscountsOk

`func (o *Coupon) GetDisabledWithVolumeDiscountsOk() (*bool, bool)`

GetDisabledWithVolumeDiscountsOk returns a tuple with the DisabledWithVolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWithVolumeDiscounts

`func (o *Coupon) SetDisabledWithVolumeDiscounts(v bool)`

SetDisabledWithVolumeDiscounts sets DisabledWithVolumeDiscounts field to given value.

### HasDisabledWithVolumeDiscounts

`func (o *Coupon) HasDisabledWithVolumeDiscounts() bool`

HasDisabledWithVolumeDiscounts returns a boolean if a field has been set.

### GetAllRecurringBillInvoices

`func (o *Coupon) GetAllRecurringBillInvoices() bool`

GetAllRecurringBillInvoices returns the AllRecurringBillInvoices field if non-nil, zero value otherwise.

### GetAllRecurringBillInvoicesOk

`func (o *Coupon) GetAllRecurringBillInvoicesOk() (*bool, bool)`

GetAllRecurringBillInvoicesOk returns a tuple with the AllRecurringBillInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecurringBillInvoices

`func (o *Coupon) SetAllRecurringBillInvoices(v bool)`

SetAllRecurringBillInvoices sets AllRecurringBillInvoices field to given value.

### HasAllRecurringBillInvoices

`func (o *Coupon) HasAllRecurringBillInvoices() bool`

HasAllRecurringBillInvoices returns a boolean if a field has been set.

### GetMaxUses

`func (o *Coupon) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *Coupon) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *Coupon) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *Coupon) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetSmartContractAddress

`func (o *Coupon) GetSmartContractAddress() string`

GetSmartContractAddress returns the SmartContractAddress field if non-nil, zero value otherwise.

### GetSmartContractAddressOk

`func (o *Coupon) GetSmartContractAddressOk() (*string, bool)`

GetSmartContractAddressOk returns a tuple with the SmartContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractAddress

`func (o *Coupon) SetSmartContractAddress(v string)`

SetSmartContractAddress sets SmartContractAddress field to given value.

### HasSmartContractAddress

`func (o *Coupon) HasSmartContractAddress() bool`

HasSmartContractAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *Coupon) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Coupon) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Coupon) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Coupon) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetBlockchain

`func (o *Coupon) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *Coupon) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *Coupon) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *Coupon) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetExpireAt

`func (o *Coupon) GetExpireAt() string`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *Coupon) GetExpireAtOk() (*string, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *Coupon) SetExpireAt(v string)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *Coupon) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Coupon) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Coupon) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Coupon) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Coupon) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Coupon) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Coupon) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Coupon) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Coupon) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Coupon) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Coupon) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Coupon) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Coupon) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProductsBound

`func (o *Coupon) GetProductsBound() []string`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *Coupon) GetProductsBoundOk() (*[]string, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *Coupon) SetProductsBound(v []string)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *Coupon) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.

### GetProductsBoundExtended

`func (o *Coupon) GetProductsBoundExtended() []string`

GetProductsBoundExtended returns the ProductsBoundExtended field if non-nil, zero value otherwise.

### GetProductsBoundExtendedOk

`func (o *Coupon) GetProductsBoundExtendedOk() (*[]string, bool)`

GetProductsBoundExtendedOk returns a tuple with the ProductsBoundExtended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBoundExtended

`func (o *Coupon) SetProductsBoundExtended(v []string)`

SetProductsBoundExtended sets ProductsBoundExtended field to given value.

### HasProductsBoundExtended

`func (o *Coupon) HasProductsBoundExtended() bool`

HasProductsBoundExtended returns a boolean if a field has been set.

### GetProductsCount

`func (o *Coupon) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *Coupon) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *Coupon) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *Coupon) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


