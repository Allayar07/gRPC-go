# First clone the repository
```git clone git@github.com:Allayar07/gRPC-go.git```
## Step-1:
* run the gRPC server by this command: ```go run cmd/main.go```
## Step-2
* Then run gRPC client Using this command: ```go run  client/client.go 3 4```


Actually app is very simple it just adding two integer and return the result. example is above :)

## For initializing pb.go files from proto: 
```protoc --go_out=gRPC --go_opt=paths=source_relative --go-grpc_out=gRPC --go-grpc_opt=paths=source_relative api.proto```

But firstly check you have been installed protoc-compile in your system!!! for installation link: https://grpc.io/docs/protoc-installation/
