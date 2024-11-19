# InvoiceListingCryptoScheduledPayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | Cryptocurrency amount sent in the payout transaction. | [optional] 
**TransactionOutput** | Pointer to **float64** | Total output of the transaction, equals to amount-transaction_fee-service_fee. | [optional] 
**ShopCut** | Pointer to **float64** | Equal to transaction_output. | [optional] 
**TransactionFee** | Pointer to **float64** | Crypto transaction fee. | [optional] 
**ServiceFee** | Pointer to **float64** | Service fee, Sellix cut. | [optional] 
**Tx** | Pointer to **string** | Crypto transaction hash. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for when the transaction was broadcast. | [optional] 

## Methods

### NewInvoiceListingCryptoScheduledPayout

`func NewInvoiceListingCryptoScheduledPayout() *InvoiceListingCryptoScheduledPayout`

NewInvoiceListingCryptoScheduledPayout instantiates a new InvoiceListingCryptoScheduledPayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListingCryptoScheduledPayoutWithDefaults

`func NewInvoiceListingCryptoScheduledPayoutWithDefaults() *InvoiceListingCryptoScheduledPayout`

NewInvoiceListingCryptoScheduledPayoutWithDefaults instantiates a new InvoiceListingCryptoScheduledPayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceListingCryptoScheduledPayout) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceListingCryptoScheduledPayout) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceListingCryptoScheduledPayout) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceListingCryptoScheduledPayout) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionOutput

`func (o *InvoiceListingCryptoScheduledPayout) GetTransactionOutput() float64`

GetTransactionOutput returns the TransactionOutput field if non-nil, zero value otherwise.

### GetTransactionOutputOk

`func (o *InvoiceListingCryptoScheduledPayout) GetTransactionOutputOk() (*float64, bool)`

GetTransactionOutputOk returns a tuple with the TransactionOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOutput

`func (o *InvoiceListingCryptoScheduledPayout) SetTransactionOutput(v float64)`

SetTransactionOutput sets TransactionOutput field to given value.

### HasTransactionOutput

`func (o *InvoiceListingCryptoScheduledPayout) HasTransactionOutput() bool`

HasTransactionOutput returns a boolean if a field has been set.

### GetShopCut

`func (o *InvoiceListingCryptoScheduledPayout) GetShopCut() float64`

GetShopCut returns the ShopCut field if non-nil, zero value otherwise.

### GetShopCutOk

`func (o *InvoiceListingCryptoScheduledPayout) GetShopCutOk() (*float64, bool)`

GetShopCutOk returns a tuple with the ShopCut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopCut

`func (o *InvoiceListingCryptoScheduledPayout) SetShopCut(v float64)`

SetShopCut sets ShopCut field to given value.

### HasShopCut

`func (o *InvoiceListingCryptoScheduledPayout) HasShopCut() bool`

HasShopCut returns a boolean if a field has been set.

### GetTransactionFee

`func (o *InvoiceListingCryptoScheduledPayout) GetTransactionFee() float64`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *InvoiceListingCryptoScheduledPayout) GetTransactionFeeOk() (*float64, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *InvoiceListingCryptoScheduledPayout) SetTransactionFee(v float64)`

SetTransactionFee sets TransactionFee field to given value.

### HasTransactionFee

`func (o *InvoiceListingCryptoScheduledPayout) HasTransactionFee() bool`

HasTransactionFee returns a boolean if a field has been set.

### GetServiceFee

`func (o *InvoiceListingCryptoScheduledPayout) GetServiceFee() float64`

GetServiceFee returns the ServiceFee field if non-nil, zero value otherwise.

### GetServiceFeeOk

`func (o *InvoiceListingCryptoScheduledPayout) GetServiceFeeOk() (*float64, bool)`

GetServiceFeeOk returns a tuple with the ServiceFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFee

`func (o *InvoiceListingCryptoScheduledPayout) SetServiceFee(v float64)`

SetServiceFee sets ServiceFee field to given value.

### HasServiceFee

`func (o *InvoiceListingCryptoScheduledPayout) HasServiceFee() bool`

HasServiceFee returns a boolean if a field has been set.

### GetTx

`func (o *InvoiceListingCryptoScheduledPayout) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *InvoiceListingCryptoScheduledPayout) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *InvoiceListingCryptoScheduledPayout) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *InvoiceListingCryptoScheduledPayout) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceListingCryptoScheduledPayout) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceListingCryptoScheduledPayout) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceListingCryptoScheduledPayout) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceListingCryptoScheduledPayout) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


