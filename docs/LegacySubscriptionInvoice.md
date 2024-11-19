# LegacySubscriptionInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**RecurringBillingId** | Pointer to **string** | The Id of the legacy subscription | [optional] 
**Total** | Pointer to **float64** | Invoice total, unit_price*quantity where quantity is applicable. | [optional] 
**TotalDisplay** | Pointer to **float64** | Invoice total in the currency chosen. | [optional] 
**ExchangeRate** | Pointer to **float64** | Exchange rate between currency chosen and USD. | [optional] 
**CryptoExchangeRate** | Pointer to **float64** | Exchange rate between the cryptocurrency chosen (if any) and USD. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this invoice belongs. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product linked to this invoice, if any. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PaypalApm** | Pointer to **string** | If gateway is PAYPAL, a paypal_apm (PayPal Alternative Payment Method) can be specified. To retrieve the available PayPal APM for a specific customer session, please refer to the PayPal SDK using window.paypal.FUNDING and fundingSource to filter out available methods. You can also use our documentation on how to process white_label payments. | [optional] 
**StripeApm** | Pointer to **string** | Stripe Alternative Payment Method name, such as iDEAL, used if gateway is STRIPE. | [optional] 
**Quantity** | Pointer to **int32** | Qauntity of product purchased. | [optional] 
**CouponId** | Pointer to **string** | Unique ID of the coupon, if used, for the discount. | [optional] 
**Status** | Pointer to **string** | Status of the invoice. | [optional] 
**StatusDetails** | Pointer to **string** | If CART_PARTIAL_OUT_OF_STOCK, the invoice has been completed but some products in the cart were out of stock. | [optional] 
**VoidDetails** | Pointer to **string** | If the invoice is VOIDED and the product (or all the products in the cart) were out of stock, this additional status is set. | [optional] 
**Discount** | Pointer to **float64** | If a coupon or volume_discount is used, the discount value presents the total amount of discount over the total cost of the invoice. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the invoice. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the blacklist has been edited. | [optional] 
**CryptoAmount** | Pointer to **float64** | Cryptocurrency amount converted based on crypto_exchange_rate. | [optional] 
**CryptoReceived** | Pointer to **float64** | Cryptocurrency amount received, paid by the customer. | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Crypto confirmations needed to process the invoice. | [optional] 
**IsVpnOrProxy** | Pointer to **bool** | Whether the customer used a VPN or proxy when purchasing the subscription | [optional] 
**DeveloperInvoice** | Pointer to **bool** | If true, this invoice has been created through the Developers API. | [optional] 
**PaypalEmailDelivery** | Pointer to **bool** | If true and gateway is PAYPAL, the product will be shipped to the PayPal email on record instead of the customer email. | [optional] 
**FeeBilled** | Pointer to **bool** | Whether or not a fee was billed to the customer | [optional] 
**CryptoScheduledPayout** | Pointer to **bool** | Whether or not the payout is scheduled for the cryptocurrency | [optional] 
**ShopPaypalCreditCard** | Pointer to **bool** | Whether or not to support credit cards on PayPal | [optional] 
**ShopForcePaypalEmailDelivery** | Pointer to **bool** | Whether or not to deliver the product to the PayPal email address | [optional] 
**PaypalFee** | Pointer to **int32** | This field is updated after the invoice is completed with the fee taken by PayPal over the invoice total. | [optional] 
**FeePercentage** | Pointer to **int32** | What cut does Sellix take out of the total. To learn more about Sellix fees please refer to https://sellix.io/fees. | [optional] 
**DayValue** | Pointer to **int32** | DEPRECATED: Day value, number. | [optional] 
**Year** | Pointer to **int32** | DEPRECATED: Year number. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the invoice. | [optional] 

## Methods

### NewLegacySubscriptionInvoice

`func NewLegacySubscriptionInvoice() *LegacySubscriptionInvoice`

NewLegacySubscriptionInvoice instantiates a new LegacySubscriptionInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegacySubscriptionInvoiceWithDefaults

`func NewLegacySubscriptionInvoiceWithDefaults() *LegacySubscriptionInvoice`

NewLegacySubscriptionInvoiceWithDefaults instantiates a new LegacySubscriptionInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LegacySubscriptionInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegacySubscriptionInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegacySubscriptionInvoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LegacySubscriptionInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *LegacySubscriptionInvoice) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *LegacySubscriptionInvoice) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *LegacySubscriptionInvoice) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *LegacySubscriptionInvoice) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetRecurringBillingId

`func (o *LegacySubscriptionInvoice) GetRecurringBillingId() string`

GetRecurringBillingId returns the RecurringBillingId field if non-nil, zero value otherwise.

### GetRecurringBillingIdOk

`func (o *LegacySubscriptionInvoice) GetRecurringBillingIdOk() (*string, bool)`

GetRecurringBillingIdOk returns a tuple with the RecurringBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringBillingId

`func (o *LegacySubscriptionInvoice) SetRecurringBillingId(v string)`

SetRecurringBillingId sets RecurringBillingId field to given value.

### HasRecurringBillingId

`func (o *LegacySubscriptionInvoice) HasRecurringBillingId() bool`

HasRecurringBillingId returns a boolean if a field has been set.

### GetTotal

`func (o *LegacySubscriptionInvoice) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LegacySubscriptionInvoice) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LegacySubscriptionInvoice) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LegacySubscriptionInvoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalDisplay

`func (o *LegacySubscriptionInvoice) GetTotalDisplay() float64`

GetTotalDisplay returns the TotalDisplay field if non-nil, zero value otherwise.

### GetTotalDisplayOk

`func (o *LegacySubscriptionInvoice) GetTotalDisplayOk() (*float64, bool)`

GetTotalDisplayOk returns a tuple with the TotalDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDisplay

`func (o *LegacySubscriptionInvoice) SetTotalDisplay(v float64)`

SetTotalDisplay sets TotalDisplay field to given value.

### HasTotalDisplay

`func (o *LegacySubscriptionInvoice) HasTotalDisplay() bool`

HasTotalDisplay returns a boolean if a field has been set.

### GetExchangeRate

`func (o *LegacySubscriptionInvoice) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *LegacySubscriptionInvoice) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *LegacySubscriptionInvoice) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *LegacySubscriptionInvoice) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetCryptoExchangeRate

`func (o *LegacySubscriptionInvoice) GetCryptoExchangeRate() float64`

GetCryptoExchangeRate returns the CryptoExchangeRate field if non-nil, zero value otherwise.

### GetCryptoExchangeRateOk

`func (o *LegacySubscriptionInvoice) GetCryptoExchangeRateOk() (*float64, bool)`

GetCryptoExchangeRateOk returns a tuple with the CryptoExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoExchangeRate

`func (o *LegacySubscriptionInvoice) SetCryptoExchangeRate(v float64)`

SetCryptoExchangeRate sets CryptoExchangeRate field to given value.

### HasCryptoExchangeRate

`func (o *LegacySubscriptionInvoice) HasCryptoExchangeRate() bool`

HasCryptoExchangeRate returns a boolean if a field has been set.

### GetCurrency

`func (o *LegacySubscriptionInvoice) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LegacySubscriptionInvoice) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LegacySubscriptionInvoice) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LegacySubscriptionInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetShopId

`func (o *LegacySubscriptionInvoice) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *LegacySubscriptionInvoice) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *LegacySubscriptionInvoice) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *LegacySubscriptionInvoice) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductId

`func (o *LegacySubscriptionInvoice) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *LegacySubscriptionInvoice) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *LegacySubscriptionInvoice) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *LegacySubscriptionInvoice) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetGateway

`func (o *LegacySubscriptionInvoice) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *LegacySubscriptionInvoice) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *LegacySubscriptionInvoice) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *LegacySubscriptionInvoice) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPaypalApm

`func (o *LegacySubscriptionInvoice) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *LegacySubscriptionInvoice) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *LegacySubscriptionInvoice) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *LegacySubscriptionInvoice) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetStripeApm

`func (o *LegacySubscriptionInvoice) GetStripeApm() string`

GetStripeApm returns the StripeApm field if non-nil, zero value otherwise.

### GetStripeApmOk

`func (o *LegacySubscriptionInvoice) GetStripeApmOk() (*string, bool)`

GetStripeApmOk returns a tuple with the StripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeApm

`func (o *LegacySubscriptionInvoice) SetStripeApm(v string)`

SetStripeApm sets StripeApm field to given value.

### HasStripeApm

`func (o *LegacySubscriptionInvoice) HasStripeApm() bool`

HasStripeApm returns a boolean if a field has been set.

### GetQuantity

`func (o *LegacySubscriptionInvoice) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LegacySubscriptionInvoice) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LegacySubscriptionInvoice) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LegacySubscriptionInvoice) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCouponId

`func (o *LegacySubscriptionInvoice) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *LegacySubscriptionInvoice) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *LegacySubscriptionInvoice) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *LegacySubscriptionInvoice) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetStatus

`func (o *LegacySubscriptionInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegacySubscriptionInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegacySubscriptionInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegacySubscriptionInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *LegacySubscriptionInvoice) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *LegacySubscriptionInvoice) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *LegacySubscriptionInvoice) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *LegacySubscriptionInvoice) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetVoidDetails

`func (o *LegacySubscriptionInvoice) GetVoidDetails() string`

GetVoidDetails returns the VoidDetails field if non-nil, zero value otherwise.

### GetVoidDetailsOk

`func (o *LegacySubscriptionInvoice) GetVoidDetailsOk() (*string, bool)`

GetVoidDetailsOk returns a tuple with the VoidDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidDetails

`func (o *LegacySubscriptionInvoice) SetVoidDetails(v string)`

SetVoidDetails sets VoidDetails field to given value.

### HasVoidDetails

`func (o *LegacySubscriptionInvoice) HasVoidDetails() bool`

HasVoidDetails returns a boolean if a field has been set.

### GetDiscount

`func (o *LegacySubscriptionInvoice) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *LegacySubscriptionInvoice) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *LegacySubscriptionInvoice) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *LegacySubscriptionInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LegacySubscriptionInvoice) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegacySubscriptionInvoice) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegacySubscriptionInvoice) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegacySubscriptionInvoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LegacySubscriptionInvoice) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LegacySubscriptionInvoice) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LegacySubscriptionInvoice) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LegacySubscriptionInvoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *LegacySubscriptionInvoice) GetCryptoAmount() float64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *LegacySubscriptionInvoice) GetCryptoAmountOk() (*float64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *LegacySubscriptionInvoice) SetCryptoAmount(v float64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *LegacySubscriptionInvoice) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoReceived

`func (o *LegacySubscriptionInvoice) GetCryptoReceived() float64`

GetCryptoReceived returns the CryptoReceived field if non-nil, zero value otherwise.

### GetCryptoReceivedOk

`func (o *LegacySubscriptionInvoice) GetCryptoReceivedOk() (*float64, bool)`

GetCryptoReceivedOk returns a tuple with the CryptoReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoReceived

`func (o *LegacySubscriptionInvoice) SetCryptoReceived(v float64)`

SetCryptoReceived sets CryptoReceived field to given value.

### HasCryptoReceived

`func (o *LegacySubscriptionInvoice) HasCryptoReceived() bool`

HasCryptoReceived returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *LegacySubscriptionInvoice) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *LegacySubscriptionInvoice) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *LegacySubscriptionInvoice) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *LegacySubscriptionInvoice) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetIsVpnOrProxy

`func (o *LegacySubscriptionInvoice) GetIsVpnOrProxy() bool`

GetIsVpnOrProxy returns the IsVpnOrProxy field if non-nil, zero value otherwise.

### GetIsVpnOrProxyOk

`func (o *LegacySubscriptionInvoice) GetIsVpnOrProxyOk() (*bool, bool)`

GetIsVpnOrProxyOk returns a tuple with the IsVpnOrProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpnOrProxy

`func (o *LegacySubscriptionInvoice) SetIsVpnOrProxy(v bool)`

SetIsVpnOrProxy sets IsVpnOrProxy field to given value.

### HasIsVpnOrProxy

`func (o *LegacySubscriptionInvoice) HasIsVpnOrProxy() bool`

HasIsVpnOrProxy returns a boolean if a field has been set.

### GetDeveloperInvoice

`func (o *LegacySubscriptionInvoice) GetDeveloperInvoice() bool`

GetDeveloperInvoice returns the DeveloperInvoice field if non-nil, zero value otherwise.

### GetDeveloperInvoiceOk

`func (o *LegacySubscriptionInvoice) GetDeveloperInvoiceOk() (*bool, bool)`

GetDeveloperInvoiceOk returns a tuple with the DeveloperInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperInvoice

`func (o *LegacySubscriptionInvoice) SetDeveloperInvoice(v bool)`

SetDeveloperInvoice sets DeveloperInvoice field to given value.

### HasDeveloperInvoice

`func (o *LegacySubscriptionInvoice) HasDeveloperInvoice() bool`

HasDeveloperInvoice returns a boolean if a field has been set.

### GetPaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) GetPaypalEmailDelivery() bool`

GetPaypalEmailDelivery returns the PaypalEmailDelivery field if non-nil, zero value otherwise.

### GetPaypalEmailDeliveryOk

`func (o *LegacySubscriptionInvoice) GetPaypalEmailDeliveryOk() (*bool, bool)`

GetPaypalEmailDeliveryOk returns a tuple with the PaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) SetPaypalEmailDelivery(v bool)`

SetPaypalEmailDelivery sets PaypalEmailDelivery field to given value.

### HasPaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) HasPaypalEmailDelivery() bool`

HasPaypalEmailDelivery returns a boolean if a field has been set.

### GetFeeBilled

`func (o *LegacySubscriptionInvoice) GetFeeBilled() bool`

GetFeeBilled returns the FeeBilled field if non-nil, zero value otherwise.

### GetFeeBilledOk

`func (o *LegacySubscriptionInvoice) GetFeeBilledOk() (*bool, bool)`

GetFeeBilledOk returns a tuple with the FeeBilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBilled

`func (o *LegacySubscriptionInvoice) SetFeeBilled(v bool)`

SetFeeBilled sets FeeBilled field to given value.

### HasFeeBilled

`func (o *LegacySubscriptionInvoice) HasFeeBilled() bool`

HasFeeBilled returns a boolean if a field has been set.

### GetCryptoScheduledPayout

`func (o *LegacySubscriptionInvoice) GetCryptoScheduledPayout() bool`

GetCryptoScheduledPayout returns the CryptoScheduledPayout field if non-nil, zero value otherwise.

### GetCryptoScheduledPayoutOk

`func (o *LegacySubscriptionInvoice) GetCryptoScheduledPayoutOk() (*bool, bool)`

GetCryptoScheduledPayoutOk returns a tuple with the CryptoScheduledPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoScheduledPayout

`func (o *LegacySubscriptionInvoice) SetCryptoScheduledPayout(v bool)`

SetCryptoScheduledPayout sets CryptoScheduledPayout field to given value.

### HasCryptoScheduledPayout

`func (o *LegacySubscriptionInvoice) HasCryptoScheduledPayout() bool`

HasCryptoScheduledPayout returns a boolean if a field has been set.

### GetShopPaypalCreditCard

`func (o *LegacySubscriptionInvoice) GetShopPaypalCreditCard() bool`

GetShopPaypalCreditCard returns the ShopPaypalCreditCard field if non-nil, zero value otherwise.

