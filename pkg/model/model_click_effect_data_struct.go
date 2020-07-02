/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 效果数据（点击）
type ClickEffectDataStruct struct {
	Count       int64   `json:"count,omitempty"`
	Ratio       float64 `json:"ratio,omitempty"`
	CategoryWin float64 `json:"category_win,omitempty"`
	CategoryAvg float64 `json:"category_avg,omitempty"`
}