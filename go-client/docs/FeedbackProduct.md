# FeedbackProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the resource | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Slug** | Pointer to **string** | The slug used to navigate to the product page on a storefront. | [optional] 
**ShopId** | Pointer to **int32** | The shop ID to which this invoice belongs. | [optional] 
**Type** | Pointer to **string** | Product type. | [optional] 
**Subtype** | Pointer to **string** | Product subtype, can be used only with type SUBSCRIPTION. | [optional] 
**Title** | Pointer to **string** | Product title. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**PayWhatYouWant** | Pointer to **int32** | Whether or not pay-what-you-want pricing is enabled. | [optional] 
**Price** | Pointer to **float64** | Product price. | [optional] 
**PriceDisplay** | Pointer to **float64** | Product price in currency. | [optional] 
**PriceDiscount** | Pointer to **float32** | The discount price on the purchased goods. | [optional] 
**AffiliateRevenuePercent** | Pointer to **float32** | The percentage of revenue to give to affiliates. -1 for disabled, 0 for default (10%), other number for a custom percent. | [optional] 
**PriceVariants** | Pointer to **map[string]interface{}** | Variants in pricing for the product. | [optional] 
**Description** | Pointer to **string** | Product description. | [optional] 
**LicensingEnabled** | Pointer to **int32** | Whether product licenses are generated on successful payments | [optional] 
**LicensePeriod** | Pointer to **int32** | Amount in days that licenses will be valid for | [optional] 
**ImageAttachment** | Pointer to [**ImageAttachment**](ImageAttachment.md) |  | [optional] 
**FileAttachment** | Pointer to **map[string]interface{}** | The file to deliver to the customer (if applicable) | [optional] 
**YoutubeLink** | Pointer to **string** | The url for the YouTube video of the product | [optional] 
**VolumeDiscounts** | Pointer to **string** | Stringified array of volume discounts. | [optional] 
**RecurringInterval** | Pointer to **string** | At which frequency the customer is billed for product type SUBSCRIPTION. | [optional] 
**RecurringIntervalCount** | Pointer to **int32** | How many recurring_interval before the customer is billed for product type SUBSCRIPTION. | [optional] 
**TrialPeriod** | Pointer to **int32** | Defines a trial period before billing the customer for product type SUBSCRIPTION. | [optional] 
**PaypalProductId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway PAYPAL chosen, a PayPal product is automatically created on the connected account. | [optional] 
**PaypalPlanId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway PAYPAL chosen, a PayPal plan is automatically created on the connected account. | [optional] 
**StripePriceId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway STRIPE chosen, a Stripe price is automatically created on the connected account. | [optional] 
**DiscordIntegration** | Pointer to **float32** | Whether or not the discord integeration is enabled for this product. | [optional] 
**DiscordOptional** | Pointer to **float32** | Whether or not the discord integration is optional. | [optional] 
**DiscordSetRole** | Pointer to **float32** | Whether to give users a role when added to the discord server. | [optional] 
**DiscordServerId** | Pointer to **string** | The id of the discord server the bot will add users to. | [optional] 
**DiscordRoleId** | Pointer to **float32** | The role to give users when added by the discord integration. | [optional] 
**DiscordRemoveRole** | Pointer to **float32** | Whether to remove a role from the user when added to the discord server. | [optional] 
**QuantityMin** | Pointer to **int32** | Minimum quantity purchasable of this product. | [optional] 
**QuantityMax** | Pointer to **int32** | Maximum quantity purchasable of this product. | [optional] 
**QuantityWarning** | Pointer to **int32** | At which product quantity should we send a webhook and email warning the merchant. | [optional] 
**Gateways** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldsArrayInner**](CustomFieldsArrayInner.md) |  | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Minimum number of confirmations for a crypto payment to be accepted. | [optional] 
**MaxRiskLevel** | Pointer to **int32** | The max fraud score to allow a customer to have when purchasing the product | [optional] 
**BlockVpnProxies** | Pointer to **bool** | Block VPN and Proxy purchases if the gateway is PAYPAL or STRIPE. | [optional] 
**DeliveryText** | Pointer to **string** | Delivery text | [optional] 
**DeliveryTime** | Pointer to **int32** | The timestamp for when the invoice was delivered | [optional] 
**ServiceText** | Pointer to **string** | If type is SERVICE, then this will be delivered as the product when an invoice is processed | [optional] 
**StockDelimiter** | Pointer to **string** | How to delimit the stock if product type is SERIALS, for example with stock_delimiter \&quot;,\&quot; and serials value first,second; the stock would have two different serials \&quot;first\&quot; and \&quot;second\&quot;. | [optional] 
**Stock** | Pointer to **int32** | Stock of the current product, can be -1 for unlimited stock. | [optional] 
**Bestseller** | Pointer to **float32** | DEPRECATED | [optional] 
**SortPriority** | Pointer to **int32** | Sort order of this product. | [optional] 
**Unlisted** | Pointer to **bool** | If unlisted is true, the product is not shown in the storefront but can be bought through a direct link. | [optional] 
**OnHold** | Pointer to **bool** | If on_hold is true, the product cannot be bought but is shown in the storefront. | [optional] 
**TermsOfService** | Pointer to **string** | Text containing the product&#39;s terms of service. | [optional] 
**Warranty** | Pointer to **int32** | Time, in seconds, of how much the warranty for this product lasts. | [optional] 
**WarrantyText** | Pointer to **string** | Clear explanation of what the warranty covers. | [optional] 
**WatermarkEnabled** | Pointer to **float32** | Whether sellix should add a watermark to your product image. | [optional] 
**WatermarkText** | Pointer to **string** | The watermark to add to the product image. | [optional] 
**RedirectLink** | Pointer to **string** | The url to redirect a customer to on successful payment | [optional] 
**LabelSingular** | Pointer to **string** |  | [optional] 
**LabelPlural** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** | If private is true, the product is hidden on the storefront and cannot be bought with a direct link. | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the feedback. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the feedback. | [optional] 
**MarketplaceCategoryId** | Pointer to **string** | The category ID the product is a part of | [optional] 
**TelegramGroupId** | Pointer to **string** |  | [optional] 
**TelegramIntegration** | Pointer to **float32** |  | [optional] 
**TelegramOptional** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** | The name of the store the invoice was created on | [optional] 
**ImageName** | Pointer to **string** | DEPRECATED: The name of the image for the store the invoice was created on | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**ImageAttachments** | Pointer to [**[]ImageAttachment**](ImageAttachment.md) |  | [optional] 
**Feedback** | Pointer to [**FeedbackProductFeedback**](FeedbackProductFeedback.md) |  | [optional] 
**Categories** | Pointer to [**[]FeedbackProductCategoriesInner**](FeedbackProductCategoriesInner.md) |  | [optional] 
**PaymentGatewaysFees** | Pointer to [**[]FeedbackProductPaymentGatewaysFeesInner**](FeedbackProductPaymentGatewaysFeesInner.md) |  | [optional] 
**PriceConversions** | Pointer to [**FeedbackProductPriceConversions**](FeedbackProductPriceConversions.md) |  | [optional] 

