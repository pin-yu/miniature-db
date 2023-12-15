# Miniature DB

My miniature implementation of database management system

## What v0.0.1 supports

### Statements

1. create tables
2. insert records
3. select record

grammar

```text
datatype        = int | char ( numeric_token )
constant        = string_token | numeric_token
constants       = constant [ , constant ]
field           = id_token
fields          = field [ , field ]
field_def       = field datatype
field_defs      = field_def [ , field_def ]
project         = field
projects        = project [ , project ]
table           = field
tables          = table [ , table ]
expression      = field | constant
term            = expression = expression
predicates      = term [ and predicates ]

update_cmd      = create_table | insert
create_table    = create table id_token ( field_defs );
insert          = insert into id_token ( fields ) values ( constants );

query           = select project from tables where predicates;
```

### Support types

- char (fixed-length character string)
- int  (32-bit signed integer)
