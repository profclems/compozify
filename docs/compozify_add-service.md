## compozify add-service

Add a service to an existing docker-compose file

### Synopsis

Converts the docker run command to docker compose and adds as a new service to an existing docker-compose file.
If no file is specified, compozify will look for a docker compose file in the current directory.
If no file is found, compozify will create one in the current directory.
Expected file names are docker-compose.[yml,yaml], compose.[yml,yaml]


```
compozify add-service [flags] DOCKER_RUN_COMMAND
```

### Examples

```

# add service to existing docker-compose file in current directory
$ compozify add-service "docker run -i -t --rm alpine"

# add service to existing docker-compose file
$ compozify add-service -f /path/to/docker-compose.yml "docker run -i -t --rm alpine"

# write to file
$ compozify add-service -w -f /path/to/docker-compose.yml "docker run -i -t --rm alpine"

# alternative usage specifying beginning of docker run command without quotes
$ compozify add-service -w -f /path/to/docker-compose.yml -- docker run -i -t --rm alpine

# add service with custom name
$ compozify add-service -w -f /path/to/docker-compose.yml -n my-service "docker run -i -t --rm alpine"

```

### Options

```
  -f, --file string           Compose file path
  -h, --help                  help for add-service
  -n, --service-name string   Name of the service
  -w, --write                 write to file
```

### Options inherited from parent commands

```
  -v, --verbose   verbose output
```

### SEE ALSO

* [compozify](compozify.md)	 - compozify is a tool mainly for converting docker run commands to docker compose files

