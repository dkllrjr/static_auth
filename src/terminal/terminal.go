package terminal

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "syscall"
    "golang.org/x/crypto/ssh/terminal"
    "golang.org/x/crypto/bcrypt"
)

// ══════════════════════════════════════════════════════════════════════════

var err error

// ══════════════════════════════════════════════════════════════════════════

func EnterPassword () string {

    var password, confirm_password = []byte{102, 111, 111}, []byte{98, 97, 114}

    for string(password) != string(confirm_password) {

        fmt.Print("Enter password: ")
        password, err = terminal.ReadPassword(int(syscall.Stdin))

        fmt.Print("\nConfirm password: ")
        confirm_password, err = terminal.ReadPassword(int(syscall.Stdin))

        fmt.Println()

        if string(password) != string(confirm_password) {
            fmt.Println("Passwords did not match, please try again.")
        }
    }

    return string(password)
}

func EnterUsername () string {

    var scanner *bufio.Scanner
    var username string

    fmt.Print("Enter username: ")
    scanner = bufio.NewScanner(os.Stdin)
    scanner.Scan()
    err = scanner.Err()
    if err != nil {
        log.Fatal(err)
    }
    username = scanner.Text()

    return username
}


func RegisterUser() (string, string) {

    username := EnterUsername()
    password := EnterPassword()
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 4)

    return username, string(hashedPassword)
}
