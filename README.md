# CuiGuoMall
A "simplified version" of a microservice Douyin shopping mall made by Kitex and Hertz

## Technology Stack
| technology | introduce |
|------------|-----------|
| cwgo       | -         |
| kitex      | -         |
| Hertz      | -         |
| MySQL      | -         |
| Redis      | -         |
| Prometheus | -         |
| Jaeger     | -         |
| Docker     | -         |
| RabbitMQ   | -         |
| Grafana    | -         |
## Make command
`Make sure you have executed "sudo chmod +x" to script files in the "scripts" directory`

| order                | introduce                                                       |
|----------------------|-----------------------------------------------------------------|
| make gen-rpc-client  | generate rpc client and stub code in rpc module                 |
| make gen-rpc-server  | generate rpc server code in app/{module_name}                   |
| make gen-http-server | generate http server code in app/gateway                        |
| make gen-http-client | generate http client code in ./client/http/{api_name}           |
| make gen-rpc         | execute "make gen-rpc-client" and "make gen-rpc-server" at once |
| make tidy            | execute "go mod tidy" in all modules                            |
