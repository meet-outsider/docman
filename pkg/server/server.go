package server

import (
	"docman/cfg"
	"errors"
	"fmt"
)

func Run() error {
	if err := Inst.Run(fmt.Sprintf(":%d", cfg.Config.Server.Port)); err != nil {
		return errors.New("server failed")
	}
	return nil
}
