# CouponListing

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
**ProductsCount** | Pointer to **int32** | How many products are present in the products_bound array | [optional] 

## Methods

### NewCouponListing

`func NewCouponListing() *CouponListing`

NewCouponListing instantiates a new CouponListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponListingWithDefaults

`func NewCouponListingWithDefaults() *CouponListing`

NewCouponListingWithDefaults instantiates a new CouponListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponListing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponListing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponListing) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CouponListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *CouponListing) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *CouponListing) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *CouponListing) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *CouponListing) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *CouponListing) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *CouponListing) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *CouponListing) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *CouponListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *CouponListing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponListing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponListing) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CouponListing) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCouponType

`func (o *CouponListing) GetCouponType() string`

GetCouponType returns the CouponType field if non-nil, zero value otherwise.

### GetCouponTypeOk

`func (o *CouponListing) GetCouponTypeOk() (*string, bool)`

GetCouponTypeOk returns a tuple with the CouponType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponType

`func (o *CouponListing) SetCouponType(v string)`

SetCouponType sets CouponType field to given value.

### HasCouponType

`func (o *CouponListing) HasCouponType() bool`

HasCouponType returns a boolean if a field has been set.

### GetCode

`func (o *CouponListing) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouponListing) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouponListing) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CouponListing) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStripePromoId

`func (o *CouponListing) GetStripePromoId() string`

GetStripePromoId returns the StripePromoId field if non-nil, zero value otherwise.

### GetStripePromoIdOk

`func (o *CouponListing) GetStripePromoIdOk() (*string, bool)`

GetStripePromoIdOk returns a tuple with the StripePromoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePromoId

`func (o *CouponListing) SetStripePromoId(v string)`

SetStripePromoId sets StripePromoId field to given value.

### HasStripePromoId

`func (o *CouponListing) HasStripePromoId() bool`

HasStripePromoId returns a boolean if a field has been set.

### GetStripeCouponId

`func (o *CouponListing) GetStripeCouponId() string`

GetStripeCouponId returns the StripeCouponId field if non-nil, zero value otherwise.

### GetStripeCouponIdOk

`func (o *CouponListing) GetStripeCouponIdOk() (*string, bool)`

GetStripeCouponIdOk returns a tuple with the StripeCouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCouponId

`func (o *CouponListing) SetStripeCouponId(v string)`

SetStripeCouponId sets StripeCouponId field to given value.

### HasStripeCouponId

`func (o *CouponListing) HasStripeCouponId() bool`

HasStripeCouponId returns a boolean if a field has been set.

### GetUseType

`func (o *CouponListing) GetUseType() string`

GetUseType returns the UseType field if non-nil, zero value otherwise.

### GetUseTypeOk

`func (o *CouponListing) GetUseTypeOk() (*string, bool)`

GetUseTypeOk returns a tuple with the UseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseType

`func (o *CouponListing) SetUseType(v string)`

SetUseType sets UseType field to given value.

### HasUseType

`func (o *CouponListing) HasUseType() bool`

HasUseType returns a boolean if a field has been set.

### GetDiscount

`func (o *CouponListing) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CouponListing) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CouponListing) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *CouponListing) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCurrency

`func (o *CouponListing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CouponListing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CouponListing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CouponListing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUsed

`func (o *CouponListing) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *CouponListing) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *CouponListing) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *CouponListing) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetDisabledWithVolumeDiscounts

`func (o *CouponListing) GetDisabledWithVolumeDiscounts() bool`

GetDisabledWithVolumeDiscounts returns the DisabledWithVolumeDiscounts field if non-nil, zero value otherwise.

### GetDisabledWithVolumeDiscountsOk

`func (o *CouponListing) GetDisabledWithVolumeDiscountsOk() (*bool, bool)`

GetDisabledWithVolumeDiscountsOk returns a tuple with the DisabledWithVolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWithVolumeDiscounts

`func (o *CouponListing) SetDisabledWithVolumeDiscounts(v bool)`

SetDisabledWithVolumeDiscounts sets DisabledWithVolumeDiscounts field to given value.

### HasDisabledWithVolumeDiscounts

`func (o *CouponListing) HasDisabledWithVolumeDiscounts() bool`

HasDisabledWithVolumeDiscounts returns a boolean if a field has been set.

### GetAllRecurringBillInvoices

`func (o *CouponListing) GetAllRecurringBillInvoices() bool`

GetAllRecurringBillInvoices returns the AllRecurringBillInvoices field if non-nil, zero value otherwise.

### GetAllRecurringBillInvoicesOk

`func (o *CouponListing) GetAllRecurringBillInvoicesOk() (*bool, bool)`

GetAllRecurringBillInvoicesOk returns a tuple with the AllRecurringBillInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecurringBillInvoices

`func (o *CouponListing) SetAllRecurringBillInvoices(v bool)`

SetAllRecurringBillInvoices sets AllRecurringBillInvoices field to given value.

### HasAllRecurringBillInvoices

`func (o *CouponListing) HasAllRecurringBillInvoices() bool`

HasAllRecurringBillInvoices returns a boolean if a field has been set.

### GetMaxUses

`func (o *CouponListing) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *CouponListing) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *CouponListing) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *CouponListing) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetSmartContractAddress

`func (o *CouponListing) GetSmartContractAddress() string`

GetSmartContractAddress returns the SmartContractAddress field if non-nil, zero value otherwise.

### GetSmartContractAddressOk

`func (o *CouponListing) GetSmartContractAddressOk() (*string, bool)`

GetSmartContractAddressOk returns a tuple with the SmartContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContractAddress

`func (o *CouponListing) SetSmartContractAddress(v string)`

SetSmartContractAddress sets SmartContractAddress field to given value.

### HasSmartContractAddress

`func (o *CouponListing) HasSmartContractAddress() bool`

HasSmartContractAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *CouponListing) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CouponListing) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CouponListing) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *CouponListing) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetBlockchain

`func (o *CouponListing) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *CouponListing) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *CouponListing) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *CouponListing) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetExpireAt

`func (o *CouponListing) GetExpireAt() string`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *CouponListing) GetExpireAtOk() (*string, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *CouponListing) SetExpireAt(v string)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *CouponListing) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CouponListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CouponListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CouponListing) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CouponListing) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CouponListing) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CouponListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CouponListing) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CouponListing) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CouponListing) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CouponListing) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetProductsCount

`func (o *CouponListing) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *CouponListing) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *CouponListing) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *CouponListing) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


