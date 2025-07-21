export ROOT_MOD=github.com/yzc/orange-review
# generate review system server & client code
.PHONY: gen-review
gen-review:
	@cd app/review_server && cwgo server \
		-I ../../idl \
		--pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" \
		--type RPC \
		--module ${ROOT_MOD}/app/review_server \
		--service review \
		--idl ../../idl/review.proto

.PHONY: gen-client
gen-review-client:
	@cd rpc_gen && cwgo client \
		-I ../idl \
		--type RPC \
		--module ${ROOT_MOD}/rpc_gen \
		--service review \
		--idl ../idl/review.proto

# generate IDGenerate Service server & client code
.PHONY: gen-generator
gen-generator:
	@cd app/id_generator && cwgo server \
		-I ../../idl \
		--pass "-use  ${ROOT_MOD}/rpc_gen/kitex_gen" \
		--type RPC \
		--module ${ROOT_MOD}/app/id_generator \
		--service generator \
		--idl ../../idl/generator.proto
	@cd rpc_gen && cwgo client \
		-I ../idl \
		--type RPC \
		--module ${ROOT_MOD}/rpc_gen \
		--service generator \
		--idl ../idl/generator.proto

.PHONY: gen-validator
gen-validator:
	@cd app/review_server && protoc \
		-I ../../idl \
		--go_out=. \
		--go_opt=module=github.com/yzc/orange-review/app/review_server \
		--validator_out=. \
		--validator_opt=module=github.com/yzc/orange-review/app/review_server \
		../../idl/review.proto

.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server \
		-I . \
		--type HTTP \
		--module ${ROOT_MOD}/app/gateway \
		--service gateway \
		--idl ../../idl/gateway.proto \
