## Description

This library was written for personal use to simplify the ability to make filters using the gorm framework.
The library is based on the principle of structures, so starting from a gorm model, a Filter structure is to be denifed into which tags are inserted to define the type of filter.

## Example

The library requires you to create a struct of type filter of this type

```golang
type UserFilter struct {
	Name      string     `json:"name" filter:"1" searchable:"1"`
	CreatedAt *time.Time `json:"created_at" filter:"2"`               // Filtro per la data di creazione
	UpdatedAt *time.Time `json:"updated_at" filter:"2"`               // Filtro per la data di aggiornamento
	Groups    []uint     `json:"groups" filter:"7" field_filter:"id"` // for query in relation field
	SortBy    string     `json:"sort_by" filter:"4"`                  // Campo su cui ordinare
	SortOrder string     `json:"sort_order" filter:"5"`               // Ordine di ordinamento (asc/desc)
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	Search    string     `json:"search"`
}
```

Parameters `Page, Size` are used for pagination, `Page` identifies the page number and `Size` identifies the number of items per page.

Parameter `Search` is used for full text search, via the `searchable` tag valued at 1 we tell which fields the full text search should be performed on.

the Tag `filter` is used to tell the type of filter we want to apply on that field, the possible values are:

1. 1 = LIKE Where type
2. 2 = Exact match
3. 3 = greater than and equal to
4. 4 = less than and equal to
5. 7 = one of these ( IN clause in the where)

the `field_filter` tag should only be used in attributes that identify a relationship and allows queries to be made using the attribute of the table referenced by the relationship.

the `SortOrder` and `SortBy` attributes are used to define the type of sorting, and the column for which to sort, SortOrder can only be worth `ASC` or `DESC`

in the example folder is an example of using filters.

**Very important the names of the Filter structure parameters must have the same name as the parameters in the Gorm model. The only exception are the additional parameters `SortBy, SortOrder, Page, Size and Search` which are mandatory parameters of the Filter**
