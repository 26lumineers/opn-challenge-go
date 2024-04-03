### Add-on ###
- use viper for reading configuration file on config.yaml,so OmisePublicKey and OmiseSecretKey are editable on config.yaml.
- Reuse Card Objects: Instead of creating a new model.Card object for each CSV row.
- use bufio's Scanner,its handy interface that reads from files 
and `\\n` will not be returned in each line.
- use "shopspring/decimal" package for more precise calculate the decimal.
### Output ###
```total received: THB 2,686,403,151.00
total received: THB 2,686,403,151.00 
successfully donated: THB 89,949,179.00
faulty donation: THB 2,596,453,972.00
average per person: THB 2,569,976.54
top donors: 
Mrs. Fatima S Maggot
Ms. Rosa B Gawkroger
Ms. Belladonna M Sackville
```
### Memory profiling ###
![mem_profiling](profile003.svg)\
Took 11MB Memory.

