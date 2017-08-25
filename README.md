# agg

**agg** is a simple command to compute numerical aggregates (e.g. averages and sums) in the shell. I constantly find myself wanting something like this, and all the [popular approaches](https://stackoverflow.com/questions/9789806/command-line-utility-to-print-statistics-of-numbers-in-linux) are a hassle.

# Installation

You must install [Go](https://golang.org/doc/install) and setup a [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH). You will also want to add `$GOPATH/bin` to your `$PATH`.

Once you have Go, simply run:

```
go get github.com/unixpickle/agg
```

You should now have an `agg` command in `$GOPATH/bin`.

# Example

Suppose we have a file containing some house prices, called `prices.txt`:

```txt
1500
2700
3200
2000
4500
```

We can compute the mean house price like so:

```
$ cat prices.txt | agg mean
2780
```

# Detailed Usage

The stdin to `agg` should be a list of real numbers, separated by whitespace. If your data does not look like this, you can likely use `sed`, `tr`, and `cut` as part of your pipeline (on UNIX).

You can see detailed usage info by running `agg` with no arguments:

```
$ agg
Usage: agg <aggregate type>

Available aggregate types:
  geommean    geometric mean
  mean        arithmetic mean
  sum         basic sum
  variance    variance (with Bessel's correction)
```
