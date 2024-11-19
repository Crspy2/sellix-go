# UpdateSubscriptionv2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**PaymentMethodUniqid** | Pointer to **string** | The unique id of the customer&#39;s payment method to set for this endpoint. | [optional] 

## Methods

### NewUpdateSubscriptionv2Request

`func NewUpdateSubscriptionv2Request() *UpdateSubscriptionv2Request`

NewUpdateSubscriptionv2Request instantiates a new UpdateSubscriptionv2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionv2RequestWithDefaults

`func NewUpdateSubscriptionv2RequestWithDefaults() *UpdateSubscriptionv2Request`

NewUpdateSubscriptionv2RequestWithDefaults instantiates a new UpdateSubscriptionv2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *UpdateSubscriptionv2Request) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateSubscriptionv2Request) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateSubscriptionv2Request) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateSubscriptionv2Request) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPaymentMethodUniqid

`func (o *UpdateSubscriptionv2Request) GetPaymentMethodUniqid() string`

GetPaymentMethodUniqid returns the PaymentMethodUniqid field if non-nil, zero value otherwise.

### GetPaymentMethodUniqidOk

`func (o *UpdateSubscriptionv2Request) GetPaymentMethodUniqidOk() (*string, bool)`

GetPaymentMethodUniqidOk returns a tuple with the PaymentMethodUniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodUniqid

`func (o *UpdateSubscriptionv2Request) SetPaymentMethodUniqid(v string)`

SetPaymentMethodUniqid sets PaymentMethodUniqid field to given value.

### HasPaymentMethodUniqid

`func (o *UpdateSubscriptionv2Request) HasPaymentMethodUniqid() bool`

HasPaymentMethodUniqid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


