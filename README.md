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

### Scan mode

In scan mode program expects provided script to be evaluated into a function. So, minimal example of such script is:

```js
(function (torrent) {})
```

For each torrent scanned torrent record trusearch invokes such function with a single argument holding object with torrent record:

| Property                    | Type   | Description                                                                                    |
| --------------------------- | ------ | ---------------------------------------------------------------------------------------------- |
| `torrent.ID`                | string | Topic ID at rutracker forum                                                                    |
| `torrent.RegisteredAt`      | string | Torrent registration date in same format as in XML                                             |
| `torrent.Size`              | string | Torrent size                                                                                   |
| `torrent.Torrent.Hash`      | string | Bittorrent info hash. Can be used to generate DHT magnet link, not depending on tracker        |
| `torrent.Torrent.TrackerID` | string | rutracker tracker server ID                                                                    |
| `torrent.Forum.ID`          | string | Forum ID                                                                                       |
| `torrent.Forum.Name`        | string | Forum name                                                                                     |
| `torrent.Content`           | string | Post contents in BBCode. Use `strip_bbcode` function if you need plain text with tags stripped |
| `torrent.Deleted`           | int    | `1` if deleted and `0` otherwise                                                               |

Script may store state between function invocations in variables or objects defined outside function. In examples presented above we reuse RegExp compiled once across all function invokations.

Also user may define optional `begin` and `end` functions in script. If defined, `begin()` will be invoked before iteration and `end()` will be invoked after. It's correct to define either of them, both, or none. Example:

```js
function begin() {
	print("begin")
}

function end() {
	print("end")
}

(function () {
	print("record")
})
```
