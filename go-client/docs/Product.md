# Product

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
**TrialPeriod** | Pointer to **int32** | Defines a trial period before billing the customer for product type SUBSCRIPTION. Will be null if product is not a subsription | [optional] 
**SetupCost** | Pointer to **string** | If the product is a subcription, this is the setup cost of the subscription | [optional] 
**PaypalProductId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway PAYPAL chosen, a PayPal product is automatically created on the connected account. | [optional] 
**PaypalPlanId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway PAYPAL chosen, a PayPal plan is automatically created on the connected account. | [optional] 
**StripePriceId** | Pointer to **string** | When a product SUBSCRIPTION is created and the gateway STRIPE chosen, a Stripe price is automatically created on the connected account. | [optional] 
**DiscordIntegration** | Pointer to **int32** | Whether or not the discord integeration is enabled for this product. | [optional] 
**DiscordOptional** | Pointer to **int32** | Whether or not the discord integration is optional. | [optional] 
**DiscordSetRole** | Pointer to **int32** | Whether to give users a role when added to the discord server. | [optional] 
**DiscordServerId** | Pointer to **string** | The id of the discord server the bot will add users to. | [optional] 
**DiscordRoleId** | Pointer to **float32** | The role to give users when added by the discord integration. | [optional] 
**DiscordRemoveRole** | Pointer to **int32** | Whether to remove a role from the user when added to the discord server. | [optional] 
**QuantityMin** | Pointer to **int32** | Minimum quantity purchasable of this product. | [optional] 
**QuantityMax** | Pointer to **int32** | Maximum quantity purchasable of this product. | [optional] 
**QuantityWarning** | Pointer to **int32** | At which product quantity should we send a webhook and email warning the merchant. | [optional] 
**Gateways** | Pointer to [**[]Gateway**](Gateway.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldsArrayInner**](CustomFieldsArrayInner.md) |  | [optional] 
**RconCommands** | Pointer to **string** | Stringified array of the commands to be executed on the RCON server | [optional] 
**RconExecutionType** | Pointer to **string** | The type of execution for the RCON commands. Only applies to RCON products | [optional] 
**RconStartTime** | Pointer to **time.Time** | The time for when the RCON product should start | [optional] 
**CryptoConfirmationsNeeded** | Pointer to **int32** | Minimum number of confirmations for a crypto payment to be accepted. | [optional] 
**MaxRiskLevel** | Pointer to **int32** | For PAYPAL and STRIPE, maximum risk level to accept payments in order to block fraud attempts. | [optional] 
**BlockVpnProxies** | Pointer to **bool** | Block VPN and Proxy purchases if the gateway is PAYPAL or STRIPE. | [optional] 
**DeliveryText** | Pointer to **string** | The text to be delivered to the customer after payment | [optional] 
**DeliveryTime** | Pointer to **int32** | The timestamp for when the invoice was delivered | [optional] 
**ServiceText** | Pointer to **string** | The text to be delivered to the customer after payment for a service | [optional] 
**StockDelimiter** | Pointer to **string** | How to delimit the stock if product type is SERIALS, for example with stock_delimiter \&quot;,\&quot; and serials value first,second; the stock would have two different serials \&quot;first\&quot; and \&quot;second\&quot;. | [optional] 
**Stock** | Pointer to **int32** | Stock of the current product, can be -1 for unlimited stock. | [optional] 
**DynamicWebhook** | Pointer to **map[string]interface{}** |  | [optional] 
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
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the product. | [optional] 
**UpdatedAt** | Pointer to **int32** | Date, available if the product has been edited. | [optional] 
**UpdatedBy** | Pointer to **int32** | User ID of the user who updated the product. | [optional] 
**MarketplaceCategoryId** | Pointer to **string** | The category ID the product is a part of | [optional] 
**TelegramGroupId** | Pointer to **string** | The Telegram group ID | [optional] 
**TelegramIntegration** | Pointer to **int32** | Whether or not the Telegram integration is enabled | [optional] 
**TelegramOptional** | Pointer to **int32** | Whether or not the Telegram integration is optional | [optional] 
**Name** | Pointer to **string** | The name of the product | [optional] 
**ImageName** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The ID of the image stored in cloudflare&#39;s CDN | [optional] 
**ImageAttachments** | Pointer to [**[]ImageAttachment**](ImageAttachment.md) |  | [optional] 
**Feedback** | Pointer to [**ListedProductFeedback**](ListedProductFeedback.md) |  | [optional] 
**Categories** | Pointer to [**[]FeedbackProductCategoriesInner**](FeedbackProductCategoriesInner.md) |  | [optional] 
**PaymentGatewaysFees** | Pointer to [**[]ListedProductPaymentGatewaysFeesInner**](ListedProductPaymentGatewaysFeesInner.md) |  | [optional] 
**PriceConversions** | Pointer to [**FeedbackProductPriceConversions**](FeedbackProductPriceConversions.md) |  | [optional] 
**Addons** | Pointer to [**[][]ProductAddonsInner**]([]ProductAddonsInner.md) |  | [optional] 
**CountryRegulations** | Pointer to **string** | The country ISO code that the business is located in | [optional] 
**AvailableStripeApm** | Pointer to [**ProductAvailableStripeApm**](ProductAvailableStripeApm.md) |  | [optional] 
**Serials** | Pointer to **[]string** |  | [optional] 
**Webhooks** | Pointer to **[]string** |  | [optional] 
**Theme** | Pointer to **string** | The current theme enabled on the Sellix storefront | [optional] 
**DarkMode** | Pointer to **int32** | Whether or not dark mode is enabled on the Sellix storefront | [optional] 
**UiStyleConfiguration** | Pointer to **bool** | Whether or not the UI style configuration is enabled on the Sellix storefront | [optional] 
**VatPercentage** | Pointer to **string** | The VAR percentage set for the store | [optional] 
**TaxDetails** | Pointer to [**ProductTaxDetails**](ProductTaxDetails.md) |  | [optional] 
**ImageAttachmentInfo** | Pointer to [**ImageAttachment**](ImageAttachment.md) |  | [optional] 
**AverageScore** | Pointer to **string** | The average rating of the product. | [optional] 
**SoldCount** | Pointer to **float32** | The amount of sales for the product. | [optional] 

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqid

