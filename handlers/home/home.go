package home

import (
	"../../gong"
)

func MainHandler(scope gong.Scope) {
	scope.WriteTemplate("handlers/home/home.html", nil)
}
