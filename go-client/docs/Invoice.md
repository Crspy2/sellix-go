# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**RecurringBillingId** | Pointer to **string** | Unique ID of the recurring bill. | [optional] 
**PayoutConfiguration** | Pointer to **string** | DEPRECATED | [optional] 
**Secret** | Pointer to **string** | An internal Sellix secret used for completing tasks with our internal API | [optional] 
**Type** | Pointer to **string** | Invoice type. For subscriptions, the invoice type is PRODUCT_SUBSCRIPTION as SUBSCRIPTION is reserved for Sellix-own plans. | [optional] 
**Subtype** | Pointer to **string** | Product type. | [optional] 
**Origin** | Pointer to **string** | How the invoice was created | [optional] 
**Total** | Pointer to **float64** | Invoice total, unit_price*quantity where quantity is applicable. | [optional] 
**TotalDisplay** | Pointer to **float64** | Invoice total in the currency chosen. | [optional] 
**ProductVariants** | Pointer to [**[]ProductVariant**](ProductVariant.md) |  | [optional] 
**ExchangeRate** | Pointer to **float64** | Exchange rate between currency chosen and USD. | [optional] 
**CryptoExchangeRate** | Pointer to **float64** | Exchange rate between the cryptocurrency chosen (if any) and USD. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this invoice belongs. | [optional] 
**ShopImageName** | Pointer to **string** | DEPRECATED: Unique id of the image attachment for this merchant with the image extension included. | [optional] 
**ShopImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**Name** | Pointer to **string** | Name of the merchant. | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer. | [optional] 
**CustomerId** | Pointer to **string** | ID of the customer. | [optional] 
**AffliateRevenueCustomerId** | Pointer to **string** | ID of the affiliate | [optional] 
**PaypalEmailDelivery** | Pointer to **bool** | If true and gateway is PAYPAL, the product will be shipped to the PayPal email on record instead of the customer email. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product linked to this invoice, if any. | [optional] 
**ProductTitle** | Pointer to **string** | Product title. | [optional] 
**ProductType** | Pointer to **string** | Product type. | [optional] 
**SubscriptionId** | Pointer to **int32** | Field reserved for Sellix-own plans. Unique ID of the subscription purchased. | [optional] 
**SubscriptionTime** | Pointer to **int32** | Field reserved for Sellix-own plans. Time, in seconds, of the subscription purchased. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**Blockchain** | Pointer to [**Blockchain**](Blockchain.md) |  | [optional] 
**PaypalApm** | Pointer to **string** | PayPal Alternative Payment Method name, such as iDEAL, used if gateway is PAYPAL. | [optional] 
**StripeApm** | Pointer to **string** | Stripe Alternative Payment Method name, such as iDEAL, used if gateway is STRIPE. | [optional] 
**PaypalEmail** | Pointer to **string** | Merchant PayPal email. | [optional] 
**PaypalOrderId** | Pointer to **string** | This field contains the PayPal order ID. | [optional] 
**PaypalPayerEmail** | Pointer to **string** | This field is updated after the invoice is completed with the PayPal&#39;s email used for the purchase. | [optional] 
**PaypalFee** | Pointer to **string** | This field is updated after the invoice is completed with the fee taken by PayPal over the invoice total. | [optional] 
**PaypalSubscriptionId** | Pointer to **int32** | ID of the paypal subscription. | [optional] 
**PaypalSubscriptionLink** | Pointer to **int32** | Link for the merchant to purchase the subscription from. | [optional] 
**LexOrderId** | Pointer to **string** | DEPRECATED | [optional] 
**LexPaymentMethod** | Pointer to **string** | DEPRECATED | [optional] 
**PaydashPaymentID** | Pointer to **string** | DEPRECATED | [optional] 
**VirtualPaymentsId** | Pointer to **string** | DEPRECATED | [optional] 
**StripeClientSecret** | Pointer to **string** | Client secret used to create the STRIPE paymentIntent. | [optional] 
**StripePriceId** | Pointer to **string** | If the invoice type is PRODUCT_SUBSCRIPTION or SUBSCRIPTION, refers to the STRIPE price ID. | [optional] 
**SkrillEmail** | Pointer to **string** | Merchant Skrill email. | [optional] 
**SkrillSid** | Pointer to **string** | Skrill unique ID linked to the invoice. | [optional] 
**SkrillLink** | Pointer to **string** | Skrill link to redirect the customer to. | [optional] 
**PerfectmoneyId** | Pointer to **string** | PerfectMoney payment ID linked to the invoice. | [optional] 
**BinanceInvoiceId** | Pointer to **string** | ID for binance invoice | [optional] 
**BinanceQrcode** | Pointer to **string** | Full Binance QRCODE string | [optional] 
**BinanceCheckoutUrl** | Pointer to **string** | Checkout URL for Binance invoice | [optional] 
**CryptoAddress** | Pointer to **string** | Cryptocurrency address linked to this invoice. | [optional] 
**CryptoAmount** | Pointer to **float64** | Cryptocurrency amount converted based on crypto_exchange_rate. | [optional] 
**CryptoReceived** | Pointer to **float64** | Cryptocurrency amount received, paid by the customer. | [optional] 
**CryptoUri** | Pointer to **string** | URI used to create the QRCODE. | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Crypto confirmations needed to process the invoice. | [optional] 
**CryptoScheduledPayout** | Pointer to **bool** | If true, a scheduled payout for this invoice&#39;s cryptocurrency address has been sent. | [optional] 
**CryptoPayout** | Pointer to **bool** | If true, an instant payout for this invoice&#39;s cryptocurrency address has been sent. | [optional] 
**FeeBilled** | Pointer to **bool** | If true, the Sellix fee_percentage has already been collected. | [optional] 
**BillInfo** | Pointer to **map[string]interface{}** | If invoice type is MONTHLY_BILL, contains a breakdown of the fees that need to be collected. | [optional] 
**CashappQrcode** | Pointer to **string** | Full CashApp QRCODE string. | [optional] 
**CashappNote** | Pointer to **string** | Unique note for the customer to leave in the CashApp payment. | [optional] 
**CashappCashtag** | Pointer to **string** | CashApp cashtag used to create the QRCODE. | [optional] 
**Country** | Pointer to **string** | Customer country. | [optional] 
**Location** | Pointer to **string** | Customer location. | [optional] 
**Ip** | Pointer to **string** | Customer IP. | [optional] 
**IsVpnOrProxy** | Pointer to **bool** | If true, a VPN or Proxy has been detected. | [optional] 
**UserAgent** | Pointer to **string** | Customer User Agent. | [optional] 
**Quantity** | Pointer to **int32** | Qauntity of product purchased. | [optional] 
**CouponId** | Pointer to **string** | Unique ID of the coupon, if used, for the discount. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**DeveloperInvoice** | Pointer to **bool** | If true, this invoice has been created through the Developers API. | [optional] 
**DeveloperTitle** | Pointer to **string** | If a product_id is not passed through the Developers API, a title must be specified. | [optional] 
**DeveloperWebhook** | Pointer to **string** | Additional webhook URL to which updates regarding this invoice will be sent. | [optional] 
**DeveloperReturnUrl** | Pointer to **string** | If present, the customer will be redirected to this URL after the invoice has been paid. | [optional] 
**Status** | Pointer to **string** | Status of the invoice. | [optional] 
**StatusDetails** | Pointer to **string** | If CART_PARTIAL_OUT_OF_STOCK, the invoice has been completed but some products in the cart were out of stock. | [optional] 
**VoidDetails** | Pointer to **string** | If the invoice is VOIDED and the product (or all the products in the cart) were out of stock, this additional status is set. | [optional] 
**Discount** | Pointer to **float64** | If a coupon or volume_discount is used, the discount value presents the total amount of discount over the total cost of the invoice. | [optional] 
**FeePercentage** | Pointer to **int32** | What cut does Sellix take out of the total. To learn more about Sellix fees please refer to https://sellix.io/fees. | [optional] 
**FeeBreakdown** | Pointer to [**FeeBreakdown**](FeeBreakdown.md) |  | [optional] 
**DiscountBreakdown** | Pointer to **map[string]interface{}** | Representation of how discount was applied to invoice | [optional] 
**DayValue** | Pointer to **int32** | DEPRECATED: Day value, number. | [optional] 
**Day** | Pointer to **string** | DEPRECATED: First three letters of the day name. | [optional] 
**Month** | Pointer to **string** | DEPRECATED: First three letters of the month name. | [optional] 
**Year** | Pointer to **int32** | DEPRECATED: Year number. | [optional] 
**ProductAddons** | Pointer to [**[]ProductAddonsInner**](ProductAddonsInner.md) |  | [optional] 
**BundleConfig** | Pointer to [**[]BundleConfig**](BundleConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the invoice. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the blacklist has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the invoice. | [optional] 
**ApprovedAddress** | Pointer to [**ApprovedAddress**](ApprovedAddress.md) |  | [optional] 
**ServiceText** | Pointer to **string** | If the product type is SERVICE, this field contains additional details on the type of service provided by the merchant. | [optional] 
**IpInfo** | Pointer to [**IpInfo**](IpInfo.md) |  | [optional] 
**Webhooks** | Pointer to [**[]Webhook**](Webhook.md) | Webhook responses for this invoice. | [optional] 
**RewardsData** | Pointer to [**[]InvoiceRewardsDataInner**](InvoiceRewardsDataInner.md) |  | [optional] 
**PaypalDispute** | Pointer to [**PaypalDispute**](PaypalDispute.md) |  | [optional] 
**ProductDownloads** | Pointer to [**[]ProductDownload**](ProductDownload.md) |  | [optional] 
**PaymentLinkId** | Pointer to **string** | The ID for the payment link, if invoice was created by a payment link | [optional] 
**CashappEmailConfigured** | Pointer to **bool** | Whether or not cashapp email is configured | [optional] 
**License** | Pointer to **bool** | True if the product is of type LICENSE | [optional] 
**StatusHistory** | Pointer to [**[]InvoiceStatus**](InvoiceStatus.md) | Additional details on the invoice status change. | [optional] 
**AmlWallets** | Pointer to [**[]AmlWallet**](AmlWallet.md) |  | [optional] 
**CryptoTransactions** | Pointer to [**[]CryptoTransaction**](CryptoTransaction.md) | Crypto transactions received to fulfill this invoice. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**TotalConversions** | Pointer to **interface{}** |  | [optional] 
**Theme** | Pointer to **string** | What Sellix theme the shop is set to. | [optional] 
**DarkMode** | Pointer to **int32** | 1 if darkmode is enabled, 0 if it is disabled. | [optional] 
**CryptoMode** | Pointer to **interface{}** |  | [optional] 
**GatewaysAvailable** | Pointer to **[]string** | gateways available for the update invoice Developers API. | [optional] 
**CountryRegulations** | Pointer to **string** | The country the shop is opperated in | [optional] 
**AvailableStripeApm** | Pointer to [**InvoiceAvailableStripeApm**](InvoiceAvailableStripeApm.md) |  | [optional] 
**Serials** | Pointer to **[]string** | If product type is SERIALS, provide the serials linked to this invoice. | [optional] 
**ShopPaymentGatewaysFees** | Pointer to [**[]GatewayFees**](GatewayFees.md) |  | [optional] 
**ShopPaypalCreditCard** | Pointer to **bool** | If true, the merchant has enabled purchase with Credit Card through PayPal. | [optional] 
**ShopForcePaypalEmailDelivery** | Pointer to **bool** | If true, the merchant has enabled invoice shipment to the PayPal customer email. | [optional] 
**ShopWalletconnectId** | Pointer to **interface{}** |  | [optional] 
**OriginalDeveloperReturnUrl** | Pointer to **string** | The original return url of the order. | [optional] 
**RatesSnapshot** | Pointer to [**InvoiceRatesSnapshot**](InvoiceRatesSnapshot.md) |  | [optional] 
**VoidTimes** | Pointer to [**[]InvoiceVoidTimesInner**](InvoiceVoidTimesInner.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Invoice) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Invoice) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Invoice) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Invoice) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetRecurringBillingId

