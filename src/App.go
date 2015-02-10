package main

import (
    "strings"
    "net/http"
    "github.com/go-martini/martini"
    "gopkg.in/mgo.v2"
)

func main() {
    m := martini.Classic()
    m.Get("/", func() string {
        return "Hello world!"
    })
    m.Post("/store/:somedata",  /*binding.Json( data{} ),*/ addSomeDataController)
    m.Use(mongoHandler())
    m.Run()
}

func mongoHandler() martini.Handler {
    session, err := mgo.Dial( "localhost/goApp" )
    if err != nil {
        panic( err )
    }

    return func (c martini.Context ) {
        reqSession := session.Clone()
        c.Map( reqSession.DB( "goApp" ) )
        defer reqSession.Close()

        c.Next()
    }
}

func addSomeDataController( params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string) {
    resource :=  strings.ToLower( params["somedata"] )
    writer.Header().Set("Content-Type", "application/json")

    return http.StatusOK, "POST placeholder " + resource
}