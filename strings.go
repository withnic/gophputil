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

// http://php.net/manual/ja/function.addcslashes.php
func AddcSlashes(str string, charSet string) string {
	r := strings.Split(charSet, "..")
	var res string
	min := []rune(r[0])
	max := []rune(r[len(r)-1])
	for _, v := range str {
		if min[0] <= v && v <= max[0] {
			res += "\\" + string(v)
		} else {
			res += string(v)
		}
	}

	return res
}

// http://php.net/manual/ja/function.addslashes.php
func AddSlashes(str string) string {
	r := strings.NewReplacer("'", "\\'", "\"", "\\\"", "\\", "\\\\")
	return r.Replace(str)
}

// http://php.net/manual/ja/function.bin2hex.php
func Bin2Hex(str string) string {
	b := []byte(str[:])
	return hex.EncodeToString(b)
}

// http://php.net/manual/ja/function.chop.php
func Chop(s ...string) string {
	return trim(s, 1)
}

// http://php.net/manual/ja/function.chr.php
func Chr(ascii int) string {
	for ascii < 0 {
		ascii += 256
	}
	ascii %= 256
	return string(rune(ascii))
}

// http://php.net/manual/ja/function.chunk-split.php
func ChunkSplit(body string, chunklen int, end string) string {
	if len(body) < chunklen {
		return body
	}

	res := body[:chunklen] + end
	tail := body[chunklen:]

	for len(tail) > chunklen {
		res += tail[:chunklen] + end
		tail = tail[chunklen:]
	}
	res += tail

	return res
}

// TODO http://php.net/manual/ja/function.convert-cyr-string.php
func ConvertCyrString(str string, from string, to string) string {
	return ""
}

// TODO http://php.net/manual/ja/function.convert-uudecode.php
func ConvertUudecode(data string) string {
	return ""
}

// TODO http://php.net/manual/ja/function.convert-uuencode.php
func ConvertUuencode(data string) string {
	return ""
}

