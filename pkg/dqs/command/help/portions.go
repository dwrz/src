package help

const portions = `The %s command %ss portions to an entry.

It accepts pairs of arguments -- a category name followed by a portion
amount.

For example:
dqs %s fruit 1

This will %s one portion to the Fruit category.

It is possible to specify a half portion, e.g.:
dqs %s vegetables 2.5

Will %s 2½ portions to the Vegetable category.

It is also possible to specify multiple pairs of arguments, e.g.:
dqs %s dairy 1 "Whole Grains" 2

This will %s 1 dairy portion, and 2 whole grain portions to an entry.
Note that category names may be cased in any manner, but names with
a space in them must be quoted, as in the "Whole Grains" example above.

To reduce typing, abbreviations are available for category names, e.g.:
dqs %s d 1 hqb 1 v 2 wg 1.5

Will %s 1 portion of dairy, 1 portion of high quality beverages, 2
portions of vegetables, and 1½ portions of whole grains.

The following abbreviations are available:
`
