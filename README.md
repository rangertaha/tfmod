# TFMod

A private terraform module registry and cli written in Go

## Requirements

To contribute or build the binaries you need to install Go and 
setup your workstation

- [Go](https://golang.org/doc/install)

### Binaries

The following builds and installs **tfmod** into the `$GOPATH/bin`. It also places a copy in the `build` directory

```bash
make build
```

### Docker Image

Build a docker image for **tfmod**

```bash
make image
```
