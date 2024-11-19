# PaymentInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID for the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique id for the payment interval; used throughout Sellix&#39;s API | [optional] 
**ProductSubscriptionId** | Pointer to **int32** | Internal ID for the product subscription | [optional] 
**ProductPlanSubscriptionId** | Pointer to **int32** | Internal ID for the subscription&#39;s product plan | [optional] 
**InvoiceUniqid** | Pointer to **string** | Unique ID of the invoice created for the payment interval. | [optional] 
**PrevPaymentIntervalId** | Pointer to **int32** | Internal ID for the previous payment interval | [optional] 
**Idx** | Pointer to **int32** | The index of the payment interval | [optional] 
**StartDate** | Pointer to **time.Time** | The time for when the payment interval started | [optional] 
**EndDate** | Pointer to **time.Time** | The time for when the payment interval ended | [optional] 
**IntervalType** | Pointer to **string** | The type of interval | [optional] 
**PaymentType** | Pointer to **string** | The billing type of the payment interval | [optional] 
**PaymentDate** | Pointer to **time.Time** | The time for when the payment was made | [optional] 
**PaymentAttempts** | Pointer to **int32** | The number of payment attempts | [optional] 
**LastPaymentAttemptDate** | Pointer to **time.Time** | The time for the last payment attempt | [optional] 
**LastPaymentDetails** | Pointer to **map[string]interface{}** | Payment details for the last payment | [optional] 
**NextPaymentAttemptDate** | Pointer to **time.Time** | The time for the next payment interval | [optional] 
**Status** | Pointer to **string** | The status of the payment interval | [optional] 
**UsedCoupons** | Pointer to **[]string** | Coupons used for the payment interval | [optional] 
**UsedCustomerBalance** | Pointer to **float64** | The amount of customer balance used for the payment interval | [optional] 
**UsedCustomerBalanceCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**UsedAddons** | Pointer to **[]string** | Addons used for the payment interval | [optional] 
**TrialEndingEmailSent** | Pointer to **bool** | Whether or not the trial ending email has been sent | [optional] 
**UpcomingRenewalEmailSent** | Pointer to **bool** | Whether or not the upcoming renewal email has been sent | [optional] 
**StartEventProcessed** | Pointer to **int32** | Whether or not the start event has been processed | [optional] 
**EndEventProcessed** | Pointer to **int32** | Whether or not the end event has been processed | [optional] 
**CompletedAt** | Pointer to **time.Time** | The timestamp for when the payment interval was completed | [optional] 
**NextIntervalInfo** | Pointer to [**PaymentIntervalNextIntervalInfo**](PaymentIntervalNextIntervalInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** | The timestamp for when the payment interval was created | [optional] 
**UpdatedAt** | Pointer to **int32** | The timestamp for when the payment interval was last updated | [optional] 
**AutoChargeBlock** | Pointer to **int32** | The number of days before the payment interval is charged | [optional] 
**InvoiceStatus** | Pointer to **string** | The status of the invoice | [optional] 
**InvoiceGateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**ShopId** | Pointer to **int32** | ID of the Sellix shop the payment interval was created for. | [optional] 
**CustomerId** | Pointer to **string** | Unique ID used within Sellix to identify customers across the API. | [optional] 
**PaymentMethodId** | Pointer to **int32** | Internal Sellix ID for the payment method | [optional] 
**PlanSubscriptionGateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**ProductId** | Pointer to **int32** | Internal Sellix ID for the product | [optional] 
**ProductUniqid** | Pointer to **string** | Unique ID of the product | [optional] 
**ProductAutoPaymentMethodRequired** | Pointer to **int32** | Whether or not the product requires an auto payment method | [optional] 

## Methods

### NewPaymentInterval

`func NewPaymentInterval() *PaymentInterval`

NewPaymentInterval instantiates a new PaymentInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntervalWithDefaults

`func NewPaymentIntervalWithDefaults() *PaymentInterval`

NewPaymentIntervalWithDefaults instantiates a new PaymentInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentInterval) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentInterval) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentInterval) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentInterval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *PaymentInterval) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *PaymentInterval) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *PaymentInterval) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *PaymentInterval) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetProductSubscriptionId

`func (o *PaymentInterval) GetProductSubscriptionId() int32`

GetProductSubscriptionId returns the ProductSubscriptionId field if non-nil, zero value otherwise.

### GetProductSubscriptionIdOk

`func (o *PaymentInterval) GetProductSubscriptionIdOk() (*int32, bool)`

GetProductSubscriptionIdOk returns a tuple with the ProductSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionId

`func (o *PaymentInterval) SetProductSubscriptionId(v int32)`

SetProductSubscriptionId sets ProductSubscriptionId field to given value.

### HasProductSubscriptionId

`func (o *PaymentInterval) HasProductSubscriptionId() bool`

HasProductSubscriptionId returns a boolean if a field has been set.

### GetProductPlanSubscriptionId

`func (o *PaymentInterval) GetProductPlanSubscriptionId() int32`

GetProductPlanSubscriptionId returns the ProductPlanSubscriptionId field if non-nil, zero value otherwise.

### GetProductPlanSubscriptionIdOk

`func (o *PaymentInterval) GetProductPlanSubscriptionIdOk() (*int32, bool)`

GetProductPlanSubscriptionIdOk returns a tuple with the ProductPlanSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanSubscriptionId

`func (o *PaymentInterval) SetProductPlanSubscriptionId(v int32)`

SetProductPlanSubscriptionId sets ProductPlanSubscriptionId field to given value.

### HasProductPlanSubscriptionId

`func (o *PaymentInterval) HasProductPlanSubscriptionId() bool`

HasProductPlanSubscriptionId returns a boolean if a field has been set.

### GetInvoiceUniqid

`func (o *PaymentInterval) GetInvoiceUniqid() string`

GetInvoiceUniqid returns the InvoiceUniqid field if non-nil, zero value otherwise.

### GetInvoiceUniqidOk

`func (o *PaymentInterval) GetInvoiceUniqidOk() (*string, bool)`

GetInvoiceUniqidOk returns a tuple with the InvoiceUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceUniqid

`func (o *PaymentInterval) SetInvoiceUniqid(v string)`

SetInvoiceUniqid sets InvoiceUniqid field to given value.

### HasInvoiceUniqid

`func (o *PaymentInterval) HasInvoiceUniqid() bool`

HasInvoiceUniqid returns a boolean if a field has been set.

### GetPrevPaymentIntervalId

`func (o *PaymentInterval) GetPrevPaymentIntervalId() int32`

GetPrevPaymentIntervalId returns the PrevPaymentIntervalId field if non-nil, zero value otherwise.

### GetPrevPaymentIntervalIdOk

`func (o *PaymentInterval) GetPrevPaymentIntervalIdOk() (*int32, bool)`

GetPrevPaymentIntervalIdOk returns a tuple with the PrevPaymentIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPaymentIntervalId

`func (o *PaymentInterval) SetPrevPaymentIntervalId(v int32)`

SetPrevPaymentIntervalId sets PrevPaymentIntervalId field to given value.

### HasPrevPaymentIntervalId

`func (o *PaymentInterval) HasPrevPaymentIntervalId() bool`

HasPrevPaymentIntervalId returns a boolean if a field has been set.

### GetIdx

`func (o *PaymentInterval) GetIdx() int32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *PaymentInterval) GetIdxOk() (*int32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *PaymentInterval) SetIdx(v int32)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *PaymentInterval) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetStartDate

`func (o *PaymentInterval) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PaymentInterval) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PaymentInterval) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PaymentInterval) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PaymentInterval) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PaymentInterval) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PaymentInterval) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PaymentInterval) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIntervalType

`func (o *PaymentInterval) GetIntervalType() string`

