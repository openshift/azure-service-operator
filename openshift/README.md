Azure Service Operator
======================

This repository contains the Azure Service Operator (ASO), which is a
dependency for Cluster API Provider for Azure (CAPZ).

Openshift uses the `v2` directory found in this repository.

The main artifacts to come from this repository will be the ASO binary and
associated container image. CustomResourceDefinitions _are explicitly omitted_
from this repository, because Openshift does not need the entire Azure API surface.

CRDs are included in the CAPZ manifests, which already filters the installed CRDs
down to only those relevant to cluster lifecycling.

If you run into issues building the container image, double check the
`openshift/Dockefile.dockerignore` file; it may be excluding something.
