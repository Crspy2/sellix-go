# FeedbackProductFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Count of all the feedback. | [optional] 
**Positive** | Pointer to **int32** | Count of positive feedback. | [optional] 
**Neutral** | Pointer to **int32** | Count of neutral feedback. | [optional] 
**Negative** | Pointer to **int32** | Count of negative feedback. | [optional] 
**Numbers** | Pointer to **map[string]int32** |  | [optional] 
**List** | Pointer to [**[]FeedbackProductFeedbackListInner**](FeedbackProductFeedbackListInner.md) |  | [optional] 

## Methods

### NewFeedbackProductFeedback

`func NewFeedbackProductFeedback() *FeedbackProductFeedback`

NewFeedbackProductFeedback instantiates a new FeedbackProductFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackProductFeedbackWithDefaults

`func NewFeedbackProductFeedbackWithDefaults() *FeedbackProductFeedback`

NewFeedbackProductFeedbackWithDefaults instantiates a new FeedbackProductFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *FeedbackProductFeedback) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FeedbackProductFeedback) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FeedbackProductFeedback) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FeedbackProductFeedback) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPositive

`func (o *FeedbackProductFeedback) GetPositive() int32`

GetPositive returns the Positive field if non-nil, zero value otherwise.

### GetPositiveOk

`func (o *FeedbackProductFeedback) GetPositiveOk() (*int32, bool)`

GetPositiveOk returns a tuple with the Positive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositive

`func (o *FeedbackProductFeedback) SetPositive(v int32)`

SetPositive sets Positive field to given value.

### HasPositive

`func (o *FeedbackProductFeedback) HasPositive() bool`

HasPositive returns a boolean if a field has been set.

### GetNeutral

`func (o *FeedbackProductFeedback) GetNeutral() int32`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *FeedbackProductFeedback) GetNeutralOk() (*int32, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *FeedbackProductFeedback) SetNeutral(v int32)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *FeedbackProductFeedback) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetNegative

`func (o *FeedbackProductFeedback) GetNegative() int32`

GetNegative returns the Negative field if non-nil, zero value otherwise.

### GetNegativeOk

`func (o *FeedbackProductFeedback) GetNegativeOk() (*int32, bool)`

GetNegativeOk returns a tuple with the Negative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegative

`func (o *FeedbackProductFeedback) SetNegative(v int32)`

SetNegative sets Negative field to given value.

### HasNegative

`func (o *FeedbackProductFeedback) HasNegative() bool`

HasNegative returns a boolean if a field has been set.

### GetNumbers

`func (o *FeedbackProductFeedback) GetNumbers() map[string]int32`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *FeedbackProductFeedback) GetNumbersOk() (*map[string]int32, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *FeedbackProductFeedback) SetNumbers(v map[string]int32)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *FeedbackProductFeedback) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetList

`func (o *FeedbackProductFeedback) GetList() []FeedbackProductFeedbackListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *FeedbackProductFeedback) GetListOk() (*[]FeedbackProductFeedbackListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *FeedbackProductFeedback) SetList(v []FeedbackProductFeedbackListInner)`

SetList sets List field to given value.

### HasList

`func (o *FeedbackProductFeedback) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


