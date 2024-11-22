/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sellix

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubscriptionInvoiceResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionInvoiceResponseData{}

// SubscriptionInvoiceResponseData struct for SubscriptionInvoiceResponseData
type SubscriptionInvoiceResponseData struct {
	// Sellix hosted payment page.
	Url string `json:"url"`
	// Unique ID of the invoice created for the subscription.
	Uniqid string `json:"uniqid"`
}

type _SubscriptionInvoiceResponseData SubscriptionInvoiceResponseData

// NewSubscriptionInvoiceResponseData instantiates a new SubscriptionInvoiceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionInvoiceResponseData(url string, uniqid string) *SubscriptionInvoiceResponseData {
	this := SubscriptionInvoiceResponseData{}
	this.Url = url
	this.Uniqid = uniqid
	return &this
}

// NewSubscriptionInvoiceResponseDataWithDefaults instantiates a new SubscriptionInvoiceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionInvoiceResponseDataWithDefaults() *SubscriptionInvoiceResponseData {
	this := SubscriptionInvoiceResponseData{}
	return &this
}

// GetUrl returns the Url field value
func (o *SubscriptionInvoiceResponseData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SubscriptionInvoiceResponseData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SubscriptionInvoiceResponseData) SetUrl(v string) {
	o.Url = v
}

// GetUniqid returns the Uniqid field value
func (o *SubscriptionInvoiceResponseData) GetUniqid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uniqid
}

// GetUniqidOk returns a tuple with the Uniqid field value
// and a boolean to check if the value has been set.
func (o *SubscriptionInvoiceResponseData) GetUniqidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uniqid, true
}

// SetUniqid sets field value
func (o *SubscriptionInvoiceResponseData) SetUniqid(v string) {
	o.Uniqid = v
}

func (o SubscriptionInvoiceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionInvoiceResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["uniqid"] = o.Uniqid
	return toSerialize, nil
}

func (o *SubscriptionInvoiceResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"uniqid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscriptionInvoiceResponseData := _SubscriptionInvoiceResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionInvoiceResponseData)

	if err != nil {
		return err
	}

	*o = SubscriptionInvoiceResponseData(varSubscriptionInvoiceResponseData)

	return err
}

type NullableSubscriptionInvoiceResponseData struct {
	value *SubscriptionInvoiceResponseData
	isSet bool
}

func (v NullableSubscriptionInvoiceResponseData) Get() *SubscriptionInvoiceResponseData {
	return v.value
}

func (v *NullableSubscriptionInvoiceResponseData) Set(val *SubscriptionInvoiceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionInvoiceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionInvoiceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionInvoiceResponseData(val *SubscriptionInvoiceResponseData) *NullableSubscriptionInvoiceResponseData {
	return &NullableSubscriptionInvoiceResponseData{value: val, isSet: true}
}

func (v NullableSubscriptionInvoiceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionInvoiceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


