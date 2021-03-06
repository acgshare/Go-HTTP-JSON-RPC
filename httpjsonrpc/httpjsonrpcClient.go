package httpjsonrpc

// Copyright 2011-2014 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

type resultContainer struct {
Error interface{}          `json:"error"`
Result json.RawMessage `json:"result"`
Id interface{} `json:"id"`
}

func Call(address string, method string, id interface{}, params []interface{})(*resultContainer, error){
    data, err := json.Marshal(map[string]interface{}{
        "method": method,
        "id":     id,
        "params": params,
    })
    if err != nil {
        log.Fatalf("Marshal: %v", err)
    	return nil, err
    }
    resp, err := http.Post(address,
        "application/json", strings.NewReader(string(data)))
    if err != nil {
        log.Fatalf("Post: %v", err)
    	return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("ReadAll: %v", err)
    	return nil, err
    }
    var result resultContainer
    err = json.Unmarshal(body, &result)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    	return nil, err
    }
    //log.Println(result)
    return &result, nil
}
