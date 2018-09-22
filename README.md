Command line tool to capture & replay grpc packets. Works only for ethernet interfaces.

## Usage
### To capture packets
`grpc-capture <ethernet interface> <mtu size, default is 1500> <port number> <output file name>`


### Dev Setup
1. `make` will download dependencies and make the binary called `grpc-replay` in $GOPATH/bin.