# ListedProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | Unique ID of the resource, used as reference across the API. | [optional] 
**Uniqid** | Pointer to **string** | Unique ID of the resource, used as reference across the API. | [optional] 
**Slug** | Pointer to **string** | The slug used to navigate to the product page on a storefront. | [optional] 
**ShopId** | Pointer to **float32** | The shop ID to which this resource belongs. | [optional] 
**Type** | Pointer to **string** | Product type. | [optional] 
**Subtype** | Pointer to **string** | Product subtype, can be used only with type SUBSCRIPTION. | [optional] 
**Title** | Pointer to **string** | Product title. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**PayWhatYouWant** | Pointer to **int32** | Whether or not pay-what-you-want pricing is enabled. | [optional] 
**Price** | Pointer to **float64** | Product price. | [optional] 
**PriceDisplay** | Pointer to **float64** | Product price in currency. | [optional] 
**PriceDiscount** | Pointer to **float64** | The discount price on the purchased goods. | [optional] 
**AffiliateRevenuePercent** | Pointer to **float32** | The percentage of revenue to give to affiliates. -1 for disabled, 0 for default (10%), other number for a custom percent. | [optional] 
**PriceVariants** | Pointer to **map[string]interface{}** | Variants in pricing for the product. | [optional] 
**Description** | Pointer to **string** | Product description. | [optional] 
**LicensingEnabled** | Pointer to **int32** | Whether product licenses are generated on successful payments | [optional] 
**LicensePeriod** | Pointer to **int32** | Amount in days that licenses will be valid for | [optional] 
**ImageAttachment** | Pointer to [**ImageAttachment**](ImageAttachment.md) |  | [optional] 
**FileAttachment** | Pointer to **string** | Unique id of the file attachment for this product if the product type is FILE. | [optional] 
**YoutubeLink** | Pointer to **string** | The url for the YouTube video of the product | [optional] 
**VolumeDiscounts** | Pointer to [**[]ListedProductVolumeDiscountsInner**](ListedProductVolumeDiscountsInner.md) | Array of volume discounts. | [optional] 
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
**RconCommands** | Pointer to **string** | Stringified array of the commands to be executed on the RCON server | [optional] 
**RconExecutionType** | Pointer to **string** | The type of execution for the RCON commands. Only applies to RCON products | [optional] 
**RconStartTime** | Pointer to **time.Time** | The time for when the RCON product should start | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Minimum number of confirmations for a crypto payment to be accepted | [optional] 
**MaxRiskLevel** | Pointer to **int32** | For PAYPAL and STRIPE, maximum risk level to accept payments in order to block fraud attempts. | [optional] 
**BlockVpnProxies** | Pointer to **bool** | Block VPN and Proxy purchases if the gateway is PAYPAL or STRIPE. | [optional] 
**DeliveryText** | Pointer to **string** | The text to be delivered to the customer after payment | [optional] 
**DeliveryTime** | Pointer to **int32** | The timestamp for when the invoice was delivered | [optional] 
**ServiceText** | Pointer to **string** | The text to be delivered to the customer after payment for a service | [optional] 
**StockDelimiter** | Pointer to **string** | How to delimit the stock if product type is SERIALS, for example with stock_delimiter \&quot;,\&quot; and serials value first,second; the stock would have two different serials \&quot;first\&quot; and \&quot;second\&quot;. | [optional] 
**Stock** | Pointer to **int32** | Stock of the current product, can be -1 for unlimited stock. | [optional] 
**DynamicWebhook** | Pointer to **map[string]interface{}** |  | [optional] 
**Bestseller** | Pointer to **int32** | DEPRECATED | [optional] 
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
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the product. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the product. | [optional] 
**MarketplaceCategoryId** | Pointer to **string** | The category ID the product is a part of | [optional] 
**TelegramGroupId** | Pointer to **string** | The Telegram group ID | [optional] 
**TelegramIntegration** | Pointer to **int32** | Whether or not the Telegram integration is enabled | [optional] 
**TelegramOptional** | Pointer to **int32** | Whether or not the Telegram integration is optional | [optional] 
**ImageName** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**ImageAttachments** | Pointer to [**[]ImageAttachment**](ImageAttachment.md) |  | [optional] 
**Feedback** | Pointer to [**ListedProductFeedback**](ListedProductFeedback.md) |  | [optional] 
**Categories** | Pointer to [**[]FeedbackProductCategoriesInner**](FeedbackProductCategoriesInner.md) |  | [optional] 
**PaymentGatewaysFees** | Pointer to [**[]ListedProductPaymentGatewaysFeesInner**](ListedProductPaymentGatewaysFeesInner.md) |  | [optional] 

