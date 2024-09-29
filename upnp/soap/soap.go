package soap

import (
    "github.com/huin/goupnp/soap"
)

func MarshalUi2(v uint16) (string, error) {
    return soap.MarshalUi2(v)
}

func UnmarshalUi2(s string) (uint16, error) {
    if s == "" {
        return 0, nil
    }
    return soap.UnmarshalUi2(s)
}

func MarshalUi4(v uint32) (string, error) {
    return soap.MarshalUi4(v)
}

func UnmarshalUi4(s string) (uint32, error) {
    if s == "" {
        return 0, nil
    }
    return soap.UnmarshalUi4(s)
}

func MarshalI4(v int32) (string, error) {
    return soap.MarshalI4(v)
}

func UnmarshalI4(s string) (int32, error) {
    if s == "" {
        return 0, nil
    }
    return soap.UnmarshalI4(s)
}

func MarshalString(v string) (string, error) {
    return soap.MarshalString(v)
}

func UnmarshalString(v string) (string, error) {
    return soap.UnmarshalString(v)
}

func MarshalBoolean(v bool) (string, error) {
    return soap.MarshalBoolean(v)
}

func UnmarshalBoolean(s string) (bool, error) {
    if s == "" {
        return false, nil
    }
    return soap.UnmarshalBoolean(s)
}
