/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告人群信息
type AdRuleSpec struct {
	RuleType       AdRuleType  `json:"rule_type,omitempty"`
	ConversionType []string    `json:"conversion_type,omitempty"`
	StartDate      string      `json:"start_date,omitempty"`
	EndDate        string      `json:"end_date,omitempty"`
	CampaignIdList []int64     `json:"campaign_id_list,omitempty"`
	ProductList    []AdProduct `json:"product_list,omitempty"`
	AdgroupIdList  []int64     `json:"adgroup_id_list,omitempty"`
}