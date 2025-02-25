// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateTextResult The generated text result to return.
type GenerateTextResult struct {

	// The OCID of the model used in this inference request.
	ModelId *string `mandatory:"true" json:"modelId"`

	// The version of the model.
	ModelVersion *string `mandatory:"true" json:"modelVersion"`

	InferenceResponse LlmInferenceResponse `mandatory:"true" json:"inferenceResponse"`
}

func (m GenerateTextResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateTextResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *GenerateTextResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelId           *string              `json:"modelId"`
		ModelVersion      *string              `json:"modelVersion"`
		InferenceResponse llminferenceresponse `json:"inferenceResponse"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelId = model.ModelId

	m.ModelVersion = model.ModelVersion

	nn, e = model.InferenceResponse.UnmarshalPolymorphicJSON(model.InferenceResponse.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.InferenceResponse = nn.(LlmInferenceResponse)
	} else {
		m.InferenceResponse = nil
	}

	return
}
