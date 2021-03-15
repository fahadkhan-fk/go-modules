module country

go 1.15

// this is how we give the path of the module we want to import in go mod file
replace cities => ../cities

// below require tag is auto generated when you will `run main.go`
require cities v0.0.0-00010101000000-000000000000 // indirect
