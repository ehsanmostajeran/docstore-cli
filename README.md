# docstore-cli

CLI tool for playing with data stored in [BlobStash](https://github.com/tsileo/blobstash) docstore extension.

## Usage

```
$ docstore-cli
docstore > help
collections => list all collections
collection [collection] => set/display current collection
query [query] => returns query result
insert [json doc] => insert the given doc (JSON expected)
docstore > collection

docstore > collection todo
todo
docstore > insert {"title": "write more docs", "done": false}
{
    "_created_at": 1.448009026e+09,
    "_hash": "c1dec7a857b541757ed31e522588990b9182cf2b5bff88f525d27d09f8335895",
    "_id": "564edd42c1dec7a857b541757ed31e522588990b9182cf2b5bff88f525d27d09f8335895",
    "done": false,
    "title": "write more docs"
}
docstore > insert {"title": "support query update", "done": false}
{
    "_created_at": 1.448009051e+09,
    "_hash": "b437e35298edb5afc4424d77faf6118981a62403c7688d28c5942e92a4810d57",
    "_id": "564edd5bb437e35298edb5afc4424d77faf6118981a62403c7688d28c5942e92a4810d57",
    "done": false,
    "title": "support query update"
}
docstore > query
[
    {
        "_created_at": 1.448009026e+09,
        "_hash": "c1dec7a857b541757ed31e522588990b9182cf2b5bff88f525d27d09f8335895",
        "_id": "564edd42c1dec7a857b541757ed31e522588990b9182cf2b5bff88f525d27d09f8335895",
        "done": false,
        "title": "write more docs"
    },
    {
        "_created_at": 1.448009051e+09,
        "_hash": "b437e35298edb5afc4424d77faf6118981a62403c7688d28c5942e92a4810d57",
        "_id": "564edd5bb437e35298edb5afc4424d77faf6118981a62403c7688d28c5942e92a4810d57",
        "done": false,
        "title": "support query update"
    }
]
docstore > insert {"title": "add example in the README", "done": true}
{
    "_created_at": 1.448009098e+09,
    "_hash": "66cbab8818b4d0f6f45821c7eedfe2e4df636d6e14f25fe08af961b327df2973",
    "_id": "564edd8a66cbab8818b4d0f6f45821c7eedfe2e4df636d6e14f25fe08af961b327df2973",
    "done": true,
    "title": "add example in the README"
}
docstore > query {"done": true}
[
    {
        "_created_at": 1.448009098e+09,
        "_hash": "66cbab8818b4d0f6f45821c7eedfe2e4df636d6e14f25fe08af961b327df2973",
        "_id": "564edd8a66cbab8818b4d0f6f45821c7eedfe2e4df636d6e14f25fe08af961b327df2973",
        "done": true,
        "title": "add example in the README"
    }
]
```