`func (o *Product) GetUniqid() string`

GetUniqid returns the Uniqid field if non-nil, zero value otherwise.

### GetUniqidOk

`func (o *Product) GetUniqidOk() (*string, bool)`

GetUniqidOk returns a tuple with the Uniqid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqid

`func (o *Product) SetUniqid(v string)`

SetUniqid sets Uniqid field to given value.

### HasUniqid

`func (o *Product) HasUniqid() bool`

HasUniqid returns a boolean if a field has been set.

### GetSlug

`func (o *Product) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Product) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Product) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Product) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetShopId

`func (o *Product) GetShopId() float32`

GetShopId returns the ShopId field if non-nil, zero value otherwise.

### GetShopIdOk

`func (o *Product) GetShopIdOk() (*float32, bool)`

GetShopIdOk returns a tuple with the ShopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopId

`func (o *Product) SetShopId(v float32)`

SetShopId sets ShopId field to given value.

### HasShopId

`func (o *Product) HasShopId() bool`

HasShopId returns a boolean if a field has been set.

### GetType

`func (o *Product) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Product) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Product) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Product) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubtype

`func (o *Product) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Product) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Product) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Product) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTitle

`func (o *Product) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Product) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Product) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Product) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCurrency

`func (o *Product) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Product) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Product) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Product) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPayWhatYouWant

`func (o *Product) GetPayWhatYouWant() int32`

GetPayWhatYouWant returns the PayWhatYouWant field if non-nil, zero value otherwise.

### GetPayWhatYouWantOk

`func (o *Product) GetPayWhatYouWantOk() (*int32, bool)`

GetPayWhatYouWantOk returns a tuple with the PayWhatYouWant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayWhatYouWant

`func (o *Product) SetPayWhatYouWant(v int32)`

SetPayWhatYouWant sets PayWhatYouWant field to given value.

### HasPayWhatYouWant

`func (o *Product) HasPayWhatYouWant() bool`

HasPayWhatYouWant returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplay

`func (o *Product) GetPriceDisplay() float64`

GetPriceDisplay returns the PriceDisplay field if non-nil, zero value otherwise.

### GetPriceDisplayOk

`func (o *Product) GetPriceDisplayOk() (*float64, bool)`

GetPriceDisplayOk returns a tuple with the PriceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplay

`func (o *Product) SetPriceDisplay(v float64)`

SetPriceDisplay sets PriceDisplay field to given value.

### HasPriceDisplay

`func (o *Product) HasPriceDisplay() bool`

HasPriceDisplay returns a boolean if a field has been set.

### GetPriceDiscount

`func (o *Product) GetPriceDiscount() float64`

GetPriceDiscount returns the PriceDiscount field if non-nil, zero value otherwise.

### GetPriceDiscountOk

`func (o *Product) GetPriceDiscountOk() (*float64, bool)`

GetPriceDiscountOk returns a tuple with the PriceDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDiscount

`func (o *Product) SetPriceDiscount(v float64)`

SetPriceDiscount sets PriceDiscount field to given value.

### HasPriceDiscount

`func (o *Product) HasPriceDiscount() bool`

HasPriceDiscount returns a boolean if a field has been set.

### GetAffiliateRevenuePercent

`func (o *Product) GetAffiliateRevenuePercent() float32`

GetAffiliateRevenuePercent returns the AffiliateRevenuePercent field if non-nil, zero value otherwise.

### GetAffiliateRevenuePercentOk

`func (o *Product) GetAffiliateRevenuePercentOk() (*float32, bool)`

GetAffiliateRevenuePercentOk returns a tuple with the AffiliateRevenuePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenuePercent

`func (o *Product) SetAffiliateRevenuePercent(v float32)`

SetAffiliateRevenuePercent sets AffiliateRevenuePercent field to given value.

### HasAffiliateRevenuePercent

`func (o *Product) HasAffiliateRevenuePercent() bool`

HasAffiliateRevenuePercent returns a boolean if a field has been set.

### GetPriceVariants

`func (o *Product) GetPriceVariants() map[string]interface{}`

GetPriceVariants returns the PriceVariants field if non-nil, zero value otherwise.

### GetPriceVariantsOk

`func (o *Product) GetPriceVariantsOk() (*map[string]interface{}, bool)`

GetPriceVariantsOk returns a tuple with the PriceVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceVariants

`func (o *Product) SetPriceVariants(v map[string]interface{})`

SetPriceVariants sets PriceVariants field to given value.

### HasPriceVariants

`func (o *Product) HasPriceVariants() bool`

HasPriceVariants returns a boolean if a field has been set.

### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicensingEnabled

`func (o *Product) GetLicensingEnabled() int32`

GetLicensingEnabled returns the LicensingEnabled field if non-nil, zero value otherwise.

### GetLicensingEnabledOk

`func (o *Product) GetLicensingEnabledOk() (*int32, bool)`

GetLicensingEnabledOk returns a tuple with the LicensingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingEnabled

`func (o *Product) SetLicensingEnabled(v int32)`

SetLicensingEnabled sets LicensingEnabled field to given value.

### HasLicensingEnabled

`func (o *Product) HasLicensingEnabled() bool`

HasLicensingEnabled returns a boolean if a field has been set.

### GetLicensePeriod

`func (o *Product) GetLicensePeriod() int32`

GetLicensePeriod returns the LicensePeriod field if non-nil, zero value otherwise.

### GetLicensePeriodOk

`func (o *Product) GetLicensePeriodOk() (*int32, bool)`

GetLicensePeriodOk returns a tuple with the LicensePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePeriod

`func (o *Product) SetLicensePeriod(v int32)`

SetLicensePeriod sets LicensePeriod field to given value.

### HasLicensePeriod

`func (o *Product) HasLicensePeriod() bool`

HasLicensePeriod returns a boolean if a field has been set.

### GetImageAttachment

`func (o *Product) GetImageAttachment() ImageAttachment`

GetImageAttachment returns the ImageAttachment field if non-nil, zero value otherwise.

### GetImageAttachmentOk

`func (o *Product) GetImageAttachmentOk() (*ImageAttachment, bool)`

GetImageAttachmentOk returns a tuple with the ImageAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachment

`func (o *Product) SetImageAttachment(v ImageAttachment)`

SetImageAttachment sets ImageAttachment field to given value.

### HasImageAttachment

`func (o *Product) HasImageAttachment() bool`

HasImageAttachment returns a boolean if a field has been set.

### GetFileAttachment

`func (o *Product) GetFileAttachment() string`

GetFileAttachment returns the FileAttachment field if non-nil, zero value otherwise.

### GetFileAttachmentOk

`func (o *Product) GetFileAttachmentOk() (*string, bool)`

GetFileAttachmentOk returns a tuple with the FileAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAttachment

`func (o *Product) SetFileAttachment(v string)`

SetFileAttachment sets FileAttachment field to given value.

### HasFileAttachment

`func (o *Product) HasFileAttachment() bool`

HasFileAttachment returns a boolean if a field has been set.

### GetYoutubeLink

`func (o *Product) GetYoutubeLink() string`

GetYoutubeLink returns the YoutubeLink field if non-nil, zero value otherwise.

### GetYoutubeLinkOk

`func (o *Product) GetYoutubeLinkOk() (*string, bool)`

GetYoutubeLinkOk returns a tuple with the YoutubeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeLink

`func (o *Product) SetYoutubeLink(v string)`

SetYoutubeLink sets YoutubeLink field to given value.

### HasYoutubeLink

`func (o *Product) HasYoutubeLink() bool`

HasYoutubeLink returns a boolean if a field has been set.

### GetVolumeDiscounts

`func (o *Product) GetVolumeDiscounts() []ListedProductVolumeDiscountsInner`

GetVolumeDiscounts returns the VolumeDiscounts field if non-nil, zero value otherwise.

### GetVolumeDiscountsOk

`func (o *Product) GetVolumeDiscountsOk() (*[]ListedProductVolumeDiscountsInner, bool)`

GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDiscounts

`func (o *Product) SetVolumeDiscounts(v []ListedProductVolumeDiscountsInner)`

SetVolumeDiscounts sets VolumeDiscounts field to given value.

### HasVolumeDiscounts

`func (o *Product) HasVolumeDiscounts() bool`

HasVolumeDiscounts returns a boolean if a field has been set.

### GetRecurringInterval

`func (o *Product) GetRecurringInterval() string`

GetRecurringInterval returns the RecurringInterval field if non-nil, zero value otherwise.

### GetRecurringIntervalOk

`func (o *Product) GetRecurringIntervalOk() (*string, bool)`

GetRecurringIntervalOk returns a tuple with the RecurringInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInterval

`func (o *Product) SetRecurringInterval(v string)`

SetRecurringInterval sets RecurringInterval field to given value.

### HasRecurringInterval

`func (o *Product) HasRecurringInterval() bool`

HasRecurringInterval returns a boolean if a field has been set.

### GetRecurringIntervalCount

`func (o *Product) GetRecurringIntervalCount() int32`

GetRecurringIntervalCount returns the RecurringIntervalCount field if non-nil, zero value otherwise.

### GetRecurringIntervalCountOk

`func (o *Product) GetRecurringIntervalCountOk() (*int32, bool)`

GetRecurringIntervalCountOk returns a tuple with the RecurringIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringIntervalCount

`func (o *Product) SetRecurringIntervalCount(v int32)`

SetRecurringIntervalCount sets RecurringIntervalCount field to given value.

### HasRecurringIntervalCount

`func (o *Product) HasRecurringIntervalCount() bool`

HasRecurringIntervalCount returns a boolean if a field has been set.

### GetTrialPeriod

`func (o *Product) GetTrialPeriod() int32`

GetTrialPeriod returns the TrialPeriod field if non-nil, zero value otherwise.

### GetTrialPeriodOk

`func (o *Product) GetTrialPeriodOk() (*int32, bool)`

GetTrialPeriodOk returns a tuple with the TrialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriod

`func (o *Product) SetTrialPeriod(v int32)`

SetTrialPeriod sets TrialPeriod field to given value.

### HasTrialPeriod

`func (o *Product) HasTrialPeriod() bool`

HasTrialPeriod returns a boolean if a field has been set.

### GetSetupCost

`func (o *Product) GetSetupCost() string`

GetSetupCost returns the SetupCost field if non-nil, zero value otherwise.

### GetSetupCostOk

`func (o *Product) GetSetupCostOk() (*string, bool)`

GetSetupCostOk returns a tuple with the SetupCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCost

`func (o *Product) SetSetupCost(v string)`

SetSetupCost sets SetupCost field to given value.

### HasSetupCost

`func (o *Product) HasSetupCost() bool`

HasSetupCost returns a boolean if a field has been set.

### GetPaypalProductId

`func (o *Product) GetPaypalProductId() string`

GetPaypalProductId returns the PaypalProductId field if non-nil, zero value otherwise.

### GetPaypalProductIdOk

`func (o *Product) GetPaypalProductIdOk() (*string, bool)`

GetPaypalProductIdOk returns a tuple with the PaypalProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalProductId

`func (o *Product) SetPaypalProductId(v string)`

SetPaypalProductId sets PaypalProductId field to given value.

### HasPaypalProductId

`func (o *Product) HasPaypalProductId() bool`

HasPaypalProductId returns a boolean if a field has been set.

### GetPaypalPlanId

`func (o *Product) GetPaypalPlanId() string`

GetPaypalPlanId returns the PaypalPlanId field if non-nil, zero value otherwise.

### GetPaypalPlanIdOk

`func (o *Product) GetPaypalPlanIdOk() (*string, bool)`

GetPaypalPlanIdOk returns a tuple with the PaypalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPlanId

`func (o *Product) SetPaypalPlanId(v string)`

SetPaypalPlanId sets PaypalPlanId field to given value.

### HasPaypalPlanId

`func (o *Product) HasPaypalPlanId() bool`

HasPaypalPlanId returns a boolean if a field has been set.

### GetStripePriceId

`func (o *Product) GetStripePriceId() string`

GetStripePriceId returns the StripePriceId field if non-nil, zero value otherwise.

### GetStripePriceIdOk

`func (o *Product) GetStripePriceIdOk() (*string, bool)`

GetStripePriceIdOk returns a tuple with the StripePriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePriceId

`func (o *Product) SetStripePriceId(v string)`

SetStripePriceId sets StripePriceId field to given value.

### HasStripePriceId

`func (o *Product) HasStripePriceId() bool`

HasStripePriceId returns a boolean if a field has been set.

### GetDiscordIntegration

`func (o *Product) GetDiscordIntegration() int32`

GetDiscordIntegration returns the DiscordIntegration field if non-nil, zero value otherwise.

### GetDiscordIntegrationOk

`func (o *Product) GetDiscordIntegrationOk() (*int32, bool)`

GetDiscordIntegrationOk returns a tuple with the DiscordIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordIntegration

`func (o *Product) SetDiscordIntegration(v int32)`

SetDiscordIntegration sets DiscordIntegration field to given value.

### HasDiscordIntegration

`func (o *Product) HasDiscordIntegration() bool`

HasDiscordIntegration returns a boolean if a field has been set.

### GetDiscordOptional

`func (o *Product) GetDiscordOptional() int32`

GetDiscordOptional returns the DiscordOptional field if non-nil, zero value otherwise.

### GetDiscordOptionalOk

`func (o *Product) GetDiscordOptionalOk() (*int32, bool)`

GetDiscordOptionalOk returns a tuple with the DiscordOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordOptional

`func (o *Product) SetDiscordOptional(v int32)`

SetDiscordOptional sets DiscordOptional field to given value.

### HasDiscordOptional

`func (o *Product) HasDiscordOptional() bool`

HasDiscordOptional returns a boolean if a field has been set.

### GetDiscordSetRole

`func (o *Product) GetDiscordSetRole() int32`

GetDiscordSetRole returns the DiscordSetRole field if non-nil, zero value otherwise.

### GetDiscordSetRoleOk

`func (o *Product) GetDiscordSetRoleOk() (*int32, bool)`

GetDiscordSetRoleOk returns a tuple with the DiscordSetRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordSetRole

`func (o *Product) SetDiscordSetRole(v int32)`

SetDiscordSetRole sets DiscordSetRole field to given value.

### HasDiscordSetRole

`func (o *Product) HasDiscordSetRole() bool`

HasDiscordSetRole returns a boolean if a field has been set.

### GetDiscordServerId

`func (o *Product) GetDiscordServerId() string`

GetDiscordServerId returns the DiscordServerId field if non-nil, zero value otherwise.

### GetDiscordServerIdOk

`func (o *Product) GetDiscordServerIdOk() (*string, bool)`

GetDiscordServerIdOk returns a tuple with the DiscordServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordServerId

`func (o *Product) SetDiscordServerId(v string)`

SetDiscordServerId sets DiscordServerId field to given value.

### HasDiscordServerId

`func (o *Product) HasDiscordServerId() bool`

HasDiscordServerId returns a boolean if a field has been set.

### GetDiscordRoleId

`func (o *Product) GetDiscordRoleId() float32`

GetDiscordRoleId returns the DiscordRoleId field if non-nil, zero value otherwise.

### GetDiscordRoleIdOk

`func (o *Product) GetDiscordRoleIdOk() (*float32, bool)`

GetDiscordRoleIdOk returns a tuple with the DiscordRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRoleId

`func (o *Product) SetDiscordRoleId(v float32)`

SetDiscordRoleId sets DiscordRoleId field to given value.

### HasDiscordRoleId

`func (o *Product) HasDiscordRoleId() bool`

HasDiscordRoleId returns a boolean if a field has been set.

### GetDiscordRemoveRole

`func (o *Product) GetDiscordRemoveRole() int32`

GetDiscordRemoveRole returns the DiscordRemoveRole field if non-nil, zero value otherwise.

### GetDiscordRemoveRoleOk

`func (o *Product) GetDiscordRemoveRoleOk() (*int32, bool)`

GetDiscordRemoveRoleOk returns a tuple with the DiscordRemoveRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordRemoveRole

`func (o *Product) SetDiscordRemoveRole(v int32)`

SetDiscordRemoveRole sets DiscordRemoveRole field to given value.

### HasDiscordRemoveRole

`func (o *Product) HasDiscordRemoveRole() bool`

HasDiscordRemoveRole returns a boolean if a field has been set.

### GetQuantityMin

`func (o *Product) GetQuantityMin() int32`

GetQuantityMin returns the QuantityMin field if non-nil, zero value otherwise.

### GetQuantityMinOk

`func (o *Product) GetQuantityMinOk() (*int32, bool)`

GetQuantityMinOk returns a tuple with the QuantityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMin

`func (o *Product) SetQuantityMin(v int32)`

SetQuantityMin sets QuantityMin field to given value.

### HasQuantityMin

`func (o *Product) HasQuantityMin() bool`

HasQuantityMin returns a boolean if a field has been set.

### GetQuantityMax

`func (o *Product) GetQuantityMax() int32`

GetQuantityMax returns the QuantityMax field if non-nil, zero value otherwise.

### GetQuantityMaxOk

`func (o *Product) GetQuantityMaxOk() (*int32, bool)`

GetQuantityMaxOk returns a tuple with the QuantityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityMax

`func (o *Product) SetQuantityMax(v int32)`

SetQuantityMax sets QuantityMax field to given value.

### HasQuantityMax

`func (o *Product) HasQuantityMax() bool`

HasQuantityMax returns a boolean if a field has been set.

### GetQuantityWarning

`func (o *Product) GetQuantityWarning() int32`

GetQuantityWarning returns the QuantityWarning field if non-nil, zero value otherwise.

### GetQuantityWarningOk

`func (o *Product) GetQuantityWarningOk() (*int32, bool)`

GetQuantityWarningOk returns a tuple with the QuantityWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWarning

`func (o *Product) SetQuantityWarning(v int32)`

SetQuantityWarning sets QuantityWarning field to given value.

### HasQuantityWarning

`func (o *Product) HasQuantityWarning() bool`

HasQuantityWarning returns a boolean if a field has been set.

### GetGateways

`func (o *Product) GetGateways() []Gateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *Product) GetGatewaysOk() (*[]Gateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *Product) SetGateways(v []Gateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *Product) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetCustomFields

`func (o *Product) GetCustomFields() []CustomFieldsArrayInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Product) GetCustomFieldsOk() (*[]CustomFieldsArrayInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Product) SetCustomFields(v []CustomFieldsArrayInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Product) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRconCommands

`func (o *Product) GetRconCommands() string`

GetRconCommands returns the RconCommands field if non-nil, zero value otherwise.

### GetRconCommandsOk

`func (o *Product) GetRconCommandsOk() (*string, bool)`

GetRconCommandsOk returns a tuple with the RconCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconCommands

`func (o *Product) SetRconCommands(v string)`

SetRconCommands sets RconCommands field to given value.

### HasRconCommands

`func (o *Product) HasRconCommands() bool`

HasRconCommands returns a boolean if a field has been set.

### GetRconExecutionType

`func (o *Product) GetRconExecutionType() string`

GetRconExecutionType returns the RconExecutionType field if non-nil, zero value otherwise.

### GetRconExecutionTypeOk

`func (o *Product) GetRconExecutionTypeOk() (*string, bool)`

GetRconExecutionTypeOk returns a tuple with the RconExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconExecutionType

`func (o *Product) SetRconExecutionType(v string)`

SetRconExecutionType sets RconExecutionType field to given value.

### HasRconExecutionType

`func (o *Product) HasRconExecutionType() bool`

HasRconExecutionType returns a boolean if a field has been set.

### GetRconStartTime

`func (o *Product) GetRconStartTime() time.Time`

GetRconStartTime returns the RconStartTime field if non-nil, zero value otherwise.

### GetRconStartTimeOk

`func (o *Product) GetRconStartTimeOk() (*time.Time, bool)`

GetRconStartTimeOk returns a tuple with the RconStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRconStartTime

`func (o *Product) SetRconStartTime(v time.Time)`

SetRconStartTime sets RconStartTime field to given value.

### HasRconStartTime

`func (o *Product) HasRconStartTime() bool`

HasRconStartTime returns a boolean if a field has been set.

### GetCryptoConfirmationsNeeded

`func (o *Product) GetCryptoConfirmationsNeeded() int32`

GetCryptoConfirmationsNeeded returns the CryptoConfirmationsNeeded field if non-nil, zero value otherwise.

### GetCryptoConfirmationsNeededOk

`func (o *Product) GetCryptoConfirmationsNeededOk() (*int32, bool)`

GetCryptoConfirmationsNeededOk returns a tuple with the CryptoConfirmationsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoConfirmationsNeeded

`func (o *Product) SetCryptoConfirmationsNeeded(v int32)`

SetCryptoConfirmationsNeeded sets CryptoConfirmationsNeeded field to given value.

### HasCryptoConfirmationsNeeded

`func (o *Product) HasCryptoConfirmationsNeeded() bool`

HasCryptoConfirmationsNeeded returns a boolean if a field has been set.

### GetMaxRiskLevel

`func (o *Product) GetMaxRiskLevel() int32`

GetMaxRiskLevel returns the MaxRiskLevel field if non-nil, zero value otherwise.

### GetMaxRiskLevelOk

`func (o *Product) GetMaxRiskLevelOk() (*int32, bool)`

GetMaxRiskLevelOk returns a tuple with the MaxRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRiskLevel

`func (o *Product) SetMaxRiskLevel(v int32)`

SetMaxRiskLevel sets MaxRiskLevel field to given value.

### HasMaxRiskLevel

`func (o *Product) HasMaxRiskLevel() bool`

HasMaxRiskLevel returns a boolean if a field has been set.

### GetBlockVpnProxies

`func (o *Product) GetBlockVpnProxies() bool`

GetBlockVpnProxies returns the BlockVpnProxies field if non-nil, zero value otherwise.

### GetBlockVpnProxiesOk

`func (o *Product) GetBlockVpnProxiesOk() (*bool, bool)`

GetBlockVpnProxiesOk returns a tuple with the BlockVpnProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVpnProxies

`func (o *Product) SetBlockVpnProxies(v bool)`

SetBlockVpnProxies sets BlockVpnProxies field to given value.

### HasBlockVpnProxies

`func (o *Product) HasBlockVpnProxies() bool`

HasBlockVpnProxies returns a boolean if a field has been set.

### GetDeliveryText

`func (o *Product) GetDeliveryText() string`

GetDeliveryText returns the DeliveryText field if non-nil, zero value otherwise.

### GetDeliveryTextOk

`func (o *Product) GetDeliveryTextOk() (*string, bool)`

GetDeliveryTextOk returns a tuple with the DeliveryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryText

`func (o *Product) SetDeliveryText(v string)`

SetDeliveryText sets DeliveryText field to given value.

### HasDeliveryText

`func (o *Product) HasDeliveryText() bool`

HasDeliveryText returns a boolean if a field has been set.

### GetDeliveryTime

`func (o *Product) GetDeliveryTime() int32`

GetDeliveryTime returns the DeliveryTime field if non-nil, zero value otherwise.

### GetDeliveryTimeOk

`func (o *Product) GetDeliveryTimeOk() (*int32, bool)`

GetDeliveryTimeOk returns a tuple with the DeliveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTime

`func (o *Product) SetDeliveryTime(v int32)`

SetDeliveryTime sets DeliveryTime field to given value.

### HasDeliveryTime

`func (o *Product) HasDeliveryTime() bool`

HasDeliveryTime returns a boolean if a field has been set.

### GetServiceText

`func (o *Product) GetServiceText() string`

GetServiceText returns the ServiceText field if non-nil, zero value otherwise.

### GetServiceTextOk

`func (o *Product) GetServiceTextOk() (*string, bool)`

GetServiceTextOk returns a tuple with the ServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceText

`func (o *Product) SetServiceText(v string)`

SetServiceText sets ServiceText field to given value.

### HasServiceText

`func (o *Product) HasServiceText() bool`

HasServiceText returns a boolean if a field has been set.

### GetStockDelimiter

`func (o *Product) GetStockDelimiter() string`

GetStockDelimiter returns the StockDelimiter field if non-nil, zero value otherwise.

### GetStockDelimiterOk

`func (o *Product) GetStockDelimiterOk() (*string, bool)`

GetStockDelimiterOk returns a tuple with the StockDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDelimiter

`func (o *Product) SetStockDelimiter(v string)`

SetStockDelimiter sets StockDelimiter field to given value.

### HasStockDelimiter

`func (o *Product) HasStockDelimiter() bool`

HasStockDelimiter returns a boolean if a field has been set.

### GetStock

`func (o *Product) GetStock() int32`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Product) GetStockOk() (*int32, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Product) SetStock(v int32)`

