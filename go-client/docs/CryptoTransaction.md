# CryptoTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoAmount** | Pointer to **float64** | Crypto amount sent in the transaction. | [optional] 
**Hash** | Pointer to **string** | Crypto transaction hash. | [optional] 
**Confirmations** | Pointer to **int32** | Crypto transaction confirmations, not updated once the invoice status is COMPLETED or VOIDED. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the crypto transaction reception date | [optional] 
**UpdatedAt** | Pointer to **int32** | Crypto transaction update date. | [optional] 
**AmlTx** | Pointer to [**CryptoTransactionAmlTx**](CryptoTransactionAmlTx.md) |  | [optional] 

## Methods

### NewCryptoTransaction

`func NewCryptoTransaction() *CryptoTransaction`

NewCryptoTransaction instantiates a new CryptoTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoTransactionWithDefaults

`func NewCryptoTransactionWithDefaults() *CryptoTransaction`

NewCryptoTransactionWithDefaults instantiates a new CryptoTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCryptoAmount

`func (o *CryptoTransaction) GetCryptoAmount() float64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *CryptoTransaction) GetCryptoAmountOk() (*float64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *CryptoTransaction) SetCryptoAmount(v float64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *CryptoTransaction) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetHash

`func (o *CryptoTransaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CryptoTransaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CryptoTransaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CryptoTransaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetConfirmations

`func (o *CryptoTransaction) GetConfirmations() int32`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *CryptoTransaction) GetConfirmationsOk() (*int32, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *CryptoTransaction) SetConfirmations(v int32)`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *CryptoTransaction) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CryptoTransaction) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CryptoTransaction) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CryptoTransaction) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CryptoTransaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CryptoTransaction) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CryptoTransaction) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CryptoTransaction) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CryptoTransaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAmlTx

`func (o *CryptoTransaction) GetAmlTx() CryptoTransactionAmlTx`

GetAmlTx returns the AmlTx field if non-nil, zero value otherwise.

### GetAmlTxOk

`func (o *CryptoTransaction) GetAmlTxOk() (*CryptoTransactionAmlTx, bool)`

GetAmlTxOk returns a tuple with the AmlTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmlTx

`func (o *CryptoTransaction) SetAmlTx(v CryptoTransactionAmlTx)`

SetAmlTx sets AmlTx field to given value.

### HasAmlTx

`func (o *CryptoTransaction) HasAmlTx() bool`

HasAmlTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


