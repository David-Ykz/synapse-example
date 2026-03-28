build-example:
	docker build -t synapse-example:latest -f example/Dockerfile .
	docker save synapse-example:latest | sudo k3s ctr images import -

deploy-example:
	kubectl delete -f deployments/example.yaml
	kubectl apply -f deployments/example.yaml

update-example: build-example deploy-example

build-agent-worker:
	docker build -t synapse-agent-worker:latest -f workers/Dockerfile .
	docker save synapse-agent-worker:latest | sudo k3s ctr images import -

deploy-agent-worker:
	kubectl delete -f deployments/agent_worker.yaml
	kubectl apply -f deployments/agent_worker.yaml

update-agent-worker: build-agent-worker deploy-agent-worker
