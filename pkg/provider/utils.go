package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func evaluateStringPtr(d *schema.ResourceData, s string) *string {
	var v string
	if _, ok := d.GetOk(s); !ok {
		return nil
	}
	val, _ := d.GetOk(s)
	v = val.(string)
	return &v
}

func evaluateUInt8Ptr(d *schema.ResourceData, s string) *uint8 {
	var v uint8
	if _, ok := d.GetOk(s); !ok {
		return nil
	}
	val, _ := d.GetOk(s)
	v = uint8(val.(int))
	return &v
}

func evaluateIntPtr(d *schema.ResourceData, s string) *int {
	var v int
	if _, ok := d.GetOk(s); !ok {
		return nil
	}
	val, _ := d.GetOk(s)
	v = val.(int)
	return &v
}

func evaluateBoolPtr(d *schema.ResourceData, s string) *bool {
	var v bool
	if _, ok := d.GetOk(s); !ok {
		return nil
	}
	val, _ := d.GetOk(s)
	v = val.(bool)
	return &v
}
