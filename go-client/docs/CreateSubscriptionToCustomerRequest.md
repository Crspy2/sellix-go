# CreateSubscriptionToCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | ID of the subscription product. | 
**CouponCode** | Pointer to **string** | Coupon code for this subscription. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | key-value JSON having as key the custom field name and as value the custom field value inserted by the customer. Custom fields can both be used as inputs from the customers but also as metadata for invoices, letting you pass hidden fields for internal referencing. | [optional] 
**CustomerId** | **string** | ID of the store customer. | 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 

## Methods

### NewCreateSubscriptionToCustomerRequest

`func NewCreateSubscriptionToCustomerRequest(productId string, customerId string, ) *CreateSubscriptionToCustomerRequest`

NewCreateSubscriptionToCustomerRequest instantiates a new CreateSubscriptionToCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionToCustomerRequestWithDefaults

`func NewCreateSubscriptionToCustomerRequestWithDefaults() *CreateSubscriptionToCustomerRequest`

NewCreateSubscriptionToCustomerRequestWithDefaults instantiates a new CreateSubscriptionToCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateSubscriptionToCustomerRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateSubscriptionToCustomerRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateSubscriptionToCustomerRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetCouponCode

`func (o *CreateSubscriptionToCustomerRequest) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *CreateSubscriptionToCustomerRequest) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *CreateSubscriptionToCustomerRequest) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *CreateSubscriptionToCustomerRequest) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateSubscriptionToCustomerRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateSubscriptionToCustomerRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateSubscriptionToCustomerRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateSubscriptionToCustomerRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCustomerId

`func (o *CreateSubscriptionToCustomerRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateSubscriptionToCustomerRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateSubscriptionToCustomerRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetGateway

`func (o *CreateSubscriptionToCustomerRequest) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateSubscriptionToCustomerRequest) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateSubscriptionToCustomerRequest) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateSubscriptionToCustomerRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


