# Basics of Go

## Dependencies tracking

- Import packages contained in other modules -> manage those dependencies on your code's own module
- The module for tracking is go.mod file
- go.mod will be included in your source code
- Enable dependencies tracking:

```bash
go mod init <module-path>
```

- If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.

```bash
go mod init example/hello
# go: creating new go.mod: module example/hello
# the created go.mod file:
# module example/hello
#
# go 1.26.1
```

- Add new modules requirements (download) and sums

```bash
go mod tidy
# go: finding module for package rsc.io/quote
# go: downloading rsc.io/quote v1.5.2
# go: found rsc.io/quote in rsc.io/quote v1.5.2
# go: downloading rsc.io/sampler v1.3.0
# go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
# the created go.mod file:
# module example/hello
#
# go 1.26.1
#
# require (
# 	rsc.io/quote v1.5.2
# 	
# 	rsc.io/sampler v1.3.0 // indirect
# )
```