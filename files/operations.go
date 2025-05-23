package files

import "os"

func Read(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Write(path string, content []byte, perm os.FileMode) error {
	return os.WriteFile(path, content, 0644)
}

func Append(path string, content []byte, perm os.FileMode) error {
	return os.WriteFile(path, content, perm)
}

func Delete(path string) error {
	return os.Remove(path)
}
