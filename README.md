# csv2json

## Description
csv2json is a program that can parse csv files to json. This program follows the Unix philosphy in that it has a very concise purpose. Because of this, the json is sent to stdout, to be consumed by another program in whatever way you see fit. 

## Getting Started

### Dependencies
* to build yourself, you must have Go v 1.13+ installed

### Installing
```
go install github.com/rnemeth90/csv2json@latest
```
### Executing program
```
./csv2json example.csv
```
## Help
If you need help, submit an issue

## To Do
- [ ] Read csv input from stdout (sent from another program)
- [ ] Add automated builds for binary distribution

## Version History
* 0.2
    * Various bug fixes and optimizations
* 0.1
    * Initial Release

## License
This project is licensed under the MIT License - see the LICENSE.md file for details
