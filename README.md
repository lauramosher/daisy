# Daisy

Daisy is a command line tool that bootstraps your day. It initially started as a way for me to run
all my of development startup commands, but morphed into a remote starter kit that integrates with
Slack to enables message posting, set status, and presence.

The idea is to streamline the day. You open your terminal to code, and you can let your team know
what you are doing by setting your status.

## Install

Run the following command. It'll prompt you when necessary.

```
bash <(curl -sL https://raw.githubusercontent.com/lauramosher/daisy/master/install.sh)
```

NOTE: This will always install the latest version.

### Alternatively, from source

```
$ go get github.com/lauramosher/daisy/cmd/day
$ go install github.com/lauramosher/daisy/cmd/day
```

[Authorize Slack](https://daisy-slack.herokuapp.com/) to take advantage of daisy's Slack integration.
daisy will post messages as you, and manipulate your Status and Presence.

#### Requirements

- Go
- $GOPATH set
- $GOPATH/bin in your path

## Usage

`day` will print usage.

### Commands

| Command | Description |
| --- | --- |
| `start` | Start your day. By default, this sets your Slack presence to "auto" (active), status to "Working remotely", and posts a Good Morning message." If you're a Planning Center employee, it will also bootstrap your development machine. |
| `end` | End your day. By default, this sets your Slack presence to "away", status to "EOD", and posts a farewell message. If you're a Planning Center employee, it will also stop your development box. |
| `break` | Go AFK. By default, this will set your Slack presence to "away", status to "Break", and posts a break message (changes depending on morning or afternoon). |
| `lunch` | Go AFK to lunch. By default, this will set your Slack presence to "away" and status to "Lunch". |
| `away` | Go AFK. Generic, you're away. This sets your Slack presence to "away" and status to "Heads down". |
| `return` | Global return from an away state. By default, this will set your Slack presence to "auto" (active) and status to "Working remotely." |


| Flag | Description |
| --- | --- |
| -s, --skip-message | Skip the default posting of message on any command. |
| -m <text>, --message <text> | Specify message to post. This will also post a message for any command, even if there are no default set. |


### Update

To install latest updates:

```
go get -u github.com/lauramosher/daisy/cmd/day
```

## Troubleshooting

> `day` command isn't working!

Ensure Go's bin is in your path:

```
export PATH=$GOPATH/bin:$PATH
```
