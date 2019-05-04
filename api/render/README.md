# Render

`render` is a simple response render helper and prepared collection of the responses 

## Usage 
```
package main

import (
    "net/http"
    "github.com/lancer-kit/armory/api/render"
)

func Get(w http.ResponseWriter, r *http.Request) {
    testMsg := r.Host
    render.WriteJSON(w, http.StatusOK, testMsg)
}

func Crap(w http.ResponseWriter, r *http.Request) {
    render.ResultBadRequest.Render(w)
}

func Ok(w http.ResponseWriter, r *http.Request)  {
    render.ResultSuccess.Render(w)
}
```

