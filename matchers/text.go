package matchers

var (
	TypeTxt = newType("txt", "text/plain")
	//TypeHTML     = newType("html", "text/html")
	//TypeXML     = newType("xml", "text/xml")
)

var Text = Map{
	TypeTxt: Txt,
}

func Txt(b []byte) bool {
	whiteByte := 0
	if len(b) < 100 {
		return false
	}
	for i := 0; i < 100; i++ {
		if (b[i] >= 0x20 && b[i] <= 0xff) ||
			b[i] == 9 ||
			b[i] == 10 ||
			b[i] == 13 {
			whiteByte++
		} else if b[i] <= 6 || (b[i] >= 14 && b[i] <= 31) {
			return false
		}
	}
	if whiteByte >= 1 {
		return true
	}
	return false
}
