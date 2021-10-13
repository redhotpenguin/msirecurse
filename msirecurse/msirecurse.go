package msirecurse

import (
	"reflect"
)

// This package was designed to search json objects for certain key value
// combinations. This function takes a map string empty interface,
// and checks the map[string]interface{} to see if the first piece is embedded
// Currently this is only tested with a single key for the needle.
// Some inherent inefficiencies with the empty interface approach, but
// intended for the case where the data structure is not known.

func NeedleInHaystack(needle map[string]interface{}, haystack map[string]interface{}) (inHaystack bool, err error) {
	for key, val := range needle {
		if _, ok := haystack[key].(map[string]interface{}); ok {

			// reference doesn't match scalar
			if reflect.ValueOf(val).Kind() != reflect.Map {
				return inHaystack, nil
			}

			// element match, recurse
			inHaystack, err = NeedleInHaystack(val.(map[string]interface{}), haystack[key].(map[string]interface{}))
			if err != nil {
				return inHaystack, err
			}
		} else {
			// hit the last element, compare values
			if val == haystack[key] {
				return true, nil
			}
		}
	}
	// needle is not in haystack
	return inHaystack, nil
}
