# aProxy
aProxy is an IoT message proxy, is deployed between a device connections and an IoT message broker and acts like PEP (Policy Enhancement Point). It enables bidirectional connection and inspects packets that flow through it - especially authorization and authentication packets.

![arch][arch]

aProxy sends authorization demands to external authentication and authorization service - AM:DM. It protects against any unauthorized traffic between devices and IoT platform.

aProxy can do other packet transformation and there is an SDK for handling the traversing packets. Users can customize in-flight packet handling and transform packets (for example - change MQTT topic, or rewrite payload).

aProxy is typically deployed on-premise, in the enterprise cloud, in front of IoT platform. This way it is insured that messages never leave the private cloud - only authorization headers are sent to AM:DM SaaS for verification. Having in mind that aProxy is very lightweight service, written in modern Go programming language, it is very portable and easily deployable.

## Usage
```bash
git clone https://github.com/absmach/aproxy.git
cd aproxy
make
make docker-image
make run
```

## Configuration

The service is configured using the environment variables presented in the following table. Note that any unset variables will be replaced with their default values.

| Variable                        | Description                                    | Default   |
|---------------------------------|------------------------------------------------|-----------|
| APROXY_WS_HOST                  | WebSocket inbound (IN) connection host         | 0.0.0.0   |
| APROXY_WS_PORT                  | WebSocket inbound (IN) connection port         | 8080      |
| APROXY_WS_PATH                  | WebSocket inbound (IN) connection path         | /mqtt     |
| APROXY_WSS_PORT                 | WebSocket Secure inbound (IN) connection port  | 8080      |
| APROXY_WSS_PATH                 | WebSocket Secure inbound (IN) connection path  | /mqtt     |
| APROXY_WS_TARGET_SCHEME         | WebSocket Target schema                        | ws        |
| APROXY_WS_TARGET_HOST           | WebSocket Target host                          | localhost |
| APROXY_WS_TARGET_PORT           | WebSocket Target port                          | 8888      |
| APROXY_WS_TARGET_PATH           | WebSocket Target path                          | /mqtt     |
| APROXY_MQTT_HOST                | MQTT inbound connection host                   | 0.0.0.0   |
| APROXY_MQTT_PORT                | MQTT inbound connection port                   | 1883      |
| APROXY_MQTTS_PORT               | MQTTS inbound connection port                  | 8883      |
| APROXY_MQTT_TARGET_HOST         | MQTT broker host                               | 0.0.0.0   |
| APROXY_MQTT_TARGET_PORT         | MQTT broker port                               | 1884      |
| APROXY_CLIENT_TLS               | Flag that indicates if TLS should be turned on | false     |
| APROXY_CA_CERTS                 | Path to trusted CAs in PEM format              |           |
| APROXY_SERVER_CERT              | Path to server certificate in pem format       |           |
| APROXY_SERVER_KEY               | Path to server key in pem format               |           |
| APROXY_LOG_LEVEL                | Log level                                      | debug     |
| APROXY_MQTT_ADAPTER_CONFIG_FILE | Config file path. This overites env if set.    |           |
| APROXY_RELEASE_TAG              | Docker release tag.                            | latest    |
| APROXY_THINGS_URL               | Things url.                                    |           |
| APROXY_THINGS_AUTH_GRPC_URL     | Things GRPC URL for authentication.            |           |
| APROXY_THINGS_AUTH_GRPC_TIMEOUT | Things GRPC timeout duration                   | 1s        |

## License
[Apache-2.0](LICENSE)

[arch]: https://github.com/absmach/docs/blob/main/img/aProxy.jpg