// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fulfillment

const host = "https://marketplaceapi.microsoft.com/api"

// APIVersion - The request must send the following parameters as a URL Encoded form; granttype - clientcredentials; resource
// - 20e940b3-4c77-4b0b-9a53-9e16a1b010a7; clientid - AAD Registered App Client ID; client
// secret - AAD Registered App Client Secret
type APIVersion string

const (
	APIVersionTwoThousandEighteen0831 APIVersion = "2018-08-31"
)

// PossibleAPIVersionValues returns the possible values for the APIVersion const type.
func PossibleAPIVersionValues() []APIVersion {
	return []APIVersion{	
		APIVersionTwoThousandEighteen0831,
	}
}

type AllowedCustomerOperationsEnum string

const (
	AllowedCustomerOperationsEnumDelete AllowedCustomerOperationsEnum = "Delete"
	AllowedCustomerOperationsEnumRead AllowedCustomerOperationsEnum = "Read"
	AllowedCustomerOperationsEnumUpdate AllowedCustomerOperationsEnum = "Update"
)

// PossibleAllowedCustomerOperationsEnumValues returns the possible values for the AllowedCustomerOperationsEnum const type.
func PossibleAllowedCustomerOperationsEnumValues() []AllowedCustomerOperationsEnum {
	return []AllowedCustomerOperationsEnum{	
		AllowedCustomerOperationsEnumDelete,
		AllowedCustomerOperationsEnumRead,
		AllowedCustomerOperationsEnumUpdate,
	}
}

type OperationActionEnum string

const (
	OperationActionEnumChangePlan OperationActionEnum = "ChangePlan"
	OperationActionEnumChangeQuantity OperationActionEnum = "ChangeQuantity"
	OperationActionEnumReinstate OperationActionEnum = "Reinstate"
	OperationActionEnumRenew OperationActionEnum = "Renew"
	OperationActionEnumSuspend OperationActionEnum = "Suspend"
	OperationActionEnumUnsubscribe OperationActionEnum = "Unsubscribe"
)

// PossibleOperationActionEnumValues returns the possible values for the OperationActionEnum const type.
func PossibleOperationActionEnumValues() []OperationActionEnum {
	return []OperationActionEnum{	
		OperationActionEnumChangePlan,
		OperationActionEnumChangeQuantity,
		OperationActionEnumReinstate,
		OperationActionEnumRenew,
		OperationActionEnumSuspend,
		OperationActionEnumUnsubscribe,
	}
}

type OperationStatusEnum string

const (
	OperationStatusEnumConflict OperationStatusEnum = "Conflict"
	OperationStatusEnumFailed OperationStatusEnum = "Failed"
	OperationStatusEnumInProgress OperationStatusEnum = "InProgress"
	OperationStatusEnumNotStarted OperationStatusEnum = "NotStarted"
	OperationStatusEnumSucceeded OperationStatusEnum = "Succeeded"
)

// PossibleOperationStatusEnumValues returns the possible values for the OperationStatusEnum const type.
func PossibleOperationStatusEnumValues() []OperationStatusEnum {
	return []OperationStatusEnum{	
		OperationStatusEnumConflict,
		OperationStatusEnumFailed,
		OperationStatusEnumInProgress,
		OperationStatusEnumNotStarted,
		OperationStatusEnumSucceeded,
	}
}

// SandboxTypeEnum - Possible Values are None, Csp (Csp sandbox purchase)
type SandboxTypeEnum string

const (
	SandboxTypeEnumCsp SandboxTypeEnum = "Csp"
	SandboxTypeEnumNone SandboxTypeEnum = "None"
)

// PossibleSandboxTypeEnumValues returns the possible values for the SandboxTypeEnum const type.
func PossibleSandboxTypeEnumValues() []SandboxTypeEnum {
	return []SandboxTypeEnum{	
		SandboxTypeEnumCsp,
		SandboxTypeEnumNone,
	}
}

// SessionModeEnum - Dry Run indicates all transactions run as Test-Mode in the commerce stack
type SessionModeEnum string

const (
	SessionModeEnumDryRun SessionModeEnum = "DryRun"
	SessionModeEnumNone SessionModeEnum = "None"
)

// PossibleSessionModeEnumValues returns the possible values for the SessionModeEnum const type.
func PossibleSessionModeEnumValues() []SessionModeEnum {
	return []SessionModeEnum{	
		SessionModeEnumDryRun,
		SessionModeEnumNone,
	}
}

// SubscriptionStatusEnum - Indicates the status of the operation.
type SubscriptionStatusEnum string

const (
	SubscriptionStatusEnumNotStarted SubscriptionStatusEnum = "NotStarted"
	SubscriptionStatusEnumPendingFulfillmentStart SubscriptionStatusEnum = "PendingFulfillmentStart"
	SubscriptionStatusEnumSubscribed SubscriptionStatusEnum = "Subscribed"
	SubscriptionStatusEnumSuspended SubscriptionStatusEnum = "Suspended"
	SubscriptionStatusEnumUnsubscribed SubscriptionStatusEnum = "Unsubscribed"
)

// PossibleSubscriptionStatusEnumValues returns the possible values for the SubscriptionStatusEnum const type.
func PossibleSubscriptionStatusEnumValues() []SubscriptionStatusEnum {
	return []SubscriptionStatusEnum{	
		SubscriptionStatusEnumNotStarted,
		SubscriptionStatusEnumPendingFulfillmentStart,
		SubscriptionStatusEnumSubscribed,
		SubscriptionStatusEnumSuspended,
		SubscriptionStatusEnumUnsubscribed,
	}
}

type TermUnitEnum string

const (
	TermUnitEnumP1M TermUnitEnum = "P1M"
	TermUnitEnumP1Y TermUnitEnum = "P1Y"
	TermUnitEnumP2Y TermUnitEnum = "P2Y"
	TermUnitEnumP3Y TermUnitEnum = "P3Y"
)

// PossibleTermUnitEnumValues returns the possible values for the TermUnitEnum const type.
func PossibleTermUnitEnumValues() []TermUnitEnum {
	return []TermUnitEnum{	
		TermUnitEnumP1M,
		TermUnitEnumP1Y,
		TermUnitEnumP2Y,
		TermUnitEnumP3Y,
	}
}

type UpdateOperationStatusEnum string

const (
	UpdateOperationStatusEnumFailure UpdateOperationStatusEnum = "Failure"
	UpdateOperationStatusEnumSuccess UpdateOperationStatusEnum = "Success"
)

// PossibleUpdateOperationStatusEnumValues returns the possible values for the UpdateOperationStatusEnum const type.
func PossibleUpdateOperationStatusEnumValues() []UpdateOperationStatusEnum {
	return []UpdateOperationStatusEnum{	
		UpdateOperationStatusEnumFailure,
		UpdateOperationStatusEnumSuccess,
	}
}

