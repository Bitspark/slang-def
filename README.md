# Slang Definition

This project is an experimental Go package for an alternative Slang definition language.

It serves two purposes:
* Provide a reference implementation for the Slang definition language
* Provide the bottom layer for the main Slang project

The current goal is to extend it until it can be used as bottom layer for the main Slang project.

## Slang Definition Language

There are 4 entities:

* Type
* Operation
* Resource
* Operator

### Type

A `type` defines a structure for data in Slang.

*Example*

```YAML
description: Name of a person
type: map
map:
  firstName:
    type: string
  lastName:
    type: string
```

### Operation

An `operation` can be thought of as an interface with a description of the desired semantics.

*Example*

```YAML
description: Convert to uppercase
in:
  type: string
out:
  type: string
```

### Resource

A `resource` is a gateway to the outside world from a Slang perspective.
It provides services that can be accessed in a Slang operator and events which are the entrance points for any data
getting into Slang.

Resources can be used to model different kinds of system resources and devices.

Examples are:
* Filesystem (see example)
* Databases
* IoT Devices
* Networking
* APIs

Services and events are ordinary operations with two important exceptions:
* Only events of resources can have a blank out-port
* Only services of resources can have a blank in-port

*Example*

```YAML
description: Filesystem
services:
  readFile:
    description: Read a file
    in:
      description: Filename
      type: string
    out:
      type: map
      map:
        found:
          type: boolean
        content:
          type: string
  writeFile:
    description: Write a file
    in:
      type: map
      map:
        filename:
          type: string
        content:
          type: string
    out:
      description: Success
      type: boolean
events:
  fileWritten:
    description: Fired when a file has been written.
    in:
      description: Filename
      type: string
    out:
      type: trigger
```

### Operator

An `operator` is an implementation of an `operation`.

It can model a system or a subsystem from the most abstract perspective to the most detailed perspective.
All such systems are built using either a top-down or a bottom-up approach or a mixture thereof.

*Example*

```YAML
operation: {}  # We have no in or out ports, because this is a complete system
description: A little IoT world
resources:
  iotSensor1:
    resource:
      events:
        measuredTemperature:
          in:
            type: number
    embedding:
      measuredTemperature: measuredTemperature
  iotActor1:
    resource:
      services:
        adjustHeating:
          out:
            type: number
operators:
  measuredTemperature:
    type: slang  # inline
    operation:
      in:
        type: number
    slang:  # inline
      operation:
        in:
          type: number
      resources:
        iotActor1:
          resource:
            services:
              adjustHeating:
                out:
                  type: number
      operators:
        adjustHeating:
          type: resource
          resource:
            resource: iotActor1
            service: adjustHeating
          handles:
            - adjustHeating
        reactOnTemperature:
          type: reference
          reference: operator.defined.SomewhereElse
          handles:
            - reactOnTemperature
      connections:
        conn1:
          from:
            handle: main  # surrounding operator (measuredTemperature)
            port: ''  # complete in-port
          to:
            handle: reactOnTemperature
            port: ''  # complete in-port
        conn2:
          from:
            handle: reactOnTemperature
            port: ''
          to:
            handle adjustHeating
            port: ''
    specification:
      embedding:
        resourceMap:
          iotActor1: iotActor1
```

In this example, two IoT devices are connected via Slang.