# ShopInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the shop | [optional] 
**UserId** | Pointer to **int32** | The ID of the user who requested the shop info | [optional] 
**Type** | Pointer to **string** | The type of the shop | [optional] 
**Name** | Pointer to **string** | The name of the shop | [optional] 
**Avatar** | Pointer to **string** | The hash for the avatar | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**DefaultCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**AvailableCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**VatPercentage** | Pointer to **string** | The VAT percentage setup for the store | [optional] 
**TaxConfiguration** | Pointer to **string** | The tax configuration for the shop | [optional] 
**DashboardFeature** | Pointer to **string** | The organization mode set for the dashboard | [optional] 
**DisplayTaxOnStorefront** | Pointer to **int32** | The tax percentage will be shown on the product cards of your storefront, and not on the checkout page and invoice PDF only. | [optional] 
**DisplayTaxCustomFields** | Pointer to **int32** | Whether the tax is displayed in custom fields | [optional] 
**ValidationOnlyForCompanies** | Pointer to **int32** | Wheter or not VAT validation is only done for companies | [optional] 
**ValidateVatNumber** | Pointer to **int32** | If enabled, we will validate the Company VAT number to be sure it’s correct. | [optional] 
**PricesTaxInclusive** | Pointer to **int32** | Whether or not product pricing incldues the VAT | [optional] 
**NotifyTrialEnding** | Pointer to **int32** | Send an email to your customers when the trial period is ending. | [optional] 
**NotifyUpcomingRenewal** | Pointer to **int32** | Send an email to your customers days before an upcoming renewal. | [optional] 
**NotifySubscriptionRenewalFailure** | Pointer to **int32** | Send an email to your customers if a subscription renewal fails. | [optional] 
**OrderEmails** | Pointer to **int32** | Whether or not the shop owner receives emails for orders | [optional] 
**SubscriptionGracePeriod** | Pointer to **int32** | In days, how much time should we wait before cancelling a subscription if no payment is completed. | [optional] 
**PaypalCreditCard** | Pointer to **int32** | Whether or not the shop accepts credit cards via PayPal | [optional] 
**ForcePaypalEmailDelivery** | Pointer to **int32** | Whether or not the shop delivers products to the PayPal email when the gateway is PAYPAL | [optional] 
**PaypalMerchantId** | Pointer to **string** | The PayPal merchant ID of the PayPal account linked to the shop | [optional] 
**BinanceId** | Pointer to **string** | The Binance ID of the Binance account linked to the shop | [optional] 
**WalletconnectId** | Pointer to **string** | The WalletConnect ID of the WalletConnect account linked to the shop | [optional] 
**NonCustodialWallet** | Pointer to **int32** | Whether or not the shop uses the Sellix non-custodial crypto wallet | [optional] 
**DarkMode** | Pointer to **int32** | Whether or not the shop has dark mode enabled | [optional] 
**DiscordLink** | Pointer to **string** | The Discord server invite link for the shop | [optional] 
**TwitterLink** | Pointer to **string** | The link to the shop&#39;s Twitter account | [optional] 
**InstagramLink** | Pointer to **string** | The link to the shop&#39;s Instagram account | [optional] 
**FacebookLink** | Pointer to **string** | The link to the shop&#39;s Facebook account | [optional] 
**TelegramLink** | Pointer to **string** | The invite link to the shop&#39;s Telegram community | [optional] 
**YoutubeLink** | Pointer to **string** | The link to the shop&#39;s YouTube channel | [optional] 
**RedditLink** | Pointer to **string** | The link to the shop&#39;s SubReddit | [optional] 
**TiktokLink** | Pointer to **string** | The link to the shop&#39;s TikTok account | [optional] 
**SearchEnabled** | Pointer to **int32** | Whether or not the shop has dark mode enabled | [optional] 
**SortEnabled** | Pointer to **int32** | Whether or not the shop has product sorting enabled | [optional] 
**CartEnabled** | Pointer to **int32** | Whether or not the shop has the cart system enabled | [optional] 
**CartMaximumCheckout** | Pointer to **string** | Set the maximum amount, in your currency, for an order made through the Shopping Cart. | [optional] 
**HideOutOfStock** | Pointer to **int32** | Automatically hide your products when out of stock. | [optional] 
**GoogleAnalyticsTrackingId** | Pointer to **string** | The google analyticd tracking id the shop uses for analytics | [optional] 
**CrispWebsiteId** | Pointer to **string** | The crisp website id the shop uses for analytics | [optional] 
**CenterProductTitles** | Pointer to **int32** | Whether or not the shop&#39;s theme centers product titles | [optional] 
**CenterGroupTitles** | Pointer to **int32** | Whether or not the shop&#39;s theme centers group titles | [optional] 
**PopupMessage** | Pointer to **string** | This message will be shown to anyone who visits your site. Do not include your Terms of Service here. | [optional] 
**Verified** | Pointer to **int32** | Whether or not the shop is verified by Sellix | [optional] 
**OnHold** | Pointer to **int32** | Whether or not the shop is on hold | [optional] 
**TermsOfService** | Pointer to **string** | The terms of service for the shop in MDX | [optional] 
**PrimaryColorCustomTheme** | Pointer to **string** | The primary color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**SecondaryColorCustomTheme** | Pointer to **string** | The secondary color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**PrimaryFontColorCustomTheme** | Pointer to **string** | The primary font color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**SecondaryFontColorCustomTheme** | Pointer to **string** | The secondary font color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**CustomEmbed** | Pointer to **int32** | Whether or not the shop uses a custom embed theme. Value is null if no custom theme is used. | [optional] 
**CustomTheme** | Pointer to **int32** | Whether or not the shop uses a custom storefront theme. Value is null if no custom theme is used. | [optional] 
**FontCustomTheme** | Pointer to **string** | The font of the shop&#39;s custom storefront theme. Value is null if no custom theme is used. | [optional] 
**StyleCustomTheme** | Pointer to **string** | The style of the shop&#39;s custom storefront theme. Value is null if no custom theme is used. | [optional] 
**EmbedStyleCustomTheme** | Pointer to **string** | The style of the shop&#39;s custom embed theme. Value is null if no custom theme is used. | [optional] 
**IndexCustomTheme** | Pointer to **string** | The index of the shop&#39;s active custom storefront theme. Value is null if no custom theme is used. | [optional] 
**ProductCardCustomTheme** | Pointer to **string** | The product_card of the shop&#39;s active custom storefront theme. Value is null if no custom theme is used. | [optional] 
**BannerCustomTheme** | Pointer to **map[string]interface{}** | The storefront banner of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**FooterCustomTheme** | Pointer to **map[string]interface{}** | The storefront footer of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**BackgroundImageCustomTheme** | Pointer to **map[string]interface{}** | The storefront background image of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**LogoCustomTheme** | Pointer to **map[string]interface{}** | The storefront logo of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**HeaderCustomTheme** | Pointer to **map[string]interface{}** | The storefront header of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**CardsInRowCustomTheme** | Pointer to **int32** | The amount of cards to display in a row on the storefront | [optional] 
**CardsAlignCustomTheme** | Pointer to **map[string]interface{}** | Value is null if no custom theme is used. | [optional] 
**GroupCardCustomTheme** | Pointer to **map[string]interface{}** | Value is null if no custom theme is used. | [optional] 
**CardAnimationCustomTheme** | Pointer to **map[string]interface{}** | Value is null if no custom theme is used. | [optional] 
**LightFontColorCustomTheme** | Pointer to **string** | The light mode font color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**DarkFontColorCustomTheme** | Pointer to **string** | The dark mode font color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**LightColorCustomTheme** | Pointer to **string** | The light mode accent color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**DarkColorCustomTheme** | Pointer to **string** | The dark mode accent color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**BorderColorCustomTheme** | Pointer to **string** | The border color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**ButtonColorCustomTheme** | Pointer to **string** | The button color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**ThinColorCustomTheme** | Pointer to **string** | The thin font color of the shop&#39;s custom theme. Value is null if no custom theme is used. | [optional] 
**SortCustomTheme** | Pointer to **string** | The default sorting method for the storefront&#39;s custom theme | [optional] 
**HelpspaceClientId** | Pointer to **string** | The helpspace client id of the helpspace account linked to the shop | [optional] 
**HelpspaceToken** | Pointer to **string** | The helpspace token of the helpspace account linked to the shop | [optional] 
**DescriptionYoutubeOnly** | Pointer to **int32** | Show only youtube video for invoice/product description. | [optional] 
**DescriptionSkipDefaultImage** | Pointer to **int32** | Skip Default Image for Product Description. | [optional] 
**HideStockCounter** | Pointer to **int32** | If enabled, the number of how many products are in stock will be removed, we will only show &#39;In Stock&#39; or &#39;Out of Stock&#39;. | [optional] 
**ImageWidth** | Pointer to **int32** | The width of the storefront image | [optional] 
**ImageHeight** | Pointer to **int32** | The height of the storefront image | [optional] 
**DefaultSort** | Pointer to **string** | The default sorting method for the storefront | [optional] 
**DescriptionImage** | Pointer to **int32** | Whether or not the shop has a description image | [optional] 
**HideProductsSold** | Pointer to **int32** | Hide the products sold counter on your storefront, only your average feedback will be displayed. | [optional] 
**ServiceDateFrom** | Pointer to **string** | Choose whether to display your business starting date from the day of your first sale or the day you have signed up to Sellix. | [optional] 
**CardsType** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**SetupWizard** | Pointer to **int32** | Whether or not the shop has completed the setup wizzard | [optional] 
**SetupCryptocurrencies** | Pointer to **int32** | Whether or not the shop has setup cryptocurrencies | [optional] 
**NoticesBanner** | Pointer to **string** | The notices for the shop | [optional] 
**AffiliateRevenuePercent** | Pointer to **int32** | The percentage of each payment given to affiliates | [optional] 
**CreatedAt** | Pointer to **int32** | Timestamp for the creation of the subscription. | [optional] 
**ImageName** | Pointer to **string** | DEPRECATED: The name of the product image with the file type | [optional] 
**ImageStorage** | Pointer to **string** | Where the image is stored in Sellix&#39;s self-hosted CDN. | [optional] 
**CloudflareImageId** | Pointer to **string** | The cloudflare image ID of this product, replaces image_attachment and image_name. Format https://imagedelivery.net/95QNzrEeP7RU5l5WdbyrKw/&lt;cloudflare_image_id&gt;/&lt;variant_name&gt; where variant_name can be shopItem, avatar, icon, imageAvatarFeedback, public, productImageCart. | [optional] 
**MarketplaceVerified** | Pointer to **int32** | Whether or not the shop is a verified marketplace | [optional] 