## Methods

### NewFeedbackProduct

`func NewFeedbackProduct() *FeedbackProduct`

NewFeedbackProduct instantiates a new FeedbackProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackProductWithDefaults

`func NewFeedbackProductWithDefaults() *FeedbackProduct`

NewFeedbackProductWithDefaults instantiates a new FeedbackProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeedbackProduct) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedbackProduct) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedbackProduct) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FeedbackProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *FeedbackProduct) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *FeedbackProduct) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *FeedbackProduct) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *FeedbackProduct) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetSlug

`func (o *FeedbackProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FeedbackProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FeedbackProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *FeedbackProduct) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetShopId

`func (o *FeedbackProduct) GetShopId() int32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *FeedbackProduct) GetShopIdOk() (*int32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *FeedbackProduct) SetShopId(v int32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *FeedbackProduct) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *FeedbackProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeedbackProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeedbackProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeedbackProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *FeedbackProduct) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *FeedbackProduct) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *FeedbackProduct) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *FeedbackProduct) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTitle

`func (o *FeedbackProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedbackProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeedbackProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FeedbackProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCurrency

`func (o *FeedbackProduct) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeedbackProduct) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeedbackProduct) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeedbackProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPayWhatYouWant

`func (o *FeedbackProduct) GetPayWhatYouWant() int32`

GetPayWhatYouWant returns the PayWhatYouWant field if non-nil, zero value otherwise.

### GetPayWhatYouWantOk

`func (o *FeedbackProduct) GetPayWhatYouWantOk() (*int32, bool)`

GetPayWhatYouWantOk returns a tuple with the PayWhatYouWant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayWhatYouWant

