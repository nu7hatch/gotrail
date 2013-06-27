# Go Trail

It's an uber simple logger with debug mode for golang.

## Installation

Goinstall the package:

    $ go install github.com/nu7hatch/gotrail

... and import it in your app:

    import "github.com/nu7hatch/gotrail"

## Usage

It's very simple to use, example:

    log := gotrail.New(os.Stdout)
    log.Debug = true

    log.Printf("hello from gotrail!")
    log.Debugf("never gonna give %s up....", "you")

You can also use named (prefixed) loggers:

    log := gotrail.NewNamed(os.Stdout, "prefix")

Enjoy!

## License

Copyright (C) 2013 by Chris Kowalik <chris@nu7hat.ch>
Released under MIT License.
