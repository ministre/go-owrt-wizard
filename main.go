package main

import (
    "encoding/json"
    "fmt"
    "os"
    "golang.org/x/crypto/ssh"
)

type Config struct {
    Ip          string
    Username    string
    Password    string
    Actions     Actions
}

type Actions struct {
    Get_model       bool
    Get_board_name  bool
    Get_sn          bool
}

func main() {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    config := new(Config)
    err := decoder.Decode(&config)
    if err != nil {
        panic(err)
    }

    ssh_config := &ssh.ClientConfig{
        User: config.Username,
        Auth: []ssh.AuthMethod{
            ssh.Password(config.Password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    fmt.Println(config.Actions.Get_model)

    fmt.Println("Trying SSH connection: " + config.Username + "@" + config.Ip)

    client, err := ssh.Dial("tcp", config.Ip + ":22", ssh_config)
    if err != nil {
        panic(err)
    }

    if config.Actions.Get_model {
        fmt.Print("Model: ")
        session, err := client.NewSession()
        if err != nil {
            panic(err)
        }
        defer session.Close()
        b, err := session.CombinedOutput("ubus call system board | grep model | cut -d \\\" -f 4\n")
        if err != nil {
            panic(err)
        }
        fmt.Print(string(b))
    }

    if config.Actions.Get_board_name {
        fmt.Print("Board Name: ")
        session, err := client.NewSession()
        if err != nil {
            panic(err)
        }
        defer session.Close()
        b, err := session.CombinedOutput("ubus call system board | grep board_name | cut -d \\\" -f 4\n")
        if err != nil {
            panic(err)
        }
        fmt.Print(string(b))
    }

    if config.Actions.Get_sn {
        fmt.Print("Serial Number: ")
        session, err := client.NewSession()
        if err != nil {
            panic(err)
        }
        defer session.Close()
        b, err := session.CombinedOutput("ubus call system board | grep serial | cut -d \\\" -f 4\n")
        if err != nil {
            panic(err)
        }
        fmt.Print(string(b))
    }


}