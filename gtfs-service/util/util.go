package util 

import (
	"path/filepath"
	"io"
	"os"
	"bufio"
)

func GetBytesFromZipFile(path string) (zbytes []byte, err error) {
	
	gtfs, err := filepath.Abs(path)
	if err != nil {
		return
	}

	file, err := os.Open(gtfs)
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat();
	if err != nil {
		return
	}

	zbytes = make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(zbytes)
	if err != nil && err != io.EOF {
		return
	}

	return
}
