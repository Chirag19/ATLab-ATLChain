package main

import (
    "encoding/pem"
    "encoding/base64"
    "crypto/x509"
    "crypto/rsa"
    "crypto/rand"
    "crypto/sha256"
    "io/ioutil"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

var privatekeyA = []byte(
`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCvWf+rchhWXxSc+kd5nioYDtdZAdVVSTlOzBoEDKyHiCrYRNVe
H1J4asas08UL1hkF0tLftrCdn4BYBtSlCkCkCOC8WTPSIRlxLfa94aEw2odV6JCF
9hoWUAwm4clBOO161RX7p4D5n8gMMrIKjo975C1Zn/YQac5HxOda8xDfQQIDAQAB
AoGAc2V97Nz8CTMvRJMsoGum9ggmTgv30dWLYkDNSibxD4xb7dF2vSdNxbM3JhuD
TGPMOdnhLppypniGJOfx3t7dZD5yd5Iz5cSM1R7Q+hOqiExdG6NjkPwEhSA4P0iK
AhHgYU3JgWpBSrgcikPVSiW8nTDnUluZW/KV/5aJ1gnYtqUCQQDi5DaKvustAQbk
NL3T0OfoVXG6P9xd7s1TGteCQPE2Xqd7kcM/iM8wAU6IUVaPos9e8HZXzT8JPjt/
PbumPp13AkEAxdkQFINBO4iejEXBUmWUVHCblCJx7pt/kXUDjtgUG5eJwI4bpKte
vUz8Z0oVFGOsQZd+k9ve0mQ91/+UUio3BwJAQGlDJp5Oi0coWq6yWSiMPYPMNnCc
sbnyZi5PkfW3xJSYfVcDE81V7C3iyoY0ybARqMUhA4oL5CbboyK2W9qYvwJBAIKM
VBv079pEr7mHXaTs+g8trrr0b5EuceKc/5gF5F7Ag1jXbE4f9gebAQF21Kn7ivJM
8GzILCNPma8pKcl9qYkCQQCOCt+/fjpymE9SjAWwR965XB0IoCT5z3KWeqA84VvY
dhyufEzlckUqNTHSQ7jLsL6BMYb9VLUa2F5q5xLdiV23
-----END RSA PRIVATE KEY-----`)

var publickeyA = []byte(
`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCvWf+rchhWXxSc+kd5nioYDtdZ
AdVVSTlOzBoEDKyHiCrYRNVeH1J4asas08UL1hkF0tLftrCdn4BYBtSlCkCkCOC8
WTPSIRlxLfa94aEw2odV6JCF9hoWUAwm4clBOO161RX7p4D5n8gMMrIKjo975C1Z
n/YQac5HxOda8xDfQQIDAQAB
-----END PUBLIC KEY-----`)


var privatekeyB = []byte(
`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC+UxqLcJPaCbIhvKzdWJUGS41E/6U3tPZJCTostUw0TeuR3QTW
ZGsYGMkAzYJhklkHnlIyVBgL9B0376xdHjOCNjxozkLdqC/gtLoTbjL4l9PZ2P6x
wG4AaXJ47JezXt7mh22ATg/rb31tLjiHd1Ne9AXIbhQEtfDg0f+ieG/AjQIDAQAB
AoGAW1MXHqejWnFil0uoiwGRaJbiL6SXy7Y6o2sZDhDkgwiMq84pHxLKTKK/+HGk
SVtm+v/eIyY0769wQcHwrDHsto7FYiP1JyUWFw1CWm2Af4+swrhQnMHNINgmgxSz
QysEeyXSHotzOVX1QCPZQiP4btqEJIkx5B7ddcwcs3oHJoUCQQDx3OgOnqm9f2XV
Xch7Ab2pamIfKYUKBbRUCwKyNZTZE897CcTXjagb18HJfg608SArJyrokU7Mk6FF
GyAUEvf/AkEAyXMCO0NvBjgghCg8VJqzSjQkm+N6LQa1qGfZx1pH5CvMy3cpPlcL
vZ4lj1NJ5lH5NOeT6auIlhSSLHInA92ncwJAUFaJenm3diuAHuyE8F72qfSdXS6E
c3zLlnMF1T45EBYlgAARs2vpYD49r3lA11eU0OC0vwWtQAT1t6e38xMN7wJAcH7v
QhUYTQrO7b5iYoS5lrijsQJJWhejHlZQQYljGEJ1bTIwMAYAInXMV8uVOy+P0UF5
UkZeUiFOt89PhlMjjQJBAOInLmhoZKPOz+MR28cagpNq7EclrR30TWWRPjWL0pB+
qfeQ3oVkQ5mLUEyirpA1KlKUmgWzhT9eWfgWLyJcJyg=
-----END RSA PRIVATE KEY-----`)

var publickeyB = []byte(
`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC+UxqLcJPaCbIhvKzdWJUGS41E
/6U3tPZJCTostUw0TeuR3QTWZGsYGMkAzYJhklkHnlIyVBgL9B0376xdHjOCNjxo
zkLdqC/gtLoTbjL4l9PZ2P6xwG4AaXJ47JezXt7mh22ATg/rb31tLjiHd1Ne9AXI
bhQEtfDg0f+ieG/AjQIDAQAB
-----END PUBLIC KEY-----`)

type TxCC struct {
}

func readFileAsByte(filePath string) []byte {
    b, _ := ioutil.ReadFile(filePath)
    return b
}

func makeAddress(publickey []byte) string {
    h := sha256.New()
    h.Write(publickey)
    addr := base64.StdEncoding.EncodeToString(h.Sum(nil))
    return addr
}

func Encryption(plaintext []byte, publickey []byte) []byte {
    pemkey, _ := pem.Decode(publickey)
    publickeyInterf, _ := x509.ParsePKIXPublicKey(pemkey.Bytes)
    newpublickey := publickeyInterf.(*rsa.PublicKey)
    cipherdata, _ := rsa.EncryptPKCS1v15(rand.Reader, newpublickey, plaintext)
    return cipherdata
}

func Decryption(cipherdata []byte, privatekey []byte) string {
    // 获取pem格式的私钥
    pemkey,_ := pem.Decode(privatekey)
    // 解析PKCS1格式的私钥
    newprivatekey,_ := x509.ParsePKCS1PrivateKey(pemkey.Bytes)
    // 解密
    plaintext, _ := rsa.DecryptPKCS1v15(rand.Reader,newprivatekey,cipherdata)
    return string(plaintext)
}

// 初始化
// args: 0-{pubkey_file_address},1-{plaintext}
func (t *TxCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
    // _, args := stub.GetFunctionAndParameters()
    // fmt.Println(args[0])
    // fmt.Println(args[1])
    // // TODO:初始化时给账户A赋值
    // publickey := readFileAsByte(args[0])
    // plaintext := []byte(args[1])
    // addressA := makeAddress(publickey)
    // cipherdata := Encryption(plaintext, publickey)
    // error := stub.PutState(addressA, cipherdata)
    // if error != nil {
    //     shim.Error("PutState failed!")
    // }
    fmt.Println("init Success")
    return shim.Success(nil)
}

func (t *TxCC) put(stub shim.ChaincodeStubInterface) pb.Response {
    addressA := makeAddress(publickeyA)
    fmt.Println("======>addressA:" + addressA)
    cipherdata := Encryption([]byte("A的数据"), publickeyA)
    error := stub.PutState(addressA, cipherdata)
    if error != nil {
        shim.Error("PutState failed!")
    }
    return shim.Success(nil)
}

// 交易
// args: 0-{pubkey_file_addressA},1-{pubkey_file_addressB},2-{prikey_file_addressA}
func (t *TxCC) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    // TODO:解密账户A的数据，同时将数据用B的公钥加密，存入B的地址
    // publickeyA := readFileAsByte(args[0])
    // publickeyB := readFileAsByte(args[1])
    // privatekeyA := readFileAsByte(args[2])
    addressA := makeAddress(publickeyA)
    addressB := makeAddress(publickeyB)
    cipherdataA, error := stub.GetState(addressA)
    if error != nil {
        shim.Error("GetState failed!")
    }
    plaintext := Decryption(cipherdataA, privatekeyA)
    cipherdataB := Encryption([]byte(plaintext + "-已转给B"), publickeyB)

    error = stub.PutState(addressB, cipherdataB)
    if error != nil {
        shim.Error("PutState failed!")
    }

    return shim.Success(nil)
}

