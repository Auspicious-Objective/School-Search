package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Combines all necessary info into the singular API url
func makeUrl(collegeName string, apiKey string) string {
	var baseUrl string = "https://api.data.gov/ed/collegescorecard/v1/schools.json?_fields=school.name,school.city,school.state,latest.admissions.admission_rate_suppressed.overal,latest.cost.avg_net_price.public,latest.cost.avg_net_price.private,latest.student.demographics.student_faculty_ratio,latest.student.size,latest.admissions.test_requirements&school.name="
	var apiKeyLink string = apiKey
	var url string = baseUrl + collegeName + apiKeyLink
	return url
}

// Makes the API call and parses the data into a map
func getAPIResponse(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer resp.Body.Close()

	buffer := make([]byte, 1024)

	var body string
	for {
		n, err := resp.Body.Read(buffer)
		if err == io.EOF {
			// Reached the end of the body
			body += string(buffer[:n])
			break
		} else if err != nil {
			fmt.Println("Error reading response:", err)
		}
		body += string(buffer[:n])
	}

	var result map[string]interface{}

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Fatal(err)
	}

	results, ok := result["results"].([]interface{})
	if !ok {
		fmt.Println("Error:", "Could not cast 'results' to a slice of interfaces")
	}

	if len(results) == 0 {
		fmt.Println("No results found in the response")
	}

	result = results[0].(map[string]interface{})
	return result
}

// extracts the name of the school from the api call
func getSchoolName(data map[string]interface{}) string {
	var schoolName string = data["school.name"].(string)
	return schoolName
}

// extracts the location of the school from the api call
func getSchoolLocation(data map[string]interface{}) string {
	var location string
	var state string = data["school.state"].(string)
	var city string = data["school.city"].(string)
	location = city + ", " + state
	return location
}

// extracts the acceptance rate of the school from the api call
func getAcceptanceRate(data map[string]interface{}) string {
	var acceptanceRate string
	if data["latest.admissions.admission_rate_suppressed.overall"] == nil {
		acceptanceRate = "N/A"
	} else {
		acceptanceRate = strconv.FormatInt(int64((data["latest.admissions.admission_rate_suppressed.overall"].(float64) * 100)), 10)
		acceptanceRate = acceptanceRate + "%"
	}
	return acceptanceRate
}

// extracts the cost of the school from the api call
func getCost(data map[string]interface{}) string {
	var cost string
	if data["latest.cost.avg_net_price.private"] == nil {
		cost = strconv.FormatInt(int64(data["latest.cost.avg_net_price.public"].(float64)), 10)
	} else {
		cost = strconv.FormatInt(int64(data["latest.cost.avg_net_price.private"].(float64)), 10)
	}
	cost = "$" + cost
	return cost
}

// extracts the student to teacher ratio of the school from the api call
func getStudentToTeacherRatio(data map[string]interface{}) string {
	var ratio string = strconv.FormatInt(int64(data["latest.student.demographics.student_faculty_ratio"].(float64)), 10)
	ratio = ratio + ":1"
	return ratio
}

// extracts the undergrad population of the school from the api call
func getTotalUndergrads(data map[string]interface{}) string {
	var amount string = strconv.FormatInt(int64(data["latest.student.size"].(float64)), 10)
	return amount
}

// extracts the testing information of the school from the api call
func getIsTestRequired(data map[string]interface{}) string {
	var Test string
	if data["latest.admissions.test_requirements"].(float64) == 0 {
		Test = "Standardized tests are not required"
	} else if data["latest.admissions.test_requirements"].(float64) == 1 {
		Test = "Standardized tests are required for attendance"
	} else if data["latest.admissions.test_requirements"].(float64) == 2 {
		Test = "Standardized tests are not required but are recommended"
	} else if data["latest.admissions.test_requirements"].(float64) == 3 {
		Test = "Standardized tests are not required nor are they recommended"
	} else if data["latest.admissions.test_requirements"].(float64) == 4 {
		Test = "Testing data is unavailable"
	} else if data["latest.admissions.test_requirements"].(float64) == 5 {
		Test = "Standardized tests are considered but not required"
	} else {
		log.Fatal("The returned value was not valid! ", data["latest.admissions.test_requirements"].(float64))
	}
	return Test
}

// compiles all of the extracted data into a slice that can then be returned for the info to be written to the file
func getCompileData(data map[string]interface{}) []string {
	var name string = getSchoolName(data)
	var location string = getSchoolLocation(data)
	var acceptanceRate string = getAcceptanceRate(data)
	var cost string = getCost(data)
	var ratio string = getStudentToTeacherRatio(data)
	var undergradPopulation string = getTotalUndergrads(data)
	var testingOptions string = getIsTestRequired(data)
	// compile data to one string group e.g. list slice etc and have it exported and acsessible by CLI schools program and have the data passed into file appending program
	compiledData := []string{name, location, acceptanceRate, cost, ratio, undergradPopulation, testingOptions}
	return compiledData
}

// A function to combine all other functions that is the only outward facing one that will be called be the main file
func CallAPI(school string, apiKey string) []string {
	var url = makeUrl(school, apiKey)
	data := getAPIResponse(url)
	compiledData := getCompileData(data)
	return compiledData
}
