# teleport_test

[![](https://img.shields.io/badge/Go-1.14.1-blue.svg)](https://golang.org/pkg/)

Contents of this repo include an optional GoLang code challenge submitted after receiving an initial offer for a
Technical Writing and Developer Advocate (Developer Relations) job.

I felt it was important to help substantiate my ability to work in the stack.

## Job Role

So, a tentative offer was made at Developer Relations IV which I turned down (since I felt it was outside my range of
experience). I was also surprised at the generosity of the offer (it's competitive and appears to be in line with both
the average and median at multiple firms). They insisted on Developer Relations II!

Per our previous conversation, the CTO felt that my previous work, articles, and certificates was sufficient to advance
forward for the Developer Relations role.

Despite that, I'd also like to complete at least one of the product engineering challenges to help substantiate my
application and ensure I'm in line with the team.

------

*Backend Engineer* - Level 1:

To be implemented in several stages per
the [requirements doc](https://github.com/gravitational/careers/blob/main/challenges/systems/worker.pdf).

### Level 1

#### Library

Worker library with methods to start/stop/query status and get an output of a running job.

#### API

Add HTTPS API to start/stop/get status of a running process. Use basic authentication.

#### Client

Client command should be able to connect to worker service and schedule several jobs.

Client should be able to query result of the job execution and fetch the logs.

## Running

Execute the following commands:

1. `go get github.com/gofrs/uuid`

For a valid self-signed SSL:

1. `openssl genrsa -out key.pem 2048`
1. `openssl req -new -sha256 -key key.pem -out csr.csr`
1. `openssl req -x509 -sha256 -days 365 -key key.pem -in csr.csr -out certificate.pem`
1. `openssl req -in csr.csr -text -noout | grep -i "Signature.*SHA256" && echo "All is well" || echo "This certificate will stop working in 2017! You must update OpenSSL to generate a widely-compatible certificate"`
1. `openssl x509 -outform der -in yourPemFilename.pem -out certfileOutName.crt`
1. `openssl rsa -in yourPemFilename.pem -out keyfileOutName.key`

## Resources

1. https://golang.org/pkg/
1. https://golangbyexample.com/singleton-design-pattern-go/
1. https://github.com/Thoughtscript/teleport_test
1. https://msol.io/blog/tech/create-a-self-signed-ssl-certificate-with-openssl/
1. https://stackoverflow.com/questions/13732826/convert-pem-to-crt-and-key
