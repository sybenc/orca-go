package authutils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Hasher interface {
	Hash(string) (string, error)
	Verify(string, string) bool
}

var Argon2id = &argon2id{}

type argon2id struct{}

const (
	hashTime    uint32 = 1
	hashMemory  uint32 = 64 * 1024
	hashThreads uint8  = 4
	hashKeyLen  uint32 = 64
	hashSaltLen        = 16
)

// Hash 将字符串进行Hash然后返回
func (a *argon2id) Hash(str string) (string, error) {
	salt := make([]byte, hashSaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(str), salt, hashTime, hashMemory, hashThreads, hashKeyLen)
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)
	hashedPassword := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		hashMemory, hashTime, hashThreads, saltB64, hashB64)
	return hashedPassword, nil
}

// Verify 验证字符串与存储的Hash值是否匹配
func (a *argon2id) Verify(str, storedHash string) bool {
	parts := strings.Split(storedHash, "$")
	if len(parts) != 6 {
		return false
	}

	// parts[0] 是空字符串，parts[1] 是算法名，parts[2] 是版本信息
	// parts[3] 是参数部分，parts[4] 是盐，parts[5] 是哈希
	if parts[1] != "argon2id" {
		return false
	}

	// 解析参数部分 m=...,t=...,p=...
	var memory, time uint32
	var threads uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false
	}

	// 解码盐和哈希
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// 使用提取的参数重新生成哈希
	newHash := argon2.IDKey([]byte(str), salt, time, memory, threads, uint32(len(hash)))

	// 使用常量时间比较哈希
	if subtle.ConstantTimeCompare(newHash, hash) == 1 {
		return true
	}
	return false
}
