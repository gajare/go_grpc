create a file proto/greet.proto

udpate the code 
hit this command to generate  2 more fils (greet_grpc.pb.go and greet.pb.go) 
$  protoc --go_out=. --go-grpc_out=. proto/greet.proto

$ fo mod tidy

update server.go
then 
update client .go