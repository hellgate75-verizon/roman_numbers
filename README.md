<p align="right">
 <img src="https://github.com/hellgate75/go-services/workflows/Go/badge.svg?branch=master"></img>
 </p>


# roman_numbers

Roman number conversion application

## Command line arguments

This command takes 1 mandatory and one optional arguments :
* in (int) mandatory input number for the conversion
* max (int) maximum number that can be converted (by default 3000)

Sample command lines:
```
> roman_numbers -in 500
Roman number for 500 is D

> roman_numbers -in 5000 -max 5000
Using upper limit 5000 instead of 3000
Roman number for 5000 is MMMMM

```


### Get the library

Library and command will be available running:

```
go get -u github.com/hellgate75-verizon/roman_numbers
```


## License

The library is licensed with [LGPL v. 3.0](/LICENSE) clauses, with prior authorization of author before any production or commercial use. Use of this library or any extension is prohibited due to high risk of damages due to improper use. No warranty is provided for improper or unauthorized use of this library or any implementation.

Any request can be prompted to the author [Fabrizio Torelli](https://www.linkedin.com/in/fabriziotorelli) at the following email address:

[hellgate75@gmail.com](mailto:hellgate75@gmail.com)
 

