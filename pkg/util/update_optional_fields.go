package util

import (
	"reflect"
)

import (
	"google.golang.org/protobuf/proto"
)

func UpdateOptionalFields(protoReq proto.Message, m map[string]any) {
	v := reflect.ValueOf(protoReq).Elem()
	idIndex := 0
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name == "Id" {
			idIndex = i
			fv := v.FieldByName("Id")
			m["Id"] = ""
			if id, ok := fv.Interface().(string); ok {
				m["Id"] = id
			}
			continue
		}
		if idIndex > 0 {
			field := v.Type().Field(i)
			fv := v.FieldByName(field.Name)
			if !fv.IsNil() {
				switch fv := fv.Interface().(type) {
				case *string:
					m[field.Name] = *fv
				case *bool:
					m[field.Name] = *fv
				case *uint32:
					m[field.Name] = *fv
				case *uint64:
					m[field.Name] = *fv
				case *int32:
					m[field.Name] = *fv
				case *int64:
					m[field.Name] = *fv
				case *float32:
					m[field.Name] = *fv
				case *float64:
					m[field.Name] = *fv
				}
			}
		}
	}
}
