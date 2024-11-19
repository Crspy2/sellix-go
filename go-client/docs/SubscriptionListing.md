# SubscriptionListing

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
**Origin** | Pointer to **string** | Origin of the subscription. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**CouponId** | Pointer to **string** | If a coupon has been applied, its ID. | [optional] 
**CurrentPeriodEnd** | Pointer to **int32** | When the current billing period will end. | [optional] 
**UpcomingEmail1WeekSent** | Pointer to **bool** | Whether or not the email for an upcoming renewal has been sent. | [optional] 
**TrialPeriodEndingEmailSent** | Pointer to **bool** | Whether or not the email for the trial period expiring has been sent. | [optional] 
**RenewalInvoiceCreated** | Pointer to **bool** | If true, the renewal invoice has already been created. | [optional] 
**PaymentLinkId** | Pointer to **string** | The ID of the payment link used to pay the initial payment for this subscription | [optional] 
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

## Methods

### NewSubscriptionListing

`func NewSubscriptionListing() *SubscriptionListing`

NewSubscriptionListing instantiates a new SubscriptionListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListingWithDefaults

`func NewSubscriptionListingWithDefaults() *SubscriptionListing`

NewSubscriptionListingWithDefaults instantiates a new SubscriptionListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductSubscriptionId

`func (o *SubscriptionListing) GetProductSubscriptionId() string`

GetProductSubscriptionId returns the ProductSubscriptionId field if non-nil, zero value otherwise.

### GetProductSubscriptionIdOk

`func (o *SubscriptionListing) GetProductSubscriptionIdOk() (*string, bool)`

GetProductSubscriptionIdOk returns a tuple with the ProductSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionId

`func (o *SubscriptionListing) SetProductSubscriptionId(v string)`

SetProductSubscriptionId sets ProductSubscriptionId field to given value.

### HasProductSubscriptionId

`func (o *SubscriptionListing) HasProductSubscriptionId() bool`

HasProductSubscriptionId returns a boolean if a field has been set.

### GetShopId

`func (o *SubscriptionListing) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *SubscriptionListing) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *SubscriptionListing) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *SubscriptionListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductId

`func (o *SubscriptionListing) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionListing) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionListing) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SubscriptionListing) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductPlanId

`func (o *SubscriptionListing) GetProductPlanId() string`

GetProductPlanId returns the ProductPlanId field if non-nil, zero value otherwise.

### GetProductPlanIdOk

`func (o *SubscriptionListing) GetProductPlanIdOk() (*string, bool)`

GetProductPlanIdOk returns a tuple with the ProductPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanId

`func (o *SubscriptionListing) SetProductPlanId(v string)`

SetProductPlanId sets ProductPlanId field to given value.

### HasProductPlanId

`func (o *SubscriptionListing) HasProductPlanId() bool`

HasProductPlanId returns a boolean if a field has been set.

### GetCustomerId

`func (o *SubscriptionListing) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionListing) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionListing) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *SubscriptionListing) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetType

`func (o *SubscriptionListing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionListing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionListing) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionListing) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriptionListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriptionListing) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionListing) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionListing) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriptionListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionListing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionListing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGateway

`func (o *SubscriptionListing) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SubscriptionListing) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SubscriptionListing) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SubscriptionListing) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetOrigin

`func (o *SubscriptionListing) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *SubscriptionListing) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *SubscriptionListing) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *SubscriptionListing) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCustomFields

`func (o *SubscriptionListing) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *SubscriptionListing) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *SubscriptionListing) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *SubscriptionListing) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCouponId

`func (o *SubscriptionListing) GetCouponId() string`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *SubscriptionListing) GetCouponIdOk() (*string, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *SubscriptionListing) SetCouponId(v string)`

SetCouponId sets CouponId field to given value.

### HasCouponId

`func (o *SubscriptionListing) HasCouponId() bool`

