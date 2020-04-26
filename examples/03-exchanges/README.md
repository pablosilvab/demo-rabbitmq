# Exchanges

En este ejemplo el productor ```emit-log``` envía un mensaje hacia un ```exchange```, cuyo mensaje posee una ```routing-key```, la cual permite establecer un binding entre la cola y el exchange.  

El receptor ```receive-log``` puede escuchar los mensajes mediante el binding que se realiza entre el exchange y la cola.

### Ejecutar Sender

```
make emit-log
```

### Ejecutar Receiver

```
make receive-log
```