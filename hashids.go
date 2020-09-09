// https://github.com/speps/go-hashids 的方法封装，数字类型的ID转换为随机字符串ID
// 用法示例

//     package main

//     import (
//         "log"

//         "github.com/axiaoxin-com/goutils"
//     )

//     func main() {
//         salt := "my-salt:appid:region:uin"
//         minLen := 8
//         prefix := ""
//         h, err := goutils.NewHashids(salt, minLen, prefix)
//         if err != nil {
//             log.Fatal(err)
//         }
//         var id int64 = 1
//         strID, err := h.Encode(id)
//         if err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("int64 id %d encode to %s", id, strID)

//         int64ID, err := h.Decode(strID)
//         if err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("string id %s decode to %d", strID, int64ID)
//     }

// 运行结果：

//     go run example.go
//     2020/02/26 13:28:11 int64 id 1 encode to 8Gnejq6A
//     2020/02/26 13:28:11 string id 8Gnejq6A decode to 1

package goutils

import (
	"strings"

	"github.com/speps/go-hashids"
)

// Hashids 封装hashids方法
type Hashids struct {
	HashID     *hashids.HashID
	HashIDData *hashids.HashIDData
	prefix     string
}

// NewHashids 创建Hashids对象
// salt可以使用用户创建记录时的用户唯一身份标识+当前时间戳的字符串作为值
// minLength指定转换后的最小长度,随着数字ID的增大长度可能会变长
func NewHashids(salt string, minLength int, prefix string) (*Hashids, error) {

	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	return &Hashids{
		HashID:     h,
		HashIDData: hd,
		prefix:     prefix,
	}, nil
}

// Encode 将数字类型的ID转换为指定长度的随机字符串ID
// int64ID为需要转换的数字id，在没有自增主键ID时，可以采用当前用户已存在的记录数+1作为数字id，保证该数字在该用户下唯一
func (h *Hashids) Encode(int64ID int64) (string, error) {
	idstr, err := h.HashID.EncodeInt64([]int64{int64ID})
	if err != nil {
		return "", err
	}
	return h.prefix + idstr, nil
}

// Decode 将生成的随机字符串ID转为为原始的数字类型ID
func (h *Hashids) Decode(hashID string) (int64, error) {
	if h.prefix != "" {
		hashID = strings.TrimPrefix(hashID, h.prefix)
	}
	idSlice, err := h.HashID.DecodeInt64WithError(hashID)
	if err != nil {
		return 0, err
	}
	return idSlice[0], nil
}
