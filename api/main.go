/*
go run *.go
*/
package main

import "net/http"
import "github.com/julienschmidt/httprouter"

func  RegisterHandlers() *httprouter.Router{
    router := httprouter.New()
    router.POST("/user",CreateUser)
    //router.POST("/user",nil)
    //router.POST("/user/:user_name",nil)
    router.POST("/user/:user_name",Login)
    return router
}
func main(){
    r:=RegisterHandlers()
    //Listen->RegisterHandlers->handlers
    http.ListenAndServe(":8080",r)
}
