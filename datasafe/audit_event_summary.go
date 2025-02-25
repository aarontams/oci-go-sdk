// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuditEventSummary The resource represents the audit events collected from the target database by Oracle Data Safe.
type AuditEventSummary struct {

	// The OCID of the audit event.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the audit event. The compartment is the same as that of audit profile of the target database resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the target database that was audited.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The name of the target database that was audited.
	TargetName *string `mandatory:"true" json:"targetName"`

	// The type of the target database that was audited. Allowed values are
	//   - DATABASE_CLOUD_SERVICE - Represents Oracle Database Cloud Services.
	//   - AUTONOMOUS_DATABASE - Represents Oracle Autonomous Databases.
	//   - INSTALLED_DATABASE - Represents databases running on-premises or on compute instances.
	DatabaseType AuditEventSummaryDatabaseTypeEnum `mandatory:"true" json:"databaseType"`

	// The time that the audit event occurs in the target database.
	AuditEventTime *common.SDKTime `mandatory:"true" json:"auditEventTime"`

	// The timestamp when this audit event was collected from the target database by Data Safe.
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Indicates whether an alert was raised for this audit event.
	IsAlerted *bool `mandatory:"true" json:"isAlerted"`

	// The name of the database user whose actions were audited.
	DbUserName *string `mandatory:"false" json:"dbUserName"`

	// The class of the target that was audited.
	TargetClass AuditEventSummaryTargetClassEnum `mandatory:"false" json:"targetClass,omitempty"`

	// The name of the operating system user for the database session.
	OsUserName *string `mandatory:"false" json:"osUserName"`

	// The name of the action executed by the user on the target database. For example ALTER, CREATE or DROP.
	Operation *string `mandatory:"false" json:"operation"`

	// Indicates whether the operation was a success or a failure.
	OperationStatus AuditEventSummaryOperationStatusEnum `mandatory:"false" json:"operationStatus,omitempty"`

	// The name of the detail action executed by the user on the target database. For example ALTER SEQUENCE, CREATE TRIGGER or CREATE INDEX.
	EventName *string `mandatory:"false" json:"eventName"`

	// Oracle Error code generated by the action. Zero indicates the action was successful.
	ErrorCode *string `mandatory:"false" json:"errorCode"`

	// The detailed message on why the error occurred.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// The type of the object in the source database affected by the action. For example PL/SQL, SYNONYM or PACKAGE BODY.
	ObjectType *string `mandatory:"false" json:"objectType"`

	// The name of the object affected by the action.
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The schema name of the object affected by the action.
	ObjectOwner *string `mandatory:"false" json:"objectOwner"`

	// The name of the host machine from which the session was spawned.
	ClientHostname *string `mandatory:"false" json:"clientHostname"`

	// The IP address of the host machine from which the session was spawned.
	ClientIp *string `mandatory:"false" json:"clientIp"`

	// The OCID of the audit trail that generated this audit event. To be noted, this field has been deprecated.
	AuditTrailId *string `mandatory:"false" json:"auditTrailId"`

	// The action taken for this audit event.
	ActionTaken *string `mandatory:"false" json:"actionTaken"`

	// The application from which the audit event was generated. For example SQL Plus or SQL Developer.
	ClientProgram *string `mandatory:"false" json:"clientProgram"`

	// The SQL associated with the audit event.
	CommandText *string `mandatory:"false" json:"commandText"`

	// List of bind variables associated with the command text.
	CommandParam *string `mandatory:"false" json:"commandParam"`

	// List of all other attributes of the audit event seperated by a colon other than the one returned in audit record.
	ExtendedEventAttributes *string `mandatory:"false" json:"extendedEventAttributes"`

	// The location of the audit. Currently the value is audit table.
	AuditLocation AuditEventSummaryAuditLocationEnum `mandatory:"false" json:"auditLocation,omitempty"`

	// The operating system terminal of the user session.
	OsTerminal *string `mandatory:"false" json:"osTerminal"`

	// The client identifier in each Oracle session.
	ClientId *string `mandatory:"false" json:"clientId"`

	// Comma-seperated list of audit policies that caused the current audit event.
	AuditPolicies *string `mandatory:"false" json:"auditPolicies"`

	// The type of the auditing.
	AuditType AuditEventSummaryAuditTypeEnum `mandatory:"false" json:"auditType,omitempty"`

	// The secondary id assigned for the peer database registered with Data Safe.
	PeerTargetDatabaseKey *int `mandatory:"false" json:"peerTargetDatabaseKey"`

	// The underlying source of unified audit trail.
	TrailSource AuditTrailSourceEnum `mandatory:"false" json:"trailSource,omitempty"`

	// Unique name of the database associated to the peer target database.
	DatabaseUniqueName *string `mandatory:"false" json:"databaseUniqueName"`

	// Semicolon-seperated list of application context namespace, attribute, value information in (APPCTX_NSPACE,APPCTX_ATTRIBUTE=<value>) format.
	ApplicationContexts *string `mandatory:"false" json:"applicationContexts"`

	// Fine-grained auditing (FGA) policy name that generated this audit record.
	FgaPolicyName *string `mandatory:"false" json:"fgaPolicyName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AuditEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAuditEventSummaryDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetAuditEventSummaryDatabaseTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAuditEventSummaryTargetClassEnum(string(m.TargetClass)); !ok && m.TargetClass != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetClass: %s. Supported values are: %s.", m.TargetClass, strings.Join(GetAuditEventSummaryTargetClassEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditEventSummaryOperationStatusEnum(string(m.OperationStatus)); !ok && m.OperationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationStatus: %s. Supported values are: %s.", m.OperationStatus, strings.Join(GetAuditEventSummaryOperationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditEventSummaryAuditLocationEnum(string(m.AuditLocation)); !ok && m.AuditLocation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditLocation: %s. Supported values are: %s.", m.AuditLocation, strings.Join(GetAuditEventSummaryAuditLocationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditEventSummaryAuditTypeEnum(string(m.AuditType)); !ok && m.AuditType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditType: %s. Supported values are: %s.", m.AuditType, strings.Join(GetAuditEventSummaryAuditTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAuditTrailSourceEnum(string(m.TrailSource)); !ok && m.TrailSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TrailSource: %s. Supported values are: %s.", m.TrailSource, strings.Join(GetAuditTrailSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuditEventSummaryDatabaseTypeEnum Enum with underlying type: string
type AuditEventSummaryDatabaseTypeEnum string

// Set of constants representing the allowable values for AuditEventSummaryDatabaseTypeEnum
const (
	AuditEventSummaryDatabaseTypeDatabaseCloudService AuditEventSummaryDatabaseTypeEnum = "DATABASE_CLOUD_SERVICE"
	AuditEventSummaryDatabaseTypeAutonomousDatabase   AuditEventSummaryDatabaseTypeEnum = "AUTONOMOUS_DATABASE"
	AuditEventSummaryDatabaseTypeInstalledDatabase    AuditEventSummaryDatabaseTypeEnum = "INSTALLED_DATABASE"
)

var mappingAuditEventSummaryDatabaseTypeEnum = map[string]AuditEventSummaryDatabaseTypeEnum{
	"DATABASE_CLOUD_SERVICE": AuditEventSummaryDatabaseTypeDatabaseCloudService,
	"AUTONOMOUS_DATABASE":    AuditEventSummaryDatabaseTypeAutonomousDatabase,
	"INSTALLED_DATABASE":     AuditEventSummaryDatabaseTypeInstalledDatabase,
}

var mappingAuditEventSummaryDatabaseTypeEnumLowerCase = map[string]AuditEventSummaryDatabaseTypeEnum{
	"database_cloud_service": AuditEventSummaryDatabaseTypeDatabaseCloudService,
	"autonomous_database":    AuditEventSummaryDatabaseTypeAutonomousDatabase,
	"installed_database":     AuditEventSummaryDatabaseTypeInstalledDatabase,
}

// GetAuditEventSummaryDatabaseTypeEnumValues Enumerates the set of values for AuditEventSummaryDatabaseTypeEnum
func GetAuditEventSummaryDatabaseTypeEnumValues() []AuditEventSummaryDatabaseTypeEnum {
	values := make([]AuditEventSummaryDatabaseTypeEnum, 0)
	for _, v := range mappingAuditEventSummaryDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventSummaryDatabaseTypeEnumStringValues Enumerates the set of values in String for AuditEventSummaryDatabaseTypeEnum
func GetAuditEventSummaryDatabaseTypeEnumStringValues() []string {
	return []string{
		"DATABASE_CLOUD_SERVICE",
		"AUTONOMOUS_DATABASE",
		"INSTALLED_DATABASE",
	}
}

// GetMappingAuditEventSummaryDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventSummaryDatabaseTypeEnum(val string) (AuditEventSummaryDatabaseTypeEnum, bool) {
	enum, ok := mappingAuditEventSummaryDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditEventSummaryTargetClassEnum Enum with underlying type: string
type AuditEventSummaryTargetClassEnum string

// Set of constants representing the allowable values for AuditEventSummaryTargetClassEnum
const (
	AuditEventSummaryTargetClassDatabase AuditEventSummaryTargetClassEnum = "DATABASE"
)

var mappingAuditEventSummaryTargetClassEnum = map[string]AuditEventSummaryTargetClassEnum{
	"DATABASE": AuditEventSummaryTargetClassDatabase,
}

var mappingAuditEventSummaryTargetClassEnumLowerCase = map[string]AuditEventSummaryTargetClassEnum{
	"database": AuditEventSummaryTargetClassDatabase,
}

// GetAuditEventSummaryTargetClassEnumValues Enumerates the set of values for AuditEventSummaryTargetClassEnum
func GetAuditEventSummaryTargetClassEnumValues() []AuditEventSummaryTargetClassEnum {
	values := make([]AuditEventSummaryTargetClassEnum, 0)
	for _, v := range mappingAuditEventSummaryTargetClassEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventSummaryTargetClassEnumStringValues Enumerates the set of values in String for AuditEventSummaryTargetClassEnum
func GetAuditEventSummaryTargetClassEnumStringValues() []string {
	return []string{
		"DATABASE",
	}
}

// GetMappingAuditEventSummaryTargetClassEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventSummaryTargetClassEnum(val string) (AuditEventSummaryTargetClassEnum, bool) {
	enum, ok := mappingAuditEventSummaryTargetClassEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditEventSummaryOperationStatusEnum Enum with underlying type: string
type AuditEventSummaryOperationStatusEnum string

// Set of constants representing the allowable values for AuditEventSummaryOperationStatusEnum
const (
	AuditEventSummaryOperationStatusSuccess AuditEventSummaryOperationStatusEnum = "SUCCESS"
	AuditEventSummaryOperationStatusFailure AuditEventSummaryOperationStatusEnum = "FAILURE"
)

var mappingAuditEventSummaryOperationStatusEnum = map[string]AuditEventSummaryOperationStatusEnum{
	"SUCCESS": AuditEventSummaryOperationStatusSuccess,
	"FAILURE": AuditEventSummaryOperationStatusFailure,
}

var mappingAuditEventSummaryOperationStatusEnumLowerCase = map[string]AuditEventSummaryOperationStatusEnum{
	"success": AuditEventSummaryOperationStatusSuccess,
	"failure": AuditEventSummaryOperationStatusFailure,
}

// GetAuditEventSummaryOperationStatusEnumValues Enumerates the set of values for AuditEventSummaryOperationStatusEnum
func GetAuditEventSummaryOperationStatusEnumValues() []AuditEventSummaryOperationStatusEnum {
	values := make([]AuditEventSummaryOperationStatusEnum, 0)
	for _, v := range mappingAuditEventSummaryOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventSummaryOperationStatusEnumStringValues Enumerates the set of values in String for AuditEventSummaryOperationStatusEnum
func GetAuditEventSummaryOperationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILURE",
	}
}

// GetMappingAuditEventSummaryOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventSummaryOperationStatusEnum(val string) (AuditEventSummaryOperationStatusEnum, bool) {
	enum, ok := mappingAuditEventSummaryOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditEventSummaryAuditLocationEnum Enum with underlying type: string
type AuditEventSummaryAuditLocationEnum string

// Set of constants representing the allowable values for AuditEventSummaryAuditLocationEnum
const (
	AuditEventSummaryAuditLocationAuditTable AuditEventSummaryAuditLocationEnum = "AUDIT_TABLE"
)

var mappingAuditEventSummaryAuditLocationEnum = map[string]AuditEventSummaryAuditLocationEnum{
	"AUDIT_TABLE": AuditEventSummaryAuditLocationAuditTable,
}

var mappingAuditEventSummaryAuditLocationEnumLowerCase = map[string]AuditEventSummaryAuditLocationEnum{
	"audit_table": AuditEventSummaryAuditLocationAuditTable,
}

// GetAuditEventSummaryAuditLocationEnumValues Enumerates the set of values for AuditEventSummaryAuditLocationEnum
func GetAuditEventSummaryAuditLocationEnumValues() []AuditEventSummaryAuditLocationEnum {
	values := make([]AuditEventSummaryAuditLocationEnum, 0)
	for _, v := range mappingAuditEventSummaryAuditLocationEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventSummaryAuditLocationEnumStringValues Enumerates the set of values in String for AuditEventSummaryAuditLocationEnum
func GetAuditEventSummaryAuditLocationEnumStringValues() []string {
	return []string{
		"AUDIT_TABLE",
	}
}

// GetMappingAuditEventSummaryAuditLocationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventSummaryAuditLocationEnum(val string) (AuditEventSummaryAuditLocationEnum, bool) {
	enum, ok := mappingAuditEventSummaryAuditLocationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AuditEventSummaryAuditTypeEnum Enum with underlying type: string
type AuditEventSummaryAuditTypeEnum string

// Set of constants representing the allowable values for AuditEventSummaryAuditTypeEnum
const (
	AuditEventSummaryAuditTypeStandard      AuditEventSummaryAuditTypeEnum = "STANDARD"
	AuditEventSummaryAuditTypeFineGrained   AuditEventSummaryAuditTypeEnum = "FINE_GRAINED"
	AuditEventSummaryAuditTypeXs            AuditEventSummaryAuditTypeEnum = "XS"
	AuditEventSummaryAuditTypeDatabaseVault AuditEventSummaryAuditTypeEnum = "DATABASE_VAULT"
	AuditEventSummaryAuditTypeLabelSecurity AuditEventSummaryAuditTypeEnum = "LABEL_SECURITY"
	AuditEventSummaryAuditTypeRman          AuditEventSummaryAuditTypeEnum = "RMAN"
	AuditEventSummaryAuditTypeDatapump      AuditEventSummaryAuditTypeEnum = "DATAPUMP"
	AuditEventSummaryAuditTypeDirectPathApi AuditEventSummaryAuditTypeEnum = "DIRECT_PATH_API"
)

var mappingAuditEventSummaryAuditTypeEnum = map[string]AuditEventSummaryAuditTypeEnum{
	"STANDARD":        AuditEventSummaryAuditTypeStandard,
	"FINE_GRAINED":    AuditEventSummaryAuditTypeFineGrained,
	"XS":              AuditEventSummaryAuditTypeXs,
	"DATABASE_VAULT":  AuditEventSummaryAuditTypeDatabaseVault,
	"LABEL_SECURITY":  AuditEventSummaryAuditTypeLabelSecurity,
	"RMAN":            AuditEventSummaryAuditTypeRman,
	"DATAPUMP":        AuditEventSummaryAuditTypeDatapump,
	"DIRECT_PATH_API": AuditEventSummaryAuditTypeDirectPathApi,
}

var mappingAuditEventSummaryAuditTypeEnumLowerCase = map[string]AuditEventSummaryAuditTypeEnum{
	"standard":        AuditEventSummaryAuditTypeStandard,
	"fine_grained":    AuditEventSummaryAuditTypeFineGrained,
	"xs":              AuditEventSummaryAuditTypeXs,
	"database_vault":  AuditEventSummaryAuditTypeDatabaseVault,
	"label_security":  AuditEventSummaryAuditTypeLabelSecurity,
	"rman":            AuditEventSummaryAuditTypeRman,
	"datapump":        AuditEventSummaryAuditTypeDatapump,
	"direct_path_api": AuditEventSummaryAuditTypeDirectPathApi,
}

// GetAuditEventSummaryAuditTypeEnumValues Enumerates the set of values for AuditEventSummaryAuditTypeEnum
func GetAuditEventSummaryAuditTypeEnumValues() []AuditEventSummaryAuditTypeEnum {
	values := make([]AuditEventSummaryAuditTypeEnum, 0)
	for _, v := range mappingAuditEventSummaryAuditTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuditEventSummaryAuditTypeEnumStringValues Enumerates the set of values in String for AuditEventSummaryAuditTypeEnum
func GetAuditEventSummaryAuditTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"FINE_GRAINED",
		"XS",
		"DATABASE_VAULT",
		"LABEL_SECURITY",
		"RMAN",
		"DATAPUMP",
		"DIRECT_PATH_API",
	}
}

// GetMappingAuditEventSummaryAuditTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuditEventSummaryAuditTypeEnum(val string) (AuditEventSummaryAuditTypeEnum, bool) {
	enum, ok := mappingAuditEventSummaryAuditTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
