# ProductPlanSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShopId** | Pointer to **int32** | The internal shop ID to which this resource belongs. | [optional] 
**CustomerId** | Pointer to **string** | The ID of the customer who purchased the subscription plan | [optional] 
**ProductSubscriptionUniqid** | Pointer to **string** | The unique ID of the subscription plan | [optional] 
**ProductId** | Pointer to **int32** | The internal ID of the subscription product | [optional] 
**ProductUniqid** | Pointer to **string** | The external ID of the subscription product | [optional] 
**ProductPlanUniqid** | Pointer to **string** | The external ID of the subscription product plan | [optional] 
**Id** | Pointer to **int32** | The internal ID of the subscription plan | [optional] 
**Uniqid** | Pointer to **string** | The unique ID of the subscription plan | [optional] 
**ProductSubscriptionId** | Pointer to **int32** | The internal ID of the product subscription | [optional] 
**ProductPlanId** | Pointer to **int32** | The internal ID of the product plan | [optional] 
**PaymentIntervalId** | Pointer to **int32** | The internal ID of the payment interval | [optional] 
**Status** | Pointer to **string** | The status of the subscription | [optional] 
**StartDate** | Pointer to **time.Time** | The date-time object of the subscription start date | [optional] 
**EndDate** | Pointer to **time.Time** | The date-time object of the subscription end date | [optional] 
**PaymentMethodId** | Pointer to **int32** | The internal ID of the payment method | [optional] 
**Origin** | Pointer to **string** | How the invoice was created | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Blockchain** | Pointer to [**Blockchain**](Blockchain.md) |  | [optional] 
**PaypalApm** | Pointer to **string** | PayPal Alternative Payment Method name, such as iDEAL, used if gateway is PAYPAL. | [optional] 
**PaypalEmailDelivery** | Pointer to **bool** | Whether or not the invoices will be sent to PayPal email | [optional] 
**StripeApm** | Pointer to **string** | Stripe Alternative Payment Method name, such as iDEAL, used if gateway is STRIPE. | [optional] 
**LexPaymentMethod** | Pointer to **string** | DEPRECATED: If gateway is LEX_HOLDINGS_GROUP, method to be used for the customer to pay. | [optional] 
**CouponId** | Pointer to **string** | If a coupon has been applied, its ID. | [optional] 
**ProductAddons** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**DeveloperInvoice** | Pointer to **bool** | If true, this invoice has been created through the Developers API. | [optional] 
**DeveloperTitle** | Pointer to **string** | If a product_id is not passed through the Developers API, a title must be specified. | [optional] 
**DeveloperWebhook** | Pointer to **string** | Additional webhook URL to which updates regarding this invoice will be sent. | [optional] 
**DeveloperReturnUrl** | Pointer to **string** | If present, the customer will be redirected to this URL after the invoice has been paid. | [optional] 
**DeveloperConfirmations** | Pointer to **bool** | Whether or not the invoice has received developer confirmation | [optional] 
**DeveloperValue** | Pointer to **int32** | Arbitraty value set for development | [optional] 
**PaymentLinkId** | Pointer to **string** | Identifier for the payment link used to pay this invoice. | [optional] 
**IsMarketplace** | Pointer to **bool** | Whether or not the purchase was made in a Sellixmarketplace | [optional] 
**PayoutConfirmation** | Pointer to **map[string]interface{}** |  | [optional] 
**GatewaysAccepted** | Pointer to **string** | Comma separated list of gateways accepted for this invoice. | [optional] 
**TaxData** | Pointer to [**ProductPlanSubscriptionTaxData**](ProductPlanSubscriptionTaxData.md) |  | [optional] 
**DiscordId** | Pointer to **string** | ID of the discord account linked to the purchase at checkout | [optional] 
**DiscordCustomerToken** | Pointer to **string** | Oauth access token for the linked discord account | [optional] 
**DiscordCustomerRefreshToken** | Pointer to **string** | Oauth refresh token used to extend access to the linked discord account | [optional] 
**DiscordCustomerTokenExpiresAt** | Pointer to **int32** | Expiration timestamp for the discord customer oauth access | [optional] 
**DiscordIntegrationRemoved** | Pointer to **bool** | Whether or not the discord integration for the product has been removed | [optional] 
**TelegramUserId** | Pointer to **string** | ID of the telegram account linked to the purchase at checkout | [optional] 
**TelegramGroupId** | Pointer to **string** | ID of the telegram group the account should be added to after payment confirmation | [optional] 
**TelegramIntegrationRemoved** | Pointer to **bool** | Whether or not the telegram integration for the product has been removed | [optional] 
**AffiliateRevenueCustomerId** | Pointer to **string** | The ID fo the customer who receives affiliate revenue from this purchase | [optional] 
**AffiliateRevenueUsed** | Pointer to **int32** | The amount of affiliate revenue used | [optional] 
**AffilitaeRevenueUsedCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**AffiliateRevenueUsedType** | Pointer to **string** | The type of affiliate revenue used | [optional] 
**StartEventProcessed** | Pointer to **bool** | Whether or not the subscription&#39;s start has been processed | [optional] 
**EndEventProcessed** | Pointer to **bool** | Whether or not the subscription&#39;s expiration event has been triggered | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the query. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the query has been updated. | [optional] 
**Intervals** | Pointer to [**[]PaymentInterval**](PaymentInterval.md) |  | [optional] 
**Stock** | Pointer to **int32** | The stock of the subscription plan | [optional] 
**Price** | Pointer to **float32** | The price of the subscription plan | [optional] 
**PriceDiscount** | Pointer to **float32** | The default discount applied to this plan | [optional] 
**TrialPeriod** | Pointer to **int32** | Whether a trial period is enabled for the subscription plan. | [optional] 
**RecurringIntervalCount** | Pointer to **int32** | How many intervals should pass before the customer is invoiced | [optional] 

