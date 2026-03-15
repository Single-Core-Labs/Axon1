CREATE TABLE IF NOT EXISTS traces (
    trace_id UUID,
    project_id String,
    model String,
    provider String,
    prompt String,
    completion String,
    prompt_tokens UInt32,
    completion_tokens UInt32,
    total_tokens UInt32,
    cost_usd Float64,
    latency_ms UInt32,
    status String,
    error_msg String,
    created_at DateTime
) ENGINE = MergeTree()
ORDER BY (project_id, created_at);