`func (o *Invoice) GetRecurringBillingId() string`

GetRecurringBillingId returns the RecurringBillingId field if non-nil, zero value otherwise.

### GetRecurringBillingIdOk

`func (o *Invoice) GetRecurringBillingIdOk() (*string, bool)`

GetRecurringBillingIdOk returns a tuple with the RecurringBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringBillingId

`func (o *Invoice) SetRecurringBillingId(v string)`

SetRecurringBillingId sets RecurringBillingId field to given value.

### HasRecurringBillingId

`func (o *Invoice) HasRecurringBillingId() bool`

HasRecurringBillingId returns a boolean if a field has been set.

### GetPayoutConfiguration

`func (o *Invoice) GetPayoutConfiguration() string`

GetPayoutConfiguration returns the PayoutConfiguration field if non-nil, zero value otherwise.

### GetPayoutConfigurationOk

`func (o *Invoice) GetPayoutConfigurationOk() (*string, bool)`

GetPayoutConfigurationOk returns a tuple with the PayoutConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutConfiguration

`func (o *Invoice) SetPayoutConfiguration(v string)`

SetPayoutConfiguration sets PayoutConfiguration field to given value.

### HasPayoutConfiguration

`func (o *Invoice) HasPayoutConfiguration() bool`

HasPayoutConfiguration returns a boolean if a field has been set.

### GetSecret

`func (o *Invoice) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Invoice) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Invoice) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Invoice) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Invoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *Invoice) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Invoice) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Invoice) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Invoice) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetOrigin

`func (o *Invoice) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Invoice) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Invoice) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Invoice) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalDisplay

`func (o *Invoice) GetTotalDisplay() float64`

GetTotalDisplay returns the TotalDisplay field if non-nil, zero value otherwise.

### GetTotalDisplayOk

`func (o *Invoice) GetTotalDisplayOk() (*float64, bool)`

GetTotalDisplayOk returns a tuple with the TotalDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisplay

`func (o *Invoice) SetTotalDisplay(v float64)`

SetTotalDisplay sets TotalDisplay field to given value.

### HasTotalDisplay

`func (o *Invoice) HasTotalDisplay() bool`

HasTotalDisplay returns a boolean if a field has been set.

### GetProductVariants

`func (o *Invoice) GetProductVariants() []ProductVariant`

GetProductVariants returns the ProductVariants field if non-nil, zero value otherwise.

### GetProductVariantsOk

`func (o *Invoice) GetProductVariantsOk() (*[]ProductVariant, bool)`

GetProductVariantsOk returns a tuple with the ProductVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariants

`func (o *Invoice) SetProductVariants(v []ProductVariant)`

SetProductVariants sets ProductVariants field to given value.

### HasProductVariants

`func (o *Invoice) HasProductVariants() bool`

HasProductVariants returns a boolean if a field has been set.

### GetExchangeRate

`func (o *Invoice) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Invoice) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Invoice) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Invoice) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetCryptoExchangeRate

`func (o *Invoice) GetCryptoExchangeRate() float64`

GetCryptoExchangeRate returns the CryptoExchangeRate field if non-nil, zero value otherwise.

### GetCryptoExchangeRateOk

`func (o *Invoice) GetCryptoExchangeRateOk() (*float64, bool)`

GetCryptoExchangeRateOk returns a tuple with the CryptoExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoExchangeRate

`func (o *Invoice) SetCryptoExchangeRate(v float64)`

SetCryptoExchangeRate sets CryptoExchangeRate field to given value.

### HasCryptoExchangeRate

`func (o *Invoice) HasCryptoExchangeRate() bool`

HasCryptoExchangeRate returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShopId

`func (o *Invoice) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Invoice) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Invoice) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Invoice) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetShopImageName

`func (o *Invoice) GetShopImageName() string`

GetShopImageName returns the ShopImageName field if non-nil, zero value otherwise.

### GetShopImageNameOk

`func (o *Invoice) GetShopImageNameOk() (*string, bool)`

GetShopImageNameOk returns a tuple with the ShopImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopImageName

`func (o *Invoice) SetShopImageName(v string)`

SetShopImageName sets ShopImageName field to given value.

