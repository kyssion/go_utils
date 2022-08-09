package crypto

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
)

type dsaSignature struct {
	R, S *big.Int
}

func readFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(f)
}

// ParseDSAPrivateKey 通过字符串解析 DSA 私钥匙
func ParseDSAPrivateKey(der []byte) (*dsa.PrivateKey, error) {
	var k struct {
		Version int
		P       *big.Int
		Q       *big.Int
		G       *big.Int
		Pub     *big.Int
		Priv    *big.Int
	}
	// 使用asn.1 方法获取dsa结构体信息
	rest, err := asn1.Unmarshal(der, &k)
	if err != nil {
		return nil, errors.New("failed to parse DSA key: " + err.Error())
	}
	if len(rest) > 0 {
		return nil, errors.New(fmt.Sprintf("garbage after DSA key ,  info : %s", rest))
	}

	return &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: dsa.Parameters{
				P: k.P,
				Q: k.Q,
				G: k.G,
			},
			Y: k.Pub,
		},
		X: k.Priv,
	}, nil
}

// ParseDSAPublicKey 通过字符串解析 DSA 公钥信息
func ParseDSAPublicKey(der []byte) (*dsa.PublicKey, error) {
	// 使用x509 公钥标准格式读取公钥信息
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *dsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("invalid type of public key")
	}
}

func ParseDSAPrivateKeyFromFile(path string) (*dsa.PrivateKey, error) {
	chunk, err := readFile(path)
	if err != nil {
		return nil, err
	}
	// 读取pem格式的编码文件
	block, rest := pem.Decode(chunk)
	if len(rest) != 0 {
		return nil, errors.New(fmt.Sprintf("failed to parse PEM block , info : %s", rest))
	}
	// 构建dsa私钥对象
	return ParseDSAPrivateKey(block.Bytes)
}

func ParseDSAPublicKeyFromFile(path string) (*dsa.PublicKey, error) {
	chunk, err := readFile(path)
	if err != nil {
		return nil, err
	}
	// 读取pem格式的编码文件
	block, rest := pem.Decode(chunk)
	if len(rest) != 0 {
		return nil, errors.New(fmt.Sprintf("failed to parse PEM block , info : %s", rest))
	}
	// 构建dsa私钥对象
	return ParseDSAPublicKey(block.Bytes)
}

func ParseSignatureFromFile(path string) (*big.Int, *big.Int, error) {
	chunk, err := readFile(path)
	if err != nil {
		return nil, nil, err
	}
	var s dsaSignature

	rest, err := asn1.Unmarshal(chunk, &s)
	if err != nil {
		return nil, nil, errors.New("failed to parse signature: " + err.Error())
	}
	if len(rest) > 0 {
		return nil, nil, errors.New("garbage after signature")
	}
	return s.R, s.S, nil
}

func hash(file string) ([]byte, error) {
	chunk, err := readFile(file)
	if err != nil {
		return nil, err
	}

	sum := sha1.Sum(chunk)
	return sum[:], nil
}

func sign(hash []byte, keyFile string) ([]byte, error) {
	priv, err := ParseDSAPrivateKeyFromFile(keyFile)
	if err != nil {
		return nil, err
	}

	var s dsaSignature
	s.R, s.S, err = dsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		return nil, err
	}

	return asn1.Marshal(s)
}

func verify(hash []byte, keyFile string, signatureFile string) ([]byte, error) {
	pub, err := ParseDSAPublicKeyFromFile(keyFile)
	if err != nil {
		return nil, err
	}

	r, s, err := ParseSignatureFromFile(signatureFile)
	if err != nil {
		return nil, err
	}

	if dsa.Verify(pub, hash, r, s) {
		return []byte("Verified OK\n"), nil
	} else {
		return nil, errors.New("Verification Failure")
	}
}