## Methods

### NewListedProduct

`func NewListedProduct() *ListedProduct`

NewListedProduct instantiates a new ListedProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListedProductWithDefaults

`func NewListedProductWithDefaults() *ListedProduct`

NewListedProductWithDefaults instantiates a new ListedProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListedProduct) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListedProduct) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListedProduct) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ListedProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *ListedProduct) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *ListedProduct) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *ListedProduct) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *ListedProduct) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetSlug

`func (o *ListedProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListedProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListedProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListedProduct) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetShopId

`func (o *ListedProduct) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *ListedProduct) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *ListedProduct) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *ListedProduct) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *ListedProduct) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListedProduct) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListedProduct) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListedProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *ListedProduct) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ListedProduct) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ListedProduct) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ListedProduct) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTitle

`func (o *ListedProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListedProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListedProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListedProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCurrency

`func (o *ListedProduct) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListedProduct) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListedProduct) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListedProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPayWhatYouWant

`func (o *ListedProduct) GetPayWhatYouWant() int32`

GetPayWhatYouWant returns the PayWhatYouWant field if non-nil, zero value otherwise.

### GetPayWhatYouWantOk

`func (o *ListedProduct) GetPayWhatYouWantOk() (*int32, bool)`

GetPayWhatYouWantOk returns a tuple with the PayWhatYouWant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayWhatYouWant

`func (o *ListedProduct) SetPayWhatYouWant(v int32)`

SetPayWhatYouWant sets PayWhatYouWant field to given value.

### HasPayWhatYouWant

`func (o *ListedProduct) HasPayWhatYouWant() bool`

HasPayWhatYouWant returns a boolean if a field has been set.

### GetPrice

`func (o *ListedProduct) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListedProduct) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListedProduct) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListedProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplay

`func (o *ListedProduct) GetPriceDisplay() float64`

GetPriceDisplay returns the PriceDisplay field if non-nil, zero value otherwise.

### GetPriceDisplayOk

`func (o *ListedProduct) GetPriceDisplayOk() (*float64, bool)`

GetPriceDisplayOk returns a tuple with the PriceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplay

`func (o *ListedProduct) SetPriceDisplay(v float64)`

SetPriceDisplay sets PriceDisplay field to given value.

### HasPriceDisplay

`func (o *ListedProduct) HasPriceDisplay() bool`

HasPriceDisplay returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *ListedProduct) GetPriceDiscount() float64`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *ListedProduct) GetPriceDiscountOk() (*float64, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *ListedProduct) SetPriceDiscount(v float64)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *ListedProduct) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetAffiliateRevenuePercent

`func (o *ListedProduct) GetAffiliateRevenuePercent() float32`

GetAffiliateRevenuePercent returns the AffiliateRevenuePercent field if non-nil, zero value otherwise.

### GetAffiliateRevenuePercentOk

`func (o *ListedProduct) GetAffiliateRevenuePercentOk() (*float32, bool)`

GetAffiliateRevenuePercentOk returns a tuple with the AffiliateRevenuePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenuePercent

`func (o *ListedProduct) SetAffiliateRevenuePercent(v float32)`

SetAffiliateRevenuePercent sets AffiliateRevenuePercent field to given value.

### HasAffiliateRevenuePercent

`func (o *ListedProduct) HasAffiliateRevenuePercent() bool`

HasAffiliateRevenuePercent returns a boolean if a field has been set.

### GetPriceVariants

