# CreatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Required if product_id and cart are null. Defines the title of the purchase, can be the digital good&#39;s name or a brief summary of what the customer is paying for. | [optional] 
**ProductId** | Pointer to **string** | If specified value, currency and custom_fields will be taken from the product details. | [optional] 
**ProductPlans** | Pointer to **map[string]interface{}** | Required for subscriptions V2. Specified which plan to use for the invoice. | [optional] 
**Cart** | Pointer to [**CreatePaymentRequestCart**](CreatePaymentRequestCart.md) |  | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**PaypalApm** | Pointer to **string** | If gateway is PAYPAL, a paypal_apm (PayPal Alternative Payment Method) can be specified. To retrieve the available PayPal APM for a specific customer session, please refer to the PayPal SDK using window.paypal.FUNDING and fundingSource to filter out available methods. You can also use our documentation on how to process white_label payments. | [optional] 
**CreditCard** | Pointer to **bool** | If gatewa is PAYPAL and no paypal_apm is passed, specify credit_card true to land the customer on the PayPal managed credit card page instead of the onboarding login. | [optional] 
**LexPaymentMethod** | Pointer to **string** | DEPRECATED: If gateway is LEX_HOLDINGS_GROUP, method to be used for the customer to pay. | [optional] 
**Value** | Pointer to **float64** | Required if product_id and cart are null. The customer will be asked to pay for this amount. | [optional] 
**Currency** | Pointer to **string** | Required if product_id and cart are null, defines the currency of value. | [optional] 
**Quantity** | Pointer to **int32** | Can be passed with either product_id null or not. The value or product&#39;s price will be multiplied by this amount. | [optional] 
**CouponCode** | Pointer to **string** | Pass a Sellix coupon code to apply a discount over the invoice. | [optional] 
**Confirmations** | Pointer to **int32** | Cryptocurrency confirmations required to count a transaction over the total crypto amount needed. | [optional] 
**Email** | **string** | Email of the customer. Should you want to handle emails on your own, pass to this field a company email to which PDF receipts of orders will be sent for accounting and log purposes. | 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**FraudShield** | Pointer to [**CreatePaymentRequestFraudShield**](CreatePaymentRequestFraudShield.md) |  | [optional] 
**Webhook** | Pointer to **string** | DEPRECATED: Webhook URL to which updates regarding this payment (invoice) will be sent. | [optional] 
**WhiteLabel** | Pointer to **bool** | Whether or not you want to handle the payments UI. If false, return_url must be specified as it is the website where we will redirect the customer once he has paid through our platform. If true, we will return a full invoice object in the response for you to handle. You can receive updates over invoices and handle subsequent logics through our webhooks. | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePaymentRequest

`func NewCreatePaymentRequest(email string, ) *CreatePaymentRequest`

NewCreatePaymentRequest instantiates a new CreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentRequestWithDefaults

`func NewCreatePaymentRequestWithDefaults() *CreatePaymentRequest`

NewCreatePaymentRequestWithDefaults instantiates a new CreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreatePaymentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePaymentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePaymentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePaymentRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProductId

`func (o *CreatePaymentRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreatePaymentRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreatePaymentRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreatePaymentRequest) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductPlans

`func (o *CreatePaymentRequest) GetProductPlans() map[string]interface{}`

GetProductPlans returns the ProductPlans field if non-nil, zero value otherwise.

### GetProductPlansOk

`func (o *CreatePaymentRequest) GetProductPlansOk() (*map[string]interface{}, bool)`

GetProductPlansOk returns a tuple with the ProductPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlans

`func (o *CreatePaymentRequest) SetProductPlans(v map[string]interface{})`

SetProductPlans sets ProductPlans field to given value.

### HasProductPlans

`func (o *CreatePaymentRequest) HasProductPlans() bool`

HasProductPlans returns a boolean if a field has been set.

### GetCart

`func (o *CreatePaymentRequest) GetCart() CreatePaymentRequestCart`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *CreatePaymentRequest) GetCartOk() (*CreatePaymentRequestCart, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *CreatePaymentRequest) SetCart(v CreatePaymentRequestCart)`

SetCart sets Cart field to given value.

### HasCart

`func (o *CreatePaymentRequest) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetGateway

`func (o *CreatePaymentRequest) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreatePaymentRequest) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreatePaymentRequest) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreatePaymentRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateways

