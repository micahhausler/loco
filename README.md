[![Build Status](https://travis-ci.org/micahhausler/loco.svg)](https://travis-ci.org/micahhausler/loco)
[![https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](http://godoc.org/github.com/micahhausler/loco/)

# Docker Login Compressor - loco

loco is a utility for bundling Docker registry credentials for
[Marathon](https://mesosphere.github.io/marathon/). Marathon can pull images
from [private repositories](https://mesosphere.github.io/marathon/docs/native-docker-private-registry.html)
but requires a `fetch` object with the bundled credentials. Rather than logging
in using the docker cli and altering a developer's `~/.docker/config.json` which
may have have credentials to other registries, loco creates the directory and
config file in memory, and generates a `docker.tgz`.

```json
{
    ...
    "fetch": [
        {
            "uri": "http://private.uri.com/docker.tar.gz",
            "extract": true
        }
    ]
}
```

## Installation

```
go get github.com/micahhausler/loco
# or
docker pull micahhausler/loco
```

## Usage

```
loco -h
Docker Login Compressor
Usage of ./loco:

  -o, --output string     The file to create. Use "-" for Stdout. (default "docker.tgz")
  -p, --password string   The password
  -r, --registry string   Specify a specific registry (default "https://index.docker.io/v1/")
  -u, --username string   The user to login as
      --version           print version and exit
```

## License
MIT License. See [License](/LICENSE) for full text
