// nanoid_test.go
package idutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGenerateNormal 测试 Generate 方法在正常输入下生成符合预期的 ID。
func TestGenerateNormal(t *testing.T) {
	alphabet := string(defaultAlphabet)
	size := defaultSize

	id, err := Nanoid.Generate(alphabet, size)
	require.NoError(t, err, "Generate 应该没有错误")

	assert.Equal(t, size, len(id), "生成的ID长度应该等于指定的大小")

	// 检查每个字符是否在字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, char, "ID字符应该在字母表中")
	}
}

// TestGenerateCustomAlphabet 测试 Generate 方法使用自定义字母表生成 ID。
func TestGenerateCustomAlphabet(t *testing.T) {
	alphabet := "ABCDEF123456"
	size := 10

	id, err := Nanoid.Generate(alphabet, size)
	require.NoError(t, err, "Generate 应该没有错误")

	assert.Equal(t, size, len(id), "生成的ID长度应该等于指定的大小")

	// 检查每个字符是否在自定义字母表中
	for _, char := range id {
		assert.Contains(t, alphabet, string(char), "ID字符应该在自定义字母表中")
	}
}

// TestGenerateEmptyAlphabet 测试 Generate 方法使用空字母表的情况。
func TestGenerateEmptyAlphabet(t *testing.T) {
	alphabet := ""
	size := 10

	id, err := Nanoid.Generate(alphabet, size)
	assert.Empty(t, id, "生成的ID应该为空")
	assert.Error(t, err, "Generate 应该返回错误")
}

// TestGenerateOversizedAlphabet 测试 Generate 方法使用超过255字符的字母表。
func TestGenerateOversizedAlphabet(t *testing.T) {
	alphabet := ""
	for i := 0; i < 256; i++ {
		alphabet += "A"
	}
	size := 10

	id, err := Nanoid.Generate(alphabet, size)
	assert.Empty(t, id, "生成的ID应该为空")
	assert.Error(t, err, "Generate 应该返回错误")
}

// TestGenerateInvalidSize 测试 Generate 方法使用无效的 ID 长度。
func TestGenerateInvalidSize(t *testing.T) {
	alphabet := string(defaultAlphabet)
	size := 0

	id, err := Nanoid.Generate(alphabet, size)
	assert.Empty(t, id, "生成的ID应该为空")
	assert.Error(t, err, "Generate 应该返回错误")
}

// TestGenerateUniqueness 测试 Generate 方法生成 ID 的唯一性。
func TestGenerateUniqueness(t *testing.T) {
	alphabet := string(defaultAlphabet)
	size := defaultSize
	count := 10000
	ids := make(map[string]struct{})

	for i := 0; i < count; i++ {
		id, err := Nanoid.Generate(alphabet, size)
		require.NoError(t, err, "Generate 应该没有错误")
		assert.Len(t, id, size, "生成的ID长度应该等于指定的大小")
		_, exists := ids[id]
		assert.False(t, exists, "生成的ID应该是唯一的")
		ids[id] = struct{}{}
	}
}

// TestMustGenerateNormal 测试 MustGenerate 方法在正常输入下生成正确的 ID。
func TestMustGenerateNormal(t *testing.T) {
	alphabet := string(defaultAlphabet)
	size := defaultSize

	id := Nanoid.MustGenerate(alphabet, size)

	assert.Equal(t, size, len(id), "生成的ID长度应该等于指定的大小")

	// 检查每个字符是否在字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, char, "ID字符应该在字母表中")
	}
}

// TestMustGeneratePanic 测试 MustGenerate 方法在错误输入下是否会 panic。
func TestMustGeneratePanic(t *testing.T) {
	alphabet := ""
	size := -1

	assert.Panics(t, func() {
		Nanoid.MustGenerate(alphabet, size)
	}, "MustGenerate 应该在错误输入时 panic")
}

// TestNewDefault 测试 New 方法使用默认参数生成 ID。
func TestNewDefault(t *testing.T) {
	id, err := Nanoid.New()
	require.NoError(t, err, "New 应该没有错误")

	assert.Equal(t, defaultSize, len(id), "生成的ID长度应该等于默认大小")

	// 检查每个字符是否在默认字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, char, "ID字符应该在字母表中")
	}
}

