package gonf
// # Gonf!!!
// Loads a configuration from a file into a map[string]string.

// Currently no support for loading the configuration options into
// their respective types, but that should all come in good time.
//
// Original consideration was to let you specify if you wanted Gonf to
// panic, print, or fatal if the file could not be read/loaded etc but
// that seemed excessive for this iddy biddy package. You're a big coder,
// I'm sure you can manage some proper error handling. :)
import (
  "io/ioutil"
  "regexp"
  "strings"
)
// ### Parameters
// - fname _string_ - the file name with absolute path to the config file in question
func GetGonf(fname string) (map[string]string, error) {
  // Default separator for now...
  sep := "\n"

  // Get the file.
  conf, err := ioutil.ReadFile(fname)

  // ### _Be responsible and check for errors!_
  if err != nil {
    return nil, err
  }

  // ### Create somewhere to keep out results
  config := make(map[string]string)

  // ### Parse the file.
  lines := strings.Split(string(conf[:]), sep)

  // ### Analyse what we have for some config...
  // This is a relatively naive process at the moment.
  for _, v := range lines {
    // Find the lines that start with valid characters and contain a value.
    // It's only one value per line at the moment. :(
    matched, _ := regexp.MatchString("^[a-z, A-Z, 0-9]+(=| = )+", strings.TrimSpace(v))
    if matched && len(v) > 0 {
      // Break out our respective values.
      // __TODO__ Fix this so we can comment inline.
      line := strings.Split(v, "=")
      // Trim our final values so we don't waste our time with gritty whitespace
      // That stuff gets in your teeth, it's horrible...
      config[strings.TrimSpace(line[0])] = strings.TrimSpace(line[1])
      // A moose might also have bit my sister. She didn't whine about it on
      // film though like some cry baby projectionist I know.
    }
  }
  // Oh yeah, actually give the values back.. Heh. That was close.
  return config, err
}
