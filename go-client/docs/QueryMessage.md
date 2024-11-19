# QueryMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Who this message belongs to; the customer or the merchant. | [optional] 
**Message** | Pointer to **string** | Message left by the customer or merchant. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for when the message was sent. | [optional] 

## Methods

### NewQueryMessage

`func NewQueryMessage() *QueryMessage`

NewQueryMessage instantiates a new QueryMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMessageWithDefaults

`func NewQueryMessageWithDefaults() *QueryMessage`

NewQueryMessageWithDefaults instantiates a new QueryMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *QueryMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *QueryMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *QueryMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *QueryMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetMessage

`func (o *QueryMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QueryMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QueryMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QueryMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueryMessage) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueryMessage) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueryMessage) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueryMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


