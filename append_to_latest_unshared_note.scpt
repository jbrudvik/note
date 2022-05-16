on run argv
    -- Get input and format as Notes app expects
    set input to first item of argv
    set formattedInput to ""
    set bulletStart to "b "
    set bulletStartLength to (count bulletStart)
    if input starts with bulletStart then
        if (count input) is bulletStartLength then
            set input to ""
        else
            set input to characters (bulletStartLength + 1) thru -1 of input as string
        end if
        set formattedInput to "<ul><li>" & input & "</li></ul>"
    else
        set formattedInput to "<div>" & input & "</div>"
    end if

    tell application "Notes"
        -- Get latest, unshared note. Abandon search past first 10 notes.
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

        -- Append the formatted input to the note
        set content to body of latestUnsharedNote
        set newContent to content & formattedInput
        set body of latestUnsharedNote to newContent
    end tell
end run
