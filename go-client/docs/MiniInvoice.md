# MiniInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**RecurringBillingId** | Pointer to **string** | Unique ID of the recurring bill. | [optional] 
**BillingPeriodStart** | Pointer to **float32** | The timestamp for when the billing for the invoice started | [optional] 
**BillingPeriodEnd** | Pointer to **float32** | The timestamp for when the billing for the invoice ended | [optional] 
**Type** | Pointer to **string** | Invoice type. For subscriptions, the invoice type is PRODUCT_SUBSCRIPTION as SUBSCRIPTION is reserved for Sellix-own plans. | [optional] 
**Subtype** | Pointer to **string** | Product type. | [optional] 
**Origin** | Pointer to **string** | How the invoice was created | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this invoice belongs. | [optional] 
**PlatformId** | Pointer to **string** | The ID of the wallet platform used during checkout | [optional] 
**CustomerId** | Pointer to **string** | ID of the customer. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product linked to this invoice, if any. | [optional] 
**ProductVariants** | Pointer to [**[]ProductVariant**](ProductVariant.md) |  | [optional] 
**AffiliateRevenueCustomerId** | Pointer to **string** | ID of the customer who is affiliated to the purchase. | [optional] 
**ProductAddons** | Pointer to [**[]ProductAddonsInner**](ProductAddonsInner.md) |  | [optional] 
**ProductType** | Pointer to **string** | Product type. | [optional] 
**ProductTitle** | Pointer to **string** | Product title. | [optional] 
**SubscriptionId** | Pointer to **int32** | Field reserved for Sellix-own plans. Unique ID of the subscription purchased. | [optional] 
**SubscriptionTime** | Pointer to **int32** | Field reserved for Sellix-own plans. Time, in seconds, of the subscription purchased. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**Blockchain** | Pointer to [**Blockchain**](Blockchain.md) |  | [optional] 
**GatewaysAccepted** | Pointer to **[]string** |  | [optional] 
**PaypalApm** | Pointer to **string** | PayPal Alternative Payment Method name, such as iDEAL, used if gateway is PAYPAL. | [optional] 
**StripeApm** | Pointer to **string** | Stripe Alternative Payment Method name, such as iDEAL, used if gateway is STRIPE. | [optional] 
**Quantity** | Pointer to **int32** | Qauntity of product purchased. | [optional] 
**Total** | Pointer to **float64** | Invoice total, unit_price*quantity where quantity is applicable. | [optional] 
**TotalDisplay** | Pointer to **float64** | Invoice total in the currency chosen. | [optional] 
**TaxPercentage** | Pointer to **float32** | The percentage of tax applied on the transaction | [optional] 
**CouponId** | Pointer to **string** | Unique ID of the coupon, if used, for the discount. | [optional] 
**Discount** | Pointer to **float64** | If a coupon or volume_discount is used, the discount value presents the total amount of discount over the total cost of the invoice. | [optional] 
**DiscountDisplay** | Pointer to **float64** | If a coupon or volume_discount is used, the displayed discount value presents the total amount of discount over the total cost of the invoice. | [optional] 
**DiscountBreakdown** | Pointer to **map[string]interface{}** | Representation of how discount was applied to invoice | [optional] 
**BundleConfig** | Pointer to [**[]BundleConfig**](BundleConfig.md) |  | [optional] 
**VolumeDiscount** | Pointer to **float64** | The discount value applied from a volume_discount. | [optional] 
**VolumeDiscountDisplay** | Pointer to **float64** | The diplayed discount value applied from a volume_discount. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ExchangeRate** | Pointer to **float64** | Exchange rate between currency chosen and USD. | [optional] 
**CryptoExchangeRate** | Pointer to **float64** | Exchange rate between the cryptocurrency chosen (if any) and USD. | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer. | [optional] 
**DiscordCustomerToken** | Pointer to **string** | The oauth access token provided by discord during checkout | [optional] 
**DiscordCustomerRefreshToken** | Pointer to **string** | The oauth refresh token provided by discord during checkout | [optional] 
**DiscordCustomerTokenExpiresAt** | Pointer to **float32** | Timstamp for the expiration of the discord oauth tokens | [optional] 
**PaypalEmailDelivery** | Pointer to **bool** | If true and gateway is PAYPAL, the product will be shipped to the PayPal email on record instead of the customer email. | [optional] 
**PaypalEmail** | Pointer to **string** | Merchant PayPal email. | [optional] 
**PaypalOrderId** | Pointer to **string** | This field contains the PayPal order ID. | [optional] 
**PaypalAuthorizationId** | Pointer to **string** | This field contains the PayPal authorization ID. | [optional] 
**PaypalCaptureId** | Pointer to **string** | This field contains the PayPal capture ID. | [optional] 
**PaypalPayerEmail** | Pointer to **string** | This field is updated after the invoice is completed with the PayPal&#39;s email used for the purchase. | [optional] 
**PaypalFee** | Pointer to **string** | This field is updated after the invoice is completed with the fee taken by PayPal over the invoice total. | [optional] 
**PaypalFeeCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**PaypalApiCredentials** | Pointer to **string** | This field contains the API version of PayPal that is used. | [optional] 
**PaypalSubscriptionId** | Pointer to **int32** | ID of the paypal subscription. | [optional] 
**PaypalSubscriptionLink** | Pointer to **int32** | Link for the merchant to purchase the subscription from. | [optional] 
**LexOrderId** | Pointer to **string** | DEPRECATED | [optional] 
**LexPaymentMethod** | Pointer to **string** | DEPRECATED | [optional] 
**LexSecret** | Pointer to **string** | DEPRECATED | [optional] 
**VirtualPaymentsId** | Pointer to **string** | DEPRECATED | [optional] 
**PaydashPaymentID** | Pointer to **string** | DEPRECATED | [optional] 
**PaydashApiKey** | Pointer to **string** | DEPRECATED | [optional] 
**SkrillEmail** | Pointer to **string** | Merchant Skrill email. | [optional] 
**SkrillSid** | Pointer to **string** | Skrill unique ID linked to the invoice. | [optional] 
**SkrillLink** | Pointer to **string** | Skrill link to redirect the customer to. | [optional] 
**StripeId** | Pointer to **string** | Client ID used to create the STRIPE paymentIntent. | [optional] 
**StripeClientSecret** | Pointer to **string** | Client secret used to create the STRIPE paymentIntent. | [optional] 
**StripePriceId** | Pointer to **string** | If the invoice type is PRODUCT_SUBSCRIPTION or SUBSCRIPTION, refers to the STRIPE price ID. | [optional] 
**StripeCustomerId** | Pointer to **string** | The ID of the stripe customer the invoice is for | [optional] 
**StripeSubscriptionId** | Pointer to **string** | If the invoice type is SUBSCRIPTION, refers to the STRIPE subscription ID. | [optional] 
**StripeInvoiceId** | Pointer to **string** | If the invoice type not is SUBSCRIPTION, refers to the STRIPE invoice ID. | [optional] 
**PerfectmoneyId** | Pointer to **string** | PerfectMoney payment ID linked to the invoice. | [optional] 
**BinanceInvoiceId** | Pointer to **string** | ID for binance invoice | [optional] 
**BinanceQrcode** | Pointer to **string** | Full Binance QRCODE string | [optional] 
**BinanceCheckoutUrl** | Pointer to **string** | Checkout URL for Binance invoice | [optional] 
**BinancePayout** | Pointer to **int32** | Whether or not the binance invoice was payed out | [optional] 
**CashappQrcode** | Pointer to **string** | Full CashApp QRCODE string. | [optional] 
**CashappNote** | Pointer to **string** | Unique note for the customer to leave in the CashApp payment. | [optional] 
**CashappCashtag** | Pointer to **string** | CashApp cashtag used to create the QRCODE. | [optional] 
**CryptoAddress** | Pointer to **string** | Cryptocurrency address linked to this invoice. | [optional] 
**CryptoAmount** | Pointer to **float64** | Cryptocurrency amount converted based on crypto_exchange_rate. | [optional] 
**CryptoReceived** | Pointer to **float64** | Cryptocurrency amount received, paid by the customer. | [optional] 
**CryptoFeePaid** | Pointer to **float32** | The amount of money already sent for the crypto transaction&#39;s fee | [optional] 
**CryptoFeeOwed** | Pointer to **float32** | The amount remaining to be paid for the crypto transaction&#39;s fee | [optional] 
**CryptoUri** | Pointer to **string** | URI used to create the QRCODE. | [optional] 
**BitcoinLnRHash** | Pointer to **string** | The bitcoin lightning r hash if applicable | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Crypto confirmations needed to process the invoice. | [optional] 
**CryptoWalletVersion** | Pointer to **string** | The version of the Sellix Crypto Wallet being used | [optional] 
**CryptoPayout** | Pointer to **bool** | If true, an instant payout for this invoice&#39;s cryptocurrency address has been sent. | [optional] 
**CryptoScheduledPayout** | Pointer to **bool** | If true, a scheduled payout for this invoice&#39;s cryptocurrency address has been sent. | [optional] 
**CryptoPayoutType** | Pointer to **string** | The payout type for the crypto | [optional] 
**FeeBilled** | Pointer to **bool** | If true, the Sellix fee_percentage has already been collected. | [optional] 
**BillInfo** | Pointer to **map[string]interface{}** | If invoice type is MONTHLY_BILL, contains a breakdown of the fees that need to be collected. | [optional] 
**Country** | Pointer to **string** | Customer country. | [optional] 
**Location** | Pointer to **string** | Customer location. | [optional] 
**Ip** | Pointer to **string** | Customer IP. | [optional] 
**IsVpnOrProxy** | Pointer to **bool** | If true, a VPN or Proxy has been detected. | [optional] 
**UserAgent** | Pointer to **string** | Customer User Agent. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**DeveloperInvoice** | Pointer to **bool** | If true, this invoice has been created through the Developers API. | [optional] 
**DeveloperTitle** | Pointer to **string** | If a product_id is not passed through the Developers API, a title must be specified. | [optional] 
**DeveloperWebhook** | Pointer to **string** | Additional webhook URL to which updates regarding this invoice will be sent. | [optional] 
**DeveloperReturnUrl** | Pointer to **string** | If present, the customer will be redirected to this URL after the invoice has been paid. | [optional] 
**PayoutConfiguration** | Pointer to **string** | DEPRECATED. Value is actually a stringified null | [optional] 
**FeePercentage** | Pointer to **int32** | What cut does Sellix take out of the total. To learn more about Sellix fees please refer to https://sellix.io/fees. | [optional] 
**FeeBreakdown** | Pointer to [**FeeBreakdown**](FeeBreakdown.md) |  | [optional] 
**ToProcess** | Pointer to **float32** | Amount of money left to process | [optional] 
**Status** | Pointer to **string** | Status of the invoice. | [optional] 
**StatusDetails** | Pointer to **string** | If CART_PARTIAL_OUT_OF_STOCK, the invoice has been completed but some products in the cart were out of stock. | [optional] 
**VoidDetails** | Pointer to **string** | If the invoice is VOIDED and the product (or all the products in the cart) were out of stock, this additional status is set. | [optional] 
**Environment** | Pointer to **string** | The environment the invoice was made under | [optional] 
**DayValue** | Pointer to **int32** | Day value, number. | [optional] 
**Day** | Pointer to **string** | First three letters of the day name. | [optional] 
**Month** | Pointer to **string** | First three letters of the month name. | [optional] 
**Year** | Pointer to **int32** | : Year number. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the invoice. | [optional] 
**GatewayUpdatedAt** | Pointer to **int32** | Timestamp, available of the last time the gateway has been updated. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the blacklist has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the invoice. | [optional] 
**IsMarketplace** | Pointer to **int32** | Whether or not the invoice was created from a marketplace | [optional] 
**SquarePaymentId** | Pointer to **string** | Payment ID for the squarespace payment method | [optional] 
**VolumePaymentId** | Pointer to **string** | The payment ID for the volume payment gateway | [optional] 
**VolumeReferenceId** | Pointer to **string** | The reference ID for the volume payment gateway | [optional] 
**ProductPlanId** | Pointer to **int32** | The internal ID of the Sellix product plan associated to the subscription v2 product | [optional] 
**Name** | Pointer to **string** | Name of the merchant. | [optional] 
**ShopPaypalCreditCard** | Pointer to **bool** | If true, the merchant has enabled purchase with Credit Card through PayPal. | [optional] 
**ShopForcePaypalEmailDelivery** | Pointer to **bool** | If true, the merchant has enabled invoice shipment to the PayPal customer email. | [optional] 

