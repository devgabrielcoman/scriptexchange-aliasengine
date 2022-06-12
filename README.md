# AliasEngine Bee
This is a command line utility to help out with managing your aliases, functions and scripts. 

## Installation

To install, execute this command in your terminal:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/devgabrielcoman/scriptexchange-aliasengine/main/install.sh)"
```

This will download [this script](https://github.com/devgabrielcoman/scriptexchange-aliasengine/blob/main/install.sh). 

In turn, this will download a binary into a new folder on your system: `$HOME/.local/bin/scripthub/`. 

The source code for the binary is [here](https://github.com/devgabrielcoman/scriptexchange-aliasengine/tree/main/aliasengine).

Finally, in your `.bashrc`, `.zshrc`, `.profile` file, add the following line, to create a shorthand for the script:

```
alias bee='$HOME/.local/bin/scripthub/bee'
```
