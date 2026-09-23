# Step 62 · Wrapping

### The problem

Long messages run past the right edge of the 40-column log box and are cut off by whatever is drawn next to them. They should wrap onto more lines instead. Wrapping at spaces is a small greedy algorithm: take words one by one and put each on the current line if it still fits, otherwise start a new line. A wrapped message takes several rows, and they have to be drawn in the right order even though the log is drawn from the bottom up.

>>> 1. In `messagelog.go`, add a function `func wrap(text string, width int) []string` that splits the text into words with `strings.Fields` and packs them into lines no longer than `width`.
>>> 2. In `renderMessages`, wrap each message to `width`, and draw its lines from the last one up to the first, each on its own row, stopping when the box is full.

!!! Messages longer than 40 characters continue on the next row inside the log box.

--- reveal

{{diff messagelog.go}}

- `strings.Fields` splits on any run of spaces and drops empty words.
- `len(line)+1+len(word) <= width` asks whether the word, plus one space, still fits on the current line.
- A message wrapped into three lines is drawn bottom line first, because the whole log is drawn upwards: the inner loop runs `j` from the last line down to 0.

--- end

%%% In `renderMessages`, loop over the wrapped lines from the first to the last instead. Long messages come out with their lines in the wrong order, the end above the beginning, because each line is drawn one row higher than the previous.
