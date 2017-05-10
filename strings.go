package phputil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"html"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

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
	return trim(s, 1)
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

func Levenshtein(s string, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}
	}
	return d[len(s)][len(t)]
}

// TODO
func Localeconv() {

}

func Ltrim(s ...string) string {
	return trim(s, 2)
}

func Md5File(s string) string {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func Md5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

// TODO: http://php.net/manual/ja/function.metaphone.php
func Metaphone(s string) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.money-format.php
func MoneyFormat(f string, m int) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.nl-langinfo.php
func NlLanginfo() string {
	return ""
}

func Nl2br(s string) string {
	r := strings.NewReplacer("\n\r", "\n", "\r\n", "\n", "\r", "\n", "\n", "<br>\n")
	return r.Replace(s)
}

// TODO:
func NumberFormat(s string) string {
	return ""
}

func Ord(s string) int {
	return int(s[0])
}

func ParseStr(s string) map[string][]string {
	res := make(map[string][]string)
	// key=v
	queries := strings.Split(s, "&")
	for _, v := range queries {
		// 0:key , 1:v
		query := strings.Split(v, "=")
		if t := strings.Index(query[0], "[]"); t != -1 {
			// ak = key
			ak := query[0][:t]
			vv := strings.Replace(query[1], "+", " ", -1)
			res[ak] = append(res[ak], vv)
		} else {
			ak := query[0]
			res[ak] = append(res[ak], query[1])
		}
	}

	return res
}

func Print(s string) int {
	fmt.Print(s)
	return 1
}

func Printf(f string, a ...interface{}) {
	fmt.Printf(f, a...)
}

// TODO: http://php.net/manual/ja/function.quoted-printable-decode.php
func QuotedPrintableDecode(s string) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.quoted-printable-encode.php
func QuotedPrintableEncode(s string) string {
	return ""
}

func Quotemeta(s string) string {
	r := strings.NewReplacer(
		`.`, `\.`,
		`\`, `\\`,
		`+`, `\+`,
		`*`, `\*`,
		`?`, `\?`,
		`[`, `\[`,
		`^`, `\^`,
		`]`, `\]`,
		`(`, `\(`,
		`$`, `\$`,
		`)`, `\)`,
	)
	return r.Replace(s)
}

func Rtrim(s ...string) string {
	return trim(s, 1)
}

// TODO: http://php.net/manual/ja/function.setlocale.php
func SetLocate(i int, l string) {
	loc := time.FixedZone(l, 9*60*60)
	time.Local = loc
}

func Sha1File(s string) string {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// TODO http://php.net/manual/ja/function.similar-text.php
func SimilarText(f string, s string) int {
	return 0
}

// TODO http://php.net/manual/ja/function.soundex.php
func Soundex(s string) string {
	return ""
}

func Sprintf(f string, a ...interface{}) string {
	return fmt.Sprintf(f, a...)
}

func Sscanf(s string, f string, a ...interface{}) int {
	n, _ := fmt.Sscanf(s, f, a...)
	return n
}

// TODO http://php.net/manual/ja/function.str-getcsv.php
/**
func StrGetcsv(s string) [][]string {
	r := csv.NewReader(strings.NewReader(s))
	records, _ := r.ReadAll()
	return records
}
**/

func StrIreplace(n string, r string, s string) string {
	res := strings.Replace(strings.ToLower(s), n, r, -1)
	res = strings.Replace(strings.ToLower(s), n, r, -1)
	return res
}

//TODO:  i 0:STR_PAD_RIGHT 1:STR_PAD_LEFT 2:STR_PAD_BOTH only STR_PAD_RIGHT support
func StrPad(s string, l int, p string, i int) string {
	if len(s) > l {
		return s
	}

	for len(s) < l {
		s += p
	}

	return s[:l]
}

func StrReplace(n string, r string, s string) string {
	return strings.Replace(s, n, r, -1)
}

func StrRepeat(s string, i int) string {
	return strings.Repeat(s, i)
}

func StrRot13(s string) string {
	var res []rune
	var t rune
	for _, v := range s {
		t = v
		if 'a' < v && v < 'z' {
			t = v + 13
			if t > 'z' {
				t -= 26
			}
		}
		if 'A' < v && v < 'Z' {
			t = v + 13
			if t > 'Z' {
				t -= 26
			}
		}
		res = append(res, t)
	}

	return string(res)
}

func StrShuffle(s string) string {
	rand.Seed(time.Now().UnixNano())
	dest := make([]byte, len(s))
	perm := rand.Perm(len(s))
	for i, v := range perm {
		dest[v] = s[i]
	}
	return string(dest)
}

func StrSplit(s string, i int) []string {
	var res []string

	if i < 1 {
		log.Fatal("please input over 1")
		//return false
	}

	if len(s) < i {
		return []string{s}
	}

	for len(s) > 0 {
		res = append(res, s[:i])
		s = s[i:]
	}

	return res
}

// TODO: http://php.net/manual/ja/function.str-word-count.php
func StrWordCount(s string, f string) {

}

func Strcasecmp(f string, s string) int {
	return Strcmp(strings.ToLower(f), strings.ToLower(s))
}

func Strchr(h string, s string) string {
	return StrStr(h, s)
}

func Strcmp(f string, s string) int {
	return strings.Compare(f, s)
}

// TODO: http://php.net/manual/ja/function.strcoll.php
func Strcoll() int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strcspn.php
func Strcspn(s string, m string) int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strip-tags.php
func StripTags(s string, allow string) string {
	return ""
}

// TODO:
func Stripcslashes(s string) string {
	return ""
}

func Stripos(s string, n string) int {
	return strings.Index(strings.ToLower(s), strings.ToLower(n))
}

func Stripslashes(s string) string {
	re := regexp.MustCompile(`\\([^\\])`)
	return re.ReplaceAllString(s, `$1`)
}

func StriStr(h string, s string) string {
	i := Stripos(h, s)
	if i >= 0 {
		return h[i:]
	}
	return string("-1")
}

func StrStr(h string, s string) string {
	i := strings.Index(h, s)
	if i >= 0 {
		return string(h[i:])
	}
	return string("-1")
}

func Strlen(s string) int {
	return len(s)
}

// TODO: http://php.net/manual/ja/function.strnatcasecmp.php
func Strnatcasecmp(f string, s string) int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strnatcmp.php
func Strnatcmp(f string, s string) int {
	return 0
}

func Strncasecmp(f string, s string, l int) int {
	t := l
	if len(f) < t {
		t = len(f)
	}
	if len(s) < t {
		t = len(s)
	}
	ff := f[:t]
	ss := s[:t]
	return Strcasecmp(ff, ss)
}

func Strncmp(f string, s string, l int) int {
	t := l
	if len(f) < t {
		t = len(f)
	}
	if len(s) < t {
		t = len(s)
	}
	ff := f[:t]
	ss := s[:t]
	return Strcmp(ff, ss)
}

func Strpbrk(h string, c string) string {
	i := strings.Index(h, c)
	if i >= 0 {
		return h[i:]
	}
	return string(-1)
}

func Strpos(h string, n string) int {
	return strings.Index(h, n)
}

func Strrchr(h string, n string) string {
	i := strings.LastIndex(h, n)
	if i >= 0 {
		return h[i:]
	}

	return string(-1)
}

func Strrev(s string) string {
	i := len(s) - 1
	var res string
	for ; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

// http://php.net/manual/ja/function.strripos.php
func Strripos(haystack string, needle string, offset int) int {
	s := haystack
	if offset >= 0 {
		s = s[offset:]
	} else {
		s = s[len(s)-offset:]
	}
	i := strings.LastIndex(strings.ToLower(s), strings.ToLower(needle))

	if i >= 0 {
		return i + offset
	}
	return i
}

// http://php.net/manual/ja/function.strripos.php
func Strrpos(haystack string, needle string, offset int) int {
	s := haystack
	if offset >= 0 {
		s = s[offset:]
	} else {
		s = s[len(s)-offset:]
	}
	i := strings.LastIndex(s, needle)

	if i >= 0 {
		return i + offset
	}
	return i
}

// TODO: http://php.net/manual/ja/function.strspn.php
func Strspn(subject string, mask string) int {
	return 0
}

// http://php.net/manual/ja/function.strtok.php
func Strtok(str string, token string) string {
	return ""
}

// http://php.net/manual/ja/function.strtolower.php
func Strtolower(s string) string {
	return strings.ToLower(s)
}

//http://php.net/manual/ja/function.strtoupper.php
func Strtoupper(s string) string {
	return strings.ToUpper(s)
}

// http://php.net/manual/ja/function.strtr.php
func Strtr(str string, from string, to string) string {
	rep := []string{}
	for i := 0; i < len(to); i++ {
		if len(from) > i {
			rep = append(rep, string(from[i]))
			rep = append(rep, string(to[i]))
		}
	}
	r := strings.NewReplacer(rep...)
	return r.Replace(str)
}

//http://php.net/manual/ja/function.substr-compare.php
func SubstrCompare(main string, str string, offset int) int {
	l := len(main) - offset
	if len(str) > l {
		l = len(str)
	}
	e := offset + l
	return strings.Compare(main[offset:e], str)
}

// http://php.net/manual/ja/function.substr-count.php
func SubstrCount(haystack string, needle string, offset int) int {
	h := haystack[offset:]
	return strings.Count(h, needle)
}

// http://php.net/manual/ja/function.substr-replace.php
func SubstrReplace(str string, replacement string, start int) string {
	i := 0
	if start >= 0 {
		i = start
	} else {
		i = len(str) + start
	}
	return str[:i] + replacement
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

func trimfunc(i int) func(s string, d string) string {
	if i == 1 {
		return func(s string, d string) string {
			return strings.TrimSuffix(s, d)
		}
	} else if i == 2 {
		return func(s string, d string) string {
			return strings.TrimPrefix(s, d)
		}
	} else {
		// TODO
		return func(s string, d string) string {
			return strings.TrimPrefix(s, d)
		}
	}
}

func trim(s []string, i int) string {
	trims := trimfunc(i)
	if len(s) == 2 {
		r := strings.Split(s[1], "..")
		if len(r) == 2 {
			min := []rune(r[0])
			max := []rune(r[len(r)-1])
			str := s[0]
			t := len(str)
			for {
				for i := min[0]; i <= max[0]; i++ {
					str = trims(str, string(i))
				}
				if t == len(str) {
					break
				}
				t = len(str)
			}
			return str
		}
		return trims(s[0], s[1])
	}

	r := s[0]
	t := len(r)
	suffix := [5]string{" ", "\t", "\n", "\r", "\x0B"}

	for {
		for _, v := range suffix {
			r = trims(r, v)
		}
		if t == len(r) {
			break
		}
		t = len(r)
	}

	return r
}