`func (o *ListedProduct) GetPriceVariants() map[string]interface{}`

GetPriceVariants returns the PriceVariants field if non-nil, zero value otherwise.

### GetPriceVariantsOk

`func (o *ListedProduct) GetPriceVariantsOk() (*map[string]interface{}, bool)`

GetPriceVariantsOk returns a tuple with the PriceVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVariants

`func (o *ListedProduct) SetPriceVariants(v map[string]interface{})`

SetPriceVariants sets PriceVariants field to given value.

### HasPriceVariants

`func (o *ListedProduct) HasPriceVariants() bool`

HasPriceVariants returns a boolean if a field has been set.

### GetDescription

`func (o *ListedProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListedProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListedProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListedProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicensingEnabled

`func (o *ListedProduct) GetLicensingEnabled() int32`

GetLicensingEnabled returns the LicensingEnabled field if non-nil, zero value otherwise.

### GetLicensingEnabledOk

`func (o *ListedProduct) GetLicensingEnabledOk() (*int32, bool)`

GetLicensingEnabledOk returns a tuple with the LicensingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingEnabled

`func (o *ListedProduct) SetLicensingEnabled(v int32)`

SetLicensingEnabled sets LicensingEnabled field to given value.

### HasLicensingEnabled

`func (o *ListedProduct) HasLicensingEnabled() bool`

HasLicensingEnabled returns a boolean if a field has been set.

### GetLicensePeriod

`func (o *ListedProduct) GetLicensePeriod() int32`

GetLicensePeriod returns the LicensePeriod field if non-nil, zero value otherwise.

### GetLicensePeriodOk

`func (o *ListedProduct) GetLicensePeriodOk() (*int32, bool)`

GetLicensePeriodOk returns a tuple with the LicensePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePeriod

`func (o *ListedProduct) SetLicensePeriod(v int32)`

SetLicensePeriod sets LicensePeriod field to given value.

### HasLicensePeriod

`func (o *ListedProduct) HasLicensePeriod() bool`

HasLicensePeriod returns a boolean if a field has been set.

### GetImageAttachment

`func (o *ListedProduct) GetImageAttachment() ImageAttachment`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *ListedProduct) GetImageAttachmentOk() (*ImageAttachment, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *ListedProduct) SetImageAttachment(v ImageAttachment)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *ListedProduct) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetFileAttachment

`func (o *ListedProduct) GetFileAttachment() string`

GetFileAttachment returns the FileAttachment field if non-nil, zero value otherwise.

### GetFileAttachmentOk

`func (o *ListedProduct) GetFileAttachmentOk() (*string, bool)`

GetFileAttachmentOk returns a tuple with the FileAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAttachment

`func (o *ListedProduct) SetFileAttachment(v string)`

SetFileAttachment sets FileAttachment field to given value.

### HasFileAttachment

`func (o *ListedProduct) HasFileAttachment() bool`

HasFileAttachment returns a boolean if a field has been set.

### GetYoutubeLink

`func (o *ListedProduct) GetYoutubeLink() string`

GetYoutubeLink returns the YoutubeLink field if non-nil, zero value otherwise.

### GetYoutubeLinkOk

`func (o *ListedProduct) GetYoutubeLinkOk() (*string, bool)`

GetYoutubeLinkOk returns a tuple with the YoutubeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeLink

`func (o *ListedProduct) SetYoutubeLink(v string)`

SetYoutubeLink sets YoutubeLink field to given value.

### HasYoutubeLink

`func (o *ListedProduct) HasYoutubeLink() bool`

HasYoutubeLink returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *ListedProduct) GetVolumeDiscounts() []ListedProductVolumeDiscountsInner`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *ListedProduct) GetVolumeDiscountsOk() (*[]ListedProductVolumeDiscountsInner, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *ListedProduct) SetVolumeDiscounts(v []ListedProductVolumeDiscountsInner)`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *ListedProduct) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *ListedProduct) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *ListedProduct) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *ListedProduct) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *ListedProduct) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *ListedProduct) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *ListedProduct) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *ListedProduct) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *ListedProduct) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *ListedProduct) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *ListedProduct) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *ListedProduct) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *ListedProduct) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetPaypalProductId

`func (o *ListedProduct) GetPaypalProductId() string`

GetPaypalProductId returns the PaypalProductId field if non-nil, zero value otherwise.

### GetPaypalProductIdOk

`func (o *ListedProduct) GetPaypalProductIdOk() (*string, bool)`

GetPaypalProductIdOk returns a tuple with the PaypalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalProductId

`func (o *ListedProduct) SetPaypalProductId(v string)`

SetPaypalProductId sets PaypalProductId field to given value.

### HasPaypalProductId

`func (o *ListedProduct) HasPaypalProductId() bool`

HasPaypalProductId returns a boolean if a field has been set.

### GetPaypalPlanId

`func (o *ListedProduct) GetPaypalPlanId() string`

GetPaypalPlanId returns the PaypalPlanId field if non-nil, zero value otherwise.

### GetPaypalPlanIdOk

`func (o *ListedProduct) GetPaypalPlanIdOk() (*string, bool)`

GetPaypalPlanIdOk returns a tuple with the PaypalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPlanId

`func (o *ListedProduct) SetPaypalPlanId(v string)`

SetPaypalPlanId sets PaypalPlanId field to given value.

### HasPaypalPlanId

`func (o *ListedProduct) HasPaypalPlanId() bool`

HasPaypalPlanId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *ListedProduct) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *ListedProduct) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *ListedProduct) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *ListedProduct) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetDiscordIntegration

`func (o *ListedProduct) GetDiscordIntegration() float32`

GetDiscordIntegration returns the DiscordIntegration field if non-nil, zero value otherwise.

### GetDiscordIntegrationOk

`func (o *ListedProduct) GetDiscordIntegrationOk() (*float32, bool)`

GetDiscordIntegrationOk returns a tuple with the DiscordIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordIntegration

`func (o *ListedProduct) SetDiscordIntegration(v float32)`

SetDiscordIntegration sets DiscordIntegration field to given value.

### HasDiscordIntegration

`func (o *ListedProduct) HasDiscordIntegration() bool`

HasDiscordIntegration returns a boolean if a field has been set.

### GetDiscordOptional

`func (o *ListedProduct) GetDiscordOptional() float32`

GetDiscordOptional returns the DiscordOptional field if non-nil, zero value otherwise.

### GetDiscordOptionalOk

`func (o *ListedProduct) GetDiscordOptionalOk() (*float32, bool)`

GetDiscordOptionalOk returns a tuple with the DiscordOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordOptional

`func (o *ListedProduct) SetDiscordOptional(v float32)`

SetDiscordOptional sets DiscordOptional field to given value.

### HasDiscordOptional

`func (o *ListedProduct) HasDiscordOptional() bool`

HasDiscordOptional returns a boolean if a field has been set.

### GetDiscordSetRole

`func (o *ListedProduct) GetDiscordSetRole() float32`

GetDiscordSetRole returns the DiscordSetRole field if non-nil, zero value otherwise.

### GetDiscordSetRoleOk

`func (o *ListedProduct) GetDiscordSetRoleOk() (*float32, bool)`

GetDiscordSetRoleOk returns a tuple with the DiscordSetRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordSetRole

`func (o *ListedProduct) SetDiscordSetRole(v float32)`

SetDiscordSetRole sets DiscordSetRole field to given value.

### HasDiscordSetRole

`func (o *ListedProduct) HasDiscordSetRole() bool`

HasDiscordSetRole returns a boolean if a field has been set.

### GetDiscordServerId

`func (o *ListedProduct) GetDiscordServerId() string`

GetDiscordServerId returns the DiscordServerId field if non-nil, zero value otherwise.

### GetDiscordServerIdOk

`func (o *ListedProduct) GetDiscordServerIdOk() (*string, bool)`

GetDiscordServerIdOk returns a tuple with the DiscordServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordServerId

`func (o *ListedProduct) SetDiscordServerId(v string)`

SetDiscordServerId sets DiscordServerId field to given value.

### HasDiscordServerId

`func (o *ListedProduct) HasDiscordServerId() bool`

HasDiscordServerId returns a boolean if a field has been set.

### GetDiscordRoleId

`func (o *ListedProduct) GetDiscordRoleId() float32`

GetDiscordRoleId returns the DiscordRoleId field if non-nil, zero value otherwise.

### GetDiscordRoleIdOk

`func (o *ListedProduct) GetDiscordRoleIdOk() (*float32, bool)`

GetDiscordRoleIdOk returns a tuple with the DiscordRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRoleId

`func (o *ListedProduct) SetDiscordRoleId(v float32)`

SetDiscordRoleId sets DiscordRoleId field to given value.

### HasDiscordRoleId

`func (o *ListedProduct) HasDiscordRoleId() bool`

HasDiscordRoleId returns a boolean if a field has been set.

### GetDiscordRemoveRole

`func (o *ListedProduct) GetDiscordRemoveRole() float32`

GetDiscordRemoveRole returns the DiscordRemoveRole field if non-nil, zero value otherwise.

### GetDiscordRemoveRoleOk

`func (o *ListedProduct) GetDiscordRemoveRoleOk() (*float32, bool)`

GetDiscordRemoveRoleOk returns a tuple with the DiscordRemoveRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRemoveRole

`func (o *ListedProduct) SetDiscordRemoveRole(v float32)`

SetDiscordRemoveRole sets DiscordRemoveRole field to given value.

### HasDiscordRemoveRole

`func (o *ListedProduct) HasDiscordRemoveRole() bool`

HasDiscordRemoveRole returns a boolean if a field has been set.

### GetQuantityMin

`func (o *ListedProduct) GetQuantityMin() int32`

GetQuantityMin returns the QuantityMin field if non-nil, zero value otherwise.

### GetQuantityMinOk

`func (o *ListedProduct) GetQuantityMinOk() (*int32, bool)`

GetQuantityMinOk returns a tuple with the QuantityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMin

`func (o *ListedProduct) SetQuantityMin(v int32)`

SetQuantityMin sets QuantityMin field to given value.

### HasQuantityMin

`func (o *ListedProduct) HasQuantityMin() bool`

HasQuantityMin returns a boolean if a field has been set.

### GetQuantityMax

`func (o *ListedProduct) GetQuantityMax() int32`

GetQuantityMax returns the QuantityMax field if non-nil, zero value otherwise.

### GetQuantityMaxOk

`func (o *ListedProduct) GetQuantityMaxOk() (*int32, bool)`

GetQuantityMaxOk returns a tuple with the QuantityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMax

`func (o *ListedProduct) SetQuantityMax(v int32)`

SetQuantityMax sets QuantityMax field to given value.

### HasQuantityMax

`func (o *ListedProduct) HasQuantityMax() bool`

HasQuantityMax returns a boolean if a field has been set.

### GetQuantityWarning

`func (o *ListedProduct) GetQuantityWarning() int32`

GetQuantityWarning returns the QuantityWarning field if non-nil, zero value otherwise.

### GetQuantityWarningOk

`func (o *ListedProduct) GetQuantityWarningOk() (*int32, bool)`

GetQuantityWarningOk returns a tuple with the QuantityWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWarning

`func (o *ListedProduct) SetQuantityWarning(v int32)`

SetQuantityWarning sets QuantityWarning field to given value.

### HasQuantityWarning

`func (o *ListedProduct) HasQuantityWarning() bool`

HasQuantityWarning returns a boolean if a field has been set.

### GetGateways

`func (o *ListedProduct) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ListedProduct) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ListedProduct) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ListedProduct) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetCustomFields

