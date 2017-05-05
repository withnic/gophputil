package phputil

import (
	"fmt"
	"os"
)

func ExampleAddcSlashes() {
	fmt.Println(AddcSlashes("Start hoge test.!!", "A..z"))
	// Output: \S\t\a\r\t \h\o\g\e \t\e\s\t.!!
}

func ExampleAddSlashes() {
	fmt.Println(AddSlashes(`Start "hoge" test.!\ww ! `))
	// Output: Start \"hoge\" test.!\\ww !
}

func ExampleBin2Hex() {
	fmt.Println(Bin2Hex("hello hex data"))
	// Output: 68656c6c6f206865782064617461
}

func ExampleChop() {
	fmt.Println(Chop("Example string\x0A\x0A", "\x00..\x1F"))
	// Output: Example string
}

func ExampleRtrim() {
	fmt.Print(Rtrim("Example string  \t     "))
	// Output: Example string
}

func ExampleChr() {
	fmt.Println(Chr(833))
	// Output: A
}

func ExampleUcfirst() {
	fmt.Println(Ucfirst("hello hogehoge."))
	// Output: Hello hogehoge.
}

func ExampleLcfirst() {
	fmt.Println(Lcfirst("Hello hogehoge."))
	// Output: hello hogehoge.
}

func ExampleUcWords() {
	fmt.Println(Ucwords("Hello fugafuga hohoho."))
	fmt.Println(Ucwords("Hello-fugafuga-hohoho.", "-"))
	// Output:
	//Hello Fugafuga Hohoho.
	//Hello-Fugafuga-Hohoho.
}

func ExampleNl2br() {
	fmt.Println(Nl2br("Hello hogehoge.\nTest hoge."))
	// Output:
	//Hello hogehoge.<br>
	//Test hoge.
}

func ExampleChunk_Split() {
	fmt.Println(ChunkSplit("Hello2hogehoge. Test hoge.", 6, "\n"))
	// Output:
	//Hello2
	//hogeho
	//ge. Te
	//st hog
	//e.
}

func ExampleCount_chars() {
	v := CountChars("Two Ts and one F.", 1)
	for k, v := range v {
		fmt.Printf("%d:%d\n", k, v)
	}
	// Unorderd output:
	//32:4
	//70:1
	//115:1
	//101:1
	//100:1
	//97:1
	//46:1
	//110:2
	//111:2
	//119:1
	//84:2
}

func ExampleCount_chars34() {
	v := CountChars34("Two Ts and one F.", 3)
	fmt.Println(v)
	// Output:
	//.FTadenosw
}

func ExampleCrc32() {
	v := Crc32("Hello, world")
	fmt.Println(v)
	// Output:
	//3885672898
}
func ExampleExplode() {
	e := Explode(",", "hogehoge,fugafuga,test", 1000)
	for _, v := range e {
		fmt.Println(v)
	}
	// Output:
	//hogehoge
	//fugafuga
	//test
}

func ExampleFprintf() {
	n := Fprintf(os.Stdout, "%s %d", "hogefuga", 1)
	fmt.Println(n)
	// Output:
	//hogefuga 110
}

func ExampleHex2Bin() {
	fmt.Println(Hex2Bin("68656c6c6f206865782064617461"))
	// Output:
	//hello hex data
}

func ExampleHtmlEntityDecode() {
	fmt.Println(HtmlEntityDecode(`&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`))
	// Output:
	// "Fran & Freddie's Diner" <tasty@example.com>
}

func ExampleHtmlEntities() {
	fmt.Println(HtmlEntities(`"Fran & Freddie's Diner" <tasty@example.com>`))
	// Output:
	//&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
}

func ExampleImplode() {
	fmt.Println(Implode(",", []string{"hoge", "fuga"}))
	// Output:
	//hoge,fuga
}

func ExampleLevenshtein() {
	fmt.Println(Levenshtein("Hello world.", "Good by world."))
	// Output:
	//7
}
