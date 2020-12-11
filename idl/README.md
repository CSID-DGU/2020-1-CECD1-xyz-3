# idl

doogeun service interface definitions

## Contribution Guide

### Prerequisites

- [Docker](https://www.docker.com)
- [Task](https://taskfile.dev)
- [pre-commit](https://pre-commit.com)

### How to

1. Install `pre-commit` hooks for validate your commit
    ```sh
    pre-commit install
    ```
1. Write service definitions in `protos` directory
1. Commit written service definitions
    - Please **do not commit local-generated code files**. Code will be automatically generated after merge.
    - You can check outputs in local by `task gen:local`
1. Open pull request
