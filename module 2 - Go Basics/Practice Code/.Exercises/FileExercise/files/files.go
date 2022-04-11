package files

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	file *os.File
	err  error
)

type Files struct {
	FilePath string
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func (s *Files) CreateFile() {
	file, err = os.Create(s.FilePath)

	check(err)

	fmt.Println("File ", s.FilePath, "created succesful")

	s.CloseFile()
}

func (s *Files) CloseFile() {
	file.Close()
	fmt.Println("File closed")
}

func (s *Files) RenameFile(newName string) {
	os.Rename(s.FilePath, newName)
	s.FilePath = newName
	fmt.Println("File name changed to", newName)
}

func (s *Files) RemoveFile() {
	os.Remove(s.FilePath)
	fmt.Println("file", s.FilePath, " was removed")
}

func (s *Files) EditFile(user, password string) {
	file, err = os.OpenFile(s.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	sliceBytes := []byte("\n User: " + user + "\n" + "Password: " + base64.StdEncoding.EncodeToString([]byte(password)) + "\n")
	_, err = file.Write(sliceBytes)

	check(err)

	fmt.Println("User ", user, "and password", password, "were added to file ", s.FilePath)

	s.CloseFile()

}

func (s *Files) ReadFile() []byte {
	file, err := os.Open(s.FilePath)

	defer s.CloseFile()

	check(err)

	data, e := ioutil.ReadAll(file)
	check(e)

	return data
}