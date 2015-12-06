[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Python: redis

- [redis](#redis)

[↑ top](#python-redis)
<br><br><br><br><hr>


#### redis

> Redis is an open source, BSD licensed, advanced **key-value cache** and
> **store**. It is often referred to as a **data structure server** since keys
> can contain [strings](http://redis.io/topics/data-types-intro#strings),
> [hashes](http://redis.io/topics/data-types-intro#hashes),
> [lists](http://redis.io/topics/data-types-intro#lists),
> [sets](http://redis.io/topics/data-types-intro#sets), [sorted
> sets](http://redis.io/topics/data-types-intro#sorted-sets),
> [bitmaps](http://redis.io/topics/data-types-intro#bitmaps),
> [hyperloglogs](http://redis.io/topics/data-types-intro#hyperloglogs).
>
> [*Redis*](http://redis.io/)

```python
#!/usr/bin/python -u
# in database.py
import redis
 
conn_dict = {
    'MY_REDIS_DATABASE': {'host': 'MY_HOST', 'port': MY_PORT, 'password': 'MY_PASSWORD'},
}
 
for name, connection_string in conn_dict.iteritems():
    globals()[name] = redis.Redis(**(connection_string))
 
 
# then import like this
import database as db
 
found_values = set(1, 2, 3)
added = db.MY_REDIS_DATABASE.sadd("MY_KEY_1", *found_values)
not_in_key_2 = db.MY_REDIS_DATABASE.sdiff("MY_KEY_1", "MY_KEY_2")
not_in_key_1 = db.MY_REDIS_DATABASE.sdiff("MY_KEY_2", "MY_KEY_1")

```

[↑ top](#python-redis)
<br><br><br><br><hr>
