## compozify convert

convert docker run command to docker compose file

```
compozify convert [flags] DOCKER_RUN_COMMAND
```

### Examples

```

# convert and write to stdout
$ compozify convert "docker run -i -t --rm alpine"

# write to file
$ compozify convert -w "docker run -i -t --rm alpine"

# write to file with custom name
$ compozify convert -w -o docker-compose.yml "docker run -i -t --rm alpine"

# alternative usage specifying beginning of docker run command
$ compozify convert -w -- docker run -i -t --rm alpine

```

### Options

```
  -a, --append-service   append service to existing compose file. Requires --out flag
  -h, --help             help for convert
  -o, --out string       output file path (default "compose.yml")
  -w, --write            write to file
```

### Options inherited from parent commands

```
  -v, --verbose   verbose output
```

### SEE ALSO

* [compozify](compozify.md)	 - compozify is a tool mainly for converting docker run commands to docker compose files

