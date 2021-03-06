# 5e

A collection of microservices for orchestrating 
a game of Dungeons & Dragons 5th Edition.

Supported controllers:

- character
  - supports objects with data filling out a character sheet
- spell
  - supports atomic spells
  - supports objects linking spells to a character
- player
  - supports JWT auth
- action
  - supports atomic actions
  - supports objects linking actions to a character
- feature
  - supports atomic features
  - supports objects linking features to a character

I.E., `5e character` starts the REST API for `character`

## Dependencies

- [docker][]

## Clean, Build, & Run

```
make up
```

## Usage

Set `${DB_URL}` for a valid postgres DB
Set `${AUTH_KEY}` for the key signing auth tokens

Run just the raw binary `5e` for monolithic mode:

```bash
out/5e

  ____     ___     ____      ____
 |  _ \   ( _ )   |  _ \    | ___|    ___
 | | | |  / _ \/\ | | | |   |___ \   / _ \
 | |_| | | (_>  < | |_| |    ___) | |  __/
 |____/   \___/\/ |____/    |____/   \___|
time="2022-07-16T18:35:03Z" level=info msg="No argument(s) found -- starting up in monolithic mode"
time="2022-07-16T18:35:03Z" level=info msg="Initializing all controllers..."
time="2022-07-16T18:35:03Z" level=info msg="CharacterController: Initialized ✅" 
time="2022-07-16T18:35:04Z" level=info msg="CharacterModel: Seeded ✅"
time="2022-07-16T18:35:04Z" level=info msg="SpellController: Initialized ✅"
time="2022-07-16T18:35:05Z" level=info msg="SpellModel: Seeded ✅"
time="2022-07-16T18:35:05Z" level=info msg="LearnedController: Initialized ✅"
time="2022-07-16T18:35:05Z" level=info msg="LearnedModel: Seeded ✅"
```

Specify the name of a controller(s) for microservice mode:

```bash
out/5e spells characters

  ____     ___     ____      ____
 |  _ \   ( _ )   |  _ \    | ___|    ___
 | | | |  / _ \/\ | | | |   |___ \   / _ \
 | |_| | | (_>  < | |_| |    ___) | |  __/
 |____/   \___/\/ |____/    |____/   \___|
time="2022-07-16T18:35:04Z" level=info msg="Initializing spells controller..."
time="2022-07-16T18:35:04Z" level=info msg="SpellController: Initialized ✅"
time="2022-07-16T18:35:05Z" level=info msg="SpellModel: Seeded ✅"
time="2022-07-16T18:35:05Z" level=info msg="LearnedController: Initialized ✅"
time="2022-07-16T18:35:05Z" level=info msg="LearnedModel: Seeded ✅"
time="2022-07-16T18:35:06Z" level=info msg="Initializing character controller..." 
time="2022-07-16T18:35:06Z" level=info msg="CharacterController: Initialized ✅" 
time="2022-07-16T18:35:06Z" level=info msg="CharacterModel: Seeded ✅"
```

[docker]: https://www.docker.com/products/docker-desktop/
