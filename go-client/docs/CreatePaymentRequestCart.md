# CreatePaymentRequestCart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]CreatePaymentRequestCartProductsInner**](CreatePaymentRequestCartProductsInner.md) |  | [optional] 

## Methods

### NewCreatePaymentRequestCart

`func NewCreatePaymentRequestCart() *CreatePaymentRequestCart`

NewCreatePaymentRequestCart instantiates a new CreatePaymentRequestCart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentRequestCartWithDefaults

`func NewCreatePaymentRequestCartWithDefaults() *CreatePaymentRequestCart`

NewCreatePaymentRequestCartWithDefaults instantiates a new CreatePaymentRequestCart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *CreatePaymentRequestCart) GetProducts() []CreatePaymentRequestCartProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CreatePaymentRequestCart) GetProductsOk() (*[]CreatePaymentRequestCartProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CreatePaymentRequestCart) SetProducts(v []CreatePaymentRequestCartProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *CreatePaymentRequestCart) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


