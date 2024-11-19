# ProductAddonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uniqid** | Pointer to **string** |  | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this addon belongs. | [optional] 
**ProductTypes** | Pointer to **[]string** | The product types for the product addon | [optional] 
**Title** | Pointer to **string** | The title of the product addon | [optional] 
**Description** | Pointer to **string** | The description of the product addon | [optional] 
**Price** | Pointer to **string** | The price of the product addon | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewProductAddonsInner

`func NewProductAddonsInner() *ProductAddonsInner`

NewProductAddonsInner instantiates a new ProductAddonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAddonsInnerWithDefaults

`func NewProductAddonsInnerWithDefaults() *ProductAddonsInner`

NewProductAddonsInnerWithDefaults instantiates a new ProductAddonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductAddonsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductAddonsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductAddonsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductAddonsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ProductAddonsInner) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ProductAddonsInner) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ProductAddonsInner) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ProductAddonsInner) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetShopId

`func (o *ProductAddonsInner) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ProductAddonsInner) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ProductAddonsInner) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ProductAddonsInner) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetProductTypes

`func (o *ProductAddonsInner) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *ProductAddonsInner) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *ProductAddonsInner) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *ProductAddonsInner) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetTitle

`func (o *ProductAddonsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductAddonsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductAddonsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProductAddonsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ProductAddonsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductAddonsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductAddonsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductAddonsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *ProductAddonsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductAddonsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductAddonsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductAddonsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductAddonsInner) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductAddonsInner) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductAddonsInner) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductAddonsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


