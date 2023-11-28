# freeminer

freeminer is a CPU-based miner for freed

## Requirements

Go 1.19 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install freed including all dependencies:

```bash
$ git clone https://github.com/hungyu99/freed/
$ cd freed/cmd/freeminer
$ go install .
```

- Kapaminer should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.
  
## Usage

The full freeminer configuration options can be seen with:

```bash
$ freeminer --help
```

But the minimum configuration needed to run it is:
```bash
$ freeminer --miningaddr=<YOUR_MINING_ADDRESS>
```