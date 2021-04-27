# dkman: OPO Docker Manager

This app will help OPO to manage docker instances:

- Generates files for each project (.e.g. , initial Dockerfile/docker-compose/.gitlab-ci, bash prompt)
- Start/stops multiple containers at once
- Starts interactive shell for selected app(s)
- Displays container status

## Build

```
$ make                        # build for your host OS
$ make PLATFORM=darwin/amd64  # build for macOS
$ make PLATFORM=windows/amd64 # build for Windows x86_64
$ make PLATFORM=linux/amd64   # build for Linux x86_64
$ make PLATFORM=linux/arm     # build for Linux ARM
```

## Development

This app is dockerized. You do not need to install GoLang or any other dependency. Just Docker.

1. Open this directory in your favorite editor
2. After editing, run `make` to build your executable (see: [Build](#build))
3. Run the executable: `bin/dkman`

Another option:
1. Open this directory in your favorite editor
2. Run `$ bin/shell` to open an interactive shell in the container, with a shared volume to /go/src.
3. Develop as you normally would in your terminal.


## References

- https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/
  - https://github.com/chris-crone/containerized-go-dev