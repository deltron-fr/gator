## go-gator

A CLI application for aggregating [RSS](https://en.wikipedia.org/wiki/RSS) feeds.
It supports a local multi-user setup and stores data in PostgreSQL.

---

## Requirements

- Postgres version 15+ -- [installation guide](https://www.postgresql.org/download/linux/ubuntu/)
- Golang toolchain(v. 1.24+) 

---

## Installation

```
go install github.com/deltron-fr/gator@latest
```

---

## Quick Start

You will need to setup a `.gatorconfig.json` (the **dot is required**) in your home directory. The file should contain only your Postgres connection string.
Example:

```json
{
  "db_url": "postgres://postgres:password@localhost:5432/gator"
}
```

The format for an actual connection string:

```
protocol://username:password@host:port/database
```


#### Create database

For debian-based systems:

```
# Set a password
sudo passwd postgres
```

```
# Start the service (if not already running)
sudo service postgresql start
```

```
# Enter psql shell
sudo -u postgres psql

# Create a new database called "gator"
CREATE DATABASE gator;

# Connect to the new database:
\c gator

# Set the user password
ALTER USER postgres PASSWORD 'password';

# To exit:
\q
```

---

## Running the program

To get started:

```
gator register <username>
```

Then log in:
```
gator login <username>
```

Add a feed:

```
gator addfeed <feed_name> <feed_url>
```
Example feeds: [Hacker News](https://news.ycombinator.com/rss), [TechCrunch](https://techcrunch.com/feed/)

Follow a feed:

```
gator follow <feed_url>
```

#### Browsing Posts

Open another terminal and run:
```
gator agg 1m
```

After a while, view the aggregated posts::
```
gator browse
```

---

## Future Extensions

- Addition of a search command for posts and feeds
- A TUI that allows for easy viewing of posts
- Addition of bookmarking and liking of posts
- Add sorting and filtering options to the browse command




