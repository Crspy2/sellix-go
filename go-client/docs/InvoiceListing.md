# InvoiceListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Total** | Pointer to **float64** | Invoice total, unit_price*quantity where quantity is applicable. | [optional] 
**TotalDisplay** | Pointer to **float64** | Invoice total in the currency chosen. | [optional] 
**ExchangeRate** | Pointer to **float64** | Exchange rate between currency chosen and USD. | [optional] 
**CryptoExchangeRate** | Pointer to **float64** | Exchange rate between the cryptocurrency chosen (if any) and USD. | [optional] 
**ProductVariants** | Pointer to [**[]ProductVariant**](ProductVariant.md) |  | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this invoice belongs. | [optional] 
**Type** | Pointer to **string** | Invoice type. For subscriptions, the invoice type is PRODUCT_SUBSCRIPTION as SUBSCRIPTION is reserved for Sellix-own plans. | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer. | [optional] 
**PaypalEmailDelivery** | Pointer to **bool** | If true and gateway is PAYPAL, the product will be shipped to the PayPal email on record instead of the customer email. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product linked to this invoice, if any. | [optional] 
**ProductTitle** | Pointer to **string** | Product title. | [optional] 
**ProductType** | Pointer to **string** | Product type. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PaypalEmail** | Pointer to **string** | DEPRECATED: Merchant PayPal email. | [optional] 
**PaypalOrderId** | Pointer to **string** | This field contains the PayPal order ID. | [optional] 
**PaypalFee** | Pointer to **string** | This field is updated after the invoice is completed with the fee taken by PayPal over the invoice total. | [optional] 
**PaypalPayerEmail** | Pointer to **string** | This field is updated after the invoice is completed with the PayPal&#39;s email used for the purchase. | [optional] 
**SkrillEmail** | Pointer to **string** | Merchant Skrill email. | [optional] 
**SkrillSid** | Pointer to **string** | Skrill unique ID linked to the invoice. | [optional] 
**SkrillLink** | Pointer to **string** | Skrill link to redirect the customer to. | [optional] 
**PerfectmoneyId** | Pointer to **string** | PerfectMoney payment ID linked to the invoice. | [optional] 
**CryptoAddress** | Pointer to **string** | Cryptocurrency address linked to this invoice. | [optional] 
**CryptoAmount** | Pointer to **float64** | Cryptocurrency amount converted based on crypto_exchange_rate. | [optional] 
**CryptoReceived** | Pointer to **float64** | Cryptocurrency amount received, paid by the customer. | [optional] 
**CryptoUri** | Pointer to **string** | URI used to create the QRCODE. | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Crypto confirmations needed to process the invoice. | [optional] 
**CashappCashtag** | Pointer to **string** | CashApp cashtag used to create the QRCODE. | [optional] 
**CashappNote** | Pointer to **string** | Unique note for the customer to leave in the CashApp payment. | [optional] 
**Country** | Pointer to **string** | Customer country. | [optional] 
**Location** | Pointer to **string** | Customer location. | [optional] 
**Ip** | Pointer to **string** | Customer IP. | [optional] 
**IsVpnOrProxy** | Pointer to **bool** | If true, a VPN or Proxy has been detected. | [optional] 
**UserAgent** | Pointer to **string** | Customer User Agent. | [optional] 
**Quantity** | Pointer to **int32** | Qauntity of product purchased. | [optional] 
**CouponId** | Pointer to **string** | Unique ID of the coupon, if used, for the discount. | [optional] 
**CouponCode** | Pointer to **string** | If coupon_id is specified, coupon_code contains the code used during the checkout. | [optional] 
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
**DayValue** | Pointer to **int32** | DEPRECATED: Day value, number. | [optional] 
**Day** | Pointer to **string** | DEPRECATED: First three letters of the day name. | [optional] 
**Month** | Pointer to **string** | DEPRECATED: First three letters of the month name. | [optional] 
**Year** | Pointer to **int32** | DEPRECATED: Year number. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the invoice. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the blacklist has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the invoice. | [optional] 
**IsEvmApproved** | Pointer to **bool** | Whether or not the EVM transaction is approved | [optional] 
**FeeBilled** | Pointer to **bool** | Whether or not a fee was billed to the customer | [optional] 
**CryptoScheduledPayout** | Pointer to [**InvoiceListingCryptoScheduledPayout**](InvoiceListingCryptoScheduledPayout.md) |  | [optional] 
**ShopPaypalCreditCard** | Pointer to **bool** | Whether or not to support credit cards on PayPal | [optional] 
**ShopForcePaypalEmailDelivery** | Pointer to **bool** | Whether or not to deliver the product to the PayPal email address | [optional] 

