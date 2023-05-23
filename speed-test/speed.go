package main
 
import (
    "bufio"
    "fmt"
	"time"
    "os"
)
 
func main() {
 
    filePath := "../input/GTFS/stop_times.txt"
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
}