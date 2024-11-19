# Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**ProductId** | Pointer to **string** | Unique ID of the product for which this feedback has been posted. | [optional] 
**InvoiceId** | Pointer to **string** | Unique ID of the invoice for which this feedback has been posted. | [optional] 
**Blocked** | Pointer to **bool** | If true, this feedback has been blocked after an appeal. | [optional] 
**Appealed** | Pointer to **bool** | If true, an appeal has been created for this feedback. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this feedback belongs. | [optional] 
**Message** | Pointer to **string** | Message left by the customer. | [optional] 
**Reply** | Pointer to **string** | Reply left by the merchant. | [optional] 
**Score** | Pointer to **int32** | Score left by the customer, if 0 no score has been left. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the feedback. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**ProductTitle** | Pointer to **string** | Product title for which this feedback has been created | [optional] 
**ProductImageName** | Pointer to **string** | Unique id of the image attachment for this product with the image extension included. | [optional] 
**ProductImageStorage** | Pointer to **string** | Where the image is stored in our self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**Invoice** | Pointer to [**MiniInvoice**](MiniInvoice.md) |  | [optional] 
**Product** | Pointer to [**FeedbackProduct**](FeedbackProduct.md) |  | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the feedback. | [optional] 

## Methods

### NewFeedback

`func NewFeedback() *Feedback`

NewFeedback instantiates a new Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithDefaults

`func NewFeedbackWithDefaults() *Feedback`

NewFeedbackWithDefaults instantiates a new Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqid

`func (o *Feedback) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Feedback) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Feedback) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Feedback) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetProductId

`func (o *Feedback) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Feedback) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Feedback) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Feedback) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Feedback) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Feedback) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Feedback) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Feedback) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetBlocked

`func (o *Feedback) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Feedback) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Feedback) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *Feedback) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetAppealed

`func (o *Feedback) GetAppealed() bool`

GetAppealed returns the Appealed field if non-nil, zero value otherwise.

### GetAppealedOk

`func (o *Feedback) GetAppealedOk() (*bool, bool)`

GetAppealedOk returns a tuple with the Appealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppealed

`func (o *Feedback) SetAppealed(v bool)`

SetAppealed sets Appealed field to given value.

### HasAppealed

`func (o *Feedback) HasAppealed() bool`

HasAppealed returns a boolean if a field has been set.

### GetShopId

`func (o *Feedback) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Feedback) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Feedback) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Feedback) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetMessage

`func (o *Feedback) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Feedback) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Feedback) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Feedback) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReply

`func (o *Feedback) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *Feedback) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *Feedback) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *Feedback) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetScore

`func (o *Feedback) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Feedback) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Feedback) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Feedback) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Feedback) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Feedback) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Feedback) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Feedback) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Feedback) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Feedback) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Feedback) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Feedback) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProductTitle

`func (o *Feedback) GetProductTitle() string`

GetProductTitle returns the ProductTitle field if non-nil, zero value otherwise.

### GetProductTitleOk

`func (o *Feedback) GetProductTitleOk() (*string, bool)`

GetProductTitleOk returns a tuple with the ProductTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTitle

`func (o *Feedback) SetProductTitle(v string)`

SetProductTitle sets ProductTitle field to given value.

### HasProductTitle

`func (o *Feedback) HasProductTitle() bool`

HasProductTitle returns a boolean if a field has been set.

### GetProductImageName

`func (o *Feedback) GetProductImageName() string`

GetProductImageName returns the ProductImageName field if non-nil, zero value otherwise.

### GetProductImageNameOk

`func (o *Feedback) GetProductImageNameOk() (*string, bool)`

GetProductImageNameOk returns a tuple with the ProductImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageName

`func (o *Feedback) SetProductImageName(v string)`

SetProductImageName sets ProductImageName field to given value.

### HasProductImageName

`func (o *Feedback) HasProductImageName() bool`

HasProductImageName returns a boolean if a field has been set.

### GetProductImageStorage

`func (o *Feedback) GetProductImageStorage() string`

GetProductImageStorage returns the ProductImageStorage field if non-nil, zero value otherwise.

### GetProductImageStorageOk

`func (o *Feedback) GetProductImageStorageOk() (*string, bool)`

GetProductImageStorageOk returns a tuple with the ProductImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageStorage

`func (o *Feedback) SetProductImageStorage(v string)`

SetProductImageStorage sets ProductImageStorage field to given value.

### HasProductImageStorage

`func (o *Feedback) HasProductImageStorage() bool`

HasProductImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *Feedback) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *Feedback) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *Feedback) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *Feedback) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetInvoice

`func (o *Feedback) GetInvoice() MiniInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Feedback) GetInvoiceOk() (*MiniInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Feedback) SetInvoice(v MiniInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *Feedback) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetProduct

`func (o *Feedback) GetProduct() FeedbackProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Feedback) GetProductOk() (*FeedbackProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Feedback) SetProduct(v FeedbackProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Feedback) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Feedback) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Feedback) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Feedback) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Feedback) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


