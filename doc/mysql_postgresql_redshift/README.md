[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# mysql, postgresql, redshift

- [Reference](#reference)
- [SQL and Relational database](#sql-and-relational-database)
- [Connect, Create, Index](#connect-create-index)
    - [MySQL](#mysql)
    - [PostgreSQL, Redshift](#postgresql-redshift)
- [Alter table](#alter-table)
    - [MySQL](#mysql-1)
    - [PostgreSQL, Redshift](#postgresql-redshift-1)
- [Update, Delete](#update-delete)
    - [MySQL](#mysql-2)
    - [PostgreSQL, Redshift](#postgresql-redshift-2)
- [Insert, Upsert](#insert-upsert)
    - [MySQL](#mysql-3)
    - [PostgreSQL, Redshift](#postgresql-redshift-3)
- [Import from csv](#import-from-csv)
    - [MySQL](#mysql-4)
    - [PostgreSQL](#postgresql)
    - [Redshift](#redshift)
- [Select](#select)
    - [MySQL](#mysql-5)
    - [PostgreSQL, Redshift](#postgresql-redshift-4)
 
[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


#### Reference

- [PostgreSQL Guide](http://www.postgresguide.com/)

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## SQL and Relational database

> **_SQL, structured query language,_** is a [*special-purpose programming
> langauge*](https://en.wikipedia.org/wiki/Special-purpose_programming_language)
> *designed for managing data held in a [relational database management
> systems](https://en.wikipedia.org/wiki/Relational_database_management_system)
> (RDBMS)*
>
> [*Wikipedia*](https://en.wikipedia.org/wiki/SQL)

SQL is perhaps the most popular programming language. SQL query reads like
plain English, and writing it doesn’t require much knowledge about inner
workings of database systems. But you **must know a few things** to prevent
performance issues. One simple query can easily stall its whole database
server.

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Connect, Create, Index
It's extremely important that you **create indexes** when you **create
tables**, or **add indexes with** *CREATE INDEX* statement. Indexing determines
how a program queries data. **Indexing makes queries run faster and more
efficiently**. Indexing does not change data: it creates a new data structure
that refers to the data. Index in database is just like an index for
dictionaries. Without indexing, you would have to read page by page until you find the
information that you need. Indexing enables you to search data in much faster
and more efficient ways. Usually indexing is done with [doubly linked
list](https://en.wikipedia.org/wiki/Doubly_linked_list) and
[b-tree](https://en.wikipedia.org/wiki/B-tree) data structures. Databases store
the minimum amount of data in leaf nodes and connect them using doubly linked
list because it’s good at handling frequent inserts and removals. And b-tree
maintains the balance of the tree, to minimize the computations of searching
data.


<br>

#### MySQL

```sql
# to install mysql-server
sudo apt-get -y install mysql-server;
 
# to help
mysql --help
 
# commands are
mysql [OPTIONS] [database]
 
# to connect to local mysql-server
sudo mysql;
 
# or to remote server
alias mydb='
mysql \
--host=MY_HOST \
--port=MY_PORT \
--user=MY_USER \
--password=MY_PASSWORD \
MY_DATABASE_NAME
'
 
# to describe table, schema
DESCRIBE myschema.example_table;
SHOW TABLES IN myschema;
 
# to create database or schema
CREATE SCHEMA IF NOT EXISTS myschema;
 
# to create a table
CREATE TABLE IF NOT EXISTS myschema.example_table (
  col1 VARCHAR(700) NOT NULL,
  col2 VARCHAR(700) DEFAULT '0',
  col3 FLOAT NOT NULL DEFAULT 0.0,
  col4 BLOB,
  col5 TINYINT(1),
  PRIMARY KEY(col1, col2)
)
;

CREATE TABLE IF NOT EXISTS myschema.example_table LIKE myschema.example_table_2
;

# to create an index
CREATE UNIQUE INDEX 'MY_INDEX_NAME'
ON myschema.example_table (col1)
;
```

<br>

#### PostgreSQL, Redshift

```sql
# to install postgresql
sudo apt-get -y install postgresql;
 
# to help
man psql;
 
# commands are
psql [option…] [dbname [username]]
 
# to connect to local postgresql server
sudo -u postgres psql
 
# or to remote server
alias mydb='
PGPASSWORD=MY_PASSWORD \
psql \
--host MY_HOST \
--port MY_PORT \
--username MY_USER \
--dbname MY_DATABASE_NAME
'
 
# to describe table, schema
\d myschema.example_table
\dt myschema.*
 
# to create database or schema
CREATE SCHEMA IF NOT EXISTS myschema;
 
# to create a table
CREATE TABLE IF NOT EXISTS myschema.example_table (
  col1 TEXT NOT NULL,
  col2 TEXT DEFAULT '0',
  col3 NUMERIC(5,4) NOT NULL DEFAULT 0.0,
  ts   TIMESTAMP,
  PRIMARY KEY(col1, col2)
)
;
 
# to create a Redshift table
CREATE TABLE IF NOT EXISTS myschema.example_table (
  day  DATE DISTKEY NOT NULL,
  ts   TIMESTAMP NOT NULL,
  col1 VARCHAR SORTKEY,
  col2 VARCHAR(9999),
  col3 SMALLINT DEFAULT 0,
  col4 VARCHAR(9999) DEFAULT '',
  col5 DOUBLE PRECISION DEFAULT 0.0
)
;
 
# to create an index
CREATE UNIQUE INDEX 'MY_INDEX_NAME'
ON myschema.example_table (col1)
;
```

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Alter table

<br>

#### MySQL

```sql
# to drop a column
CREATE TABLE myschema.example_table_backup LIKE myschema.example_table;
INSERT INTO myschema.example_table_backup SELECT * FROM myschema.example_table;
 
ALTER TABLE myschema.example_table DROP COLUMN column_to_drop;
 
# if failed, ...
DROP TABLE myschema.example_table;
CREATE TABLE myschema.example_table (LIKE myschema.example_table_backup);
INSERT INTO myschema.example_table SELECT * FROM myschema.example_table_backup;
 
# rename a table
RENAME TABLE myschema.example_table TO different_table;
 
# to rename a column
ALTER TABLE myschema.example_table CHANGE col1 col2 INT;
 
# to modify a column
ALTER TABLE myschema.example_table MODIFY col1 VARCHAR(250) default '0';
```

<br>

#### PostgreSQL, Redshift

```sql
# to drop a column
CREATE TABLE myschema.example_table_backup (LIKE myschema.example_table);
INSERT INTO myschema.example_table_backup (SELECT * FROM myschema.example_table);
 
ALTER TABLE myschema.example_table DROP COLUMN IF EXISTS column_to_drop;
 
# if failed, ...
DROP TABLE myschema.example_table;
CREATE TABLE myschema.example_table LIKE myschema.example_table_backup;
INSERT INTO myschema.example_table (SELECT * FROM myschema.example_table_backup);
 
# rename a table
ALTER TABLE myschema.example_table RENAME TO different_table;
 
# to rename a column
ALTER TABLE myschema.example_table RENAME COLUMN col1 TO col2;
 
# to modify a column
ALTER TABLE myschema.example_table ALTER COLUMN col1 TYPE TEXT;
ALTER TABLE myschema.example_table ALTER COLUMN col1 SET DEFAULT '0';
 
# to add a column
ALTER TABLE myschema.example_table ADD COLUMN col1 TEXT;
 
# to drop a column
ALTER TABLE myschema.example_table DROP COLUMN IF EXISTS col1;
 
# to modify a Redshift column
ALTER TABLE myschema.example_table ADD COLUMN col1_temp VARCHAR(9999) default '0';
UPDATE myschema.example_table SET col1_temp = col1;
ALTER TABLE myschema.example_table DROP COLUMN col1;
ALTER TABLE myschema.example_table RENAME COLUMN col1_temp TO col1;



# or use transaction
BEGIN TRANSACTION;

ALTER TABLE <TABLE_NAME> RENAME TO <TABLE_NAME>_TEMPORARY;

CREATE TABLE <TABLE_NAME> ( <NEW_COLUMN_DEFINITION> );

INSERT INTO <TABLE_NAME>
SELECT <COLUMNS> FROM <TABLE_NAME>_TEMPORARY;

DROP TABLE <TABLE_NAME>_TEMPORARY;

END TRANSACTION;
```

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Update, Delete

<br>

#### MySQL

```sql
# update rows
UPDATE myschema.example_table
SET col1 = val1, col2 = val2, col3 = val3, col4 = val4
WHERE col0 = primary_key
LIMIT 5
;
 
# delete rows
DELETE FROM myschema.example_table
WHERE col1 = val1
LIMIT 1
;
```

<br>

#### PostgreSQL, Redshift

```sql
# update values
UPDATE myschema.example_table
SET col1 = val1, col2 = val2, col3 = val3, col4 = val4
WHERE col0 = primary_key
LIMIT 5
;
 
# delete rows
DELETE FROM myschema.example_table
WHERE col1 = val1
LIMIT 1
;
```


[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Insert, Upsert

<br>
#### MySQL

```sql
# to insert
INSERT INTO myschema.example_table (col1, col2, col3)
VALUES ("a.com", "gyuho", 10.0)
;
 
INSERT INTO myschema.example_table (col1, col2, col3)
VALUES ("a.com", "gyuho", 10.0),
("b.com", "gyuho", 11.0),
("c.com", "gyuho", 12.0)
;
 
INSERT INTO myschema.example_table
SELECT * FROM myschema.different_table
;
```
```sql
CREATE TABLE myschema.example_table (
    col0 VARCHAR(150) NOT NULL, 
    col1 INTEGER,
    col2 VARCHAR(150),
    col3 INTEGER,
    col4 VARCHAR(150),
    PRIMARY KEY(col0)
);
 
REPLACE INTO myschema.example_table
SET col0 = 'a.com'
, col1 = 1
, col2 = 'A'
, col3 = 2
, col4 = 'B'
;
 
SELECT * FROM myschema.example_table;
 
REPLACE INTO myschema.example_table
SET col0 = 'a.com'
, col1 = 11
, col2 = 'A'
, col3 = 22
, col4 = 'B'
;
 
SELECT * FROM myschema.example_table;
 
DROP TABLE myschema.example_table;
```

<br>
#### PostgreSQL, Redshift

```sql
# to insert
INSERT INTO myschema.example_table (col1, col2, col3)
VALUES ("a.com", "gyuho", 10.0)
;
 
INSERT INTO myschema.example_table (col1, col2, col3)
VALUES ("a.com", "gyuho", 10.0),
("b.com", "gyuho", 11.0),
("c.com", "gyuho", 12.0)
;
 
INSERT INTO myschema.example_table 
(SELECT * FROM myschema.different_table)
;
```

Newer **PostgreSQL** versions have
[**_UPSERT_**](https://wiki.postgresql.org/wiki/UPSERT). For older versions:

```sql
CREATE TABLE myschema.example_table (
    col0 TEXT PRIMARY KEY, 
    col1 INTEGER,
    col2 TEXT,
    col3 INTEGER,
    col4 TEXT
);
 
CREATE OR REPLACE FUNCTION upsert_function(
    primary_key   TEXT, 
    val1          INTEGER,
    val2          TEXT,
    val3          INTEGER,
    val4          TEXT
) RETURNS VOID AS
$$
BEGIN
    LOOP
        UPDATE myschema.example_table
        SET col1 = val1, col2 = val2, col3 = val3, col4 = val4
        WHERE col0 = primary_key;
 
        IF found THEN
            RETURN;
        END IF;
 
        BEGIN
            INSERT INTO myschema.example_table(col0, col1, col2, col3, col4)
            VALUES (primary_key, val1, val2, val3, val4);
            RETURN;
        EXCEPTION WHEN unique_violation THEN
            -- do nothing, and loop to try the UPDATE again
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;
 
SELECT upsert_function('a.com', 1, 'A', 2, 'B');
 
SELECT * FROM myschema.example_table;
 
SELECT upsert_function('a.com', 11, 'A', 22, 'B');
 
SELECT * FROM myschema.example_table;
 
 
DROP FUNCTION IF EXISTS upsert_function(
    primary_key   TEXT, 
    val1          INTEGER,
    val2          TEXT,
    val3          INTEGER,
    val4          TEXT
);
 
DROP TABLE myschema.example_table;
```


[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Import from csv

<br>
#### MySQL

```sql
# to copy from local csv file
LOAD DATA LOCAL INFILE 'MY_CSV_FILEPATH' REPLACE INTO TABLE myschema.example_table
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\r\n'
IGNORE 1 LINES
(col1, col2, col3, col4, col5)
;
```

<br>
#### PostgreSQL

```sql
# to copy from a local csv file
# and save in a bash script
# and run sudo sh ./script.sh
 
#!/bin/bash
sudo PGPASSWORD=MY_PASSWORD \
psql --host MY_HOST \
--port MY_PORT \
--username MY_USER \
--dbname MY_DATABASE_NAME \
-c "\\copy myschema.example_table(col1, col2, col3, col4, col5) FROM 'MY_CSV_PATH' with delimiter ',' csv header;"
;
```

<br>
#### Redshift

```sql
# to copy from a local csv file
BEGIN;
 
COPY myschema.example_table(col1, col2, col3, col4, col5) FROM 'MY_S3_PATH'
CREDENTIALS 'aws_access_key_id=MY_ACCESS_KEY;aws_secret_access_key=MY_SECRET_KEY'
IGNOREHEADER 1
NULL as 'MY_NULL_STRING'
TRUNCATECOLUMNS
EMPTYASNULL
CSV
;
 
COMMIT;
```

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>


## Select

<br>

#### MySQL

```sql
SELECT tb2.customer_name
, COUNT(DISTINCT tb1.customer_id)
 
FROM her.table tb1
 
LEFT JOIN your.table tb2
ON tb2.customer_id = tb1.customer_id
 
WHERE tb1.customer_name IN ('name1', 'name2')
 
GROUP BY 1
ORDER BY 2 DESC
;
```

<br>

#### PostgreSQL, Redshift

```sql
WITH subdata1 AS (
    SELECT tb3.company_type
    , 'google' AS partner
    , tb1.company_id
    , SUM(tb2.earnings) AS profit
    , SUM(tb2.clicks) AS value2
    , SUM(tb2.ad_requests) AS value3

    FROM their.company_information tb1

    INNER JOIN her.partner_information tb2
    ON tb2.company_address = tb1.company_address
    AND DATE_TRUNC('day', tb2.date) = '{date_to_run}'

    INNER JOIN their.table tb3
    ON tb3.company_address = tb1.company_address

    WHERE DATE_TRUNC('day', tb1.day) = '{date_to_run}'

    GROUP BY 1, 2, 3
)

SELECT '{date_to_run}' AS date
, dc.customer_id
, subdata1.company_type
, subdata1.partner
, subdata1.company_id

, rvt.weight * subdata1.profit AS profit
, subdata1.value3
, subdata1.value2

, et.session
, et.ip

FROM subdata1

LEFT JOIN(
    SELECT company_type
    , company_id
    , COUNT(DISTINCT session_id) AS session
    , COUNT(DISTINCT ip) AS ip

    FROM new.table

    WHERE DATE_TRUNC('day', event_time) = '{date_to_run}'
    AND partner = 'google'

    GROUP BY 1, 2
) et
ON et.company_id = subdata1.company_id
AND et.company_type = subdata1.company_type

LEFT JOIN dm.campaign_customer dc
ON dc.company_type = subdata1.company_type

LEFT JOIN dm.partner_information rvt
ON rvt.partner = subdata1.partner
AND rvt.date1 <= '{date_to_run}'
AND rvt.date2 >= '{date_to_run}'
;







WITH subdata1 AS (
    SELECT tb2.company_type
    , CASE WHEN tb2.company_id = '' THEN '0'
      WHEN tb2.company_id IS null THEN '0'
      ELSE tb2.company_id
      end AS company_id
    , SUM(tb1.estimated_gross_revenue) AS profit
    , SUM(tb1.col2) AS value2
    , SUM(tb1.col3) AS value3

    FROM source.our_revenue tb1

    INNER JOIN source.mapping_table tb2
    ON tb2.customer_id = tb1.customer_id

    WHERE DATE_TRUNC('day', tb1.date_utc) = '{date_to_run}'

    GROUP BY 1,2
),

subdata2 AS (
    SELECT tb1.company_type
    , tb1.company_id
    , tb1.profit
    , tb1.value2
    , tb1.value3

    , COUNT(DISTINCT etb.session_id) AS session
    , COUNT(DISTINCT etb.ip) AS ip

    FROM subdata1 tb1

    LEFT JOIN new.table etb
    ON etb.company_type = tb1.company_type
    AND etb.company_id = tb1.company_id

    WHERE DATE_TRUNC('day', etb.event_time) = '{date_to_run}'
    AND etb.partner = 'yahoo'

    GROUP BY 1,2,3,4,5
),

subdata3 AS (
    SELECT dc.customer_id
    , st2.company_type
    , 'yahoo'::text AS partner
    , st2.company_id
    , st2.profit
    , st2.value2
    , st2.value3
    , st2.session
    , st2.ip

    FROM subdata2 st2

    LEFT JOIN dm.campaign_customer dc
    ON dc.company_type = st2.company_type
)

SELECT '{date_to_run}' AS date
, subdata3.customer_id
, subdata3.company_type
, subdata3.partner
, subdata3.company_id
, rvt.weight * subdata3.profit AS profit
, subdata3.value2
, subdata3.value3
, subdata3.session
, subdata3.ip

FROM subdata3

LEFT JOIN dm.partner_information rvt ON rvt.partner = subdata3.partner
AND rvt.date1 <= '{date_to_run}'
AND rvt.date2 >= '{date_to_run}'
;







WITH agg1 AS (
    SELECT customer AS customer_id
    , customer_id
    , revenue AS profit
    , col2 AS value2
    , col3 AS value3

    FROM my.table

    WHERE DATE_TRUNC('day', date) = '{date_to_run}'
    AND revenue > 0
    AND customer <> 'company'

    GROUP BY 1,2,3,4,5
),

agg2 AS (
    SELECT ss.customer_id
    , COUNT(DISTINCT ss.clientip) AS ips_old

    FROM agg1

    LEFT JOIN old.table ss
    ON ss.customer_id = agg1.customer_id
    
    WHERE ss.customer_id <> ''
    AND DATE(convert_timezone('pdt', 'utc', ss.datestart)) = '{date_to_run}'

    GROUP BY 1
),

agg3 AS (
    SELECT et.customer_id
    , COUNT(DISTINCT et.ip) AS ips_ne

    FROM agg1

    LEFT JOIN new.table et
    ON et.customer_id = agg1.customer_id

    WHERE et.customer_id <> ''
    AND DATE(et.event_time) = '{date_to_run}'

    GROUP BY 1
),

agg4 AS (
    SELECT tb1.customer_id
    , COALESCE(tb1.ips_old, tb2.ips_new) AS ips

    FROM agg2 tb1

    LEFT JOIN agg3 tb2
    ON tb2.customer_id = tb1.customer_id
),

agg5 AS (
    SELECT agg1.customer_id

    , CASE WHEN oldtt.customer_id IS NOT NULL AND newtt.customer_id IS NULL THEN 'old_customer'
        WHEN oldtt.customer_id IS NULL AND newtt.customer_id IS NOT NULL THEN 'new_customer'
        WHEN oldtt.customer_id IS NULL AND newtt.customer_id IS NULL THEN 'unknown'
        ELSE 'both'
        END AS customer_type

    , COALESCE(oldtt.machine_id, newtt.machine) AS identifier

    , SUM(agg1.profit) AS profit
    , SUM(agg1.value2) AS value2
    , SUM(agg1.value3) AS value3
    , SUM(tg4.ips) AS ips

    FROM agg1

    LEFT JOIN agg4 tg4 
    ON tg4.customer_id = agg1.customer_id

    LEFT JOIN source.customer_id_mapping oldtt
    ON oldtt.customer_id = agg1.customer_id

    LEFT JOIN source.mapping_table newtt
    ON newtt.customer_id = agg1.customer_id

    GROUP BY 1, 2, 3
) 

SELECT '{date_to_run}' AS date
, agg5.customer_id
, agg5.customer_type
, 'gyuho' AS manager
, agg5.identifier
, agg5.profit
, agg5.value2
, agg5.value3
, agg5.ips

FROM agg5

WHERE agg5.identifier != ''
```

[↑ top](#mysql-postgresql-redshift)
<br><br><br><br><hr>
