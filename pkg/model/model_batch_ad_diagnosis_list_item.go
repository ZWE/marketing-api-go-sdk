/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 单个广告的诊断明细
type BatchAdDiagnosisListItem struct {
	AdgroupId                int64                 `json:"adgroup_id,omitempty"`
	DiagnoseTime             string                `json:"diagnose_time,omitempty"`
	OperateSuggestion        OperateSuggestion     `json:"operate_suggestion,omitempty"`
	LearningStatus           LearningStatus        `json:"learning_status,omitempty"`
	CostGuaranteeStatus      CostGuaranteeStatus   `json:"cost_guarantee_status,omitempty"`
	CostGuaranteeMoney       int64                 `json:"cost_guarantee_money,omitempty"`
	ExposureCompeteScore     string                `json:"exposure_compete_score,omitempty"`
	ExposureRaiseRate        string                `json:"exposure_raise_rate,omitempty"`
	CpaBiasToday             string                `json:"cpa_bias_today,omitempty"`
	CpaBiasOverall           string                `json:"cpa_bias_overall,omitempty"`
	IsOcpx                   *bool                 `json:"is_ocpx,omitempty"`
	OptimizationGoal         string                `json:"optimization_goal,omitempty"`
	DeepOptimizationGoal     string                `json:"deep_optimization_goal,omitempty"`
	ConclusionDescription    string                `json:"conclusion_description,omitempty"`
	HasDiagnoseDetail        *bool                 `json:"has_diagnose_detail,omitempty"`
	OperateSuggestionDesc    string                `json:"operate_suggestion_desc,omitempty"`
	LearningStatusDesc       string                `json:"learning_status_desc,omitempty"`
	ExposureCompeteScoreDesc string                `json:"exposure_compete_score_desc,omitempty"`
	Detail                   *ResponseDetailStruct `json:"detail,omitempty"`
}
