# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Customer ID | [optional] 
**ShopId** | Pointer to **float32** | The shop ID to which this customer belongs. | [optional] 
**Email** | Pointer to **string** | Customer email | [optional] 
**Name** | Pointer to **string** | Customer name | [optional] 
**Surname** | Pointer to **string** | Customer surname | [optional] 
**Phone** | Pointer to **string** | Customer phone | [optional] 
**PhoneCountryCode** | Pointer to **string** | Customer phone country code | [optional] 
**CountryCode** | Pointer to **string** | Customer country code | [optional] 
**StreetAddress** | Pointer to **string** | Customer street address | [optional] 
**AdditionalAddressInfo** | Pointer to **string** | Customer street address additional info | [optional] 
**City** | Pointer to **string** | Customer city | [optional] 
**PostalCode** | Pointer to **string** | Customer postal code | [optional] 
**State** | Pointer to **string** | Customer state | [optional] 
**AffiliateRevenue** | Pointer to **string** | Revenue customer has recieved from the affiliate program | [optional] 
**AffiliateRevenueCurrency** | Pointer to **string** | The currency the customer gets reimbursed in for the affiliate program | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the customer. | [optional] 
**Subscriptions** | Pointer to [**[]SubscriptionListing**](SubscriptionListing.md) |  | [optional] 
**Invoices** | Pointer to [**[]MiniInvoice**](MiniInvoice.md) |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *Customer) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Customer) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Customer) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Customer) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Customer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *Customer) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Customer) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Customer) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Customer) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Customer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneCountryCode

`func (o *Customer) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *Customer) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *Customer) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCode

`func (o *Customer) HasPhoneCountryCode() bool`

HasPhoneCountryCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *Customer) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Customer) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Customer) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Customer) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *Customer) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *Customer) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *Customer) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *Customer) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAdditionalAddressInfo

`func (o *Customer) GetAdditionalAddressInfo() string`

GetAdditionalAddressInfo returns the AdditionalAddressInfo field if non-nil, zero value otherwise.

### GetAdditionalAddressInfoOk

`func (o *Customer) GetAdditionalAddressInfoOk() (*string, bool)`

GetAdditionalAddressInfoOk returns a tuple with the AdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAddressInfo

`func (o *Customer) SetAdditionalAddressInfo(v string)`

SetAdditionalAddressInfo sets AdditionalAddressInfo field to given value.

### HasAdditionalAddressInfo

`func (o *Customer) HasAdditionalAddressInfo() bool`

HasAdditionalAddressInfo returns a boolean if a field has been set.

### GetCity

`func (o *Customer) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Customer) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Customer) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Customer) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *Customer) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Customer) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Customer) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Customer) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *Customer) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Customer) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Customer) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Customer) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAffiliateRevenue

`func (o *Customer) GetAffiliateRevenue() string`

GetAffiliateRevenue returns the AffiliateRevenue field if non-nil, zero value otherwise.

### GetAffiliateRevenueOk

`func (o *Customer) GetAffiliateRevenueOk() (*string, bool)`

GetAffiliateRevenueOk returns a tuple with the AffiliateRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenue

`func (o *Customer) SetAffiliateRevenue(v string)`

SetAffiliateRevenue sets AffiliateRevenue field to given value.

### HasAffiliateRevenue

`func (o *Customer) HasAffiliateRevenue() bool`

HasAffiliateRevenue returns a boolean if a field has been set.

### GetAffiliateRevenueCurrency

`func (o *Customer) GetAffiliateRevenueCurrency() string`

GetAffiliateRevenueCurrency returns the AffiliateRevenueCurrency field if non-nil, zero value otherwise.

### GetAffiliateRevenueCurrencyOk

`func (o *Customer) GetAffiliateRevenueCurrencyOk() (*string, bool)`

GetAffiliateRevenueCurrencyOk returns a tuple with the AffiliateRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueCurrency

`func (o *Customer) SetAffiliateRevenueCurrency(v string)`

SetAffiliateRevenueCurrency sets AffiliateRevenueCurrency field to given value.

### HasAffiliateRevenueCurrency

`func (o *Customer) HasAffiliateRevenueCurrency() bool`

HasAffiliateRevenueCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Customer) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Customer) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Customer) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Customer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSubscriptions

`func (o *Customer) GetSubscriptions() []SubscriptionListing`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *Customer) GetSubscriptionsOk() (*[]SubscriptionListing, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *Customer) SetSubscriptions(v []SubscriptionListing)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *Customer) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetInvoices

`func (o *Customer) GetInvoices() []MiniInvoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Customer) GetInvoicesOk() (*[]MiniInvoice, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Customer) SetInvoices(v []MiniInvoice)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Customer) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


