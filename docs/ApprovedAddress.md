# ApprovedAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Address** | Pointer to **string** | The crypto address that was approved. | [optional] 
**Coin** | Pointer to **string** | The coin of the approved address. | [optional] 
**Blockchain** | Pointer to [**Blockchain**](Blockchain.md) |  | [optional] 
**Tx** | Pointer to **string** | The transaction ID of the approval. | [optional] 
**RecurringBillingId** | Pointer to **string** | The recurring billing ID of the approval. | [optional] 
**Allowance** | Pointer to **float32** | The allowance of the approved address. | [optional] 
**UpdatedAt** | Pointer to **int32** | Timestamp, available if the approved address has been updated. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp, available if the approved address has been created. | [optional] 

## Methods

### NewApprovedAddress

`func NewApprovedAddress() *ApprovedAddress`

NewApprovedAddress instantiates a new ApprovedAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovedAddressWithDefaults

`func NewApprovedAddressWithDefaults() *ApprovedAddress`

NewApprovedAddressWithDefaults instantiates a new ApprovedAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovedAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovedAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovedAddress) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovedAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *ApprovedAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApprovedAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApprovedAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApprovedAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCoin

`func (o *ApprovedAddress) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ApprovedAddress) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ApprovedAddress) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ApprovedAddress) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetBlockchain

`func (o *ApprovedAddress) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ApprovedAddress) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ApprovedAddress) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *ApprovedAddress) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetTx

`func (o *ApprovedAddress) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *ApprovedAddress) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *ApprovedAddress) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *ApprovedAddress) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetRecurringBillingId

`func (o *ApprovedAddress) GetRecurringBillingId() string`

GetRecurringBillingId returns the RecurringBillingId field if non-nil, zero value otherwise.

### GetRecurringBillingIdOk

`func (o *ApprovedAddress) GetRecurringBillingIdOk() (*string, bool)`

GetRecurringBillingIdOk returns a tuple with the RecurringBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringBillingId

`func (o *ApprovedAddress) SetRecurringBillingId(v string)`

SetRecurringBillingId sets RecurringBillingId field to given value.

### HasRecurringBillingId

`func (o *ApprovedAddress) HasRecurringBillingId() bool`

HasRecurringBillingId returns a boolean if a field has been set.

### GetAllowance

`func (o *ApprovedAddress) GetAllowance() float32`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *ApprovedAddress) GetAllowanceOk() (*float32, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *ApprovedAddress) SetAllowance(v float32)`

SetAllowance sets Allowance field to given value.

### HasAllowance

`func (o *ApprovedAddress) HasAllowance() bool`

HasAllowance returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApprovedAddress) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApprovedAddress) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApprovedAddress) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApprovedAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApprovedAddress) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApprovedAddress) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApprovedAddress) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApprovedAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


