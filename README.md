# FaaSChat
A distributed chat server build entirely using OpenFaaS functions, CockroachDB and Protocol Buffers running on Kubernetes.

### Goals
- run on one server or thousands without any work
- provide basic functionality of users, chats and messages
- easy to add new features

### Non goals
- production ready chat system
- security
- authentication
- nice client

## Architecture
- every action a client needs is a OpenFaaS function with it's own endpoint
- functions do not carry any state
- functions communicate with each other via the internal gateway e.g. to check authentication
- cockroachdb is the only source of truth

## Build & deploy the server
1. clone
1. install openfaas with helm (insecure mode)
1. install cockroachdb with helm (insecure mode)
1. configure environment variables in functions/functions.yml
1. modify and run functions/deploy script to push to your docker repository and openfaas gateway (kubectl port-forward?)
1. if that all worked you're done :)

## Build & deploy the client
1. todo