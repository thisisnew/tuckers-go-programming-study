package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("---------------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("---------------------------------------")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := filepath.Glob(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.err: ", err)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(fileList)
	recvCnt := 0

	for _, filename := range fileList {
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}

	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{
		filename: filename,
		lines:    []LineInfo{},
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		ch <- findInfo
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{
				lineNo: lineNo,
				line:   line,
			})
		}
		lineNo++
	}
	ch <- findInfo
}
