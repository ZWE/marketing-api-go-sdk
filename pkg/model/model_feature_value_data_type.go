/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// FeatureValueDataType : 特征值数据类型
type FeatureValueDataType string

// List of FeatureValueDataType
const (
	FeatureValueDataType_CATEGORICAL          FeatureValueDataType = "CATEGORICAL"
	FeatureValueDataType_DISCRETE_NUMERICAL   FeatureValueDataType = "DISCRETE_NUMERICAL"
	FeatureValueDataType_CONTINUOUS_NUMERICAL FeatureValueDataType = "CONTINUOUS_NUMERICAL"
)