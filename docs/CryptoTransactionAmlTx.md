# CryptoTransactionAmlTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | AML transaction ID. | [optional] 
**ShopId** | Pointer to **int32** | Shop ID linked to this AML transaction. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice this AML transaction is linked to. | [optional] 
**Origin** | Pointer to **string** | Origin of the AML transaction. | [optional] 
**Type** | Pointer to **string** | Type of the AML transaction. | [optional] 
**Asset** | Pointer to **string** | Asset of the AML transaction. | [optional] 
**Blockchain** | Pointer to **string** | Blockchain of the AML transaction. | [optional] 
**Hash** | Pointer to **string** | AML transaction hash. | [optional] 
**OutputType** | Pointer to **string** | AML transaction output type. | [optional] 
**OutputAddress** | Pointer to **string** | AML transaction output address. | [optional] 
**RiskScore** | Pointer to **string** | AML transaction risk score. | [optional] 
**AssetList** | Pointer to **[]string** | AML transaction asset list. | [optional] 
**Error** | Pointer to **string** | AML transaction error. | [optional] 
**EvaluationDetail** | Pointer to **[]string** | AML transaction evaluation detail. | [optional] 
**Contributions** | Pointer to [**[]CryptoTransactionAmlTxContributionsInner**](CryptoTransactionAmlTxContributionsInner.md) |  | [optional] 
**BlockchainInfo** | Pointer to [**CryptoTransactionAmlTxBlockchainInfo**](CryptoTransactionAmlTxBlockchainInfo.md) |  | [optional] 
**ClusterEntities** | Pointer to [**[]AmlWalletClusterEntitiesInner**](AmlWalletClusterEntitiesInner.md) |  | [optional] 
**RiskScoreDetail** | Pointer to [**AmlWalletRiskScoreDetail**](AmlWalletRiskScoreDetail.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the AML transaction creation date | [optional] 

## Methods

### NewCryptoTransactionAmlTx

`func NewCryptoTransactionAmlTx() *CryptoTransactionAmlTx`

NewCryptoTransactionAmlTx instantiates a new CryptoTransactionAmlTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoTransactionAmlTxWithDefaults

`func NewCryptoTransactionAmlTxWithDefaults() *CryptoTransactionAmlTx`

NewCryptoTransactionAmlTxWithDefaults instantiates a new CryptoTransactionAmlTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CryptoTransactionAmlTx) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CryptoTransactionAmlTx) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CryptoTransactionAmlTx) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CryptoTransactionAmlTx) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *CryptoTransactionAmlTx) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *CryptoTransactionAmlTx) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *CryptoTransactionAmlTx) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *CryptoTransactionAmlTx) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CryptoTransactionAmlTx) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CryptoTransactionAmlTx) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CryptoTransactionAmlTx) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CryptoTransactionAmlTx) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetOrigin

`func (o *CryptoTransactionAmlTx) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CryptoTransactionAmlTx) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CryptoTransactionAmlTx) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CryptoTransactionAmlTx) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetType

`func (o *CryptoTransactionAmlTx) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CryptoTransactionAmlTx) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CryptoTransactionAmlTx) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CryptoTransactionAmlTx) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAsset

`func (o *CryptoTransactionAmlTx) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *CryptoTransactionAmlTx) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *CryptoTransactionAmlTx) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *CryptoTransactionAmlTx) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetBlockchain

`func (o *CryptoTransactionAmlTx) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *CryptoTransactionAmlTx) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *CryptoTransactionAmlTx) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *CryptoTransactionAmlTx) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetHash

`func (o *CryptoTransactionAmlTx) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CryptoTransactionAmlTx) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CryptoTransactionAmlTx) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CryptoTransactionAmlTx) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetOutputType

`func (o *CryptoTransactionAmlTx) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *CryptoTransactionAmlTx) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *CryptoTransactionAmlTx) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *CryptoTransactionAmlTx) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetOutputAddress

`func (o *CryptoTransactionAmlTx) GetOutputAddress() string`

GetOutputAddress returns the OutputAddress field if non-nil, zero value otherwise.

### GetOutputAddressOk

`func (o *CryptoTransactionAmlTx) GetOutputAddressOk() (*string, bool)`

GetOutputAddressOk returns a tuple with the OutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputAddress

`func (o *CryptoTransactionAmlTx) SetOutputAddress(v string)`

SetOutputAddress sets OutputAddress field to given value.

### HasOutputAddress

`func (o *CryptoTransactionAmlTx) HasOutputAddress() bool`

HasOutputAddress returns a boolean if a field has been set.

### GetRiskScore

