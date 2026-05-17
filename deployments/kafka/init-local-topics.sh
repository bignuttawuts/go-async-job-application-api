kafka-topics --bootstrap-server kafka-1:29092 \
--create \
--if-not-exists \
--topic local.fct.recruitment.job-applications.v1 \
--partitions 3 \
--replication-factor 3 \
--config retention.ms=43200000 \
--config min.insync.replicas=2
