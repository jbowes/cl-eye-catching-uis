# Abstract

Command line interfaces can be as interactive, exciting, and beautiful as any
desktop or web app. This talk will cover techniques across multiple operating
systems and terminals. It will cover everything from progress bars and color to
mouse input and animated graphics on the command line.

# Description

Command line interfaces (CLIs) can be as interactive, exciting, and beautiful
as any desktop or web app. Like when building web apps, building an interactive
CLI affords the developer many features, but comes with a frightening matrix of
incompatibilities and quirks. This talk will teach you how to use the many
features and techniques available for building interactive CLs, from progress
bars and color to mouse input and animated graphics on the command line. You
will learn how to detect when these features are available in a terminal, and
how to ensure your program works across different operating systems and
terminals.

# Outline

Introduction / appetizer examples of some command line interfaces
 - Progress bars, color, prompts
 - Inputs with editing, field masking
 - Full screen vs inline interactivity
 - Images and animations (pngs/gifs)
 - Mouse input

A brief summary of incompatibilities
 - Differences in operating systems and terminal programs
 - Show how bad things can look when the control codes are wrong

Everything after will include code samples, and discuss compatibility and
cross-platform support
 - Displaying color and text formatting
 - Building a single-line updating progress bar/spinner
 - Detecting terminal size for columnar output and text wrapping
 - Multi-line interactive inline inputs
 - Cursor movement and input masking
 - Fullscreen interfaces
 - Displaying graphics in terminals that support it
 - Capturing mouse input
