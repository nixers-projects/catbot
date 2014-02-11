package last


import (
	"regexp"
	"net/http"
	"io/ioutil"
	"html"
)


func Last (user string) (string, error) {
	brl := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&limit=1&api_key=37c444dd2fb14904cca0bce2999e4a81&user="
	url := brl + user

	r,x := http.Get(url)
	if x != nil { return "", nil }
	defer r.Body.Close()

	b,x := ioutil.ReadAll(r.Body)
	if x != nil { return "", nil }
	if b == nil { return "", nil }
	buf := string(b)
	re := regexp.MustCompile("<name>.*?</name>")
	title := re.FindString(buf)
	re = regexp.MustCompile("<artist.*?\">.*?</artist>")
	artist := re.FindString(buf)
	re = regexp.MustCompile("<.*?>")
	artist = re.ReplaceAllLiteralString(artist, "")
	artist = html.UnescapeString(artist)
	title = re.ReplaceAllLiteralString(title, "")
	title = html.UnescapeString(title)

	if (title != "" && artist != "") {
		return (artist + " - " + title), nil
	}
	return "",nil
}
