# oapi-codegen-bug
Minimum example of bug in oapi-codegen


## How to (test)

Run tests in app package. Please see comments for why they shouldn't fail.

## How to (full)

Run the program ( `go run ./main.go` ) and then make a request to: `http://localhost:4000/api/v1/people`.  
**Expected behaviour**: Since `day` query parameter is defined as required on [line 17](https://github.com/Makinami/oapi-codegen-bug/blob/main/api.def.yaml#L17), a 4xx error should be returned.  
**Actual behaviour**: A string representing zero value time is returned.

## Analysis

The problem is with how date and time are internally handled.

If a query parameter is required, it's generated as a non-pointer member. When passed to `BindQueryParameter` it goes to the switch statement on [line 326 in bindparam.go](https://github.com/deepmap/oapi-codegen/blob/master/pkg/runtime/bindparam.go#L326) where it's type is checked:

```go
// abbreviated
switch k {
case reflect.Slice:
    // contains checks if param is required and found
case reflect.Struct:
    // no additional checks
    err = bindParamsToExplodedObject(paramName, queryParams, output)
default:
    // contains checks if param is required and found
}
```

Since oapi-codegen's `Date` type and `time.Time` are both structs they fall into second case where information about whether the parameter is required or not is completely lost.