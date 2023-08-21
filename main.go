package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var chapterMutex sync.Mutex     // Mutex for mutual exclusion when writing to allChapters
var allChapters strings.Builder // A string builder to store all chapters

// This function processes a paragraph, adds a chapter title, and writes to allChapters.
func processParagraph(paragraph string, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	// This creates a chapter for each paragraph with the paragraph content and index.
	chapter := fmt.Sprintf("Chapter %d:\n%s\n\n", index+1, paragraph)

	chapterMutex.Lock()
	defer chapterMutex.Unlock()
	allChapters.WriteString(chapter)
}

func main() {
	inputFilename := "/Users/dgodstand/documents/pegasus.txt"

	content, err := os.ReadFile(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	paragraphs := strings.Split(string(content), "\n\n") // If paragraphs are separated by empty lines.

	var wg sync.WaitGroup // This co-ordinates the completion of all goroutines.

	for i, paragraph := range paragraphs {
		wg.Add(1)                              // Increments the counter for each goroutine.
		go processParagraph(paragraph, i, &wg) // Starts a new goroutine to process the paragraph
	}

	wg.Wait() // Waits for all goroutines to finish

	// Writes all chapters to an output file
	outputFilename := "/Users/dgodstand/documents/chapters.txt"
	err = os.WriteFile(outputFilename, []byte(allChapters.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All chapters saved to ", outputFilename)
}
