### mexc-perpetual-api 抹茶交易所USDT永续合约逆向API

##### 开空单示例
```go
package main

import (
	"log/slog"
	mexcperpetualapi "gitlab.lnamphp.com/tsingroo/mexc-perpetual-api"
)

func main() {
	COOKID_UID := "这个是用户登录后的Cookies的u_id字段"
	PROXY := "http://localhost:7890"
	mpc, err := mexcperpetualapi.NewMexcPerpetualClient(COOKID_UID, PROXY)
	if err != nil {
		panic(err)
	}
	// 最终到交易所成交的时候vol会有一个倍数放大或缩小，放大或缩小倍数需要从官方Api查询合约信息获取
	orderId, err := mpc.OpenKong("FUEL_USDT", 2, 10)
	if err != nil {
		panic(err)
	}
	slog.Info("OpenKong return", "orderId", orderId)
}
```


##### 开空单的请求
```JSON
{
    "symbol": "SONIC_USDT",
    "side": 3,
    "openType": 2,
    "type": "5",
    "vol": 9,
    "leverage": 20,
    "marketCeiling": false,
    "priceProtect": "0",
    "p0": "kHM7ubeyY4t6pQW9xnUW19V9syfnYE6cCPPlapgyon9vOJiUU8NhB+WqeMKwxSMf7G/NjzgnFL9hYOjQvRmJ/WVsYaPD9wovKAkLIBGTcrZlxzovA+0Sfe/ewxkFR/aALRmFFToiOgDMaEwimAcTj5XoqPjg1d1ws108AWt5srujGiR3LA8ShPJpqdiEtL+dQ3dPqxB15UCf878okmZKfja2dVJe1c+ePRRReqJkrAPod9rvEf2uHJa/dSB8cQhjZOOQn/Ujl+/Ja5RxIbVyCPJbpakTaIQC/qyokvdEg+cUx3DAbeSSAxliXcCotEfV41L4SuNQr3RLF1AWM475KPEQusBy/g==",
    "k0": "ZZ8T345BLE6g6DvwxZoeWR5bNL1pXsTLbrGFo63pV3Cz9apn7hb2BZRXgn7IxooRwtMQEy7b5Q9dKD45uIMDBaWfUclqCT7haMUklkWe0xGcG6Ie/pWjyBx4/9DtjnUKLpb13/H3JqO6SpzcoaCSnQMb0Zl3FF/1TdHNU/lh75V0LeIjbenzJevbC7dMbe8HS41TM+3NwC6SpD3YrhvD+/jXgdiNIyMaTqdtT6z0lZfW87PVoXXfD48YBi/YfNTbbuhaltd6qp6FUYAis9b8LaSoGwIFnKhn8/qrvJfTG9AlseXtUrhS/MceWkRn9h0A22WrVtEQH1rM2a4yphcg/A==",
    "chash": "d6c64d28e362f314071b3f9d78ff7494d9cd7177ae0465e772d1840e9f7905d8",
    "mtoken": "624580bc74f8de8c230e34de7ca9d415",
    "ts": 1736996992679,
    "mhash": "b2d52c26ce0879d5d704893fa6e026a6"
}
```

##### 开空单的响应
```JSON
{
    "success": true,
    "code": 0,
    "data": {
        "orderId": "634326739784909376",
        "ts": 1736905679988
    }
}
```

##### 开多单的请求
```JSON
{
    "symbol": "FUEL_USDT",
    "side": 1,
    "openType": 2,
    "type": "5",
    "vol": 1,
    "leverage": 10,
    "marketCeiling": false,
    "priceProtect": "0",
    "p0": "dxuq5b/R3g5rjV1H28uSv2lgrtPv5JqvyB03rItqTekCrcTmeURjzLfkd+Aqph3u0i/aABXdh29l/rRGWOFBmKKn7/2Vur9hiArb5hrj3mJBzep9jUI+T5Rhi1EtKzR8mPQj/jYHVbxYIY9uC+VDz1xfCpAPwQkjj13dXXp0Wwi/AIX4sHphGmRpOnfU3zCgKXRxNXEhqFojq6sz20cCZEY45MOl+FxiyAtteAjYWQbZMJ7SC5wOliHCvnRcC+SIFw7+nqBb9LGOXdw/k7ACHHBp/axyHhOkDTVdBRaYUXP9FBxpZUTDln7DAzJbZX/9PTm6EEyk85Fq8MSUZD4UDX5tUuxCDw==",
    "k0": "dE4Mu3KVR6jnbqGdoq9ybb9bLNpbP3t0lb1+HMq9m+od3LP5Gnhgk6emZu4o9ZA35WrGD5KgQL2g5gAIKZSdQIt6sGxYJqsXcVdj01S8erT23OVEj4Pey1JiFDvnc9bmfpJYo6qOTTEZjo4gVSCAJLvoNPLwM37Vt2LT/CExT2bhKaTaprH+3TilcmEFOFDlRAsgMlzb9APbdV1retdKVwhx6VnAyoM7omyhDSzN3FoLUfmWW7CF0jIjS90WpvTwdXUHt+ijIpOPhaDybIm6YaWcwGMq2NfVwSgVrGshaLLXwrPW+Q7rEF0Yex9598KFFEPzsto4QhGVhIEzMpDTxw==",
    "chash": "d6c64d28e362f314071b3f9d78ff7494d9cd7177ae0465e772d1840e9f7905d8",
    "mtoken": "624580bc74f8de8c230e34de7ca9d415",
    "ts": 1737025173808,
    "mhash": "b2d52c26ce0879d5d704893fa6e026a6"
}
```