# TODO

- create common generic stack type and compose parse stack and value stack with it
  - instead of composition, stack return values are type asserted
- write unit tests for machine instruction related methods
  - unit test for every method except setc
- convert machineInstructionType to be slice of bytes
  - done conversion
- parser to receive Writer and output machine instructions with it
  - parser is using Writer to output lines of machine instruction strings
- clean up production list and parse table since epsilon is removed from statement symbol
  - done cleaning up
- clean up debugging log and use logger with level for debug
  - set the otput of the logger to use nullWriter that is associated with /dev/null
- write complete suite of unit tests for parser
  - all but jump verb and atDigits are not tested or implemented at all
- implement jump verb tokens and atDigits digit parsing
- fix: preprocess line ending \r\n to \n
- preprocess (label) and @label/@variable into symbol table
- add actions to production to insert numeric value of @lable/@variable
