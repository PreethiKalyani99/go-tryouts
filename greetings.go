package main;

import (
	"fmt"
	"errors"
);

func Greet(name string) (string, error){
	if name == "" {
		return "", errors.New("Empty name");
	}

	return fmt.Sprintf("Hi %v, good morning!", name), nil;
}