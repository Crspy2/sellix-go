# BundleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the bundle, used as reference across the API. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this bundle belongs. | [optional] 
**Title** | Pointer to **string** | The title of the bundle. | [optional] 
**Products** | Pointer to **string** | The product IDs of the products in the bundle. | [optional] 
**DiscountType** | Pointer to **string** | The type of discount applied to the bundle. | [optional] 
**DiscountAmount** | Pointer to **float32** | The amount of the discount applied to the bundle. | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp, available if the bundle has been updated. | [optional] 

## Methods

### NewBundleConfig

`func NewBundleConfig() *BundleConfig`

NewBundleConfig instantiates a new BundleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleConfigWithDefaults

`func NewBundleConfigWithDefaults() *BundleConfig`

NewBundleConfigWithDefaults instantiates a new BundleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BundleConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleConfig) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BundleConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *BundleConfig) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *BundleConfig) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *BundleConfig) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *BundleConfig) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *BundleConfig) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *BundleConfig) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *BundleConfig) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *BundleConfig) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetTitle

`func (o *BundleConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BundleConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BundleConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BundleConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProducts

`func (o *BundleConfig) GetProducts() string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *BundleConfig) GetProductsOk() (*string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *BundleConfig) SetProducts(v string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *BundleConfig) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetDiscountType

`func (o *BundleConfig) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *BundleConfig) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *BundleConfig) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *BundleConfig) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *BundleConfig) GetDiscountAmount() float32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *BundleConfig) GetDiscountAmountOk() (*float32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *BundleConfig) SetDiscountAmount(v float32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *BundleConfig) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BundleConfig) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BundleConfig) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BundleConfig) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BundleConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


