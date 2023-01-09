package main
import (
    "errors"
    "fmt"
)
func getValue(obj map[string]any, keys []string) (any, error) {
    nestedObj := obj
    n := len(keys)
    // Checking if the object is nil
    if nestedObj == nil {
        return nil, errors.New("object is nil")
    }
    // Checking if the keys are not empty
    if n == 0 {
        return nil, errors.New("no keys found in the array")
    }
    for i, k := range keys {
        // Checking if the key exists inside the object
        val, ok := nestedObj[k]
        if !ok {
            return nil, fmt.Errorf("key %s not found", k)
        }
        // checking if the the recived value for the key is an object or not.
        // If the recived value is an object, then value becomes the next object.
        if v1, ok := val.(map[string]any); !ok {
            // If the recived value for the key is not an object and
            // if this is not the last key in the array of keys, then
            // throw an error since the next key won't exist.
            if i < n-1 {
                return nil, errors.New("invalid object or keys")
            }
            return val, nil
        } else {
            nestedObj = v1
        }
    }
    return nestedObj, nil
}
func main() {
    obj := map[string]any{
        "a": map[string]any{
            "b": map[string]any{
                "c": 1,
            },
        },
    }
    fmt.Println(getValue(obj, []string{"a", "b", "c", "d"}))
    fmt.Println(getValue(obj, []string{"a", "b", "c"}))
    fmt.Println(getValue(obj, []string{"a", "b"}))
    fmt.Println(getValue(obj, []string{"c"}))
    fmt.Println(getValue(obj, []string{}))
    fmt.Println(getValue(nil, []string{}))
    fmt.Println(getValue(obj, nil))
}