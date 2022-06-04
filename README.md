# Fizz-Buzz Kata

This project aims to explore several programming languages on the classical Fizz-Buzz kata.

## Project Aim

The aim is to explore each language on these aspects:

* project structure
* language echo-system
* syntactical sugar
* execution speed
* test suite

### Implementation

Each langauge lives in its own folder, has its own Docker infrastructure (Alpine if possible).

Every Langauage should have:

* optional command line paramters:
    * `--upto=number`: runs the fizz-buzz alghoritm up to this number, `100` if not provided
    * `--no-print`: avoid on screen print of the whole series, but the whole series has to be generated
* a small set of tests

### Future Enhancement

A simple script in root folder will build and run each project taking note of time and eventually memory consumed on a simple json file.
