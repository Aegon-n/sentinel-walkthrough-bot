package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/echo"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/nodes/master-node/models"
	"github.com/Aegon-n/sentinel-bot/socks5-proxy/nodes/master-node/utils"
)

func AddNewUser(ctx echo.Context) error {
	var body models.AddUser
	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		ctx.JSONBlob(http.StatusInternalServerError, []byte("error while reading request body"))
		return err
	}
	wg := new(sync.WaitGroup)
	var cmd strings.Builder
	_, e := cmd.WriteString("printf ")
	if e != nil {
		return nil
	}
	_, e = cmd.WriteString("'" + body.Password + "\n" + body.Password + "\n" + "'" + " | adduser " + body.Username)
	if e != nil {
		return nil
	}
	//cmd := "printf 'MemsIr[OkAj4\nMemsIr[OkAj4\n' | adduser awesome"
	wg.Add(1)
	go utils.ExecCmd(cmd.String(), wg)
	wg.Wait()
	//if e != nil {
	//	log.Println("error while executing command")
	//	return nil
	//}
	ctx.JSONBlob(http.StatusCreated, []byte(`{"message": "added user"}`))
	return nil
}

func RemoveUser(ctx echo.Context) error {

	var body models.DelUser
	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		ctx.JSONBlob(http.StatusInternalServerError, []byte(`{"message": "error while deleting user"}`))
		return nil
	}
	defer ctx.Request().Body.Close()
	wg := new(sync.WaitGroup)
	cmd := "deluser " + body.Username
	wg.Add(1)
	go utils.ExecCmd(cmd, wg)
	wg.Wait()

	ctx.JSONBlob(http.StatusAccepted, []byte(`{"message": "deleted user"}`))
	return nil
}

func RootFunc(ctx echo.Context) error {
	ctx.JSONBlob(http.StatusOK, []byte(`{"status": "up"}`))
	return nil
}
