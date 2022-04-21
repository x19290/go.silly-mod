# Golang: a silly package/module example

## Naming policy

- Unique package names are really globally unique;
  these are not like "example.com/hello."
- Names of unique/short Packages, **/*.go, functions, constants
  are not related to each other except main() of main.
- All non-quallified names are identical each other except main and main().

## Unique/Short package names of dirs

- ./
  - github.com/x19290/go.silly-mod
  - silly

- ./1testdata/
  - github.com/x19290/go.silly-mod/1testdata
  - naked

- ./2demo
  - NONE: not an importable package
  - main

## **/*.go and defined functions/constants

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

- ./2demo
  - 20print.go
    - main()

## Reference chains

- Public()
  - .private()
    - Version
      - .version

- TestWhite()
  - Version
  - .private()

- TestNaked()
  - .version
  - .Version

- main()
  - Version
  - Public()

## test.sh

POSIX specific.
Expected to be run on the dir that holds itself.

A MS-DOG user or even a POSIX user
can easily follow the same steps in it manually.

## github.com/x19290/go.silly-mod-demo

(https://github.com/x19290/go.silly-mod-demo)

A repo that has github.com/x19290/go.silly-mod as a submodule.

From the former, you can access the latter locally or remotely.

## Words from a novice to novices

- [Naming/Versioning confusion](novice2novice/naming-confusion.md)
- [About demo main()s](novice2novice/demo-mains.md)
