[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sangwe11/sensu-registration-mutator)
![Go Test](https://github.com/sangwe11/sensu-registration-mutator/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/sangwe11/sensu-registration-mutator/workflows/goreleaser/badge.svg)

# Agent Registration Mutator

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Mutator definition](#mutator-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The Agent Registration Mutator is a [Sensu Mutator][6] that ...

## Files

## Usage examples

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add sangwe11/sensu-registration-mutator
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/sangwe11/sensu-registration-mutator].

### Mutator definition

```yml
---
type: Mutator
api_version: core/v2
metadata:
  name: sensu-registration-mutator
  namespace: default
spec:
  command: sensu-registration-mutator
  runtime_assets:
  - sangwe11/sensu-registration-mutator
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-registration-mutator repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/mutator-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/mutator-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/mutators/
[7]: https://github.com/sensu-community/mutator-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
