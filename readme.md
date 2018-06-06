# image-optimizer

`image-optimizer` is a [Docker](https://www.docker.com) image based on the latest [Alpine linux](https://alpinelinux.org) to provide microservice of optimizing `jpeg` images.

## Usage

Run `image-optimizer` with docker:

```bash
$ docker run \
  -d \
  -p 8080:8080 \
  --name image-optimizer \
  egorovli/image-optimizer
```

Or embed it into your `docker-compose.yml`:

```yaml
version: '3'
services:
  # ...

  image-optimizer:
    image: egorovli/image-optimizer
    expose:
      - "8080"

  some-service:
    # ...
    links:
      - image-optimizer
```

## Configuration

Configuration is supplied via environment variables and the supports the following options:

| Environment Variable | Type     | Description                                  | Default Value |
| -------------------- | -------- | -------------------------------------------- | ------------- |
| ENV                  | `string` | Environment                                  | production    |
| PORT                 | `int`    | Port to bind to                              | 8080          |
| HOST                 | `string` | Host to listen on                            | 0.0.0.0       |
| QUALITY              | `int`    | Default quality to use in `cjpeg` conversion | 80            |
| EXECUTABLE_PATH      | `string` | Executable path to call                      | cjpeg         |

## API

### POST, PUT /

Receives binary file and attempts to optimize it with `cjpeg`.

### GET /health

Get current status of the app.