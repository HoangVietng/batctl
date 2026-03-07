# batctl demo — VHS tape for recording README GIF
# Usage: sudo vhs demo.tape

Output demo.gif

Set Shell "bash"
Set FontSize 14
Set Width 1000
Set Height 600
Set Padding 20
Set Theme "Catppuccin Mocha"
Set TypingSpeed 60ms
Set WindowBar Colorful
Set WindowBarSize 40
Set CursorBlink false

# Launch TUI
Type "sudo batctl"
Enter
Sleep 2s

# Navigate to Start threshold and adjust
Sleep 500ms
Type "l"
Sleep 200ms
Type "l"
Sleep 200ms
Type "l"
Sleep 200ms
Type "l"
Sleep 200ms
Type "l"
Sleep 500ms

# Move to Stop threshold
Type "j"
Sleep 500ms

# Adjust stop threshold
Type "h"
Sleep 200ms
Type "h"
Sleep 200ms
Type "h"
Sleep 200ms
Type "h"
Sleep 200ms
Type "h"
Sleep 500ms

# Move to Presets
Type "j"
Sleep 300ms
Type "j"
Sleep 500ms

# Browse presets
Type "l"
Sleep 400ms
Type "l"
Sleep 400ms
Type "l"
Sleep 400ms

# Go back to balanced
Type "h"
Sleep 300ms
Type "h"
Sleep 500ms

# Select preset
Enter
Sleep 1s

# Apply and save
Type "a"
Sleep 2s

# Navigate to Persistence
Type "j"
Sleep 500ms

# Refresh
Type "r"
Sleep 1s

# Quit
Type "q"
Sleep 500ms
