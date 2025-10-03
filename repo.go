package go_heyframe_admin_sdk

type Repository struct {
	ClientService

	AclRole *AclRoleRepository

	AclUserRole *AclUserRoleRepository

	Category *CategoryRepository

	CategoryTag *CategoryTagRepository

	CategoryTranslation *CategoryTranslationRepository

	Channel *ChannelRepository

	ChannelCountry *ChannelCountryRepository

	ChannelCurrency *ChannelCurrencyRepository

	ChannelDomain *ChannelDomainRepository

	ChannelLanguage *ChannelLanguageRepository

	ChannelPaymentMethod *ChannelPaymentMethodRepository

	ChannelTranslation *ChannelTranslationRepository

	ChannelType *ChannelTypeRepository

	ChannelTypeTranslation *ChannelTypeTranslationRepository

	CmsPage *CmsPageRepository

	CmsPageTranslation *CmsPageTranslationRepository

	Country *CountryRepository

	CountryState *CountryStateRepository

	CountryStateTranslation *CountryStateTranslationRepository

	CountryTranslation *CountryTranslationRepository

	Currency *CurrencyRepository

	CurrencyCountryRounding *CurrencyCountryRoundingRepository

	CurrencyTranslation *CurrencyTranslationRepository

	CustomField *CustomFieldRepository

	CustomFieldSet *CustomFieldSetRepository

	CustomFieldSetRelation *CustomFieldSetRelationRepository

	Customer *CustomerRepository

	CustomerGroup *CustomerGroupRepository

	CustomerGroupTranslation *CustomerGroupTranslationRepository

	CustomerRole *CustomerRoleRepository

	CustomerRoleMapping *CustomerRoleMappingRepository

	CustomerTag *CustomerTagRepository

	Dict *DictRepository

	DictItem *DictItemRepository

	DictItemTranslation *DictItemTranslationRepository

	DictTranslation *DictTranslationRepository

	Flow *FlowRepository

	FlowSequence *FlowSequenceRepository

	FlowTemplate *FlowTemplateRepository

	ImportExportFile *ImportExportFileRepository

	ImportExportLog *ImportExportLogRepository

	ImportExportProfile *ImportExportProfileRepository

	Integration *IntegrationRepository

	IntegrationRole *IntegrationRoleRepository

	LandingPage *LandingPageRepository

	LandingPageChannel *LandingPageChannelRepository

	LandingPageTag *LandingPageTagRepository

	LandingPageTranslation *LandingPageTranslationRepository

	Language *LanguageRepository

	Locale *LocaleRepository

	LocaleTranslation *LocaleTranslationRepository

	LogEntry *LogEntryRepository

	MainCategory *MainCategoryRepository

	Media *MediaRepository

	MediaDefaultFolder *MediaDefaultFolderRepository

	MediaFolder *MediaFolderRepository

	MediaFolderConfiguration *MediaFolderConfigurationRepository

	MediaFolderConfigurationMediaThumbnailSize *MediaFolderConfigurationMediaThumbnailSizeRepository

	MediaTag *MediaTagRepository

	MediaThumbnail *MediaThumbnailRepository

	MediaThumbnailSize *MediaThumbnailSizeRepository

	MediaTranslation *MediaTranslationRepository

	Navigation *NavigationRepository

	NavigationTranslation *NavigationTranslationRepository

	NumberRange *NumberRangeRepository

	NumberRangeChannel *NumberRangeChannelRepository

	NumberRangeState *NumberRangeStateRepository

	NumberRangeTranslation *NumberRangeTranslationRepository

	NumberRangeType *NumberRangeTypeRepository

	NumberRangeTypeTranslation *NumberRangeTypeTranslationRepository

	Order *OrderRepository

	OrderCustomer *OrderCustomerRepository

	OrderLineItem *OrderLineItemRepository

	OrderTag *OrderTagRepository

	OrderTransaction *OrderTransactionRepository

	OrderTransactionCapture *OrderTransactionCaptureRepository

	OrderTransactionCaptureRefund *OrderTransactionCaptureRefundRepository

	OrderTransactionCaptureRefundPosition *OrderTransactionCaptureRefundPositionRepository

	PaymentMethod *PaymentMethodRepository

	PaymentMethodTranslation *PaymentMethodTranslationRepository

	Plugin *PluginRepository

	PluginTranslation *PluginTranslationRepository

	Product *ProductRepository

	ProductCategory *ProductCategoryRepository

	ProductCategoryTree *ProductCategoryTreeRepository

	ProductConfiguratorSetting *ProductConfiguratorSettingRepository

	ProductCustomFieldSet *ProductCustomFieldSetRepository

	ProductMedia *ProductMediaRepository

	ProductOption *ProductOptionRepository

	ProductPrice *ProductPriceRepository

	ProductProperty *ProductPropertyRepository

	ProductReview *ProductReviewRepository

	ProductSorting *ProductSortingRepository

	ProductSortingTranslation *ProductSortingTranslationRepository

	ProductTag *ProductTagRepository

	ProductTranslation *ProductTranslationRepository

	ProductVisibility *ProductVisibilityRepository

	PropertyGroup *PropertyGroupRepository

	PropertyGroupOption *PropertyGroupOptionRepository

	PropertyGroupOptionTranslation *PropertyGroupOptionTranslationRepository

	PropertyGroupTranslation *PropertyGroupTranslationRepository

	Rule *RuleRepository

	RuleCondition *RuleConditionRepository

	RuleTag *RuleTagRepository

	ScheduledTask *ScheduledTaskRepository

	SeoUrl *SeoUrlRepository

	SeoUrlTemplate *SeoUrlTemplateRepository

	Snippet *SnippetRepository

	SnippetSet *SnippetSetRepository

	StateMachine *StateMachineRepository

	StateMachineHistory *StateMachineHistoryRepository

	StateMachineState *StateMachineStateRepository

	StateMachineStateTranslation *StateMachineStateTranslationRepository

	StateMachineTransition *StateMachineTransitionRepository

	StateMachineTranslation *StateMachineTranslationRepository

	SystemConfig *SystemConfigRepository

	Tag *TagRepository

	Theme *ThemeRepository

	ThemeChannel *ThemeChannelRepository

	ThemeChild *ThemeChildRepository

	ThemeMedia *ThemeMediaRepository

	ThemeTranslation *ThemeTranslationRepository

	User *UserRepository

	UserAccessKey *UserAccessKeyRepository

	UserConfig *UserConfigRepository

	UserRecovery *UserRecoveryRepository

	Version *VersionRepository

	VersionCommit *VersionCommitRepository

	VersionCommitData *VersionCommitDataRepository
}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

	repo.AclRole = NewAclRoleRepository(client.Client)

	repo.AclUserRole = NewAclUserRoleRepository(client.Client)

	repo.Category = NewCategoryRepository(client.Client)

	repo.CategoryTag = NewCategoryTagRepository(client.Client)

	repo.CategoryTranslation = NewCategoryTranslationRepository(client.Client)

	repo.Channel = NewChannelRepository(client.Client)

	repo.ChannelCountry = NewChannelCountryRepository(client.Client)

	repo.ChannelCurrency = NewChannelCurrencyRepository(client.Client)

	repo.ChannelDomain = NewChannelDomainRepository(client.Client)

	repo.ChannelLanguage = NewChannelLanguageRepository(client.Client)

	repo.ChannelPaymentMethod = NewChannelPaymentMethodRepository(client.Client)

	repo.ChannelTranslation = NewChannelTranslationRepository(client.Client)

	repo.ChannelType = NewChannelTypeRepository(client.Client)

	repo.ChannelTypeTranslation = NewChannelTypeTranslationRepository(client.Client)

	repo.CmsPage = NewCmsPageRepository(client.Client)

	repo.CmsPageTranslation = NewCmsPageTranslationRepository(client.Client)

	repo.Country = NewCountryRepository(client.Client)

	repo.CountryState = NewCountryStateRepository(client.Client)

	repo.CountryStateTranslation = NewCountryStateTranslationRepository(client.Client)

	repo.CountryTranslation = NewCountryTranslationRepository(client.Client)

	repo.Currency = NewCurrencyRepository(client.Client)

	repo.CurrencyCountryRounding = NewCurrencyCountryRoundingRepository(client.Client)

	repo.CurrencyTranslation = NewCurrencyTranslationRepository(client.Client)

	repo.CustomField = NewCustomFieldRepository(client.Client)

	repo.CustomFieldSet = NewCustomFieldSetRepository(client.Client)

	repo.CustomFieldSetRelation = NewCustomFieldSetRelationRepository(client.Client)

	repo.Customer = NewCustomerRepository(client.Client)

	repo.CustomerGroup = NewCustomerGroupRepository(client.Client)

	repo.CustomerGroupTranslation = NewCustomerGroupTranslationRepository(client.Client)

	repo.CustomerRole = NewCustomerRoleRepository(client.Client)

	repo.CustomerRoleMapping = NewCustomerRoleMappingRepository(client.Client)

	repo.CustomerTag = NewCustomerTagRepository(client.Client)

	repo.Dict = NewDictRepository(client.Client)

	repo.DictItem = NewDictItemRepository(client.Client)

	repo.DictItemTranslation = NewDictItemTranslationRepository(client.Client)

	repo.DictTranslation = NewDictTranslationRepository(client.Client)

	repo.Flow = NewFlowRepository(client.Client)

	repo.FlowSequence = NewFlowSequenceRepository(client.Client)

	repo.FlowTemplate = NewFlowTemplateRepository(client.Client)

	repo.ImportExportFile = NewImportExportFileRepository(client.Client)

	repo.ImportExportLog = NewImportExportLogRepository(client.Client)

	repo.ImportExportProfile = NewImportExportProfileRepository(client.Client)

	repo.Integration = NewIntegrationRepository(client.Client)

	repo.IntegrationRole = NewIntegrationRoleRepository(client.Client)

	repo.LandingPage = NewLandingPageRepository(client.Client)

	repo.LandingPageChannel = NewLandingPageChannelRepository(client.Client)

	repo.LandingPageTag = NewLandingPageTagRepository(client.Client)

	repo.LandingPageTranslation = NewLandingPageTranslationRepository(client.Client)

	repo.Language = NewLanguageRepository(client.Client)

	repo.Locale = NewLocaleRepository(client.Client)

	repo.LocaleTranslation = NewLocaleTranslationRepository(client.Client)

	repo.LogEntry = NewLogEntryRepository(client.Client)

	repo.MainCategory = NewMainCategoryRepository(client.Client)

	repo.Media = NewMediaRepository(client.Client)

	repo.MediaDefaultFolder = NewMediaDefaultFolderRepository(client.Client)

	repo.MediaFolder = NewMediaFolderRepository(client.Client)

	repo.MediaFolderConfiguration = NewMediaFolderConfigurationRepository(client.Client)

	repo.MediaFolderConfigurationMediaThumbnailSize = NewMediaFolderConfigurationMediaThumbnailSizeRepository(client.Client)

	repo.MediaTag = NewMediaTagRepository(client.Client)

	repo.MediaThumbnail = NewMediaThumbnailRepository(client.Client)

	repo.MediaThumbnailSize = NewMediaThumbnailSizeRepository(client.Client)

	repo.MediaTranslation = NewMediaTranslationRepository(client.Client)

	repo.Navigation = NewNavigationRepository(client.Client)

	repo.NavigationTranslation = NewNavigationTranslationRepository(client.Client)

	repo.NumberRange = NewNumberRangeRepository(client.Client)

	repo.NumberRangeChannel = NewNumberRangeChannelRepository(client.Client)

	repo.NumberRangeState = NewNumberRangeStateRepository(client.Client)

	repo.NumberRangeTranslation = NewNumberRangeTranslationRepository(client.Client)

	repo.NumberRangeType = NewNumberRangeTypeRepository(client.Client)

	repo.NumberRangeTypeTranslation = NewNumberRangeTypeTranslationRepository(client.Client)

	repo.Order = NewOrderRepository(client.Client)

	repo.OrderCustomer = NewOrderCustomerRepository(client.Client)

	repo.OrderLineItem = NewOrderLineItemRepository(client.Client)

	repo.OrderTag = NewOrderTagRepository(client.Client)

	repo.OrderTransaction = NewOrderTransactionRepository(client.Client)

	repo.OrderTransactionCapture = NewOrderTransactionCaptureRepository(client.Client)

	repo.OrderTransactionCaptureRefund = NewOrderTransactionCaptureRefundRepository(client.Client)

	repo.OrderTransactionCaptureRefundPosition = NewOrderTransactionCaptureRefundPositionRepository(client.Client)

	repo.PaymentMethod = NewPaymentMethodRepository(client.Client)

	repo.PaymentMethodTranslation = NewPaymentMethodTranslationRepository(client.Client)

	repo.Plugin = NewPluginRepository(client.Client)

	repo.PluginTranslation = NewPluginTranslationRepository(client.Client)

	repo.Product = NewProductRepository(client.Client)

	repo.ProductCategory = NewProductCategoryRepository(client.Client)

	repo.ProductCategoryTree = NewProductCategoryTreeRepository(client.Client)

	repo.ProductConfiguratorSetting = NewProductConfiguratorSettingRepository(client.Client)

	repo.ProductCustomFieldSet = NewProductCustomFieldSetRepository(client.Client)

	repo.ProductMedia = NewProductMediaRepository(client.Client)

	repo.ProductOption = NewProductOptionRepository(client.Client)

	repo.ProductPrice = NewProductPriceRepository(client.Client)

	repo.ProductProperty = NewProductPropertyRepository(client.Client)

	repo.ProductReview = NewProductReviewRepository(client.Client)

	repo.ProductSorting = NewProductSortingRepository(client.Client)

	repo.ProductSortingTranslation = NewProductSortingTranslationRepository(client.Client)

	repo.ProductTag = NewProductTagRepository(client.Client)

	repo.ProductTranslation = NewProductTranslationRepository(client.Client)

	repo.ProductVisibility = NewProductVisibilityRepository(client.Client)

	repo.PropertyGroup = NewPropertyGroupRepository(client.Client)

	repo.PropertyGroupOption = NewPropertyGroupOptionRepository(client.Client)

	repo.PropertyGroupOptionTranslation = NewPropertyGroupOptionTranslationRepository(client.Client)

	repo.PropertyGroupTranslation = NewPropertyGroupTranslationRepository(client.Client)

	repo.Rule = NewRuleRepository(client.Client)

	repo.RuleCondition = NewRuleConditionRepository(client.Client)

	repo.RuleTag = NewRuleTagRepository(client.Client)

	repo.ScheduledTask = NewScheduledTaskRepository(client.Client)

	repo.SeoUrl = NewSeoUrlRepository(client.Client)

	repo.SeoUrlTemplate = NewSeoUrlTemplateRepository(client.Client)

	repo.Snippet = NewSnippetRepository(client.Client)

	repo.SnippetSet = NewSnippetSetRepository(client.Client)

	repo.StateMachine = NewStateMachineRepository(client.Client)

	repo.StateMachineHistory = NewStateMachineHistoryRepository(client.Client)

	repo.StateMachineState = NewStateMachineStateRepository(client.Client)

	repo.StateMachineStateTranslation = NewStateMachineStateTranslationRepository(client.Client)

	repo.StateMachineTransition = NewStateMachineTransitionRepository(client.Client)

	repo.StateMachineTranslation = NewStateMachineTranslationRepository(client.Client)

	repo.SystemConfig = NewSystemConfigRepository(client.Client)

	repo.Tag = NewTagRepository(client.Client)

	repo.Theme = NewThemeRepository(client.Client)

	repo.ThemeChannel = NewThemeChannelRepository(client.Client)

	repo.ThemeChild = NewThemeChildRepository(client.Client)

	repo.ThemeMedia = NewThemeMediaRepository(client.Client)

	repo.ThemeTranslation = NewThemeTranslationRepository(client.Client)

	repo.User = NewUserRepository(client.Client)

	repo.UserAccessKey = NewUserAccessKeyRepository(client.Client)

	repo.UserConfig = NewUserConfigRepository(client.Client)

	repo.UserRecovery = NewUserRecoveryRepository(client.Client)

	repo.Version = NewVersionRepository(client.Client)

	repo.VersionCommit = NewVersionCommitRepository(client.Client)

	repo.VersionCommitData = NewVersionCommitDataRepository(client.Client)

	return repo
}
