## v0.5.9

- Avoid unnecessarily generating diff for `BOOLEAN` type on mysqldef [#49](https://github.com/proproto/cloudsqldef/pull/49)

## v0.5.8

- Add `--skip-drop` option to skip `DROP` statements [#44](https://github.com/proproto/cloudsqldef/pull/44)

## v0.5.7

- Support `double precision` for psqldef [#42](https://github.com/proproto/cloudsqldef/pull/42)
- Support partial indexes syntax for psqldef [#41](https://github.com/proproto/cloudsqldef/pull/41)

## v0.5.6

- Fix ordering between `NOT NULL` and `WITH TIME ZONE` for psqldef, related to v0.5.4 and v0.5.5
  [#40](https://github.com/proproto/cloudsqldef/pull/40)

## v0.5.5

- Support `time` with and without timezone for psqldef [#39](https://github.com/proproto/cloudsqldef/pull/39)

## v0.5.4

- Support `timestamp` with and without timezone for psqldef [#37](https://github.com/proproto/cloudsqldef/pull/37)

## v0.5.3

- Fix output length bug of psqldef since v0.5.0 [#36](https://github.com/proproto/cloudsqldef/pull/36)

## v0.5.2

- Support `timestamp` (without timezone) for psqldef [#34](https://github.com/proproto/cloudsqldef/pull/34)

## v0.5.1

- Support `SMALLSERIAL`, `SERIAL`, `BIGSERIAL` for psqldef [#33](https://github.com/proproto/cloudsqldef/pull/33)

## v0.5.0

- Remove `pg_dump` dependency for psqldef  [#32](https://github.com/proproto/cloudsqldef/pull/32)

## v0.4.14

- Show `pg_dump` error output on failure [#30](https://github.com/proproto/cloudsqldef/pull/30)

## v0.4.13

- Preserve line feeds when using stdin [#28](https://github.com/proproto/cloudsqldef/pull/28)

## v0.4.12

- Support reordering columns with the same names [#27](https://github.com/proproto/cloudsqldef/issues/27)

## v0.4.11

- Support enum [#25](https://github.com/proproto/cloudsqldef/issues/25)

## v0.4.10

- Support `ON UPDATE CURRENT_TIMESTAMP` on MySQL

## v0.4.9

- Fix issues on handling primary key [#21](https://github.com/proproto/cloudsqldef/issues/21)

## v0.4.8

- Add `--password-prompt` option to `mysqldef`/`psqldef`
  - This may be deprecated later once `--password` without value is properly implemented

## v0.4.7

- Add `-S`/`--socket` option of `mysqldef` to use unix domain socket
- Change `-h` option of `psqldef` to allow using unix domain socket

## v0.4.6

- Add support for fulltext index

## v0.4.5

- Support including hyphen in table names

## v0.4.4

- Support UUID data type for PostgreSQL and MySQL 8+

## v0.4.3

- Do not fail when view exists but just ignore views on mysqldef
  - Views may be supported later, but it's not managed by mysqldef for now

## v0.4.2

- Support generating `AFTER` or `FIRST` on `ADD COLUMN` on mysqldef

## v0.4.1

- Support `$PGSSLMODE` environment variable to specify `sslmode` on psqldef

## v0.4.0

- Support managing non-composite foreign key by changing CREATE TABLE
  - Note: Use `CONSTRAINT xxx FOREIGN KEY (yyy) REFERENCES zzz (vvv)` for both MySQL and PostgreSQL.
    In-column `REFERENCES` for PostgreSQL is not supported.
  - Note: Always specify constraint name, which is needed to identify foreign key name.
- Fix handling of DEFAULT NULL column

## v0.3.3

- Parse PostgreSQL's `"column"` literal properly
- Dump primary key with `--export` on PostgreSQL
- Prevent unexpected DDLs caused by data type aliases (bool, integer, char, varchar)

## v0.3.2

- Support `ADD PRIMARY KEY` / `DROP PRIMARY KEY` in MySQL
- Support parsing more data types for PostgreSQL: boolean, character
- Be aware of implicit `NOT NULL` on `PRIMARY KEY`
- Use `--schema-only` on `pg_dump` in psqldef

## v0.3.1

- Support `$MYSQL_PWD` environment variable to set password on mysqldef
- Support `$PGPASS` environment variable to set password on psqldef

## v0.3.0

- Support changing index on both MySQL and PostgreSQL
- Basic support of `CHANGE COLUMN` on MySQL
- All non-SQL outputs on apply/dry-run/export are formatted like `-- comment --`

## v0.2.0

- Support handling index on PostgreSQL
- Support `ADD INDEX` by modifying `CREATE TABLE` on MySQL

## v0.1.4

- Parse column definition more flexibly
  - ex) Both `NOT NULL AUTO_INCREMENT` and `AUTO_INCREMENT NOT NULL` are now valid
- Support parsing `character varying` for PostgreSQL
- Remove ` ` (space) before `;` on generated `ADD COLUMN`

## v0.1.3

- Fix SEGV and improve error message on parse error

## v0.1.2

- Drop all dynamic-link dependency from `mysqldef`
- "-- No table exists" is printed when no table exists on `--export`
- Improve error handling of unsupported features

## v0.1.1

- Release binaries for more architectures
  - New OS: Windows
  - New arch: 386, arm, arm64

## v0.1.0

- Initial release
  - OS: Linux, macOS
  - arch: amd64
- `mysqldef` for MySQL
  - Create table, drop table
  - Add column, drop column
  - Add index, drop index
- `psqldef` for PostgreSQL
  - Create table, drop table
  - Add column, drop column