## Methods

### NewMiniInvoice

`func NewMiniInvoice() *MiniInvoice`

NewMiniInvoice instantiates a new MiniInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiniInvoiceWithDefaults

`func NewMiniInvoiceWithDefaults() *MiniInvoice`

NewMiniInvoiceWithDefaults instantiates a new MiniInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MiniInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MiniInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MiniInvoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MiniInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *MiniInvoice) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *MiniInvoice) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *MiniInvoice) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *MiniInvoice) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetRecurringBillingId

`func (o *MiniInvoice) GetRecurringBillingId() string`

GetRecurringBillingId returns the RecurringBillingId field if non-nil, zero value otherwise.

### GetRecurringBillingIdOk

`func (o *MiniInvoice) GetRecurringBillingIdOk() (*string, bool)`

GetRecurringBillingIdOk returns a tuple with the RecurringBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringBillingId

`func (o *MiniInvoice) SetRecurringBillingId(v string)`

SetRecurringBillingId sets RecurringBillingId field to given value.

### HasRecurringBillingId

`func (o *MiniInvoice) HasRecurringBillingId() bool`

HasRecurringBillingId returns a boolean if a field has been set.

### GetBillingPeriodStart

`func (o *MiniInvoice) GetBillingPeriodStart() float32`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *MiniInvoice) GetBillingPeriodStartOk() (*float32, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *MiniInvoice) SetBillingPeriodStart(v float32)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.

### HasBillingPeriodStart

`func (o *MiniInvoice) HasBillingPeriodStart() bool`

HasBillingPeriodStart returns a boolean if a field has been set.

### GetBillingPeriodEnd

`func (o *MiniInvoice) GetBillingPeriodEnd() float32`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *MiniInvoice) GetBillingPeriodEndOk() (*float32, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *MiniInvoice) SetBillingPeriodEnd(v float32)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.

### HasBillingPeriodEnd

`func (o *MiniInvoice) HasBillingPeriodEnd() bool`

HasBillingPeriodEnd returns a boolean if a field has been set.

### GetType

`func (o *MiniInvoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MiniInvoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MiniInvoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MiniInvoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *MiniInvoice) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MiniInvoice) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MiniInvoice) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MiniInvoice) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetOrigin

`func (o *MiniInvoice) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MiniInvoice) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MiniInvoice) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MiniInvoice) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetShopId

`func (o *MiniInvoice) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *MiniInvoice) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *MiniInvoice) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *MiniInvoice) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetPlatformId

`func (o *MiniInvoice) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *MiniInvoice) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *MiniInvoice) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *MiniInvoice) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetCustomerId

`func (o *MiniInvoice) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MiniInvoice) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MiniInvoice) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *MiniInvoice) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetProductId

`func (o *MiniInvoice) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *MiniInvoice) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *MiniInvoice) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *MiniInvoice) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductVariants

`func (o *MiniInvoice) GetProductVariants() []ProductVariant`

GetProductVariants returns the ProductVariants field if non-nil, zero value otherwise.

### GetProductVariantsOk

`func (o *MiniInvoice) GetProductVariantsOk() (*[]ProductVariant, bool)`

GetProductVariantsOk returns a tuple with the ProductVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariants

`func (o *MiniInvoice) SetProductVariants(v []ProductVariant)`

SetProductVariants sets ProductVariants field to given value.

### HasProductVariants

`func (o *MiniInvoice) HasProductVariants() bool`

HasProductVariants returns a boolean if a field has been set.

### GetAffiliateRevenueCustomerId

`func (o *MiniInvoice) GetAffiliateRevenueCustomerId() string`

GetAffiliateRevenueCustomerId returns the AffiliateRevenueCustomerId field if non-nil, zero value otherwise.

### GetAffiliateRevenueCustomerIdOk

`func (o *MiniInvoice) GetAffiliateRevenueCustomerIdOk() (*string, bool)`

GetAffiliateRevenueCustomerIdOk returns a tuple with the AffiliateRevenueCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueCustomerId

`func (o *MiniInvoice) SetAffiliateRevenueCustomerId(v string)`

SetAffiliateRevenueCustomerId sets AffiliateRevenueCustomerId field to given value.

### HasAffiliateRevenueCustomerId

`func (o *MiniInvoice) HasAffiliateRevenueCustomerId() bool`

HasAffiliateRevenueCustomerId returns a boolean if a field has been set.

### GetProductAddons

`func (o *MiniInvoice) GetProductAddons() []ProductAddonsInner`

GetProductAddons returns the ProductAddons field if non-nil, zero value otherwise.

### GetProductAddonsOk

`func (o *MiniInvoice) GetProductAddonsOk() (*[]ProductAddonsInner, bool)`

GetProductAddonsOk returns a tuple with the ProductAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAddons

`func (o *MiniInvoice) SetProductAddons(v []ProductAddonsInner)`

SetProductAddons sets ProductAddons field to given value.

### HasProductAddons

`func (o *MiniInvoice) HasProductAddons() bool`

HasProductAddons returns a boolean if a field has been set.

### GetProductType

`func (o *MiniInvoice) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *MiniInvoice) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *MiniInvoice) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *MiniInvoice) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductTitle

`func (o *MiniInvoice) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *MiniInvoice) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *MiniInvoice) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *MiniInvoice) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *MiniInvoice) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *MiniInvoice) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *MiniInvoice) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *MiniInvoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionTime

`func (o *MiniInvoice) GetSubscriptionTime() int32`

GetSubscriptionTime returns the SubscriptionTime field if non-nil, zero value otherwise.

### GetSubscriptionTimeOk