## Methods

### NewProductPlanSubscription

`func NewProductPlanSubscription() *ProductPlanSubscription`

NewProductPlanSubscription instantiates a new ProductPlanSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPlanSubscriptionWithDefaults

`func NewProductPlanSubscriptionWithDefaults() *ProductPlanSubscription`

NewProductPlanSubscriptionWithDefaults instantiates a new ProductPlanSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShopId

`func (o *ProductPlanSubscription) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ProductPlanSubscription) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ProductPlanSubscription) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ProductPlanSubscription) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetCustomerId

`func (o *ProductPlanSubscription) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ProductPlanSubscription) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ProductPlanSubscription) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ProductPlanSubscription) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetProductSubscriptionUniqid

`func (o *ProductPlanSubscription) GetProductSubscriptionUniqid() string`

GetProductSubscriptionUniqid returns the ProductSubscriptionUniqid field if non-nil, zero value otherwise.

### GetProductSubscriptionUniqidOk

`func (o *ProductPlanSubscription) GetProductSubscriptionUniqidOk() (*string, bool)`

GetProductSubscriptionUniqidOk returns a tuple with the ProductSubscriptionUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionUniqid

`func (o *ProductPlanSubscription) SetProductSubscriptionUniqid(v string)`

SetProductSubscriptionUniqid sets ProductSubscriptionUniqid field to given value.

### HasProductSubscriptionUniqid

`func (o *ProductPlanSubscription) HasProductSubscriptionUniqid() bool`

HasProductSubscriptionUniqid returns a boolean if a field has been set.

### GetProductId

`func (o *ProductPlanSubscription) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductPlanSubscription) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductPlanSubscription) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ProductPlanSubscription) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductUniqid

`func (o *ProductPlanSubscription) GetProductUniqid() string`

GetProductUniqid returns the ProductUniqid field if non-nil, zero value otherwise.

### GetProductUniqidOk

`func (o *ProductPlanSubscription) GetProductUniqidOk() (*string, bool)`

GetProductUniqidOk returns a tuple with the ProductUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUniqid

`func (o *ProductPlanSubscription) SetProductUniqid(v string)`

SetProductUniqid sets ProductUniqid field to given value.

### HasProductUniqid

`func (o *ProductPlanSubscription) HasProductUniqid() bool`

HasProductUniqid returns a boolean if a field has been set.

### GetProductPlanUniqid

`func (o *ProductPlanSubscription) GetProductPlanUniqid() string`

GetProductPlanUniqid returns the ProductPlanUniqid field if non-nil, zero value otherwise.

### GetProductPlanUniqidOk

`func (o *ProductPlanSubscription) GetProductPlanUniqidOk() (*string, bool)`

GetProductPlanUniqidOk returns a tuple with the ProductPlanUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanUniqid

`func (o *ProductPlanSubscription) SetProductPlanUniqid(v string)`

SetProductPlanUniqid sets ProductPlanUniqid field to given value.

### HasProductPlanUniqid

`func (o *ProductPlanSubscription) HasProductPlanUniqid() bool`

HasProductPlanUniqid returns a boolean if a field has been set.

### GetId

`func (o *ProductPlanSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductPlanSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductPlanSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductPlanSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ProductPlanSubscription) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ProductPlanSubscription) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ProductPlanSubscription) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ProductPlanSubscription) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetProductSubscriptionId

`func (o *ProductPlanSubscription) GetProductSubscriptionId() int32`

GetProductSubscriptionId returns the ProductSubscriptionId field if non-nil, zero value otherwise.

### GetProductSubscriptionIdOk

`func (o *ProductPlanSubscription) GetProductSubscriptionIdOk() (*int32, bool)`

GetProductSubscriptionIdOk returns a tuple with the ProductSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionId

`func (o *ProductPlanSubscription) SetProductSubscriptionId(v int32)`

SetProductSubscriptionId sets ProductSubscriptionId field to given value.

### HasProductSubscriptionId

`func (o *ProductPlanSubscription) HasProductSubscriptionId() bool`

HasProductSubscriptionId returns a boolean if a field has been set.

### GetProductPlanId

`func (o *ProductPlanSubscription) GetProductPlanId() int32`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *ProductPlanSubscription) GetProductPlanIdOk() (*int32, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *ProductPlanSubscription) SetProductPlanId(v int32)`

SetProductPlanId sets ProductPlanId field to given value.

### HasProductPlanId

`func (o *ProductPlanSubscription) HasProductPlanId() bool`

HasProductPlanId returns a boolean if a field has been set.

### GetPaymentIntervalId

`func (o *ProductPlanSubscription) GetPaymentIntervalId() int32`

GetPaymentIntervalId returns the PaymentIntervalId field if non-nil, zero value otherwise.

### GetPaymentIntervalIdOk

`func (o *ProductPlanSubscription) GetPaymentIntervalIdOk() (*int32, bool)`

GetPaymentIntervalIdOk returns a tuple with the PaymentIntervalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntervalId

`func (o *ProductPlanSubscription) SetPaymentIntervalId(v int32)`

SetPaymentIntervalId sets PaymentIntervalId field to given value.

### HasPaymentIntervalId

`func (o *ProductPlanSubscription) HasPaymentIntervalId() bool`

HasPaymentIntervalId returns a boolean if a field has been set.

### GetStatus

`func (o *ProductPlanSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductPlanSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductPlanSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProductPlanSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *ProductPlanSubscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProductPlanSubscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProductPlanSubscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProductPlanSubscription) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProductPlanSubscription) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProductPlanSubscription) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProductPlanSubscription) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProductPlanSubscription) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *ProductPlanSubscription) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ProductPlanSubscription) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ProductPlanSubscription) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *ProductPlanSubscription) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetOrigin

`func (o *ProductPlanSubscription) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ProductPlanSubscription) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ProductPlanSubscription) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ProductPlanSubscription) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetGateway

`func (o *ProductPlanSubscription) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ProductPlanSubscription) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ProductPlanSubscription) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ProductPlanSubscription) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductPlanSubscription) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductPlanSubscription) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductPlanSubscription) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductPlanSubscription) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBlockchain

`func (o *ProductPlanSubscription) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ProductPlanSubscription) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ProductPlanSubscription) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *ProductPlanSubscription) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetPaypalApm

`func (o *ProductPlanSubscription) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *ProductPlanSubscription) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *ProductPlanSubscription) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *ProductPlanSubscription) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetPaypalEmailDelivery

`func (o *ProductPlanSubscription) GetPaypalEmailDelivery() bool`

GetPaypalEmailDelivery returns the PaypalEmailDelivery field if non-nil, zero value otherwise.

### GetPaypalEmailDeliveryOk

`func (o *ProductPlanSubscription) GetPaypalEmailDeliveryOk() (*bool, bool)`

GetPaypalEmailDeliveryOk returns a tuple with the PaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmailDelivery

`func (o *ProductPlanSubscription) SetPaypalEmailDelivery(v bool)`

SetPaypalEmailDelivery sets PaypalEmailDelivery field to given value.

### HasPaypalEmailDelivery

`func (o *ProductPlanSubscription) HasPaypalEmailDelivery() bool`

HasPaypalEmailDelivery returns a boolean if a field has been set.

### GetStripeApm

`func (o *ProductPlanSubscription) GetStripeApm() string`

GetStripeApm returns the StripeApm field if non-nil, zero value otherwise.

### GetStripeApmOk

`func (o *ProductPlanSubscription) GetStripeApmOk() (*string, bool)`

