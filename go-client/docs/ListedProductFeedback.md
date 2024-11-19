# ListedProductFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Count of all the feedback. | [optional] 
**Positive** | Pointer to **int32** | Count of positive feedback. | [optional] 
**Neutral** | Pointer to **int32** | Count of neutral feedback. | [optional] 
**Negative** | Pointer to **int32** | Count of negative feedback. | [optional] 
**Numbers** | Pointer to **map[string]int32** |  | [optional] 
**List** | Pointer to [**[]ListedProductFeedbackListInner**](ListedProductFeedbackListInner.md) |  | [optional] 

## Methods

### NewListedProductFeedback

`func NewListedProductFeedback() *ListedProductFeedback`

NewListedProductFeedback instantiates a new ListedProductFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedProductFeedbackWithDefaults

`func NewListedProductFeedbackWithDefaults() *ListedProductFeedback`

NewListedProductFeedbackWithDefaults instantiates a new ListedProductFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListedProductFeedback) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListedProductFeedback) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListedProductFeedback) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListedProductFeedback) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPositive

`func (o *ListedProductFeedback) GetPositive() int32`

GetPositive returns the Positive field if non-nil, zero value otherwise.

### GetPositiveOk

`func (o *ListedProductFeedback) GetPositiveOk() (*int32, bool)`

GetPositiveOk returns a tuple with the Positive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositive

`func (o *ListedProductFeedback) SetPositive(v int32)`

SetPositive sets Positive field to given value.

### HasPositive

`func (o *ListedProductFeedback) HasPositive() bool`

HasPositive returns a boolean if a field has been set.

### GetNeutral

`func (o *ListedProductFeedback) GetNeutral() int32`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *ListedProductFeedback) GetNeutralOk() (*int32, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *ListedProductFeedback) SetNeutral(v int32)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *ListedProductFeedback) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetNegative

`func (o *ListedProductFeedback) GetNegative() int32`

GetNegative returns the Negative field if non-nil, zero value otherwise.

### GetNegativeOk

`func (o *ListedProductFeedback) GetNegativeOk() (*int32, bool)`

GetNegativeOk returns a tuple with the Negative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegative

`func (o *ListedProductFeedback) SetNegative(v int32)`

SetNegative sets Negative field to given value.

### HasNegative

`func (o *ListedProductFeedback) HasNegative() bool`

HasNegative returns a boolean if a field has been set.

### GetNumbers

`func (o *ListedProductFeedback) GetNumbers() map[string]int32`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ListedProductFeedback) GetNumbersOk() (*map[string]int32, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ListedProductFeedback) SetNumbers(v map[string]int32)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *ListedProductFeedback) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetList

`func (o *ListedProductFeedback) GetList() []ListedProductFeedbackListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ListedProductFeedback) GetListOk() (*[]ListedProductFeedbackListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ListedProductFeedback) SetList(v []ListedProductFeedbackListInner)`

SetList sets List field to given value.

### HasList

`func (o *ListedProductFeedback) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


