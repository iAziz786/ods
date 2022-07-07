# oms -- Open MongoDB Scanner
**oms** Checks the open MongoDB databases without any kind security layer to connect to them. Secure your DB!

## _This repo exist for educational purpose only. In no way it endorse hacking._ 

## Usage Example

**Scan all open ports through the internet**

```
$ oms --cidr=0.0.0.0/0
```

**Use a custom port for MongoDB**

```
$ oms --mongo-port=27018
```

**Use a different MongoDB dialect**

```
$ oms --mongo-dialect=mongodb+srv
```

**Show this help page**

```
$ oms --help
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
