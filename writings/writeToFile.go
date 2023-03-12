package writings

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
)

func WriteToFile(fileName, content string) {

	var AppFs = afero.NewOsFs()

	create, err := AppFs.Create(fmt.Sprintf("%v", fileName))
	if err != nil {
		log.Fatalf("failed to create file %v with error: \n%v", fileName, err.Error())
		return
	}

	writeString, err := create.WriteString(content)
	if err != nil {
		log.Fatalf(">>> ERROR >>> failed to write content to file %v", create.Name())
		return
	}

	log.Println(fmt.Sprintf("print result: %v", writeString))

}

func WriteToFileWithDirCheck(pathWithFile, pathWithNoFile string) {
	log.Printf(fmt.Sprintf("path with file: %v", pathWithFile))
	log.Printf(fmt.Sprintf("path without file: %v", pathWithNoFile))
	if _, err := os.Stat(pathWithFile); os.IsNotExist(err) {
		err := os.MkdirAll(pathWithNoFile, 0777)
		if err != nil {
			log.Fatalf("failed to create directory: %v", err.Error())
			return
		} // Create your file
	}
}
