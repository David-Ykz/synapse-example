#!/bin/bash

make undeploy-all

cd ../synapse
make build-broker
make build-agent-worker

cd ../synapse-example
make build-example
make build-handlers

make deploy-all