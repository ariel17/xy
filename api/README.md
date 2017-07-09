# Geolocalization API 

## Test it

```bash
~/xy/api$ go test ./...
```

## Run it

```bash
~/xy/api$ go run api.go 9090  # this is the port on localhost
```

## Resources

### Subjects

* URL: `/subjects`
* Methods:

  + GET: A list of all subjects currently availables.
  + POST: Creates a subject; no data needed by now.
  + OPTIONS: TODO

### Register

* URL: `/register`
* Methods:

  + POST: Catches the PIN parameter in POST and creates an unique ID to use in
    the device.
  + OPTIONS: The list of allowed methods.
