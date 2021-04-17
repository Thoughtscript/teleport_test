# teleport_test

[![](https://img.shields.io/badge/Go-1.14.1-blue.svg)](https://golang.org/pkg/)

Contents of this repo include an optional GoLang code challenge I completed for fun!

> The [Design Document](Design_Document.md) contains an overview of the application, challenge, and discussion about design decisions.

## Introduction and Objectives

I've accepted a Developer Relations II role (Technical Writer) and wanted to complete at least one of the product engineering challenges for fun!

This repo is the outcome of that effort.

### Backend Engineer - Level One

To be implemented in several stages per
the [requirements doc](https://github.com/gravitational/careers/blob/main/challenges/systems/worker.pdf).

1. **Level One**

    1. **Library** 
        1. Worker library with methods to start/stop/query status and get an output of a running job.
    1. **API** 
        1. Add HTTPS API to start/stop/get status of a running process. Use basic authentication.
    1. **Client**
        1. Client command should be able to connect to worker service and schedule several jobs. 
        1. Client should be able to query result of the job execution and fetch the logs.

## Setup, Use, and Installation

Execute the following commands to grab all the dependencies:

1. `$ go get github.com/gofrs/uuid`

For a valid self-signed SSL:

1. `$ openssl genrsa -out key.pem 2048`
1. `$ openssl req -new -sha256 -key key.pem -out csr.csr`
1. `$ openssl req -x509 -sha256 -days 365 -key key.pem -in csr.csr -out certificate.pem`
   
Navigate to [src](./src):

1. `$ go run httpsServer.go`
1. You will likely run into issues trying to access https://localhost/public/ in your browser. Safari is recommended.

### REST API

> The default authentication settings are: `User`: `test` and `Password`: `test`.

> API HTTP headers are case-insensitive.

> The API uses a self-signed TLS certificate generated through OpenSSL.

> Successful POST requests will return status code `201` and successful GET requests will return status code `200`.

1. GET - https://localhost/public/

   Brings up a simple HTML client.

   ![landingpage](./img/landingpage.png)

1. POST - https://localhost/api/create

   With headers:

    1. `Cmd` - `string` - bash command - this will be converted to `ls` so anything you pass in here is fine to send.
    1. `Scheduled` - `string` - valid go `time.RFC3339` [parsable string](https://golang.org/pkg/time/#example_Parse): `"2006-01-02T15:04:05Z"`
    1. `Content-Type` - `application/json`
    1. `User` - `string`
    1. `Password` - `string`

    ![create](./img/postcreate.png)

1. GET - https://localhost/api/pool

   With headers:

    1. `Content-Type` - `application/json`
    1. `User` - `string`
    1. `Password` - `string`

   Response:
    ```
    {
        "10a8952b-d730-447c-b1d8-b15614944246": "queued",
        "15fa758d-9c8e-4eef-877a-e332675e55fe": "completed",
        "2866b264-513f-4a68-88d2-d8ee4f294f7f": "completed",
        "4e7b801c-4fd6-4b4d-87dc-c1ce1481d4af": "completed",
        "59c904ca-bbbc-4d1b-8831-ce86725d440e": "completed",
        "86480c6d-2134-4aca-82e4-7854e3041ab1": "completed",
        "b3e8cc19-4a3c-4c83-8b11-26698d1db2a8": "executing",
        "cbaf1e15-ec3a-4182-b122-b7e235b103c0": "completed",
        "dff12a72-932b-44eb-80d3-acf4c29b2aeb": "completed"
    }
    ```

    ![pool](./img/getpool.png)

1. GET - https://localhost/api/jobs

   With headers:

    1. `Content-Type` - `application/json`
    1. `Uuid` - `string` - uuid of **Worker**
    1. `User` - `string`
    1. `Password` - `string`

   Response:
    ```
    {
        "Uuid": "2f99ae7c-992c-42bd-9df3-293101086d08",
        "Time": "2021-01-02T15:04:05Z",
        "Status": "queued",
        "Command": "ls",
        "Output": ""
    }
    ```

    ![jobs](./img/getjobs.png)

1. POST - https://localhost/api/stop

   With headers:

    1. `Content-Type` - `application/json`
    1. `Uuid` - `string` - uuid of **Worker**
    1. `User` - `string`
    1. `Password` - `string`

   Response:
   
    ```
    "stopped"
    ```

   This API will return `completed`, `failed`, or `stopped`.

   ![stop](./img/poststop.png)

1. GET - https://localhost/api/status

   With headers:

    1. `Content-Type` - `application/json`
    1. `Uuid` - `string` - uuid of **Worker**
    1. `User` - `string`
    1. `Password` - `string`

   Response:

    ```
    "stopped"
    ```

   This API will return `completed`, `failed`, or `stopped`.

   ![status](./img/getstatus.png)

### Console Examples

What you should see server-side:

```bash
=================== POLLING EVERY 5s ===================
2021/04/15 16:25:46 http: TLS handshake error from [::1]:51620: EOF
Worker added: 970801a3-7623-46dc-88c2-dac62e99de5d
Worker added: 32399647-2f91-4a8c-b479-712c30b4049a
Worker added: d054e126-d119-4a82-8733-c1d5ca9ebd71
Worker added: 30a5bb6e-e2a8-436f-b275-3c313a4f0a48
=================== POLLING EVERY 5s ===================
32399647-2f91-4a8c-b479-712c30b4049a queued ls assets
cert.pem
handlers
httpsServer.go
jobs
key.pem
models
tests

d054e126-d119-4a82-8733-c1d5ca9ebd71 queued ls assets
cert.pem
handlers
httpsServer.go
jobs
key.pem
models
tests

30a5bb6e-e2a8-436f-b275-3c313a4f0a48 queued ls assets
cert.pem
handlers
httpsServer.go
jobs
key.pem
models
tests

Worker removed: 7fdc8724-95af-4688-8ece-2618e9bb6fa4
Worker stopped: 7fdc8724-95af-4688-8ece-2618e9bb6fa4
970801a3-7623-46dc-88c2-dac62e99de5d queued ls assets
cert.pem
handlers
httpsServer.go
jobs
key.pem
models
tests

=================== POLLING EVERY 5s ===================
```

## Additional Resources

1. https://golang.org/pkg/
1. https://golangbyexample.com/singleton-design-pattern-go/
1. https://github.com/Thoughtscript/teleport_test
1. https://msol.io/blog/tech/create-a-self-signed-ssl-certificate-with-openssl/
