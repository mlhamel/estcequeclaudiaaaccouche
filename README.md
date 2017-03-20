# Est-ce que Claudia a accouché?

[http://est-ce-que-claudia-a-accouche.info](http://est-ce-que-claudia-a-accouche.info)

## Installation
```sh
git clone git@github.com:mlhamel/estcequeclaudiaaaccouche.git
cd estcequeclaudiaaaccouche
go get ./...
make
```

## Usage:
```sh
accouchement disable [--redis=<url>]
accouchement enable [--redis=<url>]
accouchement toggle [--redis=<url>]
accouchement notify [--redis=<url>] [--sid=<sid>] [--token=<token>] [--from=<from>] [--to=<to>]
accouchement serve [--port=<port>] [--redis=<url>] [--source=<source>] [--sid=<sid>] [--token=<token>]
accouchement status [--redis=<url>]
accouchement [--port=<port>] [--redis=<url>] [--sid=<sid>] [--token=<token>]
accouchement -h | --help
accouchement --version

--redis=<url>           Change Redis configuration to [default: redis://@192.168.64.42:6379].
--port=<port>           Port to serve [default: 4242].
--source=<source>       Authorized source of action [default: +15149999999].
--sid=<sid>             SID for twilio.
--token=<token>         Token for twilio.
--from=<from>           Source number for twilio.
--to=<to>               Destinatination number for twilio.
-h --help               Show this screen.
--version               Show version.
```

[![Build Status](https://travis-ci.org/mlhamel/estcequeclaudiaaaccouche.svg?branch=master)](https://travis-ci.org/mlhamel/estcequeclaudiaaaccouche)
