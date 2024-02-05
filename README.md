GO DATABASE TYPES 1
===================

This is a collection of go types to be used with databases.

## Go documentation

    package dbt1 // import "github.com/harkaitz/go-dbt1"
    
    type CSV string

## Go type CSV

    package dbt1 // import "."
    
    type CSV string
        CSV Holds a list of values separated by commas.
    
    func (c CSV) Contains(s string) bool
    func (c CSV) Get(n int) (r string)
    func (c CSV) Getl() (sl []string)
    func (c CSV) Has(n int) (b bool)
    func (c *CSV) Setl(l []string)
    func (c *CSV) Sets(s string)

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
