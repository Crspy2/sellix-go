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

// checks if the CryptoTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptoTransaction{}

// CryptoTransaction struct for CryptoTransaction
type CryptoTransaction struct {
	// Crypto amount sent in the transaction.
	CryptoAmount *float64 `json:"crypto_amount,omitempty"`
	// Crypto transaction hash.
	Hash *string `json:"hash,omitempty"`
	// Crypto transaction confirmations, not updated once the invoice status is COMPLETED or VOIDED.
	Confirmations *int32 `json:"confirmations,omitempty"`
	// Timestamp for the crypto transaction reception date
	CreatedAt *int32 `json:"created_at,omitempty"`
	// Crypto transaction update date.
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	AmlTx *CryptoTransactionAmlTx `json:"aml_tx,omitempty"`
}

// NewCryptoTransaction instantiates a new CryptoTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptoTransaction() *CryptoTransaction {
	this := CryptoTransaction{}
	return &this
}

// NewCryptoTransactionWithDefaults instantiates a new CryptoTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptoTransactionWithDefaults() *CryptoTransaction {
	this := CryptoTransaction{}
	return &this
}

// GetCryptoAmount returns the CryptoAmount field value if set, zero value otherwise.
func (o *CryptoTransaction) GetCryptoAmount() float64 {
	if o == nil || IsNil(o.CryptoAmount) {
		var ret float64
		return ret
	}
	return *o.CryptoAmount
}

// GetCryptoAmountOk returns a tuple with the CryptoAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetCryptoAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.CryptoAmount) {
		return nil, false
	}
	return o.CryptoAmount, true
}

// HasCryptoAmount returns a boolean if a field has been set.
func (o *CryptoTransaction) HasCryptoAmount() bool {
	if o != nil && !IsNil(o.CryptoAmount) {
		return true
	}

	return false
}

// SetCryptoAmount gets a reference to the given float64 and assigns it to the CryptoAmount field.
func (o *CryptoTransaction) SetCryptoAmount(v float64) {
	o.CryptoAmount = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CryptoTransaction) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CryptoTransaction) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *CryptoTransaction) SetHash(v string) {
	o.Hash = &v
}

// GetConfirmations returns the Confirmations field value if set, zero value otherwise.
func (o *CryptoTransaction) GetConfirmations() int32 {
	if o == nil || IsNil(o.Confirmations) {
		var ret int32
		return ret
	}
	return *o.Confirmations
}

// GetConfirmationsOk returns a tuple with the Confirmations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetConfirmationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Confirmations) {
		return nil, false
	}
	return o.Confirmations, true
}

// HasConfirmations returns a boolean if a field has been set.
func (o *CryptoTransaction) HasConfirmations() bool {
	if o != nil && !IsNil(o.Confirmations) {
		return true
	}

	return false
}

// SetConfirmations gets a reference to the given int32 and assigns it to the Confirmations field.
func (o *CryptoTransaction) SetConfirmations(v int32) {
	o.Confirmations = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CryptoTransaction) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CryptoTransaction) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *CryptoTransaction) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CryptoTransaction) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CryptoTransaction) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *CryptoTransaction) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetAmlTx returns the AmlTx field value if set, zero value otherwise.
func (o *CryptoTransaction) GetAmlTx() CryptoTransactionAmlTx {
	if o == nil || IsNil(o.AmlTx) {
		var ret CryptoTransactionAmlTx
		return ret
	}
	return *o.AmlTx
}

// GetAmlTxOk returns a tuple with the AmlTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CryptoTransaction) GetAmlTxOk() (*CryptoTransactionAmlTx, bool) {
	if o == nil || IsNil(o.AmlTx) {
		return nil, false
	}
	return o.AmlTx, true
}

// HasAmlTx returns a boolean if a field has been set.
func (o *CryptoTransaction) HasAmlTx() bool {
	if o != nil && !IsNil(o.AmlTx) {
		return true
	}

	return false
}

// SetAmlTx gets a reference to the given CryptoTransactionAmlTx and assigns it to the AmlTx field.
func (o *CryptoTransaction) SetAmlTx(v CryptoTransactionAmlTx) {
	o.AmlTx = &v
}

func (o CryptoTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptoTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CryptoAmount) {
		toSerialize["crypto_amount"] = o.CryptoAmount
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Confirmations) {
		toSerialize["confirmations"] = o.Confirmations
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.AmlTx) {
		toSerialize["aml_tx"] = o.AmlTx
	}
	return toSerialize, nil
}

type NullableCryptoTransaction struct {
	value *CryptoTransaction
	isSet bool
}

func (v NullableCryptoTransaction) Get() *CryptoTransaction {
	return v.value
}

func (v *NullableCryptoTransaction) Set(val *CryptoTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptoTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptoTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptoTransaction(val *CryptoTransaction) *NullableCryptoTransaction {
	return &NullableCryptoTransaction{value: val, isSet: true}
}

func (v NullableCryptoTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptoTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


