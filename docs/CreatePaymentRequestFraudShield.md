# CreatePaymentRequestFraudShield

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Customer IP. | [optional] 
**UserAgent** | Pointer to **string** | Customer user agent. | [optional] 
**UserLanguage** | Pointer to **string** | Customer user language. | [optional] 
**MaximumFraudScore** | Pointer to **int32** | The maximum fraud score a customer can have for the payment to be processed. Between 0 and 100, if set to 100 all fraud checks will be skipped and the invoice will be created. | [optional] 

## Methods

### NewCreatePaymentRequestFraudShield

`func NewCreatePaymentRequestFraudShield() *CreatePaymentRequestFraudShield`

NewCreatePaymentRequestFraudShield instantiates a new CreatePaymentRequestFraudShield object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentRequestFraudShieldWithDefaults

`func NewCreatePaymentRequestFraudShieldWithDefaults() *CreatePaymentRequestFraudShield`

NewCreatePaymentRequestFraudShieldWithDefaults instantiates a new CreatePaymentRequestFraudShield object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *CreatePaymentRequestFraudShield) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreatePaymentRequestFraudShield) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreatePaymentRequestFraudShield) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreatePaymentRequestFraudShield) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUserAgent

`func (o *CreatePaymentRequestFraudShield) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CreatePaymentRequestFraudShield) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CreatePaymentRequestFraudShield) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CreatePaymentRequestFraudShield) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserLanguage

`func (o *CreatePaymentRequestFraudShield) GetUserLanguage() string`

GetUserLanguage returns the UserLanguage field if non-nil, zero value otherwise.

### GetUserLanguageOk

`func (o *CreatePaymentRequestFraudShield) GetUserLanguageOk() (*string, bool)`

GetUserLanguageOk returns a tuple with the UserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguage

`func (o *CreatePaymentRequestFraudShield) SetUserLanguage(v string)`

SetUserLanguage sets UserLanguage field to given value.

### HasUserLanguage

`func (o *CreatePaymentRequestFraudShield) HasUserLanguage() bool`

HasUserLanguage returns a boolean if a field has been set.

### GetMaximumFraudScore

`func (o *CreatePaymentRequestFraudShield) GetMaximumFraudScore() int32`

GetMaximumFraudScore returns the MaximumFraudScore field if non-nil, zero value otherwise.

### GetMaximumFraudScoreOk

`func (o *CreatePaymentRequestFraudShield) GetMaximumFraudScoreOk() (*int32, bool)`

GetMaximumFraudScoreOk returns a tuple with the MaximumFraudScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFraudScore

`func (o *CreatePaymentRequestFraudShield) SetMaximumFraudScore(v int32)`

SetMaximumFraudScore sets MaximumFraudScore field to given value.

### HasMaximumFraudScore

`func (o *CreatePaymentRequestFraudShield) HasMaximumFraudScore() bool`

HasMaximumFraudScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


