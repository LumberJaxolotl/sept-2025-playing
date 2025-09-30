package lib

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func CheckFfmpegExists(){
	cmd := exec.Command("ffmpeg", "-version")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("ffmpeg command is NOT available or an error occurred: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		return
	}
}

func CreateDatedFolder() (string, error) {
	baseDir := "output"
	dateStr := time.Now().Format("2006-01-02") // YYYY-MM-DD

	// Ensure the base output directory exists
	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		return "", err
	}

	// Start with suffix 1 and increment if folder exists
	suffix := 1
	var folderPath string
	for {
		folderName := fmt.Sprintf("%s - %d", dateStr, suffix)
		folderPath = filepath.Join(baseDir, folderName)

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			// Folder does not exist, safe to create
			if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
				return "", err
			}
			break
		}
		suffix++
	}

	return folderPath, nil
}

func GetFilePaths(inputFolderOrFile string) ([]string, error){
	info, err := os.Stat(inputFolderOrFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Path does not exist")
		} else {
			fmt.Println("Error:", err)
		}
	
	}

	if info.IsDir() {
		res, err := getFilePathsFromInputFolder(inputFolderOrFile)
		return res, err
	} else {
		return getFilePathsFromInputFile(inputFolderOrFile), nil
	}
} 

func getFilePathsFromInputFolder(inputFolder string) ([]string, error) {
	videoExts := []string{
		"3g2",
		"3gp",
		"amv",
		"asf",
		"avi",
		"avs",
		"bik",
		"cavs",
		"divx",
		"drc",
		"dv",
		"dvr-ms",
		"f4v",
		"flv",
		"gxf",
		"ismv",
		"ivf",
		"m1v",
		"m2t",
		"m2ts",
		"m2v",
		"m4v",
		"mkv",
		"mov",
		"mp2",
		"mp4",
		"mp4v",
		"mpe",
		"mpeg",
		"mpg",
		"mpl",
		"mpv",
		"mxf",
		"nsv",
		"nut",
		"ogm",
		"ogv",
		"ps",
		"qt",
		"rm",
		"rmvb",
		"roq",
		"rpl",
		"ts",
		"vob",
		"webm",
		"wm",
		"wmv",
		"wtv",
		"yuv",
	}

	var filePaths []string

	err := filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extWithPeriod := filepath.Ext(path)
			ext := extWithPeriod[1:]
			for _, vExt := range videoExts {
				if ext == vExt {
					absPath, absErr := filepath.Abs(path)
					if absErr != nil {
						return absErr
					}
					filePaths = append(filePaths, absPath)
					break
				}
			}
		}
		return nil
	})
	return filePaths, err
}

func getFilePathsFromInputFile(inputFile string) []string {
    return []string{inputFile}
}

// CreateRandomClips generates n random clips of given length (seconds) from a video.
// If clipLength <= 0, it defaults to 0.5s.
func CreateRandomClips(inputPath string, outputFolder string, numOfClips int, clipLength float64) (int, error) {
	rand.Seed(time.Now().UnixNano())

	if clipLength <= 0 {
		clipLength = 0.5
	}

	// Get video duration using ffprobe
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		inputPath)

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to probe video: %w", err)
	}

	var duration float64
	_, err = fmt.Sscanf(string(output), "%f", &duration)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %w", err)
	}

	if duration <= clipLength {
		return 0, fmt.Errorf("video too short for %.2fs clips", clipLength)
	}

	// Prepare base name for output clips
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := base[:len(base)-len(ext)]

	for i := 0; i < numOfClips; i++ {
		// Pick a random start time
		start := rand.Float64() * (duration - clipLength)
		outputFile := fmt.Sprintf("%s\\%s_clip_%d%s", outputFolder,  name, i+1, ext)

		clipCmd := exec.Command("ffmpeg",
			"-y", // overwrite
			"-ss", fmt.Sprintf("%.2f", start),
			"-i", inputPath,
			"-t", fmt.Sprintf("%.2f", clipLength),
			outputFile,
		)

		if err := clipCmd.Run(); err != nil {
			return 0, fmt.Errorf("failed to create clip %d: %w", i+1, err)
		}

		log.Printf("Created clip: %s (start=%.2fs, length=%.2fs)", outputFile, start, clipLength)
	}

	return numOfClips, nil
}


func PrintEndOfProcessMessage(outputFolderPath string){
	
	
	
	path, err := filepath.Abs(outputFolderPath)
	if err != nil {
		log.Printf("%v", err)
	}


	fmt.Println("All done!")
	fmt.Println(" ")
	fmt.Println("Check out the new clips at: " + path)
}