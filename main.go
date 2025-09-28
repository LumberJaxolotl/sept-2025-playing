package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/LeviathanTheGreat/ffmpeg-sept-2025-playing/lib"
)

// processVideo runs an ffmpeg command on a single input file
func processVideo(inputPath, outputPath string) error {
    // Build the ffmpeg command
    // Example: scale video to 1280x720
    cmd := exec.Command(
        "ffmpeg",
        "-i", inputPath,           // input file
        "-vf", "scale=1280:720",   // video filter to resize
        outputPath,                // output file
    )

    // Connect ffmpeg's output to the console so we can see progress/errors
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Run the command
    return cmd.Run()
}

func main() {
    lib.CreateDatedFolder()
	
	// // Check if the user provided an input folder
    // if len(os.Args) < 2 {
    //     fmt.Println("Usage: go run main.go <input-folder>")
    //     os.Exit(1)
    // }

    // // Get the folder path from command-line arguments
    // inputFolder := os.Args[1]

    // // Walk through the folder and process all .mp4 files
    // err := filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
    //     if err != nil {
    //         // Handle error while walking the folder
    //         return err
    //     }

    //     // Only process regular files with .mp4 extension
    //     if !info.IsDir() && filepath.Ext(path) == ".mp4" {
    //         // Create the output file path
    //         outputPath := filepath.Join(filepath.Dir(path), "processed_"+info.Name())

    //         // Print status to console
    //         fmt.Printf("Processing %s -> %s\n", path, outputPath)

    //         // Call FFmpeg on this file
    //         if err := processVideo(path, outputPath); err != nil {
    //             log.Printf("Failed to process %s: %v\n", path, err)
    //         }
    //     }

    //     // Continue walking the folder
    //     return nil
    // })

    // // Handle error if folder walking fails
    // if err != nil {
    //     log.Fatalf("Error walking folder: %v\n", err)
    // }

    fmt.Println("All done!")
}