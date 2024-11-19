# CreateSubscriptionv2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**Email** | **string** | The customer&#39;s email address. Invoices and subscription updates will be sent to this email | 
**Currency** | Pointer to **string** | The currency to display prices in. If not specified, the shop&#39;s default currency will be used. | [optional] 
**TaxData** | Pointer to [**CreateSubscriptionv2RequestTaxData**](CreateSubscriptionv2RequestTaxData.md) |  | [optional] 
**ProductId** | **string** | The unique ID of the product to create a subscription for. | 
**Quantity** | **int32** | The quantity of the product to subscribe to. If not specified, the default quantity of the product will be used. | 
**ProductPlanId** | **string** | The id of the product plan to subscribe to. | 
**ProductAddons** | Pointer to **map[string]string** |  | [optional] 
**CouponCode** | Pointer to **int32** | The coupon code to apply to the subscription. | [optional] 
**DiscordIntegrationCode** | Pointer to **string** | The code for the Sellix discord integration. Will be null if discord integration is not enabled. | [optional] 
**CustomFields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateSubscriptionv2Request

`func NewCreateSubscriptionv2Request(email string, productId string, quantity int32, productPlanId string, ) *CreateSubscriptionv2Request`

NewCreateSubscriptionv2Request instantiates a new CreateSubscriptionv2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionv2RequestWithDefaults

`func NewCreateSubscriptionv2RequestWithDefaults() *CreateSubscriptionv2Request`

NewCreateSubscriptionv2RequestWithDefaults instantiates a new CreateSubscriptionv2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *CreateSubscriptionv2Request) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateSubscriptionv2Request) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateSubscriptionv2Request) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateSubscriptionv2Request) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetEmail

`func (o *CreateSubscriptionv2Request) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateSubscriptionv2Request) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateSubscriptionv2Request) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCurrency

`func (o *CreateSubscriptionv2Request) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateSubscriptionv2Request) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateSubscriptionv2Request) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateSubscriptionv2Request) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTaxData

`func (o *CreateSubscriptionv2Request) GetTaxData() CreateSubscriptionv2RequestTaxData`

GetTaxData returns the TaxData field if non-nil, zero value otherwise.

### GetTaxDataOk

`func (o *CreateSubscriptionv2Request) GetTaxDataOk() (*CreateSubscriptionv2RequestTaxData, bool)`

GetTaxDataOk returns a tuple with the TaxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxData

`func (o *CreateSubscriptionv2Request) SetTaxData(v CreateSubscriptionv2RequestTaxData)`

SetTaxData sets TaxData field to given value.

### HasTaxData

`func (o *CreateSubscriptionv2Request) HasTaxData() bool`

HasTaxData returns a boolean if a field has been set.

### GetProductId

`func (o *CreateSubscriptionv2Request) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateSubscriptionv2Request) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateSubscriptionv2Request) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *CreateSubscriptionv2Request) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateSubscriptionv2Request) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateSubscriptionv2Request) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetProductPlanId

`func (o *CreateSubscriptionv2Request) GetProductPlanId() string`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *CreateSubscriptionv2Request) GetProductPlanIdOk() (*string, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *CreateSubscriptionv2Request) SetProductPlanId(v string)`

SetProductPlanId sets ProductPlanId field to given value.


### GetProductAddons

`func (o *CreateSubscriptionv2Request) GetProductAddons() map[string]string`

GetProductAddons returns the ProductAddons field if non-nil, zero value otherwise.

### GetProductAddonsOk

`func (o *CreateSubscriptionv2Request) GetProductAddonsOk() (*map[string]string, bool)`

GetProductAddonsOk returns a tuple with the ProductAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAddons

`func (o *CreateSubscriptionv2Request) SetProductAddons(v map[string]string)`

SetProductAddons sets ProductAddons field to given value.

### HasProductAddons

`func (o *CreateSubscriptionv2Request) HasProductAddons() bool`

HasProductAddons returns a boolean if a field has been set.

### GetCouponCode

`func (o *CreateSubscriptionv2Request) GetCouponCode() int32`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CreateSubscriptionv2Request) GetCouponCodeOk() (*int32, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CreateSubscriptionv2Request) SetCouponCode(v int32)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *CreateSubscriptionv2Request) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetDiscordIntegrationCode

`func (o *CreateSubscriptionv2Request) GetDiscordIntegrationCode() string`

GetDiscordIntegrationCode returns the DiscordIntegrationCode field if non-nil, zero value otherwise.

### GetDiscordIntegrationCodeOk

`func (o *CreateSubscriptionv2Request) GetDiscordIntegrationCodeOk() (*string, bool)`

GetDiscordIntegrationCodeOk returns a tuple with the DiscordIntegrationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordIntegrationCode

`func (o *CreateSubscriptionv2Request) SetDiscordIntegrationCode(v string)`

SetDiscordIntegrationCode sets DiscordIntegrationCode field to given value.

### HasDiscordIntegrationCode

`func (o *CreateSubscriptionv2Request) HasDiscordIntegrationCode() bool`

HasDiscordIntegrationCode returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateSubscriptionv2Request) GetCustomFields() map[string]string`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateSubscriptionv2Request) GetCustomFieldsOk() (*map[string]string, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateSubscriptionv2Request) SetCustomFields(v map[string]string)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateSubscriptionv2Request) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


