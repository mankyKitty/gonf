gonf
====
Loads a configuration from a file into a map[string]string.

Comments are supported, just use "#" at the start of the line. Inline comments are now supported as well! Woo!

## Installation

        go get github.com/mankyKitty/gonf

## Usage

        import (
          "github.com/mankyKitty/gonf"
        )

        func main() {
          conf, err := gonf.GetGonf("/path/to/file", "\n")
        }

        $ [key:value, key:value ...etc]

Currently no support for loading the configuration options into their respective types, but that should all come in good time. Original consideration was to let you specify if you wanted Gonf to panic, print, or fatal if the file could not be read/loaded etc. But that seemed excessive for this iddy biddy package. You're a big coder, I'm sure you can manage some proper error handling.
