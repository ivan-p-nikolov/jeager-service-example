# This is the name of the system

//TODO: Fill the runbook with relevant information
Here is a long description of a system that contains the minimum detail that should be provided via a RUNBOOK.md file. The format includes the provision of text value, enum, single value and multiple value fields. Please try and avoid paragraphs as that longer content may belong in the other text fields which are contained below.

## Code

system-code-in-biz-ops

## Primary URL

https://www.sample.ft.com

## Service Tier

Bronze

## Lifecycle Stage

Preproduction

## Host Platform

AWS

## First Line Troubleshooting

Here are some details to explain how the **first line** support team can resolve issues with the system when it is running in production.
Feel free to use paragraphs and other formatting techniques.

## Second Line Troubleshooting

Here are some details to explain how the **second line** support team can resolve issues with the system when it is running in production.
Feel free to use paragraphs and other formatting techniques.

## More Information

This section should be used to capture any additional information that would otherwise have over extended the description or polluted the troubleshooting.

## Monitoring

Health Checks:

* [EU](https://{{CLUSTER_NAME}}-eu.upp.ft.com/__health/__pods-health?service-name={{SERVICE}})
* [US](https://{{CLUSTER_NAME}}-us.upp.ft.com/__health/__pods-health?service-name={{SERVICE}})

Splunk Alerts:

- [Alert Name](link.to.alert)

## Contains Personal Data

True

## Contains Sensitive Data

False

## Architecture

Here are some details about the architecture of the system.

## Failover Architecture Type

ActiveActive

## Failover Process Type

FullyAutomated

## Failback Process Type

PartiallyAutomated

## failoverDetails

Please include some text, with paragraphs, to explain how the system is failed over and how to fail back.

Failover:

- Step 1
- Step 2
- Step 3

Failback:

- Step 1
- Step 2
- Step 3.

## Data Recovery Process Type

Manual

## Data Recovery Details

Please include some text, with paragraphs, to explain how the system's data is restored after failure.

Restore:

- Step 1
- Step 2
- Step 3

## Release Process Type

FullyAutomated

## Rollback Process Type

Manual

## releaseDetails

Please include some text, with paragraphs, to explain how the system is released and how to roll it back.

Release:

- Step 1
- Step 2
- Step 3

Rollback:

- Step 1
- Step 2
- Step 3.

## Key Management Process Type

PartiallyAutomated

## Key Management Details

Please include some text, with paragraphs, to explain how the system's keys are rotated.

Key Rotation:

- Step 1
- Step 2
- Step 3
