-- +goose Up
CREATE TABLE candlesticks (
    id SERIAL PRIMARY KEY,
    symbol  TEXT NOT NULL,
	start_ts   TIMESTAMPTZ NOT NULL,
	open   NUMERIC NOT NULL,
	close   NUMERIC NOT NULL,
	high    NUMERIC NOT NULL,
	low NUMERIC NOT NULL,
-- open_time   TIMESTAMPTZ NOT NULL,
-- close_time  TIMESTAMPTZ NOT NULL,
--	trades NUMERIC NOT NULL,
	volume  NUMERIC NOT NULL
);


-- +goose Down
DROP TABLE candlesticks;
