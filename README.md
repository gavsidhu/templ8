# Templ8 - CLI Template Manager

Templ8 is a command-line tool that makes it easier to manage custom templates. It allows you to create, store, and use files or directories as templates, saving you time and effort.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)

## Installation

To use Templ8, you'll need to install it on your system. Make sure you have [Node.js](https://nodejs.org/) installed.

```bash
npm install -g @gavsidhu/templ8
```

## Usage

You can run Templ8 using the following syntax:

```bash
templ8 command [options]
```

## Commands

`add`

Add a template to the collection from a local file.

```bash
templ8 add (--dir <directory_path> | --file <file_path>) --name <template_name>
```

Example:

```bash
templ8 add --file example.md --name example-template
```

Options:

- `file` path to an existing file
- `dir` path to an existing directory
- `name` Specify a name for the template

`paste`

Paste the specified template into the current directory.

```bash
templ8 paste (--dir=true | --file=true) --name <template_name>
```

Example:

```bash
templ8 paste --file=true --name example.md
```

Options:

- `file` whether a template is a file. (default: false)
- `dir` whether a template is a directory. (default: false)
- `name` the name of the template you want to paste

`list`

List all stored templates.

```bash
templ8 list
```

`delete`

Remove a stored template.

```bash
templ8 delete (--dir=true | --file=true) --name <template_name>
```

Example:

```bash
templ8 delete --file=true --name example.md
```

Options:

- `file` whether the template to delete is a file. (default: false)
- `dir` whether the template to delete is a directory. (default: false)
- `template_name` the name of the template to delete.

`help`

Show general help information.

```bash
templ8 --help
```

Show information about a command

```bash
template command --help
```

## Example

Add a template:

```bash
templ8 add --dir exampleDir --name example-directory
```

Paste a template:

```bash
templ8 paste my-template
```

List all templates:

```bash
templ8 list
```

Delete a template:

```bash
templ8 delete --dir=true --name example-directory
```
