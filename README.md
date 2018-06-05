# RabbitMQ Metrics for monitoring

Tags

| KEY | TYPE | NOTES |
|-----|------|-------|
| name | string | 
| node | string | 


Api `Overview` Metrics

| KEY | TYPE | VALUE | NOTES |
|-----|------|-------|-------|
| overview.message_stats.publish | GAUGE | 
| overview.message_stats.publish_details.rate | GAUGE | 
| overview.message_stats.deliver_get | GAUGE | 
| overview.message_stats.deliver_get_details.rate | GAUGE | 
| overview.message_stats.deliver_no_ack | GAUGE | 
| overview.message_stats.deliver_no_ack_details.rate | GAUGE | 
| overview.queue_totals.messages | GAUGE | 
| overview.queue_totals.messages_details.rate | GAUGE | 
| overview.queue_totals.messages_ready | GAUGE | 
| overview.queue_totals.messages_ready_details.rate | GAUGE | 
| overview.queue_totals.messages_unacknowledged | GAUGE | 
| overview.queue_totals.messages_unacknowledged_details.rate | GAUGE | 
| overview.object_totals.consumers | GAUGE | 
| overview.object_totals.queues | GAUGE | 
| overview.object_totals.exchanges | GAUGE | 
| overview.object_totals.connections | GAUGE | 
| overview.object_totals.channels | GAUGE | 
| overview.statistics_db_event_queue | GAUGE | 



Api `Queues` Metrics

| KEY | TYPE | VALUE | NOTES |
|-----|------|-------|-------|
| queues.memory | GAUGE |  | total in bytes? |
| queues.message_stats.deliver_get | GAUGE | 
| queues.message_stats.deliver_get_details.rate | GAUGE | 
| queues.message_stats.deliver_no_ack | GAUGE | 
| queues.message_stats.deliver_no_ack_details.rate | GAUGE | 
| queues.message_stats.publish | GAUGE | 
| queues.message_stats.publish_details.rate | GAUGE | 
| queues.messages | GAUGE | 
| queues.messages_details.rate | GAUGE | 
| queues.messages_unacknowledged | GAUGE | 
| queues.messages_unacknowledged_details.rate | GAUGE | 
| queues.consumers | GAUGE | 
| queues.messages_ram | GAUGE |   | how much messages in RAW? |
| queues.messages_ready_ram | GAUGE | 
| queues.messages_unacknowledged_ram | GAUGE | 
| queues.message_bytes | GAUGE | 
| queues.message_bytes_ready | GAUGE | 
| queues.message_bytes_unacknowledged | GAUGE | 
| queues.message_bytes_ram | GAUGE | 
| queues.message_bytes_persistent | GAUGE | 
| queues.disk_reads | GAUGE | 
| queues.disk_writes | GAUGE | 

		 

## Build 


Build example

    go get -v github.com/MonitorMetrics/rabbitmq
    cd $GOPATH/src/github.com/MonitorMetrics/rabbitmq/examples/json.rabbitmq
    go build -o json.rabbitmq.bin


Start in debug mode

    ./json.rabbitmq.bin -debug 
   

For more detail, see source code.
