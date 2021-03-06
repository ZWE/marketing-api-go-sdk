/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdDiagnosisGetRequest struct {
	AccountId     int64     `json:"account_id,omitempty"`
	AdgroupIdList *[]int64  `json:"adgroup_id_list,omitempty"`
	DetailFields  *[]string `json:"detail_fields,omitempty"`
}
