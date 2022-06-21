- [Alias Bee](#alias-bee)
  - [Installation](#installation)
  - [Setup](#setup)
  - [Usage](#usage)
  - [Keep up to date](#keep-up-to-date)
  - [Syntax highlighting](#syntax-highlighting)
  - [Comments](#comments)
  - [Available platforms](#available-platforms)
  - [Caveats](#caveats)
  - [Build from source](#build-from-source)
 
# Alias Bee
This is a command line utility to help out with managing your aliases, functions and scripts. 

<img src="/res/screenshot.gif?raw=true" alt="Alias Bee Screenshot" width="100%"/>

![Alias Bee Screenshot](/res/screenshot.png?raw=true "Optional Title")

## Installation

To install, execute this command in your terminal:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/devgabrielcoman/scriptexchange-aliasengine/main/install.sh)"
```

This will download [this script](https://github.com/devgabrielcoman/scriptexchange-aliasengine/blob/main/install.sh). 

In turn, this will download a binary into a new folder on your system: `$HOME/.local/bin/scripthub/`. 

The binary depends on your platform. You can find all options [here](https://github.com/devgabrielcoman/scriptexchange-aliasengine/tree/main/dist).

Finally, in your `.bashrc`, `.zshrc`, `.profile` file, add the following line, to create a shorthand for the script:

```
alias bee='$HOME/.local/bin/scripthub/bee'
```

## Setup

Once installed, you can register the aliases, functions and scripts on your system. 

To register a file of aliases or functions:

```
bee --register /absolute/path/to/alias/file
```

For example, if you alread have a setup in your `.bashrc`, `.zshrc` or `.profile` file, it's as simple as:

```
bee --register ~/.bashrc
```

This will identify all of the **aliases** and **functions** and register them. 

If you want to register a whole script, it's similar, you just need to add the **-s** flag. 

```
bee --register /absolute/path/to/script.sh -s
```

## Usage

Once you've registered everything, it's a simple as typing `bee`. 

This will open the command line in interactive full screen mode. 

You can search for all of the items you registered, by name. 

If you press **ENTER** it will execute the script. 

## Keep up to date

Sometimes you may want to add or remove items from the files you registered.

Once you do that, a simple 

```
bee -u
```

will make sure Alias Bee is up to date.

## Syntax highlighting

Alias Bee doesn't do syntax highlighting by default. However, if you install [Bat](https://github.com/sharkdp/bat), Alias Bee will know how to automatically use it. The screenshots on this page were generated with this combination.

## Comments

Alias Bee can read comments associated with aliases and functions and display them in
the righ hand side code preview section.
All you need to do is add comments just above their definition. 

```
# this is an alias 
alias ll='ls -all'

# this is a function
hello() {
  echo "Hello World"
}
```

## Available platforms 

Alias Bee currently supports:
* Darwin, x86_64, arm64
* Linux, x86_64, arm64

## Caveats

When it comes to ingesting aliases from your  `.bashrc`, `.zshrc` or `.profile` files, Alias Bee can be pretty lenient. For example, the following are supported (even though they're not all valid):

```
alias my-alias = 'll -all'  # slighlty invalid
alias my-alias='ll -all'    # valid
```

However, functions need to be defined properly:

```
function test { echo "Hello" }    # improperly defined one-liner
function test() { echo "Hello"; } # improper style - either use the "function" keyword or "()" style


function test { echo "Hello"; }   # properly defined one-liner
test() { echo "Hello"; }          # another way to properly define functions 
```

## Build from source

If you'd like to build locally, these are the steps to follow:

You must have the [Go Programming Language](https://go.dev/) installed.

Additionally, you can also install [Bat](https://github.com/sharkdp/bat).

Download or clone this repo. 

Run `./local_install.sh`.

This should be useful if you run on an unsuported arch.
