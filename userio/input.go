package userio

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

func ReadInputNotBlank(msg string) (string, error) {
	var result string

	for len(result) == 0 {
		fmt.Printf("%s:", msg)
		fmt.Scanln(&result)
	}
	return result, nil
}

func GetParameterAsInt(args []string, index uint) (res int, ok bool) {
	if len(args) == 0 {
		return 0, false
	}

	if checkLength(args, index) {
		return 0, false
	}

	res, err := strconv.Atoi(args[index])
	if err != nil {
		return 0, false
	}
	return res, true
}

func checkLength(args []string, index uint) bool {
	return uint(len(args)-1) < index
}
