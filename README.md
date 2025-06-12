
# Tales of the Valiant Game Tools

A collection of tools and utilities for the Tales of the Valiant tabletop role-playing game.

## Overview

This repository contains various tools to enhance your Tales of the Valiant gaming experience, including character creation, dice rolling, and table lookup utilities.

## Tools

These are all in their infancy, use or peruse at your own risk.

### Character Creator

Create and manage game characters with various lineages and heritages.

- Access character creation API endpoints
- Manage character lineages (Human, Dwarf, Elf, etc.)
- Configure heritage options for each lineage

### Dice Roller

Utility for dice rolling operations within the game.

### Table Lookup

Tools for retrieving game tables and reference information.

## API Usage

The project provides a RESTful API with endpoints for:


- Character creation tools: `/api/v1/character/create`
- Dice rolling operations: `/api/v1/dice/roll` 
- Static table lookup: `/api/v1/table/get` 
- Lineage lookup: `/api/v1/lineages`
- Lineage information: `/api/v1/lineages/:name`
- Heritage lookup: `/api/v1/heritages`
- Heritage information: `/api/v1/heritages/:name`
- Heritage suggestions by lineage: `/api/v1/heritages/lineages`

Example HTTP requests are available in the project's HTTP client files.

## Getting Started

1. Clone the repository
2. Run `devbox shell` from the repository directory
    - This will automatically set up the correct Go version and all dependencies
    - No manual installation of Go or packages is required
3. Use the individual command tools in the `cmd` directory or run the main application

## Project Structure

- `/cmd` - Command-line applications
    - `/create_character` - Character creation utility
    - `/roll` - Dice rolling utility
    - `/get_table` - Table lookup utility
- `/pkg` - Reusable packages and libraries
- `main.go` - Main application entrypoint

## License

Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0)

This is sample code that automates existing processes. You are free to:
- Share — copy and redistribute the material in any medium or format
- Adapt — remix, transform, and build upon the material

Under the following terms:
- Attribution — You must give appropriate credit, provide a link to the license, and indicate if changes were made.
- NonCommercial — You may not use the material for commercial purposes.

The underlying game content and mechanics remain the property of their respective owners and are not covered by this license.

For more details: https://creativecommons.org/licenses/by-nc/4.0/
