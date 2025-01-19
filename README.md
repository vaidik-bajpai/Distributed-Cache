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

## Future Improvements and Features

### 1. **Implementing Sharding**

- **Sharding** will be implemented to improve the scalability and distribution of cache data across nodes. With sharding, the cache system will split the data into smaller, more manageable parts (shards), which can be distributed across multiple nodes in the cluster. This allows the system to handle larger datasets efficiently while ensuring high availability and load balancing.
- Each shard will be managed by a separate group of nodes, potentially with its own Raft consensus group. This will improve the system’s overall performance and reliability by distributing read and write operations more evenly across the nodes.

### 2. **Extensive Testing of the Application**

- To ensure the reliability and robustness of `Dcache`, **extensive testing** will be conducted. This includes unit tests, integration tests, and load testing to simulate high-throughput scenarios.
- Testing will cover:
  - **Leader election and Raft protocol functionality**: Ensuring that the leader election process works smoothly and the Raft protocol maintains consistency across the cluster.
  - **Cache consistency**: Verifying that the cache data is consistent and correctly handled between nodes, even during failures and recovery.
  - **Sharding efficiency**: Ensuring that the sharding mechanism distributes data evenly and efficiently across the cluster.
  - **Fault tolerance**: Simulating node failures and ensuring the system remains highly available, with minimal disruption to services.

These improvements will help `Dcache` scale further, handle larger workloads, and provide a more reliable, consistent caching solution.
