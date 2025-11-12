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
}
