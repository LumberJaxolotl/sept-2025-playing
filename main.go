package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LeviathanTheGreat/ffmpeg-sept-2025-playing/lib"
)


func main() {
    
	
	// Check if the user provided an input folder
    if len(os.Args) < 2 {
        fmt.Println("Error: Missing Command Argument")
		fmt.Println("Usage: go run main.go <input-folder-name>")
        os.Exit(1)
    }

	// Get the folder path from command-line arguments
    inputFolder := os.Args[1]
	// get all file paths of videos from input folder
	filePaths, filePathsErr := lib.GetFilePathsFromInputFolder(inputFolder)

	if filePathsErr != nil {
		log.Fatalf("Error getting file paths: %v\n", filePathsErr)
	}

	// create a unique output folder
	outputFolderPath, outputFolderErr := lib.CreateDatedFolder()
	if outputFolderErr != nil {
        log.Fatalf("Error creating output folder: %v\n", outputFolderErr)
    }

	filesProcessed := 0
	clipsCreated := 0

	for _, filePath := range filePaths {
		clipCount, creatingClipErr := lib.CreateRandomClips(filePath, outputFolderPath, 3, 1)
		if creatingClipErr != nil {
			log.Fatalf("Error creating clips from source video: %v\n %v", filePath, creatingClipErr)
		}

		clipsCreated += clipCount 
		filesProcessed++
	}

    lib.PrintEndOfProcessMessage(outputFolderPath)
    
}