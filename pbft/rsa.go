package pbft

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/cloudflare/cfssl/log"
	"os"
	"strconv"
)

//如果当前目录下不存在目录Keys，则创建目录，并为各个节点生成rsa公私钥
func GenRsaKeys() {
	if !isExist("./Keys") {
		log.Info("检测到还未生成公私钥目录，正在生成公私钥 ...")
		err := os.Mkdir("Keys", 0777)
		if err != nil {
			log.Error()
		}
		for i := 0; i <= 4; i++ {
			if !isExist("./Keys/N" + strconv.Itoa(i)) {
				err := os.Mkdir("./Keys/N"+strconv.Itoa(i), 0777)
				if err != nil {
					log.Error()
				}
			}
			priv, pub := GetKeyPair()
			privFileName := "Keys/N" + strconv.Itoa(i) + "/N" + strconv.Itoa(i) + "_RSA_PIV"
			file, err := os.OpenFile(privFileName, os.O_RDWR|os.O_CREATE, 0777)
			if err != nil {
				log.Error(err)
			}
			defer file.Close()
			file.Write(priv)

			pubFileName := "Keys/N" + strconv.Itoa(i) + "/N" + strconv.Itoa(i) + "_RSA_PUB"
			file2, err := os.OpenFile(pubFileName, os.O_RDWR|os.O_CREATE, 0777)
			if err != nil {
				log.Error(err)
			}
			defer file2.Close()
			file2.Write(pub)
		}
		log.Info("已为节点们生成RSA公私钥")
	}
}

//生成rsa公私钥
func GetKeyPair() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}

//判断文件或文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Info(err)
		return false
	}
	return true
}

//数字签名
func (p *pbft) RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Info("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//签名验证
func (p *pbft) RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

//数字签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Info("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

//签名验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		//panic(err)
		log.Info("验签不通过！")
		return false
	}
	return true
}
