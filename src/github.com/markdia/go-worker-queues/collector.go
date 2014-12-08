package main
 
import (
  "fmt"
  "net/http"
  //"encoding/json"
  //"log"
)
 
// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)
 
func Collector(w http.ResponseWriter, r *http.Request) {
  // Make sure we can only be called with an HTTP POST request.
  if r.Method != "POST" {
    w.Header().Set("Allow", "POST")
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }
  //check for SNS headers etc
  //Parse the request body

  //faking it now
  var uid int = 101
  var fileUrl string = "/foo/bar/baz"
  
  // Now, take the items from the message and make a WorkRequest out of them.
  work := WorkRequest{UID: uid, FileUrl: fileUrl}
  
  // Push the work onto the queue.
  WorkQueue <- work
  fmt.Println("Work request queued")
  // And let the user know their work request was created.
  w.WriteHeader(http.StatusCreated)
  return
}