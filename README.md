# branch

🌿 Enhance Git Branch Management

## 🚀 Installation

### Via Homebrew

```bash
brew tap abroudoux/tap
```

```bash
brew install abroudoux/tap/branch
```

### Via apt

Wip 🚧

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x branch
```

## 💻 Usage

`branch` allows you to manage your Git branches with various and commons actions. Once you're in interactive mode, either by using `run` / the `-r` flag or simply by running `branch`, you can manage your branches with basics options, like `Copy the name`, `Checkout`, `Delete`, `Create a new branch from` or `Merge`.

```bash
branch run
```

```bash
Choose a branch:

    develop
    feature/ui
> * main
```

```bash
Branch: name

Choose an action:

  Exit
  Delete
  Merge
  Branch
  Checkout
> Name
```

You can also list all your branches using the `--list` / `-l` flag.

```bash
branch --list
```

If you want to personnalize `branch` ui, refer to `config.json`

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request detailling your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [x] Option to create new branch
- [x] Improve UI
- [x] Enable installation via Homebrew
- [ ] Enable installation via apt
- [x] Rewrite script in `Go`
- [x] Add theme configuration

## 📑 License

This project is under [MIT License](LICENSE)