GetStripeApmOk returns a tuple with the StripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeApm

`func (o *ProductPlanSubscription) SetStripeApm(v string)`

SetStripeApm sets StripeApm field to given value.

### HasStripeApm

`func (o *ProductPlanSubscription) HasStripeApm() bool`

HasStripeApm returns a boolean if a field has been set.

### GetLexPaymentMethod

`func (o *ProductPlanSubscription) GetLexPaymentMethod() string`

GetLexPaymentMethod returns the LexPaymentMethod field if non-nil, zero value otherwise.

### GetLexPaymentMethodOk

`func (o *ProductPlanSubscription) GetLexPaymentMethodOk() (*string, bool)`

GetLexPaymentMethodOk returns a tuple with the LexPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexPaymentMethod

`func (o *ProductPlanSubscription) SetLexPaymentMethod(v string)`

SetLexPaymentMethod sets LexPaymentMethod field to given value.

### HasLexPaymentMethod

`func (o *ProductPlanSubscription) HasLexPaymentMethod() bool`

HasLexPaymentMethod returns a boolean if a field has been set.

### GetCouponId

`func (o *ProductPlanSubscription) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *ProductPlanSubscription) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *ProductPlanSubscription) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *ProductPlanSubscription) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetProductAddons

`func (o *ProductPlanSubscription) GetProductAddons() map[string]string`

GetProductAddons returns the ProductAddons field if non-nil, zero value otherwise.

### GetProductAddonsOk

`func (o *ProductPlanSubscription) GetProductAddonsOk() (*map[string]string, bool)`

GetProductAddonsOk returns a tuple with the ProductAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAddons

`func (o *ProductPlanSubscription) SetProductAddons(v map[string]string)`

SetProductAddons sets ProductAddons field to given value.

### HasProductAddons

`func (o *ProductPlanSubscription) HasProductAddons() bool`

HasProductAddons returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProductPlanSubscription) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProductPlanSubscription) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProductPlanSubscription) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProductPlanSubscription) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDeveloperInvoice

`func (o *ProductPlanSubscription) GetDeveloperInvoice() bool`

GetDeveloperInvoice returns the DeveloperInvoice field if non-nil, zero value otherwise.

### GetDeveloperInvoiceOk

`func (o *ProductPlanSubscription) GetDeveloperInvoiceOk() (*bool, bool)`

GetDeveloperInvoiceOk returns a tuple with the DeveloperInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperInvoice

`func (o *ProductPlanSubscription) SetDeveloperInvoice(v bool)`

SetDeveloperInvoice sets DeveloperInvoice field to given value.

### HasDeveloperInvoice

`func (o *ProductPlanSubscription) HasDeveloperInvoice() bool`

HasDeveloperInvoice returns a boolean if a field has been set.

### GetDeveloperTitle

`func (o *ProductPlanSubscription) GetDeveloperTitle() string`

GetDeveloperTitle returns the DeveloperTitle field if non-nil, zero value otherwise.

### GetDeveloperTitleOk

`func (o *ProductPlanSubscription) GetDeveloperTitleOk() (*string, bool)`

GetDeveloperTitleOk returns a tuple with the DeveloperTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperTitle

`func (o *ProductPlanSubscription) SetDeveloperTitle(v string)`

SetDeveloperTitle sets DeveloperTitle field to given value.

### HasDeveloperTitle

`func (o *ProductPlanSubscription) HasDeveloperTitle() bool`

HasDeveloperTitle returns a boolean if a field has been set.

### GetDeveloperWebhook

`func (o *ProductPlanSubscription) GetDeveloperWebhook() string`

GetDeveloperWebhook returns the DeveloperWebhook field if non-nil, zero value otherwise.

### GetDeveloperWebhookOk

`func (o *ProductPlanSubscription) GetDeveloperWebhookOk() (*string, bool)`

GetDeveloperWebhookOk returns a tuple with the DeveloperWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperWebhook

`func (o *ProductPlanSubscription) SetDeveloperWebhook(v string)`

SetDeveloperWebhook sets DeveloperWebhook field to given value.

### HasDeveloperWebhook

`func (o *ProductPlanSubscription) HasDeveloperWebhook() bool`

HasDeveloperWebhook returns a boolean if a field has been set.

### GetDeveloperReturnUrl

`func (o *ProductPlanSubscription) GetDeveloperReturnUrl() string`

GetDeveloperReturnUrl returns the DeveloperReturnUrl field if non-nil, zero value otherwise.

### GetDeveloperReturnUrlOk

`func (o *ProductPlanSubscription) GetDeveloperReturnUrlOk() (*string, bool)`

GetDeveloperReturnUrlOk returns a tuple with the DeveloperReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperReturnUrl

`func (o *ProductPlanSubscription) SetDeveloperReturnUrl(v string)`

SetDeveloperReturnUrl sets DeveloperReturnUrl field to given value.

### HasDeveloperReturnUrl

`func (o *ProductPlanSubscription) HasDeveloperReturnUrl() bool`

HasDeveloperReturnUrl returns a boolean if a field has been set.

### GetDeveloperConfirmations

`func (o *ProductPlanSubscription) GetDeveloperConfirmations() bool`

GetDeveloperConfirmations returns the DeveloperConfirmations field if non-nil, zero value otherwise.

### GetDeveloperConfirmationsOk

`func (o *ProductPlanSubscription) GetDeveloperConfirmationsOk() (*bool, bool)`

GetDeveloperConfirmationsOk returns a tuple with the DeveloperConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperConfirmations

`func (o *ProductPlanSubscription) SetDeveloperConfirmations(v bool)`

SetDeveloperConfirmations sets DeveloperConfirmations field to given value.

### HasDeveloperConfirmations

`func (o *ProductPlanSubscription) HasDeveloperConfirmations() bool`

HasDeveloperConfirmations returns a boolean if a field has been set.

### GetDeveloperValue

`func (o *ProductPlanSubscription) GetDeveloperValue() int32`

GetDeveloperValue returns the DeveloperValue field if non-nil, zero value otherwise.

### GetDeveloperValueOk

`func (o *ProductPlanSubscription) GetDeveloperValueOk() (*int32, bool)`

GetDeveloperValueOk returns a tuple with the DeveloperValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperValue

`func (o *ProductPlanSubscription) SetDeveloperValue(v int32)`

SetDeveloperValue sets DeveloperValue field to given value.

### HasDeveloperValue

`func (o *ProductPlanSubscription) HasDeveloperValue() bool`

HasDeveloperValue returns a boolean if a field has been set.

### GetPaymentLinkId

`func (o *ProductPlanSubscription) GetPaymentLinkId() string`

GetPaymentLinkId returns the PaymentLinkId field if non-nil, zero value otherwise.

### GetPaymentLinkIdOk

`func (o *ProductPlanSubscription) GetPaymentLinkIdOk() (*string, bool)`

GetPaymentLinkIdOk returns a tuple with the PaymentLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLinkId

`func (o *ProductPlanSubscription) SetPaymentLinkId(v string)`

SetPaymentLinkId sets PaymentLinkId field to given value.

### HasPaymentLinkId

`func (o *ProductPlanSubscription) HasPaymentLinkId() bool`

HasPaymentLinkId returns a boolean if a field has been set.

### GetIsMarketplace

`func (o *ProductPlanSubscription) GetIsMarketplace() bool`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *ProductPlanSubscription) GetIsMarketplaceOk() (*bool, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *ProductPlanSubscription) SetIsMarketplace(v bool)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *ProductPlanSubscription) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### GetPayoutConfirmation

`func (o *ProductPlanSubscription) GetPayoutConfirmation() map[string]interface{}`

GetPayoutConfirmation returns the PayoutConfirmation field if non-nil, zero value otherwise.

### GetPayoutConfirmationOk

`func (o *ProductPlanSubscription) GetPayoutConfirmationOk() (*map[string]interface{}, bool)`

GetPayoutConfirmationOk returns a tuple with the PayoutConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutConfirmation

`func (o *ProductPlanSubscription) SetPayoutConfirmation(v map[string]interface{})`

SetPayoutConfirmation sets PayoutConfirmation field to given value.

### HasPayoutConfirmation

`func (o *ProductPlanSubscription) HasPayoutConfirmation() bool`

HasPayoutConfirmation returns a boolean if a field has been set.

### GetGatewaysAccepted

`func (o *ProductPlanSubscription) GetGatewaysAccepted() string`

GetGatewaysAccepted returns the GatewaysAccepted field if non-nil, zero value otherwise.

### GetGatewaysAcceptedOk

`func (o *ProductPlanSubscription) GetGatewaysAcceptedOk() (*string, bool)`

GetGatewaysAcceptedOk returns a tuple with the GatewaysAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysAccepted

`func (o *ProductPlanSubscription) SetGatewaysAccepted(v string)`

SetGatewaysAccepted sets GatewaysAccepted field to given value.

### HasGatewaysAccepted

`func (o *ProductPlanSubscription) HasGatewaysAccepted() bool`

HasGatewaysAccepted returns a boolean if a field has been set.

### GetTaxData

`func (o *ProductPlanSubscription) GetTaxData() ProductPlanSubscriptionTaxData`

GetTaxData returns the TaxData field if non-nil, zero value otherwise.

### GetTaxDataOk

`func (o *ProductPlanSubscription) GetTaxDataOk() (*ProductPlanSubscriptionTaxData, bool)`

GetTaxDataOk returns a tuple with the TaxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxData

`func (o *ProductPlanSubscription) SetTaxData(v ProductPlanSubscriptionTaxData)`

SetTaxData sets TaxData field to given value.

### HasTaxData

`func (o *ProductPlanSubscription) HasTaxData() bool`

HasTaxData returns a boolean if a field has been set.

### GetDiscordId

`func (o *ProductPlanSubscription) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *ProductPlanSubscription) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *ProductPlanSubscription) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *ProductPlanSubscription) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### GetDiscordCustomerToken

