# CreateCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**DiscountValue** | **int32** |  | 
**MaxUses** | Pointer to **int32** |  | [optional] 
**DiscountType** | Pointer to **string** |  | [optional] 
**DisabledWithVolumeDiscounts** | Pointer to **bool** |  | [optional] 
**AllRecurringBillInvoices** | Pointer to **bool** |  | [optional] 
**ExpireAt** | Pointer to **float32** |  | [optional] 
**ProductsBound** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCouponRequest

`func NewCreateCouponRequest(code string, discountValue int32, ) *CreateCouponRequest`

NewCreateCouponRequest instantiates a new CreateCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCouponRequestWithDefaults

`func NewCreateCouponRequestWithDefaults() *CreateCouponRequest`

NewCreateCouponRequestWithDefaults instantiates a new CreateCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateCouponRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateCouponRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateCouponRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetDiscountValue

`func (o *CreateCouponRequest) GetDiscountValue() int32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *CreateCouponRequest) GetDiscountValueOk() (*int32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *CreateCouponRequest) SetDiscountValue(v int32)`

SetDiscountValue sets DiscountValue field to given value.


### GetMaxUses

`func (o *CreateCouponRequest) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *CreateCouponRequest) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *CreateCouponRequest) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *CreateCouponRequest) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetDiscountType

`func (o *CreateCouponRequest) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *CreateCouponRequest) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *CreateCouponRequest) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *CreateCouponRequest) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDisabledWithVolumeDiscounts

`func (o *CreateCouponRequest) GetDisabledWithVolumeDiscounts() bool`

GetDisabledWithVolumeDiscounts returns the DisabledWithVolumeDiscounts field if non-nil, zero value otherwise.

### GetDisabledWithVolumeDiscountsOk

`func (o *CreateCouponRequest) GetDisabledWithVolumeDiscountsOk() (*bool, bool)`

GetDisabledWithVolumeDiscountsOk returns a tuple with the DisabledWithVolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWithVolumeDiscounts

`func (o *CreateCouponRequest) SetDisabledWithVolumeDiscounts(v bool)`

SetDisabledWithVolumeDiscounts sets DisabledWithVolumeDiscounts field to given value.

### HasDisabledWithVolumeDiscounts

`func (o *CreateCouponRequest) HasDisabledWithVolumeDiscounts() bool`

HasDisabledWithVolumeDiscounts returns a boolean if a field has been set.

### GetAllRecurringBillInvoices

`func (o *CreateCouponRequest) GetAllRecurringBillInvoices() bool`

GetAllRecurringBillInvoices returns the AllRecurringBillInvoices field if non-nil, zero value otherwise.

### GetAllRecurringBillInvoicesOk

`func (o *CreateCouponRequest) GetAllRecurringBillInvoicesOk() (*bool, bool)`

GetAllRecurringBillInvoicesOk returns a tuple with the AllRecurringBillInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecurringBillInvoices

`func (o *CreateCouponRequest) SetAllRecurringBillInvoices(v bool)`

SetAllRecurringBillInvoices sets AllRecurringBillInvoices field to given value.

### HasAllRecurringBillInvoices

`func (o *CreateCouponRequest) HasAllRecurringBillInvoices() bool`

HasAllRecurringBillInvoices returns a boolean if a field has been set.

### GetExpireAt

`func (o *CreateCouponRequest) GetExpireAt() float32`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *CreateCouponRequest) GetExpireAtOk() (*float32, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *CreateCouponRequest) SetExpireAt(v float32)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *CreateCouponRequest) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetProductsBound

`func (o *CreateCouponRequest) GetProductsBound() []string`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *CreateCouponRequest) GetProductsBoundOk() (*[]string, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *CreateCouponRequest) SetProductsBound(v []string)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *CreateCouponRequest) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