// 查询
// args: 0-{pubkey_file_address},1-{prikey_file_address}
func (t *TxCC) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    // TODO:根据公钥找到地址，查询记录，同时使用私钥解密
    // publickey := readFileAsByte(args[0])
    // privatekey := readFileAsByte(args[1])
    var address string
    var privatekey []byte
    if args[0] == "A" {
        address = makeAddress(publickeyA)
        privatekey = privatekeyA
    } else if args[0] == "B"{
        address = makeAddress(publickeyB)
        privatekey = privatekeyB
    }
    cipherdata, error := stub.GetState(address)
    if error != nil {
        return shim.Error("query failed!")
    }

    plaintext := Decryption(cipherdata, privatekey)
    fmt.Printf("plaintext: %s\n", plaintext)
    fmt.Println("=====>Get plaintext:")
    fmt.Println(plaintext)
    return shim.Success([]byte(plaintext))
}

// Invoke
func (t *TxCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    if function == "invoke" {
        return t.invoke(stub, args)
    } else if function == "query" {
        return t.query(stub, args)
    } else if function == "put" {
        return t.put(stub)
    }

    return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")
}

func main(){
    err := shim.Start(new(TxCC))
    if err != nil {
        fmt.Printf("Error starting TxCC chaincode: %s", err)
    }
}