## Methods

### NewShopInfo

`func NewShopInfo() *ShopInfo`

NewShopInfo instantiates a new ShopInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShopInfoWithDefaults

`func NewShopInfoWithDefaults() *ShopInfo`

NewShopInfoWithDefaults instantiates a new ShopInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShopInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShopInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShopInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShopInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *ShopInfo) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ShopInfo) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ShopInfo) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ShopInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetType

`func (o *ShopInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShopInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShopInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShopInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ShopInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShopInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShopInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShopInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatar

`func (o *ShopInfo) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ShopInfo) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ShopInfo) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ShopInfo) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCurrency

`func (o *ShopInfo) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ShopInfo) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ShopInfo) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ShopInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *ShopInfo) GetDefaultCurrency() Currency`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *ShopInfo) GetDefaultCurrencyOk() (*Currency, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *ShopInfo) SetDefaultCurrency(v Currency)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *ShopInfo) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetAvailableCurrency

`func (o *ShopInfo) GetAvailableCurrency() Currency`

GetAvailableCurrency returns the AvailableCurrency field if non-nil, zero value otherwise.

### GetAvailableCurrencyOk

`func (o *ShopInfo) GetAvailableCurrencyOk() (*Currency, bool)`

GetAvailableCurrencyOk returns a tuple with the AvailableCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCurrency

`func (o *ShopInfo) SetAvailableCurrency(v Currency)`

SetAvailableCurrency sets AvailableCurrency field to given value.

### HasAvailableCurrency

`func (o *ShopInfo) HasAvailableCurrency() bool`

HasAvailableCurrency returns a boolean if a field has been set.

### GetVatPercentage

`func (o *ShopInfo) GetVatPercentage() string`

GetVatPercentage returns the VatPercentage field if non-nil, zero value otherwise.

### GetVatPercentageOk

`func (o *ShopInfo) GetVatPercentageOk() (*string, bool)`

GetVatPercentageOk returns a tuple with the VatPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatPercentage

`func (o *ShopInfo) SetVatPercentage(v string)`

SetVatPercentage sets VatPercentage field to given value.

### HasVatPercentage

`func (o *ShopInfo) HasVatPercentage() bool`

HasVatPercentage returns a boolean if a field has been set.

### GetTaxConfiguration

`func (o *ShopInfo) GetTaxConfiguration() string`

GetTaxConfiguration returns the TaxConfiguration field if non-nil, zero value otherwise.

### GetTaxConfigurationOk

`func (o *ShopInfo) GetTaxConfigurationOk() (*string, bool)`

GetTaxConfigurationOk returns a tuple with the TaxConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxConfiguration

`func (o *ShopInfo) SetTaxConfiguration(v string)`

SetTaxConfiguration sets TaxConfiguration field to given value.

### HasTaxConfiguration

`func (o *ShopInfo) HasTaxConfiguration() bool`

HasTaxConfiguration returns a boolean if a field has been set.

### GetDashboardFeature

`func (o *ShopInfo) GetDashboardFeature() string`

GetDashboardFeature returns the DashboardFeature field if non-nil, zero value otherwise.

### GetDashboardFeatureOk

`func (o *ShopInfo) GetDashboardFeatureOk() (*string, bool)`

GetDashboardFeatureOk returns a tuple with the DashboardFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFeature

`func (o *ShopInfo) SetDashboardFeature(v string)`

SetDashboardFeature sets DashboardFeature field to given value.

### HasDashboardFeature

`func (o *ShopInfo) HasDashboardFeature() bool`

HasDashboardFeature returns a boolean if a field has been set.

### GetDisplayTaxOnStorefront

`func (o *ShopInfo) GetDisplayTaxOnStorefront() int32`

GetDisplayTaxOnStorefront returns the DisplayTaxOnStorefront field if non-nil, zero value otherwise.

### GetDisplayTaxOnStorefrontOk

`func (o *ShopInfo) GetDisplayTaxOnStorefrontOk() (*int32, bool)`

GetDisplayTaxOnStorefrontOk returns a tuple with the DisplayTaxOnStorefront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxOnStorefront

`func (o *ShopInfo) SetDisplayTaxOnStorefront(v int32)`

SetDisplayTaxOnStorefront sets DisplayTaxOnStorefront field to given value.

### HasDisplayTaxOnStorefront

`func (o *ShopInfo) HasDisplayTaxOnStorefront() bool`

HasDisplayTaxOnStorefront returns a boolean if a field has been set.

### GetDisplayTaxCustomFields

`func (o *ShopInfo) GetDisplayTaxCustomFields() int32`

GetDisplayTaxCustomFields returns the DisplayTaxCustomFields field if non-nil, zero value otherwise.

### GetDisplayTaxCustomFieldsOk

`func (o *ShopInfo) GetDisplayTaxCustomFieldsOk() (*int32, bool)`

GetDisplayTaxCustomFieldsOk returns a tuple with the DisplayTaxCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTaxCustomFields

`func (o *ShopInfo) SetDisplayTaxCustomFields(v int32)`

SetDisplayTaxCustomFields sets DisplayTaxCustomFields field to given value.

### HasDisplayTaxCustomFields

`func (o *ShopInfo) HasDisplayTaxCustomFields() bool`

HasDisplayTaxCustomFields returns a boolean if a field has been set.

### GetValidationOnlyForCompanies

`func (o *ShopInfo) GetValidationOnlyForCompanies() int32`

GetValidationOnlyForCompanies returns the ValidationOnlyForCompanies field if non-nil, zero value otherwise.

### GetValidationOnlyForCompaniesOk

`func (o *ShopInfo) GetValidationOnlyForCompaniesOk() (*int32, bool)`

GetValidationOnlyForCompaniesOk returns a tuple with the ValidationOnlyForCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOnlyForCompanies

`func (o *ShopInfo) SetValidationOnlyForCompanies(v int32)`

SetValidationOnlyForCompanies sets ValidationOnlyForCompanies field to given value.

### HasValidationOnlyForCompanies

`func (o *ShopInfo) HasValidationOnlyForCompanies() bool`

HasValidationOnlyForCompanies returns a boolean if a field has been set.

### GetValidateVatNumber

`func (o *ShopInfo) GetValidateVatNumber() int32`

GetValidateVatNumber returns the ValidateVatNumber field if non-nil, zero value otherwise.

### GetValidateVatNumberOk

`func (o *ShopInfo) GetValidateVatNumberOk() (*int32, bool)`

GetValidateVatNumberOk returns a tuple with the ValidateVatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateVatNumber

`func (o *ShopInfo) SetValidateVatNumber(v int32)`

SetValidateVatNumber sets ValidateVatNumber field to given value.

### HasValidateVatNumber

`func (o *ShopInfo) HasValidateVatNumber() bool`

HasValidateVatNumber returns a boolean if a field has been set.

### GetPricesTaxInclusive

`func (o *ShopInfo) GetPricesTaxInclusive() int32`

GetPricesTaxInclusive returns the PricesTaxInclusive field if non-nil, zero value otherwise.

### GetPricesTaxInclusiveOk

`func (o *ShopInfo) GetPricesTaxInclusiveOk() (*int32, bool)`

GetPricesTaxInclusiveOk returns a tuple with the PricesTaxInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesTaxInclusive

`func (o *ShopInfo) SetPricesTaxInclusive(v int32)`

SetPricesTaxInclusive sets PricesTaxInclusive field to given value.

### HasPricesTaxInclusive

`func (o *ShopInfo) HasPricesTaxInclusive() bool`

HasPricesTaxInclusive returns a boolean if a field has been set.

### GetNotifyTrialEnding

`func (o *ShopInfo) GetNotifyTrialEnding() int32`

