<p align="center">
	<img src="assets/joyboy.jpg" height="200" border="0" alt="RQL">
	<br/>
</p>

<h1 align="center">CURD Boy</h1>
<p align="center">
  A tool to generate code that support dynamic curd operations based on <a href="https://github.com/ent/ent">ent</a>
</p>

--- 

# Table of contents
* [Quick Introducation](#quick-inttoducation)
  * [Setup a go environment](#setup-a-go-environment)
  * [Install ent](#install-ent)
  * [Create Schema](#create-a-schema)
  * [Generate code](#generate)
* [Specification](#specification)
  * [Query Language](#query-language)
    * [Overview](#query-language-overview)
    * [Filter](#filter)
      * [Example](#filter-example)
      * [Grammar](#filter-grammar)
    * [Fields](#fields)
      * [Example](#fields-example)
      * [Grammar](#fields-grammar)

# Quick Introducation
Before you read this introducation, i suggest you to be familary with [ent](https://github.com/ent/ent) firstly

## Setup a go environment
```shell
mkdir curdboy-playground
cd curdboy-playground
go mod init curdboy-playground
```

## Install ent
```shell
go install entgo.io/ent/cmd/ent@latest
```

## Create a schema
```shell
go run entgo.io/ent/cmd/ent init User
```

edit the User schema at ent/schema/user.go
```go

package schema

import (
   "entgo.io/ent" 
   "entgo.io/ent/schema/field" 
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
      field.String("firstname"),
      field.String("lastname"),
      field.Int("age"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return nil
}
```

Run ```go generate``` to generate ent code
```
go generate ./ent
```

## Install curdboy 
```
go install github.com/pigfall/curdboy@latest
```

## Generate
Run command to generate dynamic curd code 
```go
cbc --schemaDirPath ./ent/schema --entTargetDirPath ./ent --targetDirPath ./curd
```

Take a loot at the ./curd/curd_user.go
```go
func UserQuery (ctx context.Context,req *QueryRequest,entCli *ent.Client)([]*ent.User,error ){
	var pred predicate.User
	if len(req.Filter) > 0{
		filterExpr,err := filter.ParseFilter(req.Filter)
		if err != nil{
			return nil,err
		}
		pred,err = ToUserPredicate(filterExpr)
		if err != nil{
			return nil,err
		}
	}
	query := entCli.User.Query().Limit(req.PageSize).Offset(req.PageIndex * req.PageSize)
	if pred != nil{
		query = query.Where(pred)
	}

	return query.All(ctx)
}
```

The ```QueryRequest``` struct will be provided by the client. It will contains query conditions. The ```UserQuery``` will parse the ```QueryRequest``` then config the ent query object to query the db;


# Specification

## Query Language
### Query Language Overview
```
filter = "(name eq 'foo' and age ge 10) or (dept.age eq 50)"
fields = "name,age,lastname,parent.name,parent.age"

This will tell the ent we want to query:
predicateUser.Or(
  predicateUser.And(
    predicateUser.NameEQ("foo"),
    predicateUser.AgeGE(10),
  ),
  predicateUser.HasParentWith(
    predicateUser.AgeEQ(50),
  ),
)

the fields we want to select are:
ent.Client.User.Query().Select(name,age,lastname).WithParent(func (q *UserQuery){
  q.Select(age)
})

```

### Filter
Filter will be parsed and mapped to ent predicates object
#### Filter example
```
((name eq "foo") and (age ge 10)) or (name eq "foo2" and age leq 10)
```
#### Filter Grammar
```
binary_logical_compare -> unary_logical (("or"|"and") unary_logical)*
unary_logical -> "not" unary_logical | group 
group -> ( "(" binary_logical_compare ")" ) | compare
compare -> IDENTIFIER binary_cmp_op (STRING | NUMBER)

---

IDENTIFIER -> ALPHA (ALPHA | DIGIT)*
STRING -> "\"" <any char except "> "\""
NUMBER -> DIGIT + ( "." DIGIT+)?
ALPHA -> "a" ... "z" | "A" ... "Z" | "_"
DIGIT -> "0" ... "9"
FIELD -> IDENTIFIER ("." IDENTIFIER)*
```

### Fields
Fields will tell the ent which fields we want to query
#### Fields Example
query the user's parent id, user's children name, and the user's name
```
user.parent.id , user.children.name, name
```
#### Fields Grammar
```
FIELDS -> FIELD ("," FIELD)*
```
