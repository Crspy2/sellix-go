/*
Sellix Developers API

Sellix public API for developers to access merchant resources

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PaypalDispute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaypalDispute{}

// PaypalDispute struct for PaypalDispute
type PaypalDispute struct {
	// Unique ID of the resource, used as reference across the API.
	Id *string `json:"id,omitempty"`
	// Unique ID of the invoice linked to this dispute.
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The shop ID to which this paypal dispute belongs.
	ShopId *int32 `json:"shop_id,omitempty"`
	// The dispute reason is why the customer has opened a dispute against your order.
	Reason *string `json:"reason,omitempty"`
	// Each dispute is automatically updated when we receive an update from PayPal, the status indicates how it is going.
	Status *string `json:"status,omitempty"`
	// When a dispute it’s solved, its outcome is updated.
	Outcome *string `json:"outcome,omitempty"`
	Messages []PaypalDisputeMessage `json:"messages,omitempty"`
	// The stage in the dispute lifecycle.
	LifeCycleStage *string `json:"life_cycle_stage,omitempty"`
	// Within which date the seller needs to respond.
	SellerResponseDueDate *int32 `json:"seller_response_due_date,omitempty"`
	// Timestamp for the creation of the dispute.
	CreatedAt *int32 `json:"created_at,omitempty"`
	// When the dispute was updated.
	UpdatedAt *int32 `json:"updated_at,omitempty"`
}

// NewPaypalDispute instantiates a new PaypalDispute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalDispute() *PaypalDispute {
	this := PaypalDispute{}
	return &this
}

// NewPaypalDisputeWithDefaults instantiates a new PaypalDispute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalDisputeWithDefaults() *PaypalDispute {
	this := PaypalDispute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaypalDispute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaypalDispute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaypalDispute) SetId(v string) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *PaypalDispute) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PaypalDispute) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *PaypalDispute) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *PaypalDispute) GetShopId() int32 {
	if o == nil || IsNil(o.ShopId) {
		var ret int32
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetShopIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *PaypalDispute) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given int32 and assigns it to the ShopId field.
func (o *PaypalDispute) SetShopId(v int32) {
	o.ShopId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PaypalDispute) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PaypalDispute) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PaypalDispute) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaypalDispute) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaypalDispute) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaypalDispute) SetStatus(v string) {
	o.Status = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *PaypalDispute) GetOutcome() string {
	if o == nil || IsNil(o.Outcome) {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetOutcomeOk() (*string, bool) {
	if o == nil || IsNil(o.Outcome) {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *PaypalDispute) HasOutcome() bool {
	if o != nil && !IsNil(o.Outcome) {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *PaypalDispute) SetOutcome(v string) {
	o.Outcome = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *PaypalDispute) GetMessages() []PaypalDisputeMessage {
	if o == nil || IsNil(o.Messages) {
		var ret []PaypalDisputeMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetMessagesOk() ([]PaypalDisputeMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *PaypalDispute) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []PaypalDisputeMessage and assigns it to the Messages field.
func (o *PaypalDispute) SetMessages(v []PaypalDisputeMessage) {
	o.Messages = v
}

// GetLifeCycleStage returns the LifeCycleStage field value if set, zero value otherwise.
func (o *PaypalDispute) GetLifeCycleStage() string {
	if o == nil || IsNil(o.LifeCycleStage) {
		var ret string
		return ret
	}
	return *o.LifeCycleStage
}

// GetLifeCycleStageOk returns a tuple with the LifeCycleStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetLifeCycleStageOk() (*string, bool) {
	if o == nil || IsNil(o.LifeCycleStage) {
		return nil, false
	}
	return o.LifeCycleStage, true
}

// HasLifeCycleStage returns a boolean if a field has been set.
func (o *PaypalDispute) HasLifeCycleStage() bool {
	if o != nil && !IsNil(o.LifeCycleStage) {
		return true
	}

	return false
}

// SetLifeCycleStage gets a reference to the given string and assigns it to the LifeCycleStage field.
func (o *PaypalDispute) SetLifeCycleStage(v string) {
	o.LifeCycleStage = &v
}

// GetSellerResponseDueDate returns the SellerResponseDueDate field value if set, zero value otherwise.
func (o *PaypalDispute) GetSellerResponseDueDate() int32 {
	if o == nil || IsNil(o.SellerResponseDueDate) {
		var ret int32
		return ret
	}
	return *o.SellerResponseDueDate
}

// GetSellerResponseDueDateOk returns a tuple with the SellerResponseDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetSellerResponseDueDateOk() (*int32, bool) {
	if o == nil || IsNil(o.SellerResponseDueDate) {
		return nil, false
	}
	return o.SellerResponseDueDate, true
}

// HasSellerResponseDueDate returns a boolean if a field has been set.
func (o *PaypalDispute) HasSellerResponseDueDate() bool {
	if o != nil && !IsNil(o.SellerResponseDueDate) {
		return true
	}

	return false
}

// SetSellerResponseDueDate gets a reference to the given int32 and assigns it to the SellerResponseDueDate field.
func (o *PaypalDispute) SetSellerResponseDueDate(v int32) {
	o.SellerResponseDueDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaypalDispute) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaypalDispute) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *PaypalDispute) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaypalDispute) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaypalDispute) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaypalDispute) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *PaypalDispute) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o PaypalDispute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaypalDispute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Outcome) {
		toSerialize["outcome"] = o.Outcome
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.LifeCycleStage) {
		toSerialize["life_cycle_stage"] = o.LifeCycleStage
	}
	if !IsNil(o.SellerResponseDueDate) {
		toSerialize["seller_response_due_date"] = o.SellerResponseDueDate
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePaypalDispute struct {
	value *PaypalDispute
	isSet bool
}

func (v NullablePaypalDispute) Get() *PaypalDispute {
	return v.value
}

func (v *NullablePaypalDispute) Set(val *PaypalDispute) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalDispute) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalDispute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalDispute(val *PaypalDispute) *NullablePaypalDispute {
	return &NullablePaypalDispute{value: val, isSet: true}
}

func (v NullablePaypalDispute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalDispute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

