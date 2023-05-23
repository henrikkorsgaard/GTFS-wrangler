package gtfs

import (
	"testing"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)


// We want to test unzipping given a file

func TestUnzipGTFSReturnsDir(t *testing.T){
	gtfs, err := filepath.Abs("../data/GTFS.zip")
	if err != nil {
		t.Error("filepath.Abs returned error: " + err.Error())
	}
	gtfsDir, err := unzipGTFS(gtfs)
	if err != nil {
		t.Error("unzipGTFS failed with error: " +  err.Error())
	}

	assert.DirExists(t, gtfsDir)
}

func TestUnzipGTFSUnzippedFiles(t *testing.T){
	//this is just filenames for now. I will update test when we know the minimum required files for what we want to do
	expectedFileNames := []string{"agency.txt", "attributions.txt", "calendar.txt","shapes.txt","trips.txt","routes.txt","stops.txt","stop_times.txt","transfers.txt"}
	gtfs, err := filepath.Abs("../data/GTFS.zip")
	if err != nil {
		t.Error("filepath.Abs returned error: " + err.Error())
	}
	gtfsDir, err := unzipGTFS(gtfs)
	if err != nil {
		t.Error("unzipGTFS failed with error: " +  err.Error())
	}

	for _, name := range expectedFileNames {
		path := filepath.Join(gtfsDir, name)

		assert.FileExists(t, path)
	}

}