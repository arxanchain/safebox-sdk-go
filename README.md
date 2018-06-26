# Status
[![Build Status](https://travis-ci.org/arxanchain/safebox-sdk-go.svg?branch=master)](https://travis-ci.org/arxanchain/safebox-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arxanchain/safebox-sdk-go)](https://goreportcard.com/report/github.com/arxanchain/safebox-sdk-go)
[![GoDoc](https://godoc.org/github.com/arxanchain/safebox-sdk-go?status.svg)](https://godoc.org/github.com/arxanchain/safebox-sdk-go)

# safebox-sdk-go
This SDK enables Go developers to develop applications that interact with Arxanchain SafeBox

# Usage

## Install

Run the following command to download the Go SDK:

```code
go get github.com/arxanchain/safebox-sdk-go/api
```

## New safebox client

To invoke the SDK API, you first need to create a safebox client as follows:

```code
// Create safebox client
config := &restapi.Config{
	Address:    "http://API-Gateway-IP:PORT",
	ApiKey:     "Your-API-Access-Key",
	CryptoCfg: &restapi.CryptoConfig{
		Enable:         true,
		CertsStorePath: "/path/to/client/certs",
	},
}
safeboxClient, err = safeboxapi.NewSafeboxClient(config)
if err != nil {
	fmt.Printf("New safebox client fail: %v\n", err)
	return
}
fmt.Printf("New safebox client succ\n")
```

* When building the client configuration, the **Address** and **ApiKey** fields must
be set. The **Address** is set to the API-Gateway IP:PORT, and the
**ApiKey** is set to the API access key obtained on `ChainConsole` management page.

* If you invoke the APIs via `wasabi` service, the **Address** field should
be set to the http address of `wasabi` service, and the **CryptoCfg** field must be
set with **CryptoCfg.Enable** being `true` and **Cryptocfg.CertsStorePath** being the
path to client certificates (contains the platform public cert and user private key).

* `wasabi` service is ArxanChain BaaS API gateway with token authentication, data
encryption, and verifying signature.  For security requirement, enable crypto is
recommended for production environment.

## Trustee Key Pair

After creating safebox client, you can use this client to trustee key pair
as follows:

```code
// Build request header
header := http.Header{}

body := &safebox.SaveKeyPairRequetBody{
  UserDid:    string(userDid),
  PrivateKey: "privatekey",
  PublicKey:  "publickey",
}
resp, err := safeboxClient.TrusteeKeyPair(header, body)
if err != nil {
  fmt.Printf("trustee key pair faild.")
  return
}
fmt.Printf("trustee key pair success, %v", resp)
fmt.Printf("security code: %s", resp.Code)
```

## Query private key

After trusteeing key pair, you can query the private key as follows:

```code
// Build request header
header := http.Header{}

body := &safebox.OperateKeyInfo{
  UserDid: string(userDid),
  Code:    code,
}
resp, err := safeboxClient.QueryPrivateKey(header, body)
if err != nil {
  fmt.Printf("query private key faild.")
  return
}
fmt.Printf("query private key success, key: %v", resp.PrivateKey)
```

## Query public key

After trusteeing key pair, you can query the public key as follows:

```code
// Build request header
header := http.Header{}

body := &safebox.OperateKeyInfo{
  UserDid: string(userDid),
  Code:    code,
}
resp, err := safeboxClient.QueryPublicKey(header, body)
if err != nil {
  fmt.Printf("query public key faild.")
  return
}
fmt.Printf("query public key success, key: %v", resp.PublicKey)
```

## Recover Security Code

If you forget the security code, you need verify the user information, if success,
you can recover the security code, as follows:

```code
// Build request header
header := http.Header{}

resp, err := safeboxClient.RecoverAssistCode(header, userDid)
if err != nil {
  fmt.Printf("query code faild.")
  return
}
fmt.Printf("query code success, key: %v", resp.Code)
```

## Update Security Code

If the returns security code that trustee key pair is inconvenient to remember,
after verify user information success, you can update the security code.

```code
// Build request header
header := http.Header{}

body := &safebox.UpdateSecurityCodeRequestBody{
  UserDid:      string(userDid),
  OriginalCode: code,
  NewCode:      "我爱你中国",
}
err := safeboxClient.UpdateAssistCode(header, body)
if err != nil {
  fmt.Printf("update code faild, %v", err)
  return
}
fmt.Printf("update code success.")
```
