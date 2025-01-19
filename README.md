### mexc-perpetual-api 抹茶交易所USDT永续合约逆向API

##### 开空单示例1
```go
package main

import (
	"log/slog"
	mexcperpetualapi "github.com/testzhaoxiaofei/mexc-perpetual-api-master"
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

##### 平空的请求
```JSON
{
    "symbol": "SPELL_USDT",
    "side": 2,
    "openType": 2,
    "leverage": 10,
    "type": 5,
    "vol": 10,
    "priceProtect": "0",
    "p0": "ea3+C+1pcTvVKOpvd0T6ijI5b4/1PAb5ZbK3jPszFOKUP5c8ji5gixRzIChUeGviGvSr9RVc7Tu4nj/zOkIkjw5/GdFVvh02+jG5UnN4Agfz2o0LMfAYsfgiQFVCQiaGGYGRcmRJY08rXzeNFx4qK60payUNCRJpZTIAuMzcKR/u7OWZWmvvAtD6GZ7vwW/fGWILTL07yhggY/65JGsHaLgeud2q3MywR8LrH7kl8kF8CTqcX1WNzTdkY6bwit7Nv84M5ooEbyiLX3muld5dbkpv4T+cxPe72TlP4DaYv3gJEIcacyi7hBKKNzpm0gh7CXJwayACE2ItLSLJo3Lo7uHPnTxsAg==",
    "k0": "fSR7iSqolISRtTk1C4RLFvDFzaZFXbQxY6QHLeJqzxVrd8+n/CSI0sG4yrBxuRXrPQdTPYRSbtKFw/DL2hb1Nj5AepXLx674+nsPQ9FfmKaV96kDk+6gO9foJjuSyFDldjAr/BT5mm54dIbdlw5W6nMc7n17YEHf7v2N3+bl5rFs1TTzMMPN5fqy95cqgHk3mVaaijC9vQ4Lk0SvCdMEe/oyummJ67GetcwLgmV9YUe6TQRydMv3zM/Q8LNlkuER1bfkyRkkPt0YAKz+KOt0edvvfFytTgdK7j2Bs4GNnhdX3z8Wc3Bdk28KOpHHqOVvUVIM6bPEHd0ZV8Mnyd9Gaw==",
    "chash": "d6c64d28e362f314071b3f9d78ff7494d9cd7177ae0465e772d1840e9f7905d8",
    "mtoken": "5e39b76bd8607064225ef48f9b5833d7",
    "ts": 1737263936005,
    "mhash": "8d0726532c03c450098a8eb8734329fc"
}
```

#### 平多的请求
```JSON
{
    "symbol": "SPELL_USDT",
    "side": 4,
    "openType": 2,
    "leverage": 10,
    "type": 5,
    "vol": 1,
    "priceProtect": "0",
    "p0": "YdiCF5ULsNRWpy456qmGvgyqBCPuPfvjZL9WpgO/9SKD15DCP5wQ5VRTugPe463D/KrEmJXBFlLwX0c+Ilb8iqyd1C9plmE3B17SxQJ+6CDwYCAH2nfSnlLJ7X04P57NDdu3W/1fFprnIiZIFAWDf0fPQBfeI98h7WEdsXKtyuJ4NCm9gXvjANKYg332Icylq8mIp4ZeaCBr/WEup5RVM1X+el8jmKFpFBx2X/Zo/0/H3vPdZScza8dN7U5FMpMxwBzLY7BUPOJA43MzaxvO5qJV7It09PihRYuyLDDu3c0K2XYmqaC5Vxpn/TC88gi0bjsZQ3LB6a2I7ep31HlOzVkyzQjp8A==",
    "k0": "E9j4F9fMF6YXfVpdgQMzcAj1VzqwN+GDsnOl4FgGiuvnjw18dVKHIV9zqCJAWaJKZ2dRg66EV83aCCRJtRoavXKRG7iyYEP3MBp03X9Iww5MsWyK08Kwe2YsNMQKH4lSUyZvixOB2NhhzJ+3ZvmDufjKhV0u/LxqqxsH7kBcz4MLhYkMg9pUYE2tyt8jpf28lhhmoUifY3kKFzlLZ1PXpzD79UyDoWIMCKO+PPSSMUJHBh+kx9b/kQtG3t6tPEtSttSG4PXDP18XoNMSfl5zIwKWZcfAyxWOPDgAv3ypILX2o+NdGpdYOE3u/WobsiEPi5s55EWrxi8eEGDFovDHkw==",
    "chash": "d6c64d28e362f314071b3f9d78ff7494d9cd7177ae0465e772d1840e9f7905d8",
    "mtoken": "5e39b76bd8607064225ef48f9b5833d7",
    "ts": 1737264491676,
    "mhash": "8d0726532c03c450098a8eb8734329fc"
}
```