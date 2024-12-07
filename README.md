# branch

🪵 Enhance Git Branch Management

Version : 1.3.2

## 🚀 Installation

### Requirements

You'll need to install `fzf` to use `branch`.\
You can download it using `Hombrew` or `apt`.

```bash
brew install fzf # or apt install fzf
```

Refer to [fzp repository](https://github.com/junegunn/fzf) to more informations.

### Via Homebrew

Wip 🚧

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x branch
```

## 💻 Usage

`branch` allows you to manage your GÒit branches with various and commons actions. Once you're in interactive mode, either by using `run` / the `-r` flag or simply by running `branch`, you can manage your branches.

```bash
branch run
```

From this menu, you can select a branch and choose to delete it, move on it, and perform many more actions (more features coming soon).

You can also list all your branches using the `--list` / `-l` flag.

```bash
branch --list
```

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
- [ ] Add Git status for branches (Untracked, Renamed, Modified, etc.)
- [ ] Improve UI
- [ ] Enable installation via Homebrew
- [ ] Enable installation via apt
- [x] Create binary instead a function and `shell` alias

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
