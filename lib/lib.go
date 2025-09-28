package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

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


    