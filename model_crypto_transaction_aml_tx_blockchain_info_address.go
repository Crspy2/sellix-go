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

// checks if the CryptoTransactionAmlTxBlockchainInfoAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoTransactionAmlTxBlockchainInfoAddress{}

// CryptoTransactionAmlTxBlockchainInfoAddress struct for CryptoTransactionAmlTxBlockchainInfoAddress
type CryptoTransactionAmlTxBlockchainInfoAddress struct {
	// Has sent.
	HasSent *bool `json:"has_sent,omitempty"`
	// Has received.
	HasReceived *bool `json:"has_received,omitempty"`
}

// NewCryptoTransactionAmlTxBlockchainInfoAddress instantiates a new CryptoTransactionAmlTxBlockchainInfoAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoTransactionAmlTxBlockchainInfoAddress() *CryptoTransactionAmlTxBlockchainInfoAddress {
	this := CryptoTransactionAmlTxBlockchainInfoAddress{}
	return &this
}

// NewCryptoTransactionAmlTxBlockchainInfoAddressWithDefaults instantiates a new CryptoTransactionAmlTxBlockchainInfoAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoTransactionAmlTxBlockchainInfoAddressWithDefaults() *CryptoTransactionAmlTxBlockchainInfoAddress {
	this := CryptoTransactionAmlTxBlockchainInfoAddress{}
	return &this
}

// GetHasSent returns the HasSent field value if set, zero value otherwise.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) GetHasSent() bool {
	if o == nil || IsNil(o.HasSent) {
		var ret bool
		return ret
	}
	return *o.HasSent
}

// GetHasSentOk returns a tuple with the HasSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) GetHasSentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSent) {
		return nil, false
	}
	return o.HasSent, true
}

// HasHasSent returns a boolean if a field has been set.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) HasHasSent() bool {
	if o != nil && !IsNil(o.HasSent) {
		return true
	}

	return false
}

// SetHasSent gets a reference to the given bool and assigns it to the HasSent field.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) SetHasSent(v bool) {
	o.HasSent = &v
}

// GetHasReceived returns the HasReceived field value if set, zero value otherwise.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) GetHasReceived() bool {
	if o == nil || IsNil(o.HasReceived) {
		var ret bool
		return ret
	}
	return *o.HasReceived
}

// GetHasReceivedOk returns a tuple with the HasReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) GetHasReceivedOk() (*bool, bool) {
	if o == nil || IsNil(o.HasReceived) {
		return nil, false
	}
	return o.HasReceived, true
}

// HasHasReceived returns a boolean if a field has been set.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) HasHasReceived() bool {
	if o != nil && !IsNil(o.HasReceived) {
		return true
	}

	return false
}

// SetHasReceived gets a reference to the given bool and assigns it to the HasReceived field.
func (o *CryptoTransactionAmlTxBlockchainInfoAddress) SetHasReceived(v bool) {
	o.HasReceived = &v
}

func (o CryptoTransactionAmlTxBlockchainInfoAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoTransactionAmlTxBlockchainInfoAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HasSent) {
		toSerialize["has_sent"] = o.HasSent
	}
	if !IsNil(o.HasReceived) {
		toSerialize["has_received"] = o.HasReceived
	}
	return toSerialize, nil
}

type NullableCryptoTransactionAmlTxBlockchainInfoAddress struct {
	value *CryptoTransactionAmlTxBlockchainInfoAddress
	isSet bool
}

func (v NullableCryptoTransactionAmlTxBlockchainInfoAddress) Get() *CryptoTransactionAmlTxBlockchainInfoAddress {
	return v.value
}

func (v *NullableCryptoTransactionAmlTxBlockchainInfoAddress) Set(val *CryptoTransactionAmlTxBlockchainInfoAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoTransactionAmlTxBlockchainInfoAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoTransactionAmlTxBlockchainInfoAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoTransactionAmlTxBlockchainInfoAddress(val *CryptoTransactionAmlTxBlockchainInfoAddress) *NullableCryptoTransactionAmlTxBlockchainInfoAddress {
	return &NullableCryptoTransactionAmlTxBlockchainInfoAddress{value: val, isSet: true}
}

func (v NullableCryptoTransactionAmlTxBlockchainInfoAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoTransactionAmlTxBlockchainInfoAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


