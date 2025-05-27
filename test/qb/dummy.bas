X = 55
Y = 8.2
Print Y, X, 0

e = 1
If Not e Then
    If e = 0 Then
        Print -69 Mod 10
    Else
        Print e
    End If
End If

DECLARE FUNCTION getGrade (score1!, score2!, score3!, score4!, score5!)
Rem THIS PROGRAM USES A USER-DEFINED FUNCTION (FUNCTION PROCEDURE)
Rem TO CALCULATE A STUDENT'S GRADE BASED ON FIVE SCORES WITH
Rem SCORE EACH HAVING A DIFFERENT WEIGHTED VALUE TOWARD THE
Rem STUDENT'S OVERALL AVERAGE
Rem
Rem PROGRAMMER: MIKE WARE
Rem DATE LAST UPDATED: 4-11-01

Input "Enter the student's name: ", stuName$
Input "Enter the student's five test scores (seperated by a comma): ", score1, score2, score3, score4, score5

grade = getGrade(score1, score2, score3, score4, score5)
Print
Print stuName$; ", earned a grade of ("; grade; ")"

End

Function getGrade (score1, score2, score3, score4, score5)

    quizzes = (score1 + score2 + score3) / 300 * 100 * .4
    tests = (score4 + score5) / 200 * 100 * .6

    totalScore = quizzes + tests

    Select Case totalScore
        Case Is > 90
            getGrade = 1
        Case Is > 80
            getGrade = 2
        Case Is > 70
            getGrade = 3
        Case Is > 60
            getGrade = 4
        Case Else
            getGrade = 5
    End Select

End Function