`func (o *FeedbackProduct) SetPayWhatYouWant(v int32)`

SetPayWhatYouWant sets PayWhatYouWant field to given value.

### HasPayWhatYouWant

`func (o *FeedbackProduct) HasPayWhatYouWant() bool`

HasPayWhatYouWant returns a boolean if a field has been set.

### GetPrice

`func (o *FeedbackProduct) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FeedbackProduct) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FeedbackProduct) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FeedbackProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplay

`func (o *FeedbackProduct) GetPriceDisplay() float64`

GetPriceDisplay returns the PriceDisplay field if non-nil, zero value otherwise.

### GetPriceDisplayOk

`func (o *FeedbackProduct) GetPriceDisplayOk() (*float64, bool)`

GetPriceDisplayOk returns a tuple with the PriceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplay

`func (o *FeedbackProduct) SetPriceDisplay(v float64)`

SetPriceDisplay sets PriceDisplay field to given value.

### HasPriceDisplay

`func (o *FeedbackProduct) HasPriceDisplay() bool`

HasPriceDisplay returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *FeedbackProduct) GetPriceDiscount() float32`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *FeedbackProduct) GetPriceDiscountOk() (*float32, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *FeedbackProduct) SetPriceDiscount(v float32)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *FeedbackProduct) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetAffiliateRevenuePercent

`func (o *FeedbackProduct) GetAffiliateRevenuePercent() float32`

GetAffiliateRevenuePercent returns the AffiliateRevenuePercent field if non-nil, zero value otherwise.

### GetAffiliateRevenuePercentOk

`func (o *FeedbackProduct) GetAffiliateRevenuePercentOk() (*float32, bool)`

GetAffiliateRevenuePercentOk returns a tuple with the AffiliateRevenuePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenuePercent

`func (o *FeedbackProduct) SetAffiliateRevenuePercent(v float32)`

SetAffiliateRevenuePercent sets AffiliateRevenuePercent field to given value.

### HasAffiliateRevenuePercent

`func (o *FeedbackProduct) HasAffiliateRevenuePercent() bool`

HasAffiliateRevenuePercent returns a boolean if a field has been set.

### GetPriceVariants

`func (o *FeedbackProduct) GetPriceVariants() map[string]interface{}`

GetPriceVariants returns the PriceVariants field if non-nil, zero value otherwise.

### GetPriceVariantsOk

`func (o *FeedbackProduct) GetPriceVariantsOk() (*map[string]interface{}, bool)`

GetPriceVariantsOk returns a tuple with the PriceVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVariants

`func (o *FeedbackProduct) SetPriceVariants(v map[string]interface{})`

SetPriceVariants sets PriceVariants field to given value.

### HasPriceVariants

`func (o *FeedbackProduct) HasPriceVariants() bool`

HasPriceVariants returns a boolean if a field has been set.

### GetDescription

