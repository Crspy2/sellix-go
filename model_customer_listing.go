/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sellix

import (
	"encoding/json"
)

// checks if the CustomerListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerListing{}

// CustomerListing struct for CustomerListing
type CustomerListing struct {
	// Customer ID
	Id *string `json:"id,omitempty"`
	// The shop ID to which this customer belongs.
	ShopId *float32 `json:"shop_id,omitempty"`
	// Customer email
	Email *string `json:"email,omitempty"`
	// Customer name
	Name *string `json:"name,omitempty"`
	// Customer surname
	Surname *string `json:"surname,omitempty"`
	// Customer phone
	Phone *string `json:"phone,omitempty"`
	// Customer phone country code
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	// Customer country code
	CountryCode *string `json:"country_code,omitempty"`
	// Customer street address
	StreetAddress *string `json:"street_address,omitempty"`
	// Customer street address additional info
	AdditionalAddressInfo *string `json:"additional_address_info,omitempty"`
	// Customer city
	City *string `json:"city,omitempty"`
	// Customer postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// Customer state
	State *string `json:"state,omitempty"`
	// Revenue customer has recieved from the affiliate program
	AffiliateRevenue *string `json:"affiliate_revenue,omitempty"`
	// The currency the customer gets reimbursed in for the affiliate program
	AffiliateRevenueCurrency *string `json:"affiliate_revenue_currency,omitempty"`
	// Timestamp for the creation of the customer.
	CreatedAt *int32 `json:"created_at,omitempty"`
	// The affiliate percentage awarded to the customer
	AffiliatePercentage *string `json:"affiliate_percentage,omitempty"`
}

// NewCustomerListing instantiates a new CustomerListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerListing() *CustomerListing {
	this := CustomerListing{}
	return &this
}

// NewCustomerListingWithDefaults instantiates a new CustomerListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerListingWithDefaults() *CustomerListing {
	this := CustomerListing{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerListing) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerListing) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerListing) SetId(v string) {
	o.Id = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *CustomerListing) GetShopId() float32 {
	if o == nil || IsNil(o.ShopId) {
		var ret float32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetShopIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *CustomerListing) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given float32 and assigns it to the ShopId field.
func (o *CustomerListing) SetShopId(v float32) {
	o.ShopId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerListing) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerListing) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerListing) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerListing) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerListing) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerListing) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerListing) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerListing) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerListing) SetSurname(v string) {
	o.Surname = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerListing) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerListing) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerListing) SetPhone(v string) {
	o.Phone = &v
}

// GetPhoneCountryCode returns the PhoneCountryCode field value if set, zero value otherwise.
func (o *CustomerListing) GetPhoneCountryCode() string {
	if o == nil || IsNil(o.PhoneCountryCode) {
		var ret string
		return ret
	}
	return *o.PhoneCountryCode
}

// GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetPhoneCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneCountryCode) {
		return nil, false
	}
	return o.PhoneCountryCode, true
}

// HasPhoneCountryCode returns a boolean if a field has been set.
func (o *CustomerListing) HasPhoneCountryCode() bool {
	if o != nil && !IsNil(o.PhoneCountryCode) {
		return true
	}

	return false
}

// SetPhoneCountryCode gets a reference to the given string and assigns it to the PhoneCountryCode field.
func (o *CustomerListing) SetPhoneCountryCode(v string) {
	o.PhoneCountryCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CustomerListing) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CustomerListing) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CustomerListing) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *CustomerListing) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *CustomerListing) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *CustomerListing) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetAdditionalAddressInfo returns the AdditionalAddressInfo field value if set, zero value otherwise.
func (o *CustomerListing) GetAdditionalAddressInfo() string {
	if o == nil || IsNil(o.AdditionalAddressInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalAddressInfo
}

// GetAdditionalAddressInfoOk returns a tuple with the AdditionalAddressInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetAdditionalAddressInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalAddressInfo) {
		return nil, false
	}
	return o.AdditionalAddressInfo, true
}

// HasAdditionalAddressInfo returns a boolean if a field has been set.
func (o *CustomerListing) HasAdditionalAddressInfo() bool {
	if o != nil && !IsNil(o.AdditionalAddressInfo) {
		return true
	}

	return false
}

// SetAdditionalAddressInfo gets a reference to the given string and assigns it to the AdditionalAddressInfo field.
func (o *CustomerListing) SetAdditionalAddressInfo(v string) {
	o.AdditionalAddressInfo = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerListing) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerListing) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerListing) SetCity(v string) {
	o.City = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CustomerListing) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerListing) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CustomerListing) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerListing) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerListing) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerListing) SetState(v string) {
	o.State = &v
}

