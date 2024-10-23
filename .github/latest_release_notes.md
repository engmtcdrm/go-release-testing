### Added

- Added dependabot
- Added Action for publishing to PyPi
- Added CHANGELOG.md - Retroactively updated previous version releases

### Changed

- Refactored comments to align with more recent coding practices
- Refactored structure of method definitions to align with more recent coding practices
- Refactored the logic for converting tuples and non-tuples to its own method called `_convert_keys_to_list`. this can be useful for inherited classes that may want to use this function.
