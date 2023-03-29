# Set up

```
go mod init github.com/Tabed23/Fun-Application-Projects/tree/main/JobsAPI-GraphQL
```


# Get The Package
```
go get github.com/99designs/gqlgen
```


## Make Tools dot GO file
```
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
```


## Install dependencies
```
go mod tidy
```


## initialize the graphql generator
```
go run github.com/99designs/gqlgen init
```



## TO Generate  new graphql schema
```
go run github.com/99designs/gqlgen generate
```
