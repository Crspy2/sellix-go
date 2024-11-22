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

// checks if the InvoiceRatesSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceRatesSnapshot{}

// InvoiceRatesSnapshot struct for InvoiceRatesSnapshot
type InvoiceRatesSnapshot struct {
	// ID of the resource.
	Id *int32 `json:"id,omitempty"`
	// The rate in USD.
	USD *string `json:"USD,omitempty"`
	// The rate in CAD.
	CAD *string `json:"CAD,omitempty"`
	// The rate in HKD.
	HKD *string `json:"HKD,omitempty"`
}

// NewInvoiceRatesSnapshot instantiates a new InvoiceRatesSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceRatesSnapshot() *InvoiceRatesSnapshot {
	this := InvoiceRatesSnapshot{}
	return &this
}

// NewInvoiceRatesSnapshotWithDefaults instantiates a new InvoiceRatesSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRatesSnapshotWithDefaults() *InvoiceRatesSnapshot {
	this := InvoiceRatesSnapshot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceRatesSnapshot) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRatesSnapshot) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceRatesSnapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InvoiceRatesSnapshot) SetId(v int32) {
	o.Id = &v
}

// GetUSD returns the USD field value if set, zero value otherwise.
func (o *InvoiceRatesSnapshot) GetUSD() string {
	if o == nil || IsNil(o.USD) {
		var ret string
		return ret
	}
	return *o.USD
}

// GetUSDOk returns a tuple with the USD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRatesSnapshot) GetUSDOk() (*string, bool) {
	if o == nil || IsNil(o.USD) {
		return nil, false
	}
	return o.USD, true
}

// HasUSD returns a boolean if a field has been set.
func (o *InvoiceRatesSnapshot) HasUSD() bool {
	if o != nil && !IsNil(o.USD) {
		return true
	}

	return false
}

// SetUSD gets a reference to the given string and assigns it to the USD field.
func (o *InvoiceRatesSnapshot) SetUSD(v string) {
	o.USD = &v
}

// GetCAD returns the CAD field value if set, zero value otherwise.
func (o *InvoiceRatesSnapshot) GetCAD() string {
	if o == nil || IsNil(o.CAD) {
		var ret string
		return ret
	}
	return *o.CAD
}

// GetCADOk returns a tuple with the CAD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRatesSnapshot) GetCADOk() (*string, bool) {
	if o == nil || IsNil(o.CAD) {
		return nil, false
	}
	return o.CAD, true
}

// HasCAD returns a boolean if a field has been set.
func (o *InvoiceRatesSnapshot) HasCAD() bool {
	if o != nil && !IsNil(o.CAD) {
		return true
	}

	return false
}

// SetCAD gets a reference to the given string and assigns it to the CAD field.
func (o *InvoiceRatesSnapshot) SetCAD(v string) {
	o.CAD = &v
}

// GetHKD returns the HKD field value if set, zero value otherwise.
func (o *InvoiceRatesSnapshot) GetHKD() string {
	if o == nil || IsNil(o.HKD) {
		var ret string
		return ret
	}
	return *o.HKD
}

// GetHKDOk returns a tuple with the HKD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRatesSnapshot) GetHKDOk() (*string, bool) {
	if o == nil || IsNil(o.HKD) {
		return nil, false
	}
	return o.HKD, true
}

// HasHKD returns a boolean if a field has been set.
func (o *InvoiceRatesSnapshot) HasHKD() bool {
	if o != nil && !IsNil(o.HKD) {
		return true
	}

	return false
}

// SetHKD gets a reference to the given string and assigns it to the HKD field.
func (o *InvoiceRatesSnapshot) SetHKD(v string) {
	o.HKD = &v
}

func (o InvoiceRatesSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceRatesSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.USD) {
		toSerialize["USD"] = o.USD
	}
	if !IsNil(o.CAD) {
		toSerialize["CAD"] = o.CAD
	}
	if !IsNil(o.HKD) {
		toSerialize["HKD"] = o.HKD
	}
	return toSerialize, nil
}

type NullableInvoiceRatesSnapshot struct {
	value *InvoiceRatesSnapshot
	isSet bool
}

func (v NullableInvoiceRatesSnapshot) Get() *InvoiceRatesSnapshot {
	return v.value
}

func (v *NullableInvoiceRatesSnapshot) Set(val *InvoiceRatesSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceRatesSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceRatesSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceRatesSnapshot(val *InvoiceRatesSnapshot) *NullableInvoiceRatesSnapshot {
	return &NullableInvoiceRatesSnapshot{value: val, isSet: true}
}

func (v NullableInvoiceRatesSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceRatesSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