`func (o *ListedProduct) GetCustomFields() []CustomFieldsArrayInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ListedProduct) GetCustomFieldsOk() (*[]CustomFieldsArrayInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ListedProduct) SetCustomFields(v []CustomFieldsArrayInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ListedProduct) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRconCommands

`func (o *ListedProduct) GetRconCommands() string`

GetRconCommands returns the RconCommands field if non-nil, zero value otherwise.

### GetRconCommandsOk

`func (o *ListedProduct) GetRconCommandsOk() (*string, bool)`

GetRconCommandsOk returns a tuple with the RconCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconCommands

`func (o *ListedProduct) SetRconCommands(v string)`

SetRconCommands sets RconCommands field to given value.

### HasRconCommands

`func (o *ListedProduct) HasRconCommands() bool`

HasRconCommands returns a boolean if a field has been set.

### GetRconExecutionType

`func (o *ListedProduct) GetRconExecutionType() string`

GetRconExecutionType returns the RconExecutionType field if non-nil, zero value otherwise.

### GetRconExecutionTypeOk

`func (o *ListedProduct) GetRconExecutionTypeOk() (*string, bool)`

GetRconExecutionTypeOk returns a tuple with the RconExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconExecutionType

`func (o *ListedProduct) SetRconExecutionType(v string)`

SetRconExecutionType sets RconExecutionType field to given value.

### HasRconExecutionType

`func (o *ListedProduct) HasRconExecutionType() bool`

HasRconExecutionType returns a boolean if a field has been set.

### GetRconStartTime

`func (o *ListedProduct) GetRconStartTime() time.Time`

GetRconStartTime returns the RconStartTime field if non-nil, zero value otherwise.

### GetRconStartTimeOk

`func (o *ListedProduct) GetRconStartTimeOk() (*time.Time, bool)`

GetRconStartTimeOk returns a tuple with the RconStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconStartTime

`func (o *ListedProduct) SetRconStartTime(v time.Time)`

SetRconStartTime sets RconStartTime field to given value.

### HasRconStartTime

`func (o *ListedProduct) HasRconStartTime() bool`

HasRconStartTime returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *ListedProduct) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *ListedProduct) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *ListedProduct) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *ListedProduct) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *ListedProduct) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *ListedProduct) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *ListedProduct) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *ListedProduct) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *ListedProduct) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *ListedProduct) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *ListedProduct) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *ListedProduct) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetDeliveryText

`func (o *ListedProduct) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *ListedProduct) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *ListedProduct) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *ListedProduct) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *ListedProduct) GetDeliveryTime() int32`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *ListedProduct) GetDeliveryTimeOk() (*int32, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *ListedProduct) SetDeliveryTime(v int32)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *ListedProduct) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetServiceText

`func (o *ListedProduct) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *ListedProduct) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *ListedProduct) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *ListedProduct) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *ListedProduct) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *ListedProduct) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *ListedProduct) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *ListedProduct) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetStock

