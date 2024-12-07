# Basic operations

Access control mechanism

• Security through views

• Stored procedures

• Grant and revoke


## GRANT AND REVOKE
### GRANT
**GRANT <privilege> ON <relation> TO <user> [WITH GRANT OPTION]**
```
GRANT INSERT ON Student TO Matthews
GRANT INSERT, UPDATE(GRADE) ON Student TO FARKAS
GRANT INSERT(NAME) ON Student TO Brown
```

Notes:

• When granting INSERT at the column level, you must include all the not null columns in the row

• GRANT command applies to base relations as well as views

### REVOKE
**REVOKE <privilege> ON <relation> FROM <user>**
```
REVOKE INSERT ON Student FROM Matthews
REVOKE UPDATE ON Student FROM FARKAS
REVOKE INSERT, UPDATE(GRADE) ON Student FROM FARKAS
REVOKE INSERT(NAME) ON Student FROM Brown
```