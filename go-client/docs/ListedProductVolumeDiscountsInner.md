# ListedProductVolumeDiscountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Whether the discount value is a percentage or a fixed amount. | [optional] 
**Value** | Pointer to **int32** | Value of a percentage or fixed discount. | [optional] 
**Quantity** | Pointer to **int32** | How much of this product needs to be purchased in order to be eligible for this volume discount. | [optional] 

## Methods

### NewListedProductVolumeDiscountsInner

`func NewListedProductVolumeDiscountsInner() *ListedProductVolumeDiscountsInner`

NewListedProductVolumeDiscountsInner instantiates a new ListedProductVolumeDiscountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedProductVolumeDiscountsInnerWithDefaults

`func NewListedProductVolumeDiscountsInnerWithDefaults() *ListedProductVolumeDiscountsInner`

NewListedProductVolumeDiscountsInnerWithDefaults instantiates a new ListedProductVolumeDiscountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListedProductVolumeDiscountsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListedProductVolumeDiscountsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListedProductVolumeDiscountsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListedProductVolumeDiscountsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ListedProductVolumeDiscountsInner) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListedProductVolumeDiscountsInner) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListedProductVolumeDiscountsInner) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ListedProductVolumeDiscountsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetQuantity

`func (o *ListedProductVolumeDiscountsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListedProductVolumeDiscountsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListedProductVolumeDiscountsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ListedProductVolumeDiscountsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


