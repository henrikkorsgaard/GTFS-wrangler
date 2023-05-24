package gtfs

import (
	"os"
	"bufio"
	"io"

	"testing"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)


// We want to test unzipping given a file

func TestUnzipGTFSFromBytes(t *testing.T){

	expectedFileNames := []string{"agency.txt", "attributions.txt", "calendar.txt","shapes.txt","trips.txt","routes.txt","stops.txt","stop_times.txt","transfers.txt"}

	gtfs, err := filepath.Abs("../../data/GTFS.zip")
	if err != nil {
		t.Error("filepath.Abs returned error: " + err.Error() + "\nPlease put a GTFS zip file into the data dir for test purpose")
	}

	file, err := os.Open(gtfs)
	if err != nil {
		t.Error("File open returned error: " + err.Error())
	}
	defer file.Close()

	stat, err := file.Stat();
	if err != nil {
		t.Error("Error getting stat from file: " + err.Error())
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		t.Error("Error reading bytes from file: " + err.Error())
	}

	gtfsDir, err := UnzipGTFSFromBytes(bs)
	if err != nil {
		t.Error("unzipGTFS failed with error: " +  err.Error())
	}

	for _, name := range expectedFileNames {
		path := filepath.Join(gtfsDir, name)

		assert.FileExists(t, path)
	}

}