`func (o *ListedProduct) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *ListedProduct) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *ListedProduct) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *ListedProduct) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *ListedProduct) GetDynamicWebhook() map[string]interface{}`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *ListedProduct) GetDynamicWebhookOk() (*map[string]interface{}, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *ListedProduct) SetDynamicWebhook(v map[string]interface{})`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *ListedProduct) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetBestseller

`func (o *ListedProduct) GetBestseller() int32`

GetBestseller returns the Bestseller field if non-nil, zero value otherwise.

### GetBestsellerOk

`func (o *ListedProduct) GetBestsellerOk() (*int32, bool)`

GetBestsellerOk returns a tuple with the Bestseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestseller

`func (o *ListedProduct) SetBestseller(v int32)`

SetBestseller sets Bestseller field to given value.

### HasBestseller

`func (o *ListedProduct) HasBestseller() bool`

HasBestseller returns a boolean if a field has been set.

### GetSortPriority

`func (o *ListedProduct) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *ListedProduct) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *ListedProduct) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *ListedProduct) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetUnlisted

`func (o *ListedProduct) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *ListedProduct) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *ListedProduct) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *ListedProduct) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetOnHold

`func (o *ListedProduct) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *ListedProduct) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *ListedProduct) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *ListedProduct) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *ListedProduct) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *ListedProduct) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *ListedProduct) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *ListedProduct) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *ListedProduct) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *ListedProduct) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *ListedProduct) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *ListedProduct) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *ListedProduct) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *ListedProduct) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *ListedProduct) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *ListedProduct) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetWatermarkEnabled

