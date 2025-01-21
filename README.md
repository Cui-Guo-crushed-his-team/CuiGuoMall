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
## Make order
| order           | introduce                                               |
|-----------------|---------------------------------------------------------|
| make gen-client | generate rpc client and stub code in rpc module         |
| make gen-server | generate rpc server code in app/{module_name}           |
| make gen-rpc    | execute "make gen-client" and "make gen-server" at once |
| make tidy       | execute "go mod tidy" in all modules                    |
