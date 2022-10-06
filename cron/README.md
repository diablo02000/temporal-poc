# Cron workflow

Create simple scheduled workflow which will ping `www.google.fr`
website every minute.

## How to run

* Ensure than Temporal server is running.
* Start the worker:

```bash
$ go run cron/worker/main.go
```

* Start the starter:

```bash
$ go run cron/starter/main.go
```

* Check on Temporal server web ui
