# Kafka Calculator

A training project implementing a distributed calculator.
Uses microservice architecture.

---
```
Consists of:
- Calculation manager
  - Generates an expression
  - Sends an expression via Kafka to executor service
  - Waits for a response (result of a calculation)
  - Outputs the result
- Calculation executor
  - Receives an expression from the manager service via Kafka
  - Calculates the result of the expression
  - Sends the result via Kafka to the manager service
  
- Kafka broker
  - Contains:
    - Two topics:
        - expressions_topic
        - calculations_topic
    - Two groups:
        - exec_01
        - mgr_01
  
```

List of the services of docker-compose:
- calculation_manager
- calculation_executor
- calculation_zookeeper
- calculation_kafka

### Todo list:
- Add an expression generator service
- Add a gRPC connector to the generator and manager
- Add Redis cache for expressions
- Add Kafka router for different controllers
- Add an computation pipeline implementation (+ limiter)
