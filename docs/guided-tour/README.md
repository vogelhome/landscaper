---
title: Prerequisites
sidebar_position: 1
---

# Guided Tour

## Prerequisites and Basic Definitions

- For all examples, you need a [running Landscaper instance](../installation/install-landscaper-controller.md).

- A convenient tool we will often use in the following examples is the [Landscaper
  CLI](https://github.com/gardener/landscapercli). 

- For building reusable components you require the [OCM CLI](https://ocm.software/docs/guides/getting-started-with-ocm/#prerequisites).

- During the following exercises, you might need to change files, provided with the examples. For this, you should
  simply clone [this repository](https://github.com/gardener/landscaper) and do the required changes on your local files. You could also fork the repo and work on your fork.

- In all examples, 3 Kubernetes clusters are involved:

  - the **Landscaper Host Cluster**, on which the Landscaper runs
  - the **target cluster**, on which the deployments will be done
  - the **Landscaper Resource Cluster**, on which the various custom resources are stored. These custom resources are
    watched by the Landscaper, and define which deployments should happen on which target cluster.

  It is possible that some or all of these clusters coincide, e.g. in the most simplistic approach, you have only one
  cluster. Such a "one-cluster-setup" is the easiest way to start working with the Landscaper.

## How to follow the Tour

The Guided Tour consists of several chapters, some of which build on each other. In order to get the most out of it, 
you should be following the Guided Tour in this sequence. If you have problems with the presentation of the Guided Tour,
jump to the [original github repository](https://github.com/gardener/landscaper/tree/master/docs/guided-tour).

### A Hello World Example

[1. Hello World Example](./hello-world)

### Basics

[2. Upgrading the Hello World Example](./basics/upgrade)

[3. Manifest Deployer Example](./basics/manifest-deployer)

[4. Multiple Deployments in One Installation](./basics/multiple-deployitems)

### Recovering from Errors

[5. Handling an Immediate Error](./error-handling/immediate-error)

[6. Handling a Timeout Error](./error-handling/timeout-error)

[7. Handling a Delete Error](./error-handling/delete-error)

You can find a list of error messages and corresponding solutions [here](./error-handling/problem_analysis.md).

### Blueprints and Components

[8. An Installation with an Externally Stored Blueprint](./blueprints/external-blueprint)

[9. Helm Chart Resources in the Component Descriptor](./blueprints/helm-chart-resource)

[10. Echo Server Example](./blueprints/echo-server)

### Imports and Exports

[11. Import Parameters](./import-export/import-parameters)

[12. Import Data Mappings](./import-export/import-data-mappings)

[13. Export Parameters](./import-export/export-parameters)

### Templating

[14. Templating: Accessing Component Descriptors ](./templating/components)

## Target Maps

[15. Target Maps: Multiple Deploy Items](./target-maps/01-multiple-deploy-items)

[16. Target Maps: Target Map References](./target-maps/02-targetmap-ref)

[17. Target Maps: Multiple Subinstallations](./target-maps/03-multiple-subinst)

[18. Target Maps: Target Map on Subinstallation Level](./target-maps/04-forward-map)

[19. Target Maps: Other Target Map Examples](./target-maps/05-other-examples)

## Optimization

[20. Optimization Hints ](../usage/Optimization.md)
