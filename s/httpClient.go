package biuper

import (
    "bytes"
    "log"
    "os"
    "fmt"
    "io"
    "mime/multipart"
    "path/filepath"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

var (
    logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func FormPost(data map[string]string, dest string) string {
    v := url.Values{}

    for key, value := range data { 
        v.Add(key, value)
    }

    body := ioutil.NopCloser(strings.NewReader(v.Encode()))
    client := &http.Client{}
    req, _ := http.NewRequest("POST", dest, body)

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
    //logger.Println("%+v\n", req)

    resp, err := client.Do(req)
    defer resp.Body.Close()
    result, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(result), err)
    if err != nil {
        logger.Fatal(err)
    }
    return string(result)
}


func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
    file, err := os.Open(path)
    if err != nil {
      return nil, err
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile(paramName, filepath.Base(path))
    if err != nil {
      return nil, err
    }
    _, err = io.Copy(part, file)
    for key, val := range params {
      _ = writer.WriteField(key, val)
    }
    err = writer.Close()
    if err != nil {
      return nil, err
    }

    request, error := http.NewRequest("POST", uri, body)
    request.Header.Set("Content-Type", writer.FormDataContentType())
    // ↑坑死我了
    return request, error
}


func Upload(params map[string]string, filename string)  {

    request, err := newfileUploadRequest("http://upload.qiniu.com/", params, "file", filename)
    if err != nil {
      log.Fatal(err)
    }

    client := &http.Client{}
    resp, err := client.Do(request)
    if err != nil {
      log.Fatal(err)
    } else {
      body := &bytes.Buffer{}
      _, err := body.ReadFrom(resp.Body)
      if err != nil {
          log.Fatal(err)
      }
      resp.Body.Close()
      logger.Println(fmt.Sprintf("[INFO] Retcode: %d",resp.StatusCode))
      //logger.Println(fmt.Sprintf("[INFO] %s",resp.Header))
      logger.Println(fmt.Sprintf("[INFO] Response: %s", body))
    }
}

