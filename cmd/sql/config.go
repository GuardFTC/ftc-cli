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
}
