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

Add the `words` in the frontmatter.

```yaml
words:
- val: bird
  src: words/bird.mp3
```

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
