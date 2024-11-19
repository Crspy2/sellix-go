# CreateSubscriptionv2RequestTaxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The two letter ISO country code for your country. This is used to determine the correct tax rate to apply to the subscription. | [optional] 
**City** | Pointer to **string** | The city your business is located in. This is used to determine the correct tax rate to apply to the subscription. | [optional] 

## Methods

### NewCreateSubscriptionv2RequestTaxData

`func NewCreateSubscriptionv2RequestTaxData() *CreateSubscriptionv2RequestTaxData`

NewCreateSubscriptionv2RequestTaxData instantiates a new CreateSubscriptionv2RequestTaxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionv2RequestTaxDataWithDefaults

`func NewCreateSubscriptionv2RequestTaxDataWithDefaults() *CreateSubscriptionv2RequestTaxData`

NewCreateSubscriptionv2RequestTaxDataWithDefaults instantiates a new CreateSubscriptionv2RequestTaxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CreateSubscriptionv2RequestTaxData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateSubscriptionv2RequestTaxData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateSubscriptionv2RequestTaxData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CreateSubscriptionv2RequestTaxData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCity

`func (o *CreateSubscriptionv2RequestTaxData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateSubscriptionv2RequestTaxData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateSubscriptionv2RequestTaxData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateSubscriptionv2RequestTaxData) HasCity() bool`

HasCity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