## Methods

### NewInvoiceListing

`func NewInvoiceListing() *InvoiceListing`

NewInvoiceListing instantiates a new InvoiceListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceListingWithDefaults

`func NewInvoiceListingWithDefaults() *InvoiceListing`

NewInvoiceListingWithDefaults instantiates a new InvoiceListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceListing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceListing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceListing) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *InvoiceListing) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *InvoiceListing) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *InvoiceListing) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *InvoiceListing) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetTotal

`func (o *InvoiceListing) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceListing) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceListing) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InvoiceListing) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalDisplay

`func (o *InvoiceListing) GetTotalDisplay() float64`

GetTotalDisplay returns the TotalDisplay field if non-nil, zero value otherwise.

### GetTotalDisplayOk

`func (o *InvoiceListing) GetTotalDisplayOk() (*float64, bool)`

GetTotalDisplayOk returns a tuple with the TotalDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisplay

`func (o *InvoiceListing) SetTotalDisplay(v float64)`

SetTotalDisplay sets TotalDisplay field to given value.

### HasTotalDisplay

`func (o *InvoiceListing) HasTotalDisplay() bool`

HasTotalDisplay returns a boolean if a field has been set.

### GetExchangeRate

`func (o *InvoiceListing) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *InvoiceListing) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *InvoiceListing) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *InvoiceListing) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetCryptoExchangeRate

`func (o *InvoiceListing) GetCryptoExchangeRate() float64`

GetCryptoExchangeRate returns the CryptoExchangeRate field if non-nil, zero value otherwise.

### GetCryptoExchangeRateOk

`func (o *InvoiceListing) GetCryptoExchangeRateOk() (*float64, bool)`

GetCryptoExchangeRateOk returns a tuple with the CryptoExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoExchangeRate

`func (o *InvoiceListing) SetCryptoExchangeRate(v float64)`

SetCryptoExchangeRate sets CryptoExchangeRate field to given value.

### HasCryptoExchangeRate

`func (o *InvoiceListing) HasCryptoExchangeRate() bool`

HasCryptoExchangeRate returns a boolean if a field has been set.

### GetProductVariants

`func (o *InvoiceListing) GetProductVariants() []ProductVariant`

GetProductVariants returns the ProductVariants field if non-nil, zero value otherwise.

### GetProductVariantsOk

`func (o *InvoiceListing) GetProductVariantsOk() (*[]ProductVariant, bool)`

GetProductVariantsOk returns a tuple with the ProductVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariants

`func (o *InvoiceListing) SetProductVariants(v []ProductVariant)`

SetProductVariants sets ProductVariants field to given value.

### HasProductVariants

`func (o *InvoiceListing) HasProductVariants() bool`

HasProductVariants returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceListing) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceListing) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceListing) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceListing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShopId

`func (o *InvoiceListing) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *InvoiceListing) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *InvoiceListing) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *InvoiceListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *InvoiceListing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceListing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceListing) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceListing) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *InvoiceListing) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *InvoiceListing) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *InvoiceListing) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *InvoiceListing) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetPaypalEmailDelivery

`func (o *InvoiceListing) GetPaypalEmailDelivery() bool`

GetPaypalEmailDelivery returns the PaypalEmailDelivery field if non-nil, zero value otherwise.

### GetPaypalEmailDeliveryOk

`func (o *InvoiceListing) GetPaypalEmailDeliveryOk() (*bool, bool)`

GetPaypalEmailDeliveryOk returns a tuple with the PaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmailDelivery

`func (o *InvoiceListing) SetPaypalEmailDelivery(v bool)`

SetPaypalEmailDelivery sets PaypalEmailDelivery field to given value.

### HasPaypalEmailDelivery

`func (o *InvoiceListing) HasPaypalEmailDelivery() bool`

HasPaypalEmailDelivery returns a boolean if a field has been set.

### GetProductId

`func (o *InvoiceListing) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *InvoiceListing) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *InvoiceListing) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *InvoiceListing) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductTitle

`func (o *InvoiceListing) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *InvoiceListing) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *InvoiceListing) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *InvoiceListing) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetProductType

`func (o *InvoiceListing) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *InvoiceListing) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *InvoiceListing) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *InvoiceListing) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetGateway

`func (o *InvoiceListing) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InvoiceListing) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InvoiceListing) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InvoiceListing) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPaypalEmail

`func (o *InvoiceListing) GetPaypalEmail() string`

GetPaypalEmail returns the PaypalEmail field if non-nil, zero value otherwise.

### GetPaypalEmailOk

`func (o *InvoiceListing) GetPaypalEmailOk() (*string, bool)`

GetPaypalEmailOk returns a tuple with the PaypalEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmail

`func (o *InvoiceListing) SetPaypalEmail(v string)`

SetPaypalEmail sets PaypalEmail field to given value.

### HasPaypalEmail

`func (o *InvoiceListing) HasPaypalEmail() bool`

HasPaypalEmail returns a boolean if a field has been set.

### GetPaypalOrderId

`func (o *InvoiceListing) GetPaypalOrderId() string`

GetPaypalOrderId returns the PaypalOrderId field if non-nil, zero value otherwise.

### GetPaypalOrderIdOk

`func (o *InvoiceListing) GetPaypalOrderIdOk() (*string, bool)`

GetPaypalOrderIdOk returns a tuple with the PaypalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalOrderId

`func (o *InvoiceListing) SetPaypalOrderId(v string)`

SetPaypalOrderId sets PaypalOrderId field to given value.

### HasPaypalOrderId

`func (o *InvoiceListing) HasPaypalOrderId() bool`

HasPaypalOrderId returns a boolean if a field has been set.

### GetPaypalFee

`func (o *InvoiceListing) GetPaypalFee() string`

GetPaypalFee returns the PaypalFee field if non-nil, zero value otherwise.

### GetPaypalFeeOk

`func (o *InvoiceListing) GetPaypalFeeOk() (*string, bool)`

GetPaypalFeeOk returns a tuple with the PaypalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalFee

`func (o *InvoiceListing) SetPaypalFee(v string)`

SetPaypalFee sets PaypalFee field to given value.

### HasPaypalFee

`func (o *InvoiceListing) HasPaypalFee() bool`

HasPaypalFee returns a boolean if a field has been set.

### GetPaypalPayerEmail

`func (o *InvoiceListing) GetPaypalPayerEmail() string`

GetPaypalPayerEmail returns the PaypalPayerEmail field if non-nil, zero value otherwise.

### GetPaypalPayerEmailOk

`func (o *InvoiceListing) GetPaypalPayerEmailOk() (*string, bool)`

GetPaypalPayerEmailOk returns a tuple with the PaypalPayerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayerEmail

`func (o *InvoiceListing) SetPaypalPayerEmail(v string)`

SetPaypalPayerEmail sets PaypalPayerEmail field to given value.

### HasPaypalPayerEmail

`func (o *InvoiceListing) HasPaypalPayerEmail() bool`

HasPaypalPayerEmail returns a boolean if a field has been set.

### GetSkrillEmail

`func (o *InvoiceListing) GetSkrillEmail() string`

GetSkrillEmail returns the SkrillEmail field if non-nil, zero value otherwise.

### GetSkrillEmailOk

`func (o *InvoiceListing) GetSkrillEmailOk() (*string, bool)`

GetSkrillEmailOk returns a tuple with the SkrillEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillEmail

`func (o *InvoiceListing) SetSkrillEmail(v string)`

SetSkrillEmail sets SkrillEmail field to given value.

### HasSkrillEmail

`func (o *InvoiceListing) HasSkrillEmail() bool`

HasSkrillEmail returns a boolean if a field has been set.

### GetSkrillSid

`func (o *InvoiceListing) GetSkrillSid() string`

GetSkrillSid returns the SkrillSid field if non-nil, zero value otherwise.

### GetSkrillSidOk

`func (o *InvoiceListing) GetSkrillSidOk() (*string, bool)`

GetSkrillSidOk returns a tuple with the SkrillSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillSid

`func (o *InvoiceListing) SetSkrillSid(v string)`

SetSkrillSid sets SkrillSid field to given value.

### HasSkrillSid

`func (o *InvoiceListing) HasSkrillSid() bool`

HasSkrillSid returns a boolean if a field has been set.

### GetSkrillLink

`func (o *InvoiceListing) GetSkrillLink() string`

GetSkrillLink returns the SkrillLink field if non-nil, zero value otherwise.

### GetSkrillLinkOk

`func (o *InvoiceListing) GetSkrillLinkOk() (*string, bool)`

GetSkrillLinkOk returns a tuple with the SkrillLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkrillLink

`func (o *InvoiceListing) SetSkrillLink(v string)`

SetSkrillLink sets SkrillLink field to given value.

### HasSkrillLink

`func (o *InvoiceListing) HasSkrillLink() bool`

HasSkrillLink returns a boolean if a field has been set.

### GetPerfectmoneyId

`func (o *InvoiceListing) GetPerfectmoneyId() string`

GetPerfectmoneyId returns the PerfectmoneyId field if non-nil, zero value otherwise.

### GetPerfectmoneyIdOk

`func (o *InvoiceListing) GetPerfectmoneyIdOk() (*string, bool)`

GetPerfectmoneyIdOk returns a tuple with the PerfectmoneyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectmoneyId

`func (o *InvoiceListing) SetPerfectmoneyId(v string)`

SetPerfectmoneyId sets PerfectmoneyId field to given value.

### HasPerfectmoneyId

`func (o *InvoiceListing) HasPerfectmoneyId() bool`

HasPerfectmoneyId returns a boolean if a field has been set.

### GetCryptoAddress

`func (o *InvoiceListing) GetCryptoAddress() string`

GetCryptoAddress returns the CryptoAddress field if non-nil, zero value otherwise.

### GetCryptoAddressOk

`func (o *InvoiceListing) GetCryptoAddressOk() (*string, bool)`

GetCryptoAddressOk returns a tuple with the CryptoAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddress

`func (o *InvoiceListing) SetCryptoAddress(v string)`

SetCryptoAddress sets CryptoAddress field to given value.

### HasCryptoAddress

`func (o *InvoiceListing) HasCryptoAddress() bool`

HasCryptoAddress returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *InvoiceListing) GetCryptoAmount() float64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *InvoiceListing) GetCryptoAmountOk() (*float64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *InvoiceListing) SetCryptoAmount(v float64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *InvoiceListing) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoReceived

`func (o *InvoiceListing) GetCryptoReceived() float64`

GetCryptoReceived returns the CryptoReceived field if non-nil, zero value otherwise.

### GetCryptoReceivedOk

`func (o *InvoiceListing) GetCryptoReceivedOk() (*float64, bool)`

GetCryptoReceivedOk returns a tuple with the CryptoReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoReceived

`func (o *InvoiceListing) SetCryptoReceived(v float64)`

SetCryptoReceived sets CryptoReceived field to given value.

### HasCryptoReceived

`func (o *InvoiceListing) HasCryptoReceived() bool`

HasCryptoReceived returns a boolean if a field has been set.

### GetCryptoUri

`func (o *InvoiceListing) GetCryptoUri() string`

GetCryptoUri returns the CryptoUri field if non-nil, zero value otherwise.

### GetCryptoUriOk

`func (o *InvoiceListing) GetCryptoUriOk() (*string, bool)`

GetCryptoUriOk returns a tuple with the CryptoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoUri

`func (o *InvoiceListing) SetCryptoUri(v string)`

SetCryptoUri sets CryptoUri field to given value.

### HasCryptoUri

`func (o *InvoiceListing) HasCryptoUri() bool`

HasCryptoUri returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *InvoiceListing) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *InvoiceListing) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *InvoiceListing) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *InvoiceListing) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetCashappCashtag

`func (o *InvoiceListing) GetCashappCashtag() string`

GetCashappCashtag returns the CashappCashtag field if non-nil, zero value otherwise.

### GetCashappCashtagOk

`func (o *InvoiceListing) GetCashappCashtagOk() (*string, bool)`

GetCashappCashtagOk returns a tuple with the CashappCashtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappCashtag

`func (o *InvoiceListing) SetCashappCashtag(v string)`

SetCashappCashtag sets CashappCashtag field to given value.

### HasCashappCashtag

`func (o *InvoiceListing) HasCashappCashtag() bool`

HasCashappCashtag returns a boolean if a field has been set.

### GetCashappNote

`func (o *InvoiceListing) GetCashappNote() string`

GetCashappNote returns the CashappNote field if non-nil, zero value otherwise.

### GetCashappNoteOk

`func (o *InvoiceListing) GetCashappNoteOk() (*string, bool)`

GetCashappNoteOk returns a tuple with the CashappNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashappNote

`func (o *InvoiceListing) SetCashappNote(v string)`

SetCashappNote sets CashappNote field to given value.

### HasCashappNote

`func (o *InvoiceListing) HasCashappNote() bool`

HasCashappNote returns a boolean if a field has been set.

### GetCountry

`func (o *InvoiceListing) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *InvoiceListing) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *InvoiceListing) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *InvoiceListing) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLocation

`func (o *InvoiceListing) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InvoiceListing) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InvoiceListing) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InvoiceListing) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetIp

`func (o *InvoiceListing) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InvoiceListing) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InvoiceListing) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InvoiceListing) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsVpnOrProxy

`func (o *InvoiceListing) GetIsVpnOrProxy() bool`

GetIsVpnOrProxy returns the IsVpnOrProxy field if non-nil, zero value otherwise.

### GetIsVpnOrProxyOk

`func (o *InvoiceListing) GetIsVpnOrProxyOk() (*bool, bool)`

GetIsVpnOrProxyOk returns a tuple with the IsVpnOrProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpnOrProxy

`func (o *InvoiceListing) SetIsVpnOrProxy(v bool)`

SetIsVpnOrProxy sets IsVpnOrProxy field to given value.

### HasIsVpnOrProxy

`func (o *InvoiceListing) HasIsVpnOrProxy() bool`

HasIsVpnOrProxy returns a boolean if a field has been set.

### GetUserAgent

`func (o *InvoiceListing) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *InvoiceListing) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *InvoiceListing) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *InvoiceListing) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceListing) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceListing) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceListing) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceListing) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCouponId

`func (o *InvoiceListing) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *InvoiceListing) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *InvoiceListing) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *InvoiceListing) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCouponCode

