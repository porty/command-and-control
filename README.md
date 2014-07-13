# command-and-control

UAV Onboard Computer Command and Control Agent

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/command-and-control
```

Add your long-running agent logic to `command/agent/command.go`, and any status or action commands you need to `commands.go`.

### Testing

``make test``

## License

_Fill me in._

## Contributing

See `CONTRIBUTING.md` for more details.