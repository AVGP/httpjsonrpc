package httpjsonrpc

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

const (
  Version1 = "1.0"
  Version2 = "2.0"
)

func createRequestBody(method string, id interface{}, params []interface{}, version string) ([]byte, error) {
  jsonMap := map[string]interface{}{
      "method": method,
      "id":     id,
      "params": params,
  }

  if version == Version2 {
    jsonMap["version"] = version
  }

  return json.Marshal(jsonMap)
}

func makeRequest(address string, data []byte)(map[string]interface{}, error) {
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

  result := make(map[string]interface{})
  err = json.Unmarshal(body, &result)
  if err != nil {
      log.Fatalf("Unmarshal: %v", err)
    return nil, err
  }

  return result, nil
}

func CallV2(address string, method string, id interface{}, params []interface{})(map[string]interface{}, error){
    data, err := createRequestBody(method, id, params, Version2)
    if err != nil {
        log.Fatalf("Marshal: %v", err)
      return nil, err
    }

    return makeRequest(address, data)
}

func Call(address string, method string, id interface{}, params []interface{})(map[string]interface{}, error){
    data, err := createRequestBody(method, id, params, Version1)
    if err != nil {
        log.Fatalf("Marshal: %v", err)
    	return nil, err
    }

    return makeRequest(address, data)
}
