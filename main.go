package main

import (
	"flag"
	"os"
	"fmt"
	"path"
	"log"
	"net/url"
	"time"
	"io/ioutil"
	"net/http"
	"encoding/csv"
	"github.com/ttacon/libphonenumber"
	"io"
)

var host = "https://api.zipwhip.com"
var api = "/messaging/send"
var session = ""
var concurrency = 5 // How many concurrent web requests should be used?

const (
	RECIPIENT = iota
	BODY
)

// Given a fileName, determine the path of the executable and open the file in the current directory.
func openFile(fileName string) (*os.File, error) {

	// Determine the path to where the application is running.
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := path.Dir(ex)

	// Open the file in the directory where the application is running and the fileName provided.
	return os.Open(fmt.Sprintf("%s/%s", exPath, fileName))
}

func sendMessages(file *os.File) {

	sem := make(chan bool, concurrency)

	// Read in each line of the file.
	r := csv.NewReader(file)

	for {
		record, err := r.Read() // Reach each line out of the file.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("An error occurred parsing input string: %s\n", err)
			continue
		}

		if len(record) < 2 {
			log.Println("An error occurred parsing input string: Not enough arguements.")
			continue
		}

		sem <- true // block until a worker is ready.

		go func(input []string) {
			defer func() { <-sem }()

			phoneNumber, err := libphonenumber.Parse(input[RECIPIENT], "US")
			if err != nil{
				// If we cannot parse the number, then we'll skip over the number.
				log.Printf("An error occurred parsing phoneNumber: %s\n", input[RECIPIENT])
				return
			}

			messageSend(libphonenumber.Format(phoneNumber, libphonenumber.E164), input[BODY])
		}(record)
	}

	// Let all in-flight requests finish.
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

// Sends a Zipwhip Message.
// Outputs the result and time to process the message to the log.
func messageSend(recipient string, body string) {

	var values = url.Values{}
	values.Add("session", session)
	values.Add("to", recipient)
	values.Add("body", body)

	t1 := time.Now()
	resp, err := http.PostForm(fmt.Sprintf("%s%s", host, api), values)
	t2 := time.Now()

	diff := t2.Sub(t1) // Time passed.

	response, _ := ioutil.ReadAll(resp.Body)

	log.Printf("'%s','%d','%d','%d'\n", response, t1.UnixNano(), t2.UnixNano(), diff.Nanoseconds()/int64(time.Millisecond))

	if err != nil {
		log.Fatal(fmt.Sprintf("An error occurred sending the message: %s\n", err))
	}

}

func main() {

	sessionInput := flag.String("session", "", "session for account that numbers will be provisioned to.")
	fileInput := flag.String("fileName", "", "File with messages to send, must be in same directory.")
	concurrencyInput := flag.Int("threads", 5, "Number of concurrent requests, defaults to 5.")

	flag.Parse()

	// If either apiKey or fileName are empty then exit.
	if *sessionInput == "" {
		panic("session must be provided.")
	}

	if *fileInput == "" {
		panic("fileName must be provided.")
	}

	// Set global state to use the provide apiKey and level of concurrency.
	session = *sessionInput
	concurrency = *concurrencyInput

	// Open the file in the directory where the application is running and the fileName provided.
	file, err := openFile(*fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.Println("Starting to Send Messages:")

	sendMessages(file)
	log.Println("Finished Sending Messages!")

}
