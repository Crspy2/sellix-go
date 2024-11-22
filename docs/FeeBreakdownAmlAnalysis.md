# FeeBreakdownAmlAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The amount of the service fee. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Breakdown** | Pointer to [**FeeBreakdownAmlAnalysisBreakdown**](FeeBreakdownAmlAnalysisBreakdown.md) |  | [optional] 

## Methods

### NewFeeBreakdownAmlAnalysis

`func NewFeeBreakdownAmlAnalysis() *FeeBreakdownAmlAnalysis`

NewFeeBreakdownAmlAnalysis instantiates a new FeeBreakdownAmlAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeBreakdownAmlAnalysisWithDefaults

`func NewFeeBreakdownAmlAnalysisWithDefaults() *FeeBreakdownAmlAnalysis`

NewFeeBreakdownAmlAnalysisWithDefaults instantiates a new FeeBreakdownAmlAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FeeBreakdownAmlAnalysis) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeBreakdownAmlAnalysis) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeBreakdownAmlAnalysis) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeBreakdownAmlAnalysis) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FeeBreakdownAmlAnalysis) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeeBreakdownAmlAnalysis) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeeBreakdownAmlAnalysis) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeeBreakdownAmlAnalysis) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBreakdown

`func (o *FeeBreakdownAmlAnalysis) GetBreakdown() FeeBreakdownAmlAnalysisBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FeeBreakdownAmlAnalysis) GetBreakdownOk() (*FeeBreakdownAmlAnalysisBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FeeBreakdownAmlAnalysis) SetBreakdown(v FeeBreakdownAmlAnalysisBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FeeBreakdownAmlAnalysis) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


