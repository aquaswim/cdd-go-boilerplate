# Some Service

"Some" Service Description

## Prerequisite

- Go 1.23+
- Gnu Make

## How to run

- clone this repo
- download dependency
    ```shell
    $ make configure
    ```
- Fill the `.env` with your correct value
- Generate api server from open-api-v3 contract
    ```shell
    make generate
    ```
- Run service
    ```shell
    $ make run-local
    ```

## Database Migrations

we use [dbmate 🚀](https://github.com/amacneil/dbmate) to do database migration, setup the dbmate according their doc.
And all the migrations file is in `./db/migrations` directory.

## Generating Mock Files

We use [uber/mock](https://github.com/uber-go/mock) for generating mock files.
all the mock is in `mocks` folder, please add stuff that need to be mocked here,
and add `//go:generate` directive so it can be regenerated when running:

```shell
$ make test
```

# Todo for this boilerplate

- [ ] Update go.mod
- [ ] update package name in: generate-server.config.yaml
- [ ] fix all import in *.go files
- [ ] Update Readme.md