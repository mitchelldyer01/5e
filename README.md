# 5e

## Dependencies

- [docker][]

## Clean, Build, & Run

```
make up
```

## Usage

Set `${DB_URL}` for a valid postgres DB

Run just the raw binary `5e` for monolithic mode:

```bash
out/5e

svc   |   ____     ___     ____      ____
svc   |  |  _ \   ( _ )   |  _ \    | ___|    ___
svc   |  | | | |  / _ \/\ | | | |   |___ \   / _ \
svc   |  | |_| | | (_>  < | |_| |    ___) | |  __/
svc   |  |____/   \___/\/ |____/    |____/   \___|
svc   | time="2022-07-16T18:35:03Z" level=info msg="No argument(s) found -- starting up in monolithic mode"
svc   | time="2022-07-16T18:35:03Z" level=info msg="Initializing all controllers..."
svc   | time="2022-07-16T18:35:03Z" level=info msg="CharacterController: Initialized ✅" 
svc   | time="2022-07-16T18:35:04Z" level=info msg="CharacterModel: Seeded ✅"
svc   | time="2022-07-16T18:35:04Z" level=info msg="SpellController: Initialized ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="SpellModel: Seeded ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="LearnedController: Initialized ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="LearnedModel: Seeded ✅"
```

Specify the name of a controller(s) for microservice mode:

```bash
out/5e spells characters

svc   |   ____     ___     ____      ____
svc   |  |  _ \   ( _ )   |  _ \    | ___|    ___
svc   |  | | | |  / _ \/\ | | | |   |___ \   / _ \
svc   |  | |_| | | (_>  < | |_| |    ___) | |  __/
svc   |  |____/   \___/\/ |____/    |____/   \___|
svc   | time="2022-07-16T18:35:04Z" level=info msg="Initializing spells controller..."
svc   | time="2022-07-16T18:35:04Z" level=info msg="SpellController: Initialized ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="SpellModel: Seeded ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="LearnedController: Initialized ✅"
svc   | time="2022-07-16T18:35:05Z" level=info msg="LearnedModel: Seeded ✅"
svc   | time="2022-07-16T18:35:06Z" level=info msg="Initializing character controller..." 
svc   | time="2022-07-16T18:35:06Z" level=info msg="CharacterController: Initialized ✅" 
svc   | time="2022-07-16T18:35:06Z" level=info msg="CharacterModel: Seeded ✅"
```

Supported controllers:

- characters
- spells

[docker]: https://www.docker.com/products/docker-desktop/