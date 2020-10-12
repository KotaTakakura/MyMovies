ALTER TABLE movies ADD FULLTEXT INDEX display_namengram_idx (display_name) WITH PARSER ngram;
ALTER TABLE movies ADD FULLTEXT INDEX description_ngram_idx (description) WITH PARSER ngram;