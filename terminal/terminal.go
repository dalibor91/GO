package terminal

import (
  "os"
  "bufio"
  "fmt"
  "strings"
  "regexp"
  //"errors"
)

type Terminal struct {
  commands []string
  hello string
  cursor string
  quitCommand string
  quitText string
  started bool
}


func NewTerminal() *Terminal {
  s := new(Terminal)
  s.started = false
  s.quitText = "Bye!"
  s.quitCommand = "quit"
  return s
}

func (t *Terminal) SetHelloText(text string) *Terminal {
  t.hello = text
  return t
}

func (t *Terminal) SetCursor(text string) *Terminal {
  t.cursor = text
  return t
}

func (t *Terminal) SetQuitCommand(text string) *Terminal {
  t.quitCommand = text
  return t
}

func (t *Terminal) Start() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(t.hello)

  if t.started {
    fmt.Println("Already started")
    return
  }

  reg, err := regexp.Compile("\\s+")

  if err != nil {
    fmt.Println("Unable to compile regex for whitespaces")
    return
  }

  fmt.Println(t.hello+"\n\n")
  for {

      fmt.Printf("%s", t.cursor)

      text, _ := reader.ReadString('\n')
      text = strings.Trim(strings.ToLower(reg.ReplaceAllString(text, " ")), " ")

      fmt.Printf("'%s' == '%s'\n", text, t.quitCommand)
      if text == t.quitCommand {
        fmt.Printf("\n\n%s", t.quitText)
        break
      }

      t.ProccessCommand(strings.Split(text, " "))
  }
}

func (t *Terminal) ProccessCommand(cmd []string) {
  t.commands = cmd[:]
  fmt.Println(t.GetCommands())
}

func (t *Terminal) GetCommands() []string {
  return t.commands
}
