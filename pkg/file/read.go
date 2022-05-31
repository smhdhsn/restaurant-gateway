package file

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// ReadJsonFile is responsible for reading bytes of a JSON file.
func ReadJsonFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, errors.Wrap(err, "error on openning file")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "error on reading bytes")
	}

	return b, nil
}
