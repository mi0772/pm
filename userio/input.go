package userio

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword(msg string) ([]byte, error) {
	fmt.Printf("%s:", msg)
	mPasswordSlice, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}
	return mPasswordSlice, nil
}

func ReadNewMasterPassword(msg string) ([]byte, error) {
	fmt.Printf("%s:", msg)
	var err error
	var pwd1, pwd2 []byte

	for !bytes.Equal(pwd1, pwd2) || len(pwd1) == 0 {
		fmt.Print("\nenter new password:")
		pwd1, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			return nil, err
		}
		fmt.Print("\ntype it again:")
		pwd2, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			return nil, err
		}
	}

	return pwd1, nil
}

func ReadInput(msg string) (string, error) {
	fmt.Printf("%s:", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
