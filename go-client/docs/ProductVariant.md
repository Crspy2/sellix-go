# ProductVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **float32** | The price of the product variant. | [optional] 
**Title** | Pointer to **string** | The title of the product variant. | [optional] 
**Description** | Pointer to **string** | The description of the product variant. | [optional] 
**PriceConversions** | Pointer to [**FeedbackProductPriceConversions**](FeedbackProductPriceConversions.md) |  | [optional] 

## Methods

### NewProductVariant

`func NewProductVariant() *ProductVariant`

NewProductVariant instantiates a new ProductVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantWithDefaults

`func NewProductVariantWithDefaults() *ProductVariant`

NewProductVariantWithDefaults instantiates a new ProductVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *ProductVariant) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariant) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariant) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductVariant) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTitle

`func (o *ProductVariant) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductVariant) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductVariant) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProductVariant) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ProductVariant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductVariant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductVariant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductVariant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriceConversions

`func (o *ProductVariant) GetPriceConversions() FeedbackProductPriceConversions`

GetPriceConversions returns the PriceConversions field if non-nil, zero value otherwise.

### GetPriceConversionsOk

`func (o *ProductVariant) GetPriceConversionsOk() (*FeedbackProductPriceConversions, bool)`

GetPriceConversionsOk returns a tuple with the PriceConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceConversions

`func (o *ProductVariant) SetPriceConversions(v FeedbackProductPriceConversions)`

SetPriceConversions sets PriceConversions field to given value.

### HasPriceConversions

`func (o *ProductVariant) HasPriceConversions() bool`

HasPriceConversions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


