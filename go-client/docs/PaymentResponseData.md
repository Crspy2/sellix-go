# PaymentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Sellix hosted payment page. | 
**Uniqid** | **string** | Unique ID of the invoice created for the payment. | 
**UrlBranded** | Pointer to **string** | Sellix hosted payment page with url unique to your shop. | [optional] 

## Methods

### NewPaymentResponseData

`func NewPaymentResponseData(url string, uniqid string, ) *PaymentResponseData`

NewPaymentResponseData instantiates a new PaymentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentResponseDataWithDefaults

`func NewPaymentResponseDataWithDefaults() *PaymentResponseData`

NewPaymentResponseDataWithDefaults instantiates a new PaymentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PaymentResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUniqid

`func (o *PaymentResponseData) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *PaymentResponseData) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *PaymentResponseData) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.


### GetUrlBranded

`func (o *PaymentResponseData) GetUrlBranded() string`

GetUrlBranded returns the UrlBranded field if non-nil, zero value otherwise.

### GetUrlBrandedOk

`func (o *PaymentResponseData) GetUrlBrandedOk() (*string, bool)`

GetUrlBrandedOk returns a tuple with the UrlBranded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBranded

`func (o *PaymentResponseData) SetUrlBranded(v string)`

SetUrlBranded sets UrlBranded field to given value.

### HasUrlBranded

`func (o *PaymentResponseData) HasUrlBranded() bool`

HasUrlBranded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


