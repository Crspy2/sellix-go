# CustomerListing

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
**AffiliatePercentage** | Pointer to **string** | The affiliate percentage awarded to the customer | [optional] 

## Methods

### NewCustomerListing

`func NewCustomerListing() *CustomerListing`

NewCustomerListing instantiates a new CustomerListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListingWithDefaults

`func NewCustomerListingWithDefaults() *CustomerListing`

NewCustomerListingWithDefaults instantiates a new CustomerListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShopId

`func (o *CustomerListing) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *CustomerListing) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *CustomerListing) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *CustomerListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerListing) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerListing) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerListing) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerListing) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CustomerListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *CustomerListing) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CustomerListing) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CustomerListing) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *CustomerListing) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerListing) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerListing) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerListing) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerListing) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneCountryCode

`func (o *CustomerListing) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *CustomerListing) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *CustomerListing) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCode

`func (o *CustomerListing) HasPhoneCountryCode() bool`

HasPhoneCountryCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *CustomerListing) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CustomerListing) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CustomerListing) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CustomerListing) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *CustomerListing) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CustomerListing) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CustomerListing) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CustomerListing) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAdditionalAddressInfo

`func (o *CustomerListing) GetAdditionalAddressInfo() string`

GetAdditionalAddressInfo returns the AdditionalAddressInfo field if non-nil, zero value otherwise.

### GetAdditionalAddressInfoOk

`func (o *CustomerListing) GetAdditionalAddressInfoOk() (*string, bool)`

GetAdditionalAddressInfoOk returns a tuple with the AdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAddressInfo

`func (o *CustomerListing) SetAdditionalAddressInfo(v string)`

SetAdditionalAddressInfo sets AdditionalAddressInfo field to given value.

### HasAdditionalAddressInfo

`func (o *CustomerListing) HasAdditionalAddressInfo() bool`

HasAdditionalAddressInfo returns a boolean if a field has been set.

### GetCity

`func (o *CustomerListing) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerListing) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerListing) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerListing) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerListing) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerListing) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerListing) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomerListing) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CustomerListing) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerListing) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerListing) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerListing) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAffiliateRevenue

`func (o *CustomerListing) GetAffiliateRevenue() string`

GetAffiliateRevenue returns the AffiliateRevenue field if non-nil, zero value otherwise.

### GetAffiliateRevenueOk

`func (o *CustomerListing) GetAffiliateRevenueOk() (*string, bool)`

GetAffiliateRevenueOk returns a tuple with the AffiliateRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenue

`func (o *CustomerListing) SetAffiliateRevenue(v string)`

SetAffiliateRevenue sets AffiliateRevenue field to given value.

### HasAffiliateRevenue

`func (o *CustomerListing) HasAffiliateRevenue() bool`

HasAffiliateRevenue returns a boolean if a field has been set.

### GetAffiliateRevenueCurrency

`func (o *CustomerListing) GetAffiliateRevenueCurrency() string`

GetAffiliateRevenueCurrency returns the AffiliateRevenueCurrency field if non-nil, zero value otherwise.

### GetAffiliateRevenueCurrencyOk

`func (o *CustomerListing) GetAffiliateRevenueCurrencyOk() (*string, bool)`

GetAffiliateRevenueCurrencyOk returns a tuple with the AffiliateRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenueCurrency

`func (o *CustomerListing) SetAffiliateRevenueCurrency(v string)`

SetAffiliateRevenueCurrency sets AffiliateRevenueCurrency field to given value.

### HasAffiliateRevenueCurrency

`func (o *CustomerListing) HasAffiliateRevenueCurrency() bool`

HasAffiliateRevenueCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAffiliatePercentage

`func (o *CustomerListing) GetAffiliatePercentage() string`

GetAffiliatePercentage returns the AffiliatePercentage field if non-nil, zero value otherwise.

### GetAffiliatePercentageOk

`func (o *CustomerListing) GetAffiliatePercentageOk() (*string, bool)`

GetAffiliatePercentageOk returns a tuple with the AffiliatePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliatePercentage

`func (o *CustomerListing) SetAffiliatePercentage(v string)`

SetAffiliatePercentage sets AffiliatePercentage field to given value.

### HasAffiliatePercentage

`func (o *CustomerListing) HasAffiliatePercentage() bool`

HasAffiliatePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


