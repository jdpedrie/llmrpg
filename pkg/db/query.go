package db

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/geldata/gel-go"
	"github.com/geldata/gel-go/geltypes"
)

type Model interface {
	DBType() string
}

// Insert builds and executes the query dynamically using reflection.
// The given `model` is updated with the result of the operation.
// Note that nested object IDs may not be populated.
func Insert(ctx context.Context, client *gel.Client, model Model) error {
	query, args, err := buildInsertQuery(model, 0)
	if err != nil {
		return err
	}

	return client.QuerySingle(ctx, fmt.Sprintf(
		`with ins := (%s) select ins`, query,
	), model, args...)
}

// Recursive function to build the query and arguments list.
// 🤮
func buildInsertQuery(model Model, argIndex int) (string, []any, error) {
	modelType := model.DBType()
	fieldValues, err := extractFields(model)
	if err != nil {
		return "", nil, err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("INSERT %s {", modelType))

	var queryParts []string
	var args []any

	for fieldName, fieldValue := range fieldValues {
		var qp string
		switch v := fieldValue.(type) {
		case string, int16, int32, int64, bool, float64, nil: // Scalars & Optional Fields
			var cast string
			switch fieldValue.(type) {
			case string:
				cast = "<str>"
			case bool:
				cast = "<bool>"
			case int16:
				cast = "<int16>"
			case int32:
				cast = "<int32>"
			case int64:
				cast = "<int64>"
			case float32:
				cast = "<float32>"
			case float64:
				cast = "<float64>"
			}

			qp = fmt.Sprintf("%s := %s$%d", fieldName, cast, argIndex)
			args = append(args, v)
			queryParts = append(queryParts, qp)
			argIndex++

		case []string, []int16, []int32, []int64, []bool, []float64:
			var cast string
			switch v.(type) {
			case []string:
				cast = "<array<str>>"
			case []bool:
				cast = "<array<bool>>"
			case []int16:
				cast = "<array<int16>>"
			case []int32:
				cast = "<array<int32>>"
			case []int64:
				cast = "<array<int64>>"
			case []float64:
				cast = "<array<float64>>"
			}

			qp = fmt.Sprintf("%s := %s$%d", fieldName, cast, argIndex)
			queryParts = append(queryParts, qp)
			args = append(args, v) // Pass the array as a single argument
			argIndex++

		case Model: // Nested single entity
			subQuery, subArgs, err := buildInsertQuery(v, argIndex)
			if err != nil {
				return "", nil, err
			}
			qp = fmt.Sprintf("%s := (%s)", fieldName, subQuery)
			queryParts = append(queryParts, qp)
			args = append(args, subArgs...)
			argIndex += len(subArgs)

		case []Model: // Nested list of entities
			nestedQueries := []string{}
			nestedArgs := []any{}
			for _, nested := range v {
				subQuery, subArgs, err := buildInsertQuery(nested, argIndex)
				if err != nil {
					return "", nil, err
				}
				nestedQueries = append(nestedQueries, "("+subQuery+")")
				nestedArgs = append(nestedArgs, subArgs...)
				argIndex += len(subArgs)
			}
			qp = fmt.Sprintf("%s := {%s}", fieldName, strings.Join(nestedQueries, ", "))
			queryParts = append(queryParts, qp)
			args = append(args, nestedArgs...) // Properly append all nested args

		default:
			if cast, ok := isGelType(v); ok {
				if cast != "" {
					cast = fmt.Sprintf("<%s>", cast)
				}
				qp = fmt.Sprintf("%s := %s$%d", fieldName, cast, argIndex)
				args = append(args, v)
				queryParts = append(queryParts, qp)
				argIndex++
			} else if m, ok := v.(Model); ok {
				subQuery, subArgs, err := buildInsertQuery(m, argIndex)
				if err != nil {
					return "", nil, err
				}
				qp = fmt.Sprintf("%s := (%s)", fieldName, subQuery)
				queryParts = append(queryParts, qp)
				args = append(args, subArgs...)
				argIndex += len(subArgs)
			} else if arr, ok := v.([]any); ok {
				if _, ok := arr[0].(Model); ok && len(arr) > 0 {
					nestedQueries := []string{}
					nestedArgs := []any{}
					for _, nested := range arr {
						subQuery, subArgs, err := buildInsertQuery(nested.(Model), argIndex)
						if err != nil {
							return "", nil, err
						}
						nestedQueries = append(nestedQueries, "("+subQuery+")")
						nestedArgs = append(nestedArgs, subArgs...)
						argIndex += len(subArgs)
					}
					qp = fmt.Sprintf("%s := {%s}", fieldName, strings.Join(nestedQueries, ", "))
					queryParts = append(queryParts, qp)
					args = append(args, nestedArgs...) // Properly append all nested args
				}
			} else {
				return "", nil, fmt.Errorf("unsupported field type: %v", reflect.TypeOf(v))
			}
		}
	}

	sb.WriteString(strings.Join(queryParts, ", "))
	sb.WriteString(" }")

	return sb.String(), args, nil
}

