# Spelling

## Tasks

### Deploy

Commit your changes to the `master` branch then run the `deploy.sh` script.

```bash
./deploy.sh
```

### Create new word list

Create a new Hugo page bundle.

```bash
hugo new <yyyyMMdd>
```

Build a text file that looks like this to specify the words and sample use.

```plain
1. down (The rain came down hard. down)
2. too (There was too much rain for playing outside. too)
3. work (The Superkids like to work on projects. work)
4. many (The kids had many ideas. many)
5. first (Sal wanted to play leapfrog first. first)
6. try (Doc wanted to try all the games. try)
7. jelly (Did the kids feed the dragon peanut butter and jelly? jelly)
8. funny (Odd Socks is a funny name for a game. funny)
9. say (The winning team will say, “Hooray!” say)
10. play (What games will you play on the next rainy day? play)
```

Generate the SSML files and frontmatter:

```bash
go run bin/main.go -input <input-file> -output content/<yyyyMMdd>
```

Add the frontmatter to the `content/<yyyyMMdd>`.

Generate audio files for each word (see instructions in the section below).

### Generate audio file for a word

[Install Google Cloud SDK][3]. Follow the [Quickstart: Using the command line][4]
guide to gain access to Cloud Text-To-Speech.

Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.

```bash
export GOOGLE_APPLICATION_CREDENTIALS=<path-to-credentials-file>
```

Create an `.ssml` file for the new word.

```bash
cat << EOF > were.ssml
<speak>
  <emphasis level="strong">were</emphasis>
  <break time="400ms" />
  Sun and Wind were in a contest.
  <break time="600ms" />
  <emphasis level="strong">were</emphasis>
</speak>
EOF
```

Run `bin/ssml2mp3.sh` to create an `.mp3`.

```bash
# Creates were.mp3
bin/ssml2mp3.sh were.ssml
```

## Licenses

This project uses [revealjs][1]. See its license [here][2]. 

[1]: https://revealjs.com/
[2]: https://github.com/hakimel/reveal.js/blob/master/LICENSE
[3]: https://cloud.google.com/sdk/docs/install
[4]: https://cloud.google.com/text-to-speech/docs/quickstart-protocol
