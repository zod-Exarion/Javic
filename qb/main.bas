Dim groceryList(100) As String
Dim itemCount As Integer
itemCount = 0

Do
    Cls
    Print "=== Grocery List Maker ==="
    Print
    Print "Current List:"
    If itemCount = 0 Then
        Print "  (Empty)"
    Else
        For i = 1 To itemCount
            Print "  "; i; ". "; groceryList(i)
        Next i
    End If
    Print
    Print "Menu:"
    Print "  1. Add item"
    Print "  2. Remove item"
    Print "  3. Quit"
    Print
    Input "Choose an option (1-3): ", choice

    Select Case choice
        Case 1
            If itemCount >= 100 Then
                Print "List is full! Cannot add more items."
            Else
                Input "Enter item to add: ", newItem$
                itemCount = itemCount + 1
                groceryList(itemCount) = newItem$
            End If
        Case 2
            If itemCount = 0 Then
                Print "List is empty. Nothing to remove."
            Else
                Input "Enter item number to remove: ", remIndex
                If remIndex >= 1 And remIndex <= itemCount Then
                    For i = remIndex To itemCount - 1
                        groceryList(i) = groceryList(i + 1)
                    Next i
                    itemCount = itemCount - 1
                    Print "Item removed."
                Else
                    Print "Invalid item number."
                End If
            End If
        Case 3
            Print "Exiting..."
            Exit Do
        Case Else
            Print "Invalid choice. Try again."
    End Select

    Print
    Print "Press any key to continue..."
    Sleep

Loop

Print "succesfully executed the qbasic program"
