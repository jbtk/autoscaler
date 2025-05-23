// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"strings"
)

// NodeSourceTypeEnum Enum with underlying type: string
type NodeSourceTypeEnum string

// Set of constants representing the allowable values for NodeSourceTypeEnum
const (
	NodeSourceTypeImage NodeSourceTypeEnum = "IMAGE"
)

var mappingNodeSourceTypeEnum = map[string]NodeSourceTypeEnum{
	"IMAGE": NodeSourceTypeImage,
}

var mappingNodeSourceTypeEnumLowerCase = map[string]NodeSourceTypeEnum{
	"image": NodeSourceTypeImage,
}

// GetNodeSourceTypeEnumValues Enumerates the set of values for NodeSourceTypeEnum
func GetNodeSourceTypeEnumValues() []NodeSourceTypeEnum {
	values := make([]NodeSourceTypeEnum, 0)
	for _, v := range mappingNodeSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeSourceTypeEnumStringValues Enumerates the set of values in String for NodeSourceTypeEnum
func GetNodeSourceTypeEnumStringValues() []string {
	return []string{
		"IMAGE",
	}
}

// GetMappingNodeSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeSourceTypeEnum(val string) (NodeSourceTypeEnum, bool) {
	enum, ok := mappingNodeSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