`func (o *MiniInvoice) GetSubscriptionTimeOk() (*int32, bool)`

GetSubscriptionTimeOk returns a tuple with the SubscriptionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTime

`func (o *MiniInvoice) SetSubscriptionTime(v int32)`

SetSubscriptionTime sets SubscriptionTime field to given value.

### HasSubscriptionTime

`func (o *MiniInvoice) HasSubscriptionTime() bool`

HasSubscriptionTime returns a boolean if a field has been set.

### GetGateway

`func (o *MiniInvoice) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MiniInvoice) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MiniInvoice) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MiniInvoice) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetBlockchain

`func (o *MiniInvoice) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *MiniInvoice) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *MiniInvoice) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *MiniInvoice) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetGatewaysAccepted

`func (o *MiniInvoice) GetGatewaysAccepted() []string`

GetGatewaysAccepted returns the GatewaysAccepted field if non-nil, zero value otherwise.

### GetGatewaysAcceptedOk

`func (o *MiniInvoice) GetGatewaysAcceptedOk() (*[]string, bool)`

GetGatewaysAcceptedOk returns a tuple with the GatewaysAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysAccepted

`func (o *MiniInvoice) SetGatewaysAccepted(v []string)`

SetGatewaysAccepted sets GatewaysAccepted field to given value.

### HasGatewaysAccepted

`func (o *MiniInvoice) HasGatewaysAccepted() bool`

HasGatewaysAccepted returns a boolean if a field has been set.

### GetPaypalApm

`func (o *MiniInvoice) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *MiniInvoice) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *MiniInvoice) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *MiniInvoice) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetStripeApm

`func (o *MiniInvoice) GetStripeApm() string`

GetStripeApm returns the StripeApm field if non-nil, zero value otherwise.

### GetStripeApmOk

`func (o *MiniInvoice) GetStripeApmOk() (*string, bool)`

GetStripeApmOk returns a tuple with the StripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeApm

`func (o *MiniInvoice) SetStripeApm(v string)`

SetStripeApm sets StripeApm field to given value.

### HasStripeApm

`func (o *MiniInvoice) HasStripeApm() bool`

HasStripeApm returns a boolean if a field has been set.

### GetQuantity

`func (o *MiniInvoice) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MiniInvoice) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MiniInvoice) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MiniInvoice) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTotal

`func (o *MiniInvoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MiniInvoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MiniInvoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MiniInvoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalDisplay

`func (o *MiniInvoice) GetTotalDisplay() float64`

GetTotalDisplay returns the TotalDisplay field if non-nil, zero value otherwise.

### GetTotalDisplayOk

`func (o *MiniInvoice) GetTotalDisplayOk() (*float64, bool)`

GetTotalDisplayOk returns a tuple with the TotalDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisplay

`func (o *MiniInvoice) SetTotalDisplay(v float64)`

SetTotalDisplay sets TotalDisplay field to given value.

### HasTotalDisplay

`func (o *MiniInvoice) HasTotalDisplay() bool`

HasTotalDisplay returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *MiniInvoice) GetTaxPercentage() float32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *MiniInvoice) GetTaxPercentageOk() (*float32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *MiniInvoice) SetTaxPercentage(v float32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *MiniInvoice) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetCouponId

`func (o *MiniInvoice) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *MiniInvoice) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *MiniInvoice) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *MiniInvoice) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetDiscount

`func (o *MiniInvoice) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *MiniInvoice) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *MiniInvoice) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *MiniInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountDisplay

`func (o *MiniInvoice) GetDiscountDisplay() float64`

GetDiscountDisplay returns the DiscountDisplay field if non-nil, zero value otherwise.

### GetDiscountDisplayOk

`func (o *MiniInvoice) GetDiscountDisplayOk() (*float64, bool)`

GetDiscountDisplayOk returns a tuple with the DiscountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountDisplay

`func (o *MiniInvoice) SetDiscountDisplay(v float64)`

SetDiscountDisplay sets DiscountDisplay field to given value.

### HasDiscountDisplay

`func (o *MiniInvoice) HasDiscountDisplay() bool`

HasDiscountDisplay returns a boolean if a field has been set.

### GetDiscountBreakdown

`func (o *MiniInvoice) GetDiscountBreakdown() map[string]interface{}`

GetDiscountBreakdown returns the DiscountBreakdown field if non-nil, zero value otherwise.

### GetDiscountBreakdownOk

`func (o *MiniInvoice) GetDiscountBreakdownOk() (*map[string]interface{}, bool)`

GetDiscountBreakdownOk returns a tuple with the DiscountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBreakdown

`func (o *MiniInvoice) SetDiscountBreakdown(v map[string]interface{})`

SetDiscountBreakdown sets DiscountBreakdown field to given value.

### HasDiscountBreakdown

`func (o *MiniInvoice) HasDiscountBreakdown() bool`

HasDiscountBreakdown returns a boolean if a field has been set.

### GetBundleConfig

`func (o *MiniInvoice) GetBundleConfig() []BundleConfig`

GetBundleConfig returns the BundleConfig field if non-nil, zero value otherwise.

### GetBundleConfigOk

`func (o *MiniInvoice) GetBundleConfigOk() (*[]BundleConfig, bool)`

GetBundleConfigOk returns a tuple with the BundleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleConfig

`func (o *MiniInvoice) SetBundleConfig(v []BundleConfig)`

SetBundleConfig sets BundleConfig field to given value.

### HasBundleConfig

`func (o *MiniInvoice) HasBundleConfig() bool`

HasBundleConfig returns a boolean if a field has been set.

### GetVolumeDiscount

`func (o *MiniInvoice) GetVolumeDiscount() float64`

GetVolumeDiscount returns the VolumeDiscount field if non-nil, zero value otherwise.

### GetVolumeDiscountOk

`func (o *MiniInvoice) GetVolumeDiscountOk() (*float64, bool)`

GetVolumeDiscountOk returns a tuple with the VolumeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscount

`func (o *MiniInvoice) SetVolumeDiscount(v float64)`

SetVolumeDiscount sets VolumeDiscount field to given value.

### HasVolumeDiscount

`func (o *MiniInvoice) HasVolumeDiscount() bool`

HasVolumeDiscount returns a boolean if a field has been set.

### GetVolumeDiscountDisplay

`func (o *MiniInvoice) GetVolumeDiscountDisplay() float64`

GetVolumeDiscountDisplay returns the VolumeDiscountDisplay field if non-nil, zero value otherwise.

### GetVolumeDiscountDisplayOk

`func (o *MiniInvoice) GetVolumeDiscountDisplayOk() (*float64, bool)`

GetVolumeDiscountDisplayOk returns a tuple with the VolumeDiscountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscountDisplay

`func (o *MiniInvoice) SetVolumeDiscountDisplay(v float64)`

SetVolumeDiscountDisplay sets VolumeDiscountDisplay field to given value.

### HasVolumeDiscountDisplay

`func (o *MiniInvoice) HasVolumeDiscountDisplay() bool`

HasVolumeDiscountDisplay returns a boolean if a field has been set.

### GetCurrency

`func (o *MiniInvoice) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MiniInvoice) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MiniInvoice) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MiniInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *MiniInvoice) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *MiniInvoice) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *MiniInvoice) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *MiniInvoice) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetCryptoExchangeRate

`func (o *MiniInvoice) GetCryptoExchangeRate() float64`

GetCryptoExchangeRate returns the CryptoExchangeRate field if non-nil, zero value otherwise.

### GetCryptoExchangeRateOk

`func (o *MiniInvoice) GetCryptoExchangeRateOk() (*float64, bool)`

GetCryptoExchangeRateOk returns a tuple with the CryptoExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoExchangeRate

`func (o *MiniInvoice) SetCryptoExchangeRate(v float64)`

SetCryptoExchangeRate sets CryptoExchangeRate field to given value.

### HasCryptoExchangeRate

`func (o *MiniInvoice) HasCryptoExchangeRate() bool`

HasCryptoExchangeRate returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *MiniInvoice) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *MiniInvoice) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *MiniInvoice) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *MiniInvoice) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetDiscordCustomerToken

`func (o *MiniInvoice) GetDiscordCustomerToken() string`

GetDiscordCustomerToken returns the DiscordCustomerToken field if non-nil, zero value otherwise.

### GetDiscordCustomerTokenOk

`func (o *MiniInvoice) GetDiscordCustomerTokenOk() (*string, bool)`

GetDiscordCustomerTokenOk returns a tuple with the DiscordCustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerToken

`func (o *MiniInvoice) SetDiscordCustomerToken(v string)`

SetDiscordCustomerToken sets DiscordCustomerToken field to given value.

### HasDiscordCustomerToken

`func (o *MiniInvoice) HasDiscordCustomerToken() bool`

HasDiscordCustomerToken returns a boolean if a field has been set.

### GetDiscordCustomerRefreshToken

`func (o *MiniInvoice) GetDiscordCustomerRefreshToken() string`

GetDiscordCustomerRefreshToken returns the DiscordCustomerRefreshToken field if non-nil, zero value otherwise.

### GetDiscordCustomerRefreshTokenOk