`func (o *CreatePaymentRequest) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *CreatePaymentRequest) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *CreatePaymentRequest) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *CreatePaymentRequest) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetPaypalApm

`func (o *CreatePaymentRequest) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *CreatePaymentRequest) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *CreatePaymentRequest) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *CreatePaymentRequest) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetCreditCard

`func (o *CreatePaymentRequest) GetCreditCard() bool`

GetCreditCard returns the CreditCard field if non-nil, zero value otherwise.

### GetCreditCardOk

`func (o *CreatePaymentRequest) GetCreditCardOk() (*bool, bool)`

GetCreditCardOk returns a tuple with the CreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCard

`func (o *CreatePaymentRequest) SetCreditCard(v bool)`

SetCreditCard sets CreditCard field to given value.

### HasCreditCard

`func (o *CreatePaymentRequest) HasCreditCard() bool`

HasCreditCard returns a boolean if a field has been set.

### GetLexPaymentMethod

`func (o *CreatePaymentRequest) GetLexPaymentMethod() string`

GetLexPaymentMethod returns the LexPaymentMethod field if non-nil, zero value otherwise.

### GetLexPaymentMethodOk

`func (o *CreatePaymentRequest) GetLexPaymentMethodOk() (*string, bool)`

GetLexPaymentMethodOk returns a tuple with the LexPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexPaymentMethod

`func (o *CreatePaymentRequest) SetLexPaymentMethod(v string)`

SetLexPaymentMethod sets LexPaymentMethod field to given value.

### HasLexPaymentMethod

`func (o *CreatePaymentRequest) HasLexPaymentMethod() bool`

HasLexPaymentMethod returns a boolean if a field has been set.

### GetValue

`func (o *CreatePaymentRequest) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreatePaymentRequest) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreatePaymentRequest) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *CreatePaymentRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePaymentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetQuantity

`func (o *CreatePaymentRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreatePaymentRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreatePaymentRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreatePaymentRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCouponCode

`func (o *CreatePaymentRequest) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CreatePaymentRequest) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CreatePaymentRequest) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *CreatePaymentRequest) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetConfirmations

`func (o *CreatePaymentRequest) GetConfirmations() int32`

GetConfirmations returns the Confirmations field if non-nil, zero value otherwise.

### GetConfirmationsOk

`func (o *CreatePaymentRequest) GetConfirmationsOk() (*int32, bool)`

GetConfirmationsOk returns a tuple with the Confirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmations

`func (o *CreatePaymentRequest) SetConfirmations(v int32)`

SetConfirmations sets Confirmations field to given value.

### HasConfirmations

`func (o *CreatePaymentRequest) HasConfirmations() bool`

HasConfirmations returns a boolean if a field has been set.

### GetEmail

`func (o *CreatePaymentRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePaymentRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePaymentRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCustomFields

`func (o *CreatePaymentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreatePaymentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreatePaymentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreatePaymentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFraudShield

`func (o *CreatePaymentRequest) GetFraudShield() CreatePaymentRequestFraudShield`

GetFraudShield returns the FraudShield field if non-nil, zero value otherwise.

### GetFraudShieldOk

`func (o *CreatePaymentRequest) GetFraudShieldOk() (*CreatePaymentRequestFraudShield, bool)`

GetFraudShieldOk returns a tuple with the FraudShield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudShield

`func (o *CreatePaymentRequest) SetFraudShield(v CreatePaymentRequestFraudShield)`

SetFraudShield sets FraudShield field to given value.

### HasFraudShield

`func (o *CreatePaymentRequest) HasFraudShield() bool`

HasFraudShield returns a boolean if a field has been set.

### GetWebhook

`func (o *CreatePaymentRequest) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CreatePaymentRequest) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CreatePaymentRequest) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CreatePaymentRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWhiteLabel

`func (o *CreatePaymentRequest) GetWhiteLabel() bool`

GetWhiteLabel returns the WhiteLabel field if non-nil, zero value otherwise.

### GetWhiteLabelOk

`func (o *CreatePaymentRequest) GetWhiteLabelOk() (*bool, bool)`

GetWhiteLabelOk returns a tuple with the WhiteLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteLabel

`func (o *CreatePaymentRequest) SetWhiteLabel(v bool)`

SetWhiteLabel sets WhiteLabel field to given value.

### HasWhiteLabel

`func (o *CreatePaymentRequest) HasWhiteLabel() bool`

HasWhiteLabel returns a boolean if a field has been set.

### GetReturnUrl

`func (o *CreatePaymentRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreatePaymentRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreatePaymentRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *CreatePaymentRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


