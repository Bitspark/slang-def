# Slang Definition

This project is an experimental Go package for an alternative Slang definition language.

It serves two purposes:
* Provide a reference implementation for the Slang definition language
* Provide the bottom layer for the main Slang project

The current goal is to extend it until it can be used as bottom layer for the main Slang project.

## Slang Definition Language

There are 3 entities:

* Type
* Operation
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
It consists of two maps, one for input and one for output.

*Example*

```YAML
description: Convert to uppercase
in:
  mixedCase:
    type: string
out:
  upperCase:
    type: string
```
### Operator

An `operator` is an implementation of an `operation`.

It can model a system or a subsystem from the most abstract perspective to the most detailed perspective.
All such systems are built using either a top-down or a bottom-up approach or a mixture thereof.

*Example*

```YAML
description: A little IoT world
instances:
  iotSensor:
    elementary: iot.sensor
    embedding:
      handler:
        handle:
          instance: handler
          service: handle
  handler:
    services:
      handle:
        operation:
          reference: iot.sensor.handle
    instances:
      iotActor:
        elementary: iot.actor
    implementation:
      measuredTemperature:
        handles:
          handle:
            instance: iotActor
            service: handle
        connections:
          1:
            source:
              handle: hull
              port: value
            destination:
              handle: handle
              port: value
```

In this example, two IoT devices are connected via Slang.