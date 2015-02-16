package lib

import (
    "crypto/md5"
    "crypto/rand"
    "crypto/sha1"
    "encoding/base64"
    "encoding/hex"
    "io"
    "strings"
)

const (
//BASE64字符表,不要有重复
    base64Table        = "<>:;',./?~!@#$CDVWX%^&*ABYZabcghijklmnopqrstuvwxyz01EFGHIJKLMNOPQRSTU2345678(def)_+|{}[]9/"
    hashFunctionHeader = "com.itrustoor.zh"
    hashFunctionFooter = "80.12.05.00"
)

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s)) //使用zhifeiya名字做散列值，设定后不要变
    return hex.EncodeToString(h.Sum(nil))
}

/*获取 SHA1 字符串*/
func GetSHA1String(s string) string {
    t := sha1.New()
    t.Write([]byte(s))
    return hex.EncodeToString(t.Sum(nil))
}

/**
 * 获取一个Guid值
 */
func GetGuid() string {
    b := make([]byte, 48)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

var coder = base64.NewEncoding(base64Table)

/**
 * base64加密
 */
func Base64Encode(str string) string {
    var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
    return string([]byte(coder.EncodeToString(src)))
}

/**
 * base64解密
 */
func Base64Decode(str string) (string, error) {
    var src []byte = []byte(str)
    by, err := coder.DecodeString(string(src))
    return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}