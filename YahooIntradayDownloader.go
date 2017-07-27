package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"strconv"
)

type YahooHeader struct {
	header []string
	urlargs []string
}

func NewYahooHeader() *YahooHeader {
	return &YahooHeader{
		header: []string{"Symbol", "Last", "Date", "Time", "Change", "Open", "High", "Low", "Volume"},
		urlargs: []string{"s","l1","d1","t1","c1","o","h","g","v"}}
}

func main(){
	args := os.Args[1:]

	if len(args) != 3 {
		log.Fatal("Usage: <tickerfile> <outputfile> <chunk_size>")
	}

	tickerFile := args[0]
	outputFile := args[1]
	chunkSize,err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	tickers := loadTickers(tickerFile)
	fmt.Println("tickers: ", tickers)

	header := NewYahooHeader()

	min := func(a, b int) int { if a < b { return a }; return b }
	resultChannel := make(chan string)
	nChunks := 0

	for i:=0; i<len(tickers); i+=chunkSize {
		nChunks++
		s := tickers[i:(i+min(chunkSize, len(tickers)-i))]

		go func(){
			resultChannel <- downloadPrices(header, s)
		}()
	}

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	outWriter := bufio.NewWriter(out)
	fmt.Fprintf(outWriter, "%v\n", strings.Join(header.header,","))

	for i:=0; i<nChunks; i++ {
		chunk := <- resultChannel
		fmt.Fprint(outWriter, chunk)
	}

	outWriter.Flush()
}

func loadTickers(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := make([]string, 0, 10000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticker := strings.TrimSpace(scanner.Text())
		result = append(result, ticker)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func downloadPrices(header *YahooHeader, tickers []string) string {
	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%v&f=sl1d1t1c1ohgv&e=.csv", strings.Join(tickers,","), strings.Join(header.urlargs,""))

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := string(bytes)
	return data
}
