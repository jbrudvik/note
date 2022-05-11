on run argv
    tell application "Notes"
        set content to body of note 1 of folder "Journal"
        set newContent to content & item 1 of argv
        set body of note 1 of folder "Journal" to newContent
    end tell
end run
