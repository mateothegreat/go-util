package values

// Pick returns the value of the key in the map if it exists. Otherwise, it
// returns the default value.
func Pick[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

// PickHasValue returns the first non-zero value from the provided values.
//
// Arguments:
// - values: A variadic list of values to check.
//
// Returns:
// - The first non-zero value found, or zero value if none found.
func PickHasValue[T any](values ...T) T {
	var zero T
	for _, v := range values {
		if !IsZero(v) {
			return v
		}
	}
	return zero
}

// IsZero checks if a value is the zero value for its type without using reflection.
//
// Arguments:
// - v: The value to check.
//
// Returns:
// - true if the value is the zero value, false otherwise.
func IsZero[T any](v T) bool {
	val := any(v)

	// Handle nil interface
	if val == nil {
		return true
	}

	// Handle common pointer types
	switch ptr := val.(type) {
	case *string:
		return ptr == nil || *ptr == ""
	case *int:
		return ptr == nil || *ptr == 0
	case *int8:
		return ptr == nil || *ptr == 0
	case *int16:
		return ptr == nil || *ptr == 0
	case *int32:
		return ptr == nil || *ptr == 0
	case *int64:
		return ptr == nil || *ptr == 0
	case *uint:
		return ptr == nil || *ptr == 0
	case *uint8:
		return ptr == nil || *ptr == 0
	case *uint16:
		return ptr == nil || *ptr == 0
	case *uint32:
		return ptr == nil || *ptr == 0
	case *uint64:
		return ptr == nil || *ptr == 0
	case *float32:
		return ptr == nil || *ptr == 0
	case *float64:
		return ptr == nil || *ptr == 0
	case *bool:
		return ptr == nil || *ptr == false

	// Handle non-pointer types
	case string:
		return ptr == ""
	case int, int8, int16, int32, int64:
		return ptr == 0
	case uint, uint8, uint16, uint32, uint64:
		return ptr == 0
	case float32, float64:
		return ptr == 0
	case bool:
		return ptr == false

	// For unknown types, compare with zero value using comparable constraint
	default:
		// Create zero value of the original type
		var zero T
		zeroVal := any(zero)

		// Try to compare if both are the same type
		if zeroVal == nil {
			return val == nil
		}

		// For types we can't handle specifically, assume non-nil values are non-zero
		return false
	}
}
