deploy-all: deploy-broker deploy-handlers deploy-agent-worker deploy-example

undeploy-all: undeploy-example undeploy-agent-worker undeploy-handlers undeploy-broker

build-example:
	docker build -t synapse-example:latest -f example/Dockerfile .
	docker save synapse-example:latest | sudo k3s ctr images import -

build-handlers:
	docker build -t synapse-handlers:latest -f handlers/Dockerfile .
	docker save synapse-handlers:latest | sudo k3s ctr images import -

deploy-broker:
	kubectl apply -f deployments/broker.yaml

undeploy-broker:
	kubectl delete -f deployments/broker.yaml --ignore-not-found

deploy-agent-worker:
	kubectl apply -f deployments/agent_worker.yaml

undeploy-agent-worker:
	kubectl delete -f deployments/agent_worker.yaml --ignore-not-found

deploy-example:
	kubectl apply -f deployments/example.yaml

undeploy-example:
	kubectl delete -f deployments/example.yaml --ignore-not-found

deploy-handlers:
	kubectl apply -f deployments/handlers.yaml

undeploy-handlers:
	kubectl delete -f deployments/handlers.yaml --ignore-not-found


















# deploy-example:
# 	kubectl delete -f deployments/example.yaml --ignore-not-found
# 	kubectl apply -f deployments/example.yaml

# update-example: build-example deploy-example




# undeploy-all:
# 	kubectl delete -f deployments/agent_worker.yaml --ignore-not-found
# 	kubectl delete -f deployments/broker.yaml --ignore-not-found
# 	kubectl delete -f deployments/example.yaml --ignore-not-found



# build-broker:
# 	docker build -t synapse-broker:latest -f broker/Dockerfile .
# 	docker save synapse-broker:latest | sudo k3s ctr images import -

# deploy-broker:
# 	kubectl apply -f deployments/broker.yaml

# undeploy-broker:
# 	kubectl delete -f deployments/broker.yaml --ignore-not-found

# update-broker: build-broker undeploy-broker deploy-broker



# build-agent-worker:
# 	docker build -t synapse-agent-worker:latest -f workers/Dockerfile .
# 	docker save synapse-agent-worker:latest | sudo k3s ctr images import -

# deploy-agent-worker:
# 	kubectl apply -f deployments/agent_worker.yaml

# undeploy-agent-worker:
# 	kubectl delete -f deployments/agent_worker.yaml --ignore-not-found

# update-agent-worker: build-agent-worker undeploy-agent-worker deploy-agent-worker
