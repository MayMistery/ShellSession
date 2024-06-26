package plugin

import (
	"encoding/base64"
	"fmt"
	"github.com/KPGhat/ShellSession/session"
	"github.com/KPGhat/ShellSession/utils"
	"log"
	"os"
	"path"
)

func (plugin *commandPlugin) Upload(session *session.Session, args []string) string {
	// TODO Adapt to other platforms
	fmt.Println("[*]Warning only support for linux now...")
	if len(args) == 0 || len(args) > 2 {
		plugin.UploadHelp()
	}
	srcFileContent, err := os.ReadFile(args[0])
	if err != nil {
		log.Println(err.Error())
	}

	var result []byte
	// TODO Split the content into block
	srcFileContentBase64 := base64.StdEncoding.EncodeToString(srcFileContent)
	if len(args) == 1 {
		result = session.ExecCmd("echo " + srcFileContentBase64 + "|base64 -d >" + path.Base(args[0]))
	} else {
		result = session.ExecCmd("echo " + srcFileContentBase64 + "|base64 -d >" + args[1])
	}

	return string(result)
}

func (plugin *commandPlugin) UploadHelp() {
	utils.Congrats("Upload Usage: upload src [dst]")
	utils.Congrats("If dst is not set, will upload to the shell current dir")
}