`func (o *InvoiceListing) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *InvoiceListing) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *InvoiceListing) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *InvoiceListing) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetCustomFields

`func (o *InvoiceListing) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InvoiceListing) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InvoiceListing) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InvoiceListing) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDeveloperInvoice

`func (o *InvoiceListing) GetDeveloperInvoice() bool`

GetDeveloperInvoice returns the DeveloperInvoice field if non-nil, zero value otherwise.

### GetDeveloperInvoiceOk

`func (o *InvoiceListing) GetDeveloperInvoiceOk() (*bool, bool)`

GetDeveloperInvoiceOk returns a tuple with the DeveloperInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperInvoice

`func (o *InvoiceListing) SetDeveloperInvoice(v bool)`

SetDeveloperInvoice sets DeveloperInvoice field to given value.

### HasDeveloperInvoice

`func (o *InvoiceListing) HasDeveloperInvoice() bool`

HasDeveloperInvoice returns a boolean if a field has been set.

### GetDeveloperTitle

`func (o *InvoiceListing) GetDeveloperTitle() string`

GetDeveloperTitle returns the DeveloperTitle field if non-nil, zero value otherwise.

### GetDeveloperTitleOk

`func (o *InvoiceListing) GetDeveloperTitleOk() (*string, bool)`

GetDeveloperTitleOk returns a tuple with the DeveloperTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperTitle

`func (o *InvoiceListing) SetDeveloperTitle(v string)`

SetDeveloperTitle sets DeveloperTitle field to given value.

### HasDeveloperTitle

`func (o *InvoiceListing) HasDeveloperTitle() bool`

HasDeveloperTitle returns a boolean if a field has been set.