`func (o *FeedbackProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeedbackProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeedbackProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeedbackProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicensingEnabled

`func (o *FeedbackProduct) GetLicensingEnabled() int32`

GetLicensingEnabled returns the LicensingEnabled field if non-nil, zero value otherwise.

### GetLicensingEnabledOk

`func (o *FeedbackProduct) GetLicensingEnabledOk() (*int32, bool)`

GetLicensingEnabledOk returns a tuple with the LicensingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingEnabled

`func (o *FeedbackProduct) SetLicensingEnabled(v int32)`

SetLicensingEnabled sets LicensingEnabled field to given value.

### HasLicensingEnabled

`func (o *FeedbackProduct) HasLicensingEnabled() bool`

HasLicensingEnabled returns a boolean if a field has been set.

### GetLicensePeriod

`func (o *FeedbackProduct) GetLicensePeriod() int32`

GetLicensePeriod returns the LicensePeriod field if non-nil, zero value otherwise.

### GetLicensePeriodOk

`func (o *FeedbackProduct) GetLicensePeriodOk() (*int32, bool)`

GetLicensePeriodOk returns a tuple with the LicensePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePeriod

`func (o *FeedbackProduct) SetLicensePeriod(v int32)`

SetLicensePeriod sets LicensePeriod field to given value.

### HasLicensePeriod

`func (o *FeedbackProduct) HasLicensePeriod() bool`

HasLicensePeriod returns a boolean if a field has been set.

### GetImageAttachment

`func (o *FeedbackProduct) GetImageAttachment() ImageAttachment`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *FeedbackProduct) GetImageAttachmentOk() (*ImageAttachment, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *FeedbackProduct) SetImageAttachment(v ImageAttachment)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *FeedbackProduct) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetFileAttachment

`func (o *FeedbackProduct) GetFileAttachment() map[string]interface{}`

GetFileAttachment returns the FileAttachment field if non-nil, zero value otherwise.

### GetFileAttachmentOk

`func (o *FeedbackProduct) GetFileAttachmentOk() (*map[string]interface{}, bool)`

GetFileAttachmentOk returns a tuple with the FileAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAttachment

`func (o *FeedbackProduct) SetFileAttachment(v map[string]interface{})`

SetFileAttachment sets FileAttachment field to given value.

### HasFileAttachment

`func (o *FeedbackProduct) HasFileAttachment() bool`

HasFileAttachment returns a boolean if a field has been set.

### GetYoutubeLink

`func (o *FeedbackProduct) GetYoutubeLink() string`

GetYoutubeLink returns the YoutubeLink field if non-nil, zero value otherwise.

### GetYoutubeLinkOk

`func (o *FeedbackProduct) GetYoutubeLinkOk() (*string, bool)`

GetYoutubeLinkOk returns a tuple with the YoutubeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeLink

`func (o *FeedbackProduct) SetYoutubeLink(v string)`

SetYoutubeLink sets YoutubeLink field to given value.

### HasYoutubeLink

`func (o *FeedbackProduct) HasYoutubeLink() bool`

HasYoutubeLink returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *FeedbackProduct) GetVolumeDiscounts() string`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *FeedbackProduct) GetVolumeDiscountsOk() (*string, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *FeedbackProduct) SetVolumeDiscounts(v string)`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *FeedbackProduct) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *FeedbackProduct) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *FeedbackProduct) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *FeedbackProduct) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *FeedbackProduct) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *FeedbackProduct) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *FeedbackProduct) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *FeedbackProduct) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *FeedbackProduct) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *FeedbackProduct) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *FeedbackProduct) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *FeedbackProduct) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *FeedbackProduct) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetPaypalProductId

`func (o *FeedbackProduct) GetPaypalProductId() string`

GetPaypalProductId returns the PaypalProductId field if non-nil, zero value otherwise.

### GetPaypalProductIdOk

`func (o *FeedbackProduct) GetPaypalProductIdOk() (*string, bool)`

GetPaypalProductIdOk returns a tuple with the PaypalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalProductId

`func (o *FeedbackProduct) SetPaypalProductId(v string)`

SetPaypalProductId sets PaypalProductId field to given value.

### HasPaypalProductId

`func (o *FeedbackProduct) HasPaypalProductId() bool`

HasPaypalProductId returns a boolean if a field has been set.

### GetPaypalPlanId

`func (o *FeedbackProduct) GetPaypalPlanId() string`

GetPaypalPlanId returns the PaypalPlanId field if non-nil, zero value otherwise.

### GetPaypalPlanIdOk

`func (o *FeedbackProduct) GetPaypalPlanIdOk() (*string, bool)`

GetPaypalPlanIdOk returns a tuple with the PaypalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPlanId

`func (o *FeedbackProduct) SetPaypalPlanId(v string)`

SetPaypalPlanId sets PaypalPlanId field to given value.

### HasPaypalPlanId

`func (o *FeedbackProduct) HasPaypalPlanId() bool`

HasPaypalPlanId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *FeedbackProduct) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *FeedbackProduct) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *FeedbackProduct) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *FeedbackProduct) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetDiscordIntegration

`func (o *FeedbackProduct) GetDiscordIntegration() float32`

GetDiscordIntegration returns the DiscordIntegration field if non-nil, zero value otherwise.

### GetDiscordIntegrationOk

`func (o *FeedbackProduct) GetDiscordIntegrationOk() (*float32, bool)`

GetDiscordIntegrationOk returns a tuple with the DiscordIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordIntegration

`func (o *FeedbackProduct) SetDiscordIntegration(v float32)`

SetDiscordIntegration sets DiscordIntegration field to given value.

### HasDiscordIntegration

`func (o *FeedbackProduct) HasDiscordIntegration() bool`

HasDiscordIntegration returns a boolean if a field has been set.

### GetDiscordOptional

`func (o *FeedbackProduct) GetDiscordOptional() float32`

GetDiscordOptional returns the DiscordOptional field if non-nil, zero value otherwise.

### GetDiscordOptionalOk

`func (o *FeedbackProduct) GetDiscordOptionalOk() (*float32, bool)`

GetDiscordOptionalOk returns a tuple with the DiscordOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordOptional

`func (o *FeedbackProduct) SetDiscordOptional(v float32)`

SetDiscordOptional sets DiscordOptional field to given value.

### HasDiscordOptional

`func (o *FeedbackProduct) HasDiscordOptional() bool`

HasDiscordOptional returns a boolean if a field has been set.

### GetDiscordSetRole

`func (o *FeedbackProduct) GetDiscordSetRole() float32`

GetDiscordSetRole returns the DiscordSetRole field if non-nil, zero value otherwise.

### GetDiscordSetRoleOk

`func (o *FeedbackProduct) GetDiscordSetRoleOk() (*float32, bool)`

GetDiscordSetRoleOk returns a tuple with the DiscordSetRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordSetRole

`func (o *FeedbackProduct) SetDiscordSetRole(v float32)`

SetDiscordSetRole sets DiscordSetRole field to given value.

### HasDiscordSetRole

`func (o *FeedbackProduct) HasDiscordSetRole() bool`

HasDiscordSetRole returns a boolean if a field has been set.

### GetDiscordServerId

`func (o *FeedbackProduct) GetDiscordServerId() string`

GetDiscordServerId returns the DiscordServerId field if non-nil, zero value otherwise.

### GetDiscordServerIdOk

`func (o *FeedbackProduct) GetDiscordServerIdOk() (*string, bool)`

GetDiscordServerIdOk returns a tuple with the DiscordServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordServerId

`func (o *FeedbackProduct) SetDiscordServerId(v string)`

SetDiscordServerId sets DiscordServerId field to given value.

### HasDiscordServerId

`func (o *FeedbackProduct) HasDiscordServerId() bool`

HasDiscordServerId returns a boolean if a field has been set.

### GetDiscordRoleId

`func (o *FeedbackProduct) GetDiscordRoleId() float32`

GetDiscordRoleId returns the DiscordRoleId field if non-nil, zero value otherwise.

### GetDiscordRoleIdOk

`func (o *FeedbackProduct) GetDiscordRoleIdOk() (*float32, bool)`

GetDiscordRoleIdOk returns a tuple with the DiscordRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRoleId

`func (o *FeedbackProduct) SetDiscordRoleId(v float32)`

SetDiscordRoleId sets DiscordRoleId field to given value.

### HasDiscordRoleId

`func (o *FeedbackProduct) HasDiscordRoleId() bool`

HasDiscordRoleId returns a boolean if a field has been set.

### GetDiscordRemoveRole

`func (o *FeedbackProduct) GetDiscordRemoveRole() float32`

GetDiscordRemoveRole returns the DiscordRemoveRole field if non-nil, zero value otherwise.

### GetDiscordRemoveRoleOk

`func (o *FeedbackProduct) GetDiscordRemoveRoleOk() (*float32, bool)`

GetDiscordRemoveRoleOk returns a tuple with the DiscordRemoveRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRemoveRole

`func (o *FeedbackProduct) SetDiscordRemoveRole(v float32)`

SetDiscordRemoveRole sets DiscordRemoveRole field to given value.

### HasDiscordRemoveRole

`func (o *FeedbackProduct) HasDiscordRemoveRole() bool`

HasDiscordRemoveRole returns a boolean if a field has been set.

### GetQuantityMin

`func (o *FeedbackProduct) GetQuantityMin() int32`

GetQuantityMin returns the QuantityMin field if non-nil, zero value otherwise.

### GetQuantityMinOk

`func (o *FeedbackProduct) GetQuantityMinOk() (*int32, bool)`

GetQuantityMinOk returns a tuple with the QuantityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMin

`func (o *FeedbackProduct) SetQuantityMin(v int32)`

SetQuantityMin sets QuantityMin field to given value.

### HasQuantityMin

`func (o *FeedbackProduct) HasQuantityMin() bool`

HasQuantityMin returns a boolean if a field has been set.

### GetQuantityMax

`func (o *FeedbackProduct) GetQuantityMax() int32`

GetQuantityMax returns the QuantityMax field if non-nil, zero value otherwise.

### GetQuantityMaxOk

`func (o *FeedbackProduct) GetQuantityMaxOk() (*int32, bool)`

GetQuantityMaxOk returns a tuple with the QuantityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMax

`func (o *FeedbackProduct) SetQuantityMax(v int32)`

SetQuantityMax sets QuantityMax field to given value.

### HasQuantityMax

`func (o *FeedbackProduct) HasQuantityMax() bool`

HasQuantityMax returns a boolean if a field has been set.

### GetQuantityWarning

`func (o *FeedbackProduct) GetQuantityWarning() int32`

GetQuantityWarning returns the QuantityWarning field if non-nil, zero value otherwise.

### GetQuantityWarningOk

`func (o *FeedbackProduct) GetQuantityWarningOk() (*int32, bool)`

GetQuantityWarningOk returns a tuple with the QuantityWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWarning

`func (o *FeedbackProduct) SetQuantityWarning(v int32)`

SetQuantityWarning sets QuantityWarning field to given value.

### HasQuantityWarning

`func (o *FeedbackProduct) HasQuantityWarning() bool`

HasQuantityWarning returns a boolean if a field has been set.

### GetGateways

`func (o *FeedbackProduct) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *FeedbackProduct) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *FeedbackProduct) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *FeedbackProduct) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetCustomFields

`func (o *FeedbackProduct) GetCustomFields() []CustomFieldsArrayInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *FeedbackProduct) GetCustomFieldsOk() (*[]CustomFieldsArrayInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *FeedbackProduct) SetCustomFields(v []CustomFieldsArrayInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *FeedbackProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *FeedbackProduct) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *FeedbackProduct) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *FeedbackProduct) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *FeedbackProduct) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *FeedbackProduct) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *FeedbackProduct) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *FeedbackProduct) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *FeedbackProduct) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *FeedbackProduct) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *FeedbackProduct) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *FeedbackProduct) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *FeedbackProduct) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetDeliveryText

`func (o *FeedbackProduct) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *FeedbackProduct) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *FeedbackProduct) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *FeedbackProduct) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *FeedbackProduct) GetDeliveryTime() int32`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *FeedbackProduct) GetDeliveryTimeOk() (*int32, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *FeedbackProduct) SetDeliveryTime(v int32)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *FeedbackProduct) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetServiceText

`func (o *FeedbackProduct) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *FeedbackProduct) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *FeedbackProduct) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *FeedbackProduct) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *FeedbackProduct) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *FeedbackProduct) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *FeedbackProduct) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *FeedbackProduct) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetStock

`func (o *FeedbackProduct) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *FeedbackProduct) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *FeedbackProduct) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *FeedbackProduct) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetBestseller

`func (o *FeedbackProduct) GetBestseller() float32`

GetBestseller returns the Bestseller field if non-nil, zero value otherwise.

### GetBestsellerOk

`func (o *FeedbackProduct) GetBestsellerOk() (*float32, bool)`

GetBestsellerOk returns a tuple with the Bestseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestseller

`func (o *FeedbackProduct) SetBestseller(v float32)`

SetBestseller sets Bestseller field to given value.

### HasBestseller

`func (o *FeedbackProduct) HasBestseller() bool`

HasBestseller returns a boolean if a field has been set.

### GetSortPriority

`func (o *FeedbackProduct) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *FeedbackProduct) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *FeedbackProduct) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *FeedbackProduct) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetUnlisted

`func (o *FeedbackProduct) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *FeedbackProduct) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *FeedbackProduct) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *FeedbackProduct) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetOnHold

`func (o *FeedbackProduct) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *FeedbackProduct) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *FeedbackProduct) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *FeedbackProduct) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *FeedbackProduct) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *FeedbackProduct) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *FeedbackProduct) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *FeedbackProduct) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *FeedbackProduct) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *FeedbackProduct) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *FeedbackProduct) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *FeedbackProduct) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *FeedbackProduct) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *FeedbackProduct) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *FeedbackProduct) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *FeedbackProduct) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetWatermarkEnabled

`func (o *FeedbackProduct) GetWatermarkEnabled() float32`

GetWatermarkEnabled returns the WatermarkEnabled field if non-nil, zero value otherwise.

### GetWatermarkEnabledOk

`func (o *FeedbackProduct) GetWatermarkEnabledOk() (*float32, bool)`

GetWatermarkEnabledOk returns a tuple with the WatermarkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkEnabled

`func (o *FeedbackProduct) SetWatermarkEnabled(v float32)`

SetWatermarkEnabled sets WatermarkEnabled field to given value.

### HasWatermarkEnabled

`func (o *FeedbackProduct) HasWatermarkEnabled() bool`

HasWatermarkEnabled returns a boolean if a field has been set.

