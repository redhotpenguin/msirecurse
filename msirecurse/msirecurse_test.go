package msirecurse

import (
	"testing"
)

func TestNeedleInHaystack(t *testing.T) {
	{
		// successful test case
		needle := map[string]interface{}{"foo": map[string]interface{}{"bar": "baz"}}
		haystack := map[string]interface{}{"foo": map[string]interface{}{"bar": "baz", "zim": "zam"}}
		inHaystack, err := NeedleInHaystack(needle, haystack)
		if err != nil {
			t.Errorf("test failure %v", err)
		}
		if !inHaystack {
			t.Errorf("test failure, needle not in haystack")
		}
	}
	{
		// unsuccessful test case
		needle := map[string]interface{}{"foo": "bar"}
		haystack := map[string]interface{}{"foo": map[string]interface{}{"bar": "baz", "zim": "zam"}}
		inHaystack, err := NeedleInHaystack(needle, haystack)
		if err != nil {
			t.Errorf("test failure %v", err)
		}
		if inHaystack {
			t.Errorf("test failure, needle in haystack")
		}
	}
	{
		// value mismatch
		needle := map[string]interface{}{"foo": "bar"}
		haystack := map[string]interface{}{"foo": map[string]interface{}{"blarg": "baz", "zim": "zam"}}
		inHaystack, err := NeedleInHaystack(needle, haystack)
		if err != nil {
			t.Errorf("test failure %v", err)
		}
		if inHaystack {
			t.Errorf("test failure, needle in haystack")
		}
	}
	{
		// simple needle, not a match
		needle := map[string]interface{}{"foo": "blarg"}
		haystack := map[string]interface{}{"foo": map[string]interface{}{"blarg": "baz", "zim": "zam"}}
		inHaystack, err := NeedleInHaystack(needle, haystack)
		if err != nil {
			t.Errorf("test failure %v", err)
		}
		if inHaystack {
			t.Errorf("test failure, needle not in haystack")
		}
	}

}
