// cmd/get_table/main.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"tov_tools/pkg/character"
	"tov_tools/pkg/static_data"
)

const baseURL = "http://localhost:8080/api/v1/table/get"

func main() {
	// Define the -type command-line flag
	dataType := flag.String("type", "", "The type of data to retrieve (e.g., 'class')")
	flag.Parse()

	params := url.Values{}
	params.Add("type", *dataType)

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

	switch *dataType {
	case "class":
		printClassTable(resp.Body)
	case "damageType":
		printDamageTypeTable(resp.Body)
	case "damageModifier":
		printDamageModifierTable(resp.Body)
	default:
		fmt.Printf("Error: unsupported type '%s'\n", *dataType)
		os.Exit(1)
	}

}

// printClassTable prints out the class data in a text table
func printClassTable(respBody io.Reader) {
	var classResponse map[string]character.Class
	if err := json.NewDecoder(respBody).Decode(&classResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	prettyPrint(classResponse)
}

// printDamageTypeTable prints out the damage type data in a text table
func printDamageTypeTable(respBody io.Reader) {
	var damageTypeResponse static_data.DamageTypeTable
	if err := json.NewDecoder(respBody).Decode(&damageTypeResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	prettyPrint(damageTypeResponse)
}

// printDamageModifierTable prints out the damage modifier data in a text table
func printDamageModifierTable(respBody io.Reader) {
	var damageModifierResponse static_data.DamageModifierTable
	if err := json.NewDecoder(respBody).Decode(&damageModifierResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}
	prettyPrint(damageModifierResponse)
}

// prettyPrint prints a JSON object in a pretty format
func prettyPrint(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Println(string(prettyJSON))
}
