# Queues

En este ejemplo el productor ```new-task``` envía una tarea a la queue ```hello``` y los ```workers``` escuchan los mensajes a la cola que lleva ese nombre, los cuales pueden escuchar la cola de forma paralela, de esta forma si hay una cola de trabajo, puedes añadir workers.

### Ejecutar Sender

```
make run-sender
```

### Ejecutar Receiver

```
make run-receiver
```