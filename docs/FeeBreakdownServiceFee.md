# FeeBreakdownServiceFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The amount of the service fee. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Breakdown** | Pointer to [**FeeBreakdownServiceFeeBreakdown**](FeeBreakdownServiceFeeBreakdown.md) |  | [optional] 

## Methods

### NewFeeBreakdownServiceFee

`func NewFeeBreakdownServiceFee() *FeeBreakdownServiceFee`

NewFeeBreakdownServiceFee instantiates a new FeeBreakdownServiceFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeBreakdownServiceFeeWithDefaults

`func NewFeeBreakdownServiceFeeWithDefaults() *FeeBreakdownServiceFee`

NewFeeBreakdownServiceFeeWithDefaults instantiates a new FeeBreakdownServiceFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeBreakdownServiceFee) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeBreakdownServiceFee) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeBreakdownServiceFee) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeBreakdownServiceFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeBreakdownServiceFee) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeBreakdownServiceFee) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeBreakdownServiceFee) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeeBreakdownServiceFee) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBreakdown

`func (o *FeeBreakdownServiceFee) GetBreakdown() FeeBreakdownServiceFeeBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FeeBreakdownServiceFee) GetBreakdownOk() (*FeeBreakdownServiceFeeBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FeeBreakdownServiceFee) SetBreakdown(v FeeBreakdownServiceFeeBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FeeBreakdownServiceFee) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


