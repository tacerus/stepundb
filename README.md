Simple tool which:

- Reads a list of newline separated, encoded, certificates from stdin.
- Returns a series of JSON objects containing the decoded structure of each certificate.

Example usage:

```
env PGPASSWORD=step psql -Ustep -c 'SELECT nvalue FROM x509_certs;' --csv -t | ./stepundb
```

The `nkey` column is not used as it contains the serial numbers which are contained in the resulting data anyways.
