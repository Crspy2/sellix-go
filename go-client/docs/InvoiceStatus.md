# InvoiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID of the invoice status. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice this status is linked to. | [optional] 
**Status** | Pointer to **string** | Status of the invoice. | [optional] 
**Details** | Pointer to **string** | Additional details on the status change. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the status change date. | [optional] 

## Methods

### NewInvoiceStatus

`func NewInvoiceStatus() *InvoiceStatus`

NewInvoiceStatus instantiates a new InvoiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceStatusWithDefaults

`func NewInvoiceStatusWithDefaults() *InvoiceStatus`

NewInvoiceStatusWithDefaults instantiates a new InvoiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *InvoiceStatus) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceStatus) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceStatus) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceStatus) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetStatus

`func (o *InvoiceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetails

`func (o *InvoiceStatus) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InvoiceStatus) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InvoiceStatus) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InvoiceStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvoiceStatus) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceStatus) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceStatus) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvoiceStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