SetStock sets Stock field to given value.

### HasStock

`func (o *Product) HasStock() bool`

HasStock returns a boolean if a field has been set.

### GetDynamicWebhook

`func (o *Product) GetDynamicWebhook() map[string]interface{}`

GetDynamicWebhook returns the DynamicWebhook field if non-nil, zero value otherwise.

### GetDynamicWebhookOk

`func (o *Product) GetDynamicWebhookOk() (*map[string]interface{}, bool)`

GetDynamicWebhookOk returns a tuple with the DynamicWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWebhook

`func (o *Product) SetDynamicWebhook(v map[string]interface{})`

SetDynamicWebhook sets DynamicWebhook field to given value.

### HasDynamicWebhook

`func (o *Product) HasDynamicWebhook() bool`

HasDynamicWebhook returns a boolean if a field has been set.

### GetBestseller

`func (o *Product) GetBestseller() float32`

GetBestseller returns the Bestseller field if non-nil, zero value otherwise.

### GetBestsellerOk

`func (o *Product) GetBestsellerOk() (*float32, bool)`

GetBestsellerOk returns a tuple with the Bestseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestseller

`func (o *Product) SetBestseller(v float32)`

SetBestseller sets Bestseller field to given value.

### HasBestseller

`func (o *Product) HasBestseller() bool`

