# clrscreen
Simple package to clear screen irrespective of OS.

## Basic usage

```
package mypackage

import "github/lentregu/screen"

.......

// gets a function to clear the screen in a MAC
clear := screen.NewClearScreenFunction(screen.DARWIN)

.....

clear()
