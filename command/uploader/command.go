package uploader

import (
	"fmt"
	"os"
)

type UploadCommand struct{}

// Synopsis - Used for command line stuff
func (uc UploadCommand) Synopsis() string {
	return "Uploads a file"
}

// Help - Help text
func (uc UploadCommand) Help() string {
	helpText := `
Usage: command-and-control upload dirname filename ip

	Uploads filename.txt and filename.jpg from dirname/ using IP ip
`
	return helpText
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func (uc UploadCommand) Run(args []string) int {

	if len(args) != 3 {
		fmt.Println(uc.Help())
		return 1
	}

	dir := args[0]
	file := args[1]
	ip := args[2]

	if !IsDir(dir) {
		fmt.Printf("Dir doesn't exist: %s\n", dir)
		return 2
	}

	f := NewFileInfo(dir, file)

	if !IsFile(f.GetImagePath()) {
		fmt.Printf("File doesn't exist: %s\n", f.GetImagePath())
		return 2
	}

	if !IsFile(f.GetTextPath()) {
		fmt.Printf("File doesn't exist: %s\n", f.GetTextPath())
		return 2
	}

	fmt.Printf("Starting upload of %s using IP %s\n", file, ip)
	err := Upload(f, ip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Finished upload of %s using IP %s\n", file, ip)
	return 0
}