// extractFields - uses reflection to get field values and field names based
// on the `gel` struct tag.
func extractFields(model Model) (map[string]any, error) {
	result := map[string]any{}
	value := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)

	// Ensure we have a struct
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, errors.New("Insert() only supports struct types")
	}

	// Iterate over struct fields
	for i := range typ.NumField() {
		field := typ.Field(i)
		tag := field.Tag.Get("gel")
		if tag == "" {
			tag = strings.ToLower(field.Name) // Default to lowercase field name
		}

		if tag == "id" {
			continue
		}

		fieldValue := value.Field(i)
		fieldType := field.Type

		// Handle nil or uninitialized values
		if !fieldValue.IsValid() || (fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil()) {
			result[tag] = nil
			continue
		}

		// Detect nested Model types
		if fieldType.Kind() == reflect.Struct {
			if nestedModel, ok := fieldValue.Interface().(Model); ok {
				result[tag] = nestedModel // Single nested Model
				continue
			}
		}

		// Detect []Model slices OR arrays of scalar types
		if fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array {
			slice := fieldValue
			modelSlice := []Model{}
			for j := range slice.Len() {
				if item, ok := slice.Index(j).Interface().(Model); ok {
					modelSlice = append(modelSlice, item)
				}
			}

			if len(modelSlice) > 0 {
				result[tag] = modelSlice // Store as Model slice
			} else {
				result[tag] = slice.Interface()
			}
			continue
		}

		// Otherwise, store the field as a normal value
		result[tag] = fieldValue.Interface()
	}

	return result, nil
}

func UUIDEmpty(v geltypes.UUID) bool {
	return v.String() == "00000000-0000-0000-0000-000000000000"
}

// isGelType checks if v is one of the allowed Gel types
func isGelType(v any) (string, bool) {
	vType := reflect.TypeOf(v)

	// Direct type match
	if cast, ok := gelTypeMap[vType]; ok {
		return cast, true
	}

	// Handle pointers by dereferencing them
	if vType != nil && vType.Kind() == reflect.Ptr {
		if cast, ok := gelTypeMap[vType.Elem()]; ok {
			return cast, true
		}
	}

	return "", false
}

var gelTypeMap = map[reflect.Type]string{
	reflect.TypeOf(geltypes.DateDuration{}):               "date_duration",
	reflect.TypeOf(geltypes.Duration(0)):                  "duration",
	reflect.TypeOf(geltypes.LocalDate{}):                  "local_date",
	reflect.TypeOf(geltypes.LocalDateTime{}):              "local_datetime",
	reflect.TypeOf(geltypes.LocalTime{}):                  "local_time",
	reflect.TypeOf(geltypes.Memory(0)):                    "",
	reflect.TypeOf(geltypes.MultiRangeDateTime{}):         "",
	reflect.TypeOf(geltypes.MultiRangeFloat32{}):          "",
	reflect.TypeOf(geltypes.MultiRangeFloat64{}):          "",
	reflect.TypeOf(geltypes.MultiRangeInt32{}):            "",
	reflect.TypeOf(geltypes.MultiRangeInt64{}):            "",
	reflect.TypeOf(geltypes.MultiRangeLocalDate{}):        "",
	reflect.TypeOf(geltypes.MultiRangeLocalDateTime{}):    "",
	reflect.TypeOf(geltypes.Optional{}):                   "",
	reflect.TypeOf(geltypes.OptionalBigInt{}):             "OPTIONAL bigint",
	reflect.TypeOf(geltypes.OptionalBool{}):               "OPTIONAL bool",
	reflect.TypeOf(geltypes.OptionalBytes{}):              "OPTIONAL bytes",
	reflect.TypeOf(geltypes.OptionalDateDuration{}):       "OPTIONAL date_duration",
	reflect.TypeOf(geltypes.OptionalDateTime{}):           "OPTIONAL datetime",
	reflect.TypeOf(geltypes.OptionalDuration{}):           "OPTIONAL duration",
	reflect.TypeOf(geltypes.OptionalFloat32{}):            "OPTIONAL float32",
	reflect.TypeOf(geltypes.OptionalFloat64{}):            "OPTIONAL float64",
	reflect.TypeOf(geltypes.OptionalInt16{}):              "OPTIONAL int16",
	reflect.TypeOf(geltypes.OptionalInt32{}):              "OPTIONAL int32",
	reflect.TypeOf(geltypes.OptionalInt64{}):              "OPTIONAL int64",
	reflect.TypeOf(geltypes.OptionalLocalDate{}):          "OPTIONAL local_date",
	reflect.TypeOf(geltypes.OptionalLocalDateTime{}):      "OPTIONAL local_datetime",
	reflect.TypeOf(geltypes.OptionalLocalTime{}):          "OPTIONAL local_time",
	reflect.TypeOf(geltypes.OptionalMemory{}):             "",
	reflect.TypeOf(geltypes.OptionalRangeDateTime{}):      "",
	reflect.TypeOf(geltypes.OptionalRangeFloat32{}):       "",
	reflect.TypeOf(geltypes.OptionalRangeFloat64{}):       "",
	reflect.TypeOf(geltypes.OptionalRangeInt32{}):         "",
	reflect.TypeOf(geltypes.OptionalRangeInt64{}):         "",
	reflect.TypeOf(geltypes.OptionalRangeLocalDate{}):     "",
	reflect.TypeOf(geltypes.OptionalRangeLocalDateTime{}): "",
	reflect.TypeOf(geltypes.OptionalRelativeDuration{}):   "",
	reflect.TypeOf(geltypes.OptionalStr{}):                "OPTIONAL str",
	reflect.TypeOf(geltypes.OptionalUUID{}):               "OPTIONAL uuid",
	reflect.TypeOf(geltypes.RangeDateTime{}):              "",
	reflect.TypeOf(geltypes.RangeFloat32{}):               "",
	reflect.TypeOf(geltypes.RangeFloat64{}):               "",
	reflect.TypeOf(geltypes.RangeInt32{}):                 "",
	reflect.TypeOf(geltypes.RangeInt64{}):                 "",
	reflect.TypeOf(geltypes.RangeLocalDate{}):             "",
	reflect.TypeOf(geltypes.RangeLocalDateTime{}):         "",
	reflect.TypeOf(geltypes.RelativeDuration{}):           "relative_duration",
	reflect.TypeOf(geltypes.UUID{}):                       "uuid",
}
