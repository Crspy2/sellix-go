# AmlWallet

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
**EvaluationDetail** | Pointer to [**[]AmlWalletEvaluationDetailInner**](AmlWalletEvaluationDetailInner.md) |  | [optional] 
**Contributions** | Pointer to [**AmlWalletContributions**](AmlWalletContributions.md) |  | [optional] 
**BlockchainInfo** | Pointer to [**AmlWalletBlockchainInfo**](AmlWalletBlockchainInfo.md) |  | [optional] 
**ClusterEntities** | Pointer to [**[]AmlWalletClusterEntitiesInner**](AmlWalletClusterEntitiesInner.md) |  | [optional] 
**RiskScoreDetail** | Pointer to [**AmlWalletRiskScoreDetail**](AmlWalletRiskScoreDetail.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the AML transaction creation date. | [optional] 

## Methods

### NewAmlWallet

`func NewAmlWallet() *AmlWallet`

NewAmlWallet instantiates a new AmlWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmlWalletWithDefaults

`func NewAmlWalletWithDefaults() *AmlWallet`

NewAmlWalletWithDefaults instantiates a new AmlWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmlWallet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmlWallet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmlWallet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AmlWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *AmlWallet) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *AmlWallet) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *AmlWallet) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *AmlWallet) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *AmlWallet) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AmlWallet) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AmlWallet) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AmlWallet) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetOrigin

`func (o *AmlWallet) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AmlWallet) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AmlWallet) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AmlWallet) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetType

`func (o *AmlWallet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmlWallet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmlWallet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AmlWallet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAsset

`func (o *AmlWallet) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AmlWallet) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AmlWallet) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AmlWallet) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetBlockchain

`func (o *AmlWallet) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *AmlWallet) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *AmlWallet) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *AmlWallet) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetHash

`func (o *AmlWallet) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AmlWallet) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AmlWallet) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AmlWallet) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetOutputType

`func (o *AmlWallet) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *AmlWallet) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *AmlWallet) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *AmlWallet) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetOutputAddress

`func (o *AmlWallet) GetOutputAddress() string`

GetOutputAddress returns the OutputAddress field if non-nil, zero value otherwise.

### GetOutputAddressOk

`func (o *AmlWallet) GetOutputAddressOk() (*string, bool)`

GetOutputAddressOk returns a tuple with the OutputAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputAddress

`func (o *AmlWallet) SetOutputAddress(v string)`

SetOutputAddress sets OutputAddress field to given value.

### HasOutputAddress

`func (o *AmlWallet) HasOutputAddress() bool`

HasOutputAddress returns a boolean if a field has been set.

### GetRiskScore

`func (o *AmlWallet) GetRiskScore() string`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AmlWallet) GetRiskScoreOk() (*string, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AmlWallet) SetRiskScore(v string)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *AmlWallet) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetAssetList

`func (o *AmlWallet) GetAssetList() []string`

GetAssetList returns the AssetList field if non-nil, zero value otherwise.

### GetAssetListOk

`func (o *AmlWallet) GetAssetListOk() (*[]string, bool)`

GetAssetListOk returns a tuple with the AssetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetList

`func (o *AmlWallet) SetAssetList(v []string)`

SetAssetList sets AssetList field to given value.

### HasAssetList

`func (o *AmlWallet) HasAssetList() bool`

HasAssetList returns a boolean if a field has been set.

### GetError

`func (o *AmlWallet) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AmlWallet) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AmlWallet) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AmlWallet) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvaluationDetail

`func (o *AmlWallet) GetEvaluationDetail() []AmlWalletEvaluationDetailInner`

GetEvaluationDetail returns the EvaluationDetail field if non-nil, zero value otherwise.

### GetEvaluationDetailOk

`func (o *AmlWallet) GetEvaluationDetailOk() (*[]AmlWalletEvaluationDetailInner, bool)`

GetEvaluationDetailOk returns a tuple with the EvaluationDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDetail

`func (o *AmlWallet) SetEvaluationDetail(v []AmlWalletEvaluationDetailInner)`

SetEvaluationDetail sets EvaluationDetail field to given value.

### HasEvaluationDetail

`func (o *AmlWallet) HasEvaluationDetail() bool`

HasEvaluationDetail returns a boolean if a field has been set.

### GetContributions

`func (o *AmlWallet) GetContributions() AmlWalletContributions`

GetContributions returns the Contributions field if non-nil, zero value otherwise.

### GetContributionsOk

`func (o *AmlWallet) GetContributionsOk() (*AmlWalletContributions, bool)`

GetContributionsOk returns a tuple with the Contributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributions

`func (o *AmlWallet) SetContributions(v AmlWalletContributions)`

SetContributions sets Contributions field to given value.

### HasContributions

`func (o *AmlWallet) HasContributions() bool`

HasContributions returns a boolean if a field has been set.

### GetBlockchainInfo

`func (o *AmlWallet) GetBlockchainInfo() AmlWalletBlockchainInfo`

GetBlockchainInfo returns the BlockchainInfo field if non-nil, zero value otherwise.

### GetBlockchainInfoOk

`func (o *AmlWallet) GetBlockchainInfoOk() (*AmlWalletBlockchainInfo, bool)`

GetBlockchainInfoOk returns a tuple with the BlockchainInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainInfo

`func (o *AmlWallet) SetBlockchainInfo(v AmlWalletBlockchainInfo)`

SetBlockchainInfo sets BlockchainInfo field to given value.

### HasBlockchainInfo

`func (o *AmlWallet) HasBlockchainInfo() bool`

HasBlockchainInfo returns a boolean if a field has been set.

### GetClusterEntities

`func (o *AmlWallet) GetClusterEntities() []AmlWalletClusterEntitiesInner`

GetClusterEntities returns the ClusterEntities field if non-nil, zero value otherwise.

### GetClusterEntitiesOk

`func (o *AmlWallet) GetClusterEntitiesOk() (*[]AmlWalletClusterEntitiesInner, bool)`

GetClusterEntitiesOk returns a tuple with the ClusterEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntities

`func (o *AmlWallet) SetClusterEntities(v []AmlWalletClusterEntitiesInner)`

SetClusterEntities sets ClusterEntities field to given value.

### HasClusterEntities

`func (o *AmlWallet) HasClusterEntities() bool`

HasClusterEntities returns a boolean if a field has been set.

### GetRiskScoreDetail

`func (o *AmlWallet) GetRiskScoreDetail() AmlWalletRiskScoreDetail`

GetRiskScoreDetail returns the RiskScoreDetail field if non-nil, zero value otherwise.

### GetRiskScoreDetailOk

`func (o *AmlWallet) GetRiskScoreDetailOk() (*AmlWalletRiskScoreDetail, bool)`

GetRiskScoreDetailOk returns a tuple with the RiskScoreDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScoreDetail

`func (o *AmlWallet) SetRiskScoreDetail(v AmlWalletRiskScoreDetail)`

SetRiskScoreDetail sets RiskScoreDetail field to given value.

### HasRiskScoreDetail

`func (o *AmlWallet) HasRiskScoreDetail() bool`

HasRiskScoreDetail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AmlWallet) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmlWallet) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmlWallet) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AmlWallet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


