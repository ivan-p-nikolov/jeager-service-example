# cm-go-service

[![Coverage Status](https://coveralls.io/repos/github/Financial-Times/cm-go-service/badge.svg)](https://coveralls.io/github/Financial-Times/cm-go-service)
[![CircleCI](https://circleci.com/gh/Financial-Times/cm-go-service.svg)](https://circleci.com/gh/Financial-Times/cm-go-service)

Content&amp;Metadata platform template repository for Go microservices.

- Use this repo as template repository.
- Look for all occurrences of the string `cm-go-service` and rename them appropriately.
- Look for all `TODO` statements and fix them appropriately.
- Decide in which clusters should the service be deployed (e.g. PAC/UPP, Delivery/Publish) and leave only the needed configurations in the `helm/cm-go-service/app-configs` directory. There are some example configurations provided in the `app-configs` directory to be used as a guideline.
- Rename `helm/cm-go-service` ***and*** the corresponding app-config files e.g. `helm/cm-go-service/app-configs/cm-go-service_delivery.yaml`.
- Add the team that supports the service in the  `.github/CODEOWNERS` file, e.g. (leave only the relevant team):

    ```text
    # This repo is supported by:
    * @Financial-Times/content-team @Financial-Times/metadata-team @Financial-Times/platform-health
    ```

## Installation

Download the source code, dependencies and test dependencies:

```shell
git clone https://github.com/Financial-Times/cm-go-service.git
cd cm-go-service
go build
```

## Service endpoints

For a full description of the API endpoints for the service, please check the [Open API specification](./api/api.yml).

## Admin endpoints

The admin endpoints are:

```text
/__gtg
/__health
/__build-info
```
