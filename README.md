## Portainer's Project

This project validates the setup config via string format and returns a message if the setup is valid or not via HTTP.
Inside of the project there are 2 folders. The **client** and the **library** (portainer's) folders.

## Run locally

The project has a Makefile that allows you to do the following actions:

1.  In **library** run the command `make go-init` to install all the dependencies
2.  In **client** run the command `make go-init` to install all the dependencies

    a. Then, run `make run` to start the server

    b. Once the server started, in a Postman or browser client, send a request like:
    `localhost:10000/validate/{setup}` where **{setup}** is the configuration that wants to be validated.

Example: `localhost:10000/validate/[{()}{[()()]}{[()]}]`

## Test

In library run the command `make test`

    go test -v -cover
    === RUN   TestValidateSetup
    === RUN   TestValidateSetup/Invalid_Length
    === RUN   TestValidateSetup/Invalid_Characters
    === RUN   TestValidateSetup/Invalid_Setup
    === RUN   TestValidateSetup/Valid_Setup_Example_1
    === RUN   TestValidateSetup/Valid_Setup_Example_2
    === RUN   TestValidateSetup/Valid_Setup_Example_3
    === RUN   TestValidateSetup/Valid_Complex_Setup
    --- PASS: TestValidateSetup (0.00s)
        --- PASS: TestValidateSetup/Invalid_Length (0.00s)
        --- PASS: TestValidateSetup/Invalid_Characters (0.00s)
        --- PASS: TestValidateSetup/Invalid_Setup (0.00s)
        --- PASS: TestValidateSetup/Valid_Setup_Example_1 (0.00s)
        --- PASS: TestValidateSetup/Valid_Setup_Example_2 (0.00s)
        --- PASS: TestValidateSetup/Valid_Setup_Example_3 (0.00s)
        --- PASS: TestValidateSetup/Valid_Complex_Setup (0.00s)
    PASS
    coverage: 100.0% of statements
    ok      spark/challenge/portainer       0.444s

