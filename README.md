# Nano ID CLI
The Nano ID CLI is a command line tool for generating Nano IDs. 

With Nano ID CLI you can quickly generate one or more Nano ID strings of a specific size by means of command line. Of course, you can also specify a custom alphabet.

## Usage
By default, this allows you to use the default alphabet to generate a Nano ID string of size 21:
```bash
nanoid
```

#### Custom Size
If you need to generate a Nano ID string of the specified size, then you can use the `-s` option:
```bash
nanoid -s <size>
```

#### Custom Alphabet
If you need to use a custom alphabet to generate the Nano ID string, you can use the `-a` option:
```bash
nanoid -a <alphabet>
```
The length of the custom alphabet string must be less than or equal to 256 characters.

#### Generate Multiple
If you need to generate more than one Nano ID string at a time, then you can use the `-n` option:
```bash
nanoid -n <count>
```

## Install
You can get the latest pre-built binary executable at https://github.com/nobody-night/nanoid-cli/releases

<hr>

Released under the [MIT License](LICENSE).
