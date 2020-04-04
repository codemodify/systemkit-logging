# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg)  Logging
[![](https://img.shields.io/github/v/release/codemodify/systemkit-logging?style=flat-square)](https://github.com/codemodify/systemkit-logging/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-logging?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-logging?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-logging/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-logging?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-logging?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-logging)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-logging)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-logging?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-logging?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-logging?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-logging?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-logging?style=flat-square)


#### The missing robust, flexible, complete and advanced logging framework for Go
- #### See tests for simple and advanced concurrency usages
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-logging
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
`KeepOnlyLogs`(logUntil LogType) | 
`Tracef`(format `string`, v ...`interface`{}) | 
`Panicf`(format `string`, v ...`interface`{}) | 
`Fatalf`(format `string`, v ...`interface`{}) | 
`Errorf`(format `string`, v ...`interface`{}) | 
`Warningf`(format `string`, v ...`interface`{}) | 
`Infof`(format `string`, v ...`interface`{}) | 
`Successf`(format `string`, v ...`interface`{}) | 
`Debugf`(format `string`, v ...`interface`{}) | 
`Trace`(v ...`interface`{}) | 
`Panic`(v ...`interface`{}) | 
`Fatal`(v ...`interface`{}) | 
`Error`(v ...`interface`{}) | 
`Warning`(v ...`interface`{}) | 
`Info`(v ...`interface`{}) | 
`Success`(v ...`interface`{}) | 
`Debug`(v ...`interface`{}) | 

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) How It Works

The concept is as follows:
- `LogEntry` is logged and sent into a pipe
- Each block in the transformation pipe may apply a "transformation" to the `LogEntry` and pass it along

## Example A: consider the defined pipe below form left to right
What happens is that a `LogEntry` is sent to the first `transformer` in the pipe that will format the log message,
then the next one will save it to a file.

PIPE	|				| Formatters	| Persisters
:--		| ---:			| :---			| :---
DATA	| `LogEntry`	| `Simple`		| `File`

## Example B: consider the defined pipe below form left to right
What happens is that a `LogEntry` is sent to the first `transformer` in the pipe that will format the log message,
then the next one will send it to an array of persisters, one will save to file and the other one will display in
the console.

PIPE	|				| Formatters	| Mixers	| Persisters
:--		| ---:			| :---			| :---		| :---
DATA	| `LogEntry`	| `Simple`		| `Multi`	| `Console`, `File`

## Example C: consider the defined pipe below form left to right
What happens is that a `LogEntry` is sent to the first `transformer` in the pipe that will send the same `LogEntry`,
to multiple "loggers" each having their own transformation pipes.
`Pipe-1` will format the log entry using one formatter and then save it to a file.
`Pipe-2` will format the log entry using a different formatter and then send it to aws storage.

PIPE	|				| Pipe		| Mixers	| Pipes					| Formatters		| Persisters
:--		| ---:			| :---		| :---		| :---					| :---				| :---
DATA	| `LogEntry`	| `Pipe-0`	| `Multi`	| `Pipe-1`, `Pipe-2`	| `Simple`, `AWS`	| `File`, `AWS`

