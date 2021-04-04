package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	byterange  = "bytes"
	rHeaderKey = "Range"
	oFileName  = "output.tar.gz"
)
const exampleURL = "https://wordpress.org/latest.tar.gz"

var (
	splitsize = runtime.NumCPU()
)
var (
	client  = http.DefaultClient
	basereq *http.Request
)

type resMap struct {
	mu sync.Mutex
	rm map[int]*bytes.Buffer
}

func (r *resMap) SetRM(i int, b *bytes.Buffer) {
	r.rm[i] = b
}

func (r *resMap) At(i int) *bytes.Buffer {
	return r.rm[i]
}

func main() {
	// curl -I https://wordpress.org/latest.tar.gz
	resp, err := headreq()
	if err != nil {
		log.Fatal(fmt.Errorf("head request: %w", err))
	}

	if ars, ok := resp.Header["Accept-Ranges"]; !ok || ars[0] != byterange {
		log.Fatal(errors.New("Doesn't accept bytes"))
	}

	var length int64
	if cls, ok := resp.Header["Content-Length"]; ok {
		length, err = strconv.ParseInt(cls[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Errorf("convert %s to int: %w", cls[0], err))
		}
	}

	r := resMap{rm: make(map[int]*bytes.Buffer)}
	for i := 0; i < splitsize; i++ {
		r.SetRM(i, &bytes.Buffer{})
	}

	chunksize := length / int64(splitsize)
	g := new(errgroup.Group)
	for i := 0; i < splitsize; i++ {
		i := i
		g.Go(func() error {
			req := makereq()
			rangeFrom := chunksize * int64(i)
			if i == (splitsize - 1) {
				req.Header.Set(rHeaderKey, fmt.Sprintf("bytes=%d-", rangeFrom))
				fmt.Printf("Getting bytes=%d-\n", rangeFrom)
			} else {
				req.Header.Set(rHeaderKey, fmt.Sprintf("bytes=%d-%d", rangeFrom, chunksize*(int64(i)+1)-1))
				fmt.Printf("Getting bytes=%d-%d\n", rangeFrom, chunksize*(int64(i)+1)-1)
			}

			return rangereq(req, r.At(i))
		})
	}
	if err := g.Wait(); err != nil {
		log.Fatal(fmt.Errorf("fail %w", err))
		return
	}

	file, err := os.Create(oFileName)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create a file: %w", err))
		return
	}
	defer file.Close()
	defer func() {
		if err != nil {
			os.Remove(oFileName)
		}
	}()

	for i := 0; i < splitsize; i++ {
		_, err = file.Write(r.At(i).Bytes())
		if err != nil {
			log.Fatal(fmt.Errorf("failed to write on a file: %w", err))
			return
		}
	}
}

func headreq() (resp *http.Response, err error) {
	resp, err = http.Head(exampleURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func makereq() *http.Request {
	req, err := http.NewRequest("GET", exampleURL, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func rangereq(req *http.Request, to *bytes.Buffer) error {
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("get range from %s: %w", req.Header.Get(rHeaderKey), err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(to, resp.Body)
	return err
}
