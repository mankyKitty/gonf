package gonf

// Loads a configuration file into a _map[string]string_ of
// configuration key=value pairs. Delimiter is provided by
// you. :)
//
// Get it from [here](https://github.com/mankyKitty/gonf).
//
// Or use _go get_!
//
// <code>
//        go get github.com/mankyKitty/gonf
// </code>
//
// Import it
//
// <code>
//        import "github.com/mankyKitty/gonf"
// </code>
//
// Point it at your configuration file with your chosen method of
// delineation and marvel at the wonder of a neatly organised map
// of your configuration variables.
//
// <code>
//        conf, err := gonf.GetGonf("/path/to/file", "\n")
// </code>
//
// Currently there is no support for loading the configuration values
// into their respective types, everything is a _string_ at the moment.
//
// Original consideration was to let you specify if you wanted Gonf to
// panic, print, or fatal if the file could not be read/loaded etc but
// that seemed excessive for this iddy biddy package. You're a big coder,
// I'm sure you can manage some proper error handling. :)
import (
	"errors"
	"io/ioutil"
	"strings"
)

// Loads a configuration from a file into a map[string]string.
func GetGonf(fname, delim string) (map[string]string, error) {
	// Get the file.
	conf, err := ioutil.ReadFile(fname)

	// Check we haven't had any io errors, escape immediately
	// if anything is amiss.
	if err != nil {
		return nil, err
	}
	// ### Avoid zero length config file.
	// Be kind to our users and tell them what has happened.
	if len(conf) <= 0 {
		return nil, errors.New("Zero length configuration file!")
	}

	// Create somewhere to keep out results
	config := make(map[string]string)

	// Break out the file contents to parsing.
	lines := strings.Split(string(conf[:]), delim)

	// #### Analyse what we have for some config...
	// KISS principle is in full effect.
	for _, v := range lines {
		// Ensure we ignore commented lines.
		if len(v) > 0 && v[:1] != "#" {
			// Break out our respective values.
			line := retrieveKeyValuePair(v)
			// Trim our final values so we don't waste our time with gritty whitespace
			// That stuff gets in your teeth, it's horrible...
			config[strings.TrimSpace(line[0])] = strings.TrimSpace(line[1])
			// A moose might also have bit my sister. Mind you, she didn't whine
			// about it on film like some cry baby projectionist I know.
		}
	}
	// Oh yeah, actually give the values back.. Heh. That was close.
	return config, err
}

// Takes what is expected to be a key=value pair 'line' and
// ignores the comments and returns the split pair for input
// into our map
func retrieveKeyValuePair(line string) []string {
	// If our entry contains an inline comment
	if strings.Contains(line, "#") {
		// Change our line value to be the substring prior to the
		// start of the comment
		return strings.Split(line[:strings.Index(line, "#")], "=")
	}
	// If no comments were found then just hand it back.
	return strings.Split(strings.TrimSpace(line), "=")
}
