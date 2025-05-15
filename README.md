# Go Micro Demo [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/go-micro/demo?status.svg)](https://godoc.org/github.com/go-micro/demo)

<p align="center">
<img src="service/frontend/static/icons/Hipster_HeroLogoCyan.svg" width="300" alt="Online Boutique" />
</p>

**This application was forked from [microservices-demo](https://github.com/GoogleCloudPlatform/microservices-demo), used to demonstrate how to build micro servics with [go-micro](https://github.com/go-micro/go-micro).**

## Overview

**Online Boutique** is a cloud-native microservices demo application.
Online Boutique consists of a 11-tier microservices application. The application is a
web-based e-commerce app where users can browse items,
add them to the cart, and purchase them.

**Google uses this application to demonstrate use of technologies like
Kubernetes/GKE, Istio, Stackdriver, gRPC and OpenTelemetry**. This application
works on any Kubernetes cluster, as well as Google
Kubernetes Engine. It’s **easy to deploy with little to no configuration**.


If you’re using this demo, please **★Star** this repository to show your interest!

## Screenshots

| Home Page                                                                                                               | Checkout Screen                                                                                                          |
| ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ |
| [![Screenshot of store homepage](./docs/img/online-boutique-frontend-1.png)](./docs/img/online-boutique-frontend-1.png) | [![Screenshot of checkout screen](./docs/img/online-boutique-frontend-2.png)](./docs/img/online-boutique-frontend-2.png) |

## Other Deployment Options

- **Workload Identity**: [See these instructions.](docs/workload-identity.md)
- **Istio**: [See these instructions.](docs/service-mesh.md)
- **Anthos Service Mesh**: ASM requires Workload Identity to be enabled in your GKE cluster. [See the workload identity instructions](docs/workload-identity.md) to configure and deploy the app. Then, use the [service mesh guide](/docs/service-mesh.md).
- **non-GKE clusters (Minikube, Kind)**: see the [Development Guide](/docs/development-guide.md)
- **Memorystore**: [See these instructions](/docs/memorystore.md) to replace the in-cluster `redis` database with hosted Google Cloud Memorystore (redis).
- **Cymbal Shops Branding**: [See these instructions](/docs/cymbal-shops.md)
- **NetworkPolicies**: [See these instructions](/docs/network-policies/README.md)
- **Jaeger**: [See these instructions](/docs/jaeger.md)


## Architecture

**Online Boutique** is composed of 11 microservices written in different
languages that talk to each other over gRPC. See the [Development Principles](/docs/development-principles.md) doc for more information.

[![Architecture of
microservices](./docs/img/architecture-diagram.png)](./docs/img/architecture-diagram.png)

| Service                                              | Language      | Description                                                                                                                       |
| ---------------------------------------------------- | ------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| [frontend](./service/frontend)                           | Go            | Exposes an HTTP server to serve the website. Does not require signup/login and generates session IDs for all users automatically. |
| [cart](./service/cart)                     | Go            | Stores the items in the user's shopping cart in Redis and retrieves it.                                                           |
| [productcatalog](./service/productcatalog) | Go            | Provides the list of products from a JSON file and ability to search products and get individual products.                        |
| [currency](./service/currency)             | Go            | Converts one money amount to another currency. Uses real values fetched from European Central Bank. It's the highest QPS service. |
| [payment](./service/payment)               | Go            | Charges the given credit card info (mock) with the given amount and returns a transaction ID.                                     |
| [shipping](./service/shipping)             | Go            | Gives shipping cost estimates based on the shopping cart. Ships items to the given address (mock)                                 |
| [email](./service/email)                   | Go            | Sends users an order confirmation email (mock).                                                                                   |
| [checkout](./service/checkout)             | Go            | Retrieves user cart, prepares order and orchestrates the payment, shipping and the email notification.                            |
| [recommendation](./service/recommendation) | Go            | Recommends other products based on what's given in the cart.                                                                      |
| [ad](./service/ad)                         | Go            | Provides text ads based on given context words.                                                                                   |
| [loadgenerator](./service/loadgenerator)                 | Python+Locust | Continuously sends requests imitating realistic user shopping flows to the frontend.                                              |

## Features

- **[Kubernetes](https://kubernetes.io)/[GKE](https://cloud.google.com/kubernetes-engine/):**
  The app is designed to run on Kubernetes (both locally on "Docker for
  Desktop", as well as on the cloud with GKE).
- **[gRPC](https://grpc.io):** Microservices use a high volume of gRPC calls to
  communicate to each other.
- **[Istio](https://istio.io):** Application works on Istio service mesh.
- **[OpenTelemetry](https://opentelemetry.io/) Tracing:** Most services are
  instrumented using OpenTelemetry trace interceptors for gRPC/HTTP.
- **[Skaffold](https://skaffold.dev):** Application
  is deployed to Kubernetes with a single command using Skaffold.
- **Synthetic Load Generation:** The application demo comes with a background
  job that creates realistic usage patterns on the website using
  [Locust](https://locust.io/) load generator.

## OpenTelemetry

[![Jaeger Dependencies](./docs/img/jaeger-dependencies.png)](./docs/img/jaeger-dependencies.png)

## Local Development

If you would like to contribute features or fixes to this app, see the [Development Guide](/docs/development-guide.md) on how to build this demo locally.

## Demos featuring Online Boutique

- [From edge to mesh: Exposing service mesh applications through GKE Ingress](https://cloud.google.com/architecture/exposing-service-mesh-apps-through-gke-ingress)
- [Take the first step toward SRE with Cloud Operations Sandbox](https://cloud.google.com/blog/products/operations/on-the-road-to-sre-with-cloud-operations-sandbox)
- [Deploying the Online Boutique sample application on Anthos Service Mesh](https://cloud.google.com/service-mesh/docs/onlineboutique-install-kpt)
- [Anthos Service Mesh Workshop: Lab Guide](https://codelabs.developers.google.com/codelabs/anthos-service-mesh-workshop)
- [KubeCon EU 2019 - Reinventing Networking: A Deep Dive into Istio's Multicluster Gateways - Steve Dake, Independent](https://youtu.be/-t2BfT59zJA?t=982)
- Google Cloud Next'18 SF
  - [Day 1 Keynote](https://youtu.be/vJ9OaAqfxo4?t=2416) showing GKE On-Prem
  - [Day 3 Keynote](https://youtu.be/JQPOPV_VH5w?t=815) showing Stackdriver
    APM (Tracing, Code Search, Profiler, Google Cloud Build)
  - [Introduction to Service Management with Istio](https://www.youtube.com/watch?v=wCJrdKdD6UM&feature=youtu.be&t=586)
- [Google Cloud Next'18 London – Keynote](https://youtu.be/nIq2pkNcfEI?t=3071)
  showing Stackdriver Incident Response Management
