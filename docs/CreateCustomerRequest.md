# CreateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Customer email | 
**Name** | **string** | Customer name | 
**Surname** | **string** | Customer surname | 
**Phone** | Pointer to **string** | Customer phone | [optional] 
**PhoneCountryCode** | Pointer to **string** | Customer phone country code | [optional] 
**CountryCode** | Pointer to **string** | Customer country code | [optional] 
**StreetAddress** | Pointer to **string** | Customer street address | [optional] 
**AdditionalAddressInfo** | Pointer to **string** | Customer street address additional info | [optional] 
**City** | Pointer to **string** | Customer city | [optional] 
**PostalCode** | Pointer to **string** | Customer postal code | [optional] 
**State** | Pointer to **string** | Customer state | [optional] 

## Methods

### NewCreateCustomerRequest

`func NewCreateCustomerRequest(email string, name string, surname string, ) *CreateCustomerRequest`

NewCreateCustomerRequest instantiates a new CreateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerRequestWithDefaults

`func NewCreateCustomerRequestWithDefaults() *CreateCustomerRequest`

NewCreateCustomerRequestWithDefaults instantiates a new CreateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CreateCustomerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSurname

`func (o *CreateCustomerRequest) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CreateCustomerRequest) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CreateCustomerRequest) SetSurname(v string)`

SetSurname sets Surname field to given value.


### GetPhone

`func (o *CreateCustomerRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateCustomerRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateCustomerRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateCustomerRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneCountryCode

`func (o *CreateCustomerRequest) GetPhoneCountryCode() string`

GetPhoneCountryCode returns the PhoneCountryCode field if non-nil, zero value otherwise.

### GetPhoneCountryCodeOk

`func (o *CreateCustomerRequest) GetPhoneCountryCodeOk() (*string, bool)`

GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCode

`func (o *CreateCustomerRequest) SetPhoneCountryCode(v string)`

SetPhoneCountryCode sets PhoneCountryCode field to given value.

### HasPhoneCountryCode

`func (o *CreateCustomerRequest) HasPhoneCountryCode() bool`

HasPhoneCountryCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *CreateCustomerRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateCustomerRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateCustomerRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CreateCustomerRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetStreetAddress

`func (o *CreateCustomerRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *CreateCustomerRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *CreateCustomerRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *CreateCustomerRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetAdditionalAddressInfo

`func (o *CreateCustomerRequest) GetAdditionalAddressInfo() string`

GetAdditionalAddressInfo returns the AdditionalAddressInfo field if non-nil, zero value otherwise.

### GetAdditionalAddressInfoOk

`func (o *CreateCustomerRequest) GetAdditionalAddressInfoOk() (*string, bool)`

GetAdditionalAddressInfoOk returns a tuple with the AdditionalAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAddressInfo

`func (o *CreateCustomerRequest) SetAdditionalAddressInfo(v string)`

SetAdditionalAddressInfo sets AdditionalAddressInfo field to given value.

### HasAdditionalAddressInfo

`func (o *CreateCustomerRequest) HasAdditionalAddressInfo() bool`

HasAdditionalAddressInfo returns a boolean if a field has been set.

### GetCity

`func (o *CreateCustomerRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateCustomerRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateCustomerRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateCustomerRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *CreateCustomerRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CreateCustomerRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CreateCustomerRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CreateCustomerRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *CreateCustomerRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateCustomerRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateCustomerRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateCustomerRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


