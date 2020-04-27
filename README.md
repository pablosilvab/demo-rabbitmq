# RabbitMQ
Este proyecto tiene como objetivo realizar una introducción a RabbitMQ 🎯 

## Instalación de RabbitMQ 

Puedes ejecutar un contenedor de Docker con el siguiente comando:

```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Una vez que descargue la imagen y se levante el contenedor, puedes acceder al dashboard: http://localhost:15672. Las credenciales por defecto son ```guest/guest```.

### Ejemplos:

* [Ejemplos](examples)

## Job 

El objetivo del Job es enviar mensajes a una cola, mediante el uso de las librerías ```sender``` y ```receiver``` almacenadas en este repositorio. El ejemplo se debe correr de forma local. (por ahora)

### Requisitos:

* Elasticsearch  

```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.2  
```

* RabbitMQ  
```
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

