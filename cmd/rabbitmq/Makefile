.PHONY: add_vhost set_permissions


add_vhost:
	docker exec -it rabbitmq rabbitmqctl add_vhost auth

set_permissions:
	docker exec -it rabbitmq rabbitmqctl set_permissions -p auth root ".*" ".*" ".*"