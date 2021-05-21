.PHONE: gen
gen:
	oapi-codegen -generate "types,server,spec" -o app/openapi.gen.go -package app ./api.def.yaml