GetNotifyTrialEnding returns the NotifyTrialEnding field if non-nil, zero value otherwise.

### GetNotifyTrialEndingOk

`func (o *ShopInfo) GetNotifyTrialEndingOk() (*int32, bool)`

GetNotifyTrialEndingOk returns a tuple with the NotifyTrialEnding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTrialEnding

`func (o *ShopInfo) SetNotifyTrialEnding(v int32)`

SetNotifyTrialEnding sets NotifyTrialEnding field to given value.

### HasNotifyTrialEnding

`func (o *ShopInfo) HasNotifyTrialEnding() bool`

HasNotifyTrialEnding returns a boolean if a field has been set.

### GetNotifyUpcomingRenewal

`func (o *ShopInfo) GetNotifyUpcomingRenewal() int32`

GetNotifyUpcomingRenewal returns the NotifyUpcomingRenewal field if non-nil, zero value otherwise.

### GetNotifyUpcomingRenewalOk

`func (o *ShopInfo) GetNotifyUpcomingRenewalOk() (*int32, bool)`

GetNotifyUpcomingRenewalOk returns a tuple with the NotifyUpcomingRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUpcomingRenewal

`func (o *ShopInfo) SetNotifyUpcomingRenewal(v int32)`

SetNotifyUpcomingRenewal sets NotifyUpcomingRenewal field to given value.

### HasNotifyUpcomingRenewal

`func (o *ShopInfo) HasNotifyUpcomingRenewal() bool`

HasNotifyUpcomingRenewal returns a boolean if a field has been set.

### GetNotifySubscriptionRenewalFailure

`func (o *ShopInfo) GetNotifySubscriptionRenewalFailure() int32`

GetNotifySubscriptionRenewalFailure returns the NotifySubscriptionRenewalFailure field if non-nil, zero value otherwise.

### GetNotifySubscriptionRenewalFailureOk

`func (o *ShopInfo) GetNotifySubscriptionRenewalFailureOk() (*int32, bool)`

GetNotifySubscriptionRenewalFailureOk returns a tuple with the NotifySubscriptionRenewalFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySubscriptionRenewalFailure

`func (o *ShopInfo) SetNotifySubscriptionRenewalFailure(v int32)`

SetNotifySubscriptionRenewalFailure sets NotifySubscriptionRenewalFailure field to given value.

### HasNotifySubscriptionRenewalFailure

`func (o *ShopInfo) HasNotifySubscriptionRenewalFailure() bool`

HasNotifySubscriptionRenewalFailure returns a boolean if a field has been set.

### GetOrderEmails

`func (o *ShopInfo) GetOrderEmails() int32`

GetOrderEmails returns the OrderEmails field if non-nil, zero value otherwise.

### GetOrderEmailsOk

`func (o *ShopInfo) GetOrderEmailsOk() (*int32, bool)`

GetOrderEmailsOk returns a tuple with the OrderEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderEmails

`func (o *ShopInfo) SetOrderEmails(v int32)`

SetOrderEmails sets OrderEmails field to given value.

### HasOrderEmails

`func (o *ShopInfo) HasOrderEmails() bool`

HasOrderEmails returns a boolean if a field has been set.

### GetSubscriptionGracePeriod

`func (o *ShopInfo) GetSubscriptionGracePeriod() int32`

GetSubscriptionGracePeriod returns the SubscriptionGracePeriod field if non-nil, zero value otherwise.

### GetSubscriptionGracePeriodOk

`func (o *ShopInfo) GetSubscriptionGracePeriodOk() (*int32, bool)`

GetSubscriptionGracePeriodOk returns a tuple with the SubscriptionGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionGracePeriod

`func (o *ShopInfo) SetSubscriptionGracePeriod(v int32)`

SetSubscriptionGracePeriod sets SubscriptionGracePeriod field to given value.

### HasSubscriptionGracePeriod

`func (o *ShopInfo) HasSubscriptionGracePeriod() bool`

HasSubscriptionGracePeriod returns a boolean if a field has been set.

### GetPaypalCreditCard

`func (o *ShopInfo) GetPaypalCreditCard() int32`

GetPaypalCreditCard returns the PaypalCreditCard field if non-nil, zero value otherwise.

### GetPaypalCreditCardOk

`func (o *ShopInfo) GetPaypalCreditCardOk() (*int32, bool)`

GetPaypalCreditCardOk returns a tuple with the PaypalCreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalCreditCard

`func (o *ShopInfo) SetPaypalCreditCard(v int32)`

SetPaypalCreditCard sets PaypalCreditCard field to given value.

### HasPaypalCreditCard

`func (o *ShopInfo) HasPaypalCreditCard() bool`

HasPaypalCreditCard returns a boolean if a field has been set.

### GetForcePaypalEmailDelivery

`func (o *ShopInfo) GetForcePaypalEmailDelivery() int32`

GetForcePaypalEmailDelivery returns the ForcePaypalEmailDelivery field if non-nil, zero value otherwise.

### GetForcePaypalEmailDeliveryOk

`func (o *ShopInfo) GetForcePaypalEmailDeliveryOk() (*int32, bool)`

GetForcePaypalEmailDeliveryOk returns a tuple with the ForcePaypalEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePaypalEmailDelivery

`func (o *ShopInfo) SetForcePaypalEmailDelivery(v int32)`

SetForcePaypalEmailDelivery sets ForcePaypalEmailDelivery field to given value.

### HasForcePaypalEmailDelivery

`func (o *ShopInfo) HasForcePaypalEmailDelivery() bool`

HasForcePaypalEmailDelivery returns a boolean if a field has been set.

### GetPaypalMerchantId

`func (o *ShopInfo) GetPaypalMerchantId() string`

GetPaypalMerchantId returns the PaypalMerchantId field if non-nil, zero value otherwise.

### GetPaypalMerchantIdOk

`func (o *ShopInfo) GetPaypalMerchantIdOk() (*string, bool)`

GetPaypalMerchantIdOk returns a tuple with the PaypalMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalMerchantId

`func (o *ShopInfo) SetPaypalMerchantId(v string)`

SetPaypalMerchantId sets PaypalMerchantId field to given value.

### HasPaypalMerchantId

`func (o *ShopInfo) HasPaypalMerchantId() bool`

HasPaypalMerchantId returns a boolean if a field has been set.

### GetBinanceId

`func (o *ShopInfo) GetBinanceId() string`

GetBinanceId returns the BinanceId field if non-nil, zero value otherwise.

### GetBinanceIdOk

`func (o *ShopInfo) GetBinanceIdOk() (*string, bool)`

GetBinanceIdOk returns a tuple with the BinanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinanceId

`func (o *ShopInfo) SetBinanceId(v string)`

SetBinanceId sets BinanceId field to given value.

### HasBinanceId

`func (o *ShopInfo) HasBinanceId() bool`

HasBinanceId returns a boolean if a field has been set.

### GetWalletconnectId

`func (o *ShopInfo) GetWalletconnectId() string`

GetWalletconnectId returns the WalletconnectId field if non-nil, zero value otherwise.

### GetWalletconnectIdOk

`func (o *ShopInfo) GetWalletconnectIdOk() (*string, bool)`

GetWalletconnectIdOk returns a tuple with the WalletconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletconnectId

`func (o *ShopInfo) SetWalletconnectId(v string)`

SetWalletconnectId sets WalletconnectId field to given value.

### HasWalletconnectId

`func (o *ShopInfo) HasWalletconnectId() bool`

HasWalletconnectId returns a boolean if a field has been set.

### GetNonCustodialWallet

`func (o *ShopInfo) GetNonCustodialWallet() int32`

GetNonCustodialWallet returns the NonCustodialWallet field if non-nil, zero value otherwise.

### GetNonCustodialWalletOk

`func (o *ShopInfo) GetNonCustodialWalletOk() (*int32, bool)`

GetNonCustodialWalletOk returns a tuple with the NonCustodialWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCustodialWallet

`func (o *ShopInfo) SetNonCustodialWallet(v int32)`

SetNonCustodialWallet sets NonCustodialWallet field to given value.

### HasNonCustodialWallet

`func (o *ShopInfo) HasNonCustodialWallet() bool`

HasNonCustodialWallet returns a boolean if a field has been set.

### GetDarkMode

`func (o *ShopInfo) GetDarkMode() int32`

GetDarkMode returns the DarkMode field if non-nil, zero value otherwise.

### GetDarkModeOk

`func (o *ShopInfo) GetDarkModeOk() (*int32, bool)`

GetDarkModeOk returns a tuple with the DarkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkMode

`func (o *ShopInfo) SetDarkMode(v int32)`

SetDarkMode sets DarkMode field to given value.

### HasDarkMode

`func (o *ShopInfo) HasDarkMode() bool`

HasDarkMode returns a boolean if a field has been set.

### GetDiscordLink

`func (o *ShopInfo) GetDiscordLink() string`

GetDiscordLink returns the DiscordLink field if non-nil, zero value otherwise.

### GetDiscordLinkOk

`func (o *ShopInfo) GetDiscordLinkOk() (*string, bool)`

GetDiscordLinkOk returns a tuple with the DiscordLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordLink

`func (o *ShopInfo) SetDiscordLink(v string)`

