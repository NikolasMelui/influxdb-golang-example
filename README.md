# Influxdb-golang-example

## Requirements

* Compilers: `go`
* Utilities: `make`

## Run the project

Compile project:

```commandline
  make compile
```

Run project:

```commandline
  make run
```

Run tests:

```commandline
  make test
```

Cleanup the `./bin` directory:

```commandline
  make clean
```

### Run influxdb in the Docker container

```commandline
docker-compose --env-file test_influxdb.env up
```