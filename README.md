# branch

🪵 Bash Utility to Enhance Git Branch Management

Version : 1.0.0

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

You can copy the function in your shell's RC file. Alternatively, You can create a separate Bash script file and copy `pm.sh` into it. You'll need to load it at the beginning of your shell RC file (e.g., `.bashrc`, `.zshrc`, etc.).

```bash
source path/to/your/script.sh
```

Don't forget to resource your shell RC file:

```bash
source ~/.zshrc
```

## 💻 Usage

`branch` allows you to manage your GÒit branches with various and commons actions. Once you're in interactive mode, either by using the `--run` / `-r` flag or simply by running `branch`, you can manage your branches.

```bash
branch --run
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

- [ ] Option to create new branch
- [ ] Add Git status for branches (Untracked, Renamed, Modified, etc.)
- [ ] Improve UI
- [ ] Enable installation via Homebrew
- [ ] Enable installation via apt

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
