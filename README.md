# Dcache - A Distributed Cache with Raft-based Consensus

`Dcache` is a distributed cache system that enables horizontal scaling and high availability. It is designed with a TCP server that communicates with a gRPC cache server. The cache servers are organized into a cluster and maintain consistency through the Raft consensus protocol. In this system, candidate nodes can accept GET requests, while PUT requests are mediated by the leader node to ensure data consistency.

## Features

- **Distributed Cache**: Horizontally scalable cache system using multiple nodes.
- **Raft Consensus**: Ensures data consistency and fault tolerance across cache servers.
- **Leader Election**: Leader nodes handle PUT requests, while candidate nodes can respond to GET requests.
- **gRPC-based Communication**: Cache servers communicate using the gRPC protocol for efficient, low-latency requests.
- **TCP-based Request Handling**: The system accepts cache operations (GET, PUT) over TCP and processes them accordingly.

## Architecture Overview

### 1. **Cluster and Raft Protocol**

- The cache servers form a cluster using the Raft protocol to maintain consistency across all nodes.
- The Raft protocol handles leader election, log replication, and fault tolerance within the cache cluster.
- The leader node is responsible for mediating PUT requests (write operations) to ensure data consistency.
- Candidate nodes can handle GET requests (read operations), reducing the load on the leader.

### 2. **TCP Server**

- The TCP server acts as the interface between clients and cache nodes.
- It forwards cache operations (GET, PUT) to the appropriate node, based on the leader election.

### 3. **gRPC Cache Server**

- The gRPC-based cache server enables the coordination between nodes in the cluster.
- It allows nodes to join or leave the cluster and facilitates communication between nodes.
