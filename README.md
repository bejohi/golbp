# golbp
An implementation of the [local-binary-pattern (lbp)](https://en.wikipedia.org/wiki/Local_binary_patterns)
algorithm in go.

## Features
1. Convert a jpg or png rgba image to a grayscale image.
2. Create a lbp 2d matrix of byte values from an given img.
3. Create a uniform 2d matrix of bool values from a given img.
4. Convert byte and bool 2 matrix back to an img and save it.

## Module overview
1. __cmd:__ provides summary functions to run operations as bundles (not so important).
2. __helper:__ provides runtime helper functions, such as a lightweight logging provider (not so important).
3. __imageCalc:__ provides image conversion functions (important if you don't have grayscaled images already).
4. __lbpCalc:__ the core of this repository, where lbp and lbp-uniform calculation functions are placed (important).
5. __model:__ holds structs and interfaces for meta informations, such as img types, or lbp-uniform types.

## Supported uniform types
_A list of all lbp-uniform types we already provide (with a own struct)_
1. "End of edges" (e.g. 0b01100000, 0b11000000)