### HasShopImageName

`func (o *Invoice) HasShopImageName() bool`

HasShopImageName returns a boolean if a field has been set.

### GetShopImageStorage

`func (o *Invoice) GetShopImageStorage() string`

GetShopImageStorage returns the ShopImageStorage field if non-nil, zero value otherwise.

### GetShopImageStorageOk

`func (o *Invoice) GetShopImageStorageOk() (*string, bool)`

GetShopImageStorageOk returns a tuple with the ShopImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopImageStorage

`func (o *Invoice) SetShopImageStorage(v string)`

SetShopImageStorage sets ShopImageStorage field to given value.

### HasShopImageStorage

`func (o *Invoice) HasShopImageStorage() bool`

HasShopImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *Invoice) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *Invoice) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *Invoice) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *Invoice) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetName

`func (o *Invoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Invoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Invoice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Invoice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *Invoice) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *Invoice) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *Invoice) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *Invoice) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerId

`func (o *Invoice) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Invoice) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Invoice) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Invoice) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetAffliateRevenueCustomerId

`func (o *Invoice) GetAffliateRevenueCustomerId() string`

GetAffliateRevenueCustomerId returns the AffliateRevenueCustomerId field if non-nil, zero value otherwise.

### GetAffliateRevenueCustomerIdOk

`func (o *Invoice) GetAffliateRevenueCustomerIdOk() (*string, bool)`

GetAffliateRevenueCustomerIdOk returns a tuple with the AffliateRevenueCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffliateRevenueCustomerId

`func (o *Invoice) SetAffliateRevenueCustomerId(v string)`

SetAffliateRevenueCustomerId sets AffliateRevenueCustomerId field to given value.

### HasAffliateRevenueCustomerId

`func (o *Invoice) HasAffliateRevenueCustomerId() bool`

HasAffliateRevenueCustomerId returns a boolean if a field has been set.

### GetPaypalEmailDelivery

`func (o *Invoice) GetPaypalEmailDelivery() bool`

GetPaypalEmailDelivery returns the PaypalEmailDelivery field if non-nil, zero value otherwise.

### GetPaypalEmailDeliveryOk

`func (o *Invoice) GetPaypalEmailDeliveryOk() (*bool, bool)`

GetPaypalEmailDeliveryOk returns a tuple with the PaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmailDelivery

`func (o *Invoice) SetPaypalEmailDelivery(v bool)`

SetPaypalEmailDelivery sets PaypalEmailDelivery field to given value.

### HasPaypalEmailDelivery

`func (o *Invoice) HasPaypalEmailDelivery() bool`

HasPaypalEmailDelivery returns a boolean if a field has been set.

### GetProductId

`func (o *Invoice) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Invoice) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Invoice) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Invoice) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductTitle

`func (o *Invoice) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *Invoice) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *Invoice) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *Invoice) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetProductType

`func (o *Invoice) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Invoice) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Invoice) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Invoice) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Invoice) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Invoice) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Invoice) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Invoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubscriptionTime

`func (o *Invoice) GetSubscriptionTime() int32`

GetSubscriptionTime returns the SubscriptionTime field if non-nil, zero value otherwise.

### GetSubscriptionTimeOk

`func (o *Invoice) GetSubscriptionTimeOk() (*int32, bool)`

GetSubscriptionTimeOk returns a tuple with the SubscriptionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTime

`func (o *Invoice) SetSubscriptionTime(v int32)`

SetSubscriptionTime sets SubscriptionTime field to given value.

### HasSubscriptionTime

`func (o *Invoice) HasSubscriptionTime() bool`

HasSubscriptionTime returns a boolean if a field has been set.

### GetGateway

`func (o *Invoice) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Invoice) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Invoice) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Invoice) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetBlockchain

`func (o *Invoice) GetBlockchain() Blockchain`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *Invoice) GetBlockchainOk() (*Blockchain, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *Invoice) SetBlockchain(v Blockchain)`

SetBlockchain sets Blockchain field to given value.

### HasBlockchain

`func (o *Invoice) HasBlockchain() bool`

HasBlockchain returns a boolean if a field has been set.

### GetPaypalApm

`func (o *Invoice) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *Invoice) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *Invoice) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *Invoice) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetStripeApm

`func (o *Invoice) GetStripeApm() string`

GetStripeApm returns the StripeApm field if non-nil, zero value otherwise.

### GetStripeApmOk

`func (o *Invoice) GetStripeApmOk() (*string, bool)`

GetStripeApmOk returns a tuple with the StripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeApm

`func (o *Invoice) SetStripeApm(v string)`

SetStripeApm sets StripeApm field to given value.

### HasStripeApm

`func (o *Invoice) HasStripeApm() bool`

HasStripeApm returns a boolean if a field has been set.

### GetPaypalEmail

`func (o *Invoice) GetPaypalEmail() string`

GetPaypalEmail returns the PaypalEmail field if non-nil, zero value otherwise.

### GetPaypalEmailOk

`func (o *Invoice) GetPaypalEmailOk() (*string, bool)`

GetPaypalEmailOk returns a tuple with the PaypalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmail

`func (o *Invoice) SetPaypalEmail(v string)`

SetPaypalEmail sets PaypalEmail field to given value.

### HasPaypalEmail

`func (o *Invoice) HasPaypalEmail() bool`

HasPaypalEmail returns a boolean if a field has been set.

### GetPaypalOrderId

`func (o *Invoice) GetPaypalOrderId() string`

GetPaypalOrderId returns the PaypalOrderId field if non-nil, zero value otherwise.

### GetPaypalOrderIdOk

`func (o *Invoice) GetPaypalOrderIdOk() (*string, bool)`

GetPaypalOrderIdOk returns a tuple with the PaypalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalOrderId

`func (o *Invoice) SetPaypalOrderId(v string)`

SetPaypalOrderId sets PaypalOrderId field to given value.

### HasPaypalOrderId

`func (o *Invoice) HasPaypalOrderId() bool`

HasPaypalOrderId returns a boolean if a field has been set.

### GetPaypalPayerEmail

`func (o *Invoice) GetPaypalPayerEmail() string`

GetPaypalPayerEmail returns the PaypalPayerEmail field if non-nil, zero value otherwise.

### GetPaypalPayerEmailOk

`func (o *Invoice) GetPaypalPayerEmailOk() (*string, bool)`

GetPaypalPayerEmailOk returns a tuple with the PaypalPayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerEmail

`func (o *Invoice) SetPaypalPayerEmail(v string)`

SetPaypalPayerEmail sets PaypalPayerEmail field to given value.

### HasPaypalPayerEmail

`func (o *Invoice) HasPaypalPayerEmail() bool`

HasPaypalPayerEmail returns a boolean if a field has been set.

### GetPaypalFee

`func (o *Invoice) GetPaypalFee() string`

GetPaypalFee returns the PaypalFee field if non-nil, zero value otherwise.

### GetPaypalFeeOk

`func (o *Invoice) GetPaypalFeeOk() (*string, bool)`

GetPaypalFeeOk returns a tuple with the PaypalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalFee

`func (o *Invoice) SetPaypalFee(v string)`

SetPaypalFee sets PaypalFee field to given value.

### HasPaypalFee

`func (o *Invoice) HasPaypalFee() bool`

HasPaypalFee returns a boolean if a field has been set.

### GetPaypalSubscriptionId

`func (o *Invoice) GetPaypalSubscriptionId() int32`

GetPaypalSubscriptionId returns the PaypalSubscriptionId field if non-nil, zero value otherwise.

### GetPaypalSubscriptionIdOk

`func (o *Invoice) GetPaypalSubscriptionIdOk() (*int32, bool)`

GetPaypalSubscriptionIdOk returns a tuple with the PaypalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionId

`func (o *Invoice) SetPaypalSubscriptionId(v int32)`

SetPaypalSubscriptionId sets PaypalSubscriptionId field to given value.

### HasPaypalSubscriptionId

`func (o *Invoice) HasPaypalSubscriptionId() bool`

HasPaypalSubscriptionId returns a boolean if a field has been set.

### GetPaypalSubscriptionLink

`func (o *Invoice) GetPaypalSubscriptionLink() int32`

