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

// checks if the GatewayFeesConversions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayFeesConversions{}

// GatewayFeesConversions struct for GatewayFeesConversions
type GatewayFeesConversions struct {
	// The price of the product variant in AUD.
	AUD *float32 `json:"AUD,omitempty"`
	// The price of the product variant in BGN.
	BGN *float32 `json:"BGN,omitempty"`
	// The price of the product variant in BRL.
	BRL *float32 `json:"BRL,omitempty"`
	// The price of the product variant in ZAR.
	ZAR *float32 `json:"ZAR,omitempty"`
	Crypto *GatewayFeesConversionsCrypto `json:"crypto,omitempty"`
}

// NewGatewayFeesConversions instantiates a new GatewayFeesConversions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayFeesConversions() *GatewayFeesConversions {
	this := GatewayFeesConversions{}
	return &this
}

// NewGatewayFeesConversionsWithDefaults instantiates a new GatewayFeesConversions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayFeesConversionsWithDefaults() *GatewayFeesConversions {
	this := GatewayFeesConversions{}
	return &this
}

// GetAUD returns the AUD field value if set, zero value otherwise.
func (o *GatewayFeesConversions) GetAUD() float32 {
	if o == nil || IsNil(o.AUD) {
		var ret float32
		return ret
	}
	return *o.AUD
}

// GetAUDOk returns a tuple with the AUD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFeesConversions) GetAUDOk() (*float32, bool) {
	if o == nil || IsNil(o.AUD) {
		return nil, false
	}
	return o.AUD, true
}

// HasAUD returns a boolean if a field has been set.
func (o *GatewayFeesConversions) HasAUD() bool {
	if o != nil && !IsNil(o.AUD) {
		return true
	}

	return false
}

// SetAUD gets a reference to the given float32 and assigns it to the AUD field.
func (o *GatewayFeesConversions) SetAUD(v float32) {
	o.AUD = &v
}

// GetBGN returns the BGN field value if set, zero value otherwise.
func (o *GatewayFeesConversions) GetBGN() float32 {
	if o == nil || IsNil(o.BGN) {
		var ret float32
		return ret
	}
	return *o.BGN
}

// GetBGNOk returns a tuple with the BGN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFeesConversions) GetBGNOk() (*float32, bool) {
	if o == nil || IsNil(o.BGN) {
		return nil, false
	}
	return o.BGN, true
}

// HasBGN returns a boolean if a field has been set.
func (o *GatewayFeesConversions) HasBGN() bool {
	if o != nil && !IsNil(o.BGN) {
		return true
	}

	return false
}

// SetBGN gets a reference to the given float32 and assigns it to the BGN field.
func (o *GatewayFeesConversions) SetBGN(v float32) {
	o.BGN = &v
}

// GetBRL returns the BRL field value if set, zero value otherwise.
func (o *GatewayFeesConversions) GetBRL() float32 {
	if o == nil || IsNil(o.BRL) {
		var ret float32
		return ret
	}
	return *o.BRL
}

// GetBRLOk returns a tuple with the BRL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFeesConversions) GetBRLOk() (*float32, bool) {
	if o == nil || IsNil(o.BRL) {
		return nil, false
	}
	return o.BRL, true
}

// HasBRL returns a boolean if a field has been set.
func (o *GatewayFeesConversions) HasBRL() bool {
	if o != nil && !IsNil(o.BRL) {
		return true
	}

	return false
}

// SetBRL gets a reference to the given float32 and assigns it to the BRL field.
func (o *GatewayFeesConversions) SetBRL(v float32) {
	o.BRL = &v
}

// GetZAR returns the ZAR field value if set, zero value otherwise.
func (o *GatewayFeesConversions) GetZAR() float32 {
	if o == nil || IsNil(o.ZAR) {
		var ret float32
		return ret
	}
	return *o.ZAR
}

// GetZAROk returns a tuple with the ZAR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFeesConversions) GetZAROk() (*float32, bool) {
	if o == nil || IsNil(o.ZAR) {
		return nil, false
	}
	return o.ZAR, true
}

// HasZAR returns a boolean if a field has been set.
func (o *GatewayFeesConversions) HasZAR() bool {
	if o != nil && !IsNil(o.ZAR) {
		return true
	}

	return false
}

// SetZAR gets a reference to the given float32 and assigns it to the ZAR field.
func (o *GatewayFeesConversions) SetZAR(v float32) {
	o.ZAR = &v
}

// GetCrypto returns the Crypto field value if set, zero value otherwise.
func (o *GatewayFeesConversions) GetCrypto() GatewayFeesConversionsCrypto {
	if o == nil || IsNil(o.Crypto) {
		var ret GatewayFeesConversionsCrypto
		return ret
	}
	return *o.Crypto
}

// GetCryptoOk returns a tuple with the Crypto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayFeesConversions) GetCryptoOk() (*GatewayFeesConversionsCrypto, bool) {
	if o == nil || IsNil(o.Crypto) {
		return nil, false
	}
	return o.Crypto, true
}

// HasCrypto returns a boolean if a field has been set.
func (o *GatewayFeesConversions) HasCrypto() bool {
	if o != nil && !IsNil(o.Crypto) {
		return true
	}

	return false
}

// SetCrypto gets a reference to the given GatewayFeesConversionsCrypto and assigns it to the Crypto field.
func (o *GatewayFeesConversions) SetCrypto(v GatewayFeesConversionsCrypto) {
	o.Crypto = &v
}

func (o GatewayFeesConversions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayFeesConversions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AUD) {
		toSerialize["AUD"] = o.AUD
	}
	if !IsNil(o.BGN) {
		toSerialize["BGN"] = o.BGN
	}
	if !IsNil(o.BRL) {
		toSerialize["BRL"] = o.BRL
	}
	if !IsNil(o.ZAR) {
		toSerialize["ZAR"] = o.ZAR
	}
	if !IsNil(o.Crypto) {
		toSerialize["crypto"] = o.Crypto
	}
	return toSerialize, nil
}

type NullableGatewayFeesConversions struct {
	value *GatewayFeesConversions
	isSet bool
}

func (v NullableGatewayFeesConversions) Get() *GatewayFeesConversions {
	return v.value
}

func (v *NullableGatewayFeesConversions) Set(val *GatewayFeesConversions) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayFeesConversions) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayFeesConversions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayFeesConversions(val *GatewayFeesConversions) *NullableGatewayFeesConversions {
	return &NullableGatewayFeesConversions{value: val, isSet: true}
}

func (v NullableGatewayFeesConversions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayFeesConversions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

