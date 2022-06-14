package main

import "syscall"

var (
	api        = syscall.NewLazyDLL("coupler.dll")
	initialize = getApiProc("Api_Initialize")
)

func getApiProc(name string) *syscall.LazyProc {
	return api.NewProc(name)
}

func Initialize(session int64, jsonInfo string) int32 {
	call, _, _ := initialize.Call(int2ptr(session), str2ptr(jsonInfo))
	return int32(call)
}
