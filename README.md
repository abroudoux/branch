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

### Manual

You can paste the binary in your `bin` directory (e.g., on MacOS it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x branch
```

Enjoy!

## 💻 Usage

`branch` allows you to manage your Git branches with various and commons actions. Once you're in interactive mode, either by using `run` / the `-r` flag or simply by running `branch`, you can manage your branches with basics options, like `Copy the name`, `Checkout`, `Delete`, `Create a new branch from` or `Merge`.

```bash
branch
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

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request detailling your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- refactor `refactor/the-great-change`

## 📌 Roadmap

- [ ] Improve error handling

## 📑 License

This project is under [MIT License](LICENSE)
