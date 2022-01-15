# Payments Service
Service to handle payments via multiple different service providers.


## Generate Server
Generated server codes should be inside the infrastructures/clients package.
```shell
swagger generate server -f ./docs/spec/swagger.yaml -s ./infrastructures/server/. -m ./infrastructures/server/models/.
```


## Generate Payments Provider's clients
Generated client codes should be inside the infrastructures/clients package.
### Worldpay
#### triPOS Lane API
```shell
swagger generate client -f ./docs/spec/external/triPOS\ Lane\ API-3.1.0.yaml -m ./infrastructures/clients/triposlane/models -c ./infrastructures/clients/triposlane/.
```

#### triPOS Transaction API
```shell
swagger generate client -f ./docs/spec/external/triposcert.vantiv.com-3.1.0.yaml -m ./infrastructures/clients/tripos/models -c ./infrastructures/clients/tripos/.
```
### Windcave