HasCouponId returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *SubscriptionListing) GetCurrentPeriodEnd() int32`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionListing) GetCurrentPeriodEndOk() (*int32, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionListing) SetCurrentPeriodEnd(v int32)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *SubscriptionListing) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetUpcomingEmail1WeekSent

`func (o *SubscriptionListing) GetUpcomingEmail1WeekSent() bool`

GetUpcomingEmail1WeekSent returns the UpcomingEmail1WeekSent field if non-nil, zero value otherwise.

### GetUpcomingEmail1WeekSentOk

`func (o *SubscriptionListing) GetUpcomingEmail1WeekSentOk() (*bool, bool)`

GetUpcomingEmail1WeekSentOk returns a tuple with the UpcomingEmail1WeekSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingEmail1WeekSent

`func (o *SubscriptionListing) SetUpcomingEmail1WeekSent(v bool)`

SetUpcomingEmail1WeekSent sets UpcomingEmail1WeekSent field to given value.

### HasUpcomingEmail1WeekSent

`func (o *SubscriptionListing) HasUpcomingEmail1WeekSent() bool`

HasUpcomingEmail1WeekSent returns a boolean if a field has been set.

### GetTrialPeriodEndingEmailSent

`func (o *SubscriptionListing) GetTrialPeriodEndingEmailSent() bool`

GetTrialPeriodEndingEmailSent returns the TrialPeriodEndingEmailSent field if non-nil, zero value otherwise.

### GetTrialPeriodEndingEmailSentOk

`func (o *SubscriptionListing) GetTrialPeriodEndingEmailSentOk() (*bool, bool)`

GetTrialPeriodEndingEmailSentOk returns a tuple with the TrialPeriodEndingEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodEndingEmailSent

`func (o *SubscriptionListing) SetTrialPeriodEndingEmailSent(v bool)`

SetTrialPeriodEndingEmailSent sets TrialPeriodEndingEmailSent field to given value.

### HasTrialPeriodEndingEmailSent

`func (o *SubscriptionListing) HasTrialPeriodEndingEmailSent() bool`

HasTrialPeriodEndingEmailSent returns a boolean if a field has been set.

### GetRenewalInvoiceCreated

`func (o *SubscriptionListing) GetRenewalInvoiceCreated() bool`

GetRenewalInvoiceCreated returns the RenewalInvoiceCreated field if non-nil, zero value otherwise.

### GetRenewalInvoiceCreatedOk

`func (o *SubscriptionListing) GetRenewalInvoiceCreatedOk() (*bool, bool)`

GetRenewalInvoiceCreatedOk returns a tuple with the RenewalInvoiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalInvoiceCreated

`func (o *SubscriptionListing) SetRenewalInvoiceCreated(v bool)`

SetRenewalInvoiceCreated sets RenewalInvoiceCreated field to given value.

### HasRenewalInvoiceCreated

`func (o *SubscriptionListing) HasRenewalInvoiceCreated() bool`

HasRenewalInvoiceCreated returns a boolean if a field has been set.

### GetPaymentLinkId

`func (o *SubscriptionListing) GetPaymentLinkId() string`

GetPaymentLinkId returns the PaymentLinkId field if non-nil, zero value otherwise.

### GetPaymentLinkIdOk

`func (o *SubscriptionListing) GetPaymentLinkIdOk() (*string, bool)`

GetPaymentLinkIdOk returns a tuple with the PaymentLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLinkId

`func (o *SubscriptionListing) SetPaymentLinkId(v string)`

SetPaymentLinkId sets PaymentLinkId field to given value.

### HasPaymentLinkId

`func (o *SubscriptionListing) HasPaymentLinkId() bool`

HasPaymentLinkId returns a boolean if a field has been set.

### GetStatusDetails

`func (o *SubscriptionListing) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *SubscriptionListing) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *SubscriptionListing) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *SubscriptionListing) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetCanceledAt

`func (o *SubscriptionListing) GetCanceledAt() int32`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SubscriptionListing) GetCanceledAtOk() (*int32, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SubscriptionListing) SetCanceledAt(v int32)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *SubscriptionListing) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetStripeCustomerId

`func (o *SubscriptionListing) GetStripeCustomerId() string`

GetStripeCustomerId returns the StripeCustomerId field if non-nil, zero value otherwise.

### GetStripeCustomerIdOk

`func (o *SubscriptionListing) GetStripeCustomerIdOk() (*string, bool)`

GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCustomerId

`func (o *SubscriptionListing) SetStripeCustomerId(v string)`

SetStripeCustomerId sets StripeCustomerId field to given value.

### HasStripeCustomerId

`func (o *SubscriptionListing) HasStripeCustomerId() bool`

HasStripeCustomerId returns a boolean if a field has been set.

### GetStripeAccount

`func (o *SubscriptionListing) GetStripeAccount() string`

GetStripeAccount returns the StripeAccount field if non-nil, zero value otherwise.

### GetStripeAccountOk

`func (o *SubscriptionListing) GetStripeAccountOk() (*string, bool)`

GetStripeAccountOk returns a tuple with the StripeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccount

`func (o *SubscriptionListing) SetStripeAccount(v string)`

SetStripeAccount sets StripeAccount field to given value.

### HasStripeAccount

`func (o *SubscriptionListing) HasStripeAccount() bool`

HasStripeAccount returns a boolean if a field has been set.

### GetStripeSubscriptionId

`func (o *SubscriptionListing) GetStripeSubscriptionId() string`