### GetWatermarkText

`func (o *FeedbackProduct) GetWatermarkText() string`

GetWatermarkText returns the WatermarkText field if non-nil, zero value otherwise.

### GetWatermarkTextOk

`func (o *FeedbackProduct) GetWatermarkTextOk() (*string, bool)`

GetWatermarkTextOk returns a tuple with the WatermarkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkText

`func (o *FeedbackProduct) SetWatermarkText(v string)`

SetWatermarkText sets WatermarkText field to given value.

### HasWatermarkText

`func (o *FeedbackProduct) HasWatermarkText() bool`

HasWatermarkText returns a boolean if a field has been set.

### GetRedirectLink

`func (o *FeedbackProduct) GetRedirectLink() string`

GetRedirectLink returns the RedirectLink field if non-nil, zero value otherwise.

### GetRedirectLinkOk

`func (o *FeedbackProduct) GetRedirectLinkOk() (*string, bool)`

GetRedirectLinkOk returns a tuple with the RedirectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectLink

`func (o *FeedbackProduct) SetRedirectLink(v string)`

SetRedirectLink sets RedirectLink field to given value.

### HasRedirectLink

`func (o *FeedbackProduct) HasRedirectLink() bool`

HasRedirectLink returns a boolean if a field has been set.

### GetLabelSingular

`func (o *FeedbackProduct) GetLabelSingular() string`

GetLabelSingular returns the LabelSingular field if non-nil, zero value otherwise.

### GetLabelSingularOk

`func (o *FeedbackProduct) GetLabelSingularOk() (*string, bool)`

GetLabelSingularOk returns a tuple with the LabelSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSingular

`func (o *FeedbackProduct) SetLabelSingular(v string)`

SetLabelSingular sets LabelSingular field to given value.

### HasLabelSingular

`func (o *FeedbackProduct) HasLabelSingular() bool`

HasLabelSingular returns a boolean if a field has been set.

### GetLabelPlural

`func (o *FeedbackProduct) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *FeedbackProduct) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *FeedbackProduct) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *FeedbackProduct) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetPrivate

`func (o *FeedbackProduct) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *FeedbackProduct) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *FeedbackProduct) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *FeedbackProduct) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FeedbackProduct) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeedbackProduct) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeedbackProduct) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeedbackProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FeedbackProduct) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeedbackProduct) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeedbackProduct) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FeedbackProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FeedbackProduct) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FeedbackProduct) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FeedbackProduct) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FeedbackProduct) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetMarketplaceCategoryId

`func (o *FeedbackProduct) GetMarketplaceCategoryId() string`

GetMarketplaceCategoryId returns the MarketplaceCategoryId field if non-nil, zero value otherwise.

### GetMarketplaceCategoryIdOk

`func (o *FeedbackProduct) GetMarketplaceCategoryIdOk() (*string, bool)`

GetMarketplaceCategoryIdOk returns a tuple with the MarketplaceCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceCategoryId

`func (o *FeedbackProduct) SetMarketplaceCategoryId(v string)`

SetMarketplaceCategoryId sets MarketplaceCategoryId field to given value.

### HasMarketplaceCategoryId

`func (o *FeedbackProduct) HasMarketplaceCategoryId() bool`

HasMarketplaceCategoryId returns a boolean if a field has been set.

### GetTelegramGroupId

`func (o *FeedbackProduct) GetTelegramGroupId() string`

GetTelegramGroupId returns the TelegramGroupId field if non-nil, zero value otherwise.

### GetTelegramGroupIdOk

`func (o *FeedbackProduct) GetTelegramGroupIdOk() (*string, bool)`

GetTelegramGroupIdOk returns a tuple with the TelegramGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramGroupId

`func (o *FeedbackProduct) SetTelegramGroupId(v string)`

SetTelegramGroupId sets TelegramGroupId field to given value.

### HasTelegramGroupId

`func (o *FeedbackProduct) HasTelegramGroupId() bool`

HasTelegramGroupId returns a boolean if a field has been set.

### GetTelegramIntegration

`func (o *FeedbackProduct) GetTelegramIntegration() float32`

GetTelegramIntegration returns the TelegramIntegration field if non-nil, zero value otherwise.

### GetTelegramIntegrationOk

`func (o *FeedbackProduct) GetTelegramIntegrationOk() (*float32, bool)`