// http://php.net/manual/ja/function.count-chars.php
func CountChars(str string, mode int) map[int]int {
	r := countChars(str)

	switch mode {
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

//http://php.net/manual/ja/function.count-chars.php
func CountChars34(str string, i int) string {
	r := countChars(str)
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

// http://php.net/manual/ja/function.crc32.php
func Crc32(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// TODO http://php.net/manual/ja/function.crypt.php
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

// http://php.net/manual/ja/function.echo.php
func Echo(s string) {
	fmt.Print(s)
}

// http://php.net/manual/ja/function.explode.php
func Explode(d string, s string, l int) []string {
	return strings.SplitN(s, d, l)
}

// http://php.net/manual/ja/function.fprintf.php
func Fprintf(w io.Writer, f string, a ...interface{}) int {
	n, _ := fmt.Fprintf(w, f, a[:]...)
	return n
}

// TODO http://php.net/manual/ja/function.get-html-translation-table.php
func GetHtmlTranslationTable(t int, f int, e string) map[string]string {
	var res map[string]string
	return res
}

// TODO http://php.net/manual/ja/function.hebrev.php
func Hebrev(s string, m int) string {
	return ""
}

// TODO http://php.net/manual/ja/function.hebrevc.php
func Hebrevc(s string, m int) string {
	return ""
}

// http://php.net/manual/ja/function.hex2bin.php
func Hex2Bin(data string) string {
	b, _ := hex.DecodeString(data)
	return string(b)
}

// http://php.net/manual/ja/function.html-entity-decode.php
func HtmlEntityDecode(s string) string {
	return html.UnescapeString(s)
}

// http://php.net/manual/ja/function.htmlentities.php
func HtmlEntities(s string) string {
	return html.EscapeString(s)
}

// TODO Fix. It's not strict. http://php.net/manual/ja/function.htmlspecialchars-decode.php
func HtmlspecialcharsDecode(s string) string {
	return html.UnescapeString(s)
}

// TODO Fix. It's not strict. http://php.net/manual/ja/function.htmlspecialchars.php
func Htmlspecialchars(s string) string {
	return html.EscapeString(s)
}

// http://php.net/manual/ja/function.implode.php
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// http://php.net/manual/ja/function.join.php
func Join(glue string, pieces []string) string {
	return Implode(glue, pieces)
}

// http://php.net/manual/ja/function.lcfirst.php
func Lcfirst(s string) string {
	first := strings.ToLower(string(s[0]))
	return first + s[1:]
}

// http://php.net/manual/ja/function.levenshtein.php
func Levenshtein(str1 string, str2 string) int {
	d := make([][]int, len(str1)+1)
	for i := range d {
		d[i] = make([]int, len(str2)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(str2); j++ {
		for i := 1; i <= len(str1); i++ {
			if str1[i-1] == str2[j-1] {
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
	return d[len(str1)][len(str2)]
}

// TODO http://php.net/manual/ja/function.localeconv.php
func Localeconv() {

}

// http://php.net/manual/ja/function.ltrim.php
func Ltrim(s ...string) string {
	return trim(s, 2)
}

// http://php.net/manual/ja/function.md5-file.php
func Md5File(filename string) string {
	f, err := os.Open(filename)
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

// http://php.net/manual/ja/function.md5.php
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

// TODO: http://php.net/manual/ja/function.metaphone.php
func Metaphone(str string) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.money-format.php
func MoneyFormat(format string, number float64) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.nl-langinfo.php
func NlLanginfo(item int) string {
	return ""
}

// http://php.net/manual/ja/function.nl2br.php
func Nl2br(s string) string {
	r := strings.NewReplacer("\n\r", "\n", "\r\n", "\n", "\r", "\n", "\n", "<br>\n")
	return r.Replace(s)
}

// TODO: http://php.net/manual/ja/function.number-format.php
func NumberFormat(number float64) string {
	return ""
}

// http://php.net/manual/ja/function.ord.php
func Ord(s string) int {
	return int(s[0])
}

// http://php.net/manual/ja/function.parse-str.php
func ParseStr(encodedString string) map[string][]string {
	res := make(map[string][]string)
	// key=v
	queries := strings.Split(encodedString, "&")
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

// http://php.net/manual/ja/function.print.php
func Print(s string) int {
	fmt.Print(s)
	return 1
}

// http://php.net/manual/ja/function.printf.php
func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

// TODO: http://php.net/manual/ja/function.quoted-printable-decode.php
func QuotedPrintableDecode(s string) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.quoted-printable-encode.php
func QuotedPrintableEncode(s string) string {
	return ""
}

// http://php.net/manual/ja/function.quotemeta.php
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

// http://php.net/manual/ja/function.rtrim.php
func Rtrim(s ...string) string {
	return trim(s, 1)
}

// TODO: http://php.net/manual/ja/function.setlocale.php
func SetLocate(i int, l string) {
	loc := time.FixedZone(l, 9*60*60)
	time.Local = loc
}

//http://php.net/manual/ja/function.sha1-file.php
func Sha1File(filename string) string {
	f, err := os.Open(filename)
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

// http://php.net/manual/ja/function.sha1.php
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

// http://php.net/manual/ja/function.sprintf.php
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// http://php.net/manual/ja/function.sscanf.php
func Sscanf(s string, format string, a ...interface{}) int {
	n, _ := fmt.Sscanf(s, format, a...)
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

// http://php.net/manual/ja/function.str-ireplace.php
func StrIreplace(search string, replace string, subject string) string {
	res := strings.Replace(strings.ToLower(subject), search, replace, -1)
	res = strings.Replace(strings.ToLower(subject), search, replace, -1)
	return res
}

//TODO:  i 0:STR_PAD_RIGHT 1:STR_PAD_LEFT 2:STR_PAD_BOTH only STR_PAD_RIGHT support http://php.net/manual/ja/function.str-pad.php
func StrPad(input string, padLength int, padString string, i int) string {
	if len(input) > padLength {
		return input
	}

	for len(input) < padLength {
		input += padString
	}

	return input[:padLength]
}

// http://php.net/manual/ja/function.str-repeat.php
func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

// http://php.net/manual/ja/function.str-replace.php
func StrReplace(search string, replace string, subject string) string {
	return strings.Replace(subject, search, replace, -1)
}

// http://php.net/manual/ja/function.str-rot13.php
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

// http://php.net/manual/ja/function.str-shuffle.php
func StrShuffle(s string) string {
	rand.Seed(time.Now().UnixNano())
	dest := make([]byte, len(s))
	perm := rand.Perm(len(s))
	for i, v := range perm {
		dest[v] = s[i]
	}
	return string(dest)
}

// http://php.net/manual/ja/function.str-split.php
func StrSplit(s string, splitLength int) []string {
	var res []string

	if splitLength < 1 {
		log.Fatal("please input over 1")
		//return false
	}

	if len(s) < splitLength {
		return []string{s}
	}

	for len(s) > 0 {
		res = append(res, s[:splitLength])
		s = s[splitLength:]
	}

	return res
}

// TODO: http://php.net/manual/ja/function.str-word-count.php
func StrWordCount(s string, f string) {

}

// http://php.net/manual/ja/function.strcasecmp.php
func Strcasecmp(str1 string, str2 string) int {
	return Strcmp(strings.ToLower(str1), strings.ToLower(str2))
}

// http://php.net/manual/ja/function.strchr.php
func Strchr(haystack string, needle string) string {
	return StrStr(haystack, needle)
}

// http://php.net/manual/ja/function.strcmp.php
func Strcmp(str1 string, str2 string) int {
	return strings.Compare(str1, str2)
}

// TODO: http://php.net/manual/ja/function.strcoll.php
func Strcoll(str1 string, str2 string) int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strcspn.php
func Strcspn(subject string, mask string) int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strip-tags.php
func StripTags(s string, allowableTags string) string {
	return ""
}

// TODO: http://php.net/manual/ja/function.stripcslashes.php
func Stripcslashes(s string) string {
	return ""
}

// http://php.net/manual/ja/function.stripos.php
func Stripos(haystack string, needle string) int {
	return strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
}

// http://php.net/manual/ja/function.stripslashes.php
func Stripslashes(s string) string {
	re := regexp.MustCompile(`\\([^\\])`)
	return re.ReplaceAllString(s, `$1`)
}

// http://php.net/manual/ja/function.stristr.php
func StriStr(haystack string, needle string) string {
	i := Stripos(haystack, needle)
	if i >= 0 {
		return haystack[i:]
	}
	return string("-1")
}

// http://php.net/manual/ja/function.strlen.php
func Strlen(s string) int {
	return len(s)
}

// TODO: http://php.net/manual/ja/function.strnatcasecmp.php
func Strnatcasecmp(str1 string, str2 string) int {
	return 0
}

// TODO: http://php.net/manual/ja/function.strnatcmp.php
func Strnatcmp(f string, s string) int {
	return 0
}

// http://php.net/manual/ja/function.strncasecmp.php
func Strncasecmp(str1 string, str2 string, length int) int {
	t := length
	if len(str1) < t {
		t = len(str1)
	}
	if len(str2) < t {
		t = len(str2)
	}
	ff := str1[:t]
	ss := str2[:t]
	return Strcasecmp(ff, ss)
}

// http://php.net/manual/ja/function.strncmp.php
func Strncmp(str1 string, str2 string, length int) int {
	t := length
	if len(str1) < t {
		t = len(str1)
	}
	if len(str2) < t {
		t = len(str2)
	}
	ff := str1[:t]
	ss := str2[:t]
	return Strcmp(ff, ss)
}

// http://php.net/manual/ja/function.strpbrk.php
func Strpbrk(haystack string, charList string) string {
	i := strings.Index(haystack, charList)
	if i >= 0 {
		return haystack[i:]
	}
	return string(-1)
}

// http://php.net/manual/ja/function.strpos.php
func Strpos(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// http://php.net/manual/ja/function.strrchr.php
func Strrchr(haystack string, needle string) string {
	i := strings.LastIndex(haystack, needle)
	if i >= 0 {
		return haystack[i:]
	}

	return string(-1)
}

// http://php.net/manual/ja/function.strrev.php
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

func StrStr(h string, s string) string {
	i := strings.Index(h, s)
	if i >= 0 {
		return string(h[i:])
	}
	return string("-1")
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

// http://php.net/manual/ja/function.substr.php
func Substr(str string, start int) string {
	return str[start:]
}

//TODO: manymask http://php.net/manual/ja/function.trim.php
func Trim(str ...string) string {
	return trim(str, 0)
}

func Ucfirst(s string) string {
	first := strings.ToUpper(string(s[0]))
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

/** http://php.net/manual/ja/function.vfprintf.php
func Vfprintf(handle io.Writer, format string, a ) int {
	return
}
**/

/** http://php.net/manual/ja/function.vprintf.php
func Vprintf(format string, a) int {
	return
}
**/

/** http://php.net/manual/ja/function.vsprintf.php
func Vsprintf(format string, a) string {
	return
}
**/

// http://php.net/manual/ja/function.wordwrap.php
func Wordwrap(str string, width int, b string) []string {
	res := []string{}
	n := len(str) / width
	for i := 0; i < n; i++ {
		if len(str) > (i+1)*width {
			res = append(res, str[i*width:(i+1)*width]+b)
		} else {
			res = append(res, str[i*width:])
			break
		}
	}
	return res
}

// http://php.net/manual/ja/function.count-chars.php
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
		return func(s string, d string) string {
			return strings.Trim(s, d)
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
