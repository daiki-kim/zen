package main

import "errors"

// ユーザーデータを検証する関数
func validateUser(name string, age int, email string) (bool, error) {
	if name == "" {
		return false, errors.New("name is required")
	}
	if age < 0 {
		return false, errors.New("age cannot be negative")
	}
	if email == "" {
		return false, errors.New("email is required")
	}
	return true, nil
}