GetPaypalSubscriptionLink returns the PaypalSubscriptionLink field if non-nil, zero value otherwise.

### GetPaypalSubscriptionLinkOk

`func (o *Invoice) GetPaypalSubscriptionLinkOk() (*int32, bool)`

GetPaypalSubscriptionLinkOk returns a tuple with the PaypalSubscriptionLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionLink

`func (o *Invoice) SetPaypalSubscriptionLink(v int32)`

SetPaypalSubscriptionLink sets PaypalSubscriptionLink field to given value.

### HasPaypalSubscriptionLink

`func (o *Invoice) HasPaypalSubscriptionLink() bool`

HasPaypalSubscriptionLink returns a boolean if a field has been set.

### GetLexOrderId

`func (o *Invoice) GetLexOrderId() string`

GetLexOrderId returns the LexOrderId field if non-nil, zero value otherwise.

### GetLexOrderIdOk

`func (o *Invoice) GetLexOrderIdOk() (*string, bool)`

GetLexOrderIdOk returns a tuple with the LexOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexOrderId

`func (o *Invoice) SetLexOrderId(v string)`

SetLexOrderId sets LexOrderId field to given value.

### HasLexOrderId

`func (o *Invoice) HasLexOrderId() bool`

HasLexOrderId returns a boolean if a field has been set.

### GetLexPaymentMethod

`func (o *Invoice) GetLexPaymentMethod() string`

GetLexPaymentMethod returns the LexPaymentMethod field if non-nil, zero value otherwise.

### GetLexPaymentMethodOk

`func (o *Invoice) GetLexPaymentMethodOk() (*string, bool)`

GetLexPaymentMethodOk returns a tuple with the LexPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLexPaymentMethod

`func (o *Invoice) SetLexPaymentMethod(v string)`

SetLexPaymentMethod sets LexPaymentMethod field to given value.

### HasLexPaymentMethod

`func (o *Invoice) HasLexPaymentMethod() bool`

HasLexPaymentMethod returns a boolean if a field has been set.

### GetPaydashPaymentID

`func (o *Invoice) GetPaydashPaymentID() string`

GetPaydashPaymentID returns the PaydashPaymentID field if non-nil, zero value otherwise.

### GetPaydashPaymentIDOk

`func (o *Invoice) GetPaydashPaymentIDOk() (*string, bool)`

GetPaydashPaymentIDOk returns a tuple with the PaydashPaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaydashPaymentID

`func (o *Invoice) SetPaydashPaymentID(v string)`

SetPaydashPaymentID sets PaydashPaymentID field to given value.

### HasPaydashPaymentID

`func (o *Invoice) HasPaydashPaymentID() bool`

HasPaydashPaymentID returns a boolean if a field has been set.

### GetVirtualPaymentsId

`func (o *Invoice) GetVirtualPaymentsId() string`

GetVirtualPaymentsId returns the VirtualPaymentsId field if non-nil, zero value otherwise.

### GetVirtualPaymentsIdOk

`func (o *Invoice) GetVirtualPaymentsIdOk() (*string, bool)`

GetVirtualPaymentsIdOk returns a tuple with the VirtualPaymentsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPaymentsId

`func (o *Invoice) SetVirtualPaymentsId(v string)`

SetVirtualPaymentsId sets VirtualPaymentsId field to given value.

### HasVirtualPaymentsId

`func (o *Invoice) HasVirtualPaymentsId() bool`

HasVirtualPaymentsId returns a boolean if a field has been set.

### GetStripeClientSecret

`func (o *Invoice) GetStripeClientSecret() string`

GetStripeClientSecret returns the StripeClientSecret field if non-nil, zero value otherwise.

### GetStripeClientSecretOk

`func (o *Invoice) GetStripeClientSecretOk() (*string, bool)`

GetStripeClientSecretOk returns a tuple with the StripeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeClientSecret

`func (o *Invoice) SetStripeClientSecret(v string)`

SetStripeClientSecret sets StripeClientSecret field to given value.

### HasStripeClientSecret

`func (o *Invoice) HasStripeClientSecret() bool`

HasStripeClientSecret returns a boolean if a field has been set.

### GetStripePriceId