`func (o *ProductPlanSubscription) GetDiscordCustomerToken() string`

GetDiscordCustomerToken returns the DiscordCustomerToken field if non-nil, zero value otherwise.

### GetDiscordCustomerTokenOk

`func (o *ProductPlanSubscription) GetDiscordCustomerTokenOk() (*string, bool)`

GetDiscordCustomerTokenOk returns a tuple with the DiscordCustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerToken

`func (o *ProductPlanSubscription) SetDiscordCustomerToken(v string)`

SetDiscordCustomerToken sets DiscordCustomerToken field to given value.

### HasDiscordCustomerToken

`func (o *ProductPlanSubscription) HasDiscordCustomerToken() bool`

HasDiscordCustomerToken returns a boolean if a field has been set.

### GetDiscordCustomerRefreshToken

`func (o *ProductPlanSubscription) GetDiscordCustomerRefreshToken() string`

GetDiscordCustomerRefreshToken returns the DiscordCustomerRefreshToken field if non-nil, zero value otherwise.

### GetDiscordCustomerRefreshTokenOk

`func (o *ProductPlanSubscription) GetDiscordCustomerRefreshTokenOk() (*string, bool)`

GetDiscordCustomerRefreshTokenOk returns a tuple with the DiscordCustomerRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerRefreshToken

`func (o *ProductPlanSubscription) SetDiscordCustomerRefreshToken(v string)`

SetDiscordCustomerRefreshToken sets DiscordCustomerRefreshToken field to given value.

### HasDiscordCustomerRefreshToken

`func (o *ProductPlanSubscription) HasDiscordCustomerRefreshToken() bool`

HasDiscordCustomerRefreshToken returns a boolean if a field has been set.

### GetDiscordCustomerTokenExpiresAt

`func (o *ProductPlanSubscription) GetDiscordCustomerTokenExpiresAt() int32`

GetDiscordCustomerTokenExpiresAt returns the DiscordCustomerTokenExpiresAt field if non-nil, zero value otherwise.

### GetDiscordCustomerTokenExpiresAtOk

`func (o *ProductPlanSubscription) GetDiscordCustomerTokenExpiresAtOk() (*int32, bool)`

GetDiscordCustomerTokenExpiresAtOk returns a tuple with the DiscordCustomerTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerTokenExpiresAt

`func (o *ProductPlanSubscription) SetDiscordCustomerTokenExpiresAt(v int32)`

SetDiscordCustomerTokenExpiresAt sets DiscordCustomerTokenExpiresAt field to given value.

### HasDiscordCustomerTokenExpiresAt

`func (o *ProductPlanSubscription) HasDiscordCustomerTokenExpiresAt() bool`

HasDiscordCustomerTokenExpiresAt returns a boolean if a field has been set.

### GetDiscordIntegrationRemoved

`func (o *ProductPlanSubscription) GetDiscordIntegrationRemoved() bool`

GetDiscordIntegrationRemoved returns the DiscordIntegrationRemoved field if non-nil, zero value otherwise.

### GetDiscordIntegrationRemovedOk

`func (o *ProductPlanSubscription) GetDiscordIntegrationRemovedOk() (*bool, bool)`

GetDiscordIntegrationRemovedOk returns a tuple with the DiscordIntegrationRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordIntegrationRemoved

`func (o *ProductPlanSubscription) SetDiscordIntegrationRemoved(v bool)`

SetDiscordIntegrationRemoved sets DiscordIntegrationRemoved field to given value.

### HasDiscordIntegrationRemoved

`func (o *ProductPlanSubscription) HasDiscordIntegrationRemoved() bool`

HasDiscordIntegrationRemoved returns a boolean if a field has been set.

### GetTelegramUserId

`func (o *ProductPlanSubscription) GetTelegramUserId() string`

GetTelegramUserId returns the TelegramUserId field if non-nil, zero value otherwise.

### GetTelegramUserIdOk

`func (o *ProductPlanSubscription) GetTelegramUserIdOk() (*string, bool)`

GetTelegramUserIdOk returns a tuple with the TelegramUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramUserId

`func (o *ProductPlanSubscription) SetTelegramUserId(v string)`

SetTelegramUserId sets TelegramUserId field to given value.

### HasTelegramUserId

`func (o *ProductPlanSubscription) HasTelegramUserId() bool`

HasTelegramUserId returns a boolean if a field has been set.

### GetTelegramGroupId

`func (o *ProductPlanSubscription) GetTelegramGroupId() string`

GetTelegramGroupId returns the TelegramGroupId field if non-nil, zero value otherwise.

### GetTelegramGroupIdOk

`func (o *ProductPlanSubscription) GetTelegramGroupIdOk() (*string, bool)`

GetTelegramGroupIdOk returns a tuple with the TelegramGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramGroupId

`func (o *ProductPlanSubscription) SetTelegramGroupId(v string)`

SetTelegramGroupId sets TelegramGroupId field to given value.

### HasTelegramGroupId

`func (o *ProductPlanSubscription) HasTelegramGroupId() bool`

HasTelegramGroupId returns a boolean if a field has been set.

### GetTelegramIntegrationRemoved

`func (o *ProductPlanSubscription) GetTelegramIntegrationRemoved() bool`

GetTelegramIntegrationRemoved returns the TelegramIntegrationRemoved field if non-nil, zero value otherwise.

### GetTelegramIntegrationRemovedOk

`func (o *ProductPlanSubscription) GetTelegramIntegrationRemovedOk() (*bool, bool)`

GetTelegramIntegrationRemovedOk returns a tuple with the TelegramIntegrationRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramIntegrationRemoved

`func (o *ProductPlanSubscription) SetTelegramIntegrationRemoved(v bool)`

SetTelegramIntegrationRemoved sets TelegramIntegrationRemoved field to given value.

### HasTelegramIntegrationRemoved

`func (o *ProductPlanSubscription) HasTelegramIntegrationRemoved() bool`

HasTelegramIntegrationRemoved returns a boolean if a field has been set.

### GetAffiliateRevenueCustomerId

`func (o *ProductPlanSubscription) GetAffiliateRevenueCustomerId() string`

GetAffiliateRevenueCustomerId returns the AffiliateRevenueCustomerId field if non-nil, zero value otherwise.

### GetAffiliateRevenueCustomerIdOk

`func (o *ProductPlanSubscription) GetAffiliateRevenueCustomerIdOk() (*string, bool)`

GetAffiliateRevenueCustomerIdOk returns a tuple with the AffiliateRevenueCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueCustomerId

`func (o *ProductPlanSubscription) SetAffiliateRevenueCustomerId(v string)`

SetAffiliateRevenueCustomerId sets AffiliateRevenueCustomerId field to given value.

### HasAffiliateRevenueCustomerId

`func (o *ProductPlanSubscription) HasAffiliateRevenueCustomerId() bool`

HasAffiliateRevenueCustomerId returns a boolean if a field has been set.

### GetAffiliateRevenueUsed

`func (o *ProductPlanSubscription) GetAffiliateRevenueUsed() int32`

GetAffiliateRevenueUsed returns the AffiliateRevenueUsed field if non-nil, zero value otherwise.

### GetAffiliateRevenueUsedOk

`func (o *ProductPlanSubscription) GetAffiliateRevenueUsedOk() (*int32, bool)`

GetAffiliateRevenueUsedOk returns a tuple with the AffiliateRevenueUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueUsed

`func (o *ProductPlanSubscription) SetAffiliateRevenueUsed(v int32)`

SetAffiliateRevenueUsed sets AffiliateRevenueUsed field to given value.

### HasAffiliateRevenueUsed

`func (o *ProductPlanSubscription) HasAffiliateRevenueUsed() bool`

HasAffiliateRevenueUsed returns a boolean if a field has been set.

### GetAffilitaeRevenueUsedCurrency

`func (o *ProductPlanSubscription) GetAffilitaeRevenueUsedCurrency() Currency`

