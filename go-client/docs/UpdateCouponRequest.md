# UpdateCouponRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**DiscountValue** | Pointer to **int32** |  | [optional] 
**MaxUses** | Pointer to **int32** |  | [optional] 
**DiscountType** | Pointer to **string** |  | [optional] 
**DisabledWithVolumeDiscounts** | Pointer to **bool** |  | [optional] 
**AllRecurringBillInvoices** | Pointer to **bool** |  | [optional] 
**ProductsBound** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateCouponRequest

`func NewUpdateCouponRequest() *UpdateCouponRequest`

NewUpdateCouponRequest instantiates a new UpdateCouponRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCouponRequestWithDefaults

`func NewUpdateCouponRequestWithDefaults() *UpdateCouponRequest`

NewUpdateCouponRequestWithDefaults instantiates a new UpdateCouponRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UpdateCouponRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCouponRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCouponRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCouponRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDiscountValue

`func (o *UpdateCouponRequest) GetDiscountValue() int32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *UpdateCouponRequest) GetDiscountValueOk() (*int32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *UpdateCouponRequest) SetDiscountValue(v int32)`

SetDiscountValue sets DiscountValue field to given value.

### HasDiscountValue

`func (o *UpdateCouponRequest) HasDiscountValue() bool`

HasDiscountValue returns a boolean if a field has been set.

### GetMaxUses

`func (o *UpdateCouponRequest) GetMaxUses() int32`

GetMaxUses returns the MaxUses field if non-nil, zero value otherwise.

### GetMaxUsesOk

`func (o *UpdateCouponRequest) GetMaxUsesOk() (*int32, bool)`

GetMaxUsesOk returns a tuple with the MaxUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUses

`func (o *UpdateCouponRequest) SetMaxUses(v int32)`

SetMaxUses sets MaxUses field to given value.

### HasMaxUses

`func (o *UpdateCouponRequest) HasMaxUses() bool`

HasMaxUses returns a boolean if a field has been set.

### GetDiscountType

`func (o *UpdateCouponRequest) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UpdateCouponRequest) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UpdateCouponRequest) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UpdateCouponRequest) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDisabledWithVolumeDiscounts

`func (o *UpdateCouponRequest) GetDisabledWithVolumeDiscounts() bool`

GetDisabledWithVolumeDiscounts returns the DisabledWithVolumeDiscounts field if non-nil, zero value otherwise.

### GetDisabledWithVolumeDiscountsOk

`func (o *UpdateCouponRequest) GetDisabledWithVolumeDiscountsOk() (*bool, bool)`

GetDisabledWithVolumeDiscountsOk returns a tuple with the DisabledWithVolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledWithVolumeDiscounts

`func (o *UpdateCouponRequest) SetDisabledWithVolumeDiscounts(v bool)`

SetDisabledWithVolumeDiscounts sets DisabledWithVolumeDiscounts field to given value.

### HasDisabledWithVolumeDiscounts

`func (o *UpdateCouponRequest) HasDisabledWithVolumeDiscounts() bool`

HasDisabledWithVolumeDiscounts returns a boolean if a field has been set.

### GetAllRecurringBillInvoices

`func (o *UpdateCouponRequest) GetAllRecurringBillInvoices() bool`

GetAllRecurringBillInvoices returns the AllRecurringBillInvoices field if non-nil, zero value otherwise.

### GetAllRecurringBillInvoicesOk

`func (o *UpdateCouponRequest) GetAllRecurringBillInvoicesOk() (*bool, bool)`

GetAllRecurringBillInvoicesOk returns a tuple with the AllRecurringBillInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRecurringBillInvoices

`func (o *UpdateCouponRequest) SetAllRecurringBillInvoices(v bool)`

SetAllRecurringBillInvoices sets AllRecurringBillInvoices field to given value.

### HasAllRecurringBillInvoices

`func (o *UpdateCouponRequest) HasAllRecurringBillInvoices() bool`

HasAllRecurringBillInvoices returns a boolean if a field has been set.

### GetProductsBound

`func (o *UpdateCouponRequest) GetProductsBound() []string`

GetProductsBound returns the ProductsBound field if non-nil, zero value otherwise.

### GetProductsBoundOk

`func (o *UpdateCouponRequest) GetProductsBoundOk() (*[]string, bool)`

GetProductsBoundOk returns a tuple with the ProductsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsBound

`func (o *UpdateCouponRequest) SetProductsBound(v []string)`

SetProductsBound sets ProductsBound field to given value.

### HasProductsBound

`func (o *UpdateCouponRequest) HasProductsBound() bool`

HasProductsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


