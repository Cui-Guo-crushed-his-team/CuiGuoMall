.PHONY: gen-client
gen-client:
	@./scripts/gen_rpc_client.sh

.PHONY: gen-server
gen-server:
	@./scripts/gen_rpc_server.sh

.PHONY: gen-rpc
gen-rpc: gen-client gen-server

.PHONY: tidy
tidy:
	@./scripts/gotidy.sh
