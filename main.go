// The MIT License (MIT)
//
// Copyright (c) 2016 aerth
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// For use with tropo.com web API (point your endpoint to x.x.x.x:8080 etc)
// This server redirects calls to the number specified by the environmental variable "TRANSFER" or a command line argument.
// ./transfer 15555555555
// TRANSFER=15555555555 ./transfer
// As this is meant as a docker container, the port is fixed at 8080.


package main

import (
    "io"
    "os"
    "net/http"
    "fmt"
)


func transfer(w http.ResponseWriter, r *http.Request) {
number := os.Getenv("TRANSFER")
if number == "" { number = os.Args[1] }
    io.WriteString(w, `{
   "tropo":[
      {
         "say":[
            {
               "value":"Please hold while you are transferred."
            }
         ]
      },
      {
         "transfer":{
            "to":"`+number+`
         }
      }
   ]
}

`)
}

func main() {

if len(os.Args) < 2 && os.Getenv("TRANSFER") == "" { fmt.Println("Where to transfer? \n"+ os.Args[0] + " 15555555555"); os.Exit(1) }
number := os.Getenv("TRANSFER")
if number == "" { number = os.Args[1] }
if len(number) < 4 { fmt.Println("Seems like an invalid number. Exiting."); os.Exit(1) }
    fmt.Println("Transfering to "+number+" on 0.0.0.0:8080")
    http.HandleFunc("/", transfer)
    http.ListenAndServe(":8080", nil)
}
