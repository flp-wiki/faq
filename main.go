package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/flp-wiki/faq/gif"
)

type Items struct {
	Path     string
	Question string
	Answer   string
	Redirect string
}

func main() {

	data := []Items{
		{
			Path:     "",
			Question: "What is this flp.wiki FAQ thing?",
			Answer:   “A clever way to answer questions that are asked often.”,
			Redirect: "https://github.com/djsime1/awesome-flipperzero/blob/main/FAQ.md",
		},
		{
			Path:     "like",
			Question: "Do you like answering the same questions over and over??",
			Answer:   "We like helping people learn but it does get a little repetitive. So we made this site to help.",
			Redirect: "https://github.com/djsime1/awesome-flipperzero/blob/main/FAQ.md",
		},
		{
			Path:     "nfc/cc",
			Question: "Can I clone my Credit Card with the flipper?",
			Answer:   "No, It has encryption why would you want to clone your card? Have you thought about fraud?",
			Redirect: "https://github.com/djsime1/awesome-flipperzero/blob/main/FAQ.md#why-does-the-nfc-feature-table-say-bank-cards-can-be-read",
		},
		{
			Path:     "nfc/mfc/emulation",
			Question: "Why isn't Mifare Classic emulation working?",
			Answer:   "Some card readers operate at slightly different frequencies and the flipper can't detect that like a card.",
			Redirect: "https://github.com/djsime1/awesome-flipperzero/blob/main/FAQ.md#why-isnt-mifare-classic-emulation-working",
		},
		{
			Path:     "bettse",
			Question: "Why is bettse a jerk?",
			Answer:   "He's not a jerk, he's just a little misunderstood. Some even call him NFC Gandalf",
			Redirect: "https://ericbetts.org",
		},
		{
			Path:     "nfc/mfc/mfkey",
			Question: "What is mfkey32v2?",
			Answer:   "A way to use math to recover Mifare Classic keys from a reader.",
			Redirect: "https://github.com/equipter/mfkey32v2#mfkey32-version-2",
		},
		{
			Path:     "nfc/mfc/dict/stuck",
			Question: "Mifare Classic read is stuck!",
			Answer:   "The LED is blinking, it is still working.  It has a lot of keys to test.",
			Redirect: "https://github.com/djsime1/awesome-flipperzero/blob/main/FAQ.md#why-does-it-take-so-long-to-read-a-mifare-classic",
		},
		{
			Path:     "nfc/support",
			Question: "Why doesn't flipper support NFC-B/NFC-F/NFC-V/NTAG Write/other?",
			Answer:   "The Flipper team is not large, and they focus on the featuers they thing are most popular first",
			Redirect: "https://docs.flipperzero.one/",
		},
	}

	pageTemplate := template.Must(template.ParseFiles("template.html"))

	for _, v := range data {
		os.MkdirAll("dist/"+v.Path, 0755)
		f, _ := os.Create("dist/" + v.Path + "/index.html")
		defer f.Close()
		pageTemplate.Execute(f, v)

		if err := gif.Generate(v.Question, v.Answer, v.Path); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}

	}
}
