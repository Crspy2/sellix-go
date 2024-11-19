# FeedbackListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice for which this feedback has been posted. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product for which this feedback has been posted. | [optional] 
**ProductTitle** | Pointer to **string** | Title of the product for which this feedback has been posted. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this feedback belongs. | [optional] 
**Message** | Pointer to **string** | Message left by the customer. | [optional] 
**Reply** | Pointer to **string** | Reply left by the merchant. | [optional] 
**Score** | Pointer to **int32** | Score left by the customer, if 0 no score has been left. | [optional] 
**Appealed** | Pointer to **bool** | If true, an appeal has been created for this feedback. | [optional] 
**AppealOutcome** | Pointer to **string** | The outcome of the appeal. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the feedback. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**Blocked** | Pointer to **bool** | If true, this feedback has been blocked after an appeal. | [optional] 

## Methods

### NewFeedbackListing

`func NewFeedbackListing() *FeedbackListing`

NewFeedbackListing instantiates a new FeedbackListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackListingWithDefaults

`func NewFeedbackListingWithDefaults() *FeedbackListing`

NewFeedbackListingWithDefaults instantiates a new FeedbackListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeedbackListing) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedbackListing) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedbackListing) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FeedbackListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *FeedbackListing) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *FeedbackListing) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *FeedbackListing) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *FeedbackListing) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetInvoiceId

`func (o *FeedbackListing) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *FeedbackListing) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *FeedbackListing) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *FeedbackListing) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetProductId

`func (o *FeedbackListing) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *FeedbackListing) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *FeedbackListing) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *FeedbackListing) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductTitle

`func (o *FeedbackListing) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *FeedbackListing) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *FeedbackListing) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *FeedbackListing) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetShopId

`func (o *FeedbackListing) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *FeedbackListing) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *FeedbackListing) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *FeedbackListing) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetMessage

`func (o *FeedbackListing) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FeedbackListing) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FeedbackListing) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FeedbackListing) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReply

`func (o *FeedbackListing) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *FeedbackListing) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *FeedbackListing) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *FeedbackListing) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetScore

`func (o *FeedbackListing) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FeedbackListing) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FeedbackListing) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *FeedbackListing) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetAppealed

`func (o *FeedbackListing) GetAppealed() bool`

GetAppealed returns the Appealed field if non-nil, zero value otherwise.

### GetAppealedOk

`func (o *FeedbackListing) GetAppealedOk() (*bool, bool)`

GetAppealedOk returns a tuple with the Appealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppealed

`func (o *FeedbackListing) SetAppealed(v bool)`

SetAppealed sets Appealed field to given value.

### HasAppealed

`func (o *FeedbackListing) HasAppealed() bool`

HasAppealed returns a boolean if a field has been set.

### GetAppealOutcome

`func (o *FeedbackListing) GetAppealOutcome() string`

GetAppealOutcome returns the AppealOutcome field if non-nil, zero value otherwise.

### GetAppealOutcomeOk

`func (o *FeedbackListing) GetAppealOutcomeOk() (*string, bool)`

GetAppealOutcomeOk returns a tuple with the AppealOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppealOutcome

`func (o *FeedbackListing) SetAppealOutcome(v string)`

SetAppealOutcome sets AppealOutcome field to given value.

### HasAppealOutcome

`func (o *FeedbackListing) HasAppealOutcome() bool`

HasAppealOutcome returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedbackListing) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedbackListing) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedbackListing) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeedbackListing) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FeedbackListing) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeedbackListing) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeedbackListing) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FeedbackListing) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBlocked

`func (o *FeedbackListing) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *FeedbackListing) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *FeedbackListing) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *FeedbackListing) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


