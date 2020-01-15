package defs

//requests
type UserCredential struct{
    Username string `json:"user_name"`
    Pwd string `json:"pwd"`
}
/*
会形成json格式：
{
    user_name:xxx
    pwd:xxx
}
*/
