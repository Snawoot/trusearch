# trusearch

CLI tool to perform advanced search on [unofficial rutracker.org (ex torrents.ru) XML database](https://rutracker.org/forum/viewtopic.php?t=5591249). It doesn't require mandatory conversion of unofficial XML into own indexed format. Binary builds are available for Windows/MacOS/Linux/\*BSD/Android.

## Installation

#### Binaries

Pre-built binaries are available [here](https://github.com/Snawoot/trusearch/releases/latest).

#### Build from source

Alternatively, you may install trusearch from source. Run the following within the source directory:

```
make install
```

## Modes of operation

* `scan` - Scan XML and apply JS function defined by script file
* `forums` - Scan XML and print CSV with forum IDs and names
* `split` - Divide XML file into smaller ones by Forum ID
* `help` - Print CLI synopsis

### Split

Example:

```sh
trusearch split --dir=/arc/user/tru ~/rutracker-20210601.xml
```

It may take a long time to split whole collection (about 10 minutes). However, after splitting search over specific forums runs within few seconds.

### Forums

Example:

```sh
trusearch forums ~/rutracker-20210601.xml > ~/forums.csv
```

### Scan

Search with inline script:

```sh
trusearch scan --inline 'let re = /Жанр:.*(adventure|приключение)/iu; (function (elem) { if (elem.Content.match(re)) { print("https://rutracker.org/forum/viewtopic.php?t=" + elem.ID) } })' /arc/user/tru/forum_1992.xml
```

Same with script in file:

```sh
trusearch scan 1.js /arc/user/tru/forum_1992.xml
```

, having `1.js` content as follows:

```js
let re = /Жанр:.*(adventure|приключение)/iu;

(function (elem) {
	if (elem.Content.match(re)) {
		print("https://rutracker.org/forum/viewtopic.php?t=" + elem.ID)
	}
})
```

See `trusearch --help` for more help on commands and `trusearch COMMAND --help` for help on specific command.

## JS runtime

trusearch uses JavaScript to allow user implement any matching or aggregation logic they want. JS interpreter used by program is [goja](https://github.com/dop251/goja), pure-Go implementation of ECMAScript 5.1. trusearch extends JS runtime with few native built-ins for user's convenience.

### Functions

| Function       | Arguments    | Return value | Description                                                 |
| -------------- | ------------ | ------------ | ----------------------------------------------------------- |
| `perror`       | `value, ...` | None         | Prints values to stderr                                     |
| `print`        | `value, ...` | None         | Prints values to stdout                                     |
| `strip_bbcode` | `string`     | `string`     | Renders text with BBCode tags into plain text (strips tags) |
