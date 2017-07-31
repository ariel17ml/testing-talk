# Test methodologies

## Table driven tests

* What is it?

TODO

* Pro:

  - Easy to add new test cases.
  - Exhaustive scenarios made simple
  - Recommended to use it even in single cases.
  - Recommended to use name in test cases.

* Cons:

  - TODO

## Test fixtures

* What is it?

TODO

* Pro:

  - Useful for configuration, data models, binary data.

* Cons:

  - `go test` sets `pwd` as package directory, so relative path to fixtures is
    recommended.

## Golden files

* What is it?

TODO

* Pro:

  - 

* Cons:

## test flags

## Global state

* What is it?
* Pro:

  - Better to provide default values instead of constants for configurations
    that never changes.

* Cons:

  - Difficult to change for testing.

## Test helpers

* Cons:

  - Elegant.
  - Close is a closure that can also fail.

* Pros:

  - Never return errors; fail the test directly.
  - Easier since error checking is gone.
  - More clear on what they are testing, instead of boilerplate.
  - Return a function if cleaning is needed.

# Writing testeable code

TODO

# Recomendations

## On packages/functions

- Test only exported API, unless the relying implementation is too complex.
- Use the blackblox approach to avoid unit test what is not needed, but use it
  wisely.

## Networking

- Do not mock `net.Conn`, create a real network connection.
- Easy to test any protocol.

## Configurability

- Unconfigurable produces "untesteable" code.
- Overparametrize structs to completelly manipulate the object.

## Subprocessing

- Mock output or execute (a good thing, actually).
- Only run if binary is present.
- Verify that side effects would not affect other tests.
- `exec.Cmd` is an struct an can be replaced by a mock.

## Interfaces

- Potentila mocking point.

## Testing as a public API

## Parallelization

- Don't run test in parallel; they can be affected by others.
- 
