Output demo.gif

Require gh
Require figlet
Require lolcrab

Set Shell bash
Set Framerate 30
Set FontSize 24
# Set FontSize 23
Set Width 1280
Set Height 640
Set Padding 40
# Set Padding 80
Set WindowBar Colorful

Hide
Type@0 "cp -a ~/.config/gh ~/.config/gh-bak && find ~/.config/gh/profiles -depth 1 -not -name gabe565 -exec rm -rf {} \+"
Enter
Type@0 "PS1='\n\[\e[38;2;90;86;224m\]> \[\e[0m\]'"
Enter
Type@0 "clear && figlet -f ogre gh-profile | lolcrab --gradient=warm --seed=9 && echo"
Enter
Show
Sleep 500ms

Type "gh auth status"
Sleep 500ms
Enter
Sleep 3s

Type "gh profile create example"
Sleep 500ms
Enter
Sleep 2s

Type "gh auth status"
Sleep 500ms
Enter
Sleep 2s

Type "gh profile ls"
Sleep 500ms
Enter
Sleep 2s

Type "gh profile switch gabe565"
Sleep 500ms
Enter
Sleep 2s

Type "gh profile rm example"
Sleep 500ms
Enter
Sleep 2s

Type "gh profile --help | head -n 14"
Sleep 500ms
Enter
Sleep 10s

Hide
Type@0 "rm -rf ~/.config/gh"
Enter
Type@0 "mv ~/.config/gh-bak ~/.config/gh"
Enter