HasBestseller returns a boolean if a field has been set.

### GetSortPriority

`func (o *Product) GetSortPriority() int32`

GetSortPriority returns the SortPriority field if non-nil, zero value otherwise.

### GetSortPriorityOk

`func (o *Product) GetSortPriorityOk() (*int32, bool)`

GetSortPriorityOk returns a tuple with the SortPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPriority

`func (o *Product) SetSortPriority(v int32)`

SetSortPriority sets SortPriority field to given value.

### HasSortPriority

`func (o *Product) HasSortPriority() bool`

HasSortPriority returns a boolean if a field has been set.

### GetUnlisted

`func (o *Product) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *Product) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *Product) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *Product) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetOnHold

`func (o *Product) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *Product) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *Product) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *Product) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *Product) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *Product) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *Product) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *Product) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetWarranty

`func (o *Product) GetWarranty() int32`

GetWarranty returns the Warranty field if non-nil, zero value otherwise.

### GetWarrantyOk

`func (o *Product) GetWarrantyOk() (*int32, bool)`

GetWarrantyOk returns a tuple with the Warranty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarranty

`func (o *Product) SetWarranty(v int32)`

SetWarranty sets Warranty field to given value.

### HasWarranty

`func (o *Product) HasWarranty() bool`

HasWarranty returns a boolean if a field has been set.