GetTelegramIntegrationOk returns a tuple with the TelegramIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramIntegration

`func (o *FeedbackProduct) SetTelegramIntegration(v float32)`

SetTelegramIntegration sets TelegramIntegration field to given value.

### HasTelegramIntegration

`func (o *FeedbackProduct) HasTelegramIntegration() bool`

HasTelegramIntegration returns a boolean if a field has been set.

### GetTelegramOptional

`func (o *FeedbackProduct) GetTelegramOptional() float32`

GetTelegramOptional returns the TelegramOptional field if non-nil, zero value otherwise.

### GetTelegramOptionalOk

`func (o *FeedbackProduct) GetTelegramOptionalOk() (*float32, bool)`

GetTelegramOptionalOk returns a tuple with the TelegramOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramOptional

`func (o *FeedbackProduct) SetTelegramOptional(v float32)`

SetTelegramOptional sets TelegramOptional field to given value.

### HasTelegramOptional

`func (o *FeedbackProduct) HasTelegramOptional() bool`

HasTelegramOptional returns a boolean if a field has been set.

### GetName

`func (o *FeedbackProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeedbackProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeedbackProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeedbackProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImageName

`func (o *FeedbackProduct) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *FeedbackProduct) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *FeedbackProduct) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *FeedbackProduct) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *FeedbackProduct) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *FeedbackProduct) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *FeedbackProduct) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *FeedbackProduct) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *FeedbackProduct) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *FeedbackProduct) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *FeedbackProduct) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *FeedbackProduct) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetImageAttachments

`func (o *FeedbackProduct) GetImageAttachments() []ImageAttachment`

GetImageAttachments returns the ImageAttachments field if non-nil, zero value otherwise.

### GetImageAttachmentsOk

`func (o *FeedbackProduct) GetImageAttachmentsOk() (*[]ImageAttachment, bool)`

GetImageAttachmentsOk returns a tuple with the ImageAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachments

`func (o *FeedbackProduct) SetImageAttachments(v []ImageAttachment)`

SetImageAttachments sets ImageAttachments field to given value.

### HasImageAttachments

`func (o *FeedbackProduct) HasImageAttachments() bool`

HasImageAttachments returns a boolean if a field has been set.

### GetFeedback

`func (o *FeedbackProduct) GetFeedback() FeedbackProductFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *FeedbackProduct) GetFeedbackOk() (*FeedbackProductFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *FeedbackProduct) SetFeedback(v FeedbackProductFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *FeedbackProduct) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetCategories

`func (o *FeedbackProduct) GetCategories() []FeedbackProductCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FeedbackProduct) GetCategoriesOk() (*[]FeedbackProductCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FeedbackProduct) SetCategories(v []FeedbackProductCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FeedbackProduct) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPaymentGatewaysFees

`func (o *FeedbackProduct) GetPaymentGatewaysFees() []FeedbackProductPaymentGatewaysFeesInner`

GetPaymentGatewaysFees returns the PaymentGatewaysFees field if non-nil, zero value otherwise.

### GetPaymentGatewaysFeesOk

`func (o *FeedbackProduct) GetPaymentGatewaysFeesOk() (*[]FeedbackProductPaymentGatewaysFeesInner, bool)`

GetPaymentGatewaysFeesOk returns a tuple with the PaymentGatewaysFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewaysFees

`func (o *FeedbackProduct) SetPaymentGatewaysFees(v []FeedbackProductPaymentGatewaysFeesInner)`

SetPaymentGatewaysFees sets PaymentGatewaysFees field to given value.

### HasPaymentGatewaysFees

`func (o *FeedbackProduct) HasPaymentGatewaysFees() bool`

HasPaymentGatewaysFees returns a boolean if a field has been set.

### GetPriceConversions

`func (o *FeedbackProduct) GetPriceConversions() FeedbackProductPriceConversions`

GetPriceConversions returns the PriceConversions field if non-nil, zero value otherwise.

### GetPriceConversionsOk

`func (o *FeedbackProduct) GetPriceConversionsOk() (*FeedbackProductPriceConversions, bool)`

GetPriceConversionsOk returns a tuple with the PriceConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceConversions

`func (o *FeedbackProduct) SetPriceConversions(v FeedbackProductPriceConversions)`

SetPriceConversions sets PriceConversions field to given value.

### HasPriceConversions

`func (o *FeedbackProduct) HasPriceConversions() bool`

HasPriceConversions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


