/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 组合人群信息
type CombineSpec struct {
	Include []IncludeSimpleRule `json:"include,omitempty"`
	Exclude []ExcludeSimpleRule `json:"exclude,omitempty"`
}