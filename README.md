# rogue

Game with generated dungeons and accessible API. 

The original purpose of it is to serve as an exercise for building a dungeon and a player AI that can complete it. Dungeon generation is based on [munificent](http://journal.stuffwithstuff.com/2014/12/21/rooms-and-mazes/) method and [brad811](https://github.com/brad811/go-dungeon) implementation. Can be played in cmd or with http requests. 

![terminal](cmd.jpg?raw=true "Title")

## How to
In order to run commands, you should have [go](https://go.dev/dl/) installed.

Check [api specs](web/docs/openapi.yml) for more details.
```
# build cli and web binaries into /bin folder
make build

# start http server on 8080 port
make api

# start game in cmd
make shell
```

## Content
🟩 move to finish

🟩 seed generated map

🟧 mobs

🟧 stats

🟧 keys and chests