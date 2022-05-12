on run argv
    tell application "Notes"
        set searchLimit to 10
        set i to 1
        set latestUnsharedNote to null
        repeat with i from 1 to searchLimit
            set currentNote to note i
            if shared of currentNote is false then
                set latestUnsharedNote to currentNote
                exit repeat
            end if
        end repeat

        if latestUnsharedNote is null then
            return "Could not find unshared note in " & searchLimit & " most recent notes"
        end if

        set content to body of latestUnsharedNote
        set newContent to content & item 1 of argv
        set body of latestUnsharedNote to newContent
    end tell
end run