`func (o *ListedProduct) GetWatermarkEnabled() float32`

GetWatermarkEnabled returns the WatermarkEnabled field if non-nil, zero value otherwise.

### GetWatermarkEnabledOk

`func (o *ListedProduct) GetWatermarkEnabledOk() (*float32, bool)`

GetWatermarkEnabledOk returns a tuple with the WatermarkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkEnabled

`func (o *ListedProduct) SetWatermarkEnabled(v float32)`

SetWatermarkEnabled sets WatermarkEnabled field to given value.

### HasWatermarkEnabled

`func (o *ListedProduct) HasWatermarkEnabled() bool`

HasWatermarkEnabled returns a boolean if a field has been set.

### GetWatermarkText

`func (o *ListedProduct) GetWatermarkText() string`

GetWatermarkText returns the WatermarkText field if non-nil, zero value otherwise.

### GetWatermarkTextOk

`func (o *ListedProduct) GetWatermarkTextOk() (*string, bool)`

GetWatermarkTextOk returns a tuple with the WatermarkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkText

`func (o *ListedProduct) SetWatermarkText(v string)`

SetWatermarkText sets WatermarkText field to given value.

### HasWatermarkText

`func (o *ListedProduct) HasWatermarkText() bool`

HasWatermarkText returns a boolean if a field has been set.

