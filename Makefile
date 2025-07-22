export ROOT_MOD=github.com/yzc/orange-review
.PHONY: gen-gateway
gen-gateway:
	@cd app/gateway && cwgo server \
		-I . \
		--type HTTP \
		--module ${ROOT_MOD}/app/gateway \
		--service gateway \
		--idl ../../idl/gateway.proto \

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module ${ROOT_MOD}/app/${svc} --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto
