-- noinspection SqlDialectInspectionForFile
-- noinspection SqlNoDataSourceInspectionForFile
CREATE TABLE IF NOT EXISTS encrypted_data(id varchar(100) NOT NULL PRIMARY KEY,encrypted_text  BLOB,encryption_key varchar(100), InsertedDatetime DATETIME);