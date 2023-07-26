# dcmerge

[![Build Status](https://drone.cryptic.systems/api/badges/volker.raschek/dcmerge/status.svg)](https://drone.cryptic.systems/volker.raschek/dcmerge)
[![Docker Pulls](https://img.shields.io/docker/pulls/volkerraschek/dcmerge)](https://hub.docker.com/r/volkerraschek/dcmerge)

`dcmerge` is a small program to merge docker-compose files from multiple
sources. It is available via RPM and docker.

The dynamic pattern of a docker-compose file, that for example `environments`
can be specified as a string slice or a list of objects is currently not
supported. `dcmerge` expect a strict pattern layout. The `environments`, `ports`
and `volumes` must be declared as a slice of strings.

Dockercompose file can be read-in from different sources. Currently are the
following sources supported:

- File
- HTTP/HTTPS

Furthermore, `dcmerge` support different ways to merge multiple docker-compose
files.

- The default merge, add missing secrets, services, networks and volumes.
- The existing-win merge, add and protect existing attributes.
- The last-win merge, add or overwrite existing attributes.

## default

Merge only missing secrets, services, networks and volumes without respecting
their attributes. For example, when the service `app` is already declared, it is
not possible to add the service `app` twice. The second service will be
completely skipped.

```yaml
---
# cat ~/docker-compose-A.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=HelloWorld123
    image: example.local/app/name:0.1.0
---
# cat ~/docker-compose-B.yaml
services:
  app:
    image: app/name:2.3.0
    volume:
    - /etc/localtime:/etc/localtime
    - /dev/urandom:/etc/urandom
  db:
    image: postgres
    volume:
    - /etc/localtime:/etc/localtime
    - /dev/urandom:/etc/urandom
---
# dcmerge ~/docker-compose-A.yaml ~/docker-compose-B.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=HelloWorld123
    image: example.local/app/name:0.1.0
  db:
    image: postgres
    volume:
    - /etc/localtime:/etc/localtime
    - /dev/urandom:/etc/urandom
```

## existing-win

The existing-win merge protects existing attributes. For example there are two
different docker-compose files, but booth has the same environment variable
`CLIENT_SECRET` defined with different values. The first declaration of the
attribute wins and is for overwriting protected.

```yaml
---
# cat ~/docker-compose-A.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=HelloWorld123
    image: example.local/app/name:0.1.0
---
# cat ~/docker-compose-B.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=FooBar123
    image: example.local/app/name:0.1.0
---
# dcmerge --existing-win ~/docker-compose-A.yaml ~/docker-compose-B.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=HelloWorld123
    image: example.local/app/name:0.1.0
```

## last-win

The last-win merge overwrite recursive existing attributes. For example there
are two different docker-compose files, but booth has the same environment
variable `CLIENT_SECRET` defined with different values. The last passed
docker-compose file which contains this environment wins.

```yaml
---
# cat ~/docker-compose-A.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=HelloWorld123
    image: example.local/app/name:0.1.0
---
# cat ~/docker-compose-B.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=FooBar123
    image: example.local/app/name:0.1.0
---
# dcmerge --last-win ~/docker-compose-A.yaml ~/docker-compose-B.yaml
services:
  app:
    environments:
    - CLIENT_SECRET=FooBar123
    image: example.local/app/name:0.1.0
```