# hackedemailsapi
[![GitHub release](http://img.shields.io/github/release/jakewarren/hackedemailsapi.svg?style=flat-square)](https://github.com/jakewarren/hackedemailsapi/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/hackedemailsapi/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/hackedemailsapi)](https://goreportcard.com/report/github.com/jakewarren/hackedemailsapi)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

Unofficial API client for hacked-emails.com. Queries the API for any known breaches involving an email address.

# Example

```
$ hackedemailsapi redacted@email.com
2 breaches returned for redacted@email.com

Last.fm
        source_url: #
        date_released:2016-09-23T00:00:00+00:00
        date_leaked:2012-08-31T00:00:00+00:00
        source_network: darknet
        email_count: 37192134
        verified: true

Linkedin
        source_url: #
        date_released:2016-06-18T00:00:00+00:00
        date_leaked:2011-12-31T00:00:00+00:00
        source_network: darknet
        email_count: 159752107
        verified: true


```

# Usage

```
$ hackedemailsapi -h
usage: hackedemails [<flags>] <email>

Un-official API client for hacked-emails.com.

Optional flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
  -d, --debug                    print debug info
  -f, --filter-date=FILTER-DATE  only print breaches released after specified date
  -s, --silent                   suppress response message, only display results
  -V, --version                  Show application version.

Args:
  <email>  the email address to lookup.
```

# Installation

```
go get github.com/jakewarren/hackedemailsapi
```
