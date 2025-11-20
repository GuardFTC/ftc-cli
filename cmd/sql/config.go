// Package sql @Author:冯铁城 [17615007230@163.com] 2025-10-31 19:47:32
package sql

// 默认输出文件路径
var defaultOutput = "C:\\Users\\Administrator\\Downloads\\output.sql"

// 默认数据库
var defaultDB = "dw_tile"

// 默认表
var defaultTable = "ads_bi_af_ltvroas_d_i"

// 表名-列名Map
var tableColumnMap = map[string][]string{
	defaultTable: {
		"game_id", "data_type", "dt", "platform", "country",
		"media_source", "ad_material", "campaign", "campaign_id",
		"is_include_org", "nu", "total_rev", "click", "cost", "impression",
		"rev0", "rev1", "rev2", "rev3", "rev4", "rev5", "rev6", "rev7", "rev8", "rev9",
		"rev10", "rev11", "rev12", "rev13", "rev14", "rev15", "rev30", "rev60", "rev90",
		"rev120", "rev150", "rev180", "rev210", "rev240", "rev270", "rev300", "rev330", "rev360",
	},
	"ads_bi_af_retention_d_i": {
		"game_id", "data_type", "dt", "platform", "country",
		"media_source", "ad_material", "campaign", "campaign_id",
		"cost", "click", "impression", "nu",
		"ua1", "ua2", "ua3", "ua4", "ua5", "ua6", "ua7", "ua8", "ua9", "ua10",
		"ua11", "ua12", "ua13", "ua14", "ua15",
		"ua30", "ua60", "ua90", "ua120", "ua150",
		"ua180", "ua210", "ua240", "ua270",
		"ua300", "ua330", "ua360",
	},
	"ads_bi_af_overview_d_i": {
		"game_id", "data_type", "dt", "platform", "country",
		"media_source", "ad_material", "campaign", "campaign_id", "is_include_org",
		"nu", "ua1", "ua7", "ua30", "click", "cost", "impression", "rev0", "rev1", "rev7", "rev30",
		"td_iaarev", "td_iaprev", "device_launch", "inter_imp", "reward_imp", "inter_rev",
		"reward_rev", "cnt_level",
	},
	"ads_bi_af_material_d_i": {
		"game_id", "dt", "data_type", "platform", "country", "media_source", "ad_material",
		"nu", "rev0", "rev1", "rev7", "rev30", "ua1", "ua7", "ua30",
		"td_iaarev", "td_iaprev", "device_launch", "click", "impression", "cost",
		"total_cost", "total_days", "total_nu", "total_iaarev", "total_iaprev",
	},
	"ads_bi_af_campaign_d_i": {
		"game_id", "dt", "data_type", "platform", "country", "media_source",
		"ad_type", "campaign", "campaign_id",
		"nu", "rev0", "rev1", "rev7", "rev30", "ua1", "ua7", "ua30",
		"td_iaarev", "td_iaprev", "device_launch", "click", "impression", "cost",
		"total_cost", "total_days", "total_nu", "total_iaarev", "total_iaprev",
	},
	"ads_bi_all_overview_d_i": {
		"game_id", "target_day", "data_type", "platform", "country", "media_source",
		"total_device", "total_iaprev", "total_iaarev",
		"cnt_level", "inter_imp", "reward_imp",
		"inter_rev", "reward_rev", "new_device", "old_device",
		"new_iaprev", "old_iaprev", "new_iaarev", "old_iaarev",
		"new_duration", "old_duration",
	},
	"ads_bi_live_data_d_i": {
		"game_id", "create_dt", "data_type", "platform", "country", "media_source", "campaign",
		"create_version_code", "is_org", "living_days",
		"au", "iaa_rev", "iap_rev", "all_rev", "cnt_level",
		"duration", "inter_rev", "reward_rev", "inter_imp",
		"reward_imp", "version_spend",
	},
	"ads_bi_payinfo_product_d_i": {
		"game_id", "target_day", "channel_id", "country",
		"product_id", "product_name", "payed_user", "payed_times", "payed_total",
	},
	"ads_bi_payinfo_rn_d_i": {
		"game_id", "target_day", "channel_id", "country",
		"product_id", "product_name", "product_times", "product_count",
	},
	"ads_bi_payinfo_whale_d_i": {
		"game_id", "target_day", "channel_id", "country",
		"device_id", "server_id", "user_id", "role_id", "role_name",
		"first_login_day", "last_login_day",
		"payed_total",
	},
	"ads_bi_ad_revenue_d_i": {
		"game_id", "dt", "data_type", "platform", "ad_type", "country",
		"ad_unit_id",
		"imp", "rev", "dau",
		"load_success", "start_load",
		"display_start", "display", "display_finish",
		"nextday_user", "today_user",
	},
	"channel_material": {
		"id", "day_time", "log_time", "date_time", "ts", "game_id",
		"platform", "data_type", "country", "media_source", "account_id", "account_name",
		"campaign", "campaign_id", "ad_set_name", "ad_set_id", "ad_name", "ad_id",
		"clicks", "impressions", "spend",
	},
	"af_pull_channel": {
		"id", "day_time", "log_time", "date_time", "ts", "game_id",
		"platform", "data_type", "country", "media_source", "agency",
		"campaign", "ad_group", "ad_type", "ad_create_time",
		"impressions", "clicks", "cost",
	},
	"ads_af_s2s_event_d_i": {
		"id", "day_time",
		"event_name", "event_value", "event_time",
		"device_id", "create_dt", "appsflyer_id",
		"queued", "processed",
		"platform",
	},
	"ads_bi_instantly_d_i": {
		"game_id", "target_day", "data_type", "target_time",
		"platform", "country", "media_source",
		"total_device",
	},
	"ads_bi_contrast_version_d_i": {
		"game_id", "create_version_code", "platform", "type", "country_media", "nu", "cost", "rev0", "rev1", "rev2", "rev3", "rev4", "rev5", "rev6", "rev7", "rev14", "rev30", "rev60", "rev90", "rev120", "rev150", "rev180", "rev210", "rev240",
		"ua0", "ua1", "ua2", "ua3", "ua4", "ua5", "ua6", "ua7", "ua14", "ua30", "ua60", "ua90", "ua120", "ua150", "ua180", "ua210", "ua240",
		"cnt_level0", "cnt_level1", "cnt_level2", "cnt_level3", "cnt_level4", "cnt_level5", "cnt_level6", "cnt_level7", "cnt_level14", "cnt_level30", "cnt_level60", "cnt_level90", "cnt_level120", "cnt_level150", "cnt_level180", "cnt_level210", "cnt_level240",
		"duration0", "duration1", "duration2", "duration3", "duration4", "duration5", "duration6", "duration7", "duration14", "duration30", "duration60", "duration90", "duration120", "duration150", "duration180", "duration210", "duration240",
		"reward_imp0", "reward_imp1", "reward_imp2", "reward_imp3", "reward_imp4", "reward_imp5", "reward_imp6", "reward_imp7", "reward_imp14", "reward_imp30", "reward_imp60", "reward_imp90", "reward_imp120", "reward_imp150", "reward_imp180", "reward_imp210", "reward_imp240",
		"inter_imp0", "inter_imp1", "inter_imp2", "inter_imp3", "inter_imp4", "inter_imp5", "inter_imp6", "inter_imp7", "inter_imp14", "inter_imp30", "inter_imp60", "inter_imp90", "inter_imp120", "inter_imp150", "inter_imp180", "inter_imp210", "inter_imp240",
		"reward_rev0", "reward_rev1", "reward_rev2", "reward_rev3", "reward_rev4", "reward_rev5", "reward_rev6", "reward_rev7", "reward_rev14", "reward_rev30", "reward_rev60", "reward_rev90", "reward_rev120", "reward_rev150", "reward_rev180", "reward_rev210", "reward_rev240",
		"inter_rev0", "inter_rev1", "inter_rev2", "inter_rev3", "inter_rev4", "inter_rev5", "inter_rev6", "inter_rev7", "inter_rev14", "inter_rev30", "inter_rev60", "inter_rev90", "inter_rev120", "inter_rev150", "inter_rev180", "inter_rev210", "inter_rev240",
	},
	"ads_bi_contrast_nu_d_i": {
		"game_id", "create_dt", "platform", "type", "country_media", "nu", "cost", "rev0", "rev1", "rev2", "rev3", "rev4", "rev5", "rev6", "rev7", "rev14", "rev30", "rev60", "rev90", "rev120", "rev150", "rev180", "rev210", "rev240",
		"ua0", "ua1", "ua2", "ua3", "ua4", "ua5", "ua6", "ua7", "ua14", "ua30", "ua60", "ua90", "ua120", "ua150", "ua180", "ua210", "ua240",
		"cnt_level0", "cnt_level1", "cnt_level2", "cnt_level3", "cnt_level4", "cnt_level5", "cnt_level6", "cnt_level7", "cnt_level14", "cnt_level30", "cnt_level60", "cnt_level90", "cnt_level120", "cnt_level150", "cnt_level180", "cnt_level210", "cnt_level240",
		"duration0", "duration1", "duration2", "duration3", "duration4", "duration5", "duration6", "duration7", "duration14", "duration30", "duration60", "duration90", "duration120", "duration150", "duration180", "duration210", "duration240",
		"reward_imp0", "reward_imp1", "reward_imp2", "reward_imp3", "reward_imp4", "reward_imp5", "reward_imp6", "reward_imp7", "reward_imp14", "reward_imp30", "reward_imp60", "reward_imp90", "reward_imp120", "reward_imp150", "reward_imp180", "reward_imp210", "reward_imp240",
		"inter_imp0", "inter_imp1", "inter_imp2", "inter_imp3", "inter_imp4", "inter_imp5", "inter_imp6", "inter_imp7", "inter_imp14", "inter_imp30", "inter_imp60", "inter_imp90", "inter_imp120", "inter_imp150", "inter_imp180", "inter_imp210", "inter_imp240",
		"reward_rev0", "reward_rev1", "reward_rev2", "reward_rev3", "reward_rev4", "reward_rev5", "reward_rev6", "reward_rev7", "reward_rev14", "reward_rev30", "reward_rev60", "reward_rev90", "reward_rev120", "reward_rev150", "reward_rev180", "reward_rev210", "reward_rev240",
		"inter_rev0", "inter_rev1", "inter_rev2", "inter_rev3", "inter_rev4", "inter_rev5", "inter_rev6", "inter_rev7", "inter_rev14", "inter_rev30", "inter_rev60", "inter_rev90", "inter_rev120", "inter_rev150", "inter_rev180", "inter_rev210", "inter_rev240",
	},
}