SetDiscordLink sets DiscordLink field to given value.

### HasDiscordLink

`func (o *ShopInfo) HasDiscordLink() bool`

HasDiscordLink returns a boolean if a field has been set.

### GetTwitterLink

`func (o *ShopInfo) GetTwitterLink() string`

GetTwitterLink returns the TwitterLink field if non-nil, zero value otherwise.

### GetTwitterLinkOk

`func (o *ShopInfo) GetTwitterLinkOk() (*string, bool)`

GetTwitterLinkOk returns a tuple with the TwitterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterLink

`func (o *ShopInfo) SetTwitterLink(v string)`

SetTwitterLink sets TwitterLink field to given value.

### HasTwitterLink

`func (o *ShopInfo) HasTwitterLink() bool`

HasTwitterLink returns a boolean if a field has been set.

### GetInstagramLink

`func (o *ShopInfo) GetInstagramLink() string`

GetInstagramLink returns the InstagramLink field if non-nil, zero value otherwise.

### GetInstagramLinkOk

`func (o *ShopInfo) GetInstagramLinkOk() (*string, bool)`

GetInstagramLinkOk returns a tuple with the InstagramLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramLink

`func (o *ShopInfo) SetInstagramLink(v string)`

SetInstagramLink sets InstagramLink field to given value.

### HasInstagramLink

`func (o *ShopInfo) HasInstagramLink() bool`

HasInstagramLink returns a boolean if a field has been set.

### GetFacebookLink

`func (o *ShopInfo) GetFacebookLink() string`

GetFacebookLink returns the FacebookLink field if non-nil, zero value otherwise.

### GetFacebookLinkOk

`func (o *ShopInfo) GetFacebookLinkOk() (*string, bool)`

GetFacebookLinkOk returns a tuple with the FacebookLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookLink

`func (o *ShopInfo) SetFacebookLink(v string)`

SetFacebookLink sets FacebookLink field to given value.

### HasFacebookLink

`func (o *ShopInfo) HasFacebookLink() bool`

HasFacebookLink returns a boolean if a field has been set.

### GetTelegramLink

`func (o *ShopInfo) GetTelegramLink() string`

GetTelegramLink returns the TelegramLink field if non-nil, zero value otherwise.

### GetTelegramLinkOk

`func (o *ShopInfo) GetTelegramLinkOk() (*string, bool)`

GetTelegramLinkOk returns a tuple with the TelegramLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelegramLink

`func (o *ShopInfo) SetTelegramLink(v string)`

SetTelegramLink sets TelegramLink field to given value.

### HasTelegramLink

`func (o *ShopInfo) HasTelegramLink() bool`

HasTelegramLink returns a boolean if a field has been set.

### GetYoutubeLink

`func (o *ShopInfo) GetYoutubeLink() string`

GetYoutubeLink returns the YoutubeLink field if non-nil, zero value otherwise.

### GetYoutubeLinkOk

`func (o *ShopInfo) GetYoutubeLinkOk() (*string, bool)`

GetYoutubeLinkOk returns a tuple with the YoutubeLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeLink

`func (o *ShopInfo) SetYoutubeLink(v string)`

SetYoutubeLink sets YoutubeLink field to given value.

### HasYoutubeLink

`func (o *ShopInfo) HasYoutubeLink() bool`

HasYoutubeLink returns a boolean if a field has been set.

### GetRedditLink

`func (o *ShopInfo) GetRedditLink() string`

GetRedditLink returns the RedditLink field if non-nil, zero value otherwise.

### GetRedditLinkOk

`func (o *ShopInfo) GetRedditLinkOk() (*string, bool)`

GetRedditLinkOk returns a tuple with the RedditLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditLink

`func (o *ShopInfo) SetRedditLink(v string)`

SetRedditLink sets RedditLink field to given value.

### HasRedditLink

`func (o *ShopInfo) HasRedditLink() bool`

HasRedditLink returns a boolean if a field has been set.

### GetTiktokLink

`func (o *ShopInfo) GetTiktokLink() string`

GetTiktokLink returns the TiktokLink field if non-nil, zero value otherwise.

### GetTiktokLinkOk

`func (o *ShopInfo) GetTiktokLinkOk() (*string, bool)`

GetTiktokLinkOk returns a tuple with the TiktokLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokLink

`func (o *ShopInfo) SetTiktokLink(v string)`

SetTiktokLink sets TiktokLink field to given value.

### HasTiktokLink

`func (o *ShopInfo) HasTiktokLink() bool`

HasTiktokLink returns a boolean if a field has been set.

### GetSearchEnabled

`func (o *ShopInfo) GetSearchEnabled() int32`

GetSearchEnabled returns the SearchEnabled field if non-nil, zero value otherwise.

### GetSearchEnabledOk

`func (o *ShopInfo) GetSearchEnabledOk() (*int32, bool)`

GetSearchEnabledOk returns a tuple with the SearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchEnabled

`func (o *ShopInfo) SetSearchEnabled(v int32)`

SetSearchEnabled sets SearchEnabled field to given value.

### HasSearchEnabled

`func (o *ShopInfo) HasSearchEnabled() bool`

HasSearchEnabled returns a boolean if a field has been set.

### GetSortEnabled

`func (o *ShopInfo) GetSortEnabled() int32`

GetSortEnabled returns the SortEnabled field if non-nil, zero value otherwise.

### GetSortEnabledOk

`func (o *ShopInfo) GetSortEnabledOk() (*int32, bool)`

GetSortEnabledOk returns a tuple with the SortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortEnabled

`func (o *ShopInfo) SetSortEnabled(v int32)`

SetSortEnabled sets SortEnabled field to given value.

### HasSortEnabled

`func (o *ShopInfo) HasSortEnabled() bool`

HasSortEnabled returns a boolean if a field has been set.

### GetCartEnabled

`func (o *ShopInfo) GetCartEnabled() int32`

GetCartEnabled returns the CartEnabled field if non-nil, zero value otherwise.

### GetCartEnabledOk

`func (o *ShopInfo) GetCartEnabledOk() (*int32, bool)`

GetCartEnabledOk returns a tuple with the CartEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartEnabled

`func (o *ShopInfo) SetCartEnabled(v int32)`

SetCartEnabled sets CartEnabled field to given value.

### HasCartEnabled

`func (o *ShopInfo) HasCartEnabled() bool`

HasCartEnabled returns a boolean if a field has been set.

### GetCartMaximumCheckout

`func (o *ShopInfo) GetCartMaximumCheckout() string`

GetCartMaximumCheckout returns the CartMaximumCheckout field if non-nil, zero value otherwise.

### GetCartMaximumCheckoutOk

`func (o *ShopInfo) GetCartMaximumCheckoutOk() (*string, bool)`

GetCartMaximumCheckoutOk returns a tuple with the CartMaximumCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartMaximumCheckout

`func (o *ShopInfo) SetCartMaximumCheckout(v string)`

SetCartMaximumCheckout sets CartMaximumCheckout field to given value.

### HasCartMaximumCheckout

`func (o *ShopInfo) HasCartMaximumCheckout() bool`

HasCartMaximumCheckout returns a boolean if a field has been set.

### GetHideOutOfStock

`func (o *ShopInfo) GetHideOutOfStock() int32`

GetHideOutOfStock returns the HideOutOfStock field if non-nil, zero value otherwise.

### GetHideOutOfStockOk

`func (o *ShopInfo) GetHideOutOfStockOk() (*int32, bool)`

GetHideOutOfStockOk returns a tuple with the HideOutOfStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOutOfStock

`func (o *ShopInfo) SetHideOutOfStock(v int32)`

SetHideOutOfStock sets HideOutOfStock field to given value.

### HasHideOutOfStock

`func (o *ShopInfo) HasHideOutOfStock() bool`

HasHideOutOfStock returns a boolean if a field has been set.

### GetGoogleAnalyticsTrackingId

`func (o *ShopInfo) GetGoogleAnalyticsTrackingId() string`

GetGoogleAnalyticsTrackingId returns the GoogleAnalyticsTrackingId field if non-nil, zero value otherwise.

### GetGoogleAnalyticsTrackingIdOk

`func (o *ShopInfo) GetGoogleAnalyticsTrackingIdOk() (*string, bool)`

GetGoogleAnalyticsTrackingIdOk returns a tuple with the GoogleAnalyticsTrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalyticsTrackingId

`func (o *ShopInfo) SetGoogleAnalyticsTrackingId(v string)`

SetGoogleAnalyticsTrackingId sets GoogleAnalyticsTrackingId field to given value.

### HasGoogleAnalyticsTrackingId

`func (o *ShopInfo) HasGoogleAnalyticsTrackingId() bool`

HasGoogleAnalyticsTrackingId returns a boolean if a field has been set.

### GetCrispWebsiteId

`func (o *ShopInfo) GetCrispWebsiteId() string`

GetCrispWebsiteId returns the CrispWebsiteId field if non-nil, zero value otherwise.

### GetCrispWebsiteIdOk

`func (o *ShopInfo) GetCrispWebsiteIdOk() (*string, bool)`

GetCrispWebsiteIdOk returns a tuple with the CrispWebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrispWebsiteId

`func (o *ShopInfo) SetCrispWebsiteId(v string)`

SetCrispWebsiteId sets CrispWebsiteId field to given value.

### HasCrispWebsiteId

`func (o *ShopInfo) HasCrispWebsiteId() bool`

HasCrispWebsiteId returns a boolean if a field has been set.

### GetCenterProductTitles

`func (o *ShopInfo) GetCenterProductTitles() int32`

GetCenterProductTitles returns the CenterProductTitles field if non-nil, zero value otherwise.

### GetCenterProductTitlesOk

`func (o *ShopInfo) GetCenterProductTitlesOk() (*int32, bool)`

GetCenterProductTitlesOk returns a tuple with the CenterProductTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterProductTitles

`func (o *ShopInfo) SetCenterProductTitles(v int32)`

SetCenterProductTitles sets CenterProductTitles field to given value.

### HasCenterProductTitles

`func (o *ShopInfo) HasCenterProductTitles() bool`

HasCenterProductTitles returns a boolean if a field has been set.

### GetCenterGroupTitles

`func (o *ShopInfo) GetCenterGroupTitles() int32`

GetCenterGroupTitles returns the CenterGroupTitles field if non-nil, zero value otherwise.

### GetCenterGroupTitlesOk

`func (o *ShopInfo) GetCenterGroupTitlesOk() (*int32, bool)`

GetCenterGroupTitlesOk returns a tuple with the CenterGroupTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCenterGroupTitles

`func (o *ShopInfo) SetCenterGroupTitles(v int32)`

SetCenterGroupTitles sets CenterGroupTitles field to given value.

### HasCenterGroupTitles

`func (o *ShopInfo) HasCenterGroupTitles() bool`

HasCenterGroupTitles returns a boolean if a field has been set.

### GetPopupMessage

`func (o *ShopInfo) GetPopupMessage() string`

GetPopupMessage returns the PopupMessage field if non-nil, zero value otherwise.

### GetPopupMessageOk

`func (o *ShopInfo) GetPopupMessageOk() (*string, bool)`

GetPopupMessageOk returns a tuple with the PopupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopupMessage

`func (o *ShopInfo) SetPopupMessage(v string)`

SetPopupMessage sets PopupMessage field to given value.

### HasPopupMessage

`func (o *ShopInfo) HasPopupMessage() bool`

HasPopupMessage returns a boolean if a field has been set.

### GetVerified

`func (o *ShopInfo) GetVerified() int32`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ShopInfo) GetVerifiedOk() (*int32, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ShopInfo) SetVerified(v int32)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ShopInfo) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetOnHold

`func (o *ShopInfo) GetOnHold() int32`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *ShopInfo) GetOnHoldOk() (*int32, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *ShopInfo) SetOnHold(v int32)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *ShopInfo) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetTermsOfService

`func (o *ShopInfo) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *ShopInfo) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *ShopInfo) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *ShopInfo) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetPrimaryColorCustomTheme

`func (o *ShopInfo) GetPrimaryColorCustomTheme() string`

GetPrimaryColorCustomTheme returns the PrimaryColorCustomTheme field if non-nil, zero value otherwise.

### GetPrimaryColorCustomThemeOk

`func (o *ShopInfo) GetPrimaryColorCustomThemeOk() (*string, bool)`

GetPrimaryColorCustomThemeOk returns a tuple with the PrimaryColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorCustomTheme

`func (o *ShopInfo) SetPrimaryColorCustomTheme(v string)`

SetPrimaryColorCustomTheme sets PrimaryColorCustomTheme field to given value.

### HasPrimaryColorCustomTheme

`func (o *ShopInfo) HasPrimaryColorCustomTheme() bool`

HasPrimaryColorCustomTheme returns a boolean if a field has been set.

### GetSecondaryColorCustomTheme

`func (o *ShopInfo) GetSecondaryColorCustomTheme() string`

GetSecondaryColorCustomTheme returns the SecondaryColorCustomTheme field if non-nil, zero value otherwise.

### GetSecondaryColorCustomThemeOk

`func (o *ShopInfo) GetSecondaryColorCustomThemeOk() (*string, bool)`

GetSecondaryColorCustomThemeOk returns a tuple with the SecondaryColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryColorCustomTheme

`func (o *ShopInfo) SetSecondaryColorCustomTheme(v string)`

SetSecondaryColorCustomTheme sets SecondaryColorCustomTheme field to given value.

### HasSecondaryColorCustomTheme

`func (o *ShopInfo) HasSecondaryColorCustomTheme() bool`

HasSecondaryColorCustomTheme returns a boolean if a field has been set.

### GetPrimaryFontColorCustomTheme

`func (o *ShopInfo) GetPrimaryFontColorCustomTheme() string`

GetPrimaryFontColorCustomTheme returns the PrimaryFontColorCustomTheme field if non-nil, zero value otherwise.

### GetPrimaryFontColorCustomThemeOk

`func (o *ShopInfo) GetPrimaryFontColorCustomThemeOk() (*string, bool)`

GetPrimaryFontColorCustomThemeOk returns a tuple with the PrimaryFontColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFontColorCustomTheme

`func (o *ShopInfo) SetPrimaryFontColorCustomTheme(v string)`

SetPrimaryFontColorCustomTheme sets PrimaryFontColorCustomTheme field to given value.

### HasPrimaryFontColorCustomTheme

`func (o *ShopInfo) HasPrimaryFontColorCustomTheme() bool`

HasPrimaryFontColorCustomTheme returns a boolean if a field has been set.

### GetSecondaryFontColorCustomTheme

`func (o *ShopInfo) GetSecondaryFontColorCustomTheme() string`

GetSecondaryFontColorCustomTheme returns the SecondaryFontColorCustomTheme field if non-nil, zero value otherwise.

### GetSecondaryFontColorCustomThemeOk

`func (o *ShopInfo) GetSecondaryFontColorCustomThemeOk() (*string, bool)`

GetSecondaryFontColorCustomThemeOk returns a tuple with the SecondaryFontColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFontColorCustomTheme

`func (o *ShopInfo) SetSecondaryFontColorCustomTheme(v string)`

SetSecondaryFontColorCustomTheme sets SecondaryFontColorCustomTheme field to given value.

### HasSecondaryFontColorCustomTheme

`func (o *ShopInfo) HasSecondaryFontColorCustomTheme() bool`

HasSecondaryFontColorCustomTheme returns a boolean if a field has been set.

### GetCustomEmbed

`func (o *ShopInfo) GetCustomEmbed() int32`

GetCustomEmbed returns the CustomEmbed field if non-nil, zero value otherwise.

### GetCustomEmbedOk

`func (o *ShopInfo) GetCustomEmbedOk() (*int32, bool)`

GetCustomEmbedOk returns a tuple with the CustomEmbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmbed

`func (o *ShopInfo) SetCustomEmbed(v int32)`

SetCustomEmbed sets CustomEmbed field to given value.

### HasCustomEmbed

`func (o *ShopInfo) HasCustomEmbed() bool`

HasCustomEmbed returns a boolean if a field has been set.

### GetCustomTheme

`func (o *ShopInfo) GetCustomTheme() int32`

GetCustomTheme returns the CustomTheme field if non-nil, zero value otherwise.

### GetCustomThemeOk

`func (o *ShopInfo) GetCustomThemeOk() (*int32, bool)`

GetCustomThemeOk returns a tuple with the CustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTheme

`func (o *ShopInfo) SetCustomTheme(v int32)`

SetCustomTheme sets CustomTheme field to given value.

### HasCustomTheme

`func (o *ShopInfo) HasCustomTheme() bool`

HasCustomTheme returns a boolean if a field has been set.

### GetFontCustomTheme

`func (o *ShopInfo) GetFontCustomTheme() string`

GetFontCustomTheme returns the FontCustomTheme field if non-nil, zero value otherwise.

### GetFontCustomThemeOk

`func (o *ShopInfo) GetFontCustomThemeOk() (*string, bool)`

GetFontCustomThemeOk returns a tuple with the FontCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontCustomTheme

`func (o *ShopInfo) SetFontCustomTheme(v string)`

SetFontCustomTheme sets FontCustomTheme field to given value.

### HasFontCustomTheme

`func (o *ShopInfo) HasFontCustomTheme() bool`

HasFontCustomTheme returns a boolean if a field has been set.

### GetStyleCustomTheme

`func (o *ShopInfo) GetStyleCustomTheme() string`

GetStyleCustomTheme returns the StyleCustomTheme field if non-nil, zero value otherwise.

### GetStyleCustomThemeOk

`func (o *ShopInfo) GetStyleCustomThemeOk() (*string, bool)`

GetStyleCustomThemeOk returns a tuple with the StyleCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleCustomTheme

`func (o *ShopInfo) SetStyleCustomTheme(v string)`

SetStyleCustomTheme sets StyleCustomTheme field to given value.

### HasStyleCustomTheme

`func (o *ShopInfo) HasStyleCustomTheme() bool`

HasStyleCustomTheme returns a boolean if a field has been set.

### GetEmbedStyleCustomTheme

`func (o *ShopInfo) GetEmbedStyleCustomTheme() string`

GetEmbedStyleCustomTheme returns the EmbedStyleCustomTheme field if non-nil, zero value otherwise.

### GetEmbedStyleCustomThemeOk

`func (o *ShopInfo) GetEmbedStyleCustomThemeOk() (*string, bool)`

GetEmbedStyleCustomThemeOk returns a tuple with the EmbedStyleCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedStyleCustomTheme

`func (o *ShopInfo) SetEmbedStyleCustomTheme(v string)`

SetEmbedStyleCustomTheme sets EmbedStyleCustomTheme field to given value.

### HasEmbedStyleCustomTheme

`func (o *ShopInfo) HasEmbedStyleCustomTheme() bool`

HasEmbedStyleCustomTheme returns a boolean if a field has been set.

### GetIndexCustomTheme

`func (o *ShopInfo) GetIndexCustomTheme() string`

GetIndexCustomTheme returns the IndexCustomTheme field if non-nil, zero value otherwise.

### GetIndexCustomThemeOk

`func (o *ShopInfo) GetIndexCustomThemeOk() (*string, bool)`

GetIndexCustomThemeOk returns a tuple with the IndexCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexCustomTheme

`func (o *ShopInfo) SetIndexCustomTheme(v string)`

SetIndexCustomTheme sets IndexCustomTheme field to given value.

### HasIndexCustomTheme

`func (o *ShopInfo) HasIndexCustomTheme() bool`

HasIndexCustomTheme returns a boolean if a field has been set.

### GetProductCardCustomTheme

`func (o *ShopInfo) GetProductCardCustomTheme() string`

GetProductCardCustomTheme returns the ProductCardCustomTheme field if non-nil, zero value otherwise.

### GetProductCardCustomThemeOk

`func (o *ShopInfo) GetProductCardCustomThemeOk() (*string, bool)`

GetProductCardCustomThemeOk returns a tuple with the ProductCardCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCardCustomTheme

`func (o *ShopInfo) SetProductCardCustomTheme(v string)`

SetProductCardCustomTheme sets ProductCardCustomTheme field to given value.

### HasProductCardCustomTheme

`func (o *ShopInfo) HasProductCardCustomTheme() bool`

HasProductCardCustomTheme returns a boolean if a field has been set.

### GetBannerCustomTheme

`func (o *ShopInfo) GetBannerCustomTheme() map[string]interface{}`

GetBannerCustomTheme returns the BannerCustomTheme field if non-nil, zero value otherwise.

### GetBannerCustomThemeOk

`func (o *ShopInfo) GetBannerCustomThemeOk() (*map[string]interface{}, bool)`

GetBannerCustomThemeOk returns a tuple with the BannerCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerCustomTheme

`func (o *ShopInfo) SetBannerCustomTheme(v map[string]interface{})`

SetBannerCustomTheme sets BannerCustomTheme field to given value.

### HasBannerCustomTheme

`func (o *ShopInfo) HasBannerCustomTheme() bool`

HasBannerCustomTheme returns a boolean if a field has been set.

### GetFooterCustomTheme

`func (o *ShopInfo) GetFooterCustomTheme() map[string]interface{}`

GetFooterCustomTheme returns the FooterCustomTheme field if non-nil, zero value otherwise.

### GetFooterCustomThemeOk

`func (o *ShopInfo) GetFooterCustomThemeOk() (*map[string]interface{}, bool)`

GetFooterCustomThemeOk returns a tuple with the FooterCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterCustomTheme

`func (o *ShopInfo) SetFooterCustomTheme(v map[string]interface{})`

SetFooterCustomTheme sets FooterCustomTheme field to given value.

### HasFooterCustomTheme

`func (o *ShopInfo) HasFooterCustomTheme() bool`

HasFooterCustomTheme returns a boolean if a field has been set.

### GetBackgroundImageCustomTheme

`func (o *ShopInfo) GetBackgroundImageCustomTheme() map[string]interface{}`

GetBackgroundImageCustomTheme returns the BackgroundImageCustomTheme field if non-nil, zero value otherwise.

### GetBackgroundImageCustomThemeOk

`func (o *ShopInfo) GetBackgroundImageCustomThemeOk() (*map[string]interface{}, bool)`

GetBackgroundImageCustomThemeOk returns a tuple with the BackgroundImageCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImageCustomTheme

`func (o *ShopInfo) SetBackgroundImageCustomTheme(v map[string]interface{})`

SetBackgroundImageCustomTheme sets BackgroundImageCustomTheme field to given value.

### HasBackgroundImageCustomTheme

`func (o *ShopInfo) HasBackgroundImageCustomTheme() bool`

HasBackgroundImageCustomTheme returns a boolean if a field has been set.

### GetLogoCustomTheme

`func (o *ShopInfo) GetLogoCustomTheme() map[string]interface{}`

GetLogoCustomTheme returns the LogoCustomTheme field if non-nil, zero value otherwise.

### GetLogoCustomThemeOk

`func (o *ShopInfo) GetLogoCustomThemeOk() (*map[string]interface{}, bool)`

GetLogoCustomThemeOk returns a tuple with the LogoCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoCustomTheme

`func (o *ShopInfo) SetLogoCustomTheme(v map[string]interface{})`

SetLogoCustomTheme sets LogoCustomTheme field to given value.

### HasLogoCustomTheme

`func (o *ShopInfo) HasLogoCustomTheme() bool`

HasLogoCustomTheme returns a boolean if a field has been set.

### GetHeaderCustomTheme

`func (o *ShopInfo) GetHeaderCustomTheme() map[string]interface{}`

GetHeaderCustomTheme returns the HeaderCustomTheme field if non-nil, zero value otherwise.

### GetHeaderCustomThemeOk

`func (o *ShopInfo) GetHeaderCustomThemeOk() (*map[string]interface{}, bool)`

GetHeaderCustomThemeOk returns a tuple with the HeaderCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderCustomTheme

`func (o *ShopInfo) SetHeaderCustomTheme(v map[string]interface{})`

SetHeaderCustomTheme sets HeaderCustomTheme field to given value.

### HasHeaderCustomTheme

`func (o *ShopInfo) HasHeaderCustomTheme() bool`

HasHeaderCustomTheme returns a boolean if a field has been set.

### GetCardsInRowCustomTheme

`func (o *ShopInfo) GetCardsInRowCustomTheme() int32`

GetCardsInRowCustomTheme returns the CardsInRowCustomTheme field if non-nil, zero value otherwise.

### GetCardsInRowCustomThemeOk

`func (o *ShopInfo) GetCardsInRowCustomThemeOk() (*int32, bool)`

GetCardsInRowCustomThemeOk returns a tuple with the CardsInRowCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsInRowCustomTheme

`func (o *ShopInfo) SetCardsInRowCustomTheme(v int32)`

SetCardsInRowCustomTheme sets CardsInRowCustomTheme field to given value.

### HasCardsInRowCustomTheme

`func (o *ShopInfo) HasCardsInRowCustomTheme() bool`

HasCardsInRowCustomTheme returns a boolean if a field has been set.

### GetCardsAlignCustomTheme

`func (o *ShopInfo) GetCardsAlignCustomTheme() map[string]interface{}`

GetCardsAlignCustomTheme returns the CardsAlignCustomTheme field if non-nil, zero value otherwise.

### GetCardsAlignCustomThemeOk

`func (o *ShopInfo) GetCardsAlignCustomThemeOk() (*map[string]interface{}, bool)`

GetCardsAlignCustomThemeOk returns a tuple with the CardsAlignCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsAlignCustomTheme

`func (o *ShopInfo) SetCardsAlignCustomTheme(v map[string]interface{})`

SetCardsAlignCustomTheme sets CardsAlignCustomTheme field to given value.

### HasCardsAlignCustomTheme

`func (o *ShopInfo) HasCardsAlignCustomTheme() bool`

HasCardsAlignCustomTheme returns a boolean if a field has been set.

### GetGroupCardCustomTheme

`func (o *ShopInfo) GetGroupCardCustomTheme() map[string]interface{}`

GetGroupCardCustomTheme returns the GroupCardCustomTheme field if non-nil, zero value otherwise.

### GetGroupCardCustomThemeOk

`func (o *ShopInfo) GetGroupCardCustomThemeOk() (*map[string]interface{}, bool)`

GetGroupCardCustomThemeOk returns a tuple with the GroupCardCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCardCustomTheme

`func (o *ShopInfo) SetGroupCardCustomTheme(v map[string]interface{})`

SetGroupCardCustomTheme sets GroupCardCustomTheme field to given value.

### HasGroupCardCustomTheme

`func (o *ShopInfo) HasGroupCardCustomTheme() bool`

HasGroupCardCustomTheme returns a boolean if a field has been set.

### GetCardAnimationCustomTheme

`func (o *ShopInfo) GetCardAnimationCustomTheme() map[string]interface{}`

GetCardAnimationCustomTheme returns the CardAnimationCustomTheme field if non-nil, zero value otherwise.

### GetCardAnimationCustomThemeOk

`func (o *ShopInfo) GetCardAnimationCustomThemeOk() (*map[string]interface{}, bool)`

GetCardAnimationCustomThemeOk returns a tuple with the CardAnimationCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAnimationCustomTheme

`func (o *ShopInfo) SetCardAnimationCustomTheme(v map[string]interface{})`

SetCardAnimationCustomTheme sets CardAnimationCustomTheme field to given value.

### HasCardAnimationCustomTheme

`func (o *ShopInfo) HasCardAnimationCustomTheme() bool`

HasCardAnimationCustomTheme returns a boolean if a field has been set.

### GetLightFontColorCustomTheme

`func (o *ShopInfo) GetLightFontColorCustomTheme() string`

GetLightFontColorCustomTheme returns the LightFontColorCustomTheme field if non-nil, zero value otherwise.

### GetLightFontColorCustomThemeOk

`func (o *ShopInfo) GetLightFontColorCustomThemeOk() (*string, bool)`

GetLightFontColorCustomThemeOk returns a tuple with the LightFontColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightFontColorCustomTheme

`func (o *ShopInfo) SetLightFontColorCustomTheme(v string)`

SetLightFontColorCustomTheme sets LightFontColorCustomTheme field to given value.

### HasLightFontColorCustomTheme

`func (o *ShopInfo) HasLightFontColorCustomTheme() bool`

HasLightFontColorCustomTheme returns a boolean if a field has been set.

### GetDarkFontColorCustomTheme

`func (o *ShopInfo) GetDarkFontColorCustomTheme() string`

GetDarkFontColorCustomTheme returns the DarkFontColorCustomTheme field if non-nil, zero value otherwise.

### GetDarkFontColorCustomThemeOk

`func (o *ShopInfo) GetDarkFontColorCustomThemeOk() (*string, bool)`

GetDarkFontColorCustomThemeOk returns a tuple with the DarkFontColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkFontColorCustomTheme

`func (o *ShopInfo) SetDarkFontColorCustomTheme(v string)`

SetDarkFontColorCustomTheme sets DarkFontColorCustomTheme field to given value.

### HasDarkFontColorCustomTheme

`func (o *ShopInfo) HasDarkFontColorCustomTheme() bool`

HasDarkFontColorCustomTheme returns a boolean if a field has been set.

### GetLightColorCustomTheme

`func (o *ShopInfo) GetLightColorCustomTheme() string`

GetLightColorCustomTheme returns the LightColorCustomTheme field if non-nil, zero value otherwise.

### GetLightColorCustomThemeOk

`func (o *ShopInfo) GetLightColorCustomThemeOk() (*string, bool)`

GetLightColorCustomThemeOk returns a tuple with the LightColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightColorCustomTheme

`func (o *ShopInfo) SetLightColorCustomTheme(v string)`

SetLightColorCustomTheme sets LightColorCustomTheme field to given value.

### HasLightColorCustomTheme

`func (o *ShopInfo) HasLightColorCustomTheme() bool`

HasLightColorCustomTheme returns a boolean if a field has been set.

### GetDarkColorCustomTheme

`func (o *ShopInfo) GetDarkColorCustomTheme() string`

GetDarkColorCustomTheme returns the DarkColorCustomTheme field if non-nil, zero value otherwise.

### GetDarkColorCustomThemeOk

`func (o *ShopInfo) GetDarkColorCustomThemeOk() (*string, bool)`

GetDarkColorCustomThemeOk returns a tuple with the DarkColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkColorCustomTheme

`func (o *ShopInfo) SetDarkColorCustomTheme(v string)`

SetDarkColorCustomTheme sets DarkColorCustomTheme field to given value.

### HasDarkColorCustomTheme

`func (o *ShopInfo) HasDarkColorCustomTheme() bool`

HasDarkColorCustomTheme returns a boolean if a field has been set.

### GetBorderColorCustomTheme

`func (o *ShopInfo) GetBorderColorCustomTheme() string`

GetBorderColorCustomTheme returns the BorderColorCustomTheme field if non-nil, zero value otherwise.

### GetBorderColorCustomThemeOk

`func (o *ShopInfo) GetBorderColorCustomThemeOk() (*string, bool)`

GetBorderColorCustomThemeOk returns a tuple with the BorderColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColorCustomTheme

`func (o *ShopInfo) SetBorderColorCustomTheme(v string)`

SetBorderColorCustomTheme sets BorderColorCustomTheme field to given value.

### HasBorderColorCustomTheme

`func (o *ShopInfo) HasBorderColorCustomTheme() bool`

HasBorderColorCustomTheme returns a boolean if a field has been set.

### GetButtonColorCustomTheme

`func (o *ShopInfo) GetButtonColorCustomTheme() string`

GetButtonColorCustomTheme returns the ButtonColorCustomTheme field if non-nil, zero value otherwise.

### GetButtonColorCustomThemeOk

`func (o *ShopInfo) GetButtonColorCustomThemeOk() (*string, bool)`

GetButtonColorCustomThemeOk returns a tuple with the ButtonColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonColorCustomTheme

`func (o *ShopInfo) SetButtonColorCustomTheme(v string)`

SetButtonColorCustomTheme sets ButtonColorCustomTheme field to given value.

### HasButtonColorCustomTheme

`func (o *ShopInfo) HasButtonColorCustomTheme() bool`

HasButtonColorCustomTheme returns a boolean if a field has been set.

### GetThinColorCustomTheme

`func (o *ShopInfo) GetThinColorCustomTheme() string`

GetThinColorCustomTheme returns the ThinColorCustomTheme field if non-nil, zero value otherwise.

### GetThinColorCustomThemeOk

`func (o *ShopInfo) GetThinColorCustomThemeOk() (*string, bool)`

GetThinColorCustomThemeOk returns a tuple with the ThinColorCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinColorCustomTheme

`func (o *ShopInfo) SetThinColorCustomTheme(v string)`

SetThinColorCustomTheme sets ThinColorCustomTheme field to given value.

### HasThinColorCustomTheme

`func (o *ShopInfo) HasThinColorCustomTheme() bool`

HasThinColorCustomTheme returns a boolean if a field has been set.

### GetSortCustomTheme

`func (o *ShopInfo) GetSortCustomTheme() string`

GetSortCustomTheme returns the SortCustomTheme field if non-nil, zero value otherwise.

### GetSortCustomThemeOk

`func (o *ShopInfo) GetSortCustomThemeOk() (*string, bool)`

GetSortCustomThemeOk returns a tuple with the SortCustomTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCustomTheme

`func (o *ShopInfo) SetSortCustomTheme(v string)`

SetSortCustomTheme sets SortCustomTheme field to given value.

### HasSortCustomTheme

`func (o *ShopInfo) HasSortCustomTheme() bool`

HasSortCustomTheme returns a boolean if a field has been set.

### GetHelpspaceClientId

`func (o *ShopInfo) GetHelpspaceClientId() string`

GetHelpspaceClientId returns the HelpspaceClientId field if non-nil, zero value otherwise.

### GetHelpspaceClientIdOk

`func (o *ShopInfo) GetHelpspaceClientIdOk() (*string, bool)`

GetHelpspaceClientIdOk returns a tuple with the HelpspaceClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpspaceClientId

`func (o *ShopInfo) SetHelpspaceClientId(v string)`

SetHelpspaceClientId sets HelpspaceClientId field to given value.

### HasHelpspaceClientId

`func (o *ShopInfo) HasHelpspaceClientId() bool`

HasHelpspaceClientId returns a boolean if a field has been set.

### GetHelpspaceToken

`func (o *ShopInfo) GetHelpspaceToken() string`

GetHelpspaceToken returns the HelpspaceToken field if non-nil, zero value otherwise.

### GetHelpspaceTokenOk

`func (o *ShopInfo) GetHelpspaceTokenOk() (*string, bool)`

GetHelpspaceTokenOk returns a tuple with the HelpspaceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpspaceToken

`func (o *ShopInfo) SetHelpspaceToken(v string)`

SetHelpspaceToken sets HelpspaceToken field to given value.

### HasHelpspaceToken

`func (o *ShopInfo) HasHelpspaceToken() bool`

HasHelpspaceToken returns a boolean if a field has been set.

### GetDescriptionYoutubeOnly

`func (o *ShopInfo) GetDescriptionYoutubeOnly() int32`

GetDescriptionYoutubeOnly returns the DescriptionYoutubeOnly field if non-nil, zero value otherwise.

### GetDescriptionYoutubeOnlyOk

`func (o *ShopInfo) GetDescriptionYoutubeOnlyOk() (*int32, bool)`

GetDescriptionYoutubeOnlyOk returns a tuple with the DescriptionYoutubeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionYoutubeOnly

`func (o *ShopInfo) SetDescriptionYoutubeOnly(v int32)`

SetDescriptionYoutubeOnly sets DescriptionYoutubeOnly field to given value.

### HasDescriptionYoutubeOnly

`func (o *ShopInfo) HasDescriptionYoutubeOnly() bool`

HasDescriptionYoutubeOnly returns a boolean if a field has been set.

### GetDescriptionSkipDefaultImage

`func (o *ShopInfo) GetDescriptionSkipDefaultImage() int32`

GetDescriptionSkipDefaultImage returns the DescriptionSkipDefaultImage field if non-nil, zero value otherwise.

### GetDescriptionSkipDefaultImageOk

`func (o *ShopInfo) GetDescriptionSkipDefaultImageOk() (*int32, bool)`

GetDescriptionSkipDefaultImageOk returns a tuple with the DescriptionSkipDefaultImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionSkipDefaultImage

`func (o *ShopInfo) SetDescriptionSkipDefaultImage(v int32)`

SetDescriptionSkipDefaultImage sets DescriptionSkipDefaultImage field to given value.

### HasDescriptionSkipDefaultImage

`func (o *ShopInfo) HasDescriptionSkipDefaultImage() bool`

HasDescriptionSkipDefaultImage returns a boolean if a field has been set.

### GetHideStockCounter

`func (o *ShopInfo) GetHideStockCounter() int32`

GetHideStockCounter returns the HideStockCounter field if non-nil, zero value otherwise.

### GetHideStockCounterOk

`func (o *ShopInfo) GetHideStockCounterOk() (*int32, bool)`

GetHideStockCounterOk returns a tuple with the HideStockCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideStockCounter

`func (o *ShopInfo) SetHideStockCounter(v int32)`

SetHideStockCounter sets HideStockCounter field to given value.

### HasHideStockCounter

`func (o *ShopInfo) HasHideStockCounter() bool`

HasHideStockCounter returns a boolean if a field has been set.

### GetImageWidth

`func (o *ShopInfo) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *ShopInfo) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *ShopInfo) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *ShopInfo) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.

