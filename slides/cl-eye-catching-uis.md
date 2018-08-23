# CL Eye-Catching User Interfaces

https://github.com/jbowes/cl-eye-catching-uis

<div align="right">GopherCon</div>
<div align="right">August 29, 2018</div>

---

# About Me

<image src="/images/james.png" style="float: left; border:0; background: none" width="20%">

- James Bowes
- [@jrbowes](https://twitter.com/jrbowes)
- Technical Lead @ [Manifold](https://www.manifold.co)
- Worked on many CLI tools no-one uses any more
  (*up2date, yum*)
- Loves Cats
- Loves Terminals

<!-- .element: style="width: 70%; float: right" -->

--
<!-- .slide: data-background="/images/sandwich-cat.jpg" -->

--
<!-- .slide: data-background="/images/terminal-desktop.png" -->

---
### In 45 minutes you will learn how to
- make progress bars and spinners
- decorate text
- draw anywhere on the terminal
- collect input
- do fancy images

...safely, for any operating system<span class="fragment">*</span>

--
\* Recent versions of Mac OS, Linux, and Windows<span class="fragment">**</span>

<span class="fragment">\*\* For common terminals</span>

---
## Not just for Go

Share these system calls and escape sequences with your non-gopher friends.

<span class="fragment"><i>Use a package for real code.</i> 😬</span>

---

# Part 1
<hr />
## ☞ Characters

---
<!-- Main content begins -->
# Carriage Return
### Your new secret weapon

--
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

Progress bars with clock emoji, braille
checks and Xs.

--

## Spice up Your User Interactions

Spice it up with the pepper!

---

## What C⏻uld P⏻ssibly G⏻ Wr⏻ng?

Note:

The power symbol is used here to be a valid char that doesn't have a
representation in a font. It will still look ok if the font supports it though!

--

## Missing Characters in a Typeface

--

## Miscounting Multibyte Characters

--

## Wide Characters

--

## Single-Width Characters That Render as Wide

---

# Part 2
<hr />
## ☞ Escape Codes

---

In-band signalling that starts with `\033[`.

Thank you, American National Standards Institute! <!-- .element: class="fragment" -->

---

# Text Decoration and Color

--

## Decorations

--

## Character Size

--

## Color

---

## Multiline With Cursor Movement

--

### Linear-feedback shift register screen clearing

![wolfenstein fizzelfade](/images/wolfenstein.gif)

http://fabiensanglard.net/fizzlefade/index.php

---

# Interlude
<hr />
## ☞ Can I Use...

---

## Just how bad things can go

---

# How can we tell what is supported?

--

## Env vars

--

## terminfo

--

## Rely on the user :(

---

# Handling Windows

--

## Escape codes in Windows 10

--

## Windows Subsystem for Linux

---

## At least there's only one odd one out now!


---

# Part 3
<hr />
## ☞ System Calls

---
## The Client / Server Split

---

## Detecting terminal size for columnar output and text wrapping

---

## Cursor movement and input masking

---

## Multi-line interactive inline inputs

---

## Fullscreen interfaces

---

# Part 4
<hr />
## ☞ Potpourri

---

## Displaying graphics in terminals that support it

---

## Capturing mouse input

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

<!-- .element: style="float: left" -->


--
# Reading List <!-- .element: style="text-align: left" -->

- [Windows console reference](https://docs.microsoft.com/en-us/windows/console/console-reference)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
- [termios(3)](http://man7.org/linux/man-pages/man3/termios.3.html)

<!-- .element: style="float: left" -->

---

# Thank you! <!-- .element: style="text-align: left" -->

- [@jrbowes](https://twitter.com/jrbowes)
- [this presentation](github.com/jbowes/cl-eye-catching-uis)

<!-- .element: style="float: left" -->