### GetRedirectLink

`func (o *ListedProduct) GetRedirectLink() string`

GetRedirectLink returns the RedirectLink field if non-nil, zero value otherwise.

### GetRedirectLinkOk

`func (o *ListedProduct) GetRedirectLinkOk() (*string, bool)`

GetRedirectLinkOk returns a tuple with the RedirectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectLink

`func (o *ListedProduct) SetRedirectLink(v string)`

SetRedirectLink sets RedirectLink field to given value.

### HasRedirectLink

`func (o *ListedProduct) HasRedirectLink() bool`

HasRedirectLink returns a boolean if a field has been set.

### GetLabelSingular

`func (o *ListedProduct) GetLabelSingular() string`

GetLabelSingular returns the LabelSingular field if non-nil, zero value otherwise.

### GetLabelSingularOk

`func (o *ListedProduct) GetLabelSingularOk() (*string, bool)`

GetLabelSingularOk returns a tuple with the LabelSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSingular

`func (o *ListedProduct) SetLabelSingular(v string)`

SetLabelSingular sets LabelSingular field to given value.

### HasLabelSingular

`func (o *ListedProduct) HasLabelSingular() bool`

HasLabelSingular returns a boolean if a field has been set.

### GetLabelPlural

`func (o *ListedProduct) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *ListedProduct) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *ListedProduct) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *ListedProduct) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetPrivate

`func (o *ListedProduct) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ListedProduct) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ListedProduct) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ListedProduct) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListedProduct) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListedProduct) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListedProduct) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListedProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListedProduct) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListedProduct) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListedProduct) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListedProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListedProduct) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListedProduct) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListedProduct) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListedProduct) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetMarketplaceCategoryId

`func (o *ListedProduct) GetMarketplaceCategoryId() string`

GetMarketplaceCategoryId returns the MarketplaceCategoryId field if non-nil, zero value otherwise.

### GetMarketplaceCategoryIdOk

`func (o *ListedProduct) GetMarketplaceCategoryIdOk() (*string, bool)`

GetMarketplaceCategoryIdOk returns a tuple with the MarketplaceCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceCategoryId

`func (o *ListedProduct) SetMarketplaceCategoryId(v string)`

SetMarketplaceCategoryId sets MarketplaceCategoryId field to given value.

### HasMarketplaceCategoryId

`func (o *ListedProduct) HasMarketplaceCategoryId() bool`

HasMarketplaceCategoryId returns a boolean if a field has been set.

### GetTelegramGroupId

`func (o *ListedProduct) GetTelegramGroupId() string`

GetTelegramGroupId returns the TelegramGroupId field if non-nil, zero value otherwise.

### GetTelegramGroupIdOk

`func (o *ListedProduct) GetTelegramGroupIdOk() (*string, bool)`

