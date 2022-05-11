tell application "Notes"
    set content to body of note 1 of folder "Journal"
    set newContent to content & "<ul><li>This is a new line</li></ul>"
    set body of note 1 of folder "Journal" to newContent
end tell
