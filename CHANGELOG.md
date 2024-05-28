# Changelog for tundra

## Unreleased

### Features

* added `GetObject` method to `Object` struct.

## v0.2.1

### Fixes

* made all single directions one letter \(udio\)

## v0.2.0

### Features

* added up, down, in, and out directions.
* added `Describe()` method to `Location` struct to include the objects there.
* added take and drop command builders to commands package.

## v0.1.1

### Notes

* removed game to [external project](github.com/xaiver2910/tundragame)

### Fixes

* implemented `turnBased.InjectContext()`
* fixed location connection functionality: update context on move

## v0.1.0

### Features

* Basic framework in place for:
  * Command parsing and execution
  * Locations and their connections
  * Player data -- inventory and a spot for additional context
    as the need arises.

* No game content as yet.

