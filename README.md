# branch

🪵 Bash Utility to improve Git Branch

Version : 1.0.0 (WIP)

## 🚀 Installation

### Via Homebrew

Wip 🚧

### Manual

You can copy the function in your shell rc file. You can also create a separate bash file and copy `pm.sh` inside it. You'll need to load it at the beginning of your file, depending your shell (`.bashrc`, `.zshrc`, ...).

```bash
source path/to/your/script.sh
```

Don't forget to resource your shell rc file

```bash
source ~/.zshrc
```

## 💻 Usage

`branch` allows you to manage your git branches thanks to many actions. Once you're in the interractive mode, with the `--run` / `-r` flag or simply by use `branch`, you can manage your branches.

```bash
branch --run
```

From this menu, you can select a branch and choose to Delete it, checkout on it and many more actions which are comming soon.

You can also list all your branchs by using the flag `--list` / `-l`.

```bash
branch --list
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request with the details of your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [ ] Option to create new branch
- [ ] Add Github status on branch (Untracked, Renamed, Modified...)
- [ ] Improve UI
- [ ] Enable installation via Homebrew

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
