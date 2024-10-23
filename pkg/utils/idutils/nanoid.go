package idutils

import (
	"crypto/rand"
	"math"
	"orca/pkg/errors"
)

type nanoid struct{}

var Nanoid = &nanoid{}

// defaultAlphabet 是默认使用的字母表，用于生成 ID 字符。
var defaultAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const (
	defaultSize = 21
)

// getMask 生成用于从随机字节中获取位以获取随机字符索引的位掩码。
// 例如：如果字母表有 6 = (110)_2 个字符，使用掩码 7 = (111)_2 就足够了。
func (n *nanoid) getMask(alphabetSize int) int {
	for i := 1; i <= 8; i++ {
		mask := (2 << uint(i)) - 1
		if mask >= alphabetSize-1 {
			return mask
		}
	}
	return 0
}

// Generate 是一个底层函数，用于更改字母表和 ID 大小。
func (n *nanoid) Generate(alphabet string, size int) (string, error) {
	chars := []rune(alphabet)

	if len(alphabet) == 0 || len(alphabet) > 255 {
		return "", errors.New("字母表不能为空且长度不能超过255个字符")
	}
	if size <= 0 {
		return "", errors.New("大小必须是正整数")
	}

	mask := n.getMask(len(chars))
	// 估算我们需要多少个随机字节来生成 ID，实际上可能需要更多，但这是在平均情况和最坏情况之间的折衷
	// 以权衡效率
	ceilArg := 1.6 * float64(mask*size) / float64(len(alphabet))
	step := int(math.Ceil(ceilArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for j := 0; ; {
		_, err := rand.Read(bytes)
		if err != nil {
			return "", err
		}
		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(len(chars)) {
				id[j] = chars[currByte]
				j++
				if j == size {
					return string(id[:size]), nil
				}
			}
		}
	}
}

// MustGenerate 与 Generate 功能相同，但在出错时会 panic。
func (n *nanoid) MustGenerate(alphabet string, size int) string {
	id, err := n.Generate(alphabet, size)
	if err != nil {
		panic(err)
	}
	return id
}

// New 生成一个安全的 URL 友好的唯一 ID，接受 ID 的长度作为可选参数，
// 但是默认的 ID 长度为 21。
func (n *nanoid) New(l ...int) (string, error) {
	var size int
	switch {
	case len(l) == 0:
		size = defaultSize
	case len(l) == 1:
		size = l[0]
		if size < 0 {
			return "", errors.New("ID 长度不能为负数")
		}
	default:
		return "", errors.New("参数不符合预期")
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]&63]
	}
	return string(id[:size]), nil
}

// Must 与 New 的功能相同，但是如果出现错误程序会 panic。
func (n *nanoid) Must(l ...int) string {
	id, err := n.New(l...)
	if err != nil {
		panic(err)
	}
	return id
}
