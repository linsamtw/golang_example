package main

import (
	"fmt"
  "github.com/valyala/fasthttp"
  // "encoding/json"
  // "os"
)

var (
    str string
    contentType string
    strContentType = []byte("Content-Type")
    strApplicationJSON = []byte("application/json")
)

func main() {
  var CONTENT = "";
  str = string(CONTENT)
  _ = str
  
  h := requestHandlerText
	h = fasthttp.CompressHandler(h)
	fasthttp.ListenAndServe(":80", h)
}


func requestHandlerText(ctx *fasthttp.RequestCtx) {
  // ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
  fmt.Fprintf(ctx, str)
  fmt.Fprintf(ctx, "Hello, world!7777777\n\n")
  data := []byte(`{"id" : 1 , "name" : "Daniel"}`)
  // s := `{"text":"I'm a text.","number":1234,"floats":[1.1,2.2,3.3],
  // "innermap":{"foo":1,"bar":2}}`

  // var data map[string]interface{}
  // json.Unmarshal([]byte(s), &data)
  // jsonBody, _ := json.Marshal(data)
	ctx.SetContentType("application/json; charset=utf-8")   //  *************
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(data)
}
