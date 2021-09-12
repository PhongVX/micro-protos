order:
	protoc --go_out=. --go-grpc_out=.  order.proto
transaction:
	protoc --go_out=. --go-grpc_out=.  transaction.proto
product:
	protoc --go_out=. --go-grpc_out=.  product.proto
