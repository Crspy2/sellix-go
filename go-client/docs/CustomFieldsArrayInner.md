# CustomFieldsArrayInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the custom field. | [optional] 
**Name** | Pointer to **string** | Custom field name displayed to the customer. | [optional] 
**Regex** | Pointer to **string** | Regex that the custom field value must match. | [optional] 
**Placeholder** | Pointer to **string** | Placeholder for the custom field input. | [optional] 
**Default** | Pointer to **string** | Default value if the customer leaves the input empty. | [optional] 
**Required** | Pointer to **bool** | Whether or not this custom field is required to proceed with the purchase. | [optional] 

## Methods

### NewCustomFieldsArrayInner

`func NewCustomFieldsArrayInner() *CustomFieldsArrayInner`

NewCustomFieldsArrayInner instantiates a new CustomFieldsArrayInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsArrayInnerWithDefaults

`func NewCustomFieldsArrayInnerWithDefaults() *CustomFieldsArrayInner`

NewCustomFieldsArrayInnerWithDefaults instantiates a new CustomFieldsArrayInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomFieldsArrayInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldsArrayInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldsArrayInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomFieldsArrayInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CustomFieldsArrayInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldsArrayInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldsArrayInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFieldsArrayInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegex

`func (o *CustomFieldsArrayInner) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *CustomFieldsArrayInner) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *CustomFieldsArrayInner) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *CustomFieldsArrayInner) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetPlaceholder

`func (o *CustomFieldsArrayInner) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CustomFieldsArrayInner) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CustomFieldsArrayInner) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CustomFieldsArrayInner) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetDefault

`func (o *CustomFieldsArrayInner) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CustomFieldsArrayInner) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CustomFieldsArrayInner) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CustomFieldsArrayInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRequired

`func (o *CustomFieldsArrayInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomFieldsArrayInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomFieldsArrayInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomFieldsArrayInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


