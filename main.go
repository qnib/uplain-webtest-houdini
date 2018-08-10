package main

import (
    // Standard library packages
    "fmt"
    "os"
    "log"
    "net"
    "net/http"
    "bytes"
    "os/exec"
    // Third party packages
    "github.com/julienschmidt/httprouter"
)


func getCntName() string {
  cntName := os.Getenv("CONTAINER_NAME")
  if cntName == "" {
    cntName = "unkown"
  }
  return cntname
}

// https://blog.golang.org/context/userip/userip.go
func getIP(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
    podName := os.Getenv("POD_NAME")
    if podName == "" {
      podName = "unkown"
    }
    fmt.Fprintf(w, "You've hit cnt:%s at path:%s on pod:%s\n", getCntName(), req.URL.Path, podName)
    ip, port, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        //return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
        return
    }
    fmt.Fprintf(w, "IP: %s\n", ip)
    fmt.Fprintf(w, "Port: %s\n", port)
}

func getName(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  fmt.Fprintf(w, "container: %s\n", getCntName())
}

func getGPUs(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  cmd := exec.Command("nvidia-smi", "-L")
  cmdOutput := &bytes.Buffer{}
  cmd.Stdout = cmdOutput
  err := cmd.Run()
  if err != nil {
    fmt.Fprintf(w, "%s: %s", getCntName(), err.Error())
    return
  }
  fmt.Fprintf(w, "%s: %s", getCntName(), string(cmdOutput.Bytes()))
}

func getTask(w http.ResponseWriter, req *http.Request, _ httprouter.Params){
  srvName := os.Getenv("SERVICE_NAME")
  if srvName == "" {
    srvName = "unkown"
  }
  taskSlot := os.Getenv("TASK_SLOT")
  if taskSlot == "" {
    taskSlot = "unkown"
  }
  fmt.Fprintf(w, "%s.%s\n", srvName, taskSlot)
}
func main() {
    // Instantiate a new router
    r := httprouter.New()

    r.GET("/", getIP)
    r.GET("/cntname", getName)
    r.GET("/task", getTask)
    r.GET("/gpus", getGPUs)

    port := os.Getenv("HTTP_PORT")
    if port == "" {
      port = "8080"
    }
    addr := fmt.Sprintf("%s:%s", os.Getenv("HTTP_HOST"), port)
    l, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Start Webserver on %s (v0.1.0)", addr)
    log.Fatal(http.Serve(l, r))
}
