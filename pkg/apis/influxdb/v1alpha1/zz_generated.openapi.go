// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/influxdb/v1alpha1.InfluxDBAggregator":       schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregator(ref),
		"./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorSpec":   schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregatorSpec(ref),
		"./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorStatus": schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregatorStatus(ref),
	}
}

func schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregator(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfluxDBAggregator is the Schema for the influxdbaggregators API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorSpec", "./pkg/apis/influxdb/v1alpha1.InfluxDBAggregatorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregatorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfluxDBAggregatorSpec defines the desired state of InfluxDBAggregator",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_influxdb_v1alpha1_InfluxDBAggregatorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InfluxDBAggregatorStatus defines the observed state of InfluxDBAggregator",
				Type:        []string{"object"},
			},
		},
	}
}