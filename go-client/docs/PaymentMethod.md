# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal Sellix ID for the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique identifier used throughout Sellix&#39;s API | [optional] 
**Title** | Pointer to **string** | The title of the stored payment method | [optional] 
**SystemTitle** | Pointer to **string** | The title of the stored payment method used by our system | [optional] 
**ShopId** | Pointer to **int32** | ID of the Sellix shop the customer was created for. | [optional] 
**CustomerId** | Pointer to **string** | Unique ID used within Sellix to identify customers across the API. | [optional] 
**Gateway** | Pointer to [**Gateway**](Gateway.md) |  | [optional] 
**MethodType** | Pointer to **string** | The type of payment method used within the gateway | [optional] 
**Status** | Pointer to **string** | The status of the payment method | [optional] 
**LastConfirmDate** | Pointer to **int32** | The last time the payment method was used stored as a timestamp. | [optional] 
**CreatedAt** | Pointer to **int32** | When the payment method was saved for future reuse stored as a timestamp. | [optional] 
**UpdatedAt** | Pointer to **int32** | When the payment method was last updated stored as a timestamp. | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *PaymentMethod) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *PaymentMethod) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *PaymentMethod) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *PaymentMethod) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetTitle

`func (o *PaymentMethod) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentMethod) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentMethod) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PaymentMethod) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSystemTitle

`func (o *PaymentMethod) GetSystemTitle() string`

GetSystemTitle returns the SystemTitle field if non-nil, zero value otherwise.

### GetSystemTitleOk

`func (o *PaymentMethod) GetSystemTitleOk() (*string, bool)`

GetSystemTitleOk returns a tuple with the SystemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTitle

`func (o *PaymentMethod) SetSystemTitle(v string)`

SetSystemTitle sets SystemTitle field to given value.

### HasSystemTitle

`func (o *PaymentMethod) HasSystemTitle() bool`

HasSystemTitle returns a boolean if a field has been set.

### GetShopId

`func (o *PaymentMethod) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *PaymentMethod) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *PaymentMethod) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *PaymentMethod) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetCustomerId

`func (o *PaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentMethod) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetGateway

`func (o *PaymentMethod) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PaymentMethod) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PaymentMethod) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PaymentMethod) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMethodType

`func (o *PaymentMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *PaymentMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *PaymentMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *PaymentMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentMethod) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethod) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethod) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastConfirmDate

`func (o *PaymentMethod) GetLastConfirmDate() int32`

GetLastConfirmDate returns the LastConfirmDate field if non-nil, zero value otherwise.

### GetLastConfirmDateOk

`func (o *PaymentMethod) GetLastConfirmDateOk() (*int32, bool)`

GetLastConfirmDateOk returns a tuple with the LastConfirmDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfirmDate

`func (o *PaymentMethod) SetLastConfirmDate(v int32)`

SetLastConfirmDate sets LastConfirmDate field to given value.

### HasLastConfirmDate

`func (o *PaymentMethod) HasLastConfirmDate() bool`

HasLastConfirmDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentMethod) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethod) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethod) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentMethod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentMethod) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethod) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethod) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentMethod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