### GetWarrantyText

`func (o *Product) GetWarrantyText() string`

GetWarrantyText returns the WarrantyText field if non-nil, zero value otherwise.

### GetWarrantyTextOk

`func (o *Product) GetWarrantyTextOk() (*string, bool)`

GetWarrantyTextOk returns a tuple with the WarrantyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyText

`func (o *Product) SetWarrantyText(v string)`

SetWarrantyText sets WarrantyText field to given value.

### HasWarrantyText

`func (o *Product) HasWarrantyText() bool`

HasWarrantyText returns a boolean if a field has been set.

### GetWatermarkEnabled

`func (o *Product) GetWatermarkEnabled() float32`

GetWatermarkEnabled returns the WatermarkEnabled field if non-nil, zero value otherwise.

### GetWatermarkEnabledOk

`func (o *Product) GetWatermarkEnabledOk() (*float32, bool)`

GetWatermarkEnabledOk returns a tuple with the WatermarkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkEnabled

`func (o *Product) SetWatermarkEnabled(v float32)`

SetWatermarkEnabled sets WatermarkEnabled field to given value.

### HasWatermarkEnabled

`func (o *Product) HasWatermarkEnabled() bool`

HasWatermarkEnabled returns a boolean if a field has been set.

### GetWatermarkText

`func (o *Product) GetWatermarkText() string`

GetWatermarkText returns the WatermarkText field if non-nil, zero value otherwise.

### GetWatermarkTextOk

`func (o *Product) GetWatermarkTextOk() (*string, bool)`

GetWatermarkTextOk returns a tuple with the WatermarkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkText

`func (o *Product) SetWatermarkText(v string)`

SetWatermarkText sets WatermarkText field to given value.

### HasWatermarkText

`func (o *Product) HasWatermarkText() bool`

HasWatermarkText returns a boolean if a field has been set.

### GetRedirectLink

`func (o *Product) GetRedirectLink() string`

GetRedirectLink returns the RedirectLink field if non-nil, zero value otherwise.

### GetRedirectLinkOk

`func (o *Product) GetRedirectLinkOk() (*string, bool)`

GetRedirectLinkOk returns a tuple with the RedirectLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectLink

`func (o *Product) SetRedirectLink(v string)`

SetRedirectLink sets RedirectLink field to given value.

### HasRedirectLink

`func (o *Product) HasRedirectLink() bool`

HasRedirectLink returns a boolean if a field has been set.

### GetLabelSingular

`func (o *Product) GetLabelSingular() string`

GetLabelSingular returns the LabelSingular field if non-nil, zero value otherwise.

### GetLabelSingularOk

`func (o *Product) GetLabelSingularOk() (*string, bool)`