### GetDeveloperWebhook

`func (o *InvoiceListing) GetDeveloperWebhook() string`

GetDeveloperWebhook returns the DeveloperWebhook field if non-nil, zero value otherwise.

### GetDeveloperWebhookOk

`func (o *InvoiceListing) GetDeveloperWebhookOk() (*string, bool)`

GetDeveloperWebhookOk returns a tuple with the DeveloperWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperWebhook

`func (o *InvoiceListing) SetDeveloperWebhook(v string)`

SetDeveloperWebhook sets DeveloperWebhook field to given value.

### HasDeveloperWebhook

`func (o *InvoiceListing) HasDeveloperWebhook() bool`

HasDeveloperWebhook returns a boolean if a field has been set.

### GetDeveloperReturnUrl

`func (o *InvoiceListing) GetDeveloperReturnUrl() string`

GetDeveloperReturnUrl returns the DeveloperReturnUrl field if non-nil, zero value otherwise.

### GetDeveloperReturnUrlOk

`func (o *InvoiceListing) GetDeveloperReturnUrlOk() (*string, bool)`

GetDeveloperReturnUrlOk returns a tuple with the DeveloperReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperReturnUrl

`func (o *InvoiceListing) SetDeveloperReturnUrl(v string)`

SetDeveloperReturnUrl sets DeveloperReturnUrl field to given value.

### HasDeveloperReturnUrl

`func (o *InvoiceListing) HasDeveloperReturnUrl() bool`

HasDeveloperReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceListing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceListing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *InvoiceListing) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *InvoiceListing) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *InvoiceListing) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *InvoiceListing) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetVoidDetails

`func (o *InvoiceListing) GetVoidDetails() string`

GetVoidDetails returns the VoidDetails field if non-nil, zero value otherwise.

### GetVoidDetailsOk

`func (o *InvoiceListing) GetVoidDetailsOk() (*string, bool)`

GetVoidDetailsOk returns a tuple with the VoidDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidDetails

`func (o *InvoiceListing) SetVoidDetails(v string)`

SetVoidDetails sets VoidDetails field to given value.

### HasVoidDetails

`func (o *InvoiceListing) HasVoidDetails() bool`

HasVoidDetails returns a boolean if a field has been set.

### GetDiscount

`func (o *InvoiceListing) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoiceListing) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoiceListing) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoiceListing) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetFeePercentage

`func (o *InvoiceListing) GetFeePercentage() int32`

GetFeePercentage returns the FeePercentage field if non-nil, zero value otherwise.

### GetFeePercentageOk

`func (o *InvoiceListing) GetFeePercentageOk() (*int32, bool)`

GetFeePercentageOk returns a tuple with the FeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercentage

`func (o *InvoiceListing) SetFeePercentage(v int32)`

SetFeePercentage sets FeePercentage field to given value.

### HasFeePercentage

`func (o *InvoiceListing) HasFeePercentage() bool`

HasFeePercentage returns a boolean if a field has been set.

### GetFeeBreakdown

`func (o *InvoiceListing) GetFeeBreakdown() FeeBreakdown`

GetFeeBreakdown returns the FeeBreakdown field if non-nil, zero value otherwise.

### GetFeeBreakdownOk

`func (o *InvoiceListing) GetFeeBreakdownOk() (*FeeBreakdown, bool)`

GetFeeBreakdownOk returns a tuple with the FeeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBreakdown

`func (o *InvoiceListing) SetFeeBreakdown(v FeeBreakdown)`

SetFeeBreakdown sets FeeBreakdown field to given value.

### HasFeeBreakdown

`func (o *InvoiceListing) HasFeeBreakdown() bool`

HasFeeBreakdown returns a boolean if a field has been set.

### GetDayValue

`func (o *InvoiceListing) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *InvoiceListing) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *InvoiceListing) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *InvoiceListing) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetDay

`func (o *InvoiceListing) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *InvoiceListing) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *InvoiceListing) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *InvoiceListing) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetMonth

`func (o *InvoiceListing) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *InvoiceListing) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *InvoiceListing) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *InvoiceListing) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *InvoiceListing) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *InvoiceListing) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *InvoiceListing) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *InvoiceListing) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvoiceListing) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceListing) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceListing) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvoiceListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *InvoiceListing) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *InvoiceListing) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *InvoiceListing) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *InvoiceListing) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetIsEvmApproved

`func (o *InvoiceListing) GetIsEvmApproved() bool`

GetIsEvmApproved returns the IsEvmApproved field if non-nil, zero value otherwise.

### GetIsEvmApprovedOk

`func (o *InvoiceListing) GetIsEvmApprovedOk() (*bool, bool)`

GetIsEvmApprovedOk returns a tuple with the IsEvmApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEvmApproved

`func (o *InvoiceListing) SetIsEvmApproved(v bool)`

SetIsEvmApproved sets IsEvmApproved field to given value.

### HasIsEvmApproved

`func (o *InvoiceListing) HasIsEvmApproved() bool`

HasIsEvmApproved returns a boolean if a field has been set.

### GetFeeBilled

`func (o *InvoiceListing) GetFeeBilled() bool`

GetFeeBilled returns the FeeBilled field if non-nil, zero value otherwise.

### GetFeeBilledOk

`func (o *InvoiceListing) GetFeeBilledOk() (*bool, bool)`

GetFeeBilledOk returns a tuple with the FeeBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBilled

`func (o *InvoiceListing) SetFeeBilled(v bool)`

SetFeeBilled sets FeeBilled field to given value.

### HasFeeBilled

`func (o *InvoiceListing) HasFeeBilled() bool`

HasFeeBilled returns a boolean if a field has been set.

### GetCryptoScheduledPayout

`func (o *InvoiceListing) GetCryptoScheduledPayout() InvoiceListingCryptoScheduledPayout`

GetCryptoScheduledPayout returns the CryptoScheduledPayout field if non-nil, zero value otherwise.

### GetCryptoScheduledPayoutOk

`func (o *InvoiceListing) GetCryptoScheduledPayoutOk() (*InvoiceListingCryptoScheduledPayout, bool)`

GetCryptoScheduledPayoutOk returns a tuple with the CryptoScheduledPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoScheduledPayout

`func (o *InvoiceListing) SetCryptoScheduledPayout(v InvoiceListingCryptoScheduledPayout)`

SetCryptoScheduledPayout sets CryptoScheduledPayout field to given value.

### HasCryptoScheduledPayout

`func (o *InvoiceListing) HasCryptoScheduledPayout() bool`

HasCryptoScheduledPayout returns a boolean if a field has been set.

### GetShopPaypalCreditCard

`func (o *InvoiceListing) GetShopPaypalCreditCard() bool`

GetShopPaypalCreditCard returns the ShopPaypalCreditCard field if non-nil, zero value otherwise.

### GetShopPaypalCreditCardOk

`func (o *InvoiceListing) GetShopPaypalCreditCardOk() (*bool, bool)`

GetShopPaypalCreditCardOk returns a tuple with the ShopPaypalCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopPaypalCreditCard

`func (o *InvoiceListing) SetShopPaypalCreditCard(v bool)`

SetShopPaypalCreditCard sets ShopPaypalCreditCard field to given value.

### HasShopPaypalCreditCard

`func (o *InvoiceListing) HasShopPaypalCreditCard() bool`

HasShopPaypalCreditCard returns a boolean if a field has been set.

### GetShopForcePaypalEmailDelivery

`func (o *InvoiceListing) GetShopForcePaypalEmailDelivery() bool`

GetShopForcePaypalEmailDelivery returns the ShopForcePaypalEmailDelivery field if non-nil, zero value otherwise.

### GetShopForcePaypalEmailDeliveryOk

`func (o *InvoiceListing) GetShopForcePaypalEmailDeliveryOk() (*bool, bool)`

GetShopForcePaypalEmailDeliveryOk returns a tuple with the ShopForcePaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopForcePaypalEmailDelivery

`func (o *InvoiceListing) SetShopForcePaypalEmailDelivery(v bool)`

SetShopForcePaypalEmailDelivery sets ShopForcePaypalEmailDelivery field to given value.

### HasShopForcePaypalEmailDelivery

`func (o *InvoiceListing) HasShopForcePaypalEmailDelivery() bool`

HasShopForcePaypalEmailDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


