GoBank
=========

GoBank -- Go client for the awesome, over-the-wire test double [Mountebank](http://www.mbtest.org/).

THIS PROJECT IS CURRENTLY UNDER DEVELOPMENT.

ALTHOUGH MOST OF THE HTTP IMPOSTER FUNCTIONALITY IS COVERED, THE FLUENT API MIGHT CHANGE.

PLEASE CHECK AGAIN LATER...

## Installation

```bash
$ go get github.com/durmaze/gobank
```

## Documentation
See [Go Doc](http://godoc.org/github.com/durmaze/gobank) or [Go Walker](http://gowalker.org/github.com/durmaze/gobank) for usage and details.

## Status
**TODO** Add Travis Build Status, Go coverage, GoDoc, License, etc.

## Why should you use GoBank?

GoBank makes things simple if you'd like to use Mountebank in your tests. As recommended by Mountebank, you're better off creating and deleting your imposters in your test lifecycle (i.e. setup and teardown hooks).

GoBank provides a Go API, which directly maps to Mountebank's REST API.

With GoBank, you can simply create an Imposter via a Fluent API,

```go
// build an Imposter
imposter := NewImposterBuilder().Protocol("http").Port(4546).Build()
```
and then, you can publish your imposter to Mountebank.
```go
// publish your Imposter to Mountebank
client := mountebank.NewClient(MountebankUri)
client.CreateImposter(imposter)
```

## Contributing to GoBank:

If you want to contribute, feel free to send me a pull request with testing.

Currently, only HTTP imposters are supported. Any contributions on HTTPS, SMTP and TCP are more than welcome.

Thanks to all contributors thus far:
@alperkose

## License

GoBank is MIT License.
