/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// CalculateStatus : 计算状态
type CalculateStatus string

// List of CalculateStatus
const (
	CalculateStatus_NORMAL                 CalculateStatus = "CALCULATE_STATUS_NORMAL"
	CalculateStatus_PENDING                CalculateStatus = "CALCULATE_STATUS_PENDING"
	CalculateStatus_DENIED                 CalculateStatus = "CALCULATE_STATUS_DENIED"
	CalculateStatus_FROZEN                 CalculateStatus = "CALCULATE_STATUS_FROZEN"
	CalculateStatus_SUSPEND                CalculateStatus = "CALCULATE_STATUS_SUSPEND"
	CalculateStatus_READY                  CalculateStatus = "CALCULATE_STATUS_READY"
	CalculateStatus_ACTIVE                 CalculateStatus = "CALCULATE_STATUS_ACTIVE"
	CalculateStatus_STOP                   CalculateStatus = "CALCULATE_STATUS_STOP"
	CalculateStatus_NOT_READY_IMG          CalculateStatus = "CALCULATE_STATUS_NOT_READY_IMG"
	CalculateStatus_DELETED                CalculateStatus = "CALCULATE_STATUS_DELETED"
	CalculateStatus_NOT_READY_ACCT_DISABLE CalculateStatus = "CALCULATE_STATUS_NOT_READY_ACCT_DISABLE"
	CalculateStatus_STOP_ACCT_NO_FUND      CalculateStatus = "CALCULATE_STATUS_STOP_ACCT_NO_FUND"
	CalculateStatus_STOP_ACCT_BUDGET       CalculateStatus = "CALCULATE_STATUS_STOP_ACCT_BUDGET"
	CalculateStatus_STOP_CAMP_BUDGET       CalculateStatus = "CALCULATE_STATUS_STOP_CAMP_BUDGET"
	CalculateStatus_STOP_CAMP_PAUSE        CalculateStatus = "CALCULATE_STATUS_STOP_CAMP_PAUSE"
	CalculateStatus_ACTIVE_CAMP            CalculateStatus = "CALCULATE_STATUS_ACTIVE_CAMP"
	CalculateStatus_READY_CAMP             CalculateStatus = "CALCULATE_STATUS_READY_CAMP"
	CalculateStatus_ACTIVE_ACCP            CalculateStatus = "CALCULATE_STATUS_ACTIVE_ACCP"
	CalculateStatus_READY_ACCP             CalculateStatus = "CALCULATE_STATUS_READY_ACCP"
	CalculateStatus_SPONSORLIMIT           CalculateStatus = "CALCULATE_STATUS_SPONSORLIMIT"
	CalculateStatus_ACTIVE_ACC_FROZEN      CalculateStatus = "CALCULATE_STATUS_ACTIVE_ACC_FROZEN"
	CalculateStatus_ACTIVE_ACC_EMPTY       CalculateStatus = "CALCULATE_STATUS_ACTIVE_ACC_EMPTY"
	CalculateStatus_ACTIVE_ACC_LIMIT       CalculateStatus = "CALCULATE_STATUS_ACTIVE_ACC_LIMIT"
	CalculateStatus_ACTIVE_CAM_LIMIT       CalculateStatus = "CALCULATE_STATUS_ACTIVE_CAM_LIMIT"
	CalculateStatus_ACTIVE_CAM_PAUSED      CalculateStatus = "CALCULATE_STATUS_ACTIVE_CAM_PAUSED"
	CalculateStatus_PART_ENABLE            CalculateStatus = "CALCULATE_STATUS_PART_ENABLE"
	CalculateStatus_UNAUDIT_RE             CalculateStatus = "CALCULATE_STATUS_UNAUDIT_RE"
	CalculateStatus_PART_READY             CalculateStatus = "CALCULATE_STATUS_PART_READY"
	CalculateStatus_PART_ACTIVE            CalculateStatus = "CALCULATE_STATUS_PART_ACTIVE"
	CalculateStatus_PART_PREPARE           CalculateStatus = "CALCULATE_STATUS_PART_PREPARE"
	CalculateStatus_PART_INVALID           CalculateStatus = "CALCULATE_STATUS_PART_INVALID"
	CalculateStatus_AD_PARTIAL_NORMAL      CalculateStatus = "CALCULATE_STATUS_AD_PARTIAL_NORMAL"
	CalculateStatus_AD_PARTIAL_PENDING     CalculateStatus = "CALCULATE_STATUS_AD_PARTIAL_PENDING"
)