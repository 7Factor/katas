# GO Katas

## Assumptions

You have already installed go and understand how to create a basic project. If not please
head over [here]() and read the instructions. Note that you do not need a global `$GOPATH`
to use our projects because we think that's dumb. You can simply set your `$GOPATH` to
the top level directory (the one containing `src`) and everything will work just fine.

## Installation

Run `./install.sh` in any of the kata directories. We use glide for dependency
management. Find it here:

http://glide.sh/

## Running Tests

Run `./test.sh`.  This will run ginkgo in watch mode. Feel free
to use the native go test suite if you prefer.