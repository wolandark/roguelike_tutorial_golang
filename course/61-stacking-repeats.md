# Step 61 · Stacking repeats

### The problem

A fight fills the log with the same line: "Orc attacks Player for 1 hit points." three times in a row. Roguelikes stack repeats: one line with a counter, "Orc attacks Player for 1 hit points. (x3)". So a message gets a count, adding a message whose text equals the last one only bumps that count, and drawing shows the count when it is above 1.

>>> 1. In `messagelog.go`, add a field `Count int` to `Message`, and a method `func (m Message) FullText() string` that returns the text followed by ` (x<count>)` when `Count` is greater than 1, and just the text otherwise.
>>> 2. Give `AddMessage` a third parameter, `stack bool`. When `stack` is true and the last message has the same text, increase its `Count` and return; otherwise append a new message with `Count: 1`.
>>> 3. In `renderMessages`, draw `msg.FullText()` instead of `msg.Text`.
>>> 4. In `engine.go`, make `Log` pass `true` for `stack`.

!!! Hit an orc and let it hit back a few times: repeated lines stack with `(x2)`, `(x3)`.

--- reveal

{{diff messagelog.go}}

- `l.Messages[len(l.Messages)-1]` is the last element. `len(l.Messages) > 0` is checked first, because on an empty log there is no last element.
- `Count++` changes the element *in* the slice, because `l.Messages[i].Count` addresses the element directly. `msg := l.Messages[i]` followed by `msg.Count++` would change a copy.

{{diff engine.go}}

--- end

%%% In `AddMessage`, write `last := l.Messages[len(l.Messages)-1]` and then `last.Count++`. Repeats stop stacking: `last` is a copy of the element, and the slice still holds `Count: 1`.
