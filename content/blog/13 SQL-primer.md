+++
date = "09 sep 2018"
categories = ["SQL", "eli5"]
tags = ["SQL", "eli5"]
slug = "sqlite3-primer"
title = "Sqlite3: A Primer"
+++

# SQL Reference

## SQLite3 Specific

### Transactions

SQLite transactions are ACID (Atomic, Consistent, Isolated and Durable). This means that data is not lost in the event of crashes, power failures or operating system dumps.

**Atomic**
	- Meaning changes cannot be broken down into smaller parts; either the entire transaction is commit or nothing at all. 

**Consistent**
	- A transaction must ensure to change the database from one valid state to another. If a transaction results in invalid data, the database will revert its previous state.

**Isolated**
	- Transactions are isolated from one client to the next. When changes are made to the database the changes are only visible to that client until it is committed. If two people concurrently make changes to a database, they will be transacted sequentially. It does not ensure the order of transactions. 

**Durability**
	- If a transaction is successfully committed, the changes are permanent going forward, if it is interrupted during the commit it will revert to its previous state.

![Julia Evans ACID drawing](/images/acid.svg "ACID in cartoon format")

**Please check out Julia Evans work at https://jvns.ca**

SQLite3 uses auto-commit mode by default.

### Select

Used to query data from one or more tables. Example below:

```SQL
SELECT DISTINCT column_list
FROM table_list
  JOIN table ON join_condition
WHERE row_filter
ORDER BY column
LIMIT count OFFSET offset
GROUP BY column
HAVING group_filter;
```

A breakdown of each clause:

- ```ORDER BY``` sorts the result set,

- ```DISTINCT``` gets unique rows in a table,
	
- ```WHERE``` is a filter such as, ```WHERE employees IS NULL```,

- ```LIMIT   OFFSET``` will return a set number of rows,

- ```GROUP BY``` sorts columns into groups and applies functions to each group,

- ```HAVING``` is another filter that checks for presence of the operator.

`SELECT * FROM database_table` is acceptable usage in testing or development but in production the query should be explicit.

### Insert

Used to insert new data be it single row, multiple rows, and default data into tables.

```SQL
INSERT INTO table1 (
 column1,
 column2 ,..)
VALUES
 (
 value1,
 value2 ,...),
 (
 value1,
 value2 ,...),
        ...
 (
 value1,
 value2 ,...);
```

The above example shows how to insert multiple rows into a table. 

```SQL
INSERT INTO artists (name)
VALUES
 ("Buddy Rich"),
 ("Candido"),
 ("Charlie Byrd");
 ```

 An example of how to insert into column ```name``` within table ```artists```.

 Data can also be inserted via a SELECT statement.

### Update

Update already existing data in a table.

```SQL
UPDATE table
SET column_1 = new_value_1,
    column_2 = new_value_2
WHERE
    search_condition 
ORDER column_or_expression
LIMIT row_count OFFSET offset;
```

the ```UPDATE``` clause informs SQL that some part of the table is going to be amended. the ```SET``` clause signifies and allows the updating of the column/s that precede it. On the left of the assignment operator (=) is the column to be updated and on the right the new value, expression or data. Whilst the ```WHERE``` operator is optional if it is omitted all rows will be updated.


### Delete

The delete statement allows for deletion of one, multiple or all rows within a table.

```SQL
DELETE
FROM
 table
WHERE
 search_condition;
 ```

 Like updating a table, ```WHERE```, ```ORDER BY``` and ```LIMIT``` clauses can be applied to a delete operation. In addition deletes can be orchestrated through the ```SELECT``` operator.


## JOINS

### Inner Joins

In relation databases, data is often distributed amongst many related tables. Foreign Keys are used to associate these tables.

`INNER JOIN` clause will combine columns from correlated tables.

In fig 1.1 Table A's 'f' column is compared with table B's 'f' column. If the value of the 'f' column in the A table equals that of B's 'f' column it will return the match. More simply, `INNER JOIN` clauses return rows from table A that have corresponding rows in the B table.

![input](/images/SQLite-Inner-Join-Example.png "example of inner join")

Fig 1.1 Inner Join

`INNER JOIN` may connect more than two tables. This will require two inner join clauses form the `SELECT` statement. 

```SQL
SELECT
 trackid,
 tracks.name AS Track,
 albums.title AS Album,
 artists.name AS Artist
FROM
 tracks
INNER JOIN albums ON albums.albumid = tracks.albumid
INNER JOIN artists ON artists.artistid = albums.artistid;
```

Example code of a query joining three tables by their Artist. (from Chinook.db found at SQLite's website tutorial.)

![input](/images/SQLite-Inner-Join-3-tables.jpg "result set of inner join on three tables.")

Fig. 1.2 is the result set of the above query.

### Full Outer Join

**this command is not in SQLite3**

Theoretically, the result of a Right Join and Left Join with `NULL` values for every column of the table that does not have a matching row.

This is technically not available in SQLite3, however they do offer a workaround. 

Given that SQLite3 does not have `RIGHT JOIN` or `FULL OUTER JOIN` it uses both `UNION` and `LEFT JOIN` to emulate it.

It achieves this by switching the position of the `LEFT JOIN` clause over both columns. SQLite also uses `UNION ALL` to duplicate rows from the result set of both queries. Finally, the use of a `WHERE` clause will remove rows already included in the result set of the first `SELECT` statement. See the below code for an example.

```SQL

SELECT d.type,
         d.color,
         c.type,
         c.color
FROM dogs d
LEFT JOIN cats c USING(color)
UNION ALL
SELECT d.type,
         d.color,
         c.type,
         c.color
FROM cats c
LEFT JOIN dogs d USING(color)
WHERE d.color IS NULL;
```

## Terms/ Definitions

### Primary Key

In order to qualify as a relational table, it must have a primary key.

The **primary key** consists of one or more columns whose data contained within is used to uniquely identify each row in the table. 
*metaphor* If rows were mailboxes the primary key would be the street address.

To be a true primary key it must be:
- Unique (data, not the name),
- Must not be NULL or ""

Tables must have primary keys, and these are stored in an index which is used to enforce the uniqueness requirement. Being indexed, accessing it does not necessitate scanning the entire table.

### Foreign Key

A **foreign key** is one or more columns in a table that refers to the primary key or unique identifier in another table. 

They unlike primary keys **can contain NULL, blank or duplicate values**. As they allow duplication it is best not to use a foreign key as a primary key, unless it is a one-to-one relationship.

![input](/images/foreign_key.png "example foreign key") 

In the above picture we have a Students and Cities table. PERSON_ID and CITY_ID are unique, and therefore make good Primary keys (which they are). Inside the Students table it is possible for several students to be born in the same city. As such the BIRTH_PLACE column is the foreign key to the Cities table. This ensures that there is relationship between the tables and that relation points to something unique in the other table.

### Union's

*BLUF: combines result set of two or more queries into a single result set.*

`UNION` by default removes duplicate rows whereas `UNION ALL` does not. As `UNION ALL` does not remove duplicates it will process faster. Both statements have the following rules:

- Must have the same number of columns,
- Corresponding columns must be same data type,
- The column names of the first query will determine the column of the combined result set,
- Any ```GROUP BY``` and ```HAVING``` clauses are applied to each individual query, not the final result set,
- The ```ORDER BY``` is applied to the combined result set, not the individual result sets.

### Transaction

The basis of all interactions with the database. Inserts, updates, deletes, commits and table creation and deletion are all transactions.
