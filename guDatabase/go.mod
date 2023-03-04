module github.com/moondevgo/goUtils/guDatabase

go 1.19

replace github.com/moondevgo/goUtils/guBasic => ../guBasic

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/moondevgo/goUtils/guBasic v0.0.0-00010101000000-000000000000
)

require (
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
