[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Python: sql

- [Reference](#reference)
- [`sql`](#sql)

[↑ top](#python-sql)
<br><br><br><br>
<hr>










#### Reference

- [SQLAlchemy](http://docs.sqlalchemy.org)

[↑ top](#python-sql)
<br><br><br><br>
<hr>









#### `sql`

```python
# http://docs.sqlalchemy.org/en/rel_0_9/core/connections.html
import sqlalchemy
 
class _db:
    engine = None
    name = None
    connection_string = None
 
    def __init__(self, n, cs):
        self.name = n
        self.connection_string = cs
 
    def __call__(self):
        return self.db()
 
    def db(self):
 
        connect_args = {}
            
        if not self.engine:
            self.engine = sqlalchemy.create_engine(
                self.connection_string,
                echo=False,
                logging_name=name,
                # http://docs.sqlalchemy.org/en/rel_0_9/core/pooling.html
                pool_recycle=3600,
            )
        return self.engine
 
    def execute(self, *a, **ka):
        return self.db().execute(*a, **ka)
 
 
conn_dict = {
	"MY_MYSQL_DATABASE_NAME": 'mysql://MY_USER:MY_PASSWORD@MY_HOST:MY_PORT/MY_DATABASE_NAME?connect_timeout=3',
	"MY_POSTGRES_DATABASE_NAME": "postgres://MY_USER:MY_PASSWORD@MY_HOST:MY_PORT/MY_DATABASE_NAME",
}
 
for name, connection_string in conn_dict.iteritems():
    globals()[name] = _db(name, connection_string)
 
 
# then import like this
import database as db
 
data_list = []
select_query = """
SELECT col1, col2

FROM schema.table

WHERE col3 = %(value3)s
AND col4 = %(value4)s
"""
 
cur = db.MY_MYSQL_DATABASE_NAME.execute(select_query, {'value3': "hello", 'value4': 123})
for r in cur:
	data_list.append({
		'col1' : r['col1'] or '',
		'col2' : r['col2'] or 0,
	})
 
import csv
columns = ['col1', 'col2']
 
fpath = "data.csv"
with open(fpath, 'wb') as f:
    w = csv.DictWriter(f, fieldnames=columns, delimiter=',')
    w.writeheader()
    for elem in data_list:
        w.writerow(elem)
```

[↑ top](#python-sql)
<br><br><br><br>
<hr>