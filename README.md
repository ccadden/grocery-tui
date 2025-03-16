# Grocery TUI
Toy Go app to play around with [BubbleTea](https://github.com/charmbracelet/bubbletea) and help with the weekly grocery order.

## Recipe Definitions
For now recipies go in `./recipes`. Format is as follows:

```json
{
    "description": "Carrots and Hummus",
    "ingredients": {
        "carrots": 2,
        "hummus": 1
    }
}
```
## TODO
- Some sort of config file/flag for where recipe files are stored
- Refactor into separate model/views
- Make it better looking
- Searching for recipes to add
    - Related, but can we do pagination/grouping?
