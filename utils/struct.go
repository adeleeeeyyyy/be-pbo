package utils

import "reflect"

func StructToMap(data interface{}) map[string]interface{} {
    result := make(map[string]interface{})

    v := reflect.ValueOf(data)
    t := reflect.TypeOf(data)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }

    for i := 0; i < v.NumField(); i++ {
        fieldValue := v.Field(i)
        fieldType := t.Field(i)

        formTag := fieldType.Tag.Get("form")
        if formTag == "" {
            continue
        }

        if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
            result[formTag] = fieldValue.Elem().Interface()
        }
    }

    return result
}
