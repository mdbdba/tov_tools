package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"tov_tools/pkg/dice"
)

const baseURL = "http://localhost:8080/api/v1/dice/roll"

func main() {
	sides := flag.Int("sides", 20, "number of sides on the dice")
	times := flag.Int("times", 1, "number of times to roll the dice")
	options := flag.String("options", "", "comma-separated options for the roll (e.g., 'advantage,disadvantage,drop lowest')")
	ctxRef := flag.String("ctx", "default", "context reference for the roll")
	full := flag.Bool("full", false, "print the full Roll object")

	flag.Parse()

	var opts []string
	if *options != "" {
		opts = strings.Split(*options, ",")
	}

	// Build the query parameters
	params := url.Values{}
	params.Add("sides", fmt.Sprintf("%d", *sides))
	params.Add("timesToRoll", fmt.Sprintf("%d", *times))
	if len(opts) > 0 {
		params.Add("options", strings.Join(opts, ","))
	}
	params.Add("ctxRef", fmt.Sprintf("%s", *ctxRef))

	// Complete URL
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-200 response: %d", resp.StatusCode)
	}

	var rollResponse dice.Roll
	if err := json.NewDecoder(resp.Body).Decode(&rollResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	if *full {
		fmt.Println(rollResponse.ToPrettyString())
	} else {
		fmt.Printf("%d\n", rollResponse.Result)
	}
}
