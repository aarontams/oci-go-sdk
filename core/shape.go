// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// Shape. A compute instance shape that can be used in LaunchInstance.
// For more information, see [Overview of the Compute Service]({{DOC_SERVER_URL}}/Content/Compute/Concepts/computeoverview.htm).
type Shape struct {

	// The name of the shape. You can enumerate all available shapes by calling
	// ListShapes.
	Shape *string `mandatory:"true" json:"shape,omitempty"`
}

func (model Shape) String() string {
	return common.PointerString(model)
}