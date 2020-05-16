# eszero - exit status tester

eszero is a small tool for testing an exit status, also known as return code or exit code.

Examples (on Linux bash):

- `eszero` returns zero (and print exit status)

    ```bash
    $ ./eszero
    0
    ```
- `eszero` returns a number of the first argument

    ```bash
    $ ./eszero 1 ; echo $?
    1
    1
    $ ./eszero -1 ; echo $?
    -1
    255
    ```
- `eszero_n` returns a number specified by `n`

    ```bash
    $ cp eszero eszero_1 ; ./eszero_1 ; echo $?
    1
    1
    $ cp eszero eszero_-1 ; ./eszero_-1 ; echo $?
    -1
    255
    ```
- `n` returns a number specified by `n`

    ```bash
    $ cp eszero 1 ; ./1 ; echo $?
    1
    1
    $ cp eszero -- -1 ; ./-1 ; echo $?
    -1
    255
    ```
- `n` is able to be a number of binary, octal, hexadecimal

    ```bash
    $ cp eszero 0b10 ; ./0b10 ; echo $?
    2
    2
    $ cp eszero 010 ; ./010 ; echo $?
    8
    8
    $ cp eszero 0x10 ; ./0x10 ; echo $?
    16
    16
    ```
