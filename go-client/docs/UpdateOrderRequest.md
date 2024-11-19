# UpdateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PaypalApm** | Pointer to **string** | If gateway is PAYPAL, a paypal_apm (PayPal Alternative Payment Method) can be specified. To retrieve the available PayPal APM for a specific customer session, please refer to the PayPal SDK using window.paypal.FUNDING and fundingSource to filter out available methods. You can also use our documentation on how to process white_label payments. | [optional] 
**Name** | Pointer to **string** | The name of the recipient for the order | [optional] 
**Surname** | Pointer to **string** | The surname of the recipient for the order | [optional] 
**AddressLine1** | Pointer to **string** | The address line of the the order | [optional] 
**AddressCity** | Pointer to **string** | The city the address is located in | [optional] 
**AddressCountry** | Pointer to **string** | The ISO country code country the address is located in | [optional] 
**AddressPostalCode** | Pointer to **string** | The zipcode for the address | [optional] 
**AddressState** | Pointer to **string** | The state the address is located | [optional] 
**StripeApm** | Pointer to **string** | Stripe Alternative Payment Method name, such as iDEAL, used if gateway is STRIPE. | [optional] 

## Methods

### NewUpdateOrderRequest

`func NewUpdateOrderRequest() *UpdateOrderRequest`

NewUpdateOrderRequest instantiates a new UpdateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderRequestWithDefaults

`func NewUpdateOrderRequestWithDefaults() *UpdateOrderRequest`

NewUpdateOrderRequestWithDefaults instantiates a new UpdateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *UpdateOrderRequest) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateOrderRequest) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateOrderRequest) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateOrderRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPaypalApm

`func (o *UpdateOrderRequest) GetPaypalApm() string`

GetPaypalApm returns the PaypalApm field if non-nil, zero value otherwise.

### GetPaypalApmOk

`func (o *UpdateOrderRequest) GetPaypalApmOk() (*string, bool)`

GetPaypalApmOk returns a tuple with the PaypalApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalApm

`func (o *UpdateOrderRequest) SetPaypalApm(v string)`

SetPaypalApm sets PaypalApm field to given value.

### HasPaypalApm

`func (o *UpdateOrderRequest) HasPaypalApm() bool`

HasPaypalApm returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *UpdateOrderRequest) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UpdateOrderRequest) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UpdateOrderRequest) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *UpdateOrderRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetAddressLine1

`func (o *UpdateOrderRequest) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *UpdateOrderRequest) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *UpdateOrderRequest) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *UpdateOrderRequest) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressCity

`func (o *UpdateOrderRequest) GetAddressCity() string`

GetAddressCity returns the AddressCity field if non-nil, zero value otherwise.

### GetAddressCityOk

`func (o *UpdateOrderRequest) GetAddressCityOk() (*string, bool)`

GetAddressCityOk returns a tuple with the AddressCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCity

`func (o *UpdateOrderRequest) SetAddressCity(v string)`

SetAddressCity sets AddressCity field to given value.

### HasAddressCity

`func (o *UpdateOrderRequest) HasAddressCity() bool`

HasAddressCity returns a boolean if a field has been set.

### GetAddressCountry

`func (o *UpdateOrderRequest) GetAddressCountry() string`

GetAddressCountry returns the AddressCountry field if non-nil, zero value otherwise.

### GetAddressCountryOk

`func (o *UpdateOrderRequest) GetAddressCountryOk() (*string, bool)`

GetAddressCountryOk returns a tuple with the AddressCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCountry

`func (o *UpdateOrderRequest) SetAddressCountry(v string)`

SetAddressCountry sets AddressCountry field to given value.

### HasAddressCountry

`func (o *UpdateOrderRequest) HasAddressCountry() bool`

HasAddressCountry returns a boolean if a field has been set.

### GetAddressPostalCode

`func (o *UpdateOrderRequest) GetAddressPostalCode() string`

GetAddressPostalCode returns the AddressPostalCode field if non-nil, zero value otherwise.

### GetAddressPostalCodeOk

`func (o *UpdateOrderRequest) GetAddressPostalCodeOk() (*string, bool)`

GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPostalCode

`func (o *UpdateOrderRequest) SetAddressPostalCode(v string)`

SetAddressPostalCode sets AddressPostalCode field to given value.

### HasAddressPostalCode

`func (o *UpdateOrderRequest) HasAddressPostalCode() bool`

HasAddressPostalCode returns a boolean if a field has been set.

### GetAddressState

`func (o *UpdateOrderRequest) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *UpdateOrderRequest) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *UpdateOrderRequest) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *UpdateOrderRequest) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetStripeApm

`func (o *UpdateOrderRequest) GetStripeApm() string`

GetStripeApm returns the StripeApm field if non-nil, zero value otherwise.

### GetStripeApmOk

`func (o *UpdateOrderRequest) GetStripeApmOk() (*string, bool)`

GetStripeApmOk returns a tuple with the StripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeApm

`func (o *UpdateOrderRequest) SetStripeApm(v string)`

SetStripeApm sets StripeApm field to given value.

### HasStripeApm

`func (o *UpdateOrderRequest) HasStripeApm() bool`

HasStripeApm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