GetLabelSingularOk returns a tuple with the LabelSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSingular

`func (o *Product) SetLabelSingular(v string)`

SetLabelSingular sets LabelSingular field to given value.

### HasLabelSingular

`func (o *Product) HasLabelSingular() bool`

HasLabelSingular returns a boolean if a field has been set.

### GetLabelPlural

`func (o *Product) GetLabelPlural() string`

GetLabelPlural returns the LabelPlural field if non-nil, zero value otherwise.

### GetLabelPluralOk

`func (o *Product) GetLabelPluralOk() (*string, bool)`

GetLabelPluralOk returns a tuple with the LabelPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPlural

`func (o *Product) SetLabelPlural(v string)`

SetLabelPlural sets LabelPlural field to given value.

### HasLabelPlural

`func (o *Product) HasLabelPlural() bool`

HasLabelPlural returns a boolean if a field has been set.

### GetPrivate

`func (o *Product) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Product) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Product) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Product) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Product) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Product) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Product) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Product) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Product) GetUpdatedBy() int32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Product) GetUpdatedByOk() (*int32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Product) SetUpdatedBy(v int32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Product) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetMarketplaceCategoryId

`func (o *Product) GetMarketplaceCategoryId() string`

GetMarketplaceCategoryId returns the MarketplaceCategoryId field if non-nil, zero value otherwise.

### GetMarketplaceCategoryIdOk

`func (o *Product) GetMarketplaceCategoryIdOk() (*string, bool)`

GetMarketplaceCategoryIdOk returns a tuple with the MarketplaceCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceCategoryId

`func (o *Product) SetMarketplaceCategoryId(v string)`

SetMarketplaceCategoryId sets MarketplaceCategoryId field to given value.

### HasMarketplaceCategoryId

`func (o *Product) HasMarketplaceCategoryId() bool`

HasMarketplaceCategoryId returns a boolean if a field has been set.

### GetTelegramGroupId

`func (o *Product) GetTelegramGroupId() string`

GetTelegramGroupId returns the TelegramGroupId field if non-nil, zero value otherwise.

### GetTelegramGroupIdOk

`func (o *Product) GetTelegramGroupIdOk() (*string, bool)`

GetTelegramGroupIdOk returns a tuple with the TelegramGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramGroupId

`func (o *Product) SetTelegramGroupId(v string)`

SetTelegramGroupId sets TelegramGroupId field to given value.

### HasTelegramGroupId

`func (o *Product) HasTelegramGroupId() bool`

HasTelegramGroupId returns a boolean if a field has been set.

### GetTelegramIntegration

`func (o *Product) GetTelegramIntegration() int32`

GetTelegramIntegration returns the TelegramIntegration field if non-nil, zero value otherwise.

### GetTelegramIntegrationOk

`func (o *Product) GetTelegramIntegrationOk() (*int32, bool)`

GetTelegramIntegrationOk returns a tuple with the TelegramIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramIntegration

`func (o *Product) SetTelegramIntegration(v int32)`

SetTelegramIntegration sets TelegramIntegration field to given value.

### HasTelegramIntegration

`func (o *Product) HasTelegramIntegration() bool`

HasTelegramIntegration returns a boolean if a field has been set.

### GetTelegramOptional

`func (o *Product) GetTelegramOptional() int32`

GetTelegramOptional returns the TelegramOptional field if non-nil, zero value otherwise.

### GetTelegramOptionalOk

`func (o *Product) GetTelegramOptionalOk() (*int32, bool)`

GetTelegramOptionalOk returns a tuple with the TelegramOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramOptional

`func (o *Product) SetTelegramOptional(v int32)`

SetTelegramOptional sets TelegramOptional field to given value.

### HasTelegramOptional

`func (o *Product) HasTelegramOptional() bool`

HasTelegramOptional returns a boolean if a field has been set.

### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Product) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImageName

`func (o *Product) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *Product) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *Product) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *Product) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *Product) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *Product) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *Product) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *Product) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *Product) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *Product) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *Product) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *Product) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetImageAttachments

`func (o *Product) GetImageAttachments() []ImageAttachment`

GetImageAttachments returns the ImageAttachments field if non-nil, zero value otherwise.

### GetImageAttachmentsOk

`func (o *Product) GetImageAttachmentsOk() (*[]ImageAttachment, bool)`

GetImageAttachmentsOk returns a tuple with the ImageAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachments

`func (o *Product) SetImageAttachments(v []ImageAttachment)`

SetImageAttachments sets ImageAttachments field to given value.

### HasImageAttachments

`func (o *Product) HasImageAttachments() bool`

HasImageAttachments returns a boolean if a field has been set.

### GetFeedback

`func (o *Product) GetFeedback() ListedProductFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *Product) GetFeedbackOk() (*ListedProductFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *Product) SetFeedback(v ListedProductFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *Product) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetCategories

`func (o *Product) GetCategories() []FeedbackProductCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Product) GetCategoriesOk() (*[]FeedbackProductCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Product) SetCategories(v []FeedbackProductCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Product) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPaymentGatewaysFees

`func (o *Product) GetPaymentGatewaysFees() []ListedProductPaymentGatewaysFeesInner`

GetPaymentGatewaysFees returns the PaymentGatewaysFees field if non-nil, zero value otherwise.

### GetPaymentGatewaysFeesOk

`func (o *Product) GetPaymentGatewaysFeesOk() (*[]ListedProductPaymentGatewaysFeesInner, bool)`

GetPaymentGatewaysFeesOk returns a tuple with the PaymentGatewaysFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewaysFees

`func (o *Product) SetPaymentGatewaysFees(v []ListedProductPaymentGatewaysFeesInner)`

SetPaymentGatewaysFees sets PaymentGatewaysFees field to given value.

### HasPaymentGatewaysFees

`func (o *Product) HasPaymentGatewaysFees() bool`

HasPaymentGatewaysFees returns a boolean if a field has been set.

### GetPriceConversions

`func (o *Product) GetPriceConversions() FeedbackProductPriceConversions`

GetPriceConversions returns the PriceConversions field if non-nil, zero value otherwise.

