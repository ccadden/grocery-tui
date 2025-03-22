# Grocery TUI
Toy Go app to play around with [BubbleTea](https://github.com/charmbracelet/bubbletea) and help with the weekly grocery order.

## Recipe Definitions
### Recipe File Example
```json
{
    "description": "Carrots and Hummus",
    "ingredients": {
        "Carrots": 2,
        "Hummus": 1,
        "Olive Oil (Tbsp)": 1.5
    }
}
```

### Sourcing Recipes
By default, recipe files will be read from `./recipes`.

If a `RECIPE_DIRECTORY` environment variable is specified, the recipes will be read from the defined path if it is valid and accessable.

The `-r` flag can be specified to point to a directory where recipe files are located.

```sh
grocery-tui -r ~/my-best-recipes
```

## TODO
- Refactor into separate model/views
- Make it better looking
- Searching for recipes to add
    - Related, but can we do pagination/grouping?