GetAffilitaeRevenueUsedCurrency returns the AffilitaeRevenueUsedCurrency field if non-nil, zero value otherwise.

### GetAffilitaeRevenueUsedCurrencyOk

`func (o *ProductPlanSubscription) GetAffilitaeRevenueUsedCurrencyOk() (*Currency, bool)`

GetAffilitaeRevenueUsedCurrencyOk returns a tuple with the AffilitaeRevenueUsedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffilitaeRevenueUsedCurrency

`func (o *ProductPlanSubscription) SetAffilitaeRevenueUsedCurrency(v Currency)`

SetAffilitaeRevenueUsedCurrency sets AffilitaeRevenueUsedCurrency field to given value.

### HasAffilitaeRevenueUsedCurrency

`func (o *ProductPlanSubscription) HasAffilitaeRevenueUsedCurrency() bool`

HasAffilitaeRevenueUsedCurrency returns a boolean if a field has been set.

### GetAffiliateRevenueUsedType

`func (o *ProductPlanSubscription) GetAffiliateRevenueUsedType() string`

GetAffiliateRevenueUsedType returns the AffiliateRevenueUsedType field if non-nil, zero value otherwise.

### GetAffiliateRevenueUsedTypeOk

`func (o *ProductPlanSubscription) GetAffiliateRevenueUsedTypeOk() (*string, bool)`

GetAffiliateRevenueUsedTypeOk returns a tuple with the AffiliateRevenueUsedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueUsedType

`func (o *ProductPlanSubscription) SetAffiliateRevenueUsedType(v string)`

SetAffiliateRevenueUsedType sets AffiliateRevenueUsedType field to given value.

### HasAffiliateRevenueUsedType

`func (o *ProductPlanSubscription) HasAffiliateRevenueUsedType() bool`

HasAffiliateRevenueUsedType returns a boolean if a field has been set.

### GetStartEventProcessed

`func (o *ProductPlanSubscription) GetStartEventProcessed() bool`

GetStartEventProcessed returns the StartEventProcessed field if non-nil, zero value otherwise.

### GetStartEventProcessedOk

`func (o *ProductPlanSubscription) GetStartEventProcessedOk() (*bool, bool)`

GetStartEventProcessedOk returns a tuple with the StartEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartEventProcessed

`func (o *ProductPlanSubscription) SetStartEventProcessed(v bool)`

SetStartEventProcessed sets StartEventProcessed field to given value.

### HasStartEventProcessed

`func (o *ProductPlanSubscription) HasStartEventProcessed() bool`

HasStartEventProcessed returns a boolean if a field has been set.

### GetEndEventProcessed

`func (o *ProductPlanSubscription) GetEndEventProcessed() bool`

GetEndEventProcessed returns the EndEventProcessed field if non-nil, zero value otherwise.

### GetEndEventProcessedOk

`func (o *ProductPlanSubscription) GetEndEventProcessedOk() (*bool, bool)`

GetEndEventProcessedOk returns a tuple with the EndEventProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndEventProcessed

`func (o *ProductPlanSubscription) SetEndEventProcessed(v bool)`

SetEndEventProcessed sets EndEventProcessed field to given value.

### HasEndEventProcessed

`func (o *ProductPlanSubscription) HasEndEventProcessed() bool`

HasEndEventProcessed returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProductPlanSubscription) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductPlanSubscription) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductPlanSubscription) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProductPlanSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProductPlanSubscription) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductPlanSubscription) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductPlanSubscription) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProductPlanSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIntervals

`func (o *ProductPlanSubscription) GetIntervals() []PaymentInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ProductPlanSubscription) GetIntervalsOk() (*[]PaymentInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ProductPlanSubscription) SetIntervals(v []PaymentInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ProductPlanSubscription) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetStock

`func (o *ProductPlanSubscription) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ProductPlanSubscription) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ProductPlanSubscription) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ProductPlanSubscription) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetPrice

`func (o *ProductPlanSubscription) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductPlanSubscription) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductPlanSubscription) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductPlanSubscription) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *ProductPlanSubscription) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *ProductPlanSubscription) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *ProductPlanSubscription) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *ProductPlanSubscription) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *ProductPlanSubscription) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *ProductPlanSubscription) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *ProductPlanSubscription) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *ProductPlanSubscription) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *ProductPlanSubscription) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *ProductPlanSubscription) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *ProductPlanSubscription) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *ProductPlanSubscription) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


