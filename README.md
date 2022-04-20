# Golang a silly package/module example

## Unique/Short packages of dirs

- ./
  - [github.com/x19290/go.silly-mod](https://github.com/x19290/go.silly-mod)
  - silly

- ./1testdata/
  - [github.com/x19290/go.silly-mod](https://github.com/x19290/go.silly-mod)/1testdata
  - naked

## Contents of dirs

- ./
  - 00public.go
    - .Public()
  - 01private.go
    - .private()
  - 02whitebox_test.go
    - .TestWhite()

- ./1testdata
  - 10public.go
    - .Version
  - 11private.go
    - .version
  - 12naked_test.go
    - .TestNaked()

## Reference chain

  - Public()
    - private()
      - Version
        - version

  - TestWhite()
    - naked.Version
    - .private()

  - TestNaked()
    - .version
    - .Version

## test.sh

POSIX specific.
Expected to be run on the dir that holds itself.

A MS-DOG user or even a POSIX user easily follow the same steps in it manually.

## [github.com/x19290/go.silly-mod-demo](https://github.com/x19290/go.silly-mod-demo)

A repo that has [github.com/x19290/go.silly-mod](https://github.com/x19290/go.silly-mod) as a submodule.

From the former, you can access the latter locally or remotely.
