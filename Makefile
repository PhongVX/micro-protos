order:
	protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. ./order/order.proto
transaction:
	protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. ./transaction/transaction.proto
product:
	protoc --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. ./product/product.proto
