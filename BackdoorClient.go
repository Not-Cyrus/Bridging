package main 

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	NeedsUsername := true
	NeedsPassword := true
	var Username string
	var Password string
	fmt.Print("Insert your username: ")
	for {
		text,_ := reader.ReadString('\n')
		text = strings.Replace(text,"\r\n","",-1)
		if NeedsUsername {
			Username = text
			NeedsUsername = false
			fmt.Print("Insert your password: ")
		} else if NeedsPassword {
			Password = text
			NeedsPassword = false
			cmd := exec.Command("cmd","/c","cls");cmd.Stdout = os.Stdout;cmd.Run()
			fmt.Print("Command & Command Args: ")
		}else if !NeedsUsername && !NeedsPassword {
			fmt.Print("Command & Command Args: ")
			res,err := http.Post("YourURL?Username=" + url.QueryEscape(Username) + "&Password=" + url.QueryEscape(Password) + "&Command=" + url.QueryEscape(text),"text/html",nil)
			if err != nil {}
			res.Body.Close()
		}
	}
}