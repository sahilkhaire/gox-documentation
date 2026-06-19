# Kafka (kafkajs) → Go

gox does not ship a Kafka client — use a mature Go library. This guide maps kafkajs patterns.

## Library choice

| Node.js | Go | Notes |
|---------|-----|-------|
| `kafkajs` | [segmentio/kafka-go](https://github.com/segmentio/kafka-go) | Most popular, simple API |
| `node-rdkafka` | [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | librdkafka bindings, production-grade |

**Recommendation:** `kafka-go` for most migrations; Confluent client for advanced ops.

```bash
go get github.com/segmentio/kafka-go
```

## Producer

### kafkajs

```js
const { Kafka } = require('kafkajs');
const kafka = new Kafka({ brokers: ['localhost:9092'] });
const producer = kafka.producer();
await producer.connect();
await producer.send({ topic: 'events', messages: [{ value: JSON.stringify({ id: 1 }) }] });
```

### kafka-go

```go
import "github.com/segmentio/kafka-go"

w := kafka.NewWriter(kafka.WriterConfig{
    Brokers: []string{"localhost:9092"},
    Topic:   "events",
})
defer w.Close()

err := w.WriteMessages(ctx, kafka.Message{
    Value: []byte(`{"id":1}`),
})
```

## Consumer

### kafkajs

```js
const consumer = kafka.consumer({ groupId: 'my-group' });
await consumer.connect();
await consumer.subscribe({ topic: 'events' });
await consumer.run({
  eachMessage: async ({ message }) => {
    console.log(message.value.toString());
  },
});
```

### kafka-go

```go
r := kafka.NewReader(kafka.ReaderConfig{
    Brokers: []string{"localhost:9092"},
    Topic:   "events",
    GroupID: "my-group",
})
defer r.Close()

for {
    m, err := r.ReadMessage(ctx)
    if err != nil {
        break
    }
    fmt.Println(string(m.Value))
}
```

## Pattern mapping

| kafkajs | kafka-go |
|---------|----------|
| `producer.send()` | `Writer.WriteMessages()` |
| `consumer.run({ eachMessage })` | `Reader.ReadMessage()` loop |
| `admin.createTopics()` | `kafka.Dial` + `CreateTopics` |
| `compression: CompressionTypes.GZIP` | `kafka.Gzip` codec in WriterConfig |
| Idempotent producer | `Writer.RequiredAcks`, `Writer.Async` config |

## With gox/queue

For Redis-backed job queues (Bull/BullMQ style), use `gox/queue` instead of Kafka when:

- Lower throughput is acceptable
- You already run Redis
- Jobs need retries/DLQ with less ops overhead

Use Kafka when you need event streaming, replay, or multi-consumer fan-out at scale.

## Further reading

- [kafka-go README](https://github.com/segmentio/kafka-go)
- [Confluent Go client docs](https://docs.confluent.io/kafka-clients/go/current/overview.html)
