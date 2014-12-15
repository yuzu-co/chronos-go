Chronos Client
==============

`chronos-go` is an API client for [Chronos](https://github.com/airbnb/chronos) written in Go.

    go get github.com/gocardless/chronos-go

Example
-------

Create a task that executed inside a Docker container:

    package main

    import (
        "log"

        "github.com/gocardless/chronos-go"
    )

    func main() {
        j := chronos.NewJob()

        j.Name = "DockerTest"
        j.Owner = "hello@example.com"
        j.Command = "/usr/bin/uptime"
        j.Container = chronos.NewContainer("ubuntu:trusty")
        j.Schedule = "R/2014-12-15T16:00:00/PT24H"
        j.Epsilon = "PT24H"

        err := j.Create()
        if err != nil {
            log.Fatalf("unable to create job: %s\n", err)
        }

        log.Println(j)
    }
