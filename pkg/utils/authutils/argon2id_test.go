// auth_test.go
package authutils

import (
	"fmt"
	"strings"
	"testing"
)

// TestHash 测试 Hash 方法是否正确生成哈希字符串
func TestHash(t *testing.T) {
	password := "TestPassword123!"

	hash, err := Argon2id.Hash(password)
	if err != nil {
		t.Fatalf("Hash() 返回了一个错误: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash() 返回了一个空字符串")
	}

	// 检查哈希字符串是否以正确的前缀开头
	expectedPrefix := "$argon2id$"
	if !strings.HasPrefix(hash, expectedPrefix) {
		t.Errorf("Hash() 返回的哈希前缀不正确。获得: %s, 预期前缀: %s", hash, expectedPrefix)
	}

	// 使用 "$" 分割哈希字符串，确保其包含正确的部分数量
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		t.Errorf("Hash() 返回的哈希部分数量不正确。获得: %d, 预期: 6", len(parts))
	}

	// 可选：进一步检查参数部分是否正确
	// 例如，检查 parts[3] 是否包含 m=65536,t=1,p=4
	expectedParams := fmt.Sprintf("m=%d,t=%d,p=%d", hashMemory, hashTime, hashThreads)
	if parts[3] != expectedParams {
		t.Errorf("Hash() 返回的参数部分不正确。获得: %s, 预期: %s", parts[3], expectedParams)
	}
}

// TestVerify 测试 Verify 方法在不同情况下的表现
func TestVerify(t *testing.T) {
	password := "TestPassword123!"
	wrongPassword := "WrongPassword!"
	hash, err := Argon2id.Hash(password)
	if err != nil {
		t.Fatalf("Hash() 返回了一个错误: %v", err)
	}

	tests := []struct {
		name       string
		password   string
		storedHash string
		want       bool
	}{
		{
			name:       "正确密码验证",
			password:   password,
			storedHash: hash,
			want:       true,
		},
		{
			name:       "错误密码验证",
			password:   wrongPassword,
			storedHash: hash,
			want:       false,
		},
		{
			name:       "空密码验证",
			password:   "",
			storedHash: hash,
			want:       false,
		},
		{
			name:       "空哈希验证",
			password:   password,
			storedHash: "",
			want:       false,
		},
		{
			name:       "无效哈希格式",
			password:   password,
			storedHash: "invalid$hash$format",
			want:       false,
		},
		{
			name:       "不支持的算法",
			password:   password,
			storedHash: "$bcrypt$v=2a$12$somebcrypthash",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Argon2id.Verify(tt.password, tt.storedHash)
			if got != tt.want {
				t.Errorf("Verify() = %v, 预期 %v", got, tt.want)
			}
		})
	}
}

// TestHashUniqueness 测试 Hash 方法对相同密码生成不同哈希的能力
func TestHashUniqueness(t *testing.T) {
	password := "TestPassword123!"
	hashes := make(map[string]bool)
	numHashes := 100

	for i := 0; i < numHashes; i++ {
		hash, err := Argon2id.Hash(password)
		if err != nil {
			t.Fatalf("Hash() 返回了一个错误: %v", err)
		}

		if hashes[hash] {
			t.Errorf("Hash() 返回了重复的哈希 (第 %d 次): %s", i, hash)
		}

		hashes[hash] = true
	}
}

// TestEmptyPassword 测试对空密码的哈希和验证
func TestEmptyPassword(t *testing.T) {
	password := ""
	hash, err := Argon2id.Hash(password)
	if err != nil {
		t.Fatalf("Hash() 对空密码返回了一个错误: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash() 对空密码返回了一个空字符串")
	}

	// 验证空密码是否与生成的哈希匹配
	if !Argon2id.Verify(password, hash) {
		t.Error("Verify() 未能正确验证空密码")
	}

	// 验证非空密码不匹配空密码的哈希
	if Argon2id.Verify("nonempty", hash) {
		t.Error("Verify() 错误地验证了错误的密码")
	}
}

// TestEmptyHash 测试使用空哈希进行验证的行为
func TestEmptyHash(t *testing.T) {
	password := "TestPassword123!"
	emptyHash := ""

	if Argon2id.Verify(password, emptyHash) {
		t.Error("Verify() 错误地返回了空哈希的验证结果为真")
	}
}

// TestMalformedHash 测试对各种格式错误的哈希进行验证
func TestMalformedHash(t *testing.T) {
	password := "TestPassword123!"
	malformedHashes := []string{
		"$argon2id$v=19$m=65536,t=1,p=4$salt",            // 缺少哈希部分
		"$argon2id$v=19$m=65536,t=1,p=4",                 // 缺少盐和哈希部分
		"$argon2id$v=19$m=65536,t=1,p=4$salt$hash$extra", // 多余的部分
		"argon2id$v=19$m=65536,t=1,p=4$salt$hash",        // 缺少起始 $
		"$argon2id$v=19$m=65536,t=1,p=4$invalidsalt$",    // 无效的盐（假设）
		"$argon2id$v=19$m=65536,t=1,p=4$$hash",           // 缺少盐
		"$argon2id$v=19$m=invalid,t=1,p=4$salt$hash",     // 非整数内存参数
		"$argon2id$v=19$m=65536,t=invalid,p=4$salt$hash", // 非整数时间参数
		"$argon2id$v=19$m=65536,t=1,p=invalid$salt$hash", // 非整数线程参数
	}

	for _, storedHash := range malformedHashes {
		t.Run(fmt.Sprintf("MalformedHash_%s", storedHash), func(t *testing.T) {
			result := Argon2id.Verify(password, storedHash)
			if result {
				t.Errorf("Verify() 错误地对格式错误的哈希返回了真: %s", storedHash)
			}
		})
	}
}

// TestUnsupportedAlgorithm 测试对不支持的算法的哈希进行验证
func TestUnsupportedAlgorithm(t *testing.T) {
	password := "TestPassword123!"
	unsupportedHash := "$bcrypt$v=2a$12$somebcrypthash"

	if Argon2id.Verify(password, unsupportedHash) {
		t.Error("Verify() 错误地对不支持的算法返回了真")
	}
}
