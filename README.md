# Temporal POC

Several sample Temporal's workflows application.

## Introduction

[Temporal](https://temporal.io/) is a developer-first, open source platform that ensures
the successful execution of services and applications (using workflows).

## Requirement

This repository required:

* [Temporal Go SDK](https://github.com/temporalio/sdk-go)
* [Temporal Server](https://github.com/temporalio/temporal)

## How to use

You can use:

* With [Docker Temporal dev server](https://github.com/temporalio/docker-compose).
* With gitops.

### Docker Temporal

* Get docker-compose YAML file from submodule repository:

```bash
$ git submodule update --init --recursive
```

* Change directory to `local-engine` folder.
* Run docker compose command:

```
$ docker compose up -d
```

* Ensure Docker containers is running:

```
$ docker ps
CONTAINER ID   IMAGE                           COMMAND                  CREATED         STATUS         PORTS                                                                                         NAMES
3ccdfa5652d9   temporalio/admin-tools:1.18.0   "tail -f /dev/null"      4 seconds ago   Up 1 second                                                                                                  temporal-admin-tools
bc369762cea7   temporalio/ui:2.6.2             "./start-ui-server.sh"   4 seconds ago   Up 1 second    0.0.0.0:8080->8080/tcp, :::8080->8080/tcp                                                     temporal-ui
adac8fa0b1b4   temporalio/web:1.15.0           "docker-entrypoint.s…"   4 seconds ago   Up 1 second    0.0.0.0:8088->8088/tcp, :::8088->8088/tcp                                                     temporal-web
474f26ace192   temporalio/auto-setup:1.18.0    "/etc/temporal/entry…"   4 seconds ago   Up 2 seconds   6933-6935/tcp, 6939/tcp, 7234-7235/tcp, 7239/tcp, 0.0.0.0:7233->7233/tcp, :::7233->7233/tcp   temporal
5e869bbcd0e5   elasticsearch:7.16.2            "/bin/tini -- /usr/l…"   2 days ago      Up 2 seconds   9200/tcp, 9300/tcp                                                                            temporal-elasticsearch
186275c15f53   postgres:13                     "docker-entrypoint.s…"   2 days ago      Up 2 seconds   5432/tcp                                                                                      temporal-postgresql
```
