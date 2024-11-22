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

// checks if the FeedbackProductPriceConversions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackProductPriceConversions{}

// FeedbackProductPriceConversions struct for FeedbackProductPriceConversions
type FeedbackProductPriceConversions struct {
	// The price of the product variant in AUD.
	AUD *float32 `json:"AUD,omitempty"`
	// The price of the product variant in BGN.
	BGN *float32 `json:"BGN,omitempty"`
	// The price of the product variant in BRL.
	BRL *float32 `json:"BRL,omitempty"`
	// The price of the product variant in ZAR.
	ZAR *float32 `json:"ZAR,omitempty"`
	Crypto *FeedbackProductPriceConversionsCrypto `json:"crypto,omitempty"`
}

// NewFeedbackProductPriceConversions instantiates a new FeedbackProductPriceConversions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackProductPriceConversions() *FeedbackProductPriceConversions {
	this := FeedbackProductPriceConversions{}
	return &this
}

// NewFeedbackProductPriceConversionsWithDefaults instantiates a new FeedbackProductPriceConversions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackProductPriceConversionsWithDefaults() *FeedbackProductPriceConversions {
	this := FeedbackProductPriceConversions{}
	return &this
}

// GetAUD returns the AUD field value if set, zero value otherwise.
func (o *FeedbackProductPriceConversions) GetAUD() float32 {
	if o == nil || IsNil(o.AUD) {
		var ret float32
		return ret
	}
	return *o.AUD
}

// GetAUDOk returns a tuple with the AUD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackProductPriceConversions) GetAUDOk() (*float32, bool) {
	if o == nil || IsNil(o.AUD) {
		return nil, false
	}
	return o.AUD, true
}

// HasAUD returns a boolean if a field has been set.
func (o *FeedbackProductPriceConversions) HasAUD() bool {
	if o != nil && !IsNil(o.AUD) {
		return true
	}

	return false
}

// SetAUD gets a reference to the given float32 and assigns it to the AUD field.
func (o *FeedbackProductPriceConversions) SetAUD(v float32) {
	o.AUD = &v
}

// GetBGN returns the BGN field value if set, zero value otherwise.
func (o *FeedbackProductPriceConversions) GetBGN() float32 {
	if o == nil || IsNil(o.BGN) {
		var ret float32
		return ret
	}
	return *o.BGN
}

// GetBGNOk returns a tuple with the BGN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackProductPriceConversions) GetBGNOk() (*float32, bool) {
	if o == nil || IsNil(o.BGN) {
		return nil, false
	}
	return o.BGN, true
}

// HasBGN returns a boolean if a field has been set.
func (o *FeedbackProductPriceConversions) HasBGN() bool {
	if o != nil && !IsNil(o.BGN) {
		return true
	}

	return false
}

// SetBGN gets a reference to the given float32 and assigns it to the BGN field.
func (o *FeedbackProductPriceConversions) SetBGN(v float32) {
	o.BGN = &v
}

// GetBRL returns the BRL field value if set, zero value otherwise.
func (o *FeedbackProductPriceConversions) GetBRL() float32 {
	if o == nil || IsNil(o.BRL) {
		var ret float32
		return ret
	}
	return *o.BRL
}

// GetBRLOk returns a tuple with the BRL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackProductPriceConversions) GetBRLOk() (*float32, bool) {
	if o == nil || IsNil(o.BRL) {
		return nil, false
	}
	return o.BRL, true
}

// HasBRL returns a boolean if a field has been set.
func (o *FeedbackProductPriceConversions) HasBRL() bool {
	if o != nil && !IsNil(o.BRL) {
		return true
	}

	return false
}

// SetBRL gets a reference to the given float32 and assigns it to the BRL field.
func (o *FeedbackProductPriceConversions) SetBRL(v float32) {
	o.BRL = &v
}

// GetZAR returns the ZAR field value if set, zero value otherwise.
func (o *FeedbackProductPriceConversions) GetZAR() float32 {
	if o == nil || IsNil(o.ZAR) {
		var ret float32
		return ret
	}
	return *o.ZAR
}

// GetZAROk returns a tuple with the ZAR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackProductPriceConversions) GetZAROk() (*float32, bool) {
	if o == nil || IsNil(o.ZAR) {
		return nil, false
	}
	return o.ZAR, true
}

// HasZAR returns a boolean if a field has been set.
func (o *FeedbackProductPriceConversions) HasZAR() bool {
	if o != nil && !IsNil(o.ZAR) {
		return true
	}

	return false
}

// SetZAR gets a reference to the given float32 and assigns it to the ZAR field.
func (o *FeedbackProductPriceConversions) SetZAR(v float32) {
	o.ZAR = &v
}

// GetCrypto returns the Crypto field value if set, zero value otherwise.
func (o *FeedbackProductPriceConversions) GetCrypto() FeedbackProductPriceConversionsCrypto {
	if o == nil || IsNil(o.Crypto) {
		var ret FeedbackProductPriceConversionsCrypto
		return ret
	}
	return *o.Crypto
}

// GetCryptoOk returns a tuple with the Crypto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackProductPriceConversions) GetCryptoOk() (*FeedbackProductPriceConversionsCrypto, bool) {
	if o == nil || IsNil(o.Crypto) {
		return nil, false
	}
	return o.Crypto, true
}

// HasCrypto returns a boolean if a field has been set.
func (o *FeedbackProductPriceConversions) HasCrypto() bool {
	if o != nil && !IsNil(o.Crypto) {
		return true
	}

	return false
}

// SetCrypto gets a reference to the given FeedbackProductPriceConversionsCrypto and assigns it to the Crypto field.
func (o *FeedbackProductPriceConversions) SetCrypto(v FeedbackProductPriceConversionsCrypto) {
	o.Crypto = &v
}

func (o FeedbackProductPriceConversions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackProductPriceConversions) ToMap() (map[string]interface{}, error) {
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

type NullableFeedbackProductPriceConversions struct {
	value *FeedbackProductPriceConversions
	isSet bool
}

func (v NullableFeedbackProductPriceConversions) Get() *FeedbackProductPriceConversions {
	return v.value
}

func (v *NullableFeedbackProductPriceConversions) Set(val *FeedbackProductPriceConversions) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackProductPriceConversions) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackProductPriceConversions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackProductPriceConversions(val *FeedbackProductPriceConversions) *NullableFeedbackProductPriceConversions {
	return &NullableFeedbackProductPriceConversions{value: val, isSet: true}
}

func (v NullableFeedbackProductPriceConversions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackProductPriceConversions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


