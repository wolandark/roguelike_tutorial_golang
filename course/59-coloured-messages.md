# Step 59 · Coloured messages

### The problem

Every message looks the same, so an important one ("You died!") is as easy to miss as a routine one. Each message should carry a colour, chosen by whoever logs it. A message is then two things, a text and a colour, which makes it a small struct. And the colours themselves are best named once, in one file, so the palette can be changed in one place.

>>> 1. Create `colors.go` with a `var ( )` block of named colours: `colorWhite`, `colorBlack`, `colorRed`, `colorPlayerAtk`, `colorEnemyAtk`, `colorPlayerDie`, `colorEnemyDie` and `colorWelcomeText`, each a `tcell.NewRGBColor(...)`.
>>> 2. Create `messagelog.go` with a struct `Message` with fields `Text string` and `Color tcell.Color`.
>>> 3. In `engine.go`, change `Messages` to a `[]Message`, change `Log` to `func (e *Engine) Log(msg string, color tcell.Color)`, and draw each message in its colour.
>>> 4. Pass a colour to every `Log` call: `colorPlayerAtk` or `colorEnemyAtk` in `MeleeAction` (depending on who attacks), `colorPlayerDie` or `colorEnemyDie` in `Die`, and `colorWelcomeText` for the welcome in `main.go`.

!!! The welcome message in blue. Attacks in light grey (yours) and pink (theirs); deaths in red and orange.

--- reveal

{{file colors.go}}

- `0xFF` is a hexadecimal literal, 255. Colours are often written in hex because each pair of digits is one of red, green and blue.

{{file messagelog.go}}

{{diff engine.go}}

{{diff actions.go}}

- `color := colorEnemyAtk` then `if entity == engine.Player { color = colorPlayerAtk }` picks the colour by who is attacking.

{{diff fighter.go}}

{{diff main.go}}

--- end

%%% Change `colorWelcomeText` in `colors.go` to `tcell.NewRGBColor(0xFF, 0, 0xFF)`. The welcome turns magenta and nothing else changes: that is the point of naming colours in one place.
