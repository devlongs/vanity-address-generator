# Ethereum Vanity Address Generator
The Ethereum Vanity Address Generator is a Go-based tool that allows you to generate Ethereum addresses with a desired prefix. It provides a simple and efficient way to create personalized and memorable Ethereum addresses.

### Features
- Generate Ethereum addresses with a specified prefix
- Customizable prefix length and characters
- Fast and efficient generation algorithm
- Easy-to-use command-line interface
- Generates corresponding private keys for the vanity addresses

### Installation
1. Clone the repository:

```bash
git clone https://github.com/devlongs/ethereum-vanity-address-generator.git
```
2. Navigate to the project directory:
```bash
cd ethereum-vanity-address-generator
```
3. Build the project:

```bash
go build
```
4. Usage
To generate a vanity address, run the following command:

```bash
./ethereum-vanity-address-generator <prefix>
```
Replace <prefix> with the desired prefix for your vanity address. The prefix should only contain valid hexadecimal characters (0-9, a-f).