`func (o *Invoice) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *Invoice) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *Invoice) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *Invoice) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetSkrillEmail

`func (o *Invoice) GetSkrillEmail() string`

GetSkrillEmail returns the SkrillEmail field if non-nil, zero value otherwise.

### GetSkrillEmailOk

`func (o *Invoice) GetSkrillEmailOk() (*string, bool)`

GetSkrillEmailOk returns a tuple with the SkrillEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillEmail

`func (o *Invoice) SetSkrillEmail(v string)`

SetSkrillEmail sets SkrillEmail field to given value.

### HasSkrillEmail

`func (o *Invoice) HasSkrillEmail() bool`

HasSkrillEmail returns a boolean if a field has been set.

### GetSkrillSid

`func (o *Invoice) GetSkrillSid() string`

GetSkrillSid returns the SkrillSid field if non-nil, zero value otherwise.

### GetSkrillSidOk

`func (o *Invoice) GetSkrillSidOk() (*string, bool)`

GetSkrillSidOk returns a tuple with the SkrillSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillSid

`func (o *Invoice) SetSkrillSid(v string)`

SetSkrillSid sets SkrillSid field to given value.

### HasSkrillSid

`func (o *Invoice) HasSkrillSid() bool`

HasSkrillSid returns a boolean if a field has been set.

### GetSkrillLink

`func (o *Invoice) GetSkrillLink() string`

GetSkrillLink returns the SkrillLink field if non-nil, zero value otherwise.

### GetSkrillLinkOk

`func (o *Invoice) GetSkrillLinkOk() (*string, bool)`

GetSkrillLinkOk returns a tuple with the SkrillLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillLink

`func (o *Invoice) SetSkrillLink(v string)`

SetSkrillLink sets SkrillLink field to given value.

### HasSkrillLink

`func (o *Invoice) HasSkrillLink() bool`

HasSkrillLink returns a boolean if a field has been set.

### GetPerfectmoneyId

`func (o *Invoice) GetPerfectmoneyId() string`

GetPerfectmoneyId returns the PerfectmoneyId field if non-nil, zero value otherwise.

### GetPerfectmoneyIdOk

`func (o *Invoice) GetPerfectmoneyIdOk() (*string, bool)`

GetPerfectmoneyIdOk returns a tuple with the PerfectmoneyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectmoneyId

`func (o *Invoice) SetPerfectmoneyId(v string)`

SetPerfectmoneyId sets PerfectmoneyId field to given value.

### HasPerfectmoneyId

`func (o *Invoice) HasPerfectmoneyId() bool`

HasPerfectmoneyId returns a boolean if a field has been set.

### GetBinanceInvoiceId

`func (o *Invoice) GetBinanceInvoiceId() string`

GetBinanceInvoiceId returns the BinanceInvoiceId field if non-nil, zero value otherwise.

### GetBinanceInvoiceIdOk

`func (o *Invoice) GetBinanceInvoiceIdOk() (*string, bool)`

GetBinanceInvoiceIdOk returns a tuple with the BinanceInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceInvoiceId

`func (o *Invoice) SetBinanceInvoiceId(v string)`

SetBinanceInvoiceId sets BinanceInvoiceId field to given value.

### HasBinanceInvoiceId

`func (o *Invoice) HasBinanceInvoiceId() bool`

HasBinanceInvoiceId returns a boolean if a field has been set.

### GetBinanceQrcode

`func (o *Invoice) GetBinanceQrcode() string`

GetBinanceQrcode returns the BinanceQrcode field if non-nil, zero value otherwise.

### GetBinanceQrcodeOk

`func (o *Invoice) GetBinanceQrcodeOk() (*string, bool)`

GetBinanceQrcodeOk returns a tuple with the BinanceQrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceQrcode

`func (o *Invoice) SetBinanceQrcode(v string)`

SetBinanceQrcode sets BinanceQrcode field to given value.

### HasBinanceQrcode

`func (o *Invoice) HasBinanceQrcode() bool`

HasBinanceQrcode returns a boolean if a field has been set.

### GetBinanceCheckoutUrl

`func (o *Invoice) GetBinanceCheckoutUrl() string`

GetBinanceCheckoutUrl returns the BinanceCheckoutUrl field if non-nil, zero value otherwise.

### GetBinanceCheckoutUrlOk

`func (o *Invoice) GetBinanceCheckoutUrlOk() (*string, bool)`

GetBinanceCheckoutUrlOk returns a tuple with the BinanceCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceCheckoutUrl

`func (o *Invoice) SetBinanceCheckoutUrl(v string)`

SetBinanceCheckoutUrl sets BinanceCheckoutUrl field to given value.

### HasBinanceCheckoutUrl

`func (o *Invoice) HasBinanceCheckoutUrl() bool`

HasBinanceCheckoutUrl returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *Invoice) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *Invoice) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *Invoice) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *Invoice) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *Invoice) GetCryptoAmount() float64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *Invoice) GetCryptoAmountOk() (*float64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *Invoice) SetCryptoAmount(v float64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *Invoice) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoReceived

`func (o *Invoice) GetCryptoReceived() float64`

GetCryptoReceived returns the CryptoReceived field if non-nil, zero value otherwise.

### GetCryptoReceivedOk

`func (o *Invoice) GetCryptoReceivedOk() (*float64, bool)`

GetCryptoReceivedOk returns a tuple with the CryptoReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoReceived

`func (o *Invoice) SetCryptoReceived(v float64)`

SetCryptoReceived sets CryptoReceived field to given value.

### HasCryptoReceived

`func (o *Invoice) HasCryptoReceived() bool`

HasCryptoReceived returns a boolean if a field has been set.

### GetCryptoUri

`func (o *Invoice) GetCryptoUri() string`

GetCryptoUri returns the CryptoUri field if non-nil, zero value otherwise.

### GetCryptoUriOk

`func (o *Invoice) GetCryptoUriOk() (*string, bool)`

GetCryptoUriOk returns a tuple with the CryptoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoUri

`func (o *Invoice) SetCryptoUri(v string)`

SetCryptoUri sets CryptoUri field to given value.

### HasCryptoUri

`func (o *Invoice) HasCryptoUri() bool`

HasCryptoUri returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *Invoice) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *Invoice) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *Invoice) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *Invoice) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetCryptoScheduledPayout

`func (o *Invoice) GetCryptoScheduledPayout() bool`

GetCryptoScheduledPayout returns the CryptoScheduledPayout field if non-nil, zero value otherwise.

### GetCryptoScheduledPayoutOk

`func (o *Invoice) GetCryptoScheduledPayoutOk() (*bool, bool)`

GetCryptoScheduledPayoutOk returns a tuple with the CryptoScheduledPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoScheduledPayout

`func (o *Invoice) SetCryptoScheduledPayout(v bool)`

SetCryptoScheduledPayout sets CryptoScheduledPayout field to given value.

### HasCryptoScheduledPayout

`func (o *Invoice) HasCryptoScheduledPayout() bool`

HasCryptoScheduledPayout returns a boolean if a field has been set.

### GetCryptoPayout

`func (o *Invoice) GetCryptoPayout() bool`

GetCryptoPayout returns the CryptoPayout field if non-nil, zero value otherwise.

### GetCryptoPayoutOk

`func (o *Invoice) GetCryptoPayoutOk() (*bool, bool)`

GetCryptoPayoutOk returns a tuple with the CryptoPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoPayout

`func (o *Invoice) SetCryptoPayout(v bool)`

SetCryptoPayout sets CryptoPayout field to given value.

### HasCryptoPayout

`func (o *Invoice) HasCryptoPayout() bool`

HasCryptoPayout returns a boolean if a field has been set.

### GetFeeBilled

`func (o *Invoice) GetFeeBilled() bool`

GetFeeBilled returns the FeeBilled field if non-nil, zero value otherwise.

### GetFeeBilledOk

`func (o *Invoice) GetFeeBilledOk() (*bool, bool)`

GetFeeBilledOk returns a tuple with the FeeBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBilled

`func (o *Invoice) SetFeeBilled(v bool)`

SetFeeBilled sets FeeBilled field to given value.

### HasFeeBilled

`func (o *Invoice) HasFeeBilled() bool`

HasFeeBilled returns a boolean if a field has been set.

### GetBillInfo

`func (o *Invoice) GetBillInfo() map[string]interface{}`

GetBillInfo returns the BillInfo field if non-nil, zero value otherwise.

### GetBillInfoOk

`func (o *Invoice) GetBillInfoOk() (*map[string]interface{}, bool)`

GetBillInfoOk returns a tuple with the BillInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillInfo

`func (o *Invoice) SetBillInfo(v map[string]interface{})`

SetBillInfo sets BillInfo field to given value.

### HasBillInfo

`func (o *Invoice) HasBillInfo() bool`

HasBillInfo returns a boolean if a field has been set.

### GetCashappQrcode

`func (o *Invoice) GetCashappQrcode() string`

GetCashappQrcode returns the CashappQrcode field if non-nil, zero value otherwise.

### GetCashappQrcodeOk

`func (o *Invoice) GetCashappQrcodeOk() (*string, bool)`

GetCashappQrcodeOk returns a tuple with the CashappQrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappQrcode

`func (o *Invoice) SetCashappQrcode(v string)`

SetCashappQrcode sets CashappQrcode field to given value.

### HasCashappQrcode

`func (o *Invoice) HasCashappQrcode() bool`

HasCashappQrcode returns a boolean if a field has been set.

### GetCashappNote

`func (o *Invoice) GetCashappNote() string`

GetCashappNote returns the CashappNote field if non-nil, zero value otherwise.

### GetCashappNoteOk

`func (o *Invoice) GetCashappNoteOk() (*string, bool)`

GetCashappNoteOk returns a tuple with the CashappNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappNote

`func (o *Invoice) SetCashappNote(v string)`

SetCashappNote sets CashappNote field to given value.

### HasCashappNote

`func (o *Invoice) HasCashappNote() bool`

HasCashappNote returns a boolean if a field has been set.

### GetCashappCashtag

`func (o *Invoice) GetCashappCashtag() string`

GetCashappCashtag returns the CashappCashtag field if non-nil, zero value otherwise.

### GetCashappCashtagOk

`func (o *Invoice) GetCashappCashtagOk() (*string, bool)`

GetCashappCashtagOk returns a tuple with the CashappCashtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappCashtag

`func (o *Invoice) SetCashappCashtag(v string)`

SetCashappCashtag sets CashappCashtag field to given value.

### HasCashappCashtag

`func (o *Invoice) HasCashappCashtag() bool`

HasCashappCashtag returns a boolean if a field has been set.

### GetCountry

`func (o *Invoice) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Invoice) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Invoice) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Invoice) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocation

`func (o *Invoice) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Invoice) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Invoice) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Invoice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIp

`func (o *Invoice) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Invoice) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Invoice) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Invoice) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsVpnOrProxy

`func (o *Invoice) GetIsVpnOrProxy() bool`

GetIsVpnOrProxy returns the IsVpnOrProxy field if non-nil, zero value otherwise.

### GetIsVpnOrProxyOk

`func (o *Invoice) GetIsVpnOrProxyOk() (*bool, bool)`

GetIsVpnOrProxyOk returns a tuple with the IsVpnOrProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpnOrProxy

`func (o *Invoice) SetIsVpnOrProxy(v bool)`

SetIsVpnOrProxy sets IsVpnOrProxy field to given value.

### HasIsVpnOrProxy

`func (o *Invoice) HasIsVpnOrProxy() bool`

HasIsVpnOrProxy returns a boolean if a field has been set.

### GetUserAgent

`func (o *Invoice) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Invoice) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Invoice) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Invoice) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetQuantity

`func (o *Invoice) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Invoice) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Invoice) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Invoice) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCouponId

`func (o *Invoice) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *Invoice) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *Invoice) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *Invoice) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCustomFields

`func (o *Invoice) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Invoice) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Invoice) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Invoice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDeveloperInvoice

`func (o *Invoice) GetDeveloperInvoice() bool`

GetDeveloperInvoice returns the DeveloperInvoice field if non-nil, zero value otherwise.

