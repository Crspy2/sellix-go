# FeeBreakdownPlatformFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The amount of the service fee. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Breakdown** | Pointer to [**FeeBreakdownPlatformFeeBreakdown**](FeeBreakdownPlatformFeeBreakdown.md) |  | [optional] 

## Methods

### NewFeeBreakdownPlatformFee

`func NewFeeBreakdownPlatformFee() *FeeBreakdownPlatformFee`

NewFeeBreakdownPlatformFee instantiates a new FeeBreakdownPlatformFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeBreakdownPlatformFeeWithDefaults

`func NewFeeBreakdownPlatformFeeWithDefaults() *FeeBreakdownPlatformFee`

NewFeeBreakdownPlatformFeeWithDefaults instantiates a new FeeBreakdownPlatformFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeBreakdownPlatformFee) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeBreakdownPlatformFee) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeBreakdownPlatformFee) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeBreakdownPlatformFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeBreakdownPlatformFee) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeBreakdownPlatformFee) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeBreakdownPlatformFee) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeeBreakdownPlatformFee) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBreakdown

`func (o *FeeBreakdownPlatformFee) GetBreakdown() FeeBreakdownPlatformFeeBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FeeBreakdownPlatformFee) GetBreakdownOk() (*FeeBreakdownPlatformFeeBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FeeBreakdownPlatformFee) SetBreakdown(v FeeBreakdownPlatformFeeBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FeeBreakdownPlatformFee) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


