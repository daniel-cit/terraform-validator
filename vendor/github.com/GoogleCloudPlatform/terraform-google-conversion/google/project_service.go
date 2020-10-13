package google

import (
	"fmt"
	"reflect"
)

func GetServiceUsageCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//serviceusage.googleapis.com/projects/{{project}}/services/{{service}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetServiceUsageApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "serviceusage.googleapis.com/Service",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/serviceusage/v1/rest",
				DiscoveryName:        "Service",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetServiceUsageApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentProp, err := expandServiceUsageParent(d.Get("project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("project"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	serviceProp, err := expandServiceUsageService(d.Get("service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service"); !isEmptyValue(reflect.ValueOf(serviceProp)) && (ok || !reflect.DeepEqual(v, serviceProp)) {
		obj["name"] = serviceProp
	}

	obj["state"] = "ENABLED"

	return obj, nil
}

func expandServiceUsageParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		// It does not try to construct anything from empty.
		return "", nil
	}
	// Ideally we should use project_number, but since that is generated server-side,
	// we substitute project_id.
	return fmt.Sprintf("projects/%s", v.(string)), nil
}

func expandServiceUsageService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
