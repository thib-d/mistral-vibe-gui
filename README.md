# Mistral Vibe

> **Projet non officiel — not affiliated with or endorsed by Mistral AI.**

Application de bureau légère pour [chat.mistral.ai](https://chat.mistral.ai/work) — fenêtre native, session persistante, sans onglets de navigateur.

---

## Français

### Fonctionnalités

- **Fenêtre native** — vit dans le Dock et ⌘+Tab comme n'importe quelle app
- **Session persistante** — connecte-toi une fois, reste connecté
- **Coller (⌘+V) fonctionnel** — contournement de la restriction clipboard de WKWebView
- **Binaire unique ~2 Mo** — pas d'Electron, pas de Chromium embarqué
- **Moteur WebView système** — WKWebKit sur macOS, WebView2 sur Windows, WebKitGTK sur Linux

### Téléchargement

Voir la [page de téléchargement](https://thib-d.github.io/mistral-vibe-gui) ou les [releases GitHub](https://github.com/thib-d/mistral-vibe-gui/releases/latest).

| Plateforme | Architecture | Fichier |
|---|---|---|
| macOS | Apple Silicon (M1/M2/M3/M4) | `mistral-vibe-mac-arm64` |
| macOS | Intel | `mistral-vibe-mac-intel` |
| Windows | x64 | `mistral-vibe-windows.exe` |
| Linux | x64 | `mistral-vibe-linux` |

### Premier lancement sur macOS

macOS bloque les binaires non signés. Pour autoriser :

```bash
xattr -dr com.apple.quarantine mistral-vibe-mac-arm64
./mistral-vibe-mac-arm64
```

Ou clic droit sur le fichier dans le Finder → **Ouvrir** → **Ouvrir** dans la boîte de dialogue.

### Compiler depuis les sources

Prérequis : [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/thib-d/mistral-vibe-gui
cd mistral-vibe-gui
go build -o mistral-vibe .
./mistral-vibe
```

**Linux** — installer la dépendance WebKit :
```bash
# Debian / Ubuntu
sudo apt install libwebkit2gtk-4.0-dev

# Fedora
sudo dnf install webkit2gtk4.0-devel
```

**Windows** — nécessite le [runtime WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) (inclus dans Windows 11, à télécharger pour Windows 10).

---

## English

### Features

- **Native window** — sits in your Dock and ⌘+Tab like any app
- **Persistent session** — log in once, stay logged in
- **Paste (⌘+V / Ctrl+V) works** — WKWebView clipboard restriction bypassed natively
- **~2 MB single binary** — no Electron, no bundled Chromium
- **System WebView engine** — WKWebKit on macOS, WebView2 on Windows, WebKitGTK on Linux

### Download

See the [download page](https://thib-d.github.io/mistral-vibe-gui) or [GitHub Releases](https://github.com/thib-d/mistral-vibe-gui/releases/latest).

| Platform | Architecture | File |
|---|---|---|
| macOS | Apple Silicon (M1/M2/M3/M4) | `mistral-vibe-mac-arm64` |
| macOS | Intel | `mistral-vibe-mac-intel` |
| Windows | x64 | `mistral-vibe-windows.exe` |
| Linux | x64 | `mistral-vibe-linux` |

### First launch on macOS

macOS blocks unsigned binaries. To allow it:

```bash
xattr -dr com.apple.quarantine mistral-vibe-mac-arm64
./mistral-vibe-mac-arm64
```

Or right-click the file in Finder → **Open** → **Open** in the dialog.

### Build from source

Requires [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/thib-d/mistral-vibe-gui
cd mistral-vibe-gui
go build -o mistral-vibe .
./mistral-vibe
```

**Linux** — install WebKit dependency:
```bash
# Debian / Ubuntu
sudo apt install libwebkit2gtk-4.0-dev

# Fedora
sudo dnf install webkit2gtk4.0-devel
```

**Windows** — requires the [WebView2 runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) (included in Windows 11, download for Windows 10).

---

## Disclaimer

This project is **not affiliated with, endorsed by, or associated with Mistral AI** in any way.
It is a personal open-source tool that loads the public website [chat.mistral.ai](https://chat.mistral.ai) in a native window.

## License

MIT
