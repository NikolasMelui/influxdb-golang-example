# Influxdb-golang-example

## Requirements

* Compilers: `go`
* Utilities: `make`

## Run the project

Compile:

```commandline
  make compile
```

Run:

```commandline
  make run
```

Cleanup the `./bin` directory:

```commandline
  make clean
```

### Run influxdb in the Docker container

```commandline
docker-compose --env-file test_influxdb.env up
```