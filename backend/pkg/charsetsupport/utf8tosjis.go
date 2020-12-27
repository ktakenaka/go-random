package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// covert utf8 to sjis
// whwn a letter doesn't exist in sjis, like 〜, it will be converted to substitution of `?`

// inputとoutputなににしよかなー
// ユースケースから考えるか
// あとどうやってその文字が使えるかどうかdetectするべきか

// Convertロジック "〜
// これらはUnicodeのコードポイントでは？16進数の表現とは別物？
// from_chr = "\u{301C 2212 00A2 00A3 00AC 2013 2014 2016 203E 00A0 00F8 203A}"
// to_chr   = "\u{FF5E FF0D FFE0 FFE1 FFE2 FF0D 2015 2225 FFE3 0020 03A6 3009}"
// str&.tr(from_chr, to_chr)

// 上記以外は ? に変換

// string => byte => 置換 => byte
// reader in utf8 => writer in sjis

// 例外文字以外を変換できる ([]byteを受け取って、標準出力に出す => ここはnkfで変換しよう)

// io.Pipeをつかって、そのまま標準出力に書き込む

//

func main() {
	out := os.Stdout
	defer out.Close()

	w := transform.NewWriter(out, japanese.ShiftJIS.NewEncoder())

	bfw := bufio.NewWriter(w)
	defer bfw.Flush()

	testString := "〜−¢£¬–—‖‾ ø›"
	newString := replaceToAvailableLetters([]byte(testString))
	_, err := bfw.Write(newString)
	if err != nil {
		fmt.Println(err)
	}
}

func replaceToAvailableLetters(targetBytes []byte) []byte {
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("〜"), []byte("～"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("−"), []byte("－"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("¢"), []byte("￠"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("£"), []byte("￡"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("¬"), []byte("￢"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("–"), []byte("－"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("—"), []byte("―"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("‖"), []byte("∥"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("‾"), []byte("￣"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("ø"), []byte("Φ"))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte(" "), []byte(" "))
	targetBytes = bytes.ReplaceAll(targetBytes, []byte("›"), []byte("〉"))
	return targetBytes
}
