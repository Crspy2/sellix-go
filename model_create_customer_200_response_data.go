/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sellix

import (
	"encoding/json"
)

// checks if the CreateCustomer200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomer200ResponseData{}

// CreateCustomer200ResponseData struct for CreateCustomer200ResponseData
type CreateCustomer200ResponseData struct {
	// Customer ID
	CustomerId *string `json:"customer_id,omitempty"`
}

// NewCreateCustomer200ResponseData instantiates a new CreateCustomer200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomer200ResponseData() *CreateCustomer200ResponseData {
	this := CreateCustomer200ResponseData{}
	return &this
}

// NewCreateCustomer200ResponseDataWithDefaults instantiates a new CreateCustomer200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomer200ResponseDataWithDefaults() *CreateCustomer200ResponseData {
	this := CreateCustomer200ResponseData{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CreateCustomer200ResponseData) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomer200ResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateCustomer200ResponseData) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CreateCustomer200ResponseData) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o CreateCustomer200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomer200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullableCreateCustomer200ResponseData struct {
	value *CreateCustomer200ResponseData
	isSet bool
}

func (v NullableCreateCustomer200ResponseData) Get() *CreateCustomer200ResponseData {
	return v.value
}

func (v *NullableCreateCustomer200ResponseData) Set(val *CreateCustomer200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomer200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomer200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomer200ResponseData(val *CreateCustomer200ResponseData) *NullableCreateCustomer200ResponseData {
	return &NullableCreateCustomer200ResponseData{value: val, isSet: true}
}

func (v NullableCreateCustomer200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomer200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


