run-job:
	go run job/main.go

run-elastic:
	@echo "Run elasticsearch container"
	docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.6.2 

run-rabbit:
	@echo "Run RabbitMQ"
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management