### GetShopPaypalCreditCardOk

`func (o *LegacySubscriptionInvoice) GetShopPaypalCreditCardOk() (*bool, bool)`

GetShopPaypalCreditCardOk returns a tuple with the ShopPaypalCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopPaypalCreditCard

`func (o *LegacySubscriptionInvoice) SetShopPaypalCreditCard(v bool)`

SetShopPaypalCreditCard sets ShopPaypalCreditCard field to given value.

### HasShopPaypalCreditCard

`func (o *LegacySubscriptionInvoice) HasShopPaypalCreditCard() bool`

HasShopPaypalCreditCard returns a boolean if a field has been set.

### GetShopForcePaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) GetShopForcePaypalEmailDelivery() bool`

GetShopForcePaypalEmailDelivery returns the ShopForcePaypalEmailDelivery field if non-nil, zero value otherwise.

### GetShopForcePaypalEmailDeliveryOk

`func (o *LegacySubscriptionInvoice) GetShopForcePaypalEmailDeliveryOk() (*bool, bool)`

GetShopForcePaypalEmailDeliveryOk returns a tuple with the ShopForcePaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopForcePaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) SetShopForcePaypalEmailDelivery(v bool)`

SetShopForcePaypalEmailDelivery sets ShopForcePaypalEmailDelivery field to given value.

### HasShopForcePaypalEmailDelivery

`func (o *LegacySubscriptionInvoice) HasShopForcePaypalEmailDelivery() bool`

HasShopForcePaypalEmailDelivery returns a boolean if a field has been set.

### GetPaypalFee

`func (o *LegacySubscriptionInvoice) GetPaypalFee() int32`

GetPaypalFee returns the PaypalFee field if non-nil, zero value otherwise.

### GetPaypalFeeOk

`func (o *LegacySubscriptionInvoice) GetPaypalFeeOk() (*int32, bool)`

GetPaypalFeeOk returns a tuple with the PaypalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalFee

`func (o *LegacySubscriptionInvoice) SetPaypalFee(v int32)`

SetPaypalFee sets PaypalFee field to given value.

### HasPaypalFee

`func (o *LegacySubscriptionInvoice) HasPaypalFee() bool`

HasPaypalFee returns a boolean if a field has been set.

### GetFeePercentage

`func (o *LegacySubscriptionInvoice) GetFeePercentage() int32`

GetFeePercentage returns the FeePercentage field if non-nil, zero value otherwise.

### GetFeePercentageOk

`func (o *LegacySubscriptionInvoice) GetFeePercentageOk() (*int32, bool)`

GetFeePercentageOk returns a tuple with the FeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercentage

`func (o *LegacySubscriptionInvoice) SetFeePercentage(v int32)`

SetFeePercentage sets FeePercentage field to given value.

### HasFeePercentage

`func (o *LegacySubscriptionInvoice) HasFeePercentage() bool`

HasFeePercentage returns a boolean if a field has been set.

### GetDayValue

`func (o *LegacySubscriptionInvoice) GetDayValue() int32`

GetDayValue returns the DayValue field if non-nil, zero value otherwise.

### GetDayValueOk

`func (o *LegacySubscriptionInvoice) GetDayValueOk() (*int32, bool)`

GetDayValueOk returns a tuple with the DayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayValue

`func (o *LegacySubscriptionInvoice) SetDayValue(v int32)`

SetDayValue sets DayValue field to given value.

### HasDayValue

`func (o *LegacySubscriptionInvoice) HasDayValue() bool`

HasDayValue returns a boolean if a field has been set.

### GetYear

`func (o *LegacySubscriptionInvoice) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *LegacySubscriptionInvoice) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *LegacySubscriptionInvoice) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *LegacySubscriptionInvoice) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *LegacySubscriptionInvoice) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *LegacySubscriptionInvoice) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *LegacySubscriptionInvoice) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *LegacySubscriptionInvoice) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


