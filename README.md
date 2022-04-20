# go-tr

Implements basic features of the `tr` Unix command. Made for UNR's Spring 2022 CS 491 (Testing &amp; DevOps) course.

Unit tests are placed within their respective packages. Integration tests are placed in the `test` directory.

## Usage

```
Usage of tr:
tr [-c] [-s] string1 string2
tr -s [-c] string1
tr -d [-c] string1
tr -d -s [-c] string1 string2
After giving arguments, program reads from stdin.
To submit input, pass in the EOF character.
  -c    Complement the set of values specified by string1.
  -d    Delete all occurrences of input characters that are specified by string1.
  -s    Replace instances of repeated characters with a single character.
```

Pull from DockerHub: `docker pull quincinia/go-tr`

Running the image: `docker run -it quincinia/go-tr [flags] <string1> <string2>`

Running `docker run quincinia/go-tr [-h|-help]` will show the message above.

## Examples

Note: pressing `enter` WILL NOT progress the program. In order to proceed, you must press `enter` THEN `Ctrl+D`.

### Capitalize String

```
$ docker run -it quincinia/go-tr "a-z" "A-Z"
the quick brown fox jumped over the lazy dog
THE QUICK BROWN FOX JUMPED OVER THE LAZY DOG
```

### Delete all non-digits

```
$ docker run -it quincinia/go-tr -d -c "0-9"
my phone is 123-456-7890
1234567890
```
