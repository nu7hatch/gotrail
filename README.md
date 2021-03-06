# Go Trail

It's an uber simple logger with debug mode for golang.

Pic of sniffing cat:

![sniffing kitten](http://i.imgur.com/ilF3bhq.jpg)

## Installation

Goinstall the package:

    $ go install github.com/nu7hatch/gotrail

... and import it in your app:

    import "github.com/nu7hatch/gotrail"

## Usage

It's very simple to use, example:

```go
log := gotrail.New(os.Stdout)
log.Debug = true

log.Printf("hello from gotrail!")
log.Debugf("never gonna give %s up....", "you")
```

You can also use named (prefixed) loggers:

```go
log := gotrail.NewNamed(os.Stdout, "prefix")
```

Enjoy!

## License

Copyright (C) 2013 by Chris Kowalik <chris@nu7hat.ch><br />
Released under MIT License.
