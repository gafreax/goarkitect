# GoArkitect

This project gives developers the ability to describe and check the architecture of a project and check it is respected at any time.

## Desired usecases

if a folder:
- [ ] exist
- [ ] does not exist
- [ ] contains a specific file
- [ ] contains a specific set of files
- [ ] contains a files matching a regex
- [ ] contains a files matching a glob pattern
- [ ] is gitignored
- [ ] is gitcrypted

if a set of files:
- [x] exists
- [x] does not exist
- [ ] is gitignored
- [ ] is gitcrypted

if a file:
- [x] exists
- [x] does not exist
- [x] name matches a regex
- [ ] name matches a glob pattern
- [ ] content matches value
- [ ] content matches regex
- [ ] content matches template
- [ ] is gitignored
- [ ] is gitcrypted

if a package:
- [ ] depends on another package
- [ ] contains symbols matching a regex
- [ ] does not contain symbols matching a regex

if a struct:
- [ ] depends on a namespace
- [ ] embeds another struct
- [ ] embeds another interface
- [ ] have a name matching a pattern
- [ ] implements an interface
- [ ] depends on a namespace
- [ ] don't have dependency outside a namespace
- [ ] reside in a package

if a makefile:
- [ ] contains a set of targets
- [ ] does not contain a set of targets

if a json file:
- [ ] matches a json schema

if a yaml file:
- [ ] matches a json schema

if a go module:
- [ ] contains up-to-date dependencies

## Json schema to Golang structs libraries

- https://github.com/atombender/go-jsonschema
- https://github.com/a-h/generate
