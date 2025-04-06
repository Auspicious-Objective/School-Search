package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"

	"schoolSearch/libs/api"
	"schoolSearch/libs/fileWorker"
)

// add API key here
var APIKey string = ""

func getAPIKey() string {
	if len(APIKey) > 0 {
		return APIKey
	} else {
		keyPrompt := promptui.Prompt{
			Label: "Enter US Department of Education API key:",
			Validate: func(input string) error {
				if input == "" {
					return fmt.Errorf("API key cannot be empty")
				}
				return nil
			},
		}

		// Run the prompt and capture the result
		APIKeyValue, err := keyPrompt.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Prompt failed: %v\n", err)
			os.Exit(1)
		}

		return APIKeyValue
	}
}

func getFileName() string {
	filePrompt := promptui.Prompt{
		Label: "Enter the name of the file",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("file name cannot be empty")
			}
			return nil
		},
	}

	// Run the prompt and capture the result
	fileName, err := filePrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed: %v\n", err)
		os.Exit(1)
	}

	fileName = strings.ReplaceAll(fileName, ".csv", "")
	fileName = fileName + ".csv"
	return fileName

}

func main() {
	responses := []string{}

	APIKey = getAPIKey()

	fmt.Println("Hello! Welcome to the program.")

	for {
		// Configure the prompt
		prompt := promptui.Prompt{
			Label: "Enter The Schools Name ->",
			Templates: &promptui.PromptTemplates{
				Prompt: "{{ . | faint }} ",
				Valid:  "{{ . | green }} ",
			},
			Validate: func(input string) error {
				if input == "" {
					return fmt.Errorf("input cannot be empty")
				}
				return nil
			},
			Default: "    enter school name or use /e to submit or use /h for more options",
		}

		// Run the prompt and get input
		result, err := prompt.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Prompt failed %v\n", err)
			os.Exit(1)
		}

		// Handle the input
		if result == "/e" || result == "/enter" {
			fmt.Println("Processing your schools.")
			break
		}

		if result == "/q" || result == "/quit" {
			fmt.Println("You have exited the program.")
			os.Exit(1)
		}

		if result == "/help" || result == "/h" {
			fmt.Println("Commands: \n/e or /enter:    Used to end the session and submit the schools for file creation. \n/q or /quit:    Used to exit the program without executing.")
		}

		if result != "    enter school name or use /e to submit or use /h for more options" {
			result = strings.ToLower(result)
			result = strings.ReplaceAll(result, " ", "%20")

			if len(responses) > 0 {
				duplicate := false
				for _, response_check := range responses {
					if result == response_check {
						fmt.Println("The school you entered has already been added to your list, please add a different one or use /help for more options.")
						duplicate = true
						break
					}
				}
				if !duplicate {
					responses = append(responses, result)
				}
			} else {
				responses = append(responses, result)
			}
		}
	}

	var nameOfFile string = getFileName()
	fileWorker.GenerateFile(nameOfFile)

	fmt.Println("Generating file this may take a moment. Please be patient.")

	for _, response := range responses {
		collegeInfo := api.CallAPI(response, APIKey)
		fileWorker.AddDataToFile(nameOfFile, collegeInfo)
	}
}
