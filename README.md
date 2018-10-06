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

A `resource` is a gateway to the outside world from a Slang perspective. Examples are filesystem, networking, databases, etc.

*Example*

```YAML
description: Filesystem
operations:
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
```

### Operator

An `operator` is an implementation of an `operation`.
