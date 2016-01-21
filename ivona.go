package main
import (
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	ivona "github.com/jpadilla/ivona-go"
)
func main() {
	client := ivona.New("", "")
	text, err := ioutil.ReadFile("/home/aliaksandr/work/ivona/book.txt")
	if err != nil {
		log.Fatal(err)
	}

	arrayOfParagraphs := strings.Split(string(text), "\n")
	i := 0
	for _,paragraph := range arrayOfParagraphs {
		if (len(paragraph) < 4) { // against empty lines
			continue
		}
		log.Printf("%v\n", paragraph)
		options := ivona.NewSpeechOptions(paragraph)
		options.Voice.Language = "ru-RU"
		options.Voice.Name = "Maxim"
		options.Voice.Gender = "Male"
		options.OutputFormat.Codec = "OGG"
		r, err := client.CreateSpeech(options)
		if err != nil {
			log.Fatal(err)
		}

		i++
		file := fmt.Sprintf("/home/aliaksandr/work/ivona/audiobook/ivona/tts%04d.ogg", i) // files like 0001.ogg
		ioutil.WriteFile(file, r.Audio, 0644)
	}
}