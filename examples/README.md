# GoPHPServer Examples

This directory contains example applications that demonstrate how to use GoPHPServer.

## Basic Example

The `basic` directory contains a simple PHP application that shows basic PHP functionality with GoPHPServer.

To run this example:

```bash
cd examples/basic
../../build/gophpserver -docroot .
```

Then open your browser to http://localhost:8080/

## API Example

The `api` directory contains a simple REST API implemented in PHP.

To run this example:

```bash
cd examples/api
../../build/gophpserver -docroot .
```

Then you can test the API with:

```bash
# Get all users
curl http://localhost:8080/

# Get a specific user
curl http://localhost:8080/?id=1

# Create a new user (simulated)
curl -X POST http://localhost:8080/
```

## Using These Examples

These examples are meant to demonstrate how GoPHPServer works with different types of PHP applications. You can use them as starting points for your own projects.

To copy an example to your project:

```bash
cp -r examples/basic/* public/
```

Then start the server:

```bash
./run.sh
```