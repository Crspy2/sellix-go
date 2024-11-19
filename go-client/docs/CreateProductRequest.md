# CreateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Price** | **float64** |  | 
**Description** | **string** |  | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Product type. | 
**Serials** | Pointer to **[]string** |  | [optional] 
**SerialsRemoveDuplicates** | Pointer to **bool** |  | [optional] 
**ServiceText** | Pointer to **string** |  | [optional] 
**Stock** | Pointer to **int32** |  | [optional] 
**DynamicWebhook** | Pointer to **string** |  | [optional] 
**StockDelimiter** | Pointer to **string** |  | [optional] 
**MinQuantity** | Pointer to **int32** |  | [optional] 
**MaxQuantity** | Pointer to **int32** |  | [optional] 
**DeliveryText** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **[]map[string]interface{}** | key-value pairs for custom metadata to be included in the invoice throughout its lifetime | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** |  | [optional] 
**MaxRiskLevel** | Pointer to **int32** | For PAYPAL and STRIPE, maximum risk level to accept payments in order to block fraud attempts. | [optional] 
**Unlisted** | Pointer to **bool** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**BlockVpnProxies** | Pointer to **bool** | Block VPN and Proxy purchases if the gateway is PAYPAL or STRIPE. | [optional] 
**SortPriority** | Pointer to **int32** |  | [optional] 
**Webhooks** | Pointer to **[]string** |  | [optional] 
**OnHold** | Pointer to **bool** |  | [optional] 
**TermsOfService** | Pointer to **string** |  | [optional] 
**Warranty** | Pointer to **int32** | The length of the warranty in seconds | [optional] 
**WarrantyText** | Pointer to **string** |  | [optional] 
**RemoveImage** | Pointer to **bool** |  | [optional] 
**RemoveFile** | Pointer to **bool** |  | [optional] 
**VolumeDiscounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RecurringInterval** | Pointer to **string** |  | [optional] 
**RecurringIntervalCount** | Pointer to **int32** |  | [optional] 
**TrialPeriod** | Pointer to **int32** |  | [optional] 
**ProductPlans** | Pointer to [**[]ProductPlanRequest**](ProductPlanRequest.md) |  | [optional] 

## Methods

### NewCreateProductRequest

`func NewCreateProductRequest(title string, price float64, description string, type_ string, ) *CreateProductRequest`

NewCreateProductRequest instantiates a new CreateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductRequestWithDefaults

`func NewCreateProductRequestWithDefaults() *CreateProductRequest`

NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateProductRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateProductRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateProductRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrice

`func (o *CreateProductRequest) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductRequest) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductRequest) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetDescription

`func (o *CreateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *CreateProductRequest) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateProductRequest) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateProductRequest) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGateways

`func (o *CreateProductRequest) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CreateProductRequest) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CreateProductRequest) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CreateProductRequest) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetType

`func (o *CreateProductRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateProductRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateProductRequest) SetType(v string)`

SetType sets Type field to given value.


### GetSerials

`func (o *CreateProductRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *CreateProductRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *CreateProductRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *CreateProductRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSerialsRemoveDuplicates

`func (o *CreateProductRequest) GetSerialsRemoveDuplicates() bool`

GetSerialsRemoveDuplicates returns the SerialsRemoveDuplicates field if non-nil, zero value otherwise.

### GetSerialsRemoveDuplicatesOk

`func (o *CreateProductRequest) GetSerialsRemoveDuplicatesOk() (*bool, bool)`

GetSerialsRemoveDuplicatesOk returns a tuple with the SerialsRemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialsRemoveDuplicates

`func (o *CreateProductRequest) SetSerialsRemoveDuplicates(v bool)`

SetSerialsRemoveDuplicates sets SerialsRemoveDuplicates field to given value.

### HasSerialsRemoveDuplicates

`func (o *CreateProductRequest) HasSerialsRemoveDuplicates() bool`

HasSerialsRemoveDuplicates returns a boolean if a field has been set.

### GetServiceText

`func (o *CreateProductRequest) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *CreateProductRequest) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *CreateProductRequest) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *CreateProductRequest) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStock

`func (o *CreateProductRequest) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductRequest) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductRequest) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CreateProductRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *CreateProductRequest) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *CreateProductRequest) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *CreateProductRequest) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *CreateProductRequest) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *CreateProductRequest) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *CreateProductRequest) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *CreateProductRequest) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *CreateProductRequest) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetMinQuantity

