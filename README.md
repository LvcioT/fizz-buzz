# Fizz-Buzz Kata

This project aims to explore several programming languages on the classical Fizz-Buzz kata.

## Project Aim

The aim is to explore each language on these aspects:

* project structure
* language echo-system
* syntactical sugar
* execution speed
* test suite

Thus, the actual implementation could be overthinked and not optimal.

### Implementation

Each langauge lives in its own folder, has its own Docker infrastructure (Alpine if possible).

Every Langauage should have:

* optional command line paramters:
    * `--n-from`: runs the fizz-buzz alghoritm startingo from this number, `1` if not provided
    * `--n-to`: runs the fizz-buzz alghoritm up to this number, `100` if not providedwhole series has to be generated
    * `--print-last`: prints just the last number and not the whole series
* save the series in a data structure and print all when finished
* a small set of tests

### Future Enhancement

A simple script in root folder will build and run each project taking note of time and eventually memory consumed on a simple json file.
