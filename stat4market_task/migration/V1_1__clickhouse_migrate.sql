CREATE DATABASE IF NOT EXISTS test;

DROP TABLE test.sequence_table;
DROP TABLE test.events;

CREATE TABLE IF NOT EXISTS test.events
(
    eventID Int64,
    eventType String,
    userID Int64,
    eventTime DateTime,
    payload String
) ENGINE = MergeTree
ORDER BY (eventID, eventTime)

CREATE TABLE IF NOT EXISTS test.sequence_table (
    next_id Int64
) ENGINE = Memory();

INSERT INTO test.sequence_table VALUES (1);