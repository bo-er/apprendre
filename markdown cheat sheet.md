# Heading 1

```
Markup :  # Heading 1 #

-OR-

Markup :  ============= (below H1 text)
```

##

## Heading 2

```
Markup :  ## Heading 2 ##

-OR-

Markup: --------------- (below H2 text)
```

###

### Heading 3

```
Markup :  ### Heading 3 ###
```

####

#### Heading 4

```
Markup :  #### Heading 4 ####
```

Common text

```
Markup :  Common text
```

_Emphasized text_

```
Markup :  _Emphasized text_ or *Emphasized text*
```

~~Strikethrough text~~

```
Markup :  ~~Strikethrough text~~
```

**Strong text**

```
Markup :  __Strong text__ or **Strong text**
```

**_Strong emphasized text_**

```
Markup :  ___Strong emphasized text___ or ***Strong emphasized text***
```

[Named Link](http://www.google.fr/) and http://www.google.fr/ or http://example.com/

```
Markup :  [Named Link](http://www.google.fr/ "Named link title") and http://www.google.fr/ or <http://example.com/>
```

[heading-1](https://github.com/tchapi/markdown-cheatsheet/blob/master/README.md#heading-1)

```
Markup: [heading-1](#heading-1 "Goto heading-1")
```

Table, like this one :

| First Header | Second Header |
| ------------ | ------------- |
| Content Cell | Content Cell  |
| Content Cell | Content Cell  |

```
First Header  | Second Header
------------- | -------------
Content Cell  | Content Cell
Content Cell  | Content Cell
```

Adding a pipe `|` in a cell :

| First Header | Second Header |
| ------------ | ------------- |
| Content Cell | Content Cell  |
| Content Cell | \|            |

```
First Header  | Second Header
------------- | -------------
Content Cell  | Content Cell
Content Cell  |  \|
```

Left, right and center aligned table

| Left aligned Header | Right aligned Header | Center aligned Header |
| ------------------- | -------------------- | --------------------- |
| Content Cell        | Content Cell         | Content Cell          |
| Content Cell        | Content Cell         | Content Cell          |

````
Left aligned Header | Right aligned Header | Center aligned Header
| :--- | ---: | :---:
Content Cell  | Content Cell | Content Cell
Content Cell  | Content Cell | Content Cell
code()
Markup :  `code()`
    var specificLanguage_code =
    {
        "data": {
            "lookedUpPlatform": 1,
            "query": "Kasabian+Test+Transmission",
            "lookedUpItem": {
                "name": "Test Transmission",
                "artist": "Kasabian",
                "album": "Kasabian",
                "picture": null,
                "link": "http://open.spotify.com/track/5jhJur5n4fasblLSCOcrTp"
            }
        }
    }
Markup : ```javascript
         ```
````

- Bullet list
  - Nested bullet
    - Sub-nested bullet etc
- Bullet list item 2

```
 Markup : * Bullet list
              * Nested bullet
                  * Sub-nested bullet etc
          * Bullet list item 2

-OR-

 Markup : - Bullet list
              - Nested bullet
                  - Sub-nested bullet etc
          - Bullet list item 2
```

1. A numbered list
   1. A nested numbered list
   2. Which is numbered
2. Which is numbered

```
 Markup : 1. A numbered list
              1. A nested numbered list
              2. Which is numbered
          2. Which is numbered
```

- An uncompleted task
- A completed task

```
 Markup : - [ ] An uncompleted task
          - [x] A completed task
```

- An uncompleted task
- A subtask

```
 Markup : - [ ] An uncompleted task
              - [ ] A subtask
```

> Blockquote
>
> > Nested blockquote

```
Markup :  > Blockquote
          >> Nested Blockquote
```

_Horizontal line :_

---

```
Markup :  - - - -
```

_Image with alt :_

[![picture alt](https://camo.githubusercontent.com/2d4a8f835fecf8bee4caa27930ddd7c012ea4bb8023909ee093ee9f5a327ca06/687474703a2f2f7669612e706c616365686f6c6465722e636f6d2f32303078313530)](https://camo.githubusercontent.com/2d4a8f835fecf8bee4caa27930ddd7c012ea4bb8023909ee093ee9f5a327ca06/687474703a2f2f7669612e706c616365686f6c6465722e636f6d2f32303078313530)

```
Markup : ![picture alt](http://via.placeholder.com/200x150 "Title is optional")
```

Foldable text:

<details>
  <summary>Title 1</summary>
  
</details>

<details>
  <summary>Title 2</summary>
  
</details>

```
Markup : <details>
           <summary>Title 1</summary>
           <p>Content 1 Content 1 Content 1 Content 1 Content 1</p>
         </details>
<h3>HTML</h3>
<p> Some HTML code here </p>
```

Link to a specific part of the page:

[Go To TOP](https://github.com/tchapi/markdown-cheatsheet/blob/master/README.md#TOP)

```
Markup : [text goes here](#section_name)
          section_title<a name="section_name"></a>
```

Hotkey:

⌘F

⇧⌘F

```
Markup : <kbd>⌘F</kbd>
```

Hotkey list:

| Key       | Symbol |
| --------- | ------ |
| Option    | ⌥      |
| Control   | ⌃      |
| Command   | ⌘      |
| Shift     | ⇧      |
| Caps Lock | ⇪      |
| Tab       | ⇥      |
| Esc       | ⎋      |
| Power     | ⌽      |
| Return    | ↩      |
| Delete    | ⌫      |
| Up        | ↑      |
| Down      | ↓      |
| Left      | ←      |
| Right     | →      |

Emoji:

❗ Use emoji icons to enhance text. 👍 Look up emoji codes at [emoji-cheat-sheet.com](http://emoji-cheat-sheet.com/)
