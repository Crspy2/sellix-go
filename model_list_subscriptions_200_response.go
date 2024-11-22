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

// checks if the ListSubscriptions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSubscriptions200Response{}

// ListSubscriptions200Response struct for ListSubscriptions200Response
type ListSubscriptions200Response struct {
	Status *int32 `json:"status,omitempty"`
	Data *ListSubscriptions200ResponseData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Log map[string]interface{} `json:"log,omitempty"`
	Error *string `json:"error,omitempty"`
	Env *string `json:"env,omitempty"`
}

// NewListSubscriptions200Response instantiates a new ListSubscriptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSubscriptions200Response() *ListSubscriptions200Response {
	this := ListSubscriptions200Response{}
	return &this
}

// NewListSubscriptions200ResponseWithDefaults instantiates a new ListSubscriptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSubscriptions200ResponseWithDefaults() *ListSubscriptions200Response {
	this := ListSubscriptions200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ListSubscriptions200Response) SetStatus(v int32) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetData() ListSubscriptions200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret ListSubscriptions200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetDataOk() (*ListSubscriptions200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSubscriptions200ResponseData and assigns it to the Data field.
func (o *ListSubscriptions200Response) SetData(v ListSubscriptions200ResponseData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ListSubscriptions200Response) SetMessage(v string) {
	o.Message = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetLog() map[string]interface{} {
	if o == nil || IsNil(o.Log) {
		var ret map[string]interface{}
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetLogOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Log) {
		return map[string]interface{}{}, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given map[string]interface{} and assigns it to the Log field.
func (o *ListSubscriptions200Response) SetLog(v map[string]interface{}) {
	o.Log = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ListSubscriptions200Response) SetError(v string) {
	o.Error = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ListSubscriptions200Response) GetEnv() string {
	if o == nil || IsNil(o.Env) {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSubscriptions200Response) GetEnvOk() (*string, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ListSubscriptions200Response) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ListSubscriptions200Response) SetEnv(v string) {
	o.Env = &v
}

func (o ListSubscriptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSubscriptions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	return toSerialize, nil
}

type NullableListSubscriptions200Response struct {
	value *ListSubscriptions200Response
	isSet bool
}

func (v NullableListSubscriptions200Response) Get() *ListSubscriptions200Response {
	return v.value
}

func (v *NullableListSubscriptions200Response) Set(val *ListSubscriptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListSubscriptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSubscriptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSubscriptions200Response(val *ListSubscriptions200Response) *NullableListSubscriptions200Response {
	return &NullableListSubscriptions200Response{value: val, isSet: true}
}

func (v NullableListSubscriptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSubscriptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