// TestNewCustomSize 测试 New 方法使用自定义长度生成 ID。
func TestNewCustomSize(t *testing.T) {
	size := 30

	id, err := Nanoid.New(size)
	require.NoError(t, err, "New 应该没有错误")

	assert.Equal(t, size, len(id), "生成的ID长度应该等于指定的大小")

	// 检查每个字符是否在默认字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, char, "ID字符应该在字母表中")
	}
}

// TestNewInvalidParameters 测试 New 方法使用无效参数的情况。
func TestNewInvalidParameters(t *testing.T) {
	// 测试负数长度
	_, err := Nanoid.New(-5)
	assert.Error(t, err, "New 应该返回错误对于负数长度")

	// 测试多个参数
	_, err = Nanoid.New(10, 20)
	assert.Error(t, err, "New 应该返回错误对于多个参数")
}

// TestNewUniqueness 测试 New 方法生成 ID 的唯一性。
func TestNewUniqueness(t *testing.T) {
	size := defaultSize
	count := 10000
	ids := make(map[string]struct{})

	for i := 0; i < count; i++ {
		id, err := Nanoid.New(size)
		require.NoError(t, err, "New 应该没有错误")
		assert.Len(t, id, size, "生成的ID长度应该等于指定的大小")
		_, exists := ids[id]
		assert.False(t, exists, "生成的ID应该是唯一的")
		ids[id] = struct{}{}
	}
}

// TestMustNewNormal 测试 Must 方法在正常输入下生成正确的 ID。
func TestMustNewNormal(t *testing.T) {
	id := Nanoid.Must()

	assert.Equal(t, defaultSize, len(id), "生成的ID长度应该等于默认大小")

	// 检查每个字符是否在默认字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, char, "ID字符应该在字母表中")
	}
}

// TestMustNewPanic 测试 Must 方法在错误输入下是否会 panic。
func TestMustNewPanic(t *testing.T) {
	// 由于 Must 方法调用 New 方法时会传递参数，以下测试将通过调用 Must 方法传递无效参数来触发 panic。
	assert.Panics(t, func() {
		// 这里调用 Nanoid.Must() 并传递一个无效的负数长度。
		// 由于 Must 方法的实现中参数是可变参数，因此需要传递负数长度。
		// 例如：Nanoid.Must(-10)
		Nanoid.Must(-10)
	}, "Must 应该在错误输入时 panic")
}

// TestConcurrency 测试 ID 生成方法在高并发环境下的表现。
func TestConcurrency(t *testing.T) {
	alphabet := string(defaultAlphabet)
	size := defaultSize
	concurrency := 100
	generationsPerGoroutine := 1000
	ids := make(chan string, concurrency*generationsPerGoroutine)

	// 使用 Generate 方法生成 ID
	for i := 0; i < concurrency; i++ {
		go func() {
			for j := 0; j < generationsPerGoroutine; j++ {
				id, err := Nanoid.Generate(alphabet, size)
				if err == nil {
					ids <- id
				}
			}
		}()
	}

	// 收集 ID 并检查唯一性
	idMap := make(map[string]struct{})
	for i := 0; i < concurrency*generationsPerGoroutine; i++ {
		id := <-ids
		_, exists := idMap[id]
		assert.False(t, exists, "生成的ID应该是唯一的")
		idMap[id] = struct{}{}
	}
}

// TestGenerateMask 测试 getMask 方法在各种字母表大小下的表现。
func TestGenerateMask(t *testing.T) {
	tests := []struct {
		alphabetSize int
		expectedMask int
	}{
		{alphabetSize: 2, expectedMask: 3},
		{alphabetSize: 3, expectedMask: 3},
		{alphabetSize: 4, expectedMask: 3},
		{alphabetSize: 5, expectedMask: 7},
		{alphabetSize: 8, expectedMask: 7},
		{alphabetSize: 9, expectedMask: 15},
		{alphabetSize: 16, expectedMask: 15},
		{alphabetSize: 17, expectedMask: 31},
		{alphabetSize: 32, expectedMask: 31},
		{alphabetSize: 33, expectedMask: 63},
		{alphabetSize: 64, expectedMask: 63},
		{alphabetSize: 255, expectedMask: 255},
	}

	for _, test := range tests {
		mask := Nanoid.getMask(test.alphabetSize)
		assert.Equal(t, test.expectedMask, mask, "对于字母表大小 %d，掩码应该是 %d", test.alphabetSize, test.expectedMask)
	}
}
