package cloudflare

import (
	"reflect"
	"testing"
)

func testResourceCloudflareAccessPolicyStateDataV0() map[string]interface{} {
	return map[string]interface{}{
		"allocated_storage": 10,
		"engine":            "mariadb",
		"identifier":        "my-test-instance",
		"instance_class":    "db.t2.micro",
		"password":          "avoid-plaintext-passwords",
		"username":          "tfacctest",
		"tags":              map[string]interface{}{"key1": "value1"},
	}
}

func testResourceCloudflareAccessPolicyStateDataV1() map[string]interface{} {
	v0 := testResourceCloudflareAccessPolicyStateDataV0()
	return map[string]interface{}{
		"allocated_storage":        v0["allocated_storage"],
		"delete_automated_backups": true,
		"engine":                   v0["engine"],
		"identifier":               v0["identifier"],
		"instance_class":           v0["instance_class"],
		"password":                 v0["password"],
		"username":                 v0["username"],
		"tags":                     v0["tags"],
	}
}

func TestResourceCloudflareAccessPolicyStateUpgradeV0(t *testing.T) {
	expected := testResourceCloudflareAccessPolicyStateDataV1()
	actual, err := resourceCloudflareAccessPolicyStateUpgradeV0(testResourceCloudflareAccessPolicyStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