`func (o *MiniInvoice) GetDiscordCustomerRefreshTokenOk() (*string, bool)`

GetDiscordCustomerRefreshTokenOk returns a tuple with the DiscordCustomerRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerRefreshToken

`func (o *MiniInvoice) SetDiscordCustomerRefreshToken(v string)`

SetDiscordCustomerRefreshToken sets DiscordCustomerRefreshToken field to given value.

### HasDiscordCustomerRefreshToken

`func (o *MiniInvoice) HasDiscordCustomerRefreshToken() bool`

HasDiscordCustomerRefreshToken returns a boolean if a field has been set.

### GetDiscordCustomerTokenExpiresAt

`func (o *MiniInvoice) GetDiscordCustomerTokenExpiresAt() float32`

GetDiscordCustomerTokenExpiresAt returns the DiscordCustomerTokenExpiresAt field if non-nil, zero value otherwise.

### GetDiscordCustomerTokenExpiresAtOk

`func (o *MiniInvoice) GetDiscordCustomerTokenExpiresAtOk() (*float32, bool)`

GetDiscordCustomerTokenExpiresAtOk returns a tuple with the DiscordCustomerTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordCustomerTokenExpiresAt

`func (o *MiniInvoice) SetDiscordCustomerTokenExpiresAt(v float32)`

SetDiscordCustomerTokenExpiresAt sets DiscordCustomerTokenExpiresAt field to given value.

### HasDiscordCustomerTokenExpiresAt

`func (o *MiniInvoice) HasDiscordCustomerTokenExpiresAt() bool`

HasDiscordCustomerTokenExpiresAt returns a boolean if a field has been set.

### GetPaypalEmailDelivery

`func (o *MiniInvoice) GetPaypalEmailDelivery() bool`

GetPaypalEmailDelivery returns the PaypalEmailDelivery field if non-nil, zero value otherwise.

### GetPaypalEmailDeliveryOk

`func (o *MiniInvoice) GetPaypalEmailDeliveryOk() (*bool, bool)`

GetPaypalEmailDeliveryOk returns a tuple with the PaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmailDelivery

`func (o *MiniInvoice) SetPaypalEmailDelivery(v bool)`

SetPaypalEmailDelivery sets PaypalEmailDelivery field to given value.

### HasPaypalEmailDelivery

`func (o *MiniInvoice) HasPaypalEmailDelivery() bool`

HasPaypalEmailDelivery returns a boolean if a field has been set.

### GetPaypalEmail

`func (o *MiniInvoice) GetPaypalEmail() string`

GetPaypalEmail returns the PaypalEmail field if non-nil, zero value otherwise.

### GetPaypalEmailOk

`func (o *MiniInvoice) GetPaypalEmailOk() (*string, bool)`

GetPaypalEmailOk returns a tuple with the PaypalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmail

`func (o *MiniInvoice) SetPaypalEmail(v string)`

SetPaypalEmail sets PaypalEmail field to given value.

### HasPaypalEmail

`func (o *MiniInvoice) HasPaypalEmail() bool`

HasPaypalEmail returns a boolean if a field has been set.

### GetPaypalOrderId

`func (o *MiniInvoice) GetPaypalOrderId() string`

GetPaypalOrderId returns the PaypalOrderId field if non-nil, zero value otherwise.

### GetPaypalOrderIdOk

`func (o *MiniInvoice) GetPaypalOrderIdOk() (*string, bool)`

GetPaypalOrderIdOk returns a tuple with the PaypalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalOrderId

`func (o *MiniInvoice) SetPaypalOrderId(v string)`

SetPaypalOrderId sets PaypalOrderId field to given value.

### HasPaypalOrderId

`func (o *MiniInvoice) HasPaypalOrderId() bool`

HasPaypalOrderId returns a boolean if a field has been set.

### GetPaypalAuthorizationId

`func (o *MiniInvoice) GetPaypalAuthorizationId() string`

GetPaypalAuthorizationId returns the PaypalAuthorizationId field if non-nil, zero value otherwise.

### GetPaypalAuthorizationIdOk

`func (o *MiniInvoice) GetPaypalAuthorizationIdOk() (*string, bool)`

GetPaypalAuthorizationIdOk returns a tuple with the PaypalAuthorizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalAuthorizationId

`func (o *MiniInvoice) SetPaypalAuthorizationId(v string)`

SetPaypalAuthorizationId sets PaypalAuthorizationId field to given value.

### HasPaypalAuthorizationId

`func (o *MiniInvoice) HasPaypalAuthorizationId() bool`

HasPaypalAuthorizationId returns a boolean if a field has been set.

### GetPaypalCaptureId

`func (o *MiniInvoice) GetPaypalCaptureId() string`

GetPaypalCaptureId returns the PaypalCaptureId field if non-nil, zero value otherwise.

### GetPaypalCaptureIdOk

`func (o *MiniInvoice) GetPaypalCaptureIdOk() (*string, bool)`

GetPaypalCaptureIdOk returns a tuple with the PaypalCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalCaptureId

`func (o *MiniInvoice) SetPaypalCaptureId(v string)`

SetPaypalCaptureId sets PaypalCaptureId field to given value.

### HasPaypalCaptureId

`func (o *MiniInvoice) HasPaypalCaptureId() bool`

HasPaypalCaptureId returns a boolean if a field has been set.

### GetPaypalPayerEmail

`func (o *MiniInvoice) GetPaypalPayerEmail() string`

GetPaypalPayerEmail returns the PaypalPayerEmail field if non-nil, zero value otherwise.

### GetPaypalPayerEmailOk

`func (o *MiniInvoice) GetPaypalPayerEmailOk() (*string, bool)`

GetPaypalPayerEmailOk returns a tuple with the PaypalPayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerEmail

`func (o *MiniInvoice) SetPaypalPayerEmail(v string)`

SetPaypalPayerEmail sets PaypalPayerEmail field to given value.

### HasPaypalPayerEmail

`func (o *MiniInvoice) HasPaypalPayerEmail() bool`

HasPaypalPayerEmail returns a boolean if a field has been set.

### GetPaypalFee

`func (o *MiniInvoice) GetPaypalFee() string`

GetPaypalFee returns the PaypalFee field if non-nil, zero value otherwise.

### GetPaypalFeeOk

`func (o *MiniInvoice) GetPaypalFeeOk() (*string, bool)`

GetPaypalFeeOk returns a tuple with the PaypalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalFee

`func (o *MiniInvoice) SetPaypalFee(v string)`

SetPaypalFee sets PaypalFee field to given value.

### HasPaypalFee

`func (o *MiniInvoice) HasPaypalFee() bool`

HasPaypalFee returns a boolean if a field has been set.

### GetPaypalFeeCurrency

`func (o *MiniInvoice) GetPaypalFeeCurrency() Currency`

GetPaypalFeeCurrency returns the PaypalFeeCurrency field if non-nil, zero value otherwise.

### GetPaypalFeeCurrencyOk

`func (o *MiniInvoice) GetPaypalFeeCurrencyOk() (*Currency, bool)`

GetPaypalFeeCurrencyOk returns a tuple with the PaypalFeeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalFeeCurrency

`func (o *MiniInvoice) SetPaypalFeeCurrency(v Currency)`

SetPaypalFeeCurrency sets PaypalFeeCurrency field to given value.

### HasPaypalFeeCurrency

`func (o *MiniInvoice) HasPaypalFeeCurrency() bool`

HasPaypalFeeCurrency returns a boolean if a field has been set.

### GetPaypalApiCredentials

`func (o *MiniInvoice) GetPaypalApiCredentials() string`

GetPaypalApiCredentials returns the PaypalApiCredentials field if non-nil, zero value otherwise.

### GetPaypalApiCredentialsOk

`func (o *MiniInvoice) GetPaypalApiCredentialsOk() (*string, bool)`

GetPaypalApiCredentialsOk returns a tuple with the PaypalApiCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApiCredentials

`func (o *MiniInvoice) SetPaypalApiCredentials(v string)`

SetPaypalApiCredentials sets PaypalApiCredentials field to given value.

### HasPaypalApiCredentials

`func (o *MiniInvoice) HasPaypalApiCredentials() bool`

HasPaypalApiCredentials returns a boolean if a field has been set.

### GetPaypalSubscriptionId

`func (o *MiniInvoice) GetPaypalSubscriptionId() int32`

GetPaypalSubscriptionId returns the PaypalSubscriptionId field if non-nil, zero value otherwise.

### GetPaypalSubscriptionIdOk

`func (o *MiniInvoice) GetPaypalSubscriptionIdOk() (*int32, bool)`

GetPaypalSubscriptionIdOk returns a tuple with the PaypalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionId

`func (o *MiniInvoice) SetPaypalSubscriptionId(v int32)`

SetPaypalSubscriptionId sets PaypalSubscriptionId field to given value.

### HasPaypalSubscriptionId

`func (o *MiniInvoice) HasPaypalSubscriptionId() bool`

HasPaypalSubscriptionId returns a boolean if a field has been set.

### GetPaypalSubscriptionLink

`func (o *MiniInvoice) GetPaypalSubscriptionLink() int32`