GetIntervalType returns the IntervalType field if non-nil, zero value otherwise.

### GetIntervalTypeOk

`func (o *PaymentInterval) GetIntervalTypeOk() (*string, bool)`

GetIntervalTypeOk returns a tuple with the IntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalType

`func (o *PaymentInterval) SetIntervalType(v string)`

SetIntervalType sets IntervalType field to given value.

### HasIntervalType

`func (o *PaymentInterval) HasIntervalType() bool`

HasIntervalType returns a boolean if a field has been set.

### GetPaymentType

`func (o *PaymentInterval) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentInterval) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentInterval) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentInterval) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentDate

`func (o *PaymentInterval) GetPaymentDate() time.Time`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentInterval) GetPaymentDateOk() (*time.Time, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentInterval) SetPaymentDate(v time.Time)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *PaymentInterval) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPaymentAttempts

`func (o *PaymentInterval) GetPaymentAttempts() int32`

GetPaymentAttempts returns the PaymentAttempts field if non-nil, zero value otherwise.

### GetPaymentAttemptsOk

`func (o *PaymentInterval) GetPaymentAttemptsOk() (*int32, bool)`

GetPaymentAttemptsOk returns a tuple with the PaymentAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttempts

`func (o *PaymentInterval) SetPaymentAttempts(v int32)`

SetPaymentAttempts sets PaymentAttempts field to given value.

### HasPaymentAttempts

`func (o *PaymentInterval) HasPaymentAttempts() bool`

HasPaymentAttempts returns a boolean if a field has been set.

### GetLastPaymentAttemptDate

`func (o *PaymentInterval) GetLastPaymentAttemptDate() time.Time`

GetLastPaymentAttemptDate returns the LastPaymentAttemptDate field if non-nil, zero value otherwise.

### GetLastPaymentAttemptDateOk

`func (o *PaymentInterval) GetLastPaymentAttemptDateOk() (*time.Time, bool)`

GetLastPaymentAttemptDateOk returns a tuple with the LastPaymentAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAttemptDate

`func (o *PaymentInterval) SetLastPaymentAttemptDate(v time.Time)`

SetLastPaymentAttemptDate sets LastPaymentAttemptDate field to given value.

### HasLastPaymentAttemptDate

`func (o *PaymentInterval) HasLastPaymentAttemptDate() bool`

HasLastPaymentAttemptDate returns a boolean if a field has been set.

### GetLastPaymentDetails

`func (o *PaymentInterval) GetLastPaymentDetails() map[string]interface{}`

GetLastPaymentDetails returns the LastPaymentDetails field if non-nil, zero value otherwise.

### GetLastPaymentDetailsOk

`func (o *PaymentInterval) GetLastPaymentDetailsOk() (*map[string]interface{}, bool)`

GetLastPaymentDetailsOk returns a tuple with the LastPaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDetails

`func (o *PaymentInterval) SetLastPaymentDetails(v map[string]interface{})`

SetLastPaymentDetails sets LastPaymentDetails field to given value.

### HasLastPaymentDetails

`func (o *PaymentInterval) HasLastPaymentDetails() bool`

HasLastPaymentDetails returns a boolean if a field has been set.

### GetNextPaymentAttemptDate

`func (o *PaymentInterval) GetNextPaymentAttemptDate() time.Time`

GetNextPaymentAttemptDate returns the NextPaymentAttemptDate field if non-nil, zero value otherwise.

### GetNextPaymentAttemptDateOk

`func (o *PaymentInterval) GetNextPaymentAttemptDateOk() (*time.Time, bool)`

GetNextPaymentAttemptDateOk returns a tuple with the NextPaymentAttemptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentAttemptDate

`func (o *PaymentInterval) SetNextPaymentAttemptDate(v time.Time)`

SetNextPaymentAttemptDate sets NextPaymentAttemptDate field to given value.

### HasNextPaymentAttemptDate

`func (o *PaymentInterval) HasNextPaymentAttemptDate() bool`

HasNextPaymentAttemptDate returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentInterval) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInterval) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInterval) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentInterval) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsedCoupons

`func (o *PaymentInterval) GetUsedCoupons() []string`

GetUsedCoupons returns the UsedCoupons field if non-nil, zero value otherwise.

### GetUsedCouponsOk

`func (o *PaymentInterval) GetUsedCouponsOk() (*[]string, bool)`

GetUsedCouponsOk returns a tuple with the UsedCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCoupons

`func (o *PaymentInterval) SetUsedCoupons(v []string)`

SetUsedCoupons sets UsedCoupons field to given value.

### HasUsedCoupons

`func (o *PaymentInterval) HasUsedCoupons() bool`

HasUsedCoupons returns a boolean if a field has been set.

### GetUsedCustomerBalance

`func (o *PaymentInterval) GetUsedCustomerBalance() float64`

GetUsedCustomerBalance returns the UsedCustomerBalance field if non-nil, zero value otherwise.

### GetUsedCustomerBalanceOk

`func (o *PaymentInterval) GetUsedCustomerBalanceOk() (*float64, bool)`

GetUsedCustomerBalanceOk returns a tuple with the UsedCustomerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCustomerBalance

`func (o *PaymentInterval) SetUsedCustomerBalance(v float64)`

SetUsedCustomerBalance sets UsedCustomerBalance field to given value.

### HasUsedCustomerBalance

`func (o *PaymentInterval) HasUsedCustomerBalance() bool`

HasUsedCustomerBalance returns a boolean if a field has been set.

### GetUsedCustomerBalanceCurrency

`func (o *PaymentInterval) GetUsedCustomerBalanceCurrency() Currency`

GetUsedCustomerBalanceCurrency returns the UsedCustomerBalanceCurrency field if non-nil, zero value otherwise.

### GetUsedCustomerBalanceCurrencyOk

`func (o *PaymentInterval) GetUsedCustomerBalanceCurrencyOk() (*Currency, bool)`

GetUsedCustomerBalanceCurrencyOk returns a tuple with the UsedCustomerBalanceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCustomerBalanceCurrency

`func (o *PaymentInterval) SetUsedCustomerBalanceCurrency(v Currency)`

SetUsedCustomerBalanceCurrency sets UsedCustomerBalanceCurrency field to given value.

### HasUsedCustomerBalanceCurrency

`func (o *PaymentInterval) HasUsedCustomerBalanceCurrency() bool`

HasUsedCustomerBalanceCurrency returns a boolean if a field has been set.

### GetUsedAddons

`func (o *PaymentInterval) GetUsedAddons() []string`

GetUsedAddons returns the UsedAddons field if non-nil, zero value otherwise.

### GetUsedAddonsOk

`func (o *PaymentInterval) GetUsedAddonsOk() (*[]string, bool)`

GetUsedAddonsOk returns a tuple with the UsedAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAddons

`func (o *PaymentInterval) SetUsedAddons(v []string)`

SetUsedAddons sets UsedAddons field to given value.

### HasUsedAddons

`func (o *PaymentInterval) HasUsedAddons() bool`

HasUsedAddons returns a boolean if a field has been set.

### GetTrialEndingEmailSent

`func (o *PaymentInterval) GetTrialEndingEmailSent() bool`

GetTrialEndingEmailSent returns the TrialEndingEmailSent field if non-nil, zero value otherwise.

### GetTrialEndingEmailSentOk

`func (o *PaymentInterval) GetTrialEndingEmailSentOk() (*bool, bool)`

GetTrialEndingEmailSentOk returns a tuple with the TrialEndingEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndingEmailSent

`func (o *PaymentInterval) SetTrialEndingEmailSent(v bool)`

SetTrialEndingEmailSent sets TrialEndingEmailSent field to given value.

### HasTrialEndingEmailSent

`func (o *PaymentInterval) HasTrialEndingEmailSent() bool`

HasTrialEndingEmailSent returns a boolean if a field has been set.

### GetUpcomingRenewalEmailSent

`func (o *PaymentInterval) GetUpcomingRenewalEmailSent() bool`

GetUpcomingRenewalEmailSent returns the UpcomingRenewalEmailSent field if non-nil, zero value otherwise.

### GetUpcomingRenewalEmailSentOk

`func (o *PaymentInterval) GetUpcomingRenewalEmailSentOk() (*bool, bool)`

GetUpcomingRenewalEmailSentOk returns a tuple with the UpcomingRenewalEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingRenewalEmailSent

`func (o *PaymentInterval) SetUpcomingRenewalEmailSent(v bool)`

SetUpcomingRenewalEmailSent sets UpcomingRenewalEmailSent field to given value.

### HasUpcomingRenewalEmailSent

`func (o *PaymentInterval) HasUpcomingRenewalEmailSent() bool`

HasUpcomingRenewalEmailSent returns a boolean if a field has been set.

### GetStartEventProcessed

`func (o *PaymentInterval) GetStartEventProcessed() int32`

GetStartEventProcessed returns the StartEventProcessed field if non-nil, zero value otherwise.

### GetStartEventProcessedOk

`func (o *PaymentInterval) GetStartEventProcessedOk() (*int32, bool)`

GetStartEventProcessedOk returns a tuple with the StartEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartEventProcessed

`func (o *PaymentInterval) SetStartEventProcessed(v int32)`

SetStartEventProcessed sets StartEventProcessed field to given value.

### HasStartEventProcessed

`func (o *PaymentInterval) HasStartEventProcessed() bool`

HasStartEventProcessed returns a boolean if a field has been set.

### GetEndEventProcessed

`func (o *PaymentInterval) GetEndEventProcessed() int32`

GetEndEventProcessed returns the EndEventProcessed field if non-nil, zero value otherwise.

### GetEndEventProcessedOk

`func (o *PaymentInterval) GetEndEventProcessedOk() (*int32, bool)`

GetEndEventProcessedOk returns a tuple with the EndEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEventProcessed

`func (o *PaymentInterval) SetEndEventProcessed(v int32)`

SetEndEventProcessed sets EndEventProcessed field to given value.

### HasEndEventProcessed

`func (o *PaymentInterval) HasEndEventProcessed() bool`

HasEndEventProcessed returns a boolean if a field has been set.

### GetCompletedAt

`func (o *PaymentInterval) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *PaymentInterval) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *PaymentInterval) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *PaymentInterval) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetNextIntervalInfo

`func (o *PaymentInterval) GetNextIntervalInfo() PaymentIntervalNextIntervalInfo`

GetNextIntervalInfo returns the NextIntervalInfo field if non-nil, zero value otherwise.

### GetNextIntervalInfoOk

`func (o *PaymentInterval) GetNextIntervalInfoOk() (*PaymentIntervalNextIntervalInfo, bool)`

GetNextIntervalInfoOk returns a tuple with the NextIntervalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextIntervalInfo

`func (o *PaymentInterval) SetNextIntervalInfo(v PaymentIntervalNextIntervalInfo)`

SetNextIntervalInfo sets NextIntervalInfo field to given value.

### HasNextIntervalInfo

`func (o *PaymentInterval) HasNextIntervalInfo() bool`

HasNextIntervalInfo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentInterval) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentInterval) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentInterval) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentInterval) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentInterval) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentInterval) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentInterval) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentInterval) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAutoChargeBlock

`func (o *PaymentInterval) GetAutoChargeBlock() int32`

GetAutoChargeBlock returns the AutoChargeBlock field if non-nil, zero value otherwise.

### GetAutoChargeBlockOk

`func (o *PaymentInterval) GetAutoChargeBlockOk() (*int32, bool)`

GetAutoChargeBlockOk returns a tuple with the AutoChargeBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeBlock

`func (o *PaymentInterval) SetAutoChargeBlock(v int32)`

SetAutoChargeBlock sets AutoChargeBlock field to given value.

### HasAutoChargeBlock

`func (o *PaymentInterval) HasAutoChargeBlock() bool`

HasAutoChargeBlock returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *PaymentInterval) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *PaymentInterval) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *PaymentInterval) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *PaymentInterval) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceGateway

`func (o *PaymentInterval) GetInvoiceGateway() Gateway`

GetInvoiceGateway returns the InvoiceGateway field if non-nil, zero value otherwise.

### GetInvoiceGatewayOk

`func (o *PaymentInterval) GetInvoiceGatewayOk() (*Gateway, bool)`

GetInvoiceGatewayOk returns a tuple with the InvoiceGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceGateway

`func (o *PaymentInterval) SetInvoiceGateway(v Gateway)`

SetInvoiceGateway sets InvoiceGateway field to given value.

### HasInvoiceGateway

`func (o *PaymentInterval) HasInvoiceGateway() bool`

HasInvoiceGateway returns a boolean if a field has been set.

### GetShopId

`func (o *PaymentInterval) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *PaymentInterval) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *PaymentInterval) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *PaymentInterval) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetCustomerId

`func (o *PaymentInterval) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentInterval) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentInterval) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentInterval) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *PaymentInterval) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentInterval) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentInterval) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PaymentInterval) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPlanSubscriptionGateway

`func (o *PaymentInterval) GetPlanSubscriptionGateway() Gateway`

GetPlanSubscriptionGateway returns the PlanSubscriptionGateway field if non-nil, zero value otherwise.

### GetPlanSubscriptionGatewayOk

`func (o *PaymentInterval) GetPlanSubscriptionGatewayOk() (*Gateway, bool)`

GetPlanSubscriptionGatewayOk returns a tuple with the PlanSubscriptionGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanSubscriptionGateway

`func (o *PaymentInterval) SetPlanSubscriptionGateway(v Gateway)`

SetPlanSubscriptionGateway sets PlanSubscriptionGateway field to given value.

### HasPlanSubscriptionGateway

`func (o *PaymentInterval) HasPlanSubscriptionGateway() bool`

HasPlanSubscriptionGateway returns a boolean if a field has been set.

### GetProductId

`func (o *PaymentInterval) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PaymentInterval) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PaymentInterval) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PaymentInterval) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductUniqid

`func (o *PaymentInterval) GetProductUniqid() string`

GetProductUniqid returns the ProductUniqid field if non-nil, zero value otherwise.

### GetProductUniqidOk

`func (o *PaymentInterval) GetProductUniqidOk() (*string, bool)`

GetProductUniqidOk returns a tuple with the ProductUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUniqid

`func (o *PaymentInterval) SetProductUniqid(v string)`

SetProductUniqid sets ProductUniqid field to given value.

### HasProductUniqid

`func (o *PaymentInterval) HasProductUniqid() bool`

HasProductUniqid returns a boolean if a field has been set.

### GetProductAutoPaymentMethodRequired

`func (o *PaymentInterval) GetProductAutoPaymentMethodRequired() int32`

GetProductAutoPaymentMethodRequired returns the ProductAutoPaymentMethodRequired field if non-nil, zero value otherwise.

### GetProductAutoPaymentMethodRequiredOk

`func (o *PaymentInterval) GetProductAutoPaymentMethodRequiredOk() (*int32, bool)`

GetProductAutoPaymentMethodRequiredOk returns a tuple with the ProductAutoPaymentMethodRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAutoPaymentMethodRequired

`func (o *PaymentInterval) SetProductAutoPaymentMethodRequired(v int32)`

SetProductAutoPaymentMethodRequired sets ProductAutoPaymentMethodRequired field to given value.

### HasProductAutoPaymentMethodRequired

`func (o *PaymentInterval) HasProductAutoPaymentMethodRequired() bool`

HasProductAutoPaymentMethodRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