`func (o *CryptoTransactionAmlTx) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *CryptoTransactionAmlTx) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *CryptoTransactionAmlTx) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *CryptoTransactionAmlTx) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetAssetList

`func (o *CryptoTransactionAmlTx) GetAssetList() []string`

GetAssetList returns the AssetList field if non-nil, zero value otherwise.

### GetAssetListOk

`func (o *CryptoTransactionAmlTx) GetAssetListOk() (*[]string, bool)`

GetAssetListOk returns a tuple with the AssetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetList

`func (o *CryptoTransactionAmlTx) SetAssetList(v []string)`

SetAssetList sets AssetList field to given value.

### HasAssetList

`func (o *CryptoTransactionAmlTx) HasAssetList() bool`

HasAssetList returns a boolean if a field has been set.

### GetError

`func (o *CryptoTransactionAmlTx) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CryptoTransactionAmlTx) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CryptoTransactionAmlTx) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CryptoTransactionAmlTx) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvaluationDetail

`func (o *CryptoTransactionAmlTx) GetEvaluationDetail() []string`

GetEvaluationDetail returns the EvaluationDetail field if non-nil, zero value otherwise.

### GetEvaluationDetailOk

`func (o *CryptoTransactionAmlTx) GetEvaluationDetailOk() (*[]string, bool)`

GetEvaluationDetailOk returns a tuple with the EvaluationDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDetail

`func (o *CryptoTransactionAmlTx) SetEvaluationDetail(v []string)`

SetEvaluationDetail sets EvaluationDetail field to given value.

### HasEvaluationDetail

`func (o *CryptoTransactionAmlTx) HasEvaluationDetail() bool`

HasEvaluationDetail returns a boolean if a field has been set.

### GetContributions

`func (o *CryptoTransactionAmlTx) GetContributions() []CryptoTransactionAmlTxContributionsInner`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *CryptoTransactionAmlTx) GetContributionsOk() (*[]CryptoTransactionAmlTxContributionsInner, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *CryptoTransactionAmlTx) SetContributions(v []CryptoTransactionAmlTxContributionsInner)`

SetContributions sets Contributions field to given value.

### HasContributions

`func (o *CryptoTransactionAmlTx) HasContributions() bool`

HasContributions returns a boolean if a field has been set.

### GetBlockchainInfo

`func (o *CryptoTransactionAmlTx) GetBlockchainInfo() CryptoTransactionAmlTxBlockchainInfo`

GetBlockchainInfo returns the BlockchainInfo field if non-nil, zero value otherwise.

### GetBlockchainInfoOk

`func (o *CryptoTransactionAmlTx) GetBlockchainInfoOk() (*CryptoTransactionAmlTxBlockchainInfo, bool)`

GetBlockchainInfoOk returns a tuple with the BlockchainInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainInfo

`func (o *CryptoTransactionAmlTx) SetBlockchainInfo(v CryptoTransactionAmlTxBlockchainInfo)`

SetBlockchainInfo sets BlockchainInfo field to given value.

### HasBlockchainInfo

`func (o *CryptoTransactionAmlTx) HasBlockchainInfo() bool`

HasBlockchainInfo returns a boolean if a field has been set.

### GetClusterEntities

`func (o *CryptoTransactionAmlTx) GetClusterEntities() []AmlWalletClusterEntitiesInner`

GetClusterEntities returns the ClusterEntities field if non-nil, zero value otherwise.

### GetClusterEntitiesOk

`func (o *CryptoTransactionAmlTx) GetClusterEntitiesOk() (*[]AmlWalletClusterEntitiesInner, bool)`

GetClusterEntitiesOk returns a tuple with the ClusterEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntities

`func (o *CryptoTransactionAmlTx) SetClusterEntities(v []AmlWalletClusterEntitiesInner)`

SetClusterEntities sets ClusterEntities field to given value.

### HasClusterEntities

`func (o *CryptoTransactionAmlTx) HasClusterEntities() bool`

HasClusterEntities returns a boolean if a field has been set.

### GetRiskScoreDetail

`func (o *CryptoTransactionAmlTx) GetRiskScoreDetail() AmlWalletRiskScoreDetail`

GetRiskScoreDetail returns the RiskScoreDetail field if non-nil, zero value otherwise.

### GetRiskScoreDetailOk

`func (o *CryptoTransactionAmlTx) GetRiskScoreDetailOk() (*AmlWalletRiskScoreDetail, bool)`

GetRiskScoreDetailOk returns a tuple with the RiskScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreDetail

`func (o *CryptoTransactionAmlTx) SetRiskScoreDetail(v AmlWalletRiskScoreDetail)`

SetRiskScoreDetail sets RiskScoreDetail field to given value.

### HasRiskScoreDetail

`func (o *CryptoTransactionAmlTx) HasRiskScoreDetail() bool`

HasRiskScoreDetail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CryptoTransactionAmlTx) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CryptoTransactionAmlTx) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CryptoTransactionAmlTx) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CryptoTransactionAmlTx) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


