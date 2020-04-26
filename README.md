# RabbitMQ
Este proyecto tiene como objetivo realizar una introducción a RabbitMQ 🎯 

## Instalación de RabbitMQ 

Puedes ejecutar un contenedor de Docker con el siguiente comando:

```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Una vez que descargue la imagen y se levante el contenedor, puedes acceder al dashboard: http://localhost:15672. Las credenciales por defecto son ```guest/guest```.

## Ejemplos:

* [Hello World](#01-hello-world)
* [Workers](#02-queues)


