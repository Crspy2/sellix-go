# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the subscription. | [optional] 
**ProductSubscriptionId** | Pointer to **string** | ID of product subscription for subscriptions v2 | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this subscription belongs. | [optional] 
**ProductId** | Pointer to **string** | ID of the product subscription. | [optional] 
**ProductPlanId** | Pointer to **string** | ID of the subscriptions v2 active product plan | [optional] 
**CustomerId** | Pointer to **string** | ID of the store&#39;s customer. | [optional] 
**Type** | Pointer to **string** | The subscription version | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the subscription. | [optional] 
**UpdatedAt** | Pointer to **int32** | When this subscription was last updated. | [optional] 
**Status** | Pointer to **string** | Subscription status. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**CouponId** | Pointer to **string** | The ID of the coupon used when paying for the subscription. Will be null if no coupon was used | [optional] 
**CurrentPeriodEnd** | Pointer to **int32** | When the current billing period will end. | [optional] 
**UpcomingEmail1WeekSent** | Pointer to **bool** | Whether or not the email for an upcoming renewal has been sent. | [optional] 
**TrialPeriodEndingEmailSent** | Pointer to **bool** | Whether or not the email for the trial period expiring has been sent. | [optional] 
**RenewalInvoiceCreated** | Pointer to **bool** | If true, the renewal invoice has already been created. | [optional] 
**StatusDetails** | Pointer to **string** | Details about the subscription status. | [optional] 
**CanceledAt** | Pointer to **int32** | When this subscription was canceled. | [optional] 
**StripeCustomerId** | Pointer to **string** | ID of the customer created on STRIPE. | [optional] 
**StripeAccount** | Pointer to **string** | ID of the Stripe connected account. | [optional] 
**StripeSubscriptionId** | Pointer to **string** | ID of the Stripe subscription. | [optional] 
**PaypalAccount** | Pointer to **string** | The ID of the PayPal account used to complete the subscription payments | [optional] 
**PaypalSubscriptionId** | Pointer to **string** | The ID of the PayPal subscription | [optional] 
**ProductTitle** | Pointer to **string** | Digital Software | [optional] 
**ProductPlanTitle** | Pointer to **string** | The title of the product plan for this subscription (only for type: v2) | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**CustomerName** | Pointer to **string** | Customer name. | [optional] 
**CustomerSurname** | Pointer to **string** | Customer surname. | [optional] 
**CustomerPhone** | Pointer to **string** | Customer phone. | [optional] 
**CustomerPhoneCountryCode** | Pointer to **string** | Customer phone country code. | [optional] 
**CustomerCountryCode** | Pointer to **string** | Customer country code. | [optional] 
**CustomerStreetAddress** | Pointer to **string** | Customer street address. | [optional] 
**CustomerAdditionalAddressInfo** | Pointer to **string** | Customer street address additional info. | [optional] 
**CustomerCity** | Pointer to **string** | Customer city. | [optional] 
**CustomerPostalCode** | Pointer to **string** | Customer postal code. | [optional] 
**CustomerState** | Pointer to **string** | Customer state. | [optional] 
**CustomerEmail** | Pointer to **string** | Customer email. | [optional] 
**ShopName** | Pointer to **string** | The name of the shop the subscription was made for | [optional] 
**Invoices** | Pointer to [**[]LegacySubscriptionInvoice**](LegacySubscriptionInvoice.md) |  | [optional] 
**ApprovedAddress** | Pointer to [**ApprovedAddress**](ApprovedAddress.md) |  | [optional] 
**GatewaysAvailable** | Pointer to [**[]Gateway**](Gateway.md) | Gateways available for this invoice | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductSubscriptionId

`func (o *Subscription) GetProductSubscriptionId() string`

GetProductSubscriptionId returns the ProductSubscriptionId field if non-nil, zero value otherwise.

### GetProductSubscriptionIdOk

`func (o *Subscription) GetProductSubscriptionIdOk() (*string, bool)`

GetProductSubscriptionIdOk returns a tuple with the ProductSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionId

`func (o *Subscription) SetProductSubscriptionId(v string)`

SetProductSubscriptionId sets ProductSubscriptionId field to given value.

### HasProductSubscriptionId

`func (o *Subscription) HasProductSubscriptionId() bool`

HasProductSubscriptionId returns a boolean if a field has been set.

### GetShopId

`func (o *Subscription) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Subscription) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Subscription) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Subscription) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductId

`func (o *Subscription) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Subscription) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Subscription) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Subscription) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductPlanId

`func (o *Subscription) GetProductPlanId() string`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *Subscription) GetProductPlanIdOk() (*string, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *Subscription) SetProductPlanId(v string)`

SetProductPlanId sets ProductPlanId field to given value.

### HasProductPlanId

`func (o *Subscription) HasProductPlanId() bool`

HasProductPlanId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Subscription) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Subscription) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Subscription) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Subscription) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetType

`func (o *Subscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Subscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGateway

`func (o *Subscription) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Subscription) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Subscription) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Subscription) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetCustomFields

`func (o *Subscription) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Subscription) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Subscription) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Subscription) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCouponId

`func (o *Subscription) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *Subscription) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *Subscription) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *Subscription) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() int32`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*int32, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v int32)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *Subscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetUpcomingEmail1WeekSent

`func (o *Subscription) GetUpcomingEmail1WeekSent() bool`

GetUpcomingEmail1WeekSent returns the UpcomingEmail1WeekSent field if non-nil, zero value otherwise.

### GetUpcomingEmail1WeekSentOk

`func (o *Subscription) GetUpcomingEmail1WeekSentOk() (*bool, bool)`

GetUpcomingEmail1WeekSentOk returns a tuple with the UpcomingEmail1WeekSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingEmail1WeekSent

`func (o *Subscription) SetUpcomingEmail1WeekSent(v bool)`

SetUpcomingEmail1WeekSent sets UpcomingEmail1WeekSent field to given value.

### HasUpcomingEmail1WeekSent

`func (o *Subscription) HasUpcomingEmail1WeekSent() bool`

HasUpcomingEmail1WeekSent returns a boolean if a field has been set.

### GetTrialPeriodEndingEmailSent

`func (o *Subscription) GetTrialPeriodEndingEmailSent() bool`

GetTrialPeriodEndingEmailSent returns the TrialPeriodEndingEmailSent field if non-nil, zero value otherwise.

### GetTrialPeriodEndingEmailSentOk

`func (o *Subscription) GetTrialPeriodEndingEmailSentOk() (*bool, bool)`

GetTrialPeriodEndingEmailSentOk returns a tuple with the TrialPeriodEndingEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodEndingEmailSent

`func (o *Subscription) SetTrialPeriodEndingEmailSent(v bool)`

SetTrialPeriodEndingEmailSent sets TrialPeriodEndingEmailSent field to given value.

### HasTrialPeriodEndingEmailSent

`func (o *Subscription) HasTrialPeriodEndingEmailSent() bool`

HasTrialPeriodEndingEmailSent returns a boolean if a field has been set.

### GetRenewalInvoiceCreated

`func (o *Subscription) GetRenewalInvoiceCreated() bool`

GetRenewalInvoiceCreated returns the RenewalInvoiceCreated field if non-nil, zero value otherwise.

### GetRenewalInvoiceCreatedOk

`func (o *Subscription) GetRenewalInvoiceCreatedOk() (*bool, bool)`

GetRenewalInvoiceCreatedOk returns a tuple with the RenewalInvoiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalInvoiceCreated

`func (o *Subscription) SetRenewalInvoiceCreated(v bool)`

SetRenewalInvoiceCreated sets RenewalInvoiceCreated field to given value.

### HasRenewalInvoiceCreated

`func (o *Subscription) HasRenewalInvoiceCreated() bool`

HasRenewalInvoiceCreated returns a boolean if a field has been set.

### GetStatusDetails

`func (o *Subscription) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Subscription) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Subscription) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Subscription) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetCanceledAt

`func (o *Subscription) GetCanceledAt() int32`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*int32, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v int32)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *Subscription) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *Subscription) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *Subscription) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *Subscription) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeAccount

`func (o *Subscription) GetStripeAccount() string`

GetStripeAccount returns the StripeAccount field if non-nil, zero value otherwise.

### GetStripeAccountOk

`func (o *Subscription) GetStripeAccountOk() (*string, bool)`

GetStripeAccountOk returns a tuple with the StripeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccount

`func (o *Subscription) SetStripeAccount(v string)`

SetStripeAccount sets StripeAccount field to given value.

### HasStripeAccount

`func (o *Subscription) HasStripeAccount() bool`

HasStripeAccount returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *Subscription) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *Subscription) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *Subscription) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *Subscription) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetPaypalAccount

`func (o *Subscription) GetPaypalAccount() string`

GetPaypalAccount returns the PaypalAccount field if non-nil, zero value otherwise.

### GetPaypalAccountOk

`func (o *Subscription) GetPaypalAccountOk() (*string, bool)`

GetPaypalAccountOk returns a tuple with the PaypalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalAccount

`func (o *Subscription) SetPaypalAccount(v string)`

SetPaypalAccount sets PaypalAccount field to given value.

### HasPaypalAccount

`func (o *Subscription) HasPaypalAccount() bool`

HasPaypalAccount returns a boolean if a field has been set.

### GetPaypalSubscriptionId

`func (o *Subscription) GetPaypalSubscriptionId() string`

GetPaypalSubscriptionId returns the PaypalSubscriptionId field if non-nil, zero value otherwise.

### GetPaypalSubscriptionIdOk

`func (o *Subscription) GetPaypalSubscriptionIdOk() (*string, bool)`

GetPaypalSubscriptionIdOk returns a tuple with the PaypalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionId

`func (o *Subscription) SetPaypalSubscriptionId(v string)`

SetPaypalSubscriptionId sets PaypalSubscriptionId field to given value.

### HasPaypalSubscriptionId

`func (o *Subscription) HasPaypalSubscriptionId() bool`

HasPaypalSubscriptionId returns a boolean if a field has been set.

### GetProductTitle

`func (o *Subscription) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *Subscription) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *Subscription) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *Subscription) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetProductPlanTitle

`func (o *Subscription) GetProductPlanTitle() string`

GetProductPlanTitle returns the ProductPlanTitle field if non-nil, zero value otherwise.

### GetProductPlanTitleOk

`func (o *Subscription) GetProductPlanTitleOk() (*string, bool)`

GetProductPlanTitleOk returns a tuple with the ProductPlanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanTitle

`func (o *Subscription) SetProductPlanTitle(v string)`

SetProductPlanTitle sets ProductPlanTitle field to given value.

### HasProductPlanTitle

`func (o *Subscription) HasProductPlanTitle() bool`

HasProductPlanTitle returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *Subscription) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *Subscription) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *Subscription) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *Subscription) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetCustomerName

`func (o *Subscription) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *Subscription) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *Subscription) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *Subscription) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerSurname

`func (o *Subscription) GetCustomerSurname() string`

GetCustomerSurname returns the CustomerSurname field if non-nil, zero value otherwise.

### GetCustomerSurnameOk

`func (o *Subscription) GetCustomerSurnameOk() (*string, bool)`

GetCustomerSurnameOk returns a tuple with the CustomerSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSurname

`func (o *Subscription) SetCustomerSurname(v string)`

SetCustomerSurname sets CustomerSurname field to given value.

### HasCustomerSurname

`func (o *Subscription) HasCustomerSurname() bool`

HasCustomerSurname returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *Subscription) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *Subscription) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *Subscription) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *Subscription) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerPhoneCountryCode

`func (o *Subscription) GetCustomerPhoneCountryCode() string`

GetCustomerPhoneCountryCode returns the CustomerPhoneCountryCode field if non-nil, zero value otherwise.

### GetCustomerPhoneCountryCodeOk

`func (o *Subscription) GetCustomerPhoneCountryCodeOk() (*string, bool)`

GetCustomerPhoneCountryCodeOk returns a tuple with the CustomerPhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhoneCountryCode

`func (o *Subscription) SetCustomerPhoneCountryCode(v string)`

SetCustomerPhoneCountryCode sets CustomerPhoneCountryCode field to given value.

### HasCustomerPhoneCountryCode

`func (o *Subscription) HasCustomerPhoneCountryCode() bool`

HasCustomerPhoneCountryCode returns a boolean if a field has been set.

### GetCustomerCountryCode

`func (o *Subscription) GetCustomerCountryCode() string`

GetCustomerCountryCode returns the CustomerCountryCode field if non-nil, zero value otherwise.

### GetCustomerCountryCodeOk

`func (o *Subscription) GetCustomerCountryCodeOk() (*string, bool)`

GetCustomerCountryCodeOk returns a tuple with the CustomerCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountryCode

`func (o *Subscription) SetCustomerCountryCode(v string)`

SetCustomerCountryCode sets CustomerCountryCode field to given value.

### HasCustomerCountryCode

`func (o *Subscription) HasCustomerCountryCode() bool`

HasCustomerCountryCode returns a boolean if a field has been set.

### GetCustomerStreetAddress

`func (o *Subscription) GetCustomerStreetAddress() string`

GetCustomerStreetAddress returns the CustomerStreetAddress field if non-nil, zero value otherwise.

### GetCustomerStreetAddressOk

`func (o *Subscription) GetCustomerStreetAddressOk() (*string, bool)`

GetCustomerStreetAddressOk returns a tuple with the CustomerStreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerStreetAddress

`func (o *Subscription) SetCustomerStreetAddress(v string)`

SetCustomerStreetAddress sets CustomerStreetAddress field to given value.

### HasCustomerStreetAddress

`func (o *Subscription) HasCustomerStreetAddress() bool`

HasCustomerStreetAddress returns a boolean if a field has been set.

### GetCustomerAdditionalAddressInfo

`func (o *Subscription) GetCustomerAdditionalAddressInfo() string`

GetCustomerAdditionalAddressInfo returns the CustomerAdditionalAddressInfo field if non-nil, zero value otherwise.

### GetCustomerAdditionalAddressInfoOk

`func (o *Subscription) GetCustomerAdditionalAddressInfoOk() (*string, bool)`

GetCustomerAdditionalAddressInfoOk returns a tuple with the CustomerAdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAdditionalAddressInfo

`func (o *Subscription) SetCustomerAdditionalAddressInfo(v string)`

SetCustomerAdditionalAddressInfo sets CustomerAdditionalAddressInfo field to given value.

### HasCustomerAdditionalAddressInfo

`func (o *Subscription) HasCustomerAdditionalAddressInfo() bool`

HasCustomerAdditionalAddressInfo returns a boolean if a field has been set.

### GetCustomerCity

`func (o *Subscription) GetCustomerCity() string`

GetCustomerCity returns the CustomerCity field if non-nil, zero value otherwise.

### GetCustomerCityOk

`func (o *Subscription) GetCustomerCityOk() (*string, bool)`

GetCustomerCityOk returns a tuple with the CustomerCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCity

`func (o *Subscription) SetCustomerCity(v string)`

SetCustomerCity sets CustomerCity field to given value.

### HasCustomerCity

`func (o *Subscription) HasCustomerCity() bool`

HasCustomerCity returns a boolean if a field has been set.

### GetCustomerPostalCode

`func (o *Subscription) GetCustomerPostalCode() string`

GetCustomerPostalCode returns the CustomerPostalCode field if non-nil, zero value otherwise.

### GetCustomerPostalCodeOk

`func (o *Subscription) GetCustomerPostalCodeOk() (*string, bool)`

GetCustomerPostalCodeOk returns a tuple with the CustomerPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPostalCode

`func (o *Subscription) SetCustomerPostalCode(v string)`

SetCustomerPostalCode sets CustomerPostalCode field to given value.

### HasCustomerPostalCode

`func (o *Subscription) HasCustomerPostalCode() bool`

HasCustomerPostalCode returns a boolean if a field has been set.

### GetCustomerState

`func (o *Subscription) GetCustomerState() string`

GetCustomerState returns the CustomerState field if non-nil, zero value otherwise.

### GetCustomerStateOk

`func (o *Subscription) GetCustomerStateOk() (*string, bool)`

GetCustomerStateOk returns a tuple with the CustomerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerState

`func (o *Subscription) SetCustomerState(v string)`

SetCustomerState sets CustomerState field to given value.

### HasCustomerState

`func (o *Subscription) HasCustomerState() bool`

HasCustomerState returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *Subscription) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *Subscription) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *Subscription) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *Subscription) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetShopName

`func (o *Subscription) GetShopName() string`

GetShopName returns the ShopName field if non-nil, zero value otherwise.

### GetShopNameOk

`func (o *Subscription) GetShopNameOk() (*string, bool)`

GetShopNameOk returns a tuple with the ShopName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopName

`func (o *Subscription) SetShopName(v string)`

SetShopName sets ShopName field to given value.

### HasShopName

`func (o *Subscription) HasShopName() bool`

HasShopName returns a boolean if a field has been set.

### GetInvoices

`func (o *Subscription) GetInvoices() []LegacySubscriptionInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Subscription) GetInvoicesOk() (*[]LegacySubscriptionInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Subscription) SetInvoices(v []LegacySubscriptionInvoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Subscription) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetApprovedAddress

`func (o *Subscription) GetApprovedAddress() ApprovedAddress`

GetApprovedAddress returns the ApprovedAddress field if non-nil, zero value otherwise.

### GetApprovedAddressOk

`func (o *Subscription) GetApprovedAddressOk() (*ApprovedAddress, bool)`

GetApprovedAddressOk returns a tuple with the ApprovedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAddress

`func (o *Subscription) SetApprovedAddress(v ApprovedAddress)`

SetApprovedAddress sets ApprovedAddress field to given value.

### HasApprovedAddress

`func (o *Subscription) HasApprovedAddress() bool`

HasApprovedAddress returns a boolean if a field has been set.

### GetGatewaysAvailable

`func (o *Subscription) GetGatewaysAvailable() []Gateway`

GetGatewaysAvailable returns the GatewaysAvailable field if non-nil, zero value otherwise.

### GetGatewaysAvailableOk

`func (o *Subscription) GetGatewaysAvailableOk() (*[]Gateway, bool)`

GetGatewaysAvailableOk returns a tuple with the GatewaysAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaysAvailable

`func (o *Subscription) SetGatewaysAvailable(v []Gateway)`

SetGatewaysAvailable sets GatewaysAvailable field to given value.

### HasGatewaysAvailable

`func (o *Subscription) HasGatewaysAvailable() bool`

HasGatewaysAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


