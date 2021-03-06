<p align="center">

  <img src="docs/assets/images/icon.png" width="80" height="80" alt="flake logo">
  <h1 align="center">Flake</h1>

  <p align="center">
    Generate k-sortable 64-bit unique flake ids.
    <br>
    <br>
    <a href="https://github.com/ybt195/flake/issues">Issues</a>
    ·
    <a href="https://godoc.org/github.com/ybt195/flake/pkg/flake">Documentation</a>
  </p>

  <p align="center">
    <a href="https://godoc.org/github.com/ybt195/flake/pkg/flake">
      <img alt="GoDoc Reference" src="https://godoc.org/github.com/ybt195/flake/pkg/flake?status.svg">
    </a>
    <a href="https://travis-ci.org/ybt195/flake">
      <img alt="Build Status" src="https://travis-ci.org/ybt195/flake.svg?branch=master">
    </a>
    <a href="https://codecov.io/gh/ybt195/flake">
      <img src="https://codecov.io/gh/ybt195/flake/branch/master/graph/badge.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/ybt195/flake">
      <img alt="Go Report Status" src="https://goreportcard.com/badge/github.com/ybt195/flake">
    </a>
    <a href="https://opensource.org/licenses/Apache-2.0">
      <img alt="Apache License, Version 2.0" src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
    </a>
  </p>
</p>

## Table of contents

- [Overview](#overview)
- [Contributing](#contributing)
- [License](#license)

## Overview

Flake ids are represented in 64-bit unsigned integers. Ids can be broken down into three components:

- Bucket: A 10-bit bucket id for sharding ids over 1024 different buckets.
- Timestamp: A 42-bit millisecond encoded timestamp of when the id was generated.
- Sequence: A 12-bit  incrementing sequence number that is reset at each timestamp increment.

With this formation, flake can generate 4096 unique ids per millisecond per bucket, or in other words, 4 **billion** unique ids per second. For reference, YouTube gets about 75 **thousand** unique views per second.

## Contributing

Contributions are very much welcomed. Please read the [Contribution Guide](CONTRIBUTING.md) for how to get started.

Note that all contributors are expected to follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## License

Flake is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0). For more information, see the [License](LICENSE) file. 
