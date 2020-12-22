/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type DailyReportsGetListStruct struct {
	AccountId                           int64   `json:"account_id,omitempty"`
	Date                                string  `json:"date,omitempty"`
	ViewCount                           int64   `json:"view_count,omitempty"`
	DownloadCount                       int64   `json:"download_count,omitempty"`
	ActivatedCount                      int64   `json:"activated_count,omitempty"`
	ActivatedRate                       float64 `json:"activated_rate,omitempty"`
	ThousandDisplayPrice                int64   `json:"thousand_display_price,omitempty"`
	ValidClickCount                     int64   `json:"valid_click_count,omitempty"`
	Ctr                                 float64 `json:"ctr,omitempty"`
	Cpc                                 int64   `json:"cpc,omitempty"`
	Cost                                int64   `json:"cost,omitempty"`
	KeyPageViewCost                     int64   `json:"key_page_view_cost,omitempty"`
	CouponClickCount                    int64   `json:"coupon_click_count,omitempty"`
	CouponIssueCount                    int64   `json:"coupon_issue_count,omitempty"`
	CouponGetCount                      int64   `json:"coupon_get_count,omitempty"`
	PlatformPageViewCount               int64   `json:"platform_page_view_count,omitempty"`
	PlatformPageViewRate                float64 `json:"platform_page_view_rate,omitempty"`
	WebCommodityPageViewCount           int64   `json:"web_commodity_page_view_count,omitempty"`
	WebCommodityPageViewCost            int64   `json:"web_commodity_page_view_cost,omitempty"`
	WebRegisterCount                    int64   `json:"web_register_count,omitempty"`
	PageConsultCount                    int64   `json:"page_consult_count,omitempty"`
	PageConsultCost                     int64   `json:"page_consult_cost,omitempty"`
	PagePhoneCallDirectCount            int64   `json:"page_phone_call_direct_count,omitempty"`
	PagePhoneCallDirectCost             int64   `json:"page_phone_call_direct_cost,omitempty"`
	PagePhoneCallBackCount              int64   `json:"page_phone_call_back_count,omitempty"`
	PagePhoneCallBackCost               int64   `json:"page_phone_call_back_cost,omitempty"`
	OwnPageNavigationCount              int64   `json:"own_page_navigation_count,omitempty"`
	OwnPageNaviCost                     int64   `json:"own_page_navi_cost,omitempty"`
	PlatformPageNavigationCount         int64   `json:"platform_page_navigation_count,omitempty"`
	PlatformPageNavigationCost          int64   `json:"platform_page_navigation_cost,omitempty"`
	PlatformShopNavigationCount         int64   `json:"platform_shop_navigation_count,omitempty"`
	PlatformShopNavigationCost          int64   `json:"platform_shop_navigation_cost,omitempty"`
	WebApplicationCount                 int64   `json:"web_application_count,omitempty"`
	WebApplicationCost                  int64   `json:"web_application_cost,omitempty"`
	PageReservationCount                int64   `json:"page_reservation_count,omitempty"`
	PageReservationRate                 float64 `json:"page_reservation_rate,omitempty"`
	PageReservationCost                 int64   `json:"page_reservation_cost,omitempty"`
	WebAddToCartCount                   int64   `json:"web_add_to_cart_count,omitempty"`
	WebAddToCartCost                    int64   `json:"web_add_to_cart_cost,omitempty"`
	AddToCartPrice                      int64   `json:"add_to_cart_price,omitempty"`
	OwnPageCouponGetCount               int64   `json:"own_page_coupon_get_count,omitempty"`
	OwnPageCouponGetCost                int64   `json:"own_page_coupon_get_cost,omitempty"`
	PlatformCouponGetCount              int64   `json:"platform_coupon_get_count,omitempty"`
	PlatformCouponGetCost               int64   `json:"platform_coupon_get_cost,omitempty"`
	WebOrderCount                       int64   `json:"web_order_count,omitempty"`
	WebOrderRate                        float64 `json:"web_order_rate,omitempty"`
	AppOrderRate                        float64 `json:"app_order_rate,omitempty"`
	WebOrderCost                        int64   `json:"web_order_cost,omitempty"`
	WebCheckoutAmount                   int64   `json:"web_checkout_amount,omitempty"`
	WebCheckoutCount                    int64   `json:"web_checkout_count,omitempty"`
	WebCheckoutCost                     int64   `json:"web_checkout_cost,omitempty"`
	OrderAmount                         int64   `json:"order_amount,omitempty"`
	OrderUnitPrice                      int64   `json:"order_unit_price,omitempty"`
	OrderRoi                            float64 `json:"order_roi,omitempty"`
	DeliverCount                        int64   `json:"deliver_count,omitempty"`
	DeliverCost                         int64   `json:"deliver_cost,omitempty"`
	SignInCount                         int64   `json:"sign_in_count,omitempty"`
	SignInCost                          int64   `json:"sign_in_cost,omitempty"`
	DownloadRate                        float64 `json:"download_rate,omitempty"`
	DownloadCost                        int64   `json:"download_cost,omitempty"`
	InstallCount                        int64   `json:"install_count,omitempty"`
	InstallCost                         int64   `json:"install_cost,omitempty"`
	ClickActivatedRate                  float64 `json:"click_activated_rate,omitempty"`
	ActivatedCost                       int64   `json:"activated_cost,omitempty"`
	RetentionCount                      int64   `json:"retention_count,omitempty"`
	RetentionRate                       float64 `json:"retention_rate,omitempty"`
	RetentionCost                       int64   `json:"retention_cost,omitempty"`
	KeyPageViewCount                    int64   `json:"key_page_view_count,omitempty"`
	AppCommodityPageViewCount           int64   `json:"app_commodity_page_view_count,omitempty"`
	AppCommodityPageViewRate            float64 `json:"app_commodity_page_view_rate,omitempty"`
	WebCommodityPageViewRate            float64 `json:"web_commodity_page_view_rate,omitempty"`
	AppCommodityPageViewCost            int64   `json:"app_commodity_page_view_cost,omitempty"`
	AppRegisterCount                    int64   `json:"app_register_count,omitempty"`
	AppRegisterCost                     int64   `json:"app_register_cost,omitempty"`
	WebRegisterCost                     int64   `json:"web_register_cost,omitempty"`
	AppApplicationCount                 int64   `json:"app_application_count,omitempty"`
	AppApplicationCost                  int64   `json:"app_application_cost,omitempty"`
	AppAddToCartCount                   int64   `json:"app_add_to_cart_count,omitempty"`
	AddToCartAmount                     int64   `json:"add_to_cart_amount,omitempty"`
	AppAddToCartCost                    int64   `json:"app_add_to_cart_cost,omitempty"`
	AppOrderCount                       int64   `json:"app_order_count,omitempty"`
	AppOrderCost                        int64   `json:"app_order_cost,omitempty"`
	AppCheckoutCount                    int64   `json:"app_checkout_count,omitempty"`
	AppCheckoutAmount                   int64   `json:"app_checkout_amount,omitempty"`
	AppCheckoutCost                     int64   `json:"app_checkout_cost,omitempty"`
	PlatformCouponClickCount            int64   `json:"platform_coupon_click_count,omitempty"`
	PlatformCouponGetRate               float64 `json:"platform_coupon_get_rate,omitempty"`
	FollowCount                         int64   `json:"follow_count,omitempty"`
	FollowCost                          int64   `json:"follow_cost,omitempty"`
	ForwardCount                        int64   `json:"forward_count,omitempty"`
	ForwardCost                         int64   `json:"forward_cost,omitempty"`
	ReadCount                           int64   `json:"read_count,omitempty"`
	ReadCost                            int64   `json:"read_cost,omitempty"`
	PraiseCount                         int64   `json:"praise_count,omitempty"`
	PraiseCost                          int64   `json:"praise_cost,omitempty"`
	CommentCount                        int64   `json:"comment_count,omitempty"`
	CommentCost                         int64   `json:"comment_cost,omitempty"`
	IntePhoneCount                      int64   `json:"inte_phone_count,omitempty"`
	PhoneCallCount                      int64   `json:"phone_call_count,omitempty"`
	ExternalFormReservationCount        int64   `json:"external_form_reservation_count,omitempty"`
	AppAdPayingUsers                    int64   `json:"app_ad_paying_users,omitempty"`
	AdPurValWeb                         int64   `json:"ad_pur_val_web,omitempty"`
	AdPurValApp                         int64   `json:"ad_pur_val_app,omitempty"`
	GameCreateRoleCount                 int64   `json:"game_create_role_count,omitempty"`
	GameAuthorizeCount                  int64   `json:"game_authorize_count,omitempty"`
	GameTutorialFinishCount             int64   `json:"game_tutorial_finish_count,omitempty"`
	EffectiveLeadsCount                 int64   `json:"effective_leads_count,omitempty"`
	EffectiveCost                       int64   `json:"effective_cost,omitempty"`
	EffectiveReserveCount               int64   `json:"effective_reserve_count,omitempty"`
	EffectiveConsultCount               int64   `json:"effective_consult_count,omitempty"`
	EffectivePhoneCount                 int64   `json:"effective_phone_count,omitempty"`
	PotentialReserveCount               int64   `json:"potential_reserve_count,omitempty"`
	PotentialConsultCount               int64   `json:"potential_consult_count,omitempty"`
	PotentialPhoneCount                 int64   `json:"potential_phone_count,omitempty"`
	AppCheckoutRate                     float64 `json:"app_checkout_rate,omitempty"`
	WebCheckoutRate                     float64 `json:"web_checkout_rate,omitempty"`
	AppActivatedCheckoutRate            float64 `json:"app_activated_checkout_rate,omitempty"`
	WebActivatedCheckoutRate            float64 `json:"web_activated_checkout_rate,omitempty"`
	AppRegisterRate                     float64 `json:"app_register_rate,omitempty"`
	WebRegRate                          float64 `json:"web_reg_rate,omitempty"`
	PagePhoneCallDirectRate             float64 `json:"page_phone_call_direct_rate,omitempty"`
	PagePhoneCallBackRate               float64 `json:"page_phone_call_back_rate,omitempty"`
	PageConsultRate                     float64 `json:"page_consult_rate,omitempty"`
	DeliverRate                         float64 `json:"deliver_rate,omitempty"`
	InstallRate                         float64 `json:"install_rate,omitempty"`
	ArppuCost                           int64   `json:"arppu_cost,omitempty"`
	ArpuCost                            int64   `json:"arpu_cost,omitempty"`
	WebArppuCost                        int64   `json:"web_arppu_cost,omitempty"`
	WebArpuCost                         int64   `json:"web_arpu_cost,omitempty"`
	AppAdPurArpuCost                    int64   `json:"app_ad_pur_arpu_cost,omitempty"`
	AppAdPurArppuCost                   int64   `json:"app_ad_pur_arppu_cost,omitempty"`
	WebAdPurArpuCost                    int64   `json:"web_ad_pur_arpu_cost,omitempty"`
	CheoutFd                            int64   `json:"cheout_fd,omitempty"`
	CheoutTd                            int64   `json:"cheout_td,omitempty"`
	CheoutOw                            int64   `json:"cheout_ow,omitempty"`
	CheoutTw                            int64   `json:"cheout_tw,omitempty"`
	CheoutOm                            int64   `json:"cheout_om,omitempty"`
	CheoutFdReward                      float64 `json:"cheout_fd_reward,omitempty"`
	CheoutTdReward                      float64 `json:"cheout_td_reward,omitempty"`
	CheoutOwReward                      float64 `json:"cheout_ow_reward,omitempty"`
	CheoutTwReward                      float64 `json:"cheout_tw_reward,omitempty"`
	CheoutOmReward                      float64 `json:"cheout_om_reward,omitempty"`
	CheoutTotalReward                   float64 `json:"cheout_total_reward,omitempty"`
	FromFollowUv                        int64   `json:"from_follow_uv,omitempty"`
	FromFollowCost                      int64   `json:"from_follow_cost,omitempty"`
	AddDesktopPv                        int64   `json:"add_desktop_pv,omitempty"`
	AddDesktopCost                      int64   `json:"add_desktop_cost,omitempty"`
	FirstPayCount                       int64   `json:"first_pay_count,omitempty"`
	FirstPayRate                        float64 `json:"first_pay_rate,omitempty"`
	PreCreWeb                           int64   `json:"pre_cre_web,omitempty"`
	PreCreApp                           int64   `json:"pre_cre_app,omitempty"`
	PreCreWebVal                        int64   `json:"pre_cre_web_val,omitempty"`
	PreCreAppVal                        int64   `json:"pre_cre_app_val,omitempty"`
	CreWeb                              int64   `json:"cre_web,omitempty"`
	CreApp                              int64   `json:"cre_app,omitempty"`
	CreWebVal                           int64   `json:"cre_web_val,omitempty"`
	CreAppVal                           int64   `json:"cre_app_val,omitempty"`
	WithdrDepWeb                        int64   `json:"withdr_dep_web,omitempty"`
	WithdrDepApp                        int64   `json:"withdr_dep_app,omitempty"`
	WithdrDepWebVal                     int64   `json:"withdr_dep_web_val,omitempty"`
	WithdrDepAppVal                     int64   `json:"withdr_dep_app_val,omitempty"`
	FirstPayCost                        int64   `json:"first_pay_cost,omitempty"`
	LandingPageClickCount               int64   `json:"landing_page_click_count,omitempty"`
	WebCartAmount                       int64   `json:"web_cart_amount,omitempty"`
	ScanFollowCount                     int64   `json:"scan_follow_count,omitempty"`
	ScanFollowUserCount                 int64   `json:"scan_follow_user_count,omitempty"`
	ScanFollowUserCost                  int64   `json:"scan_follow_user_cost,omitempty"`
	ScanFollowUserRate                  float64 `json:"scan_follow_user_rate,omitempty"`
	AppOrderAmount                      int64   `json:"app_order_amount,omitempty"`
	WebOrderAmount                      int64   `json:"web_order_amount,omitempty"`
	PhoneConsultCount                   int64   `json:"phone_consult_count,omitempty"`
	ToolConsultCount                    int64   `json:"tool_consult_count,omitempty"`
	LotteryLeadsCount                   int64   `json:"lottery_leads_count,omitempty"`
	LotteryLeadsCost                    int64   `json:"lottery_leads_cost,omitempty"`
	ConversionsCount                    int64   `json:"conversions_count,omitempty"`
	ConversionsRate                     float64 `json:"conversions_rate,omitempty"`
	ConversionsCost                     int64   `json:"conversions_cost,omitempty"`
	DeepConversionsCount                int64   `json:"deep_conversions_count,omitempty"`
	DeepConversionsRate                 float64 `json:"deep_conversions_rate,omitempty"`
	DeepConversionsCost                 int64   `json:"deep_conversions_cost,omitempty"`
	FirstMemcardWebCount                int64   `json:"first_memcard_web_count,omitempty"`
	FirstMemcardAppCount                int64   `json:"first_memcard_app_count,omitempty"`
	MemcardWebCount                     int64   `json:"memcard_web_count,omitempty"`
	MemcardAppCount                     int64   `json:"memcard_app_count,omitempty"`
	FirstMemcardWebRate                 float64 `json:"first_memcard_web_rate,omitempty"`
	FirstMemcardAppRate                 float64 `json:"first_memcard_app_rate,omitempty"`
	FirstMemcardWebCost                 int64   `json:"first_memcard_web_cost,omitempty"`
	FirstMemcardAppCost                 int64   `json:"first_memcard_app_cost,omitempty"`
	ValuableClickCount                  int64   `json:"valuable_click_count,omitempty"`
	ValuableClickRate                   float64 `json:"valuable_click_rate,omitempty"`
	ValuableClickCost                   int64   `json:"valuable_click_cost,omitempty"`
	VideoPlayCount                      int64   `json:"video_play_count,omitempty"`
	ClickImageCount                     int64   `json:"click_image_count,omitempty"`
	ClickDetailCount                    int64   `json:"click_detail_count,omitempty"`
	ClickHeadCount                      int64   `json:"click_head_count,omitempty"`
	ClickNickCount                      int64   `json:"click_nick_count,omitempty"`
	ClickPoiCount                       int64   `json:"click_poi_count,omitempty"`
	VideoInnerPlayCount                 int64   `json:"video_inner_play_count,omitempty"`
	LanButtonClickCount                 int64   `json:"lan_button_click_count,omitempty"`
	LanJumpButtonClickers               int64   `json:"lan_jump_button_clickers,omitempty"`
	LanJumpButtonClickCost              int64   `json:"lan_jump_button_click_cost,omitempty"`
	LanJumpButtonCtr                    float64 `json:"lan_jump_button_ctr,omitempty"`
	LanButtonClickCost                  int64   `json:"lan_button_click_cost,omitempty"`
	CpnClickButtonCount                 int64   `json:"cpn_click_button_count,omitempty"`
	CpnClickButtonUv                    int64   `json:"cpn_click_button_uv,omitempty"`
	KeyPageUv                           int64   `json:"key_page_uv,omitempty"`
	SpecialPageExpUv                    int64   `json:"special_page_exp_uv,omitempty"`
	SpecialPageExpCost                  int64   `json:"special_page_exp_cost,omitempty"`
	ViewCommodityPageUv                 int64   `json:"view_commodity_page_uv,omitempty"`
	EffectLeadsPurchaseCount            int64   `json:"effect_leads_purchase_count,omitempty"`
	ReservationUv                       int64   `json:"reservation_uv,omitempty"`
	OverallLeadsPurchaseCount           int64   `json:"overall_leads_purchase_count,omitempty"`
	LeadsPurchaseCount                  int64   `json:"leads_purchase_count,omitempty"`
	LeadsPurchaseRate                   float64 `json:"leads_purchase_rate,omitempty"`
	LeadsPurchaseCost                   int64   `json:"leads_purchase_cost,omitempty"`
	LeadsPurchaseUv                     int64   `json:"leads_purchase_uv,omitempty"`
	ValidLeadsUv                        int64   `json:"valid_leads_uv,omitempty"`
	PhoneCallUv                         int64   `json:"phone_call_uv,omitempty"`
	ValidPhoneUv                        int64   `json:"valid_phone_uv,omitempty"`
	PotentialCustomerPhoneUv            int64   `json:"potential_customer_phone_uv,omitempty"`
	WebRegisterUv                       int64   `json:"web_register_uv,omitempty"`
	WebApplyUv                          int64   `json:"web_apply_uv,omitempty"`
	WebCreditUv                         int64   `json:"web_credit_uv,omitempty"`
	AppApplyUv                          int64   `json:"app_apply_uv,omitempty"`
	AppPreCreditUv                      int64   `json:"app_pre_credit_uv,omitempty"`
	AppCreditUv                         int64   `json:"app_credit_uv,omitempty"`
	AppWithdrawUv                       int64   `json:"app_withdraw_uv,omitempty"`
	WechatAppRegisterUv                 int64   `json:"wechat_app_register_uv,omitempty"`
	NoInterestCount                     int64   `json:"no_interest_count,omitempty"`
	FirstDayOrderCount                  int64   `json:"first_day_order_count,omitempty"`
	FirstDayOrderAmount                 int64   `json:"first_day_order_amount,omitempty"`
	AddWishlistCount                    int64   `json:"add_wishlist_count,omitempty"`
	VideoOuterPlay10Count               int64   `json:"video_outer_play10_count,omitempty"`
	VideoOuterPlay25Count               int64   `json:"video_outer_play25_count,omitempty"`
	VideoOuterPlay50Count               int64   `json:"video_outer_play50_count,omitempty"`
	VideoOuterPlay75Count               int64   `json:"video_outer_play75_count,omitempty"`
	VideoOuterPlay95Count               int64   `json:"video_outer_play95_count,omitempty"`
	VideoOuterPlay100Count              int64   `json:"video_outer_play100_count,omitempty"`
	VideoOuterPlayTimeCount             float64 `json:"video_outer_play_time_count,omitempty"`
	VideoOuterPlayTimeAvgRate           float64 `json:"video_outer_play_time_avg_rate,omitempty"`
	VideoOuterPlayRate                  float64 `json:"video_outer_play_rate,omitempty"`
	VideoOuterPlayCost                  int64   `json:"video_outer_play_cost,omitempty"`
	VideoOuterPlayCount                 int64   `json:"video_outer_play_count,omitempty"`
	VideoOuterPlay3sCount               int64   `json:"video_outer_play3s_count,omitempty"`
	VideoOuterPlay5sCount               int64   `json:"video_outer_play5s_count,omitempty"`
	VideoOuterPlay7sCount               int64   `json:"video_outer_play7s_count,omitempty"`
	EffectLeadsPurchaseCost             int64   `json:"effect_leads_purchase_cost,omitempty"`
	CreWebCost                          int64   `json:"cre_web_cost,omitempty"`
	CreAppCost                          int64   `json:"cre_app_cost,omitempty"`
	PreCreWebCost                       int64   `json:"pre_cre_web_cost,omitempty"`
	PreCreAppCost                       int64   `json:"pre_cre_app_cost,omitempty"`
	StoreVisitor                        int64   `json:"store_visitor,omitempty"`
	TryOutUser                          int64   `json:"try_out_user,omitempty"`
	ConsultLeaveInfoUsers               int64   `json:"consult_leave_info_users,omitempty"`
	ActivePageViews                     int64   `json:"active_page_views,omitempty"`
	ActivePageViewers                   int64   `json:"active_page_viewers,omitempty"`
	ActivePageInteractionAmount         int64   `json:"active_page_interaction_amount,omitempty"`
	ActivePageInteractionUsers          int64   `json:"active_page_interaction_users,omitempty"`
	JoinChatGroupAmount                 int64   `json:"join_chat_group_amount,omitempty"`
	GuideToFollowPageViews              int64   `json:"guide_to_follow_page_views,omitempty"`
	GuideToFollowPageViewers            int64   `json:"guide_to_follow_page_viewers,omitempty"`
	GuideToFollowPageInteractionAmount  int64   `json:"guide_to_follow_page_interaction_amount,omitempty"`
	GuideToFollowPageInteractionUsers   int64   `json:"guide_to_follow_page_interaction_users,omitempty"`
	RequestConversionsCount             int64   `json:"request_conversions_count,omitempty"`
	RequestConversionsCost              int64   `json:"request_conversions_cost,omitempty"`
	IncomeRoi1                          float64 `json:"income_roi_1,omitempty"`
	IncomeRoi3                          float64 `json:"income_roi_3,omitempty"`
	IncomeRoi7                          float64 `json:"income_roi_7,omitempty"`
	IncomeRoi14                         float64 `json:"income_roi_14,omitempty"`
	ActivatedTotalPaymentCost           int64   `json:"activated_total_payment_cost,omitempty"`
	PaymentAmountActivatedD3            int64   `json:"payment_amount_activated_d3,omitempty"`
	PaymentAmountActivatedD7            int64   `json:"payment_amount_activated_d7,omitempty"`
	PaymentAmountActivatedD14           int64   `json:"payment_amount_activated_d14,omitempty"`
	PaymentAmountActivatedD30           int64   `json:"payment_amount_activated_d30,omitempty"`
	FirstDayPayCount                    int64   `json:"first_day_pay_count,omitempty"`
	FirstDayPayCost                     int64   `json:"first_day_pay_cost,omitempty"`
	FirstDayFirstPayCost                int64   `json:"first_day_first_pay_cost,omitempty"`
	FirstDayFirstPayCount               int64   `json:"first_day_first_pay_count,omitempty"`
	PaymentCostActivatedD1              int64   `json:"payment_cost_activated_d1,omitempty"`
	FirstDayPayAmount                   int64   `json:"first_day_pay_amount,omitempty"`
	RoiActivatedD1                      float64 `json:"roi_activated_d1,omitempty"`
	RoiActivatedD3                      float64 `json:"roi_activated_d3,omitempty"`
	RoiActivatedD7                      float64 `json:"roi_activated_d7,omitempty"`
	RoiActivatedD14                     float64 `json:"roi_activated_d14,omitempty"`
	RoiActivatedD30                     float64 `json:"roi_activated_d30,omitempty"`
	FirstDayFirstPayRate                float64 `json:"first_day_first_pay_rate,omitempty"`
	CampaignId                          int64   `json:"campaign_id,omitempty"`
	CampaignName                        string  `json:"campaign_name,omitempty"`
	AdgroupId                           int64   `json:"adgroup_id,omitempty"`
	AdgroupName                         string  `json:"adgroup_name,omitempty"`
	AdId                                int64   `json:"ad_id,omitempty"`
	AdName                              string  `json:"ad_name,omitempty"`
	PromotedObjectType                  string  `json:"promoted_object_type,omitempty"`
	PromotedObjectId                    string  `json:"promoted_object_id,omitempty"`
	UnionPositionId                     int64   `json:"union_position_id,omitempty"`
	IndustryParentName                  string  `json:"industry_parent_name,omitempty"`
	PlacementName                       string  `json:"placement_name,omitempty"`
	TemplateId                          int64   `json:"template_id,omitempty"`
	AddToCartCost                       int64   `json:"add_to_cart_cost,omitempty"`
	WebAddToCartRate                    float64 `json:"web_add_to_cart_rate,omitempty"`
	AppAddToCartRate                    float64 `json:"app_add_to_cart_rate,omitempty"`
	IsExpandTargeting                   bool    `json:"is_expand_targeting,omitempty"`
	SiteSet                             string  `json:"site_set,omitempty"`
	MaterialId                          int64   `json:"material_id,omitempty"`
	Signature                           string  `json:"signature,omitempty"`
	ProductCatalogId                    int64   `json:"product_catalog_id,omitempty"`
	ProductOuterId                      string  `json:"product_outer_id,omitempty"`
	BrandId                             int64   `json:"brand_id,omitempty"`
	ProductSetId                        int64   `json:"product_set_id,omitempty"`
	ShopId                              int64   `json:"shop_id,omitempty"`
	FirstCategoryId                     int64   `json:"first_category_id,omitempty"`
	SecondCategoryId                    int64   `json:"second_category_id,omitempty"`
	ThirdCategoryId                     int64   `json:"third_category_id,omitempty"`
	WechatAccountId                     string  `json:"wechat_account_id,omitempty"`
	WechatAgencyId                      string  `json:"wechat_agency_id,omitempty"`
	CostDeviationRate                   float64 `json:"cost_deviation_rate,omitempty"`
	CompensationAmount                  int64   `json:"compensation_amount,omitempty"`
	ViewUserCount                       int64   `json:"view_user_count,omitempty"`
	AvgViewPerUser                      float64 `json:"avg_view_per_user,omitempty"`
	ClickUserCount                      int64   `json:"click_user_count,omitempty"`
	ImageClickUserCount                 int64   `json:"image_click_user_count,omitempty"`
	VideoClickUserCount                 int64   `json:"video_click_user_count,omitempty"`
	LinkClickUserCount                  int64   `json:"link_click_user_count,omitempty"`
	PortraitClickUserCount              int64   `json:"portrait_click_user_count,omitempty"`
	NicknameClickUserCount              int64   `json:"nickname_click_user_count,omitempty"`
	PoiClickUserCount                   int64   `json:"poi_click_user_count,omitempty"`
	PlatformKeyPageViewUserCount        int64   `json:"platform_key_page_view_user_count,omitempty"`
	PlatformKeyPageAvgViewPerUser       float64 `json:"platform_key_page_avg_view_per_user,omitempty"`
	PlatformKeyPageViewDuration         float64 `json:"platform_key_page_view_duration,omitempty"`
	CpnClickButtonCost                  int64   `json:"cpn_click_button_cost,omitempty"`
	PraiseUserCount                     int64   `json:"praise_user_count,omitempty"`
	CommentUserCount                    int64   `json:"comment_user_count,omitempty"`
	OrderCount                          int64   `json:"order_count,omitempty"`
	OrderRate                           float64 `json:"order_rate,omitempty"`
	QuestReservationPvCost              int64   `json:"quest_reservation_pv_cost,omitempty"`
	LeadsRate                           float64 `json:"leads_rate,omitempty"`
	LeadsUserCount                      int64   `json:"leads_user_count,omitempty"`
	LeadsCost                           int64   `json:"leads_cost,omitempty"`
	LeadsUserRate                       float64 `json:"leads_user_rate,omitempty"`
	ValidLeadsCost                      int64   `json:"valid_leads_cost,omitempty"`
	ValidLeadsRate                      float64 `json:"valid_leads_rate,omitempty"`
	PageConsultUserCount                int64   `json:"page_consult_user_count,omitempty"`
	ValidPageConsultUserCount           int64   `json:"valid_page_consult_user_count,omitempty"`
	WithdrDepWebUserCount               int64   `json:"withdr_dep_web_user_count,omitempty"`
	WechatMinigameRegisterCost          int64   `json:"wechat_minigame_register_cost,omitempty"`
	WechatMinigameRegisterRate          float64 `json:"wechat_minigame_register_rate,omitempty"`
	WechatMinigameArpu                  float64 `json:"wechat_minigame_arpu,omitempty"`
	WechatMinigameRetentionCount        int64   `json:"wechat_minigame_retention_count,omitempty"`
	WechatMinigameCheckoutCount         int64   `json:"wechat_minigame_checkout_count,omitempty"`
	WechatMinigameCheckoutAmount        int64   `json:"wechat_minigame_checkout_amount,omitempty"`
	OfficialAccountFollowCount          int64   `json:"official_account_follow_count,omitempty"`
	OfficialAccountFollowCost           int64   `json:"official_account_follow_cost,omitempty"`
	OfficialAccountFollowRate           float64 `json:"official_account_follow_rate,omitempty"`
	OfficialAccountRegisterUserCount    int64   `json:"official_account_register_user_count,omitempty"`
	OfficialAccountRegisterRate         float64 `json:"official_account_register_rate,omitempty"`
	OfficialAccountRegisterCost         int64   `json:"official_account_register_cost,omitempty"`
	OfficialAccountRegisterAmount       int64   `json:"official_account_register_amount,omitempty"`
	OfficialAccountRegisterRoi          int64   `json:"official_account_register_roi,omitempty"`
	OfficialAccountApplyCount           int64   `json:"official_account_apply_count,omitempty"`
	OfficialAccountApplyUserCount       int64   `json:"official_account_apply_user_count,omitempty"`
	OfficialAccountApplyRate            float64 `json:"official_account_apply_rate,omitempty"`
	OfficialAccountApplyCost            int64   `json:"official_account_apply_cost,omitempty"`
	OfficialAccountApplyAmount          int64   `json:"official_account_apply_amount,omitempty"`
	OfficialAccountApplyRoi             int64   `json:"official_account_apply_roi,omitempty"`
	OfficialAccountOrderCount           int64   `json:"official_account_order_count,omitempty"`
	OfficialAccountFirstDayOrderCount   int64   `json:"official_account_first_day_order_count,omitempty"`
	OfficialAccountOrderUserCount       int64   `json:"official_account_order_user_count,omitempty"`
	OfficialAccountOrderRate            float64 `json:"official_account_order_rate,omitempty"`
	OfficialAccountOrderCost            int64   `json:"official_account_order_cost,omitempty"`
	OfficialAccountOrderAmount          int64   `json:"official_account_order_amount,omitempty"`
	OfficialAccountFirstDayOrderAmount  int64   `json:"official_account_first_day_order_amount,omitempty"`
	OfficialAccountOrderRoi             int64   `json:"official_account_order_roi,omitempty"`
	OfficialAccountConsultCount         int64   `json:"official_account_consult_count,omitempty"`
	OfficialAccountReaderCount          int64   `json:"official_account_reader_count,omitempty"`
	OfficialAccountCreditApplyUserCount int64   `json:"official_account_credit_apply_user_count,omitempty"`
	OfficialAccountCreditUserCount      int64   `json:"official_account_credit_user_count,omitempty"`
	CouponGetCost                       int64   `json:"coupon_get_cost,omitempty"`
	CouponGetRate                       float64 `json:"coupon_get_rate,omitempty"`
	CouponUseCount                      int64   `json:"coupon_use_count,omitempty"`
	ForwardUserCount                    int64   `json:"forward_user_count,omitempty"`
	VideoOuterPlayUserCount             int64   `json:"video_outer_play_user_count,omitempty"`
	VideoInnerPlayUserCount             int64   `json:"video_inner_play_user_count,omitempty"`
}
