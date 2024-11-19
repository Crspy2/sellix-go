# InvoiceRewardsDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clause** | Pointer to **string** |  | [optional] 
**ClauseValue** | Pointer to **string** |  | [optional] 
**ClauseCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**ActionValue** | Pointer to **string** |  | [optional] 
**ActionCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ActionSendProductVariant** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | Creation date of the invoice. | [optional] 
**UpdatedAt** | Pointer to **int32** | Timestamp for when resource was last updated. | [optional] 
**Log** | Pointer to **map[string]interface{}** |  | [optional] 
**RewardInfo** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInvoiceRewardsDataInner

`func NewInvoiceRewardsDataInner() *InvoiceRewardsDataInner`

NewInvoiceRewardsDataInner instantiates a new InvoiceRewardsDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRewardsDataInnerWithDefaults

`func NewInvoiceRewardsDataInnerWithDefaults() *InvoiceRewardsDataInner`

NewInvoiceRewardsDataInnerWithDefaults instantiates a new InvoiceRewardsDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClause

`func (o *InvoiceRewardsDataInner) GetClause() string`

GetClause returns the Clause field if non-nil, zero value otherwise.

### GetClauseOk

`func (o *InvoiceRewardsDataInner) GetClauseOk() (*string, bool)`

GetClauseOk returns a tuple with the Clause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClause

`func (o *InvoiceRewardsDataInner) SetClause(v string)`

SetClause sets Clause field to given value.

### HasClause

`func (o *InvoiceRewardsDataInner) HasClause() bool`

HasClause returns a boolean if a field has been set.

### GetClauseValue

`func (o *InvoiceRewardsDataInner) GetClauseValue() string`

GetClauseValue returns the ClauseValue field if non-nil, zero value otherwise.

### GetClauseValueOk

`func (o *InvoiceRewardsDataInner) GetClauseValueOk() (*string, bool)`

GetClauseValueOk returns a tuple with the ClauseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseValue

`func (o *InvoiceRewardsDataInner) SetClauseValue(v string)`

SetClauseValue sets ClauseValue field to given value.

### HasClauseValue

`func (o *InvoiceRewardsDataInner) HasClauseValue() bool`

HasClauseValue returns a boolean if a field has been set.

### GetClauseCurrency

`func (o *InvoiceRewardsDataInner) GetClauseCurrency() Currency`

GetClauseCurrency returns the ClauseCurrency field if non-nil, zero value otherwise.

### GetClauseCurrencyOk

`func (o *InvoiceRewardsDataInner) GetClauseCurrencyOk() (*Currency, bool)`

GetClauseCurrencyOk returns a tuple with the ClauseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseCurrency

`func (o *InvoiceRewardsDataInner) SetClauseCurrency(v Currency)`

SetClauseCurrency sets ClauseCurrency field to given value.

### HasClauseCurrency

`func (o *InvoiceRewardsDataInner) HasClauseCurrency() bool`

HasClauseCurrency returns a boolean if a field has been set.

### GetAction

`func (o *InvoiceRewardsDataInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InvoiceRewardsDataInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InvoiceRewardsDataInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InvoiceRewardsDataInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionValue

`func (o *InvoiceRewardsDataInner) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *InvoiceRewardsDataInner) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *InvoiceRewardsDataInner) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *InvoiceRewardsDataInner) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionCurrency

`func (o *InvoiceRewardsDataInner) GetActionCurrency() Currency`

GetActionCurrency returns the ActionCurrency field if non-nil, zero value otherwise.

### GetActionCurrencyOk

`func (o *InvoiceRewardsDataInner) GetActionCurrencyOk() (*Currency, bool)`

GetActionCurrencyOk returns a tuple with the ActionCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCurrency

`func (o *InvoiceRewardsDataInner) SetActionCurrency(v Currency)`

SetActionCurrency sets ActionCurrency field to given value.

### HasActionCurrency

`func (o *InvoiceRewardsDataInner) HasActionCurrency() bool`

HasActionCurrency returns a boolean if a field has been set.

### GetActionSendProductVariant

`func (o *InvoiceRewardsDataInner) GetActionSendProductVariant() string`

GetActionSendProductVariant returns the ActionSendProductVariant field if non-nil, zero value otherwise.

### GetActionSendProductVariantOk

`func (o *InvoiceRewardsDataInner) GetActionSendProductVariantOk() (*string, bool)`

GetActionSendProductVariantOk returns a tuple with the ActionSendProductVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSendProductVariant

`func (o *InvoiceRewardsDataInner) SetActionSendProductVariant(v string)`

SetActionSendProductVariant sets ActionSendProductVariant field to given value.

### HasActionSendProductVariant

`func (o *InvoiceRewardsDataInner) HasActionSendProductVariant() bool`

HasActionSendProductVariant returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceRewardsDataInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceRewardsDataInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceRewardsDataInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceRewardsDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceRewardsDataInner) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceRewardsDataInner) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceRewardsDataInner) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceRewardsDataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLog

`func (o *InvoiceRewardsDataInner) GetLog() map[string]interface{}`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *InvoiceRewardsDataInner) GetLogOk() (*map[string]interface{}, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *InvoiceRewardsDataInner) SetLog(v map[string]interface{})`

SetLog sets Log field to given value.

### HasLog

`func (o *InvoiceRewardsDataInner) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetRewardInfo

`func (o *InvoiceRewardsDataInner) GetRewardInfo() map[string]interface{}`

GetRewardInfo returns the RewardInfo field if non-nil, zero value otherwise.

### GetRewardInfoOk

`func (o *InvoiceRewardsDataInner) GetRewardInfoOk() (*map[string]interface{}, bool)`

GetRewardInfoOk returns a tuple with the RewardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardInfo

`func (o *InvoiceRewardsDataInner) SetRewardInfo(v map[string]interface{})`

SetRewardInfo sets RewardInfo field to given value.

### HasRewardInfo

`func (o *InvoiceRewardsDataInner) HasRewardInfo() bool`

HasRewardInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


