package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	fmt.Println("hellow")
}

func IsPNG(r io.ReadSeeker) (bool, error) {
	// PNGファイルのマジックナンバー
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	_, err := io.ReadAtLeast(r, buf, len(buf))
	if err != nil {
		return false, err
	}

	// Seekすることで、読みはじめの位置にもどって、そのままimage.Decodeとかに渡してもエラーにならないようにする
	// ちなみに、io.SeekEndを使うと、tailのようにファイルの末尾から読みこむことができる
	_, err = r.Seek(0, io.SeekStart)
	if err != nil {
		return false, err
	}
	return bytes.Equal(magicnum, buf), nil
}

func NewPNG(r io.Reader) (io.Reader, error) {
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	if !bytes.Equal(magicnum, buf) {
		return nil, errors.New("not PNG file")
	}

	// MultiReaderを使って、Readerを連結することで、Seekを使わなくても、読み進めた分を付けたして完成品として返せる
	pngImg := io.MultiReader(bytes.NewReader(magicnum), r)
	return pngImg, nil
}

func Post(m string) (rerr error) {
	// Pipeを使うと、読み込んだものを、bytes介してメモリ上に書かなくても、そのまま書き込める
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		enc := json.NewEncoder(pw)
		err := enc.Encode(m)
		if err != nil {
			rerr = err
		}
	}()

	const url = "http://example.com"
	const contentType = "application/json"

	_, err := http.Post(url, contentType, pr)
	if err != nil {
		return err
	}
	return nil
}

func Walk(dir string, f func(b []byte, err error) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			b, err := ioutil.ReadFile(path)
			if err := f(b, err); err != nil {
				return err
			}
			continue
		}
		if err := Walk(path, f); err != nil {
			return err
		}
	}
	return nil
}
