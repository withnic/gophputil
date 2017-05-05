package phputil

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"html"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/blowfish"
)

func AddcSlashes(s string, c string) string {
	r := strings.Split(c, "..")
	var res string
	min := []rune(r[0])
	max := []rune(r[len(r)-1])
	for _, v := range s {
		if min[0] <= v && v <= max[0] {
			res += "\\" + string(v)
		} else {
			res += string(v)
		}
	}

	return res
}

func AddSlashes(s string) string {
	r := strings.NewReplacer("'", "\\'", "\"", "\\\"", "\\", "\\\\")
	return r.Replace(s)
}

func Bin2Hex(s string) string {
	b := []byte(s[:])
	return hex.EncodeToString(b)
}

func Chop(s ...string) string {
	return rtrim(s)
}

func Chr(i int) string {
	for i < 0 {
		i += 256
	}
	i %= 256
	return string(rune(i))
}

func ChunkSplit(s string, l int, sep string) string {
	if len(s) < l {
		return s
	}

	res := s[:l] + sep
	tail := s[l:]

	for len(tail) > l {
		res += tail[:l] + sep
		tail = tail[l:]
	}
	res += tail

	return res
}

// TODO
func ConvertCyrString(s string, from string, to string) string {
	return ""
}

// TODO
func ConvertUudecode(s string) string {
	return ""
}

// TODO
func ConvertUuencode(s string) string {
	return ""
}

func CountChars(s string, i int) map[int]int {

	r := countChars(s)

	switch i {
	case 0:
		return r
	case 1:
		for n := 0; n < 255; n++ {
			if r[n] == 0 {
				delete(r, n)
			}
		}
		return r
	case 2:
		for n := 0; n < 255; n++ {
			if r[n] > 0 {
				delete(r, n)
			}
		}
		return r
	}

	return r
}

func CountChars34(s string, i int) string {
	r := countChars(s)
	var res string
	switch i {
	case 3:
		for n := 0; n < 255; n++ {
			if r[n] > 0 {
				res += string(n)
			}
		}
		return res
	case 4:
		for n := 0; n < 255; n++ {
			if r[n] == 0 {
				res += string(n)
			}
		}
		return res
	}
	return res
}

func countChars(s string) map[int]int {
	r := make(map[int]int)
	for i := 0; i < 255; i++ {
		r[i] = 0
	}

	for _, v := range s {
		r[int(v)]++
	}
	return r
}

func Crc32(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// TODO
func Crypt(s string, salt string) string {
	var res string

	switch len(salt) {
	/**
	case 2:
		// CRYPT_STD_DES
		c, _ := des.NewCipher([]byte(salt))
		encrypted := make([]byte, des.BlockSize)
		// DES で暗号化をおこなう
		c.Encrypt(encrypted, []byte(s))
		return string(encrypted)
	case 9:
	// CRYPT_EXT_DES
	**/
	case 12:
		// CRYPT_MD5
		h := md5.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	case 22:
		// TODO: PADDING salt. CRYPT_BLOWFISH
		cipher, _ := blowfish.NewSaltedCipher([]byte("key"), []byte(salt))
		if len(s)%blowfish.BlockSize != 0 {
			os.Exit(1)
		}
		var encrypted []byte
		cipher.Encrypt(encrypted, []byte(s))
		return string(encrypted)
	case 16:
		// TODO: if condition. CRYPT_SHA256
		if true {
			c := sha256.Sum256([]byte(s))
			return hex.EncodeToString(c[:])
		} else {
			// CRYPT_SHA51
			c := sha512.Sum512([]byte(s))
			return hex.EncodeToString(c[:])
		}

	}
	return res
}

func echo(s string) {
	fmt.Print(s)
}

func Explode(d string, s string, l int) []string {
	return strings.SplitN(s, d, l)
}

func Fprintf(w io.Writer, f string, a ...interface{}) int {
	n, _ := fmt.Fprintf(w, f, a[:]...)
	return n
}

// TODO
func GetHtmlTranslationTable(t int, f int, e string) map[string]string {
	var res map[string]string
	return res
}

// TODO
func Hebrev(s string, m int) string {
	return ""
}

// TODO
func Hebrevc(s string, m int) string {
	return ""
}

func Hex2Bin(s string) string {
	b, _ := hex.DecodeString(s)
	return string(b)
}

func HtmlEntityDecode(s string) string {
	return html.UnescapeString(s)
}

func HtmlEntities(s string) string {
	return html.EscapeString(s)
}

// TODO Fix. It's not strict.
func HtmlspecialcharsDecode(s string) string {
	return html.UnescapeString(s)
}

// TODO Fix. It's not strict.
func Htmlspecialchars(s string) string {
	return html.EscapeString(s)
}

func Implode(d string, s []string) string {
	return strings.Join(s, d)
}

func Join(d string, s []string) string {
	return Implode(d, s)
}

func Rtrim(s ...string) string {
	return rtrim(s)
}

func rtrim(s []string) string {
	if len(s) == 2 {
		r := strings.Split(s[1], "..")
		if len(r) == 2 {
			min := []rune(r[0])
			max := []rune(r[len(r)-1])
			str := s[0]
			t := len(str)
			for {
				for i := min[0]; i <= max[0]; i++ {
					str = strings.TrimSuffix(str, string(i))
				}
				if t == len(str) {
					break
				}
				t = len(str)
			}
			return str
		}
		return strings.TrimSuffix(s[0], s[1])
	}

	r := s[0]
	t := len(r)
	suffix := [5]string{" ", "\t", "\n", "\r", "\x0B"}

	for {
		for _, v := range suffix {
			r = strings.TrimSuffix(r, v)
		}
		if t == len(r) {
			break
		}
		t = len(r)
	}

	return r
}

func Ucfirst(s string) string {
	first := strings.ToUpper(string(s[0]))
	return first + s[1:]
}

func Lcfirst(s string) string {
	first := strings.ToLower(string(s[0]))
	return first + s[1:]
}

func Ucwords(p ...string) string {
	d := " "
	if len(p) > 1 {
		d = p[1]
	}
	words := strings.Split(p[0], d)
	var res string

	if len(words) > 1 {
		for i := 0; i < len(words); i++ {
			res += Ucfirst(words[i]) + d
		}
		res = strings.TrimSuffix(res, d)
	} else {
		res = p[0]
	}

	return res
}

func Nl2br(s string) string {
	r := strings.NewReplacer("\n\r", "\n", "\r\n", "\n", "\r", "\n", "\n", "<br>\n")
	return r.Replace(s)
}
