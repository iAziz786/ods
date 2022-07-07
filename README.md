# ods -- Open Datebase Scanner
**ods** checks for the open database which doesn't have any kind security layer. Secure your DB!

Currently only MongoDB is supported but other database will follow if requested.

## _This repo exist for educational purpose only. In no way it endorse hacking._ 

## Usage Example

**Scan all open ports through IPv4 addresses**

```
$ ods --cidr=0.0.0.0/0
```

**Use a custom port for MongoDB**

```
$ ods --db-port=27018
```

**Specify the dialect to use. Only `mongodb` and `mongodb+srv` supported.**

```
$ ods --dialect=mongodb+srv
```

**Show this help page**

```
$ ods --help
```

## License

MIT License

Copyright (c) 2022 Mohammad Aziz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
