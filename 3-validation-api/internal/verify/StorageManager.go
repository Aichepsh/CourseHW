package verify

import (
	"encoding/json"
	"os"
)

type Storage struct {
	File string // путь к JSON файлу
	Data map[string]string
}

func NewStorage(path string) *Storage {
	return &Storage{
		File: path,
		Data: make(map[string]string),
	}
}

func (s *Storage) Load() error {
	bytes, err := os.ReadFile(s.File)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(bytes, &s.Data)
}

func (s *Storage) Save() error {
	dataJSON, err := json.MarshalIndent(s.Data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.File, dataJSON, 0644)
}

func (s *Storage) Set(code, email string) error {
	s.Data[code] = email
	return s.Save()
}
