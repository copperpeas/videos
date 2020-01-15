package main
import "net/http"
import "github.com/julienschmidt/httprouter"
//import "./handler"

func  RegisterHandlers() *httprouter.Router{
    router := httprouter.New()
    router.POST("/user",CreateUser)
    router.POST("/user/:user_name",Login)
    return router
}
func main(){
    r:=RegisterHandlers()
    http.ListenAndServe(":8080",r)
}
