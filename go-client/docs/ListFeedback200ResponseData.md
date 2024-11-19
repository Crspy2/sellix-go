# ListFeedback200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feedback** | Pointer to [**[]FeedbackListing**](FeedbackListing.md) |  | [optional] 
**Details** | Pointer to [**ListFeedback200ResponseDataDetails**](ListFeedback200ResponseDataDetails.md) |  | [optional] 

## Methods

### NewListFeedback200ResponseData

`func NewListFeedback200ResponseData() *ListFeedback200ResponseData`

NewListFeedback200ResponseData instantiates a new ListFeedback200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeedback200ResponseDataWithDefaults

`func NewListFeedback200ResponseDataWithDefaults() *ListFeedback200ResponseData`

NewListFeedback200ResponseDataWithDefaults instantiates a new ListFeedback200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedback

`func (o *ListFeedback200ResponseData) GetFeedback() []FeedbackListing`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *ListFeedback200ResponseData) GetFeedbackOk() (*[]FeedbackListing, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *ListFeedback200ResponseData) SetFeedback(v []FeedbackListing)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *ListFeedback200ResponseData) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetDetails

`func (o *ListFeedback200ResponseData) GetDetails() ListFeedback200ResponseDataDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListFeedback200ResponseData) GetDetailsOk() (*ListFeedback200ResponseDataDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListFeedback200ResponseData) SetDetails(v ListFeedback200ResponseDataDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListFeedback200ResponseData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