GetPaypalSubscriptionLink returns the PaypalSubscriptionLink field if non-nil, zero value otherwise.

### GetPaypalSubscriptionLinkOk

`func (o *MiniInvoice) GetPaypalSubscriptionLinkOk() (*int32, bool)`

GetPaypalSubscriptionLinkOk returns a tuple with the PaypalSubscriptionLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionLink

`func (o *MiniInvoice) SetPaypalSubscriptionLink(v int32)`

SetPaypalSubscriptionLink sets PaypalSubscriptionLink field to given value.

### HasPaypalSubscriptionLink

`func (o *MiniInvoice) HasPaypalSubscriptionLink() bool`

HasPaypalSubscriptionLink returns a boolean if a field has been set.

### GetLexOrderId

`func (o *MiniInvoice) GetLexOrderId() string`

GetLexOrderId returns the LexOrderId field if non-nil, zero value otherwise.

### GetLexOrderIdOk

`func (o *MiniInvoice) GetLexOrderIdOk() (*string, bool)`

GetLexOrderIdOk returns a tuple with the LexOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexOrderId

`func (o *MiniInvoice) SetLexOrderId(v string)`

SetLexOrderId sets LexOrderId field to given value.

### HasLexOrderId

`func (o *MiniInvoice) HasLexOrderId() bool`

HasLexOrderId returns a boolean if a field has been set.

### GetLexPaymentMethod

`func (o *MiniInvoice) GetLexPaymentMethod() string`

GetLexPaymentMethod returns the LexPaymentMethod field if non-nil, zero value otherwise.

### GetLexPaymentMethodOk

`func (o *MiniInvoice) GetLexPaymentMethodOk() (*string, bool)`

GetLexPaymentMethodOk returns a tuple with the LexPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexPaymentMethod

`func (o *MiniInvoice) SetLexPaymentMethod(v string)`

SetLexPaymentMethod sets LexPaymentMethod field to given value.

### HasLexPaymentMethod

`func (o *MiniInvoice) HasLexPaymentMethod() bool`

HasLexPaymentMethod returns a boolean if a field has been set.

### GetLexSecret

`func (o *MiniInvoice) GetLexSecret() string`

GetLexSecret returns the LexSecret field if non-nil, zero value otherwise.

### GetLexSecretOk

`func (o *MiniInvoice) GetLexSecretOk() (*string, bool)`

GetLexSecretOk returns a tuple with the LexSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexSecret

`func (o *MiniInvoice) SetLexSecret(v string)`

SetLexSecret sets LexSecret field to given value.

### HasLexSecret

`func (o *MiniInvoice) HasLexSecret() bool`

HasLexSecret returns a boolean if a field has been set.

### GetVirtualPaymentsId

`func (o *MiniInvoice) GetVirtualPaymentsId() string`

GetVirtualPaymentsId returns the VirtualPaymentsId field if non-nil, zero value otherwise.

### GetVirtualPaymentsIdOk

`func (o *MiniInvoice) GetVirtualPaymentsIdOk() (*string, bool)`

GetVirtualPaymentsIdOk returns a tuple with the VirtualPaymentsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPaymentsId

`func (o *MiniInvoice) SetVirtualPaymentsId(v string)`

SetVirtualPaymentsId sets VirtualPaymentsId field to given value.

### HasVirtualPaymentsId

`func (o *MiniInvoice) HasVirtualPaymentsId() bool`

HasVirtualPaymentsId returns a boolean if a field has been set.

### GetPaydashPaymentID

`func (o *MiniInvoice) GetPaydashPaymentID() string`

GetPaydashPaymentID returns the PaydashPaymentID field if non-nil, zero value otherwise.

### GetPaydashPaymentIDOk

`func (o *MiniInvoice) GetPaydashPaymentIDOk() (*string, bool)`

GetPaydashPaymentIDOk returns a tuple with the PaydashPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaydashPaymentID

`func (o *MiniInvoice) SetPaydashPaymentID(v string)`

SetPaydashPaymentID sets PaydashPaymentID field to given value.

### HasPaydashPaymentID

`func (o *MiniInvoice) HasPaydashPaymentID() bool`

HasPaydashPaymentID returns a boolean if a field has been set.

### GetPaydashApiKey

`func (o *MiniInvoice) GetPaydashApiKey() string`

GetPaydashApiKey returns the PaydashApiKey field if non-nil, zero value otherwise.

### GetPaydashApiKeyOk

`func (o *MiniInvoice) GetPaydashApiKeyOk() (*string, bool)`

GetPaydashApiKeyOk returns a tuple with the PaydashApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaydashApiKey

`func (o *MiniInvoice) SetPaydashApiKey(v string)`

SetPaydashApiKey sets PaydashApiKey field to given value.

### HasPaydashApiKey

`func (o *MiniInvoice) HasPaydashApiKey() bool`

HasPaydashApiKey returns a boolean if a field has been set.

### GetSkrillEmail

`func (o *MiniInvoice) GetSkrillEmail() string`

GetSkrillEmail returns the SkrillEmail field if non-nil, zero value otherwise.

### GetSkrillEmailOk

`func (o *MiniInvoice) GetSkrillEmailOk() (*string, bool)`

GetSkrillEmailOk returns a tuple with the SkrillEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillEmail

`func (o *MiniInvoice) SetSkrillEmail(v string)`

SetSkrillEmail sets SkrillEmail field to given value.

### HasSkrillEmail

`func (o *MiniInvoice) HasSkrillEmail() bool`

HasSkrillEmail returns a boolean if a field has been set.

### GetSkrillSid

`func (o *MiniInvoice) GetSkrillSid() string`

GetSkrillSid returns the SkrillSid field if non-nil, zero value otherwise.

### GetSkrillSidOk

`func (o *MiniInvoice) GetSkrillSidOk() (*string, bool)`

GetSkrillSidOk returns a tuple with the SkrillSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillSid

`func (o *MiniInvoice) SetSkrillSid(v string)`

SetSkrillSid sets SkrillSid field to given value.

### HasSkrillSid

`func (o *MiniInvoice) HasSkrillSid() bool`

HasSkrillSid returns a boolean if a field has been set.

### GetSkrillLink

`func (o *MiniInvoice) GetSkrillLink() string`

GetSkrillLink returns the SkrillLink field if non-nil, zero value otherwise.

### GetSkrillLinkOk

`func (o *MiniInvoice) GetSkrillLinkOk() (*string, bool)`

GetSkrillLinkOk returns a tuple with the SkrillLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillLink

`func (o *MiniInvoice) SetSkrillLink(v string)`

SetSkrillLink sets SkrillLink field to given value.

### HasSkrillLink

`func (o *MiniInvoice) HasSkrillLink() bool`

HasSkrillLink returns a boolean if a field has been set.

### GetStripeId

`func (o *MiniInvoice) GetStripeId() string`

GetStripeId returns the StripeId field if non-nil, zero value otherwise.

### GetStripeIdOk

`func (o *MiniInvoice) GetStripeIdOk() (*string, bool)`

GetStripeIdOk returns a tuple with the StripeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeId

`func (o *MiniInvoice) SetStripeId(v string)`

SetStripeId sets StripeId field to given value.

### HasStripeId

`func (o *MiniInvoice) HasStripeId() bool`

HasStripeId returns a boolean if a field has been set.

### GetStripeClientSecret

`func (o *MiniInvoice) GetStripeClientSecret() string`

GetStripeClientSecret returns the StripeClientSecret field if non-nil, zero value otherwise.

### GetStripeClientSecretOk

`func (o *MiniInvoice) GetStripeClientSecretOk() (*string, bool)`

GetStripeClientSecretOk returns a tuple with the StripeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeClientSecret

`func (o *MiniInvoice) SetStripeClientSecret(v string)`

SetStripeClientSecret sets StripeClientSecret field to given value.

### HasStripeClientSecret

`func (o *MiniInvoice) HasStripeClientSecret() bool`

HasStripeClientSecret returns a boolean if a field has been set.

### GetStripePriceId

