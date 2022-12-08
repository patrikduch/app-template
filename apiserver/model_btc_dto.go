/*
 * App template API
 *
 * API to access and configure the app template
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apiserver

// BtcDto - Simple BTC exchange rate object
type BtcDto struct {

	// A id identifying the example configuration
	Id *int64 `json:"id,omitempty"`
}

// AssertBtcDtoRequired checks if the required fields are not zero-ed
func AssertBtcDtoRequired(obj BtcDto) error {
	return nil
}

// AssertRecurseBtcDtoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of BtcDto (e.g. [][]BtcDto), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBtcDtoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBtcDto, ok := obj.(BtcDto)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBtcDtoRequired(aBtcDto)
	})
}
