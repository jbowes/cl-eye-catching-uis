# CL Eye-Catching User Interfaces
<!-- .element: style="text-align: left" -->

https://bit.ly/cli-ui
<!-- .element: style="float: left" -->

<div align="right">GopherCon</div>
<div align="right">August 29, 2018</div>

---

# About Me

<image src="/images/james.png" style="float: left; border:0; background: none" width="20%">

- James Bowes
- [<i class="fab fa-twitter"></i> @jrbowes](https://twitter.com/jrbowes)
- Technical Lead @ [Manifold](https://www.manifold.co)
- Worked on many CLI tools no-one uses any more
  (*up2date, yum*)
- Loves Cats
- Loves Terminals

<!-- .element: style="width: 70%; float: right" -->

--
<!-- .slide: data-menu-title="Spooky on her thinkpad" -->
<!-- .slide: data-background="/images/sandwich-cat.jpg" -->

--
<!-- .slide: data-menu-title="My terminals" -->
<!-- .slide: data-background="/images/terminal-desktop.png" -->

---
### You will learn how to
- make progress bars and spinners
- decorate text
- draw anywhere on the terminal
- collect input
- do fancy things with mice and images

...safely, for any operating system<span class="fragment">*</span>

--
<!-- .slide: data-menu-title="Caveats" -->
\* Recent versions of Mac OS, Linux, and Windows<span class="fragment">**</span>

<span class="fragment">\*\* For common terminals</span>

--
## Notably Absent

- Command line argument parsing
- Environment variables

Note:

command line argument / flag parsing is very important, but also divisive. In
layout and API. Pick your favorite! You can't go wrong.

---
## Not just for Go

Share these system calls and escape sequences with your non-gopher friends.

<span class="fragment"><i>Use a package for real code.</i> 😬</span>

--

## Go is awesome for CLIs

- Great support for Windows, Mac OS, Linux
- Easy cross compilation with `GOOS` and `GOARCH`
- Self-contained binary (as long as you don't need Cgo)

---

## Hierarchy of User Interfaces

- Command Line Interface (CLI)
- Text-based User Interface (TUI) <span class="fragment">☜</span>
- Graphical User Interface (GUI)

---
<!-- .slide: data-menu-title="Part 1: Characters" -->

# Part 1
<hr />
## ☞ Characters

---
<!-- Main content begins -->
# Carriage Return
### Your new secret weapon

--
<!-- .slide: data-menu-title="A runaway carriage" -->
<!-- .slide: data-background="/images/runaway-carriage.jpg" -->

Note:

Describe that it's the part on a typewriter that moved the paper back and
forth. Line feed would move down to the next line.

You could just carriage return without line feed and overwrite text.

--
### Building a progress bar
```go
func drawBar(percent int) {
    cols := 10
    prog := strings.Repeat("=", cols*percent/100.0)
    fmt.Printf("\rdemo progress: %3[1]d%% |%-[3]*[2]s|",
        percent, prog, cols)
}
```

<!-- .element: class="fragment" -->

--
### Building a progress bar (con't)
```go
fmt.Printf("\rdemo progress: %3[1]d%% |%-[3]*[2]s|",
    percent, prog, cols)
```
<ul>
<li class="fragment">`\r` Move to start of line </li>
<li class="fragment">`%3[1]d` Print the int value of arg 1</li>
<li class="fragment">`%%` Literal percent</li>
<li class="fragment">`-[3]*` Left justify and pad by the value of arg 3</li>
<li class="fragment">`[2]s` Print the string value of arg 2</li>
</ul>


---
# Unicode
### Your *other* New Secret Weapon

--

## Unicode Spinners

Carriage return ❤️ unicode

```go
var (
	braille = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
	clock   = []rune{'🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚', '🕛', '🕐', '🕑'}
)
```
<!-- .element: class="fragment" -->

--

## Spice up Your User Interactions

```shell
✅ Success!
⚠️ Your password needs more entropy
🌶 ???
```

---

## What C⏻uld P⏻ssibly G⏻ Wr⏻ng?

Note:

The power symbol is used here to be a valid char that doesn't have a
representation in a font. It will still look ok if the font supports it though!

--

## Missing Characters in a Typeface

```go
missing := "The power symbol is new in Unicode 9: ⏻"
fmt.Println(missing)
```

--

## Miscounting Multibyte Characters

```go
miscounted := "😀"
fmt.Println(miscounted, "has len", len(miscounted))
fmt.Println(miscounted, "has runes",
	utf8.RuneCount([]byte(miscounted)))
```

```shell
has len 4
has runes 1
```

--

## Wide Characters

```go
wide := "十"
fmt.Println(wide)
fmt.Println("01")

fmt.Println("runes:", utf8.RuneCount([]byte(wide)))
fmt.Println("width:", runewidth.StringWidth(wide))
```

```shell
十
01
runes: 1
width: 2
```

--

### Single-Width Characters That Render as Wide
```go
wrongWidth := "☛"
fmt.Println(wrongWidth)
fmt.Println("01")

fmt.Println("runes:", utf8.RuneCount([]byte(wrongWidth)))
fmt.Println("width:", runewidth.StringWidth(wrongWidth))

fmt.Printf("%sOver top\n", wrongWidth)
```

```shell
☛
01
runes: 1
width: 1
☛Over top
```

---
<!-- .slide: data-menu-title="Part 2: Escape Codes" -->

# Part 2
<hr />
## ☞ Escape Codes

---
<!-- .slide: data-menu-title="In-band signalling" -->

In-band signalling that starts with `\033[`.

Thank you, American National Standards Institute! <!-- .element: class="fragment" -->

---

# Text Decoration and Color

--

## Decorations
Mostly supported

--

## Character Size
Not well supported, but really neat!

--

## Color
So many options!

---

# Multiline Output
## With Cursor Movement
--

### Linear-feedback shift register screen clearing

![wolfenstein fizzelfade](/images/wolfenstein.gif) <!-- .element: style="border: 0;" width="45%" -->

http://fabiensanglard.net/fizzlefade/index.php

--

## So much more!

- Relative cursor movement
- Partial screen clearing

---

# Interlude
<hr />
## ☞ Can I Use...

---

## How bad can it be?

```
0kc5N/d|NkuKboHmhDdc4_.f~Y`zTevNbsHipCjn3^`mZlVkkRdkPkkM`
lLnlJkmIinIhoIlpImqIorImsJotJiuKkvKawLnwLdxMixMayMfyMjyM
nyMoyMczMMgzMMMhzNNizNjzNhzOizOkzOizPjzPhzQfzQczQmyRjyRg
yR`yShxSmwSmvSivS`vShuS`uSmtSftSctSlsTjsTesUlrVhrWdrXarY
jqZbq[gp\ko]in^gm^cl^ej^dh]f[mcXbaT-f~O{IbxBlt2ZiqPjnEnk
1YjiLmg0^jfOdl@nl/QomCjo.VmqJet-_`wUzLa}D.b`,]obWieSchOi
jMlKenJooIgqIdrJgsJctKotLguMluO`vPdvQkvQnvRawSgwShwTmwTa
xTexTixTlxT`yTgySfySiySmyS`zSczRRRbzRRnyRRjyReyRayRlxRkx
QfxQaxQlwQnwPiwPgwOawOovNivNgvMevM`vMkuLfuLntLjtLosLgsLi
```

That's a teapot in ReGIS

---

# How can we tell what is supported?

--

## Environment Variables

- `TERM=xterm-256color`
- `COLORTERM=truecolor`
- Terminal specific values

--

## terminfo

A database of terminal names (from `TERM`) and capabilities that started in
`ncurses`. Includes escape codes.

For better portability, use a terminfo database for in-band signalling.
<!-- .element: class="fragment" -->

--

## Rely on the user

Provide flags and configuration for color and interactivity

---

## Handling Windows

- Windows 10 can enable VT processing
- For other versions, wrap the output and parse escape codes

--

## Windows Subsystem for Linux

- ANSI escape codes and POSIX syscalls under Windows 👍
- The console related APIS don't work across environments 👎

---
<!-- .slide: data-menu-title="Part 3: System Calls" -->

# Part 3
<hr />
## ☞ System Calls

---
<!-- .slide: data-menu-title="Out-of-band signalling" -->
Out-of-band signalling through system calls.

Thank you, IEEE Computer Society 🤷 <!-- .element: class="fragment" -->

And Microsoft 🤷 <!-- .element: class="fragment" -->

---

## Detecting terminal size

Great for columnar output and wrapping.

---

## Multi-line interactive inline inputs

Raw mode

---
<!-- .slide: data-menu-title="Part 4: Potpourri" -->

# Part 4
<hr />
## ☞ Potpourri

---

## Fullscreen interfaces

Combine:
- Raw mode
- Direct tty access
- Alternate buffer

---

## Displaying graphics

- `SIXEL`: raster graphics from DEC
- `ReGIS`: vector graphics from DEC
- Custom formats for assorted modern terminals

---

## Capturing mouse input

- Broad support 👍
- Not supported in Windows VT Processing. You have to fall back to the input
  buffer API. 👎

Note:
Show the final example for 11-potpourri here.

---
<!-- Closing section begins -->

# Appendix
<hr />
## ☞ Learn More

--
# Great Libraries <!-- .element: style="text-align: left" -->

- [chzyer/readline](https://github.com/chzyer/readline)
- [manifoldco/promptui](https://github.com/manifoldco/promptui)
- [gdamore/tcell](https://github.com/gdamore/tcell)
- [mattn/go-runewidth](https://github.com/mattn/go-runewidth)

<!-- .element: style="float: left" -->

--
# Reading List
<!-- .element: style="text-align: left" -->

- [Windows console reference](https://docs.microsoft.com/en-us/windows/console/console-reference)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
- [termios(3)](http://man7.org/linux/man-pages/man3/termios.3.html)
- [XTerm Control Sequences](http://invisible-island.net/xterm/ctlseqs/ctlseqs.html)

<!-- .element: style="float: left" -->

---

# Thank you!
<!-- .element: style="text-align: left" -->

- [<i class="fab fa-twitter"></i> @jrbowes](https://twitter.com/jrbowes)
- [<i class="fab fa-github"></i> this presentation (https://bit.ly/cli-ui)](https://bit.ly/cli-ui)

<!-- .element: style="float: left" -->
