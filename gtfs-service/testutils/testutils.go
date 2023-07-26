package testutils

import (
	"path/filepath"
	"io"
	"os"
	"bufio"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/jackc/pgx/v5/pgxpool"
	
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
// TODO: Replace with standard db package
func ResetDatabase(sqlFilePath string) (err error){
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,pass,host,port,name)
	pool, err := pgxpool.New(context.Background(), connectionString)
	defer pool.Close()
	if err != nil {
		return
	}

	file, err := ioutil.ReadFile(sqlFilePath)
    if err != nil {
		return
    }

	_, err = pool.Exec(context.Background(), string(file))
	return
}


