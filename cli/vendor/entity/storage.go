package entity

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-errors/errors"
)

type storage struct {
	path string
}

// Panic :tips error
func Panic(err error) {
	panic(errors.Wrap(err.Error()+"\n"+string(errors.Wrap(err, 1).Stack()), 1))
}

func (files *storage) read(itf interface{}) {
	file, err := os.Open(files.path)
	defer file.Close()
	if os.IsNotExist(err) {
		logger.Printf("'%s' does not exist, please new an empty storage file \n", files.path)
		return
	}
	if err != nil {
		Panic(err)
	}

	if err := json.NewDecoder(file).Decode(itf); err != nil {
		fmt.Fprintf(os.Stderr, "decode'%s' failed.\n", files.path)
		Panic(err)
	}
	logger.Printf("successfully decoded \n")
}

func (files *storage) write(ptr interface{}) {
	file, err := os.OpenFile(files.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		Panic(err)
	}
	if err := json.NewEncoder(file).Encode(ptr); err != nil {
		Panic(err)
	}
	logger.Printf("successfully written\n")
}
