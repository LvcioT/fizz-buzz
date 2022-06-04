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
    * `nFrom`: runs the fizz-buzz alghoritm startingo from this number, `1` if not provided
    * `nTo`: runs the fizz-buzz alghoritm up to this number, `100` if not providedwhole series has to be generated
    * `print`: prints `series`|`end` just the last number or the whole series (series is however generated but not printed)
* no check for paramters correctness
* save the series in a data structure and print all when finished
* minimize externale libraries usage
* a small set of tests
* no explicit use of mod 15

### Future Enhancement

A simple script in root folder will build and run each project taking note of time and eventually memory consumed on a simple json file.
