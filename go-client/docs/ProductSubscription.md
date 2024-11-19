# ProductSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | The ID of the product subscription | [optional] 
**ShopId** | Pointer to **int32** | The ID of the shop to which this resource belongs. | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer who made the purchase | [optional] 
**CustomerEmail** | Pointer to **string** | The email of the customer to which the invoice was sent to | [optional] 
**ProductId** | Pointer to **int32** | The internal ID of the product | [optional] 
**ProductPlanId** | Pointer to **int32** | The internal ID of the product subscription plan | [optional] 
**Status** | Pointer to **string** | The status of the subscription | [optional] 
**CreditNodeUsd** | Pointer to **int32** |  | [optional] 
**StartDate** | Pointer to **int32** | The timestamp of the subscription start date | [optional] 
**EndDate** | Pointer to **int32** | The timestamp of the subscription end date | [optional] 
**SkipTrial** | Pointer to **bool** | Whether or not the trial period was skipped | [optional] 
**SkipTrialReason** | Pointer to **string** | The reason the trial period was skipped | [optional] 
**StartEventProcessed** | Pointer to **bool** | Whether or not the subscription&#39;s start has been processed | [optional] 
**EndEventProcessed** | Pointer to **bool** | Whether or not the subscription&#39;s expiration event has been triggered | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the query. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the query has been updated. | [optional] 
**ShopName** | Pointer to **string** | The name of the shop the subscription was made for | [optional] 
**ProductUniqid** | Pointer to **string** | The ID of the product the subscription is for | [optional] 
**Plan** | Pointer to [**[]ProductPlanSubscription**](ProductPlanSubscription.md) |  | [optional] 

## Methods

### NewProductSubscription

`func NewProductSubscription() *ProductSubscription`

NewProductSubscription instantiates a new ProductSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSubscriptionWithDefaults

`func NewProductSubscriptionWithDefaults() *ProductSubscription`

NewProductSubscriptionWithDefaults instantiates a new ProductSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ProductSubscription) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ProductSubscription) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ProductSubscription) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ProductSubscription) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *ProductSubscription) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ProductSubscription) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ProductSubscription) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ProductSubscription) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetCustomerId

`func (o *ProductSubscription) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ProductSubscription) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ProductSubscription) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ProductSubscription) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *ProductSubscription) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *ProductSubscription) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *ProductSubscription) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *ProductSubscription) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetProductId

`func (o *ProductSubscription) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductSubscription) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductSubscription) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductSubscription) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductPlanId

`func (o *ProductSubscription) GetProductPlanId() int32`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *ProductSubscription) GetProductPlanIdOk() (*int32, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *ProductSubscription) SetProductPlanId(v int32)`

SetProductPlanId sets ProductPlanId field to given value.

### HasProductPlanId

`func (o *ProductSubscription) HasProductPlanId() bool`

HasProductPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *ProductSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreditNodeUsd

`func (o *ProductSubscription) GetCreditNodeUsd() int32`

GetCreditNodeUsd returns the CreditNodeUsd field if non-nil, zero value otherwise.

### GetCreditNodeUsdOk

`func (o *ProductSubscription) GetCreditNodeUsdOk() (*int32, bool)`

GetCreditNodeUsdOk returns a tuple with the CreditNodeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNodeUsd

`func (o *ProductSubscription) SetCreditNodeUsd(v int32)`

SetCreditNodeUsd sets CreditNodeUsd field to given value.

### HasCreditNodeUsd

`func (o *ProductSubscription) HasCreditNodeUsd() bool`

HasCreditNodeUsd returns a boolean if a field has been set.

### GetStartDate

`func (o *ProductSubscription) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProductSubscription) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProductSubscription) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProductSubscription) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProductSubscription) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProductSubscription) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProductSubscription) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProductSubscription) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetSkipTrial

`func (o *ProductSubscription) GetSkipTrial() bool`

GetSkipTrial returns the SkipTrial field if non-nil, zero value otherwise.

### GetSkipTrialOk

`func (o *ProductSubscription) GetSkipTrialOk() (*bool, bool)`

GetSkipTrialOk returns a tuple with the SkipTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTrial

`func (o *ProductSubscription) SetSkipTrial(v bool)`

SetSkipTrial sets SkipTrial field to given value.

### HasSkipTrial

`func (o *ProductSubscription) HasSkipTrial() bool`

HasSkipTrial returns a boolean if a field has been set.

### GetSkipTrialReason

`func (o *ProductSubscription) GetSkipTrialReason() string`

GetSkipTrialReason returns the SkipTrialReason field if non-nil, zero value otherwise.

### GetSkipTrialReasonOk

`func (o *ProductSubscription) GetSkipTrialReasonOk() (*string, bool)`

GetSkipTrialReasonOk returns a tuple with the SkipTrialReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTrialReason

`func (o *ProductSubscription) SetSkipTrialReason(v string)`

SetSkipTrialReason sets SkipTrialReason field to given value.

### HasSkipTrialReason

`func (o *ProductSubscription) HasSkipTrialReason() bool`

HasSkipTrialReason returns a boolean if a field has been set.

### GetStartEventProcessed

`func (o *ProductSubscription) GetStartEventProcessed() bool`

GetStartEventProcessed returns the StartEventProcessed field if non-nil, zero value otherwise.

### GetStartEventProcessedOk

`func (o *ProductSubscription) GetStartEventProcessedOk() (*bool, bool)`

GetStartEventProcessedOk returns a tuple with the StartEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartEventProcessed

`func (o *ProductSubscription) SetStartEventProcessed(v bool)`

SetStartEventProcessed sets StartEventProcessed field to given value.

### HasStartEventProcessed

`func (o *ProductSubscription) HasStartEventProcessed() bool`

HasStartEventProcessed returns a boolean if a field has been set.

### GetEndEventProcessed

`func (o *ProductSubscription) GetEndEventProcessed() bool`

GetEndEventProcessed returns the EndEventProcessed field if non-nil, zero value otherwise.

### GetEndEventProcessedOk

`func (o *ProductSubscription) GetEndEventProcessedOk() (*bool, bool)`

GetEndEventProcessedOk returns a tuple with the EndEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEventProcessed

`func (o *ProductSubscription) SetEndEventProcessed(v bool)`

SetEndEventProcessed sets EndEventProcessed field to given value.

### HasEndEventProcessed

`func (o *ProductSubscription) HasEndEventProcessed() bool`

HasEndEventProcessed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductSubscription) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductSubscription) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductSubscription) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProductSubscription) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductSubscription) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductSubscription) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetShopName

`func (o *ProductSubscription) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *ProductSubscription) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *ProductSubscription) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *ProductSubscription) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetProductUniqid

`func (o *ProductSubscription) GetProductUniqid() string`

GetProductUniqid returns the ProductUniqid field if non-nil, zero value otherwise.

### GetProductUniqidOk

`func (o *ProductSubscription) GetProductUniqidOk() (*string, bool)`

GetProductUniqidOk returns a tuple with the ProductUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUniqid

`func (o *ProductSubscription) SetProductUniqid(v string)`

SetProductUniqid sets ProductUniqid field to given value.

### HasProductUniqid

`func (o *ProductSubscription) HasProductUniqid() bool`

HasProductUniqid returns a boolean if a field has been set.

### GetPlan

`func (o *ProductSubscription) GetPlan() []ProductPlanSubscription`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ProductSubscription) GetPlanOk() (*[]ProductPlanSubscription, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ProductSubscription) SetPlan(v []ProductPlanSubscription)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ProductSubscription) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


