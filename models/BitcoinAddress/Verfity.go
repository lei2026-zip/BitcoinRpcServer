package BitcoinAddress

import (
	base58 "BitcoinRpcServer/BTCUtils/Base58"
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func New_elliptic_publickey()[]byte{
	curve := elliptic.P256()
	pri, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//ecdsa.GenerateKey(curve,rand.Reader)
	fmt.Println(pri)
	fmt.Println(len(pri))

	//拼接x和y，生成公钥
	pubKey :=elliptic.Marshal(curve,x, y)
	fmt.Println(hex.EncodeToString(pubKey))
	fmt.Println(len(pubKey))
	return pubKey
}

/**
   返回新比特币地址
 */
func NewBitcoinAddress(version byte, PublicKey []byte)(string){
	//sha256 计算
	sha256Hash := sha256.New()
	sha256Hash.Write(PublicKey)
	pub256Hash := sha256Hash.Sum(nil)

	//ripemd160 hash计算 （github下载ripemd160库)
	rpmd160 := ripemd160.New()
	rpmd160.Write(pub256Hash)
	pubRpmd160 := rpmd160.Sum(nil)

	fmt.Println(len(hex.EncodeToString(pubRpmd160)))

	//第二步、添加版本号
	versionPubRmpd160 := append([]byte{version}, pubRpmd160...)

	//第三步、计算校验位
	sha256Hash.Reset() //重置
	sha256Hash.Write(versionPubRmpd160)
	hash1 := sha256Hash.Sum(nil)

	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	hash2 := sha256Hash.Sum(nil)

	check := hash2[0:4]

	//第四步、拼接
	addBytes := append(versionPubRmpd160, check...)

	//第五步、对地址的拼接结果进行base58编码
	addr := base58.Encode(addBytes)
	fmt.Println("生成了新的比特币的地址：", addr)
	return addr
}

func CheckAddress(Address string)(bool){

	//第一步，base58解码
	addBytes := base58.Decode(Address)
	if len(addBytes) <=4 {
		return false
	}
	//第二步，截取后4位，得到校验位
	check := addBytes[len(addBytes)-4 : len(addBytes)]

	//第三步，双hash计算
	//1、获取到base58解码后的数据除去后4个字节的数据
	versionRipemd160 := addBytes[0 : len(addBytes)-4]
	//2、sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(versionRipemd160)
	hash1 := sha256Hash.Sum(nil)

	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	hash2 := sha256Hash.Sum(nil)
	//3、获取前4个字节
	deCheck := hash2[0:4]

	//第四步，将校验位与再次双hash以后得到的校验位进行比较
	return bytes.Compare(check, deCheck) == 0
}