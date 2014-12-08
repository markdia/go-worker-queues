package main
 
import (
  "fmt"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
  "log"
)
 
// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
  // Create, and return the worker.
  worker := Worker {
    ID: id,
    Work: make(chan WorkRequest),
    WorkerQueue: workerQueue,
    QuitChan: make(chan bool),
  }
  return worker
}
 
type Worker struct {
  ID int
  Work chan WorkRequest
  WorkerQueue chan chan WorkRequest
  QuitChan chan bool
}
 
// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w Worker) Start() {
  go func() {
    for {
      // Add ourselves into the worker queue.
      w.WorkerQueue <- w.Work
      select {
        case work := <-w.Work:
          // Receive a work request.
          fmt.Printf("worker%d: Received work request, some meaningful chunk of text\n", w.ID)
          //connect to AWS and do something to test
          auth := aws.Auth{
	        AccessKey: "ASDFASDFASDFASDK",
	        SecretKey: "DSFSDFDWESDADSFASDFADFDSFASDF",
		  }
		  euwest := aws.EUWest

		  connection := s3.New(auth, euwest)
		  mybucket := connection.Bucket("mytotallysecretbucket")
		  res, err := mybucket.List("", "", "", 1000)
		  if err != nil {
		    log.Fatal(err)
		  }
		  for _, v := range res.Contents {
		    fmt.Println(v.Key)
		  }


          fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.UID)
        case <-w.QuitChan:
          // We have been asked to stop.
          fmt.Printf("worker%d stopping\n", w.ID)
        return
      }
    }
  }()
}
 
// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
  go func() {
    w.QuitChan <- true
  }()
}