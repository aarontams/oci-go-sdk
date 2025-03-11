// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ModelVersionSetCategoryEnum Enum with underlying type: string
type ModelVersionSetCategoryEnum string

// Set of constants representing the allowable values for ModelVersionSetCategoryEnum
const (
	ModelVersionSetCategoryUser    ModelVersionSetCategoryEnum = "USER"
	ModelVersionSetCategoryService ModelVersionSetCategoryEnum = "SERVICE"
)

var mappingModelVersionSetCategoryEnum = map[string]ModelVersionSetCategoryEnum{
	"USER":    ModelVersionSetCategoryUser,
	"SERVICE": ModelVersionSetCategoryService,
}

var mappingModelVersionSetCategoryEnumLowerCase = map[string]ModelVersionSetCategoryEnum{
	"user":    ModelVersionSetCategoryUser,
	"service": ModelVersionSetCategoryService,
}

// GetModelVersionSetCategoryEnumValues Enumerates the set of values for ModelVersionSetCategoryEnum
func GetModelVersionSetCategoryEnumValues() []ModelVersionSetCategoryEnum {
	values := make([]ModelVersionSetCategoryEnum, 0)
	for _, v := range mappingModelVersionSetCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetModelVersionSetCategoryEnumStringValues Enumerates the set of values in String for ModelVersionSetCategoryEnum
func GetModelVersionSetCategoryEnumStringValues() []string {
	return []string{
		"USER",
		"SERVICE",
	}
}

// GetMappingModelVersionSetCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelVersionSetCategoryEnum(val string) (ModelVersionSetCategoryEnum, bool) {
	enum, ok := mappingModelVersionSetCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
