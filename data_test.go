package appconf

import (
	"testing"
)

func TestAppConf_StringValue_ToInt(t *testing.T) {
	sv := StringValue("123")
	intValue, err := sv.ToInt()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if intValue != 123 {
		t.Errorf("StringValue.ToInt() = %d; expected %d", intValue, 123)
	}
}

func TestAppConf_StringValue_ToFloat64(t *testing.T) {
	sv := StringValue("123.456")
	float64Value, err := sv.ToFloat64()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !almostEqual(float64Value, 123.456) {
		t.Errorf("StringValue.ToFloat64() = %f; expected %f", float64Value, 123.456)
	}
}

func TestAppConf_StringValue_ToBool(t *testing.T) {
	sv := StringValue("true")
	boolValue, err := sv.ToBool()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !boolValue {
		t.Errorf("StringValue.ToBool() = %t; expected %t", boolValue, true)
	}
}

func TestAppConf_IntValue_ToFloat64(t *testing.T) {
	iv := IntValue(456)
	floatValue, err := iv.ToFloat64()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !almostEqual(floatValue, 456.0) {
		t.Errorf("IntValue.ToFloat64() = %f; expected %f", floatValue, 456.0)
	}
}

func TestAppConf_IntValue_ToBool(t *testing.T) {
	iv := IntValue(0)
	boolValue, err := iv.ToBool()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if boolValue {
		t.Errorf("IntValue.ToBool() = %t; expected %t", boolValue, false)
	}
}

func TestAppConf_IntValue_ToString(t *testing.T) {
	iv := IntValue(456)
	stringValue := iv.ToString()
	if stringValue != "456" {
		t.Errorf("IntValue.ToString() = %s; expected %s", stringValue, "456")
	}
}

func TestAppConf_FloatValue_ToBool(t *testing.T) {
	fv := FloatValue(0.0)
	boolValue, err := fv.ToBool()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if boolValue != false {
		t.Errorf("FloatValue.ToBool() = %t; expected %t", boolValue, false)
	}
}

func TestAppConf_FloatValue_ToInt(t *testing.T) {
	fv := FloatValue(456.0)
	intValue, err := fv.ToInt()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if intValue != 456 {
		t.Errorf("FloatValue.ToBool() = %d; expected %d", intValue, 456)
	}
}

func TestAppConf_FloatValue_ToString(t *testing.T) {
	fv := FloatValue(456.1)
	stringValue := fv.ToString()
	if stringValue != "456.1" {
		t.Errorf("FloatValue.ToBool() = %s; expected %s", stringValue, "456.1")
	}
}