GetStripeSubscriptionId returns the StripeSubscriptionId field if non-nil, zero value otherwise.

### GetStripeSubscriptionIdOk

`func (o *SubscriptionListing) GetStripeSubscriptionIdOk() (*string, bool)`

GetStripeSubscriptionIdOk returns a tuple with the StripeSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionId

`func (o *SubscriptionListing) SetStripeSubscriptionId(v string)`

SetStripeSubscriptionId sets StripeSubscriptionId field to given value.

### HasStripeSubscriptionId

`func (o *SubscriptionListing) HasStripeSubscriptionId() bool`

HasStripeSubscriptionId returns a boolean if a field has been set.

### GetPaypalAccount

`func (o *SubscriptionListing) GetPaypalAccount() string`

GetPaypalAccount returns the PaypalAccount field if non-nil, zero value otherwise.

### GetPaypalAccountOk

`func (o *SubscriptionListing) GetPaypalAccountOk() (*string, bool)`

GetPaypalAccountOk returns a tuple with the PaypalAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalAccount

`func (o *SubscriptionListing) SetPaypalAccount(v string)`

SetPaypalAccount sets PaypalAccount field to given value.

### HasPaypalAccount

`func (o *SubscriptionListing) HasPaypalAccount() bool`

HasPaypalAccount returns a boolean if a field has been set.

### GetPaypalSubscriptionId

`func (o *SubscriptionListing) GetPaypalSubscriptionId() string`

GetPaypalSubscriptionId returns the PaypalSubscriptionId field if non-nil, zero value otherwise.

### GetPaypalSubscriptionIdOk

`func (o *SubscriptionListing) GetPaypalSubscriptionIdOk() (*string, bool)`

GetPaypalSubscriptionIdOk returns a tuple with the PaypalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalSubscriptionId

`func (o *SubscriptionListing) SetPaypalSubscriptionId(v string)`

SetPaypalSubscriptionId sets PaypalSubscriptionId field to given value.

### HasPaypalSubscriptionId

`func (o *SubscriptionListing) HasPaypalSubscriptionId() bool`

HasPaypalSubscriptionId returns a boolean if a field has been set.

### GetProductTitle

`func (o *SubscriptionListing) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *SubscriptionListing) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *SubscriptionListing) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *SubscriptionListing) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetProductPlanTitle

`func (o *SubscriptionListing) GetProductPlanTitle() string`

GetProductPlanTitle returns the ProductPlanTitle field if non-nil, zero value otherwise.

### GetProductPlanTitleOk

`func (o *SubscriptionListing) GetProductPlanTitleOk() (*string, bool)`

GetProductPlanTitleOk returns a tuple with the ProductPlanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPlanTitle

`func (o *SubscriptionListing) SetProductPlanTitle(v string)`

SetProductPlanTitle sets ProductPlanTitle field to given value.

### HasProductPlanTitle

`func (o *SubscriptionListing) HasProductPlanTitle() bool`

HasProductPlanTitle returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *SubscriptionListing) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *SubscriptionListing) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *SubscriptionListing) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *SubscriptionListing) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetCustomerName

`func (o *SubscriptionListing) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *SubscriptionListing) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *SubscriptionListing) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *SubscriptionListing) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerSurname

`func (o *SubscriptionListing) GetCustomerSurname() string`

GetCustomerSurname returns the CustomerSurname field if non-nil, zero value otherwise.

### GetCustomerSurnameOk

`func (o *SubscriptionListing) GetCustomerSurnameOk() (*string, bool)`

GetCustomerSurnameOk returns a tuple with the CustomerSurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSurname

`func (o *SubscriptionListing) SetCustomerSurname(v string)`

SetCustomerSurname sets CustomerSurname field to given value.

### HasCustomerSurname

`func (o *SubscriptionListing) HasCustomerSurname() bool`

HasCustomerSurname returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *SubscriptionListing) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *SubscriptionListing) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *SubscriptionListing) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *SubscriptionListing) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerPhoneCountryCode

`func (o *SubscriptionListing) GetCustomerPhoneCountryCode() string`

GetCustomerPhoneCountryCode returns the CustomerPhoneCountryCode field if non-nil, zero value otherwise.

### GetCustomerPhoneCountryCodeOk

`func (o *SubscriptionListing) GetCustomerPhoneCountryCodeOk() (*string, bool)`

GetCustomerPhoneCountryCodeOk returns a tuple with the CustomerPhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhoneCountryCode

`func (o *SubscriptionListing) SetCustomerPhoneCountryCode(v string)`

SetCustomerPhoneCountryCode sets CustomerPhoneCountryCode field to given value.

### HasCustomerPhoneCountryCode