### GetDeveloperInvoiceOk

`func (o *Invoice) GetDeveloperInvoiceOk() (*bool, bool)`

GetDeveloperInvoiceOk returns a tuple with the DeveloperInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperInvoice

`func (o *Invoice) SetDeveloperInvoice(v bool)`

SetDeveloperInvoice sets DeveloperInvoice field to given value.

### HasDeveloperInvoice

`func (o *Invoice) HasDeveloperInvoice() bool`

HasDeveloperInvoice returns a boolean if a field has been set.

### GetDeveloperTitle

`func (o *Invoice) GetDeveloperTitle() string`

GetDeveloperTitle returns the DeveloperTitle field if non-nil, zero value otherwise.

### GetDeveloperTitleOk

`func (o *Invoice) GetDeveloperTitleOk() (*string, bool)`

GetDeveloperTitleOk returns a tuple with the DeveloperTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperTitle

`func (o *Invoice) SetDeveloperTitle(v string)`

SetDeveloperTitle sets DeveloperTitle field to given value.

### HasDeveloperTitle

`func (o *Invoice) HasDeveloperTitle() bool`

HasDeveloperTitle returns a boolean if a field has been set.

### GetDeveloperWebhook

`func (o *Invoice) GetDeveloperWebhook() string`

GetDeveloperWebhook returns the DeveloperWebhook field if non-nil, zero value otherwise.

### GetDeveloperWebhookOk

`func (o *Invoice) GetDeveloperWebhookOk() (*string, bool)`

GetDeveloperWebhookOk returns a tuple with the DeveloperWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperWebhook

`func (o *Invoice) SetDeveloperWebhook(v string)`

SetDeveloperWebhook sets DeveloperWebhook field to given value.

### HasDeveloperWebhook

`func (o *Invoice) HasDeveloperWebhook() bool`

HasDeveloperWebhook returns a boolean if a field has been set.

### GetDeveloperReturnUrl

`func (o *Invoice) GetDeveloperReturnUrl() string`

GetDeveloperReturnUrl returns the DeveloperReturnUrl field if non-nil, zero value otherwise.

### GetDeveloperReturnUrlOk

`func (o *Invoice) GetDeveloperReturnUrlOk() (*string, bool)`

GetDeveloperReturnUrlOk returns a tuple with the DeveloperReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperReturnUrl

`func (o *Invoice) SetDeveloperReturnUrl(v string)`

SetDeveloperReturnUrl sets DeveloperReturnUrl field to given value.

### HasDeveloperReturnUrl

`func (o *Invoice) HasDeveloperReturnUrl() bool`

HasDeveloperReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *Invoice) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Invoice) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Invoice) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Invoice) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetVoidDetails

`func (o *Invoice) GetVoidDetails() string`

GetVoidDetails returns the VoidDetails field if non-nil, zero value otherwise.

### GetVoidDetailsOk

`func (o *Invoice) GetVoidDetailsOk() (*string, bool)`

GetVoidDetailsOk returns a tuple with the VoidDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidDetails

`func (o *Invoice) SetVoidDetails(v string)`

SetVoidDetails sets VoidDetails field to given value.

### HasVoidDetails

`func (o *Invoice) HasVoidDetails() bool`

HasVoidDetails returns a boolean if a field has been set.

### GetDiscount

`func (o *Invoice) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Invoice) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Invoice) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Invoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetFeePercentage

`func (o *Invoice) GetFeePercentage() int32`

GetFeePercentage returns the FeePercentage field if non-nil, zero value otherwise.

### GetFeePercentageOk

`func (o *Invoice) GetFeePercentageOk() (*int32, bool)`

GetFeePercentageOk returns a tuple with the FeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercentage

`func (o *Invoice) SetFeePercentage(v int32)`

SetFeePercentage sets FeePercentage field to given value.

### HasFeePercentage

`func (o *Invoice) HasFeePercentage() bool`

HasFeePercentage returns a boolean if a field has been set.

### GetFeeBreakdown

`func (o *Invoice) GetFeeBreakdown() FeeBreakdown`

GetFeeBreakdown returns the FeeBreakdown field if non-nil, zero value otherwise.

### GetFeeBreakdownOk

`func (o *Invoice) GetFeeBreakdownOk() (*FeeBreakdown, bool)`

GetFeeBreakdownOk returns a tuple with the FeeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBreakdown

`func (o *Invoice) SetFeeBreakdown(v FeeBreakdown)`

SetFeeBreakdown sets FeeBreakdown field to given value.

### HasFeeBreakdown

`func (o *Invoice) HasFeeBreakdown() bool`

HasFeeBreakdown returns a boolean if a field has been set.

### GetDiscountBreakdown

`func (o *Invoice) GetDiscountBreakdown() map[string]interface{}`

GetDiscountBreakdown returns the DiscountBreakdown field if non-nil, zero value otherwise.

### GetDiscountBreakdownOk

`func (o *Invoice) GetDiscountBreakdownOk() (*map[string]interface{}, bool)`

GetDiscountBreakdownOk returns a tuple with the DiscountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBreakdown

`func (o *Invoice) SetDiscountBreakdown(v map[string]interface{})`

SetDiscountBreakdown sets DiscountBreakdown field to given value.

### HasDiscountBreakdown

`func (o *Invoice) HasDiscountBreakdown() bool`

HasDiscountBreakdown returns a boolean if a field has been set.

### GetDayValue

`func (o *Invoice) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *Invoice) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *Invoice) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *Invoice) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetDay

`func (o *Invoice) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Invoice) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Invoice) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *Invoice) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetMonth

`func (o *Invoice) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Invoice) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Invoice) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Invoice) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *Invoice) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Invoice) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Invoice) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *Invoice) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetProductAddons

`func (o *Invoice) GetProductAddons() []ProductAddonsInner`

GetProductAddons returns the ProductAddons field if non-nil, zero value otherwise.

### GetProductAddonsOk

`func (o *Invoice) GetProductAddonsOk() (*[]ProductAddonsInner, bool)`

GetProductAddonsOk returns a tuple with the ProductAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAddons

`func (o *Invoice) SetProductAddons(v []ProductAddonsInner)`

SetProductAddons sets ProductAddons field to given value.

### HasProductAddons

`func (o *Invoice) HasProductAddons() bool`

HasProductAddons returns a boolean if a field has been set.

### GetBundleConfig

`func (o *Invoice) GetBundleConfig() []BundleConfig`

GetBundleConfig returns the BundleConfig field if non-nil, zero value otherwise.

### GetBundleConfigOk

`func (o *Invoice) GetBundleConfigOk() (*[]BundleConfig, bool)`

GetBundleConfigOk returns a tuple with the BundleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleConfig

`func (o *Invoice) SetBundleConfig(v []BundleConfig)`

SetBundleConfig sets BundleConfig field to given value.

### HasBundleConfig

`func (o *Invoice) HasBundleConfig() bool`

HasBundleConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Invoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Invoice) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invoice) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invoice) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Invoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Invoice) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Invoice) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Invoice) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Invoice) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetApprovedAddress

`func (o *Invoice) GetApprovedAddress() ApprovedAddress`

GetApprovedAddress returns the ApprovedAddress field if non-nil, zero value otherwise.

### GetApprovedAddressOk

`func (o *Invoice) GetApprovedAddressOk() (*ApprovedAddress, bool)`

GetApprovedAddressOk returns a tuple with the ApprovedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAddress

`func (o *Invoice) SetApprovedAddress(v ApprovedAddress)`

SetApprovedAddress sets ApprovedAddress field to given value.

### HasApprovedAddress

`func (o *Invoice) HasApprovedAddress() bool`

HasApprovedAddress returns a boolean if a field has been set.

### GetServiceText

`func (o *Invoice) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *Invoice) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *Invoice) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *Invoice) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetIpInfo

`func (o *Invoice) GetIpInfo() IpInfo`

GetIpInfo returns the IpInfo field if non-nil, zero value otherwise.

### GetIpInfoOk

`func (o *Invoice) GetIpInfoOk() (*IpInfo, bool)`

GetIpInfoOk returns a tuple with the IpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpInfo

`func (o *Invoice) SetIpInfo(v IpInfo)`

SetIpInfo sets IpInfo field to given value.

### HasIpInfo

`func (o *Invoice) HasIpInfo() bool`

HasIpInfo returns a boolean if a field has been set.

### GetWebhooks

`func (o *Invoice) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *Invoice) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *Invoice) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *Invoice) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetRewardsData

