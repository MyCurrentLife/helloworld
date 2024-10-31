package order

import (
	"os"
)

func getBytesFromFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()

	bytesFile := make([]byte, filesize)
	_, err = file.Read(bytesFile)
	if err != nil {
		return nil, err
	}
	return bytesFile, nil
}
func writeTextInFile(name string, b []byte) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}
