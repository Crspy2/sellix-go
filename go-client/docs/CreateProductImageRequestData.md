# CreateProductImageRequestData

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
**CustomFields** | Pointer to **[]map[string]interface{}** |  | [optional] 
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

## Methods

### NewCreateProductImageRequestData

`func NewCreateProductImageRequestData(title string, price float64, description string, type_ string, ) *CreateProductImageRequestData`

NewCreateProductImageRequestData instantiates a new CreateProductImageRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductImageRequestDataWithDefaults

`func NewCreateProductImageRequestDataWithDefaults() *CreateProductImageRequestData`

NewCreateProductImageRequestDataWithDefaults instantiates a new CreateProductImageRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateProductImageRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateProductImageRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateProductImageRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrice

`func (o *CreateProductImageRequestData) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateProductImageRequestData) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateProductImageRequestData) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetDescription

`func (o *CreateProductImageRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductImageRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductImageRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *CreateProductImageRequestData) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateProductImageRequestData) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateProductImageRequestData) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateProductImageRequestData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGateways

`func (o *CreateProductImageRequestData) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CreateProductImageRequestData) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CreateProductImageRequestData) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CreateProductImageRequestData) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetType

`func (o *CreateProductImageRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateProductImageRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateProductImageRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetSerials

`func (o *CreateProductImageRequestData) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *CreateProductImageRequestData) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *CreateProductImageRequestData) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *CreateProductImageRequestData) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSerialsRemoveDuplicates

`func (o *CreateProductImageRequestData) GetSerialsRemoveDuplicates() bool`

GetSerialsRemoveDuplicates returns the SerialsRemoveDuplicates field if non-nil, zero value otherwise.

### GetSerialsRemoveDuplicatesOk

`func (o *CreateProductImageRequestData) GetSerialsRemoveDuplicatesOk() (*bool, bool)`

GetSerialsRemoveDuplicatesOk returns a tuple with the SerialsRemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialsRemoveDuplicates

`func (o *CreateProductImageRequestData) SetSerialsRemoveDuplicates(v bool)`

SetSerialsRemoveDuplicates sets SerialsRemoveDuplicates field to given value.

### HasSerialsRemoveDuplicates

`func (o *CreateProductImageRequestData) HasSerialsRemoveDuplicates() bool`

HasSerialsRemoveDuplicates returns a boolean if a field has been set.

### GetServiceText

`func (o *CreateProductImageRequestData) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *CreateProductImageRequestData) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *CreateProductImageRequestData) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *CreateProductImageRequestData) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStock

`func (o *CreateProductImageRequestData) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *CreateProductImageRequestData) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *CreateProductImageRequestData) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *CreateProductImageRequestData) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *CreateProductImageRequestData) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *CreateProductImageRequestData) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *CreateProductImageRequestData) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *CreateProductImageRequestData) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *CreateProductImageRequestData) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *CreateProductImageRequestData) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *CreateProductImageRequestData) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *CreateProductImageRequestData) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetMinQuantity

`func (o *CreateProductImageRequestData) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *CreateProductImageRequestData) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *CreateProductImageRequestData) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *CreateProductImageRequestData) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *CreateProductImageRequestData) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *CreateProductImageRequestData) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *CreateProductImageRequestData) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *CreateProductImageRequestData) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetDeliveryText

`func (o *CreateProductImageRequestData) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *CreateProductImageRequestData) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *CreateProductImageRequestData) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *CreateProductImageRequestData) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateProductImageRequestData) GetCustomFields() []map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateProductImageRequestData) GetCustomFieldsOk() (*[]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateProductImageRequestData) SetCustomFields(v []map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateProductImageRequestData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *CreateProductImageRequestData) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *CreateProductImageRequestData) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *CreateProductImageRequestData) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *CreateProductImageRequestData) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *CreateProductImageRequestData) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *CreateProductImageRequestData) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *CreateProductImageRequestData) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *CreateProductImageRequestData) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetUnlisted

`func (o *CreateProductImageRequestData) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *CreateProductImageRequestData) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *CreateProductImageRequestData) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *CreateProductImageRequestData) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetPrivate

`func (o *CreateProductImageRequestData) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CreateProductImageRequestData) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CreateProductImageRequestData) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *CreateProductImageRequestData) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *CreateProductImageRequestData) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *CreateProductImageRequestData) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *CreateProductImageRequestData) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *CreateProductImageRequestData) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetSortPriority

`func (o *CreateProductImageRequestData) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *CreateProductImageRequestData) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *CreateProductImageRequestData) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *CreateProductImageRequestData) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetWebhooks

`func (o *CreateProductImageRequestData) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *CreateProductImageRequestData) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *CreateProductImageRequestData) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *CreateProductImageRequestData) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetOnHold

`func (o *CreateProductImageRequestData) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *CreateProductImageRequestData) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *CreateProductImageRequestData) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *CreateProductImageRequestData) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *CreateProductImageRequestData) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *CreateProductImageRequestData) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *CreateProductImageRequestData) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *CreateProductImageRequestData) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *CreateProductImageRequestData) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *CreateProductImageRequestData) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *CreateProductImageRequestData) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *CreateProductImageRequestData) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *CreateProductImageRequestData) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *CreateProductImageRequestData) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *CreateProductImageRequestData) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *CreateProductImageRequestData) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetRemoveImage

`func (o *CreateProductImageRequestData) GetRemoveImage() bool`

GetRemoveImage returns the RemoveImage field if non-nil, zero value otherwise.

### GetRemoveImageOk

`func (o *CreateProductImageRequestData) GetRemoveImageOk() (*bool, bool)`

GetRemoveImageOk returns a tuple with the RemoveImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveImage

`func (o *CreateProductImageRequestData) SetRemoveImage(v bool)`

SetRemoveImage sets RemoveImage field to given value.

### HasRemoveImage

`func (o *CreateProductImageRequestData) HasRemoveImage() bool`

HasRemoveImage returns a boolean if a field has been set.

### GetRemoveFile

`func (o *CreateProductImageRequestData) GetRemoveFile() bool`

GetRemoveFile returns the RemoveFile field if non-nil, zero value otherwise.

### GetRemoveFileOk

`func (o *CreateProductImageRequestData) GetRemoveFileOk() (*bool, bool)`

GetRemoveFileOk returns a tuple with the RemoveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFile

`func (o *CreateProductImageRequestData) SetRemoveFile(v bool)`

SetRemoveFile sets RemoveFile field to given value.

### HasRemoveFile

`func (o *CreateProductImageRequestData) HasRemoveFile() bool`

HasRemoveFile returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *CreateProductImageRequestData) GetVolumeDiscounts() []map[string]interface{}`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *CreateProductImageRequestData) GetVolumeDiscountsOk() (*[]map[string]interface{}, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *CreateProductImageRequestData) SetVolumeDiscounts(v []map[string]interface{})`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *CreateProductImageRequestData) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *CreateProductImageRequestData) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *CreateProductImageRequestData) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *CreateProductImageRequestData) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *CreateProductImageRequestData) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *CreateProductImageRequestData) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *CreateProductImageRequestData) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *CreateProductImageRequestData) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *CreateProductImageRequestData) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *CreateProductImageRequestData) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *CreateProductImageRequestData) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *CreateProductImageRequestData) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *CreateProductImageRequestData) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