`func (o *MiniInvoice) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *MiniInvoice) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *MiniInvoice) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *MiniInvoice) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *MiniInvoice) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *MiniInvoice) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *MiniInvoice) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *MiniInvoice) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *MiniInvoice) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *MiniInvoice) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *MiniInvoice) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *MiniInvoice) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetStripeInvoiceId

`func (o *MiniInvoice) GetStripeInvoiceId() string`

GetStripeInvoiceId returns the StripeInvoiceId field if non-nil, zero value otherwise.

### GetStripeInvoiceIdOk

`func (o *MiniInvoice) GetStripeInvoiceIdOk() (*string, bool)`

GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoiceId

`func (o *MiniInvoice) SetStripeInvoiceId(v string)`

SetStripeInvoiceId sets StripeInvoiceId field to given value.

### HasStripeInvoiceId

`func (o *MiniInvoice) HasStripeInvoiceId() bool`

HasStripeInvoiceId returns a boolean if a field has been set.

### GetPerfectmoneyId

`func (o *MiniInvoice) GetPerfectmoneyId() string`

GetPerfectmoneyId returns the PerfectmoneyId field if non-nil, zero value otherwise.

### GetPerfectmoneyIdOk

`func (o *MiniInvoice) GetPerfectmoneyIdOk() (*string, bool)`

GetPerfectmoneyIdOk returns a tuple with the PerfectmoneyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectmoneyId

`func (o *MiniInvoice) SetPerfectmoneyId(v string)`

SetPerfectmoneyId sets PerfectmoneyId field to given value.

### HasPerfectmoneyId

`func (o *MiniInvoice) HasPerfectmoneyId() bool`

HasPerfectmoneyId returns a boolean if a field has been set.

### GetBinanceInvoiceId

`func (o *MiniInvoice) GetBinanceInvoiceId() string`

GetBinanceInvoiceId returns the BinanceInvoiceId field if non-nil, zero value otherwise.

### GetBinanceInvoiceIdOk

`func (o *MiniInvoice) GetBinanceInvoiceIdOk() (*string, bool)`

GetBinanceInvoiceIdOk returns a tuple with the BinanceInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceInvoiceId

`func (o *MiniInvoice) SetBinanceInvoiceId(v string)`

SetBinanceInvoiceId sets BinanceInvoiceId field to given value.

### HasBinanceInvoiceId

`func (o *MiniInvoice) HasBinanceInvoiceId() bool`

HasBinanceInvoiceId returns a boolean if a field has been set.

### GetBinanceQrcode

`func (o *MiniInvoice) GetBinanceQrcode() string`

GetBinanceQrcode returns the BinanceQrcode field if non-nil, zero value otherwise.

### GetBinanceQrcodeOk

`func (o *MiniInvoice) GetBinanceQrcodeOk() (*string, bool)`

GetBinanceQrcodeOk returns a tuple with the BinanceQrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceQrcode

`func (o *MiniInvoice) SetBinanceQrcode(v string)`

SetBinanceQrcode sets BinanceQrcode field to given value.

### HasBinanceQrcode

`func (o *MiniInvoice) HasBinanceQrcode() bool`

HasBinanceQrcode returns a boolean if a field has been set.

### GetBinanceCheckoutUrl

`func (o *MiniInvoice) GetBinanceCheckoutUrl() string`

GetBinanceCheckoutUrl returns the BinanceCheckoutUrl field if non-nil, zero value otherwise.

### GetBinanceCheckoutUrlOk

`func (o *MiniInvoice) GetBinanceCheckoutUrlOk() (*string, bool)`

GetBinanceCheckoutUrlOk returns a tuple with the BinanceCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceCheckoutUrl

`func (o *MiniInvoice) SetBinanceCheckoutUrl(v string)`

SetBinanceCheckoutUrl sets BinanceCheckoutUrl field to given value.

### HasBinanceCheckoutUrl

`func (o *MiniInvoice) HasBinanceCheckoutUrl() bool`

HasBinanceCheckoutUrl returns a boolean if a field has been set.

### GetBinancePayout

`func (o *MiniInvoice) GetBinancePayout() int32`

GetBinancePayout returns the BinancePayout field if non-nil, zero value otherwise.

### GetBinancePayoutOk

`func (o *MiniInvoice) GetBinancePayoutOk() (*int32, bool)`

GetBinancePayoutOk returns a tuple with the BinancePayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinancePayout

`func (o *MiniInvoice) SetBinancePayout(v int32)`

SetBinancePayout sets BinancePayout field to given value.

### HasBinancePayout

`func (o *MiniInvoice) HasBinancePayout() bool`

HasBinancePayout returns a boolean if a field has been set.

### GetCashappQrcode

`func (o *MiniInvoice) GetCashappQrcode() string`

GetCashappQrcode returns the CashappQrcode field if non-nil, zero value otherwise.

### GetCashappQrcodeOk

`func (o *MiniInvoice) GetCashappQrcodeOk() (*string, bool)`

GetCashappQrcodeOk returns a tuple with the CashappQrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappQrcode

`func (o *MiniInvoice) SetCashappQrcode(v string)`

SetCashappQrcode sets CashappQrcode field to given value.

### HasCashappQrcode

`func (o *MiniInvoice) HasCashappQrcode() bool`

HasCashappQrcode returns a boolean if a field has been set.

### GetCashappNote

`func (o *MiniInvoice) GetCashappNote() string`

GetCashappNote returns the CashappNote field if non-nil, zero value otherwise.

### GetCashappNoteOk

`func (o *MiniInvoice) GetCashappNoteOk() (*string, bool)`

GetCashappNoteOk returns a tuple with the CashappNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappNote

`func (o *MiniInvoice) SetCashappNote(v string)`

SetCashappNote sets CashappNote field to given value.

### HasCashappNote

`func (o *MiniInvoice) HasCashappNote() bool`

HasCashappNote returns a boolean if a field has been set.

### GetCashappCashtag

`func (o *MiniInvoice) GetCashappCashtag() string`

GetCashappCashtag returns the CashappCashtag field if non-nil, zero value otherwise.

### GetCashappCashtagOk

`func (o *MiniInvoice) GetCashappCashtagOk() (*string, bool)`

GetCashappCashtagOk returns a tuple with the CashappCashtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappCashtag

`func (o *MiniInvoice) SetCashappCashtag(v string)`

SetCashappCashtag sets CashappCashtag field to given value.

### HasCashappCashtag

`func (o *MiniInvoice) HasCashappCashtag() bool`

HasCashappCashtag returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *MiniInvoice) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *MiniInvoice) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *MiniInvoice) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *MiniInvoice) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *MiniInvoice) GetCryptoAmount() float64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *MiniInvoice) GetCryptoAmountOk() (*float64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *MiniInvoice) SetCryptoAmount(v float64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *MiniInvoice) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoReceived

`func (o *MiniInvoice) GetCryptoReceived() float64`

GetCryptoReceived returns the CryptoReceived field if non-nil, zero value otherwise.

### GetCryptoReceivedOk

`func (o *MiniInvoice) GetCryptoReceivedOk() (*float64, bool)`

GetCryptoReceivedOk returns a tuple with the CryptoReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoReceived

`func (o *MiniInvoice) SetCryptoReceived(v float64)`

SetCryptoReceived sets CryptoReceived field to given value.

### HasCryptoReceived

`func (o *MiniInvoice) HasCryptoReceived() bool`

HasCryptoReceived returns a boolean if a field has been set.

### GetCryptoFeePaid

`func (o *MiniInvoice) GetCryptoFeePaid() float32`

GetCryptoFeePaid returns the CryptoFeePaid field if non-nil, zero value otherwise.

### GetCryptoFeePaidOk

`func (o *MiniInvoice) GetCryptoFeePaidOk() (*float32, bool)`

GetCryptoFeePaidOk returns a tuple with the CryptoFeePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoFeePaid

`func (o *MiniInvoice) SetCryptoFeePaid(v float32)`

SetCryptoFeePaid sets CryptoFeePaid field to given value.

### HasCryptoFeePaid

`func (o *MiniInvoice) HasCryptoFeePaid() bool`

HasCryptoFeePaid returns a boolean if a field has been set.

### GetCryptoFeeOwed

`func (o *MiniInvoice) GetCryptoFeeOwed() float32`

GetCryptoFeeOwed returns the CryptoFeeOwed field if non-nil, zero value otherwise.

### GetCryptoFeeOwedOk

`func (o *MiniInvoice) GetCryptoFeeOwedOk() (*float32, bool)`

GetCryptoFeeOwedOk returns a tuple with the CryptoFeeOwed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoFeeOwed

`func (o *MiniInvoice) SetCryptoFeeOwed(v float32)`

SetCryptoFeeOwed sets CryptoFeeOwed field to given value.

### HasCryptoFeeOwed

`func (o *MiniInvoice) HasCryptoFeeOwed() bool`

HasCryptoFeeOwed returns a boolean if a field has been set.

### GetCryptoUri

`func (o *MiniInvoice) GetCryptoUri() string`

GetCryptoUri returns the CryptoUri field if non-nil, zero value otherwise.

### GetCryptoUriOk

`func (o *MiniInvoice) GetCryptoUriOk() (*string, bool)`

GetCryptoUriOk returns a tuple with the CryptoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoUri

`func (o *MiniInvoice) SetCryptoUri(v string)`

SetCryptoUri sets CryptoUri field to given value.

### HasCryptoUri

`func (o *MiniInvoice) HasCryptoUri() bool`

HasCryptoUri returns a boolean if a field has been set.

### GetBitcoinLnRHash

`func (o *MiniInvoice) GetBitcoinLnRHash() string`

GetBitcoinLnRHash returns the BitcoinLnRHash field if non-nil, zero value otherwise.

### GetBitcoinLnRHashOk

`func (o *MiniInvoice) GetBitcoinLnRHashOk() (*string, bool)`

GetBitcoinLnRHashOk returns a tuple with the BitcoinLnRHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitcoinLnRHash

`func (o *MiniInvoice) SetBitcoinLnRHash(v string)`

SetBitcoinLnRHash sets BitcoinLnRHash field to given value.

### HasBitcoinLnRHash

`func (o *MiniInvoice) HasBitcoinLnRHash() bool`

HasBitcoinLnRHash returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *MiniInvoice) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *MiniInvoice) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *MiniInvoice) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *MiniInvoice) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetCryptoWalletVersion

`func (o *MiniInvoice) GetCryptoWalletVersion() string`

GetCryptoWalletVersion returns the CryptoWalletVersion field if non-nil, zero value otherwise.

### GetCryptoWalletVersionOk

`func (o *MiniInvoice) GetCryptoWalletVersionOk() (*string, bool)`

GetCryptoWalletVersionOk returns a tuple with the CryptoWalletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoWalletVersion

`func (o *MiniInvoice) SetCryptoWalletVersion(v string)`

SetCryptoWalletVersion sets CryptoWalletVersion field to given value.

### HasCryptoWalletVersion

`func (o *MiniInvoice) HasCryptoWalletVersion() bool`

HasCryptoWalletVersion returns a boolean if a field has been set.

### GetCryptoPayout

`func (o *MiniInvoice) GetCryptoPayout() bool`

GetCryptoPayout returns the CryptoPayout field if non-nil, zero value otherwise.

### GetCryptoPayoutOk

`func (o *MiniInvoice) GetCryptoPayoutOk() (*bool, bool)`

GetCryptoPayoutOk returns a tuple with the CryptoPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPayout

`func (o *MiniInvoice) SetCryptoPayout(v bool)`

SetCryptoPayout sets CryptoPayout field to given value.

### HasCryptoPayout

`func (o *MiniInvoice) HasCryptoPayout() bool`

HasCryptoPayout returns a boolean if a field has been set.

### GetCryptoScheduledPayout

`func (o *MiniInvoice) GetCryptoScheduledPayout() bool`

GetCryptoScheduledPayout returns the CryptoScheduledPayout field if non-nil, zero value otherwise.

### GetCryptoScheduledPayoutOk

`func (o *MiniInvoice) GetCryptoScheduledPayoutOk() (*bool, bool)`

GetCryptoScheduledPayoutOk returns a tuple with the CryptoScheduledPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoScheduledPayout

`func (o *MiniInvoice) SetCryptoScheduledPayout(v bool)`

SetCryptoScheduledPayout sets CryptoScheduledPayout field to given value.

### HasCryptoScheduledPayout

`func (o *MiniInvoice) HasCryptoScheduledPayout() bool`

HasCryptoScheduledPayout returns a boolean if a field has been set.

### GetCryptoPayoutType

`func (o *MiniInvoice) GetCryptoPayoutType() string`

GetCryptoPayoutType returns the CryptoPayoutType field if non-nil, zero value otherwise.

### GetCryptoPayoutTypeOk

`func (o *MiniInvoice) GetCryptoPayoutTypeOk() (*string, bool)`

GetCryptoPayoutTypeOk returns a tuple with the CryptoPayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPayoutType

`func (o *MiniInvoice) SetCryptoPayoutType(v string)`

SetCryptoPayoutType sets CryptoPayoutType field to given value.

### HasCryptoPayoutType

`func (o *MiniInvoice) HasCryptoPayoutType() bool`

HasCryptoPayoutType returns a boolean if a field has been set.

### GetFeeBilled

`func (o *MiniInvoice) GetFeeBilled() bool`

GetFeeBilled returns the FeeBilled field if non-nil, zero value otherwise.

### GetFeeBilledOk

`func (o *MiniInvoice) GetFeeBilledOk() (*bool, bool)`

GetFeeBilledOk returns a tuple with the FeeBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBilled

`func (o *MiniInvoice) SetFeeBilled(v bool)`

SetFeeBilled sets FeeBilled field to given value.

### HasFeeBilled

`func (o *MiniInvoice) HasFeeBilled() bool`

HasFeeBilled returns a boolean if a field has been set.

### GetBillInfo

`func (o *MiniInvoice) GetBillInfo() map[string]interface{}`

GetBillInfo returns the BillInfo field if non-nil, zero value otherwise.

### GetBillInfoOk

`func (o *MiniInvoice) GetBillInfoOk() (*map[string]interface{}, bool)`

GetBillInfoOk returns a tuple with the BillInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillInfo

`func (o *MiniInvoice) SetBillInfo(v map[string]interface{})`

SetBillInfo sets BillInfo field to given value.

### HasBillInfo

`func (o *MiniInvoice) HasBillInfo() bool`

HasBillInfo returns a boolean if a field has been set.

### GetCountry

`func (o *MiniInvoice) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MiniInvoice) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MiniInvoice) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MiniInvoice) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocation

`func (o *MiniInvoice) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MiniInvoice) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MiniInvoice) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MiniInvoice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIp

`func (o *MiniInvoice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MiniInvoice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MiniInvoice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MiniInvoice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsVpnOrProxy

`func (o *MiniInvoice) GetIsVpnOrProxy() bool`

GetIsVpnOrProxy returns the IsVpnOrProxy field if non-nil, zero value otherwise.

### GetIsVpnOrProxyOk

`func (o *MiniInvoice) GetIsVpnOrProxyOk() (*bool, bool)`

GetIsVpnOrProxyOk returns a tuple with the IsVpnOrProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpnOrProxy

`func (o *MiniInvoice) SetIsVpnOrProxy(v bool)`

SetIsVpnOrProxy sets IsVpnOrProxy field to given value.

### HasIsVpnOrProxy

`func (o *MiniInvoice) HasIsVpnOrProxy() bool`

HasIsVpnOrProxy returns a boolean if a field has been set.

### GetUserAgent

`func (o *MiniInvoice) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *MiniInvoice) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *MiniInvoice) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *MiniInvoice) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCustomFields

`func (o *MiniInvoice) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MiniInvoice) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MiniInvoice) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MiniInvoice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDeveloperInvoice

`func (o *MiniInvoice) GetDeveloperInvoice() bool`

GetDeveloperInvoice returns the DeveloperInvoice field if non-nil, zero value otherwise.

### GetDeveloperInvoiceOk

`func (o *MiniInvoice) GetDeveloperInvoiceOk() (*bool, bool)`

GetDeveloperInvoiceOk returns a tuple with the DeveloperInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperInvoice

`func (o *MiniInvoice) SetDeveloperInvoice(v bool)`

SetDeveloperInvoice sets DeveloperInvoice field to given value.

### HasDeveloperInvoice

`func (o *MiniInvoice) HasDeveloperInvoice() bool`

HasDeveloperInvoice returns a boolean if a field has been set.

### GetDeveloperTitle

`func (o *MiniInvoice) GetDeveloperTitle() string`

GetDeveloperTitle returns the DeveloperTitle field if non-nil, zero value otherwise.

### GetDeveloperTitleOk

`func (o *MiniInvoice) GetDeveloperTitleOk() (*string, bool)`

GetDeveloperTitleOk returns a tuple with the DeveloperTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperTitle

`func (o *MiniInvoice) SetDeveloperTitle(v string)`

SetDeveloperTitle sets DeveloperTitle field to given value.

### HasDeveloperTitle

`func (o *MiniInvoice) HasDeveloperTitle() bool`

HasDeveloperTitle returns a boolean if a field has been set.

### GetDeveloperWebhook

`func (o *MiniInvoice) GetDeveloperWebhook() string`

GetDeveloperWebhook returns the DeveloperWebhook field if non-nil, zero value otherwise.

### GetDeveloperWebhookOk

`func (o *MiniInvoice) GetDeveloperWebhookOk() (*string, bool)`

GetDeveloperWebhookOk returns a tuple with the DeveloperWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperWebhook

`func (o *MiniInvoice) SetDeveloperWebhook(v string)`

SetDeveloperWebhook sets DeveloperWebhook field to given value.

### HasDeveloperWebhook

`func (o *MiniInvoice) HasDeveloperWebhook() bool`

HasDeveloperWebhook returns a boolean if a field has been set.

### GetDeveloperReturnUrl

`func (o *MiniInvoice) GetDeveloperReturnUrl() string`

GetDeveloperReturnUrl returns the DeveloperReturnUrl field if non-nil, zero value otherwise.

### GetDeveloperReturnUrlOk

`func (o *MiniInvoice) GetDeveloperReturnUrlOk() (*string, bool)`

GetDeveloperReturnUrlOk returns a tuple with the DeveloperReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperReturnUrl

`func (o *MiniInvoice) SetDeveloperReturnUrl(v string)`

SetDeveloperReturnUrl sets DeveloperReturnUrl field to given value.

### HasDeveloperReturnUrl

`func (o *MiniInvoice) HasDeveloperReturnUrl() bool`

HasDeveloperReturnUrl returns a boolean if a field has been set.

### GetPayoutConfiguration

`func (o *MiniInvoice) GetPayoutConfiguration() string`

GetPayoutConfiguration returns the PayoutConfiguration field if non-nil, zero value otherwise.

### GetPayoutConfigurationOk

`func (o *MiniInvoice) GetPayoutConfigurationOk() (*string, bool)`

GetPayoutConfigurationOk returns a tuple with the PayoutConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutConfiguration

`func (o *MiniInvoice) SetPayoutConfiguration(v string)`

