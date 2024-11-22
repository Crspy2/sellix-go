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

// checks if the ListQueries200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListQueries200ResponseData{}

// ListQueries200ResponseData struct for ListQueries200ResponseData
type ListQueries200ResponseData struct {
	Queries []QueryListing `json:"queries,omitempty"`
}

// NewListQueries200ResponseData instantiates a new ListQueries200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListQueries200ResponseData() *ListQueries200ResponseData {
	this := ListQueries200ResponseData{}
	return &this
}

// NewListQueries200ResponseDataWithDefaults instantiates a new ListQueries200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListQueries200ResponseDataWithDefaults() *ListQueries200ResponseData {
	this := ListQueries200ResponseData{}
	return &this
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *ListQueries200ResponseData) GetQueries() []QueryListing {
	if o == nil || IsNil(o.Queries) {
		var ret []QueryListing
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListQueries200ResponseData) GetQueriesOk() ([]QueryListing, bool) {
	if o == nil || IsNil(o.Queries) {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *ListQueries200ResponseData) HasQueries() bool {
	if o != nil && !IsNil(o.Queries) {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []QueryListing and assigns it to the Queries field.
func (o *ListQueries200ResponseData) SetQueries(v []QueryListing) {
	o.Queries = v
}

func (o ListQueries200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListQueries200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Queries) {
		toSerialize["queries"] = o.Queries
	}
	return toSerialize, nil
}

type NullableListQueries200ResponseData struct {
	value *ListQueries200ResponseData
	isSet bool
}

func (v NullableListQueries200ResponseData) Get() *ListQueries200ResponseData {
	return v.value
}

func (v *NullableListQueries200ResponseData) Set(val *ListQueries200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListQueries200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListQueries200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListQueries200ResponseData(val *ListQueries200ResponseData) *NullableListQueries200ResponseData {
	return &NullableListQueries200ResponseData{value: val, isSet: true}
}

func (v NullableListQueries200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListQueries200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


