# ProductTaxDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatPercentage** | Pointer to **string** |  | [optional] 
**TaxConfiguration** | Pointer to **string** |  | [optional] 
**TaxConfigurationData** | Pointer to **[]string** |  | [optional] 
**DisplayTaxOnStorefront** | Pointer to **int32** |  | [optional] 
**DisplayTaxCustomFields** | Pointer to **int32** |  | [optional] 
**ValidationOnlyForCompanies** | Pointer to **int32** |  | [optional] 
**ValidateVatNumber** | Pointer to **int32** |  | [optional] 
**PricesTaxInclusive** | Pointer to **int32** |  | [optional] 

## Methods

### NewProductTaxDetails

`func NewProductTaxDetails() *ProductTaxDetails`

NewProductTaxDetails instantiates a new ProductTaxDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTaxDetailsWithDefaults

`func NewProductTaxDetailsWithDefaults() *ProductTaxDetails`

NewProductTaxDetailsWithDefaults instantiates a new ProductTaxDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatPercentage

`func (o *ProductTaxDetails) GetVatPercentage() string`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *ProductTaxDetails) GetVatPercentageOk() (*string, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *ProductTaxDetails) SetVatPercentage(v string)`

SetVatPercentage sets VatPercentage field to given value.

### HasVatPercentage

`func (o *ProductTaxDetails) HasVatPercentage() bool`

HasVatPercentage returns a boolean if a field has been set.

### GetTaxConfiguration

`func (o *ProductTaxDetails) GetTaxConfiguration() string`

GetTaxConfiguration returns the TaxConfiguration field if non-nil, zero value otherwise.

### GetTaxConfigurationOk

`func (o *ProductTaxDetails) GetTaxConfigurationOk() (*string, bool)`

GetTaxConfigurationOk returns a tuple with the TaxConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfiguration

`func (o *ProductTaxDetails) SetTaxConfiguration(v string)`

SetTaxConfiguration sets TaxConfiguration field to given value.

### HasTaxConfiguration

`func (o *ProductTaxDetails) HasTaxConfiguration() bool`

HasTaxConfiguration returns a boolean if a field has been set.

### GetTaxConfigurationData

`func (o *ProductTaxDetails) GetTaxConfigurationData() []string`

GetTaxConfigurationData returns the TaxConfigurationData field if non-nil, zero value otherwise.

### GetTaxConfigurationDataOk

`func (o *ProductTaxDetails) GetTaxConfigurationDataOk() (*[]string, bool)`

GetTaxConfigurationDataOk returns a tuple with the TaxConfigurationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfigurationData

`func (o *ProductTaxDetails) SetTaxConfigurationData(v []string)`

SetTaxConfigurationData sets TaxConfigurationData field to given value.

### HasTaxConfigurationData

`func (o *ProductTaxDetails) HasTaxConfigurationData() bool`

HasTaxConfigurationData returns a boolean if a field has been set.

### GetDisplayTaxOnStorefront

`func (o *ProductTaxDetails) GetDisplayTaxOnStorefront() int32`

GetDisplayTaxOnStorefront returns the DisplayTaxOnStorefront field if non-nil, zero value otherwise.

### GetDisplayTaxOnStorefrontOk

`func (o *ProductTaxDetails) GetDisplayTaxOnStorefrontOk() (*int32, bool)`

GetDisplayTaxOnStorefrontOk returns a tuple with the DisplayTaxOnStorefront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxOnStorefront

`func (o *ProductTaxDetails) SetDisplayTaxOnStorefront(v int32)`

SetDisplayTaxOnStorefront sets DisplayTaxOnStorefront field to given value.

### HasDisplayTaxOnStorefront

`func (o *ProductTaxDetails) HasDisplayTaxOnStorefront() bool`

HasDisplayTaxOnStorefront returns a boolean if a field has been set.

### GetDisplayTaxCustomFields

`func (o *ProductTaxDetails) GetDisplayTaxCustomFields() int32`

GetDisplayTaxCustomFields returns the DisplayTaxCustomFields field if non-nil, zero value otherwise.

### GetDisplayTaxCustomFieldsOk

`func (o *ProductTaxDetails) GetDisplayTaxCustomFieldsOk() (*int32, bool)`

GetDisplayTaxCustomFieldsOk returns a tuple with the DisplayTaxCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxCustomFields

`func (o *ProductTaxDetails) SetDisplayTaxCustomFields(v int32)`

SetDisplayTaxCustomFields sets DisplayTaxCustomFields field to given value.

### HasDisplayTaxCustomFields

`func (o *ProductTaxDetails) HasDisplayTaxCustomFields() bool`

HasDisplayTaxCustomFields returns a boolean if a field has been set.

### GetValidationOnlyForCompanies

`func (o *ProductTaxDetails) GetValidationOnlyForCompanies() int32`

GetValidationOnlyForCompanies returns the ValidationOnlyForCompanies field if non-nil, zero value otherwise.

### GetValidationOnlyForCompaniesOk

`func (o *ProductTaxDetails) GetValidationOnlyForCompaniesOk() (*int32, bool)`

GetValidationOnlyForCompaniesOk returns a tuple with the ValidationOnlyForCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOnlyForCompanies

`func (o *ProductTaxDetails) SetValidationOnlyForCompanies(v int32)`

SetValidationOnlyForCompanies sets ValidationOnlyForCompanies field to given value.

### HasValidationOnlyForCompanies

`func (o *ProductTaxDetails) HasValidationOnlyForCompanies() bool`

HasValidationOnlyForCompanies returns a boolean if a field has been set.

### GetValidateVatNumber

`func (o *ProductTaxDetails) GetValidateVatNumber() int32`

GetValidateVatNumber returns the ValidateVatNumber field if non-nil, zero value otherwise.

### GetValidateVatNumberOk

`func (o *ProductTaxDetails) GetValidateVatNumberOk() (*int32, bool)`

GetValidateVatNumberOk returns a tuple with the ValidateVatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateVatNumber

`func (o *ProductTaxDetails) SetValidateVatNumber(v int32)`

SetValidateVatNumber sets ValidateVatNumber field to given value.

### HasValidateVatNumber

`func (o *ProductTaxDetails) HasValidateVatNumber() bool`

HasValidateVatNumber returns a boolean if a field has been set.

### GetPricesTaxInclusive

`func (o *ProductTaxDetails) GetPricesTaxInclusive() int32`

GetPricesTaxInclusive returns the PricesTaxInclusive field if non-nil, zero value otherwise.

### GetPricesTaxInclusiveOk

`func (o *ProductTaxDetails) GetPricesTaxInclusiveOk() (*int32, bool)`

GetPricesTaxInclusiveOk returns a tuple with the PricesTaxInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesTaxInclusive

`func (o *ProductTaxDetails) SetPricesTaxInclusive(v int32)`

SetPricesTaxInclusive sets PricesTaxInclusive field to given value.

### HasPricesTaxInclusive

`func (o *ProductTaxDetails) HasPricesTaxInclusive() bool`

HasPricesTaxInclusive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