`func (o *Invoice) GetRewardsData() []InvoiceRewardsDataInner`

GetRewardsData returns the RewardsData field if non-nil, zero value otherwise.

### GetRewardsDataOk

`func (o *Invoice) GetRewardsDataOk() (*[]InvoiceRewardsDataInner, bool)`

GetRewardsDataOk returns a tuple with the RewardsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsData

`func (o *Invoice) SetRewardsData(v []InvoiceRewardsDataInner)`

SetRewardsData sets RewardsData field to given value.

### HasRewardsData

`func (o *Invoice) HasRewardsData() bool`

HasRewardsData returns a boolean if a field has been set.

### GetPaypalDispute

`func (o *Invoice) GetPaypalDispute() PaypalDispute`

GetPaypalDispute returns the PaypalDispute field if non-nil, zero value otherwise.

### GetPaypalDisputeOk

`func (o *Invoice) GetPaypalDisputeOk() (*PaypalDispute, bool)`

GetPaypalDisputeOk returns a tuple with the PaypalDispute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalDispute

`func (o *Invoice) SetPaypalDispute(v PaypalDispute)`

SetPaypalDispute sets PaypalDispute field to given value.

### HasPaypalDispute

`func (o *Invoice) HasPaypalDispute() bool`

HasPaypalDispute returns a boolean if a field has been set.

### GetProductDownloads

`func (o *Invoice) GetProductDownloads() []ProductDownload`

GetProductDownloads returns the ProductDownloads field if non-nil, zero value otherwise.

### GetProductDownloadsOk

`func (o *Invoice) GetProductDownloadsOk() (*[]ProductDownload, bool)`

GetProductDownloadsOk returns a tuple with the ProductDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDownloads

`func (o *Invoice) SetProductDownloads(v []ProductDownload)`

SetProductDownloads sets ProductDownloads field to given value.

### HasProductDownloads

`func (o *Invoice) HasProductDownloads() bool`

HasProductDownloads returns a boolean if a field has been set.

### GetPaymentLinkId

`func (o *Invoice) GetPaymentLinkId() string`

GetPaymentLinkId returns the PaymentLinkId field if non-nil, zero value otherwise.

### GetPaymentLinkIdOk

`func (o *Invoice) GetPaymentLinkIdOk() (*string, bool)`

GetPaymentLinkIdOk returns a tuple with the PaymentLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLinkId

`func (o *Invoice) SetPaymentLinkId(v string)`

SetPaymentLinkId sets PaymentLinkId field to given value.

### HasPaymentLinkId

`func (o *Invoice) HasPaymentLinkId() bool`

HasPaymentLinkId returns a boolean if a field has been set.

### GetCashappEmailConfigured

`func (o *Invoice) GetCashappEmailConfigured() bool`

GetCashappEmailConfigured returns the CashappEmailConfigured field if non-nil, zero value otherwise.

### GetCashappEmailConfiguredOk

`func (o *Invoice) GetCashappEmailConfiguredOk() (*bool, bool)`

GetCashappEmailConfiguredOk returns a tuple with the CashappEmailConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappEmailConfigured

`func (o *Invoice) SetCashappEmailConfigured(v bool)`

SetCashappEmailConfigured sets CashappEmailConfigured field to given value.

### HasCashappEmailConfigured

`func (o *Invoice) HasCashappEmailConfigured() bool`

HasCashappEmailConfigured returns a boolean if a field has been set.

### GetLicense

`func (o *Invoice) GetLicense() bool`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Invoice) GetLicenseOk() (*bool, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Invoice) SetLicense(v bool)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Invoice) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetStatusHistory

`func (o *Invoice) GetStatusHistory() []InvoiceStatus`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *Invoice) GetStatusHistoryOk() (*[]InvoiceStatus, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *Invoice) SetStatusHistory(v []InvoiceStatus)`

SetStatusHistory sets StatusHistory field to given value.

### HasStatusHistory

`func (o *Invoice) HasStatusHistory() bool`

HasStatusHistory returns a boolean if a field has been set.

### GetAmlWallets

`func (o *Invoice) GetAmlWallets() []AmlWallet`

GetAmlWallets returns the AmlWallets field if non-nil, zero value otherwise.

### GetAmlWalletsOk

`func (o *Invoice) GetAmlWalletsOk() (*[]AmlWallet, bool)`

GetAmlWalletsOk returns a tuple with the AmlWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmlWallets

`func (o *Invoice) SetAmlWallets(v []AmlWallet)`

SetAmlWallets sets AmlWallets field to given value.

### HasAmlWallets

`func (o *Invoice) HasAmlWallets() bool`

HasAmlWallets returns a boolean if a field has been set.

### GetCryptoTransactions

`func (o *Invoice) GetCryptoTransactions() []CryptoTransaction`

GetCryptoTransactions returns the CryptoTransactions field if non-nil, zero value otherwise.

### GetCryptoTransactionsOk

`func (o *Invoice) GetCryptoTransactionsOk() (*[]CryptoTransaction, bool)`

GetCryptoTransactionsOk returns a tuple with the CryptoTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoTransactions

`func (o *Invoice) SetCryptoTransactions(v []CryptoTransaction)`

SetCryptoTransactions sets CryptoTransactions field to given value.

### HasCryptoTransactions

`func (o *Invoice) HasCryptoTransactions() bool`

HasCryptoTransactions returns a boolean if a field has been set.

### GetProduct

`func (o *Invoice) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Invoice) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Invoice) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Invoice) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTotalConversions

`func (o *Invoice) GetTotalConversions() interface{}`

GetTotalConversions returns the TotalConversions field if non-nil, zero value otherwise.

### GetTotalConversionsOk

`func (o *Invoice) GetTotalConversionsOk() (*interface{}, bool)`

GetTotalConversionsOk returns a tuple with the TotalConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConversions

`func (o *Invoice) SetTotalConversions(v interface{})`

SetTotalConversions sets TotalConversions field to given value.

### HasTotalConversions

`func (o *Invoice) HasTotalConversions() bool`

HasTotalConversions returns a boolean if a field has been set.

### SetTotalConversionsNil

`func (o *Invoice) SetTotalConversionsNil(b bool)`

 SetTotalConversionsNil sets the value for TotalConversions to be an explicit nil

### UnsetTotalConversions
`func (o *Invoice) UnsetTotalConversions()`

UnsetTotalConversions ensures that no value is present for TotalConversions, not even an explicit nil
### GetTheme

