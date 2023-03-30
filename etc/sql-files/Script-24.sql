DROP TABLE IF EXISTS store CASCADE;

CREATE TABLE IF NOT EXISTS store (
	id SERIAL PRIMARY KEY,
	filename TEXT UNIQUE NOT NULL,
	uploaded TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	bytes BYTEA NOT NULL,
	byte_size INT NOT NULL
);

ALTER TABLE store ALTER bytes SET STORAGE EXTERNAL;

INSERT INTO store (filename, bytes) VALUES (
	'my-file.png', 'file-contents'::bytea
);

SELECT * FROM store;

SELECT reltoastrelid::regclass from pg_class where relname = 'store';
SELECT * FROM pg_toast.pg_toast_17846;

SELECT DISTINCT chunk_id, count(chunk_seq) AS chunk_count FROM pg_toast.pg_toast_17846 GROUP BY chunk_id;