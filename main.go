package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

var filename = fmt.Sprintf("resume_%v", time.Now().Unix())
var templatefilename = fmt.Sprintf("%s.tex", filename)

func main() {

	// process the json file
	r, err := jsonToStruct()
	if err != nil {
		log.Fatal(err)
	}

	generateLaTeX(r)

	err = runPdflatex()
	if err != nil {
		log.Fatal("Failed to generate PDF:", err)
	}

	removeLaTeXFiles()

}

func jsonToStruct() (Resume, error) {
	var resume Resume

	if len(os.Args) < 2 {
		return Resume{}, errors.New("No filename provided.")
	}

	/* header */
	headerFile, err := os.Open("header.json")
	if err != nil {
		return Resume{}, err
	}
	defer headerFile.Close()

	headerInfo, err := ioutil.ReadAll(headerFile)
	if err != nil {
		return Resume{}, nil
	}

	// Replace '%' with '\\%' in the header JSON data

	if err := json.Unmarshal(headerInfo, &resume); err != nil {
		return Resume{}, err
	}

	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		return Resume{}, err
	}
	defer file.Close()

	// Read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return Resume{}, err
	}

	if err := json.Unmarshal(data, &resume); err != nil {
		return Resume{}, err
	}

	return resume, nil
}

func generateLaTeX(resume Resume) {

	tmpl, err := template.New("template.tex").Delims("[[", "]]").ParseFiles("template.tex")
	if err != nil {
		panic(err)
	}

	// Your code to execute the template will go here
	// Open a new LaTeX file
	newFile, err := os.Create(templatefilename)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	// Execute the template and inject the Resume struct
	err = tmpl.Execute(newFile, resume)
	if err != nil {
		panic(err)
	}
}

// Add this function to your code
func runPdflatex() error {
	cmd := exec.Command("xelatex", templatefilename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func removeLaTeXFiles() {
	extensions := []string{".aux", ".tex", ".out", ".log"}

	for _, ext := range extensions {
		filename := filename + ext
		err := os.Remove(filename)
		if err != nil {
			fmt.Printf("Error removing %s: %s\n", filename, err)
			return
		} else {
			fmt.Printf("Successfully removed %s\n", filename)
		}
	}
}