### GetImageHeight

`func (o *ShopInfo) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *ShopInfo) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *ShopInfo) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *ShopInfo) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetDefaultSort

`func (o *ShopInfo) GetDefaultSort() string`

GetDefaultSort returns the DefaultSort field if non-nil, zero value otherwise.

### GetDefaultSortOk

`func (o *ShopInfo) GetDefaultSortOk() (*string, bool)`

GetDefaultSortOk returns a tuple with the DefaultSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSort

`func (o *ShopInfo) SetDefaultSort(v string)`

SetDefaultSort sets DefaultSort field to given value.

### HasDefaultSort

`func (o *ShopInfo) HasDefaultSort() bool`

HasDefaultSort returns a boolean if a field has been set.

### GetDescriptionImage

`func (o *ShopInfo) GetDescriptionImage() int32`

GetDescriptionImage returns the DescriptionImage field if non-nil, zero value otherwise.

### GetDescriptionImageOk

`func (o *ShopInfo) GetDescriptionImageOk() (*int32, bool)`

GetDescriptionImageOk returns a tuple with the DescriptionImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionImage

`func (o *ShopInfo) SetDescriptionImage(v int32)`

SetDescriptionImage sets DescriptionImage field to given value.

### HasDescriptionImage

`func (o *ShopInfo) HasDescriptionImage() bool`

HasDescriptionImage returns a boolean if a field has been set.

### GetHideProductsSold

`func (o *ShopInfo) GetHideProductsSold() int32`

GetHideProductsSold returns the HideProductsSold field if non-nil, zero value otherwise.

### GetHideProductsSoldOk

`func (o *ShopInfo) GetHideProductsSoldOk() (*int32, bool)`

GetHideProductsSoldOk returns a tuple with the HideProductsSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideProductsSold

`func (o *ShopInfo) SetHideProductsSold(v int32)`

SetHideProductsSold sets HideProductsSold field to given value.

### HasHideProductsSold

`func (o *ShopInfo) HasHideProductsSold() bool`

HasHideProductsSold returns a boolean if a field has been set.

### GetServiceDateFrom

`func (o *ShopInfo) GetServiceDateFrom() string`

GetServiceDateFrom returns the ServiceDateFrom field if non-nil, zero value otherwise.

### GetServiceDateFromOk

`func (o *ShopInfo) GetServiceDateFromOk() (*string, bool)`

GetServiceDateFromOk returns a tuple with the ServiceDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDateFrom

`func (o *ShopInfo) SetServiceDateFrom(v string)`

SetServiceDateFrom sets ServiceDateFrom field to given value.

### HasServiceDateFrom

`func (o *ShopInfo) HasServiceDateFrom() bool`

HasServiceDateFrom returns a boolean if a field has been set.

### GetCardsType

`func (o *ShopInfo) GetCardsType() string`

GetCardsType returns the CardsType field if non-nil, zero value otherwise.

### GetCardsTypeOk

`func (o *ShopInfo) GetCardsTypeOk() (*string, bool)`

GetCardsTypeOk returns a tuple with the CardsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsType

`func (o *ShopInfo) SetCardsType(v string)`

SetCardsType sets CardsType field to given value.

### HasCardsType

`func (o *ShopInfo) HasCardsType() bool`

HasCardsType returns a boolean if a field has been set.

### GetSetupWizard

`func (o *ShopInfo) GetSetupWizard() int32`

GetSetupWizard returns the SetupWizard field if non-nil, zero value otherwise.

### GetSetupWizardOk

`func (o *ShopInfo) GetSetupWizardOk() (*int32, bool)`

GetSetupWizardOk returns a tuple with the SetupWizard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupWizard

`func (o *ShopInfo) SetSetupWizard(v int32)`

SetSetupWizard sets SetupWizard field to given value.

### HasSetupWizard

`func (o *ShopInfo) HasSetupWizard() bool`

HasSetupWizard returns a boolean if a field has been set.

### GetSetupCryptocurrencies

`func (o *ShopInfo) GetSetupCryptocurrencies() int32`

GetSetupCryptocurrencies returns the SetupCryptocurrencies field if non-nil, zero value otherwise.

### GetSetupCryptocurrenciesOk

`func (o *ShopInfo) GetSetupCryptocurrenciesOk() (*int32, bool)`

GetSetupCryptocurrenciesOk returns a tuple with the SetupCryptocurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCryptocurrencies

`func (o *ShopInfo) SetSetupCryptocurrencies(v int32)`

SetSetupCryptocurrencies sets SetupCryptocurrencies field to given value.

### HasSetupCryptocurrencies

`func (o *ShopInfo) HasSetupCryptocurrencies() bool`

HasSetupCryptocurrencies returns a boolean if a field has been set.

### GetNoticesBanner

`func (o *ShopInfo) GetNoticesBanner() string`

GetNoticesBanner returns the NoticesBanner field if non-nil, zero value otherwise.

### GetNoticesBannerOk

`func (o *ShopInfo) GetNoticesBannerOk() (*string, bool)`

GetNoticesBannerOk returns a tuple with the NoticesBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticesBanner

`func (o *ShopInfo) SetNoticesBanner(v string)`

SetNoticesBanner sets NoticesBanner field to given value.

### HasNoticesBanner

`func (o *ShopInfo) HasNoticesBanner() bool`

HasNoticesBanner returns a boolean if a field has been set.

### GetAffiliateRevenuePercent

`func (o *ShopInfo) GetAffiliateRevenuePercent() int32`

GetAffiliateRevenuePercent returns the AffiliateRevenuePercent field if non-nil, zero value otherwise.

### GetAffiliateRevenuePercentOk

`func (o *ShopInfo) GetAffiliateRevenuePercentOk() (*int32, bool)`

GetAffiliateRevenuePercentOk returns a tuple with the AffiliateRevenuePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateRevenuePercent

`func (o *ShopInfo) SetAffiliateRevenuePercent(v int32)`

SetAffiliateRevenuePercent sets AffiliateRevenuePercent field to given value.

### HasAffiliateRevenuePercent

`func (o *ShopInfo) HasAffiliateRevenuePercent() bool`

HasAffiliateRevenuePercent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ShopInfo) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShopInfo) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShopInfo) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShopInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImageName

`func (o *ShopInfo) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ShopInfo) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ShopInfo) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ShopInfo) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageStorage

`func (o *ShopInfo) GetImageStorage() string`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *ShopInfo) GetImageStorageOk() (*string, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *ShopInfo) SetImageStorage(v string)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *ShopInfo) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetCloudflareImageId

`func (o *ShopInfo) GetCloudflareImageId() string`

GetCloudflareImageId returns the CloudflareImageId field if non-nil, zero value otherwise.

### GetCloudflareImageIdOk

`func (o *ShopInfo) GetCloudflareImageIdOk() (*string, bool)`

GetCloudflareImageIdOk returns a tuple with the CloudflareImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudflareImageId

`func (o *ShopInfo) SetCloudflareImageId(v string)`

SetCloudflareImageId sets CloudflareImageId field to given value.

### HasCloudflareImageId

`func (o *ShopInfo) HasCloudflareImageId() bool`

HasCloudflareImageId returns a boolean if a field has been set.

### GetMarketplaceVerified

`func (o *ShopInfo) GetMarketplaceVerified() int32`

GetMarketplaceVerified returns the MarketplaceVerified field if non-nil, zero value otherwise.

### GetMarketplaceVerifiedOk

`func (o *ShopInfo) GetMarketplaceVerifiedOk() (*int32, bool)`

GetMarketplaceVerifiedOk returns a tuple with the MarketplaceVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceVerified

`func (o *ShopInfo) SetMarketplaceVerified(v int32)`

SetMarketplaceVerified sets MarketplaceVerified field to given value.

### HasMarketplaceVerified

`func (o *ShopInfo) HasMarketplaceVerified() bool`

HasMarketplaceVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


