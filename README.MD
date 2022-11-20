What is Viper?
Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:

setting defaults
reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
live watching and re-reading of config files (optional)
reading from environment variables
reading from remote config systems (etcd or Consul), and watching changes
reading from command line flags
reading from buffer
setting explicit values
Viper can be thought of as a registry for all of your applications configuration needs.

Why Viper?
When building a modern application, you don’t want to worry about configuration file formats; you want to focus on building awesome software. Viper is here to help with that.

Viper does the following for you:

Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile or Java properties formats.
Provide a mechanism to set default values for your different configuration options.
Provide a mechanism to set override values for options specified through command line flags.
Provide an alias system to easily rename parameters without breaking existing code.
Make it easy to tell the difference between when a user has provided a command line or config file which is the same as the default.
Viper uses the following precedence order. Each item takes precedence over the item below it:

explicit call to Set
flag
env
config
key/value store
default

What is GORM?
GORM is a great ORM library for Go developers. with incredible feature lists and speed, it is considered the standard GO ORM.
Overview
Full-Featured ORM
Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
Hooks (Before/After Create/Save/Update/Delete/Find)
Eager loading with Preload, Joins
Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
Context, Prepared Statment Mode, DryRun Mode
Batch Insert, FindInBatches, Find To Map
SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, NamedArg, Search/Update/Create with SQL Expr
Composite Primary Key
Auto Migrations
Logger
Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
Every feature comes with tests
Developer Friendly