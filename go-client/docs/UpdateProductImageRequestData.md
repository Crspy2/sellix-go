# UpdateProductImageRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | Product type. | [optional] 
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

### NewUpdateProductImageRequestData

`func NewUpdateProductImageRequestData() *UpdateProductImageRequestData`

NewUpdateProductImageRequestData instantiates a new UpdateProductImageRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductImageRequestDataWithDefaults

`func NewUpdateProductImageRequestDataWithDefaults() *UpdateProductImageRequestData`

NewUpdateProductImageRequestDataWithDefaults instantiates a new UpdateProductImageRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateProductImageRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateProductImageRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateProductImageRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateProductImageRequestData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *UpdateProductImageRequestData) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateProductImageRequestData) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateProductImageRequestData) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateProductImageRequestData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductImageRequestData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductImageRequestData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductImageRequestData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductImageRequestData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateProductImageRequestData) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateProductImageRequestData) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateProductImageRequestData) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateProductImageRequestData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGateways

`func (o *UpdateProductImageRequestData) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *UpdateProductImageRequestData) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *UpdateProductImageRequestData) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *UpdateProductImageRequestData) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetType

`func (o *UpdateProductImageRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateProductImageRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateProductImageRequestData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateProductImageRequestData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSerials

`func (o *UpdateProductImageRequestData) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *UpdateProductImageRequestData) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *UpdateProductImageRequestData) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *UpdateProductImageRequestData) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSerialsRemoveDuplicates

`func (o *UpdateProductImageRequestData) GetSerialsRemoveDuplicates() bool`

GetSerialsRemoveDuplicates returns the SerialsRemoveDuplicates field if non-nil, zero value otherwise.

### GetSerialsRemoveDuplicatesOk

`func (o *UpdateProductImageRequestData) GetSerialsRemoveDuplicatesOk() (*bool, bool)`

GetSerialsRemoveDuplicatesOk returns a tuple with the SerialsRemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialsRemoveDuplicates

`func (o *UpdateProductImageRequestData) SetSerialsRemoveDuplicates(v bool)`

SetSerialsRemoveDuplicates sets SerialsRemoveDuplicates field to given value.

### HasSerialsRemoveDuplicates

`func (o *UpdateProductImageRequestData) HasSerialsRemoveDuplicates() bool`

HasSerialsRemoveDuplicates returns a boolean if a field has been set.

### GetServiceText

`func (o *UpdateProductImageRequestData) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *UpdateProductImageRequestData) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *UpdateProductImageRequestData) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *UpdateProductImageRequestData) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStock

`func (o *UpdateProductImageRequestData) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *UpdateProductImageRequestData) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *UpdateProductImageRequestData) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *UpdateProductImageRequestData) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *UpdateProductImageRequestData) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *UpdateProductImageRequestData) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *UpdateProductImageRequestData) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *UpdateProductImageRequestData) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *UpdateProductImageRequestData) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *UpdateProductImageRequestData) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *UpdateProductImageRequestData) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *UpdateProductImageRequestData) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetMinQuantity

`func (o *UpdateProductImageRequestData) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *UpdateProductImageRequestData) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *UpdateProductImageRequestData) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *UpdateProductImageRequestData) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *UpdateProductImageRequestData) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *UpdateProductImageRequestData) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *UpdateProductImageRequestData) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *UpdateProductImageRequestData) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetDeliveryText

`func (o *UpdateProductImageRequestData) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *UpdateProductImageRequestData) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *UpdateProductImageRequestData) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *UpdateProductImageRequestData) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateProductImageRequestData) GetCustomFields() []map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateProductImageRequestData) GetCustomFieldsOk() (*[]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateProductImageRequestData) SetCustomFields(v []map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateProductImageRequestData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *UpdateProductImageRequestData) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *UpdateProductImageRequestData) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *UpdateProductImageRequestData) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *UpdateProductImageRequestData) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *UpdateProductImageRequestData) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *UpdateProductImageRequestData) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *UpdateProductImageRequestData) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *UpdateProductImageRequestData) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetUnlisted

`func (o *UpdateProductImageRequestData) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *UpdateProductImageRequestData) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *UpdateProductImageRequestData) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *UpdateProductImageRequestData) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetPrivate

`func (o *UpdateProductImageRequestData) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *UpdateProductImageRequestData) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *UpdateProductImageRequestData) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *UpdateProductImageRequestData) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *UpdateProductImageRequestData) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *UpdateProductImageRequestData) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *UpdateProductImageRequestData) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *UpdateProductImageRequestData) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetSortPriority

`func (o *UpdateProductImageRequestData) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *UpdateProductImageRequestData) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *UpdateProductImageRequestData) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *UpdateProductImageRequestData) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetWebhooks

`func (o *UpdateProductImageRequestData) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *UpdateProductImageRequestData) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *UpdateProductImageRequestData) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *UpdateProductImageRequestData) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetOnHold

`func (o *UpdateProductImageRequestData) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *UpdateProductImageRequestData) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *UpdateProductImageRequestData) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *UpdateProductImageRequestData) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *UpdateProductImageRequestData) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *UpdateProductImageRequestData) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *UpdateProductImageRequestData) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *UpdateProductImageRequestData) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *UpdateProductImageRequestData) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *UpdateProductImageRequestData) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *UpdateProductImageRequestData) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *UpdateProductImageRequestData) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *UpdateProductImageRequestData) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *UpdateProductImageRequestData) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *UpdateProductImageRequestData) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *UpdateProductImageRequestData) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetRemoveImage

`func (o *UpdateProductImageRequestData) GetRemoveImage() bool`

GetRemoveImage returns the RemoveImage field if non-nil, zero value otherwise.

### GetRemoveImageOk

`func (o *UpdateProductImageRequestData) GetRemoveImageOk() (*bool, bool)`

GetRemoveImageOk returns a tuple with the RemoveImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveImage

`func (o *UpdateProductImageRequestData) SetRemoveImage(v bool)`

SetRemoveImage sets RemoveImage field to given value.

### HasRemoveImage

`func (o *UpdateProductImageRequestData) HasRemoveImage() bool`

HasRemoveImage returns a boolean if a field has been set.

### GetRemoveFile

`func (o *UpdateProductImageRequestData) GetRemoveFile() bool`

GetRemoveFile returns the RemoveFile field if non-nil, zero value otherwise.

### GetRemoveFileOk

`func (o *UpdateProductImageRequestData) GetRemoveFileOk() (*bool, bool)`

GetRemoveFileOk returns a tuple with the RemoveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFile

`func (o *UpdateProductImageRequestData) SetRemoveFile(v bool)`

SetRemoveFile sets RemoveFile field to given value.

### HasRemoveFile

`func (o *UpdateProductImageRequestData) HasRemoveFile() bool`

HasRemoveFile returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *UpdateProductImageRequestData) GetVolumeDiscounts() []map[string]interface{}`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *UpdateProductImageRequestData) GetVolumeDiscountsOk() (*[]map[string]interface{}, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *UpdateProductImageRequestData) SetVolumeDiscounts(v []map[string]interface{})`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *UpdateProductImageRequestData) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *UpdateProductImageRequestData) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *UpdateProductImageRequestData) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *UpdateProductImageRequestData) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *UpdateProductImageRequestData) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *UpdateProductImageRequestData) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *UpdateProductImageRequestData) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *UpdateProductImageRequestData) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *UpdateProductImageRequestData) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *UpdateProductImageRequestData) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *UpdateProductImageRequestData) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *UpdateProductImageRequestData) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *UpdateProductImageRequestData) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


