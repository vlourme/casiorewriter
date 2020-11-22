# Casio OS Rewriter
This project is a PoC showing that we can write our text into the system, allowing us to bypass the examination mode.

If you want to learn more about this subject, I wrote an entire article about this here: https://vlour.me/p/casio-exam-mode/.

**Disclaimer**: Don't use this project to cheat. I highly discourage you from doing this.

## Building
Follow those commands: 
```shell script
# Clone repository
git clone https://github.com/vlourme/casiorewriter

# Navigate to the directory
cd casiorewriter

# Sync libraries
go mod download

# Build
go build main.go
```

## Requirements
- An OS file, I can't give a link (due to DMCA takedown). You can most likely find this on forums. I recommend you using Casio 75+E OS.
- Offsets. Can be found using a hex-editor, search for texts, and note the start offset and end offset.

## How to use
You'll need first to dump the text from offsets :
```shell scipt
./main -input <os_path> -fetch -in <offset_start> -out <offset_end>
```

You'll get a `slice.txt` file, edit this file then you're good to rewrite the OS like this:
```shell script
./main -input <os_path> -in <offset_start> -out <offset_end>
```

Once done, you'll get a `patched_os.bin file.

*Note: You can get help menu by running `./main -h`* 