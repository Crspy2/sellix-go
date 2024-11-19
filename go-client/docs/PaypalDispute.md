# PaypalDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice linked to this dispute. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this paypal dispute belongs. | [optional] 
**Reason** | Pointer to **string** | The dispute reason is why the customer has opened a dispute against your order. | [optional] 
**Status** | Pointer to **string** | Each dispute is automatically updated when we receive an update from PayPal, the status indicates how it is going. | [optional] 
**Outcome** | Pointer to **string** | When a dispute it’s solved, its outcome is updated. | [optional] 
**Messages** | Pointer to [**[]PaypalDisputeMessage**](PaypalDisputeMessage.md) |  | [optional] 
**LifeCycleStage** | Pointer to **string** | The stage in the dispute lifecycle. | [optional] 
**SellerResponseDueDate** | Pointer to **int32** | Within which date the seller needs to respond. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the dispute. | [optional] 
**UpdatedAt** | Pointer to **int32** | When the dispute was updated. | [optional] 

## Methods

### NewPaypalDispute

`func NewPaypalDispute() *PaypalDispute`

NewPaypalDispute instantiates a new PaypalDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalDisputeWithDefaults

`func NewPaypalDisputeWithDefaults() *PaypalDispute`

NewPaypalDisputeWithDefaults instantiates a new PaypalDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaypalDispute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaypalDispute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaypalDispute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaypalDispute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaypalDispute) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaypalDispute) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaypalDispute) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaypalDispute) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetShopId

`func (o *PaypalDispute) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *PaypalDispute) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *PaypalDispute) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *PaypalDispute) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetReason

`func (o *PaypalDispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PaypalDispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PaypalDispute) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PaypalDispute) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *PaypalDispute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaypalDispute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaypalDispute) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaypalDispute) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOutcome

`func (o *PaypalDispute) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *PaypalDispute) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *PaypalDispute) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *PaypalDispute) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetMessages

`func (o *PaypalDispute) GetMessages() []PaypalDisputeMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PaypalDispute) GetMessagesOk() (*[]PaypalDisputeMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PaypalDispute) SetMessages(v []PaypalDisputeMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *PaypalDispute) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetLifeCycleStage

`func (o *PaypalDispute) GetLifeCycleStage() string`

GetLifeCycleStage returns the LifeCycleStage field if non-nil, zero value otherwise.

### GetLifeCycleStageOk

`func (o *PaypalDispute) GetLifeCycleStageOk() (*string, bool)`

GetLifeCycleStageOk returns a tuple with the LifeCycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStage

`func (o *PaypalDispute) SetLifeCycleStage(v string)`

SetLifeCycleStage sets LifeCycleStage field to given value.

### HasLifeCycleStage

`func (o *PaypalDispute) HasLifeCycleStage() bool`

HasLifeCycleStage returns a boolean if a field has been set.

### GetSellerResponseDueDate

`func (o *PaypalDispute) GetSellerResponseDueDate() int32`

GetSellerResponseDueDate returns the SellerResponseDueDate field if non-nil, zero value otherwise.

### GetSellerResponseDueDateOk

`func (o *PaypalDispute) GetSellerResponseDueDateOk() (*int32, bool)`

GetSellerResponseDueDateOk returns a tuple with the SellerResponseDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerResponseDueDate

`func (o *PaypalDispute) SetSellerResponseDueDate(v int32)`

SetSellerResponseDueDate sets SellerResponseDueDate field to given value.

### HasSellerResponseDueDate

`func (o *PaypalDispute) HasSellerResponseDueDate() bool`

HasSellerResponseDueDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaypalDispute) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaypalDispute) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaypalDispute) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaypalDispute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaypalDispute) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaypalDispute) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaypalDispute) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaypalDispute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


