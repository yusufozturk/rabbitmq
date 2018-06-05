package metricsRabbitMQ

const (
	MetricPrefix = "rmq."
)

var (
	keysQueue = []string{
		// string type value keys: name, node
		"memory", // total in bytes?

		"message_stats.deliver_get",
		"message_stats.deliver_get_details.rate",

		"message_stats.deliver_no_ack",
		"message_stats.deliver_no_ack_details.rate",

		"message_stats.publish",
		"message_stats.publish_details.rate",

		"messages",
		"messages_details.rate",

		"messages_unacknowledged",
		"messages_unacknowledged_details.rate",

		"consumers",

		"messages_ram", // how much messages in RAW?
		"messages_ready_ram",
		"messages_unacknowledged_ram",

		"message_bytes",
		"message_bytes_ready",
		"message_bytes_unacknowledged",
		"message_bytes_ram",
		"message_bytes_persistent",

		// NOTICE: following keys are not available > v3.5.7
		"disk_reads",
		"disk_writes",
	}

	keysOverview = []string{
		// string type value keys: name, node
		"message_stats.publish",
		"message_stats.publish_details.rate",

		"message_stats.deliver_get",
		"message_stats.deliver_get_details.rate",

		"message_stats.deliver_no_ack",
		"message_stats.deliver_no_ack_details.rate",

		"queue_totals.messages",
		"queue_totals.messages_details.rate",

		"queue_totals.messages_ready",
		"queue_totals.messages_ready_details.rate",

		"queue_totals.messages_unacknowledged",
		"queue_totals.messages_unacknowledged_details.rate",

		"object_totals.consumers",
		"object_totals.queues",
		"object_totals.exchanges",
		"object_totals.connections",
		"object_totals.channels",

		"statistics_db_event_queue",
	}
)
