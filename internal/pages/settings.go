package pages

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type FolderRequest struct {
	Folder string `json:"folder"`
}

func RenderSettingsPage(ctx *fiber.Ctx) error {
	settingData := fiber.Map{}
	return ctx.Render("layouts/settings", settingData)
}

func GetFilesMessage(ctx *fiber.Ctx) error {
	message := "You dont have any storage assigned yet"
	return ctx.SendString(message)
}

func FolderLists(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)
	userDocRootPath := defaultUserDocPath() + "/" + username
	isFolderExists := checkFolderExists(userDocRootPath)

	if !isFolderExists {
		createFolder(userDocRootPath)
	}

	folders, err := GetFolders(userDocRootPath)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error getting folders")
	}
	dropdownHTML := "<select class='dropdown-btn' 'id='folderSelect' name='folder'><option value='' selected disabled>Select Folder</option>"
	for _, folder := range folders {
		dropdownHTML += "<option value='" + folder + "'>" + folder + "</option>"
	}
	dropdownHTML += "</select>"

	return ctx.SendString(dropdownHTML)
}

func GetFolders(folderPath string) ([]string, error) {
	var folders []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			print(err)
			return err
		}
		if info.IsDir() {
			folders = append(folders, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return folders, nil
}

func ProcessFolders(ctx *fiber.Ctx) error {
	selectedFolder := ctx.FormValue("folder")

	// Return the selected folder in the response
	return ctx.SendString("Selected folder: " + selectedFolder)
}

func checkFolderExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return true
	}

	return true
}

func defaultUserDocPath() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading in .env file")
	}

	return os.Getenv("DEFAULT_USER_PATH")
}

func createFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	return err
}
