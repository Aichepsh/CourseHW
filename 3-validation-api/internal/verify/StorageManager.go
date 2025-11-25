package verify

import (
	"encoding/json"
	"os"
)

type Storage struct {
	File        string            // путь к JSON файлу
	EmailToCode map[string]string // email → code
	CodeToEmail map[string]string
}

func NewStorage(path string) *Storage {
	return &Storage{
		File:        path,
		EmailToCode: make(map[string]string),
		CodeToEmail: make(map[string]string),
	}
}

func (s *Storage) Load() error {
	bytes, err := os.ReadFile(s.File)
	if err != nil {
		if os.IsNotExist(err) {
			// создаём пустые карты
			s.EmailToCode = make(map[string]string)
			s.CodeToEmail = make(map[string]string)
			return nil
		}
		return err
	}

	if len(bytes) == 0 {
		// если файл пустой → тоже создаём пустые карты
		s.EmailToCode = make(map[string]string)
		s.CodeToEmail = make(map[string]string)
		return nil
	}

	// временная структура
	temp := struct {
		EmailToCode map[string]string `json:"email_to_code"`
		CodeToEmail map[string]string `json:"code_to_email"`
	}{}

	err = json.Unmarshal(bytes, &temp)
	if err != nil {
		return err
	}

	// если в файле этих карт нет → тоже создаём
	if temp.EmailToCode == nil {
		temp.EmailToCode = make(map[string]string)
	}
	if temp.CodeToEmail == nil {
		temp.CodeToEmail = make(map[string]string)
	}

	s.EmailToCode = temp.EmailToCode
	s.CodeToEmail = temp.CodeToEmail
	return nil
}

func (s *Storage) Save() error {
	temp := struct {
		EmailToCode map[string]string `json:"email_to_code"`
		CodeToEmail map[string]string `json:"code_to_email"`
	}{
		EmailToCode: s.EmailToCode,
		CodeToEmail: s.CodeToEmail,
	}

	bytes, err := json.MarshalIndent(temp, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.File, bytes, 0644)
}

func (s *Storage) Set(email, code string) error {
	if oldCode, ok := s.EmailToCode[email]; ok {
		delete(s.CodeToEmail, oldCode)
	}
	s.EmailToCode[email] = code
	s.CodeToEmail[code] = email
	s.Clean()
	return s.Save()
}

func (s *Storage) Clean() {
	// удаляем все коды, которые не соответствуют EmailToCode
	validCodes := make(map[string]bool)

	for email, code := range s.EmailToCode {
		validCodes[code] = true
		_ = email // ignore
	}

	for code := range s.CodeToEmail {
		if !validCodes[code] {
			delete(s.CodeToEmail, code)
		}
	}
}

func (s *Storage) GetCodeByEmail(email string) string {
	if code, ok := s.EmailToCode[email]; ok {
		return code
	}
	return ""
}