`func (o *CreateProductRequest) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *CreateProductRequest) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *CreateProductRequest) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *CreateProductRequest) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *CreateProductRequest) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *CreateProductRequest) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *CreateProductRequest) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *CreateProductRequest) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetDeliveryText

`func (o *CreateProductRequest) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *CreateProductRequest) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *CreateProductRequest) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *CreateProductRequest) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateProductRequest) GetCustomFields() []map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateProductRequest) GetCustomFieldsOk() (*[]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateProductRequest) SetCustomFields(v []map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateProductRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *CreateProductRequest) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *CreateProductRequest) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *CreateProductRequest) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *CreateProductRequest) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *CreateProductRequest) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *CreateProductRequest) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *CreateProductRequest) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *CreateProductRequest) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetUnlisted

`func (o *CreateProductRequest) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *CreateProductRequest) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *CreateProductRequest) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *CreateProductRequest) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetPrivate

`func (o *CreateProductRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateProductRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateProductRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateProductRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *CreateProductRequest) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *CreateProductRequest) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *CreateProductRequest) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *CreateProductRequest) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetSortPriority

`func (o *CreateProductRequest) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *CreateProductRequest) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *CreateProductRequest) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *CreateProductRequest) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetWebhooks

`func (o *CreateProductRequest) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *CreateProductRequest) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *CreateProductRequest) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *CreateProductRequest) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetOnHold

`func (o *CreateProductRequest) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *CreateProductRequest) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *CreateProductRequest) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *CreateProductRequest) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *CreateProductRequest) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *CreateProductRequest) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *CreateProductRequest) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *CreateProductRequest) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *CreateProductRequest) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *CreateProductRequest) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *CreateProductRequest) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *CreateProductRequest) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *CreateProductRequest) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *CreateProductRequest) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *CreateProductRequest) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *CreateProductRequest) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetRemoveImage

`func (o *CreateProductRequest) GetRemoveImage() bool`

GetRemoveImage returns the RemoveImage field if non-nil, zero value otherwise.

### GetRemoveImageOk

`func (o *CreateProductRequest) GetRemoveImageOk() (*bool, bool)`

GetRemoveImageOk returns a tuple with the RemoveImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveImage

`func (o *CreateProductRequest) SetRemoveImage(v bool)`

SetRemoveImage sets RemoveImage field to given value.

### HasRemoveImage

`func (o *CreateProductRequest) HasRemoveImage() bool`

HasRemoveImage returns a boolean if a field has been set.

### GetRemoveFile

`func (o *CreateProductRequest) GetRemoveFile() bool`

GetRemoveFile returns the RemoveFile field if non-nil, zero value otherwise.

### GetRemoveFileOk

`func (o *CreateProductRequest) GetRemoveFileOk() (*bool, bool)`

GetRemoveFileOk returns a tuple with the RemoveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFile

`func (o *CreateProductRequest) SetRemoveFile(v bool)`

SetRemoveFile sets RemoveFile field to given value.

### HasRemoveFile

`func (o *CreateProductRequest) HasRemoveFile() bool`

HasRemoveFile returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *CreateProductRequest) GetVolumeDiscounts() []map[string]interface{}`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *CreateProductRequest) GetVolumeDiscountsOk() (*[]map[string]interface{}, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *CreateProductRequest) SetVolumeDiscounts(v []map[string]interface{})`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *CreateProductRequest) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *CreateProductRequest) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *CreateProductRequest) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *CreateProductRequest) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *CreateProductRequest) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *CreateProductRequest) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *CreateProductRequest) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *CreateProductRequest) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *CreateProductRequest) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *CreateProductRequest) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *CreateProductRequest) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *CreateProductRequest) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *CreateProductRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetProductPlans

`func (o *CreateProductRequest) GetProductPlans() []ProductPlanRequest`

GetProductPlans returns the ProductPlans field if non-nil, zero value otherwise.

### GetProductPlansOk

`func (o *CreateProductRequest) GetProductPlansOk() (*[]ProductPlanRequest, bool)`

GetProductPlansOk returns a tuple with the ProductPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlans

`func (o *CreateProductRequest) SetProductPlans(v []ProductPlanRequest)`

SetProductPlans sets ProductPlans field to given value.

### HasProductPlans

`func (o *CreateProductRequest) HasProductPlans() bool`

HasProductPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


