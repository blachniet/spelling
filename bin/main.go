package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

var ssmlTmpl *template.Template

func main() {
	ssmlTmpl = template.Must(template.New("ssml").Parse(`<speak>
  <emphasis level="strong">{{.Word}}</emphasis>
  <break time="400ms" />
  {{.Sentence}}
  <break time="600ms" />
  <emphasis level="strong">{{.Word}}</emphasis>
</speak>
`))

	var input string
	var outputDir string
	flag.StringVar(&input, "input", "", "input file")
	flag.StringVar(&outputDir, "output", "", "output directory")
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("err opening input file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		entry, err := parseLine(line)
		if err != nil {
			log.Fatalf("failed to parse line: '%v'", line)
		}

		err = writeSsmlFile(entry, outputDir)
		if err != nil {
			log.Fatalf("err writing ssml file: '%v'", line)
		}

		// Print the content that should go in index.md
		fmt.Println("- val: " + entry.Word)
		fmt.Println("  src: words/" + entry.Word + ".mp3")
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("err reading lines: %v", err)
	}
}

func parseLine(line string) (Entry, error) {
	spaceIdx := strings.Index(line, " ")
	if spaceIdx == -1 || spaceIdx < 2 || spaceIdx > 3 {
		return Entry{}, fmt.Errorf("unexpected index for first space: %v", spaceIdx)
	}

	openParenIdx := strings.Index(line, "(")
	if openParenIdx == -1 {
		return Entry{}, fmt.Errorf("missing opening paren")
	}

	word := strings.TrimSpace(line[spaceIdx+1 : openParenIdx])
	if len(word) == 0 {
		return Entry{}, fmt.Errorf("couldn't find word")
	}

	closeParenIdx := strings.Index(line, ")")
	if closeParenIdx == -1 {
		return Entry{}, fmt.Errorf("missing closing paren")
	}

	sentence := strings.TrimSpace(line[openParenIdx+1 : closeParenIdx-len(word)])
	if len(sentence) == 0 {
		return Entry{}, fmt.Errorf("couldn't find sentence")
	}

	return Entry{word, sentence}, nil
}

func writeSsmlFile(e Entry, output string) error {
	f, err := os.Create(path.Join(output, "words", e.Word+".ssml"))
	if err != nil {
		return fmt.Errorf("err creating file: %w", err)
	}
	defer f.Close()

	return ssmlTmpl.Execute(f, e)
}

type Entry struct {
	Word     string
	Sentence string
}
