order:
	protoc --go_out=. --go-grpc_out=.  ./order/order.proto
transaction:
	protoc --go_out=. --go-grpc_out=.  ./transaction/transaction.proto
product:
	protoc --go_out=. --go-grpc_out=.  ./product/product.proto
