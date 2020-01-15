//package main
package handler

import "io"
import "net/http"
import "github.com/julienschmidt/httprouter"

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
    io.WriteString(w,"create user handler")
}
func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
    uname := p.ByName("user_name")
    io.WriteString(w,uname)
}