### GetPriceConversionsOk

`func (o *Product) GetPriceConversionsOk() (*FeedbackProductPriceConversions, bool)`

GetPriceConversionsOk returns a tuple with the PriceConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceConversions

`func (o *Product) SetPriceConversions(v FeedbackProductPriceConversions)`

SetPriceConversions sets PriceConversions field to given value.

### HasPriceConversions

`func (o *Product) HasPriceConversions() bool`

HasPriceConversions returns a boolean if a field has been set.

### GetAddons

`func (o *Product) GetAddons() [][]ProductAddonsInner`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *Product) GetAddonsOk() (*[][]ProductAddonsInner, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *Product) SetAddons(v [][]ProductAddonsInner)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *Product) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCountryRegulations

`func (o *Product) GetCountryRegulations() string`

GetCountryRegulations returns the CountryRegulations field if non-nil, zero value otherwise.

### GetCountryRegulationsOk

`func (o *Product) GetCountryRegulationsOk() (*string, bool)`

GetCountryRegulationsOk returns a tuple with the CountryRegulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRegulations

`func (o *Product) SetCountryRegulations(v string)`

SetCountryRegulations sets CountryRegulations field to given value.

### HasCountryRegulations

`func (o *Product) HasCountryRegulations() bool`

HasCountryRegulations returns a boolean if a field has been set.

### GetAvailableStripeApm

`func (o *Product) GetAvailableStripeApm() ProductAvailableStripeApm`

GetAvailableStripeApm returns the AvailableStripeApm field if non-nil, zero value otherwise.

### GetAvailableStripeApmOk

`func (o *Product) GetAvailableStripeApmOk() (*ProductAvailableStripeApm, bool)`

GetAvailableStripeApmOk returns a tuple with the AvailableStripeApm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableStripeApm

`func (o *Product) SetAvailableStripeApm(v ProductAvailableStripeApm)`

SetAvailableStripeApm sets AvailableStripeApm field to given value.

### HasAvailableStripeApm

`func (o *Product) HasAvailableStripeApm() bool`

HasAvailableStripeApm returns a boolean if a field has been set.

### GetSerials

`func (o *Product) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *Product) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *Product) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *Product) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetWebhooks

`func (o *Product) GetWebhooks() []string`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *Product) GetWebhooksOk() (*[]string, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *Product) SetWebhooks(v []string)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *Product) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### GetTheme

`func (o *Product) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Product) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Product) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Product) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetDarkMode

`func (o *Product) GetDarkMode() int32`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *Product) GetDarkModeOk() (*int32, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *Product) SetDarkMode(v int32)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *Product) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetUiStyleConfiguration

`func (o *Product) GetUiStyleConfiguration() bool`

GetUiStyleConfiguration returns the UiStyleConfiguration field if non-nil, zero value otherwise.

### GetUiStyleConfigurationOk

`func (o *Product) GetUiStyleConfigurationOk() (*bool, bool)`

GetUiStyleConfigurationOk returns a tuple with the UiStyleConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiStyleConfiguration

`func (o *Product) SetUiStyleConfiguration(v bool)`

SetUiStyleConfiguration sets UiStyleConfiguration field to given value.

### HasUiStyleConfiguration

`func (o *Product) HasUiStyleConfiguration() bool`

HasUiStyleConfiguration returns a boolean if a field has been set.

### GetVatPercentage

`func (o *Product) GetVatPercentage() string`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *Product) GetVatPercentageOk() (*string, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *Product) SetVatPercentage(v string)`

SetVatPercentage sets VatPercentage field to given value.

### HasVatPercentage

`func (o *Product) HasVatPercentage() bool`

HasVatPercentage returns a boolean if a field has been set.

### GetTaxDetails

`func (o *Product) GetTaxDetails() ProductTaxDetails`

GetTaxDetails returns the TaxDetails field if non-nil, zero value otherwise.

### GetTaxDetailsOk

`func (o *Product) GetTaxDetailsOk() (*ProductTaxDetails, bool)`

GetTaxDetailsOk returns a tuple with the TaxDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDetails

`func (o *Product) SetTaxDetails(v ProductTaxDetails)`

SetTaxDetails sets TaxDetails field to given value.

### HasTaxDetails

`func (o *Product) HasTaxDetails() bool`

HasTaxDetails returns a boolean if a field has been set.

### GetImageAttachmentInfo

`func (o *Product) GetImageAttachmentInfo() ImageAttachment`

GetImageAttachmentInfo returns the ImageAttachmentInfo field if non-nil, zero value otherwise.

### GetImageAttachmentInfoOk

`func (o *Product) GetImageAttachmentInfoOk() (*ImageAttachment, bool)`

GetImageAttachmentInfoOk returns a tuple with the ImageAttachmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAttachmentInfo

`func (o *Product) SetImageAttachmentInfo(v ImageAttachment)`

SetImageAttachmentInfo sets ImageAttachmentInfo field to given value.

### HasImageAttachmentInfo

`func (o *Product) HasImageAttachmentInfo() bool`

HasImageAttachmentInfo returns a boolean if a field has been set.

### GetAverageScore

`func (o *Product) GetAverageScore() string`

GetAverageScore returns the AverageScore field if non-nil, zero value otherwise.

### GetAverageScoreOk

`func (o *Product) GetAverageScoreOk() (*string, bool)`

GetAverageScoreOk returns a tuple with the AverageScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageScore

`func (o *Product) SetAverageScore(v string)`

SetAverageScore sets AverageScore field to given value.

### HasAverageScore

`func (o *Product) HasAverageScore() bool`

HasAverageScore returns a boolean if a field has been set.

### GetSoldCount

`func (o *Product) GetSoldCount() float32`

GetSoldCount returns the SoldCount field if non-nil, zero value otherwise.

### GetSoldCountOk

`func (o *Product) GetSoldCountOk() (*float32, bool)`

GetSoldCountOk returns a tuple with the SoldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldCount

`func (o *Product) SetSoldCount(v float32)`

SetSoldCount sets SoldCount field to given value.

### HasSoldCount

`func (o *Product) HasSoldCount() bool`

HasSoldCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


