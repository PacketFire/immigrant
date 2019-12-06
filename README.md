# Immigrant
![immigrant master build status](https://travis-ci.org/PacketFire/immigrant.svg?branch=master)

## Setup
To setup locally, simply clone the repository and run `make`.
After running make, for usage with docker execute, `docker run --rm packetfire/immigrant --command`.

### Dependencies
- go 1.13+
- make

## Testing
### Running the Test Suite
The unit test suite leverages the standard libraries testing pkg and requires no additional dependencies outside of the ones listed above.

The test suite can be executed accross all subpackages by running `make test`.

### Linting
Code quality and linting is handled by the golint package and the linting suite can be executed accross all subpackages by running `make lint`.

### Questions, Issues and Contributions
Any questions in regards to this software, feel free to send an email to, [admin@packetfire.org](mailto:admin@packetfire.org).

If any issues occur when accessing/using this application, please file a [bug report issue](https://github.com/packetfire/immigrant/issues/new).

To contribute, fork this repository, apply changes to a local branch and [create a pull request](https://github.com/packetfire/immigrant/compare). All contributions are welcome and will be reviewed accordingly.
