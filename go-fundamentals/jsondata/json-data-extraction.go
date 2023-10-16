package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	data := `
{
  "id": "53a74f0c-bb18-48bc-a39f-43d1c5e5525c",
  "verb": "POST",
  "uri": "/files",
  "request": "{\"id\":\"cae48fec-abe0-469e-9ff9-8a0430f88b02\",\"name\":\"test_file3\",\"fileName\":\"test_file3\",\"description\":\"\",\"uri\":\"file://opt/f5/mbip/subsystem/csm/shared/persisted/files/cae48fec-abe0-469e-9ff9-8a0430f88b02-6ab85d05-ae16-4d25-8668-3956c93d36b0\",\"hash\":\"2682f1e979272ac69c2180c07ba640bb7578ff7324eada348d5259f4bafed277\",\"size\":393459712,\"type\":\"GENERIC\"}",
  "role": "AGENTCONFIG",
  "status": "PENDING",
  "message": {
    "code": "",
    "title": "",
    "detail": "",
    "links": null
  },
  "creationTime": "2023-10-11T05:02:48.078064Z",
  "updateTime": "2023-10-11T05:02:48.078064Z",
  "owner": "admin-cm",
  "authn": "local",
  "authz": "global"
}
`

	requestJson := gjson.Get(data, "request").String()
	id := gjson.Get(requestJson, "id").String()
	fmt.Printf("request_json = %s", requestJson)
	fmt.Printf("id = %s", id)

}
