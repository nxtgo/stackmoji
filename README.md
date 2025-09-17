# stackmoji

a simple go utility to slice any image into multiple square discord emojis that stack perfectly in chat.

## features

* slice an image into **nxm grid** (configurable rows and columns).
* **automatically pads** slices to square cells.
* **resizes each slice to 128×128 px**, ready for discord.
* fully configurable via flags: input, output folder, grid size, and emoji prefix.

## usage

```bash
go run main.go -input myimage.png -output emojis -rows 3 -cols 3 -prefix e
```

### flags

| flag      | description                    | default  |
| --------- | ------------------------------ | -------- |
| `-input`  | path to the input image        | (none)   |
| `-output` | directory to save emoji slices | `emojis` |
| `-rows`   | number of rows in the grid     | 3        |
| `-cols`   | number of columns in the grid  | 3        |
| `-prefix` | prefix for emoji filenames     | `e`      |

## example

if your image is sliced into 3×3 grid with prefix `e`, the output will be:

```
emojis/e0.png
emojis/e1.png
...
emojis/e8.png
```

you can stack them in discord like this:

```
:e0::e1::e2:
:e3::e4::e5:
:e6::e7::e8:
```

# licensed

under CC0 1.0 (public domain) + ip waiver. <3
