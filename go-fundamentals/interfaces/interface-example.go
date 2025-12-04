package main

var config = `{
  "serverPools": [
    {
      "name": "DemoPolicy_531-AD-Server-46d2f5eb",
      "endpoints": [
        {
          "address": "10.22.32.43:46"
        },
        {
          "address": "11.22.32.43:45"
        }
      ]
    }
  ],
  "externalServers": [
    {
      "name": "AD-Server-46d2f5eb",
      "serverType": "AAAConnector123",
      "poolName": "DemoPolicy_531-AD-Server-46d2f5eb",
      "dnsResolverName": "global_f5_internal_net_resolver"
    }
  ]
}`

func main() {
	//jsonData := []byte(config)
	//var m map[string]interface{}
	//err := json.Unmarshal(jsonData, &m)
	//if err != nil {
	//	panic(err)
	//}
	//
	//config1 := m["serverPools"]
	//serverPools, serverPoolsExist := config1.([]interface{})
	//println(serverPoolsExist)
	//println(serverPools)

	tmp := "[]"
	serverPools, serverPoolsExist := tmp.([]interface{})

}
