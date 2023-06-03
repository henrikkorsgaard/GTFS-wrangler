package main
 
import (
    "bufio"
    "fmt"
	"time"
    "os"
    "encoding/csv"
    "io"
)
 
func main() {
    readWithFileScanner()
    readWithCSVAll()
    readwithCSVRow()
}

func readWithFileScanner(){
    fmt.Println("\nReading file with filescanner")
    filePath := "../data/GTFS/stop_times.txt"
	start := time.Now()
    readFile, err := os.Open(filePath)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("time %d\n", diff.Milliseconds())
	fmt.Println(len(fileLines))
    fmt.Println("Done reading file with filescanner\n")
}

func readWithCSVAll(){

    fmt.Println("\nReading file with CSVAll")
    filepath := "../data/GTFS/stop_times.txt"
	start := time.Now()

    csvfile, err := os.Open(filepath)
	if err != nil {
		return
	}

    var fileLines [][]string

	r := csv.NewReader(csvfile)
	data, err := r.ReadAll()
    if err != nil {
        return
    }

    for _,row := range data {
        fileLines = append(fileLines, row)
    }
    
    csvfile.Close()
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("time %d\n", diff.Milliseconds())
	fmt.Println(len(fileLines))
    fmt.Println("Done reading file with CSVAll\n")
}

func readwithCSVRow(){
    fmt.Println("\nReading file with CSVRow")
    filepath := "../data/GTFS/stop_times.txt"
	start := time.Now()

    csvfile, err := os.Open(filepath)
	if err != nil {
		return
	}

    var fileLines [][]string

	r := csv.NewReader(csvfile)
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }

        if err != nil {
            panic(err)
        }
        fileLines = append(fileLines, record)
    }

    csvfile.Close()
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("time %d\n", diff.Milliseconds())
	fmt.Println(len(fileLines))
    fmt.Println("Done reading file with CSVRow\n")
}