.PHONY: gen-rpc-client gen-rpc-server gen-http-client gen-http-server tidy gen-rpc

gen-rpc-client:
	@./scripts/gen_rpc.sh client

gen-rpc-server:
	@./scripts/gen_rpc.sh server

gen-http-client:
	@./scripts/gen_http.sh client

gen-http-server:
	@./scripts/gen_http.sh server
gen-rpc: gen-rpc-client gen-rpc-server
tidy:
	@./scripts/gotidy.sh
