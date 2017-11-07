// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateSaml2IdentityProviderDetails struct {

	// The OCID of your tenancy.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	ProductType CreateSaml2IdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType,omitempty"`

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol CreateSaml2IdentityProviderDetailsProtocolEnum `mandatory:"true" json:"protocol,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataURL *string `mandatory:"true" json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata *string `mandatory:"true" json:"metadata,omitempty"`
}

func (model CreateSaml2IdentityProviderDetails) String() string {
	return common.PointerString(model)
}

type CreateSaml2IdentityProviderDetailsProductTypeEnum string

const (
	CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_IDCS    CreateSaml2IdentityProviderDetailsProductTypeEnum = "IDCS"
	CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS    CreateSaml2IdentityProviderDetailsProductTypeEnum = "ADFS"
	CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN CreateSaml2IdentityProviderDetailsProductTypeEnum = "UNKNOWN"
)

var mapping_createsaml2identityproviderdetails_productType = map[string]CreateSaml2IdentityProviderDetailsProductTypeEnum{
	"IDCS":    CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_IDCS,
	"ADFS":    CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS,
	"UNKNOWN": CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN,
}

func GetCreateSaml2IdentityProviderDetailsProductTypeEnumValues() []CreateSaml2IdentityProviderDetailsProductTypeEnum {
	values := make([]CreateSaml2IdentityProviderDetailsProductTypeEnum, 0)
	for _, v := range mapping_createsaml2identityproviderdetails_productType {
		if v != CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type CreateSaml2IdentityProviderDetailsProtocolEnum string

const (
	CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2   CreateSaml2IdentityProviderDetailsProtocolEnum = "SAML2"
	CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN CreateSaml2IdentityProviderDetailsProtocolEnum = "UNKNOWN"
)

var mapping_createsaml2identityproviderdetails_protocol = map[string]CreateSaml2IdentityProviderDetailsProtocolEnum{
	"SAML2":   CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2,
	"UNKNOWN": CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN,
}

func GetCreateSaml2IdentityProviderDetailsProtocolEnumValues() []CreateSaml2IdentityProviderDetailsProtocolEnum {
	values := make([]CreateSaml2IdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mapping_createsaml2identityproviderdetails_protocol {
		if v != CREATE_SAML2_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