SetPayoutConfiguration sets PayoutConfiguration field to given value.

### HasPayoutConfiguration

`func (o *MiniInvoice) HasPayoutConfiguration() bool`

HasPayoutConfiguration returns a boolean if a field has been set.

### GetFeePercentage

`func (o *MiniInvoice) GetFeePercentage() int32`

GetFeePercentage returns the FeePercentage field if non-nil, zero value otherwise.

### GetFeePercentageOk

`func (o *MiniInvoice) GetFeePercentageOk() (*int32, bool)`

GetFeePercentageOk returns a tuple with the FeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercentage

`func (o *MiniInvoice) SetFeePercentage(v int32)`

SetFeePercentage sets FeePercentage field to given value.

### HasFeePercentage

`func (o *MiniInvoice) HasFeePercentage() bool`

HasFeePercentage returns a boolean if a field has been set.

### GetFeeBreakdown

`func (o *MiniInvoice) GetFeeBreakdown() FeeBreakdown`

GetFeeBreakdown returns the FeeBreakdown field if non-nil, zero value otherwise.

### GetFeeBreakdownOk

`func (o *MiniInvoice) GetFeeBreakdownOk() (*FeeBreakdown, bool)`

GetFeeBreakdownOk returns a tuple with the FeeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBreakdown

`func (o *MiniInvoice) SetFeeBreakdown(v FeeBreakdown)`

SetFeeBreakdown sets FeeBreakdown field to given value.

### HasFeeBreakdown

`func (o *MiniInvoice) HasFeeBreakdown() bool`

HasFeeBreakdown returns a boolean if a field has been set.

### GetToProcess

`func (o *MiniInvoice) GetToProcess() float32`

GetToProcess returns the ToProcess field if non-nil, zero value otherwise.

### GetToProcessOk

`func (o *MiniInvoice) GetToProcessOk() (*float32, bool)`

GetToProcessOk returns a tuple with the ToProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToProcess

`func (o *MiniInvoice) SetToProcess(v float32)`

SetToProcess sets ToProcess field to given value.

### HasToProcess

`func (o *MiniInvoice) HasToProcess() bool`

HasToProcess returns a boolean if a field has been set.

### GetStatus

`func (o *MiniInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MiniInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MiniInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MiniInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *MiniInvoice) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *MiniInvoice) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *MiniInvoice) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *MiniInvoice) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetVoidDetails

`func (o *MiniInvoice) GetVoidDetails() string`

GetVoidDetails returns the VoidDetails field if non-nil, zero value otherwise.

### GetVoidDetailsOk

`func (o *MiniInvoice) GetVoidDetailsOk() (*string, bool)`

GetVoidDetailsOk returns a tuple with the VoidDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidDetails

`func (o *MiniInvoice) SetVoidDetails(v string)`

SetVoidDetails sets VoidDetails field to given value.

### HasVoidDetails

`func (o *MiniInvoice) HasVoidDetails() bool`

HasVoidDetails returns a boolean if a field has been set.

### GetEnvironment

`func (o *MiniInvoice) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MiniInvoice) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MiniInvoice) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MiniInvoice) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetDayValue

`func (o *MiniInvoice) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *MiniInvoice) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *MiniInvoice) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *MiniInvoice) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetDay

`func (o *MiniInvoice) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MiniInvoice) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MiniInvoice) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *MiniInvoice) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetMonth

`func (o *MiniInvoice) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MiniInvoice) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MiniInvoice) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *MiniInvoice) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *MiniInvoice) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MiniInvoice) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MiniInvoice) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MiniInvoice) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MiniInvoice) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MiniInvoice) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MiniInvoice) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MiniInvoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGatewayUpdatedAt

`func (o *MiniInvoice) GetGatewayUpdatedAt() int32`

GetGatewayUpdatedAt returns the GatewayUpdatedAt field if non-nil, zero value otherwise.

### GetGatewayUpdatedAtOk

`func (o *MiniInvoice) GetGatewayUpdatedAtOk() (*int32, bool)`

GetGatewayUpdatedAtOk returns a tuple with the GatewayUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUpdatedAt

`func (o *MiniInvoice) SetGatewayUpdatedAt(v int32)`

SetGatewayUpdatedAt sets GatewayUpdatedAt field to given value.

### HasGatewayUpdatedAt

`func (o *MiniInvoice) HasGatewayUpdatedAt() bool`

HasGatewayUpdatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MiniInvoice) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MiniInvoice) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MiniInvoice) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MiniInvoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *MiniInvoice) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *MiniInvoice) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *MiniInvoice) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *MiniInvoice) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetIsMarketplace

`func (o *MiniInvoice) GetIsMarketplace() int32`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *MiniInvoice) GetIsMarketplaceOk() (*int32, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *MiniInvoice) SetIsMarketplace(v int32)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *MiniInvoice) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### GetSquarePaymentId

`func (o *MiniInvoice) GetSquarePaymentId() string`

GetSquarePaymentId returns the SquarePaymentId field if non-nil, zero value otherwise.

### GetSquarePaymentIdOk

`func (o *MiniInvoice) GetSquarePaymentIdOk() (*string, bool)`

GetSquarePaymentIdOk returns a tuple with the SquarePaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarePaymentId

`func (o *MiniInvoice) SetSquarePaymentId(v string)`

SetSquarePaymentId sets SquarePaymentId field to given value.

### HasSquarePaymentId

`func (o *MiniInvoice) HasSquarePaymentId() bool`

HasSquarePaymentId returns a boolean if a field has been set.

### GetVolumePaymentId

`func (o *MiniInvoice) GetVolumePaymentId() string`

GetVolumePaymentId returns the VolumePaymentId field if non-nil, zero value otherwise.

### GetVolumePaymentIdOk

`func (o *MiniInvoice) GetVolumePaymentIdOk() (*string, bool)`

GetVolumePaymentIdOk returns a tuple with the VolumePaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePaymentId

`func (o *MiniInvoice) SetVolumePaymentId(v string)`

SetVolumePaymentId sets VolumePaymentId field to given value.

### HasVolumePaymentId

`func (o *MiniInvoice) HasVolumePaymentId() bool`

HasVolumePaymentId returns a boolean if a field has been set.

### GetVolumeReferenceId

`func (o *MiniInvoice) GetVolumeReferenceId() string`

GetVolumeReferenceId returns the VolumeReferenceId field if non-nil, zero value otherwise.

### GetVolumeReferenceIdOk

`func (o *MiniInvoice) GetVolumeReferenceIdOk() (*string, bool)`

GetVolumeReferenceIdOk returns a tuple with the VolumeReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeReferenceId

`func (o *MiniInvoice) SetVolumeReferenceId(v string)`

SetVolumeReferenceId sets VolumeReferenceId field to given value.

### HasVolumeReferenceId

`func (o *MiniInvoice) HasVolumeReferenceId() bool`

HasVolumeReferenceId returns a boolean if a field has been set.

### GetProductPlanId

`func (o *MiniInvoice) GetProductPlanId() int32`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *MiniInvoice) GetProductPlanIdOk() (*int32, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *MiniInvoice) SetProductPlanId(v int32)`

SetProductPlanId sets ProductPlanId field to given value.

### HasProductPlanId

`func (o *MiniInvoice) HasProductPlanId() bool`

HasProductPlanId returns a boolean if a field has been set.

### GetName

`func (o *MiniInvoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MiniInvoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MiniInvoice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MiniInvoice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShopPaypalCreditCard

`func (o *MiniInvoice) GetShopPaypalCreditCard() bool`

GetShopPaypalCreditCard returns the ShopPaypalCreditCard field if non-nil, zero value otherwise.

### GetShopPaypalCreditCardOk

`func (o *MiniInvoice) GetShopPaypalCreditCardOk() (*bool, bool)`

GetShopPaypalCreditCardOk returns a tuple with the ShopPaypalCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopPaypalCreditCard

`func (o *MiniInvoice) SetShopPaypalCreditCard(v bool)`

SetShopPaypalCreditCard sets ShopPaypalCreditCard field to given value.

### HasShopPaypalCreditCard

`func (o *MiniInvoice) HasShopPaypalCreditCard() bool`

HasShopPaypalCreditCard returns a boolean if a field has been set.

### GetShopForcePaypalEmailDelivery

`func (o *MiniInvoice) GetShopForcePaypalEmailDelivery() bool`

GetShopForcePaypalEmailDelivery returns the ShopForcePaypalEmailDelivery field if non-nil, zero value otherwise.

### GetShopForcePaypalEmailDeliveryOk

`func (o *MiniInvoice) GetShopForcePaypalEmailDeliveryOk() (*bool, bool)`

GetShopForcePaypalEmailDeliveryOk returns a tuple with the ShopForcePaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopForcePaypalEmailDelivery

`func (o *MiniInvoice) SetShopForcePaypalEmailDelivery(v bool)`

SetShopForcePaypalEmailDelivery sets ShopForcePaypalEmailDelivery field to given value.

### HasShopForcePaypalEmailDelivery

`func (o *MiniInvoice) HasShopForcePaypalEmailDelivery() bool`

HasShopForcePaypalEmailDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


