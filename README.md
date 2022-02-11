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

Switch to the `build` directory:

```commandline
cd build
```

Create the `influxdb.env` file with credentials:

```
INFLUXDB_USERNAME=username
INFLUXDB_PASSWORD=password
INFLUXDB_TOKEN=your_super_secret_token
INFLUXDB_URL=http://localhost:8086
```

Run the composition:

```commandline
docker-compose --env-file .env up
```