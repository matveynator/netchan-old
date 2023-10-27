# Package netchan

The `netchan` package implements type-safe networked channels, allowing two ends of a channel to appear on different computers connected by a network. It accomplishes this by transporting data sent to a channel on one machine so that it can be recovered by a receive operation on a channel of the same type on the other machine.

## Overview

- An **exporter** is responsible for publishing a set of channels by name.
- An **importer** connects to the exporting machine and imports the channels by name.
- After importing the channels, the two machines can use them in the usual way.

## Usage

1. **Exporter**:

   To share channels over the network, an exporter needs to:
   - Publish the desired channels by name.
   - Allow remote importers to connect to its machine.

2. **Importer**:

   To use channels from a remote machine, an importer needs to:
   - Connect to the exporting machine.
   - Import the desired channels by name.
   - Use the imported channels as if they were local.

## Note

It's important to understand that networked channels are not synchronized. They always behave as if they are buffered channels with at least one element. This means that sends and receives can be asynchronous, and the order of operations may not be guaranteed.

Please refer to the package documentation and examples for detailed information on using the `netchan` package effectively.


```