// GetAffiliateRevenue returns the AffiliateRevenue field value if set, zero value otherwise.
func (o *CustomerListing) GetAffiliateRevenue() string {
	if o == nil || IsNil(o.AffiliateRevenue) {
		var ret string
		return ret
	}
	return *o.AffiliateRevenue
}

// GetAffiliateRevenueOk returns a tuple with the AffiliateRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetAffiliateRevenueOk() (*string, bool) {
	if o == nil || IsNil(o.AffiliateRevenue) {
		return nil, false
	}
	return o.AffiliateRevenue, true
}

// HasAffiliateRevenue returns a boolean if a field has been set.
func (o *CustomerListing) HasAffiliateRevenue() bool {
	if o != nil && !IsNil(o.AffiliateRevenue) {
		return true
	}

	return false
}

// SetAffiliateRevenue gets a reference to the given string and assigns it to the AffiliateRevenue field.
func (o *CustomerListing) SetAffiliateRevenue(v string) {
	o.AffiliateRevenue = &v
}

// GetAffiliateRevenueCurrency returns the AffiliateRevenueCurrency field value if set, zero value otherwise.
func (o *CustomerListing) GetAffiliateRevenueCurrency() string {
	if o == nil || IsNil(o.AffiliateRevenueCurrency) {
		var ret string
		return ret
	}
	return *o.AffiliateRevenueCurrency
}

// GetAffiliateRevenueCurrencyOk returns a tuple with the AffiliateRevenueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetAffiliateRevenueCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.AffiliateRevenueCurrency) {
		return nil, false
	}
	return o.AffiliateRevenueCurrency, true
}

// HasAffiliateRevenueCurrency returns a boolean if a field has been set.
func (o *CustomerListing) HasAffiliateRevenueCurrency() bool {
	if o != nil && !IsNil(o.AffiliateRevenueCurrency) {
		return true
	}

	return false
}

// SetAffiliateRevenueCurrency gets a reference to the given string and assigns it to the AffiliateRevenueCurrency field.
func (o *CustomerListing) SetAffiliateRevenueCurrency(v string) {
	o.AffiliateRevenueCurrency = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerListing) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerListing) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *CustomerListing) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetAffiliatePercentage returns the AffiliatePercentage field value if set, zero value otherwise.
func (o *CustomerListing) GetAffiliatePercentage() string {
	if o == nil || IsNil(o.AffiliatePercentage) {
		var ret string
		return ret
	}
	return *o.AffiliatePercentage
}

// GetAffiliatePercentageOk returns a tuple with the AffiliatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListing) GetAffiliatePercentageOk() (*string, bool) {
	if o == nil || IsNil(o.AffiliatePercentage) {
		return nil, false
	}
	return o.AffiliatePercentage, true
}

// HasAffiliatePercentage returns a boolean if a field has been set.
func (o *CustomerListing) HasAffiliatePercentage() bool {
	if o != nil && !IsNil(o.AffiliatePercentage) {
		return true
	}

	return false
}

// SetAffiliatePercentage gets a reference to the given string and assigns it to the AffiliatePercentage field.
func (o *CustomerListing) SetAffiliatePercentage(v string) {
	o.AffiliatePercentage = &v
}

func (o CustomerListing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PhoneCountryCode) {
		toSerialize["phone_country_code"] = o.PhoneCountryCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.AdditionalAddressInfo) {
		toSerialize["additional_address_info"] = o.AdditionalAddressInfo
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.AffiliateRevenue) {
		toSerialize["affiliate_revenue"] = o.AffiliateRevenue
	}
	if !IsNil(o.AffiliateRevenueCurrency) {
		toSerialize["affiliate_revenue_currency"] = o.AffiliateRevenueCurrency
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.AffiliatePercentage) {
		toSerialize["affiliate_percentage"] = o.AffiliatePercentage
	}
	return toSerialize, nil
}

type NullableCustomerListing struct {
	value *CustomerListing
	isSet bool
}

func (v NullableCustomerListing) Get() *CustomerListing {
	return v.value
}

func (v *NullableCustomerListing) Set(val *CustomerListing) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerListing) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerListing(val *CustomerListing) *NullableCustomerListing {
	return &NullableCustomerListing{value: val, isSet: true}
}

func (v NullableCustomerListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


