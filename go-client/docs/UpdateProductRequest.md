# UpdateProductRequest

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
**ProductPlans** | Pointer to [**[]ProductPlanRequest**](ProductPlanRequest.md) |  | [optional] 

## Methods

### NewUpdateProductRequest

`func NewUpdateProductRequest() *UpdateProductRequest`

NewUpdateProductRequest instantiates a new UpdateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductRequestWithDefaults

`func NewUpdateProductRequestWithDefaults() *UpdateProductRequest`

NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateProductRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateProductRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateProductRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateProductRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *UpdateProductRequest) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateProductRequest) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateProductRequest) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateProductRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateProductRequest) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateProductRequest) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateProductRequest) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGateways

`func (o *UpdateProductRequest) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *UpdateProductRequest) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *UpdateProductRequest) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *UpdateProductRequest) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetType

`func (o *UpdateProductRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateProductRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateProductRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateProductRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSerials

`func (o *UpdateProductRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *UpdateProductRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *UpdateProductRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *UpdateProductRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSerialsRemoveDuplicates

`func (o *UpdateProductRequest) GetSerialsRemoveDuplicates() bool`

GetSerialsRemoveDuplicates returns the SerialsRemoveDuplicates field if non-nil, zero value otherwise.

### GetSerialsRemoveDuplicatesOk

`func (o *UpdateProductRequest) GetSerialsRemoveDuplicatesOk() (*bool, bool)`

GetSerialsRemoveDuplicatesOk returns a tuple with the SerialsRemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialsRemoveDuplicates

`func (o *UpdateProductRequest) SetSerialsRemoveDuplicates(v bool)`

SetSerialsRemoveDuplicates sets SerialsRemoveDuplicates field to given value.

### HasSerialsRemoveDuplicates

`func (o *UpdateProductRequest) HasSerialsRemoveDuplicates() bool`

HasSerialsRemoveDuplicates returns a boolean if a field has been set.

### GetServiceText

`func (o *UpdateProductRequest) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *UpdateProductRequest) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *UpdateProductRequest) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *UpdateProductRequest) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStock

`func (o *UpdateProductRequest) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *UpdateProductRequest) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *UpdateProductRequest) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *UpdateProductRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *UpdateProductRequest) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *UpdateProductRequest) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *UpdateProductRequest) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *UpdateProductRequest) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *UpdateProductRequest) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *UpdateProductRequest) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *UpdateProductRequest) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *UpdateProductRequest) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetMinQuantity

`func (o *UpdateProductRequest) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *UpdateProductRequest) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *UpdateProductRequest) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *UpdateProductRequest) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *UpdateProductRequest) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *UpdateProductRequest) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *UpdateProductRequest) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *UpdateProductRequest) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetDeliveryText

`func (o *UpdateProductRequest) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *UpdateProductRequest) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *UpdateProductRequest) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *UpdateProductRequest) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateProductRequest) GetCustomFields() []map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateProductRequest) GetCustomFieldsOk() (*[]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateProductRequest) SetCustomFields(v []map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateProductRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *UpdateProductRequest) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *UpdateProductRequest) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *UpdateProductRequest) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *UpdateProductRequest) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *UpdateProductRequest) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *UpdateProductRequest) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *UpdateProductRequest) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *UpdateProductRequest) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetUnlisted

`func (o *UpdateProductRequest) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *UpdateProductRequest) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *UpdateProductRequest) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *UpdateProductRequest) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetPrivate

`func (o *UpdateProductRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *UpdateProductRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *UpdateProductRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *UpdateProductRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *UpdateProductRequest) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *UpdateProductRequest) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *UpdateProductRequest) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *UpdateProductRequest) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetSortPriority

`func (o *UpdateProductRequest) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *UpdateProductRequest) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *UpdateProductRequest) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *UpdateProductRequest) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetWebhooks

`func (o *UpdateProductRequest) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *UpdateProductRequest) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *UpdateProductRequest) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *UpdateProductRequest) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetOnHold

`func (o *UpdateProductRequest) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *UpdateProductRequest) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *UpdateProductRequest) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *UpdateProductRequest) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *UpdateProductRequest) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *UpdateProductRequest) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *UpdateProductRequest) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *UpdateProductRequest) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *UpdateProductRequest) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *UpdateProductRequest) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *UpdateProductRequest) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *UpdateProductRequest) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *UpdateProductRequest) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *UpdateProductRequest) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *UpdateProductRequest) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *UpdateProductRequest) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetRemoveImage

`func (o *UpdateProductRequest) GetRemoveImage() bool`

GetRemoveImage returns the RemoveImage field if non-nil, zero value otherwise.

### GetRemoveImageOk

`func (o *UpdateProductRequest) GetRemoveImageOk() (*bool, bool)`

GetRemoveImageOk returns a tuple with the RemoveImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveImage

`func (o *UpdateProductRequest) SetRemoveImage(v bool)`

SetRemoveImage sets RemoveImage field to given value.

### HasRemoveImage

`func (o *UpdateProductRequest) HasRemoveImage() bool`

HasRemoveImage returns a boolean if a field has been set.

### GetRemoveFile

`func (o *UpdateProductRequest) GetRemoveFile() bool`

GetRemoveFile returns the RemoveFile field if non-nil, zero value otherwise.

### GetRemoveFileOk

`func (o *UpdateProductRequest) GetRemoveFileOk() (*bool, bool)`

GetRemoveFileOk returns a tuple with the RemoveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFile

`func (o *UpdateProductRequest) SetRemoveFile(v bool)`

SetRemoveFile sets RemoveFile field to given value.

### HasRemoveFile

`func (o *UpdateProductRequest) HasRemoveFile() bool`

HasRemoveFile returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *UpdateProductRequest) GetVolumeDiscounts() []map[string]interface{}`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *UpdateProductRequest) GetVolumeDiscountsOk() (*[]map[string]interface{}, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *UpdateProductRequest) SetVolumeDiscounts(v []map[string]interface{})`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *UpdateProductRequest) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *UpdateProductRequest) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *UpdateProductRequest) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *UpdateProductRequest) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *UpdateProductRequest) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *UpdateProductRequest) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *UpdateProductRequest) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *UpdateProductRequest) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *UpdateProductRequest) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *UpdateProductRequest) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *UpdateProductRequest) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *UpdateProductRequest) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *UpdateProductRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetProductPlans

`func (o *UpdateProductRequest) GetProductPlans() []ProductPlanRequest`

GetProductPlans returns the ProductPlans field if non-nil, zero value otherwise.

### GetProductPlansOk

`func (o *UpdateProductRequest) GetProductPlansOk() (*[]ProductPlanRequest, bool)`

GetProductPlansOk returns a tuple with the ProductPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlans

`func (o *UpdateProductRequest) SetProductPlans(v []ProductPlanRequest)`

SetProductPlans sets ProductPlans field to given value.

### HasProductPlans

`func (o *UpdateProductRequest) HasProductPlans() bool`

HasProductPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


