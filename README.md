# golang-merge-templates

This program is used to merge multiple Go template files.

## Installation

```shell
go install github.com/Napat/golang-merge-templates@latest
```

## Usage

```shell
golang-merge-templates $OUTOUT <MAIN_TEMPLATE> <TEMPLATE_1> [<$TEMPLATE_2> ... ]
```

### Sample

```shell
golang-merge-templates output.html sampledata/main.tpl sampledata/template_1.tpl sampledata/template_2.tpl
```
