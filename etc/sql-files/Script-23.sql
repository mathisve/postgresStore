CREATE TABLE test_store (
	a INTEGER,
	b INTEGER,
	c VARCHAR
);

BEGIN;
INSERT INTO test_store (A, B, C) VALUES (1, 1, 'la la la');
COMMIT;

SELECT * FROM test_store;

SELECT tableoid, * FROM test_store;
--16557

--SELECT * FROM get_raw_page('public.test_store'::text, 0::bigint);

SELECT * FROM pg_class WHERE relname LIKE 'pg_toast_%';
SELECT * FROM public.pg_toast_16557;