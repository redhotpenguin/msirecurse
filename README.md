# msirecurse

## Recursively searches a `map[string]interface{}` structure for existence of a `map[string]interface{}` structure

### Motivation
I wrote this package when faced with the problem of searching a json object
for existence of a subset json object. It could be there's a package out
there which already does this, but I wasn't able to find it. If you are 
aware of one, please reach out!

Rather than being json specific, I just provided the recursive logic
to search the `map[string]interface{}` structures which are unmarshaled
from json. The `NeedleInHaystack` function is provided which accepts
two `map[string]interface{}` arguments, and returns a `bool` indicating
if the needle was found in the haystack.


```
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
    "redhotpenguin/msirecurse"
)

func main() {

	jsonFile, err := os.Open("needle.json")
	if err != nil {
		panic(err)
	}
	needleBytes, _ := ioutil.ReadAll(jsonFile)

	var needle map[string]interface{}
	json.Unmarshal(needleBytes, &needle)

	jsonFile, err = os.Open("haystack.json")
	if err != nil {
		panic(err)
	}

	haystackBytes, _ := ioutil.ReadAll(jsonFile)
	var haystack map[string]interface{}
	json.Unmarshal(haystackBytes, &haystack)

	fmt.Println("needle ", needle)
	fmt.Println("haystack ", haystack)
	isInHaystack, err := msirecurse.NeedleInHaystack(needle, haystack)
	fmt.Println("is in haystack", isInHaystack)
}
```
