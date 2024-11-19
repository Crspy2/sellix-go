# ProductPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | The unique identifier for the payment plan. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this resource belongs. | [optional] 
**ProductId** | Pointer to **int32** | ID of the product subscription. | [optional] 
**Status** | Pointer to **string** | The status of the subscription | [optional] 
**Stock** | Pointer to **int32** | The stock of the subscription plan | [optional] 
**DynamicWebhook** | Pointer to **string** | DEPRECATED: The webhook to be sent for events related to thsi product. | [optional] 
**ServiceText** | Pointer to **string** | The service text to be delievered upon delivery (it product type is SERVICE). | [optional] 
**Title** | Pointer to **string** | The title of the plan. | [optional] 
**Price** | Pointer to **float32** | The price of the subscription plan | [optional] 
**PriceDiscount** | Pointer to **float32** | The default discount applied to this plan | [optional] 
**Description** | Pointer to **string** | The description for the subscription plan | [optional] 
**TrialPeriod** | Pointer to **int32** | Whether a trial period is enabled for the subscription plan. Will be null if product is not a subscription | [optional] 
**RecurringInterval** | Pointer to **string** | How often the customer should be invoiced for the subscription. | [optional] 
**RecurringIntervalCount** | Pointer to **int32** | How many times the customer should be invoiced per interval | [optional] 
**PaymentType** | Pointer to **string** | When the customer will be billed for the subscription. | [optional] 
**RenewType** | Pointer to **string** | How subscription renewals should be handled. | [optional] 
**BillingType** | Pointer to **string** | How customers should be billed for the subscription if they purchase through the interval. | [optional] 
**CreatedAt** | Pointer to **int32** | When the subscription plan was created | [optional] 
**UpdatedAt** | Pointer to **int32** | When the subscription plan was last updated | [optional] 
**ProductCurrency** | Pointer to **string** | Default currency for the subscription product | [optional] 
**ProductStock** | Pointer to **int32** | The stock of the subscription plan | [optional] 
**ProductType** | Pointer to **string** | The type of the product | [optional] 
**ProductSubtype** | Pointer to **string** |  | [optional] 
**RealStock** | Pointer to **int32** | The real stock of the subscription plan | [optional] 
**PriceConversions** | Pointer to [**ProductPlanPriceConversions**](ProductPlanPriceConversions.md) |  | [optional] 

## Methods

### NewProductPlan

`func NewProductPlan() *ProductPlan`

NewProductPlan instantiates a new ProductPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPlanWithDefaults

`func NewProductPlanWithDefaults() *ProductPlan`

NewProductPlanWithDefaults instantiates a new ProductPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductPlan) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductPlan) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductPlan) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ProductPlan) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ProductPlan) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ProductPlan) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ProductPlan) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *ProductPlan) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ProductPlan) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ProductPlan) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ProductPlan) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductId

`func (o *ProductPlan) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductPlan) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductPlan) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductPlan) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStatus

`func (o *ProductPlan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductPlan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductPlan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStock

`func (o *ProductPlan) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductPlan) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductPlan) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductPlan) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *ProductPlan) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *ProductPlan) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *ProductPlan) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *ProductPlan) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetServiceText

`func (o *ProductPlan) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *ProductPlan) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *ProductPlan) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *ProductPlan) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetTitle

`func (o *ProductPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductPlan) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductPlan) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProductPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *ProductPlan) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductPlan) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductPlan) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductPlan) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *ProductPlan) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *ProductPlan) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *ProductPlan) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *ProductPlan) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetDescription

`func (o *ProductPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *ProductPlan) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *ProductPlan) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *ProductPlan) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *ProductPlan) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *ProductPlan) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *ProductPlan) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *ProductPlan) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *ProductPlan) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *ProductPlan) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *ProductPlan) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *ProductPlan) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *ProductPlan) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetPaymentType

