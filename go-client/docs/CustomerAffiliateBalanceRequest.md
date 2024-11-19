# CustomerAffiliateBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | How the amount should be applied to the customer | [optional] 
**Amount** | Pointer to **float32** | The amount the action should apply to the customer | [optional] 

## Methods

### NewCustomerAffiliateBalanceRequest

`func NewCustomerAffiliateBalanceRequest() *CustomerAffiliateBalanceRequest`

NewCustomerAffiliateBalanceRequest instantiates a new CustomerAffiliateBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAffiliateBalanceRequestWithDefaults

`func NewCustomerAffiliateBalanceRequestWithDefaults() *CustomerAffiliateBalanceRequest`

NewCustomerAffiliateBalanceRequestWithDefaults instantiates a new CustomerAffiliateBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CustomerAffiliateBalanceRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CustomerAffiliateBalanceRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CustomerAffiliateBalanceRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CustomerAffiliateBalanceRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAmount

`func (o *CustomerAffiliateBalanceRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CustomerAffiliateBalanceRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CustomerAffiliateBalanceRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CustomerAffiliateBalanceRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


