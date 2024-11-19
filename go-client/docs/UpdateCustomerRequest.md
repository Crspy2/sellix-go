# UpdateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewUpdateCustomerRequest

`func NewUpdateCustomerRequest() *UpdateCustomerRequest`

NewUpdateCustomerRequest instantiates a new UpdateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerRequestWithDefaults

`func NewUpdateCustomerRequestWithDefaults() *UpdateCustomerRequest`

NewUpdateCustomerRequestWithDefaults instantiates a new UpdateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCustomerRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *UpdateCustomerRequest) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UpdateCustomerRequest) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UpdateCustomerRequest) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *UpdateCustomerRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateCustomerRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateCustomerRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateCustomerRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateCustomerRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneCountryCode

`func (o *UpdateCustomerRequest) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *UpdateCustomerRequest) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *UpdateCustomerRequest) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCode

`func (o *UpdateCustomerRequest) HasPhoneCountryCode() bool`

HasPhoneCountryCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *UpdateCustomerRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UpdateCustomerRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UpdateCustomerRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UpdateCustomerRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UpdateCustomerRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UpdateCustomerRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UpdateCustomerRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UpdateCustomerRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAdditionalAddressInfo

`func (o *UpdateCustomerRequest) GetAdditionalAddressInfo() string`

GetAdditionalAddressInfo returns the AdditionalAddressInfo field if non-nil, zero value otherwise.

### GetAdditionalAddressInfoOk

`func (o *UpdateCustomerRequest) GetAdditionalAddressInfoOk() (*string, bool)`

GetAdditionalAddressInfoOk returns a tuple with the AdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAddressInfo

`func (o *UpdateCustomerRequest) SetAdditionalAddressInfo(v string)`

SetAdditionalAddressInfo sets AdditionalAddressInfo field to given value.

### HasAdditionalAddressInfo

`func (o *UpdateCustomerRequest) HasAdditionalAddressInfo() bool`

HasAdditionalAddressInfo returns a boolean if a field has been set.

### GetCity

`func (o *UpdateCustomerRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateCustomerRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateCustomerRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateCustomerRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *UpdateCustomerRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UpdateCustomerRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UpdateCustomerRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UpdateCustomerRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *UpdateCustomerRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateCustomerRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateCustomerRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateCustomerRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