`func (o *ProductPlan) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *ProductPlan) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *ProductPlan) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *ProductPlan) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetRenewType

`func (o *ProductPlan) GetRenewType() string`

GetRenewType returns the RenewType field if non-nil, zero value otherwise.

### GetRenewTypeOk

`func (o *ProductPlan) GetRenewTypeOk() (*string, bool)`

GetRenewTypeOk returns a tuple with the RenewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewType

`func (o *ProductPlan) SetRenewType(v string)`

SetRenewType sets RenewType field to given value.

### HasRenewType

`func (o *ProductPlan) HasRenewType() bool`

HasRenewType returns a boolean if a field has been set.

### GetBillingType

`func (o *ProductPlan) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *ProductPlan) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *ProductPlan) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *ProductPlan) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductPlan) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductPlan) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductPlan) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProductPlan) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductPlan) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductPlan) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductPlan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProductCurrency

`func (o *ProductPlan) GetProductCurrency() string`

GetProductCurrency returns the ProductCurrency field if non-nil, zero value otherwise.

### GetProductCurrencyOk

`func (o *ProductPlan) GetProductCurrencyOk() (*string, bool)`

GetProductCurrencyOk returns a tuple with the ProductCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCurrency

`func (o *ProductPlan) SetProductCurrency(v string)`

SetProductCurrency sets ProductCurrency field to given value.

### HasProductCurrency

`func (o *ProductPlan) HasProductCurrency() bool`

HasProductCurrency returns a boolean if a field has been set.

### GetProductStock

`func (o *ProductPlan) GetProductStock() int32`

GetProductStock returns the ProductStock field if non-nil, zero value otherwise.

### GetProductStockOk

`func (o *ProductPlan) GetProductStockOk() (*int32, bool)`

GetProductStockOk returns a tuple with the ProductStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStock

`func (o *ProductPlan) SetProductStock(v int32)`

SetProductStock sets ProductStock field to given value.

### HasProductStock

`func (o *ProductPlan) HasProductStock() bool`

HasProductStock returns a boolean if a field has been set.

### GetProductType

`func (o *ProductPlan) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductPlan) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductPlan) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *ProductPlan) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductSubtype

`func (o *ProductPlan) GetProductSubtype() string`

GetProductSubtype returns the ProductSubtype field if non-nil, zero value otherwise.

### GetProductSubtypeOk

`func (o *ProductPlan) GetProductSubtypeOk() (*string, bool)`

GetProductSubtypeOk returns a tuple with the ProductSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubtype

`func (o *ProductPlan) SetProductSubtype(v string)`

SetProductSubtype sets ProductSubtype field to given value.

### HasProductSubtype

`func (o *ProductPlan) HasProductSubtype() bool`

HasProductSubtype returns a boolean if a field has been set.

### GetRealStock

`func (o *ProductPlan) GetRealStock() int32`

GetRealStock returns the RealStock field if non-nil, zero value otherwise.

### GetRealStockOk

`func (o *ProductPlan) GetRealStockOk() (*int32, bool)`

GetRealStockOk returns a tuple with the RealStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealStock

`func (o *ProductPlan) SetRealStock(v int32)`

SetRealStock sets RealStock field to given value.

### HasRealStock

`func (o *ProductPlan) HasRealStock() bool`

HasRealStock returns a boolean if a field has been set.

### GetPriceConversions

`func (o *ProductPlan) GetPriceConversions() ProductPlanPriceConversions`

GetPriceConversions returns the PriceConversions field if non-nil, zero value otherwise.

### GetPriceConversionsOk

`func (o *ProductPlan) GetPriceConversionsOk() (*ProductPlanPriceConversions, bool)`

GetPriceConversionsOk returns a tuple with the PriceConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceConversions

`func (o *ProductPlan) SetPriceConversions(v ProductPlanPriceConversions)`

SetPriceConversions sets PriceConversions field to given value.

### HasPriceConversions

`func (o *ProductPlan) HasPriceConversions() bool`

HasPriceConversions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


