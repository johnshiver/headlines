BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;

DROP TABLE headline CASCADE;
DROP TABLE data_source CASCADE;

END TRANSACTION;