/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告组结构
type AdgroupCreateSpec struct {
	AdgroupName       string                   `json:"adgroup_name,omitempty"`
	BeginDate         string                   `json:"begin_date,omitempty"`
	Targeting         *DpWriteTargetingSetting `json:"targeting,omitempty"`
	PoiList           *[]string                `json:"poi_list,omitempty"`
	PackageLevel      PackageLevel             `json:"package_level,omitempty"`
	ColdStartAudience *[]int64                 `json:"cold_start_audience,omitempty"`
}