GetTelegramGroupIdOk returns a tuple with the TelegramGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramGroupId

`func (o *ListedProduct) SetTelegramGroupId(v string)`

SetTelegramGroupId sets TelegramGroupId field to given value.

### HasTelegramGroupId

`func (o *ListedProduct) HasTelegramGroupId() bool`

HasTelegramGroupId returns a boolean if a field has been set.

### GetTelegramIntegration

`func (o *ListedProduct) GetTelegramIntegration() int32`

GetTelegramIntegration returns the TelegramIntegration field if non-nil, zero value otherwise.

### GetTelegramIntegrationOk

`func (o *ListedProduct) GetTelegramIntegrationOk() (*int32, bool)`

GetTelegramIntegrationOk returns a tuple with the TelegramIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramIntegration

`func (o *ListedProduct) SetTelegramIntegration(v int32)`

SetTelegramIntegration sets TelegramIntegration field to given value.

### HasTelegramIntegration

`func (o *ListedProduct) HasTelegramIntegration() bool`

HasTelegramIntegration returns a boolean if a field has been set.

### GetTelegramOptional

`func (o *ListedProduct) GetTelegramOptional() int32`

GetTelegramOptional returns the TelegramOptional field if non-nil, zero value otherwise.

### GetTelegramOptionalOk

`func (o *ListedProduct) GetTelegramOptionalOk() (*int32, bool)`

GetTelegramOptionalOk returns a tuple with the TelegramOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramOptional

`func (o *ListedProduct) SetTelegramOptional(v int32)`

SetTelegramOptional sets TelegramOptional field to given value.

### HasTelegramOptional

`func (o *ListedProduct) HasTelegramOptional() bool`

HasTelegramOptional returns a boolean if a field has been set.

### GetImageName

`func (o *ListedProduct) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListedProduct) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListedProduct) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ListedProduct) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *ListedProduct) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *ListedProduct) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *ListedProduct) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *ListedProduct) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *ListedProduct) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *ListedProduct) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *ListedProduct) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *ListedProduct) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetImageAttachments

`func (o *ListedProduct) GetImageAttachments() []ImageAttachment`

GetImageAttachments returns the ImageAttachments field if non-nil, zero value otherwise.

### GetImageAttachmentsOk

`func (o *ListedProduct) GetImageAttachmentsOk() (*[]ImageAttachment, bool)`

GetImageAttachmentsOk returns a tuple with the ImageAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachments

`func (o *ListedProduct) SetImageAttachments(v []ImageAttachment)`

SetImageAttachments sets ImageAttachments field to given value.

### HasImageAttachments

`func (o *ListedProduct) HasImageAttachments() bool`

HasImageAttachments returns a boolean if a field has been set.

### GetFeedback

`func (o *ListedProduct) GetFeedback() ListedProductFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *ListedProduct) GetFeedbackOk() (*ListedProductFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *ListedProduct) SetFeedback(v ListedProductFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *ListedProduct) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetCategories

`func (o *ListedProduct) GetCategories() []FeedbackProductCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ListedProduct) GetCategoriesOk() (*[]FeedbackProductCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ListedProduct) SetCategories(v []FeedbackProductCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ListedProduct) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPaymentGatewaysFees

`func (o *ListedProduct) GetPaymentGatewaysFees() []ListedProductPaymentGatewaysFeesInner`

GetPaymentGatewaysFees returns the PaymentGatewaysFees field if non-nil, zero value otherwise.

### GetPaymentGatewaysFeesOk

`func (o *ListedProduct) GetPaymentGatewaysFeesOk() (*[]ListedProductPaymentGatewaysFeesInner, bool)`

GetPaymentGatewaysFeesOk returns a tuple with the PaymentGatewaysFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewaysFees

`func (o *ListedProduct) SetPaymentGatewaysFees(v []ListedProductPaymentGatewaysFeesInner)`

SetPaymentGatewaysFees sets PaymentGatewaysFees field to given value.

### HasPaymentGatewaysFees

`func (o *ListedProduct) HasPaymentGatewaysFees() bool`

HasPaymentGatewaysFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


