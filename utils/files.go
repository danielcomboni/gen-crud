package utils

import (
	"fmt"
	"log"
	"os"
)

// DoesFileExist function to check if file exists
func DoesFileExist(fileName string) bool {
	log.Println("check file existence: " + fileName)
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		log.Fatal("file does not exist")
		return false
	}
	log.Println("file exist")
	return true
}

// DoesDirectoryExist check if directory exists
// returns a bool
func DoesDirectoryExist(dirPath string) bool {
	log.Println("check directory existence: " + dirPath)
	_, err := os.Stat(dirPath)
	var flag = true
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("directory does not exist: %v", dirPath)
		} else {
			log.Println(fmt.Sprintf("directory exists: %v", dirPath))
			flag = true
		}
	}
	if flag {
		log.Println(fmt.Sprintf("directory `%v` exists", dirPath))
	}
	return flag
}

// WriteBase64Image writes the incoming base64 image to a file
// filename is imagedata.txt
func WriteBase64Image(directoryWithoutFileName, base64Image string) {
	f, err := os.Create(directoryWithoutFileName + "/imagedata.txt")
	if err != nil {
		log.Fatalf("base64 image data not saved to file")
	} else {
		log.Println(fmt.Sprintf("writing image base64 to file: %v", directoryWithoutFileName))
		_, err := f.WriteString(base64Image)
		if err != nil {
			log.Fatalf(">>> ERROR >>> problem occurred during file creation")
			return
		}
		err = f.Close()
		if err != nil {
			log.Fatalf(">>> ERROR >>> problem occurred when closing file creation")
			return
		}
	}
}

func WriteFile(directoryWithoutFileName, fileName, content string) {
	//f, err := os.Create(directoryWithoutFileName + "/imagedata.txt")

	f, err := os.Create(fmt.Sprintf("%v/%v", directoryWithoutFileName, fileName))

	if err != nil {
		//log.Fatalf("base64 image data not saved to file")
		log.Fatalf("failed to create file: %v\n%v", fileName, err.Error())
	} else {
		log.Println(fmt.Sprintf("writing struct file package: %v", directoryWithoutFileName))
		_, err := f.WriteString(content)
		if err != nil {
			log.Fatalf(">>> ERROR >>> problem occurred during file creation")
			return
		}
		err = f.Close()
		if err != nil {
			log.Fatalf(">>> ERROR >>> problem occurred when closing file creation")
			return
		}
	}
}

// deletes folder containing the user image in base64 string format
// it should only be used whenever a user is being deleted
//func DeleteUserImageFolder(directoryWithoutFileName string) {
//	err := os.RemoveAll(directoryWithoutFileName)
//	if err != nil {
//		Logger.Error("user image folder was not deleted: " + directoryWithoutFileName)
//		Logger.Error("error message: " + err.Error())
//	} else {
//		Logger.Error("user image folder was successfully deleted: " + directoryWithoutFileName)
//	}
//}
