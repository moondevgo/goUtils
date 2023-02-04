module github.com/moondevgo/go-mods/network

go 1.18

replace github.com/moondevgo/go-mods/basic => ../basic

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/machinebox/graphql v0.2.2
	github.com/moondevgo/go-mods/basic v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
)

require (
	github.com/matryer/is v1.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.mongodb.org/mongo-driver v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