`func (o *SubscriptionListing) HasCustomerPhoneCountryCode() bool`

HasCustomerPhoneCountryCode returns a boolean if a field has been set.

### GetCustomerCountryCode

`func (o *SubscriptionListing) GetCustomerCountryCode() string`

GetCustomerCountryCode returns the CustomerCountryCode field if non-nil, zero value otherwise.

### GetCustomerCountryCodeOk

`func (o *SubscriptionListing) GetCustomerCountryCodeOk() (*string, bool)`

GetCustomerCountryCodeOk returns a tuple with the CustomerCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCountryCode

`func (o *SubscriptionListing) SetCustomerCountryCode(v string)`

SetCustomerCountryCode sets CustomerCountryCode field to given value.

### HasCustomerCountryCode

`func (o *SubscriptionListing) HasCustomerCountryCode() bool`

HasCustomerCountryCode returns a boolean if a field has been set.

### GetCustomerStreetAddress

`func (o *SubscriptionListing) GetCustomerStreetAddress() string`

GetCustomerStreetAddress returns the CustomerStreetAddress field if non-nil, zero value otherwise.

### GetCustomerStreetAddressOk

`func (o *SubscriptionListing) GetCustomerStreetAddressOk() (*string, bool)`

GetCustomerStreetAddressOk returns a tuple with the CustomerStreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerStreetAddress

`func (o *SubscriptionListing) SetCustomerStreetAddress(v string)`

SetCustomerStreetAddress sets CustomerStreetAddress field to given value.

### HasCustomerStreetAddress

`func (o *SubscriptionListing) HasCustomerStreetAddress() bool`

HasCustomerStreetAddress returns a boolean if a field has been set.

### GetCustomerAdditionalAddressInfo

`func (o *SubscriptionListing) GetCustomerAdditionalAddressInfo() string`

GetCustomerAdditionalAddressInfo returns the CustomerAdditionalAddressInfo field if non-nil, zero value otherwise.

### GetCustomerAdditionalAddressInfoOk

`func (o *SubscriptionListing) GetCustomerAdditionalAddressInfoOk() (*string, bool)`

GetCustomerAdditionalAddressInfoOk returns a tuple with the CustomerAdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAdditionalAddressInfo

`func (o *SubscriptionListing) SetCustomerAdditionalAddressInfo(v string)`

SetCustomerAdditionalAddressInfo sets CustomerAdditionalAddressInfo field to given value.

### HasCustomerAdditionalAddressInfo

`func (o *SubscriptionListing) HasCustomerAdditionalAddressInfo() bool`

HasCustomerAdditionalAddressInfo returns a boolean if a field has been set.

### GetCustomerCity

`func (o *SubscriptionListing) GetCustomerCity() string`

GetCustomerCity returns the CustomerCity field if non-nil, zero value otherwise.

### GetCustomerCityOk

`func (o *SubscriptionListing) GetCustomerCityOk() (*string, bool)`

GetCustomerCityOk returns a tuple with the CustomerCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCity

`func (o *SubscriptionListing) SetCustomerCity(v string)`

SetCustomerCity sets CustomerCity field to given value.

### HasCustomerCity

`func (o *SubscriptionListing) HasCustomerCity() bool`

HasCustomerCity returns a boolean if a field has been set.

### GetCustomerPostalCode

`func (o *SubscriptionListing) GetCustomerPostalCode() string`

GetCustomerPostalCode returns the CustomerPostalCode field if non-nil, zero value otherwise.

### GetCustomerPostalCodeOk

`func (o *SubscriptionListing) GetCustomerPostalCodeOk() (*string, bool)`

GetCustomerPostalCodeOk returns a tuple with the CustomerPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPostalCode

`func (o *SubscriptionListing) SetCustomerPostalCode(v string)`

SetCustomerPostalCode sets CustomerPostalCode field to given value.

### HasCustomerPostalCode

`func (o *SubscriptionListing) HasCustomerPostalCode() bool`

HasCustomerPostalCode returns a boolean if a field has been set.

### GetCustomerState

`func (o *SubscriptionListing) GetCustomerState() string`

GetCustomerState returns the CustomerState field if non-nil, zero value otherwise.

### GetCustomerStateOk

`func (o *SubscriptionListing) GetCustomerStateOk() (*string, bool)`

GetCustomerStateOk returns a tuple with the CustomerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerState

`func (o *SubscriptionListing) SetCustomerState(v string)`

SetCustomerState sets CustomerState field to given value.

### HasCustomerState

`func (o *SubscriptionListing) HasCustomerState() bool`

HasCustomerState returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *SubscriptionListing) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *SubscriptionListing) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *SubscriptionListing) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *SubscriptionListing) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