`func (o *Invoice) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Invoice) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Invoice) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Invoice) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetDarkMode

`func (o *Invoice) GetDarkMode() int32`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *Invoice) GetDarkModeOk() (*int32, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *Invoice) SetDarkMode(v int32)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *Invoice) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetCryptoMode

`func (o *Invoice) GetCryptoMode() interface{}`

GetCryptoMode returns the CryptoMode field if non-nil, zero value otherwise.

### GetCryptoModeOk

`func (o *Invoice) GetCryptoModeOk() (*interface{}, bool)`

GetCryptoModeOk returns a tuple with the CryptoMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoMode

`func (o *Invoice) SetCryptoMode(v interface{})`

SetCryptoMode sets CryptoMode field to given value.

### HasCryptoMode

`func (o *Invoice) HasCryptoMode() bool`

HasCryptoMode returns a boolean if a field has been set.

### SetCryptoModeNil

`func (o *Invoice) SetCryptoModeNil(b bool)`

 SetCryptoModeNil sets the value for CryptoMode to be an explicit nil

### UnsetCryptoMode
`func (o *Invoice) UnsetCryptoMode()`

UnsetCryptoMode ensures that no value is present for CryptoMode, not even an explicit nil
### GetGatewaysAvailable

`func (o *Invoice) GetGatewaysAvailable() []string`

GetGatewaysAvailable returns the GatewaysAvailable field if non-nil, zero value otherwise.

### GetGatewaysAvailableOk

`func (o *Invoice) GetGatewaysAvailableOk() (*[]string, bool)`

GetGatewaysAvailableOk returns a tuple with the GatewaysAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysAvailable

`func (o *Invoice) SetGatewaysAvailable(v []string)`

SetGatewaysAvailable sets GatewaysAvailable field to given value.

### HasGatewaysAvailable

`func (o *Invoice) HasGatewaysAvailable() bool`

HasGatewaysAvailable returns a boolean if a field has been set.

### GetCountryRegulations

`func (o *Invoice) GetCountryRegulations() string`

GetCountryRegulations returns the CountryRegulations field if non-nil, zero value otherwise.

### GetCountryRegulationsOk

`func (o *Invoice) GetCountryRegulationsOk() (*string, bool)`

GetCountryRegulationsOk returns a tuple with the CountryRegulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRegulations

`func (o *Invoice) SetCountryRegulations(v string)`

SetCountryRegulations sets CountryRegulations field to given value.

### HasCountryRegulations

`func (o *Invoice) HasCountryRegulations() bool`

HasCountryRegulations returns a boolean if a field has been set.

### GetAvailableStripeApm

`func (o *Invoice) GetAvailableStripeApm() InvoiceAvailableStripeApm`

GetAvailableStripeApm returns the AvailableStripeApm field if non-nil, zero value otherwise.

### GetAvailableStripeApmOk

`func (o *Invoice) GetAvailableStripeApmOk() (*InvoiceAvailableStripeApm, bool)`

GetAvailableStripeApmOk returns a tuple with the AvailableStripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStripeApm

`func (o *Invoice) SetAvailableStripeApm(v InvoiceAvailableStripeApm)`

SetAvailableStripeApm sets AvailableStripeApm field to given value.

### HasAvailableStripeApm

`func (o *Invoice) HasAvailableStripeApm() bool`

HasAvailableStripeApm returns a boolean if a field has been set.

### GetSerials

`func (o *Invoice) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *Invoice) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *Invoice) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *Invoice) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetShopPaymentGatewaysFees

`func (o *Invoice) GetShopPaymentGatewaysFees() []GatewayFees`

GetShopPaymentGatewaysFees returns the ShopPaymentGatewaysFees field if non-nil, zero value otherwise.

### GetShopPaymentGatewaysFeesOk

`func (o *Invoice) GetShopPaymentGatewaysFeesOk() (*[]GatewayFees, bool)`

GetShopPaymentGatewaysFeesOk returns a tuple with the ShopPaymentGatewaysFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopPaymentGatewaysFees

`func (o *Invoice) SetShopPaymentGatewaysFees(v []GatewayFees)`

SetShopPaymentGatewaysFees sets ShopPaymentGatewaysFees field to given value.

### HasShopPaymentGatewaysFees

`func (o *Invoice) HasShopPaymentGatewaysFees() bool`

HasShopPaymentGatewaysFees returns a boolean if a field has been set.

### GetShopPaypalCreditCard

`func (o *Invoice) GetShopPaypalCreditCard() bool`

GetShopPaypalCreditCard returns the ShopPaypalCreditCard field if non-nil, zero value otherwise.

### GetShopPaypalCreditCardOk

`func (o *Invoice) GetShopPaypalCreditCardOk() (*bool, bool)`

GetShopPaypalCreditCardOk returns a tuple with the ShopPaypalCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopPaypalCreditCard

`func (o *Invoice) SetShopPaypalCreditCard(v bool)`

SetShopPaypalCreditCard sets ShopPaypalCreditCard field to given value.

### HasShopPaypalCreditCard

`func (o *Invoice) HasShopPaypalCreditCard() bool`

HasShopPaypalCreditCard returns a boolean if a field has been set.

### GetShopForcePaypalEmailDelivery

`func (o *Invoice) GetShopForcePaypalEmailDelivery() bool`

GetShopForcePaypalEmailDelivery returns the ShopForcePaypalEmailDelivery field if non-nil, zero value otherwise.

### GetShopForcePaypalEmailDeliveryOk

`func (o *Invoice) GetShopForcePaypalEmailDeliveryOk() (*bool, bool)`

GetShopForcePaypalEmailDeliveryOk returns a tuple with the ShopForcePaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopForcePaypalEmailDelivery

`func (o *Invoice) SetShopForcePaypalEmailDelivery(v bool)`

SetShopForcePaypalEmailDelivery sets ShopForcePaypalEmailDelivery field to given value.

### HasShopForcePaypalEmailDelivery

`func (o *Invoice) HasShopForcePaypalEmailDelivery() bool`

HasShopForcePaypalEmailDelivery returns a boolean if a field has been set.

### GetShopWalletconnectId

`func (o *Invoice) GetShopWalletconnectId() interface{}`

GetShopWalletconnectId returns the ShopWalletconnectId field if non-nil, zero value otherwise.

### GetShopWalletconnectIdOk

`func (o *Invoice) GetShopWalletconnectIdOk() (*interface{}, bool)`

GetShopWalletconnectIdOk returns a tuple with the ShopWalletconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopWalletconnectId

`func (o *Invoice) SetShopWalletconnectId(v interface{})`

SetShopWalletconnectId sets ShopWalletconnectId field to given value.

### HasShopWalletconnectId

`func (o *Invoice) HasShopWalletconnectId() bool`

HasShopWalletconnectId returns a boolean if a field has been set.

### SetShopWalletconnectIdNil

`func (o *Invoice) SetShopWalletconnectIdNil(b bool)`

 SetShopWalletconnectIdNil sets the value for ShopWalletconnectId to be an explicit nil

### UnsetShopWalletconnectId
`func (o *Invoice) UnsetShopWalletconnectId()`

UnsetShopWalletconnectId ensures that no value is present for ShopWalletconnectId, not even an explicit nil
### GetOriginalDeveloperReturnUrl

`func (o *Invoice) GetOriginalDeveloperReturnUrl() string`

GetOriginalDeveloperReturnUrl returns the OriginalDeveloperReturnUrl field if non-nil, zero value otherwise.

### GetOriginalDeveloperReturnUrlOk

`func (o *Invoice) GetOriginalDeveloperReturnUrlOk() (*string, bool)`

GetOriginalDeveloperReturnUrlOk returns a tuple with the OriginalDeveloperReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDeveloperReturnUrl

`func (o *Invoice) SetOriginalDeveloperReturnUrl(v string)`

SetOriginalDeveloperReturnUrl sets OriginalDeveloperReturnUrl field to given value.

### HasOriginalDeveloperReturnUrl

`func (o *Invoice) HasOriginalDeveloperReturnUrl() bool`

HasOriginalDeveloperReturnUrl returns a boolean if a field has been set.

### GetRatesSnapshot

`func (o *Invoice) GetRatesSnapshot() InvoiceRatesSnapshot`

GetRatesSnapshot returns the RatesSnapshot field if non-nil, zero value otherwise.

### GetRatesSnapshotOk

`func (o *Invoice) GetRatesSnapshotOk() (*InvoiceRatesSnapshot, bool)`

GetRatesSnapshotOk returns a tuple with the RatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatesSnapshot

`func (o *Invoice) SetRatesSnapshot(v InvoiceRatesSnapshot)`

SetRatesSnapshot sets RatesSnapshot field to given value.

### HasRatesSnapshot

`func (o *Invoice) HasRatesSnapshot() bool`

HasRatesSnapshot returns a boolean if a field has been set.

### GetVoidTimes

`func (o *Invoice) GetVoidTimes() []InvoiceVoidTimesInner`

GetVoidTimes returns the VoidTimes field if non-nil, zero value otherwise.

### GetVoidTimesOk

`func (o *Invoice) GetVoidTimesOk() (*[]InvoiceVoidTimesInner, bool)`

GetVoidTimesOk returns a tuple with the VoidTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidTimes

`func (o *Invoice) SetVoidTimes(v []InvoiceVoidTimesInner)`

SetVoidTimes sets VoidTimes field to given value.

### HasVoidTimes

`func (o *Invoice) HasVoidTimes() bool`

HasVoidTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


