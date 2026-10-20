# Gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```sh
go install github.com/MansurEh99/gator@latest
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable",
  "current_user_name": "username"
}
```

Replace the values with your database connection string.

## Database Setup

Clone this repository so you have access to the schema files. From the repository root, create a PostgreSQL database and run:

```sh
createdb gator
psql -d gator -f sql/schema/001_users.sql
psql -d gator -f sql/schema/002_feeds.sql
psql -d gator -f sql/schema/003_feed_follows.sql
psql -d gator -f sql/schema/004_feed_lastfetched.sql
psql -d gator -f sql/schema/005_posts.sql
```

Set `db_url` in `.gatorconfig.json` to point to that database.

## Usage

Create a new user:

```sh
gator register <name>
```

Add a feed:

```sh
gator addfeed <name> <url>
```

Start the aggregator:

```sh
gator agg 30s
```

View the posts:

```sh
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
