# full-cycle-2.0-kafka

Files I produced during the Apache Kafka classes of my [Microservices Full Cycle 2.0 course](https://drive.google.com/file/d/1MdN-qK_8Pfg6YI3TSfSa5_2-FHmqGxEP/view?usp=sharing).

# Theory

## Kafka basic functioning

![Producers, Kafka Cluster, Brokers and Consumers](./images/kafka-basic-functioning.png)

## Topics

Usually you will want to have the topic partitions inside different brokers. So you can imagine that the partitions 1, 2 and 3 are in different machines. If one machine dies, you have 2/3 of the messages of this topic in other brokers.

![Shows that a topic has multiple partitions inside different brokers](./images/kafka-topic-partitions.png)

You can also choose to duplicate the partitions and distribute them along all brokers using Replication Factor = 2. In this way, if one broker dies, you will have the same partition in another one.

![Shows that a topic has multiple partitions inside different brokers](./images/kafka-topic-partitions-replication.png)

## Partitions

Each topic has associated partitions. Kafka saves messages sequentially to disk!

![Shows that a partition of a kafka topic is like an array of messages](./images/kafka-partition.png)

## Message structure

- Headers are metadata you can include in the message
- Key is used to send the same type of messages to the same partition
- Value is the payload itself
- Timestamp can be set by the producer or the kafka cluster when the message arrive

![Shows that a message has 4 fields: Headers, Key, Value and Timestamp](./images/kafka-message-registry.png)

