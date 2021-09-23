order:
	protoc --proto_path=. --go_out=. --go-grpc_out=.  ./order/order.proto
transaction:
	protoc --proto_path=. --go_out=. --go-grpc_out=.  ./transaction/transaction.proto
product:
	protoc --proto_path=. --go_out=. --go-grpc_out=.  ./product/product.proto
