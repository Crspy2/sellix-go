# ProductPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the plan. | [optional] 
**Price** | Pointer to **float32** | The price of the subscription plan | [optional] 
**PriceDiscount** | Pointer to **float32** | The default discount applied to this plan as a percent | [optional] 
**Description** | Pointer to **string** | The description for the subscription plan | [optional] 
**TrialPeriod** | Pointer to **string** | The trial period of the subscription plan in days | [optional] 
**RecurringInterval** | Pointer to **string** | The recurring interval type. How often the customer should be invoiced for the subscription. | [optional] 
**RecurringIntervalCount** | Pointer to **string** | Count of days, weeks, months or years in each subscription&#39;s billing period | [optional] 
**PaymentType** | Pointer to **string** | When the customer will be billed for the subscription. ADVANCE — payment will happen on the first day of the period; ARREARS - payment will happen on the last day of the period. | [optional] 
**RenewType** | Pointer to **string** | How subscription renewals should be handled. CALENDAR - payment will happen on the first day of the period (Monday for WEEK, 1st day for MONTH, January 1st for YEAR); ANNIVERSARY - payment will happen on the same day for all invoices (if subscription was created on March 14th, payment will happen on the 14th of each month (if interval is MONTH)) | [optional] 
**BillingType** | Pointer to **string** | How customers should be billed for the subscription if they purchase through the interval. FULL - full price will be invoiced even if only one day left in subscription period; PRORRATED - price will be prorated to used pariod part. If only 50% of the period is left, invoice&#39;s total will be 50% | [optional] 
**Serials** | Pointer to **string** | UNIMPLEMENTED | [optional] 
**ServiceText** | Pointer to **string** | UNIMPLEMENTED | [optional] 
**DynamicWebhook** | Pointer to **string** | DEPRECATED | [optional] 
**Stock** | Pointer to **string** | DEPRECATED | [optional] 
**File** | Pointer to **string** | DEPRECATED | [optional] 
**FileUniqid** | Pointer to **string** | DEPRECATED | [optional] 

## Methods

### NewProductPlanRequest

`func NewProductPlanRequest() *ProductPlanRequest`

NewProductPlanRequest instantiates a new ProductPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPlanRequestWithDefaults

`func NewProductPlanRequestWithDefaults() *ProductPlanRequest`

NewProductPlanRequestWithDefaults instantiates a new ProductPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ProductPlanRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductPlanRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductPlanRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProductPlanRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *ProductPlanRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductPlanRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductPlanRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductPlanRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *ProductPlanRequest) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *ProductPlanRequest) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *ProductPlanRequest) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *ProductPlanRequest) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetDescription

`func (o *ProductPlanRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductPlanRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductPlanRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductPlanRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *ProductPlanRequest) GetTrialPeriod() string`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *ProductPlanRequest) GetTrialPeriodOk() (*string, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *ProductPlanRequest) SetTrialPeriod(v string)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *ProductPlanRequest) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *ProductPlanRequest) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *ProductPlanRequest) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *ProductPlanRequest) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *ProductPlanRequest) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *ProductPlanRequest) GetRecurringIntervalCount() string`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *ProductPlanRequest) GetRecurringIntervalCountOk() (*string, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *ProductPlanRequest) SetRecurringIntervalCount(v string)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *ProductPlanRequest) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetPaymentType

`func (o *ProductPlanRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *ProductPlanRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *ProductPlanRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *ProductPlanRequest) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetRenewType

`func (o *ProductPlanRequest) GetRenewType() string`

GetRenewType returns the RenewType field if non-nil, zero value otherwise.

### GetRenewTypeOk

`func (o *ProductPlanRequest) GetRenewTypeOk() (*string, bool)`

GetRenewTypeOk returns a tuple with the RenewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewType

`func (o *ProductPlanRequest) SetRenewType(v string)`

SetRenewType sets RenewType field to given value.

### HasRenewType

`func (o *ProductPlanRequest) HasRenewType() bool`

HasRenewType returns a boolean if a field has been set.

### GetBillingType

`func (o *ProductPlanRequest) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *ProductPlanRequest) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *ProductPlanRequest) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *ProductPlanRequest) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetSerials

`func (o *ProductPlanRequest) GetSerials() string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ProductPlanRequest) GetSerialsOk() (*string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ProductPlanRequest) SetSerials(v string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ProductPlanRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetServiceText

`func (o *ProductPlanRequest) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *ProductPlanRequest) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *ProductPlanRequest) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *ProductPlanRequest) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *ProductPlanRequest) GetDynamicWebhook() string`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *ProductPlanRequest) GetDynamicWebhookOk() (*string, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *ProductPlanRequest) SetDynamicWebhook(v string)`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *ProductPlanRequest) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetStock

`func (o *ProductPlanRequest) GetStock() string`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductPlanRequest) GetStockOk() (*string, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductPlanRequest) SetStock(v string)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductPlanRequest) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetFile

`func (o *ProductPlanRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ProductPlanRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ProductPlanRequest) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *ProductPlanRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFileUniqid

`func (o *ProductPlanRequest) GetFileUniqid() string`

GetFileUniqid returns the FileUniqid field if non-nil, zero value otherwise.

### GetFileUniqidOk

`func (o *ProductPlanRequest) GetFileUniqidOk() (*string, bool)`

GetFileUniqidOk returns a tuple with the FileUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUniqid

`func (o *ProductPlanRequest) SetFileUniqid(v string)`

SetFileUniqid sets FileUniqid field to given value.

### HasFileUniqid

`func (o *ProductPlanRequest) HasFileUniqid() bool`

HasFileUniqid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


